package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Yifangmo/micro-shop-services/order/global"
	"github.com/Yifangmo/micro-shop-services/order/models"
	"github.com/Yifangmo/micro-shop-services/order/proto"
)

func (*OrderServer) ShoppingCartItemList(ctx context.Context, req *proto.ShoppingCartRequest) (*proto.ShoppingCartListResponse, error) {
	var shopCarts []models.ShoppingCart
	var resp proto.ShoppingCartListResponse
	if dbres := global.DB.Where(&models.ShoppingCart{UserID: req.Id}).Find(&shopCarts); dbres.Error != nil {
		return nil, dbres.Error
	} else {
		resp.Total = dbres.RowsAffected
	}

	for _, shopCart := range shopCarts {
		resp.Data = append(resp.Data, &proto.ShoppingCartItemResponse{
			Id:          shopCart.ID,
			UserId:      shopCart.UserID,
			GoodsId:     shopCart.GoodsID,
			GoodsNumber: shopCart.GoodNumber,
			Checked:     shopCart.Checked,
		})
	}
	return &resp, nil
}

// 将商品添加到购物车，若购物车中原本没有这件商品，则新建一个记录，否则更新购物车商品数量
func (*OrderServer) CreateShoppingCartItem(ctx context.Context, req *proto.ShoppingCartRequest) (*proto.ShoppingCartItemResponse, error) {
	var shopCart models.ShoppingCart
	if dbres := global.DB.Where(&models.ShoppingCart{GoodsID: req.GoodsId, UserID: req.UserId}).First(&shopCart); dbres.RowsAffected == 1 {
		shopCart.GoodNumber += req.GoodsNumber
	} else {
		shopCart.UserID = req.UserId
		shopCart.GoodsID = req.GoodsId
		shopCart.GoodNumber = req.GoodsNumber
		shopCart.Checked = false
	}
	if dbres := global.DB.Save(&shopCart); dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	return &proto.ShoppingCartItemResponse{Id: shopCart.ID}, nil
}

// 要保证购物车商品数量大于0
func (*OrderServer) UpdateShoppingCartItem(ctx context.Context, req *proto.ShoppingCartRequest) (*emptypb.Empty, error) {
	if req.GoodsNumber < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "the goods number less than 1")
	}
	var shopCart models.ShoppingCart
	if dbres := global.DB.Where("user_id=? and goods_id=?", req.UserId, req.GoodsId).First(&shopCart); dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "the Shopping cart record does not exist")
	}
	if req.Checked != nil {
		shopCart.Checked = req.Checked.Value
	}
	shopCart.GoodNumber = req.GoodsNumber
	if dbres := global.DB.Save(&shopCart); dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	return &emptypb.Empty{}, nil
}

func (*OrderServer) DeleteShoppingCartItem(ctx context.Context, req *proto.ShoppingCartRequest) (*emptypb.Empty, error) {
	if dbres := global.DB.Where("user_id=? and goods_id=?", req.UserId, req.GoodsId).Delete(&models.ShoppingCart{}); dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "the Shopping cart record does not exist")
	}
	return &emptypb.Empty{}, nil
}
