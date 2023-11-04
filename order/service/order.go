package service

import (
	"context"
	"encoding/json"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Yifangmo/micro-shop-services/order/global"
	"github.com/Yifangmo/micro-shop-services/order/models"
	"github.com/Yifangmo/micro-shop-services/order/proto"
	"github.com/Yifangmo/micro-shop-services/order/utils"
)

type OrderServer struct {
	proto.UnimplementedOrderServer
}

/*
新建订单流程：
 1. 从购物车中获取到选中的商品
 2. 访问商品服务查询商品价格
 3. 访问库存服务扣减库存
 4. 向订单基本信息表和订单商品信息表添加记录
 5. 从购物车中删除已购买的记录
*/
func (*OrderServer) CreateOrder(ctx context.Context, req *proto.OrderRequest) (*proto.OrderInfoResponse, error) {
	orderListener := CreateOrderListener{Ctx: ctx}
	tp, err := rocketmq.NewTransactionProducer(
		&orderListener,
		producer.WithNameServer(global.ServerConfig.RocketMQConfig.NameServer),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "NewTransactionProducer error: %v", err)
	}
	if err = tp.Start(); err != nil {
		return nil, status.Errorf(codes.Internal, "Start TransactionProducer error: %v", err)
	}
	defer tp.Shutdown()
	order := models.Order{
		UserID:           req.UserId,
		OrderSN:          utils.GenerateOrderSN(req.UserId),
		Status:           models.ORDER_STATUS_WAITING_PAY,
		ConsigneeAddress: req.ConsigneeAddress,
		ConsigneeName:    req.ConsigneeName,
		ConsigneeMobile:  req.ConsigneeMobile,
		Remark:           req.Remark,
	}
	orderMsg, _ := json.Marshal(order)
	_, err = tp.SendMessageInTransaction(ctx, primitive.NewMessage("order_reback", orderMsg))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "SendMessageInTransaction error: %v", err)
	}
	if orderListener.Code != codes.OK {
		return nil, status.Error(orderListener.Code, orderListener.ErrMsg)
	}
	return &proto.OrderInfoResponse{
		Id:      orderListener.OrderID,
		OrderSn: order.OrderSN,
		Status:  proto.OrderStatus_ORDER_STATUS_WAITING_PAY,
		Amount:  orderListener.OrderAmount,
	}, nil
}

func (*OrderServer) OrderList(ctx context.Context, req *proto.OrderListRequest) (*proto.OrderListResponse, error) {
	var (
		orders []models.Order
		resp   proto.OrderListResponse
	)
	global.DB.Model(&models.Order{}).Where(&models.Order{UserID: req.UserId}).Count(&resp.Total)
	global.DB.Model(&models.Order{}).
		Scopes(utils.Paginate(int(req.PageInfo.PageNumber), int(req.PageInfo.PageSize))).
		Where(&models.Order{UserID: req.UserId}).
		Find(&orders)
	for _, order := range orders {
		resp.Data = append(resp.Data, &proto.OrderInfoResponse{
			Id:               order.ID,
			UserId:           order.UserID,
			OrderSn:          order.OrderSN,
			PayType:          proto.PayType(order.PayType),
			Status:           proto.OrderStatus(order.Status),
			Remark:           order.Remark,
			Amount:           order.Amount,
			ConsigneeAddress: order.ConsigneeAddress,
			ConsigneeName:    order.ConsigneeName,
			ConsigneeMobile:  order.ConsigneeMobile,
			CreateAt:         order.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &resp, nil
}

func (*OrderServer) OrderDetail(ctx context.Context, req *proto.OrderRequest) (*proto.OrderDetailResponse, error) {
	var (
		order      models.Order
		orderGoods []models.OrderGoods
	)
	if dbres := global.DB.Where(&models.Order{BaseModel: models.BaseModel{ID: req.Id}, UserID: req.UserId}).First(&order); dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "the order does not exist")
	}
	if dbres := global.DB.Where(&models.OrderGoods{OrderID: order.ID}).Find(&orderGoods); dbres.Error != nil {
		return nil, dbres.Error
	}
	resp := proto.OrderDetailResponse{
		OrderInfo: &proto.OrderInfoResponse{
			Id:               order.ID,
			UserId:           order.UserID,
			OrderSn:          order.OrderSN,
			PayType:          proto.PayType(order.PayType),
			Status:           proto.OrderStatus(order.Status),
			Remark:           order.Remark,
			Amount:           order.Amount,
			ConsigneeAddress: order.ConsigneeAddress,
			ConsigneeName:    order.ConsigneeName,
			ConsigneeMobile:  order.ConsigneeMobile,
		},
	}
	for _, item := range orderGoods {
		resp.OrderGoods = append(resp.OrderGoods, &proto.OrderGoods{
			GoodsId:     item.GoodsID,
			GoodsName:   item.GoodsName,
			GoodsPrice:  item.GoodsPrice,
			GoodsImage:  item.GoodsImage,
			GoodsNumber: item.GoodsNumber,
		})
	}

	return &resp, nil
}

func (*OrderServer) UpdateOrderStatus(ctx context.Context, req *proto.UpdateOrderStatusRequest) (*emptypb.Empty, error) {
	if dbres := global.DB.Model(&models.Order{}).Where("order_sn = ?", req.OrderSn).Update("status", req.Status); dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "the order does not exist")
	}
	return &emptypb.Empty{}, nil
}
