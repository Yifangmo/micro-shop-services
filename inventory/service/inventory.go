package service

import (
	"context"

	"github.com/Yifangmo/micro-shop-services/inventory/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"

	"github.com/Yifangmo/micro-shop-services/inventory/global"
	"github.com/Yifangmo/micro-shop-services/inventory/proto"
)

type InventoryServer struct {
	proto.UnimplementedInventoryServer
}

func (*InventoryServer) SetInventory(ctx context.Context, req *proto.GoodsInventory) (*emptypb.Empty, error) {
	inventory := models.Inventory{
		GoodsID: req.GoodsId,
		Stock:   req.Num,
	}
	if err := global.DB.Save(&inventory).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func (*InventoryServer) GetInventory(ctx context.Context, req *proto.GoodsInventory) (*proto.GoodsInventory, error) {
	var res models.Inventory
	if dbres := global.DB.Where(&models.Inventory{GoodsID: req.GoodsId}).Take(&res); dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "the goods inventory not exist")
	}
	return &proto.GoodsInventory{
		GoodsId: res.GoodsID,
		Num:     res.Stock,
	}, nil
}

// 扣减库存
func (*InventoryServer) Sell(ctx context.Context, req *proto.OrderInfo) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, global.DB.Transaction(
		func(tx *gorm.DB) error {
			orderDetail := models.OrderDetail{
				OrderSn: req.OrderSn,
				Status:  models.SOLD,
			}
			for _, require := range req.GoodsNumDelta {
				if require.Num <= 0 {
					return status.Errorf(codes.InvalidArgument, "goods[%d] sold num less than 1: %d", require.GoodsId, require.Num)
				}
				orderDetail.Detail = append(orderDetail.Detail, models.GoodsDetail{
					GoodsID: require.GoodsId,
					Num:     require.Num,
				})
				// 注意检查库存是否充足和扣减库存操作必须放在一个SQL语句中来保证原子性（或者使用select...for update 来显式加写锁）
				// 同时由Mysql的写锁机制来保证各并发事务相互隔离
				dbres := tx.Model(&models.Inventory{}).
					Where("goods_id = ?", require.GoodsId).
					Where("stock >= ?", require.Num).
					Update("stock", gorm.Expr("stock-?", require.Num))
				if dbres.Error != nil {
					return status.Errorf(codes.Internal, "update Inventory db error: %v", dbres.Error)
				}
				if dbres.RowsAffected == 0 {
					return status.Errorf(codes.ResourceExhausted, "goods[%d] stock insufficient", require.GoodsId)
				}
			}
			// 为该订单号创建 OrderDetail 记录，用于后来可能的库存归还操作
			if dbres := tx.Create(&orderDetail); dbres.Error != nil {
				return status.Errorf(codes.Internal, "create orderDetail db error: %v", dbres.Error)
			}
			return nil
		},
	)
}

func (*InventoryServer) GiveBack(ctx context.Context, req *proto.OrderInfo) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, global.DB.Transaction(
		func(tx *gorm.DB) error {
			for _, delta := range req.GoodsNumDelta {
				dbres := tx.Model(&models.Inventory{}).Where("goods_id", delta.GoodsId).Update("stock", gorm.Expr("stock+?", delta.Num))
				if dbres.Error != nil {
					return status.Errorf(codes.Internal, "update Inventory db error: %v", dbres.Error)
				}
				if dbres.RowsAffected == 0 {
					return status.Errorf(codes.InvalidArgument, "goods[%d] inventory not exist", delta.GoodsId)
				}
			}
			return nil
		},
	)
}
