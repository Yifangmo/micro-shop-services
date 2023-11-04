package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Yifangmo/micro-shop-services/user/global"
	"github.com/Yifangmo/micro-shop-services/user/models"
	"github.com/Yifangmo/micro-shop-services/user/proto"
)

func (*UserServer) GetUserFavList(ctx context.Context, req *proto.UserFavRequest) (*proto.UserFavListResponse, error) {
	var resp proto.UserFavListResponse
	var userFav []models.UserFav
	var userFavProto []*proto.UserFavResponse
	dbres := global.DB.Where(&models.UserFav{UserID: req.UserId}).Find(&userFav)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	resp.Total = dbres.RowsAffected

	for _, userFav := range userFav {
		userFavProto = append(userFavProto, &proto.UserFavResponse{
			UserId:  userFav.UserID,
			GoodsId: userFav.GoodsID,
		})
	}

	resp.Data = userFavProto

	return &resp, nil
}

func (*UserServer) AddUserFav(ctx context.Context, req *proto.UserFavRequest) (*emptypb.Empty, error) {
	userFav := models.UserFav{
		UserID:  req.UserId,
		GoodsID: req.GoodsId,
	}
	dbres := global.DB.Create(&userFav)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}

	return &emptypb.Empty{}, nil
}

func (*UserServer) DeleteUserFav(ctx context.Context, req *proto.UserFavRequest) (*emptypb.Empty, error) {
	dbres := global.DB.Unscoped().Where("goods=? and user=?", req.GoodsId, req.UserId).Delete(&models.UserFav{})
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	if dbres.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "the user favor record does not exist")
	}
	return &emptypb.Empty{}, nil
}

func (*UserServer) GetUserFavDetail(ctx context.Context, req *proto.UserFavRequest) (*emptypb.Empty, error) {
	var userfav models.UserFav
	dbres := global.DB.Where("goods=? and user=?", req.GoodsId, req.UserId).Find(&userfav)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	if dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "the user favor record does not exist")
	}
	return &emptypb.Empty{}, nil
}
