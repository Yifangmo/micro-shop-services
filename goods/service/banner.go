package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Yifangmo/micro-shop-services/goods/global"
	"github.com/Yifangmo/micro-shop-services/goods/models"
	"github.com/Yifangmo/micro-shop-services/goods/proto"
)

func (s *GoodsServer) BannerList(ctx context.Context, req *emptypb.Empty) (*proto.BannerListResponse, error) {
	resp := proto.BannerListResponse{}

	var banners []models.Banner
	dbres := global.DB.Find(&banners)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	resp.Total = dbres.RowsAffected

	var bannersProto []*proto.BannerResponse
	for _, banner := range banners {
		bannersProto = append(bannersProto, &proto.BannerResponse{
			Id:    banner.ID,
			Image: banner.Image,
			Index: banner.Index,
			Url:   banner.Url,
		})
	}

	resp.Data = bannersProto

	return &resp, nil
}

func (s *GoodsServer) CreateBanner(ctx context.Context, req *proto.BannerRequest) (*proto.BannerResponse, error) {
	banner := &models.Banner{
		Image: req.Image,
		Index: req.Index,
		Url:   req.Url,
	}
	dbres := global.DB.Save(banner)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	return &proto.BannerResponse{Id: banner.ID}, nil
}

func (s *GoodsServer) DeleteBanner(ctx context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	dbres := global.DB.Delete(&models.Banner{}, req.Id)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	if dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "the banner does not exist, id: %v", req.Id)
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateBanner(ctx context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	var banner models.Banner
	dbres := global.DB.Take(&banner, req.Id)
	if dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "the banner does not exist, id: %v", req.Id)
	}

	if req.Url != "" {
		banner.Url = req.Url
	}
	if req.Image != "" {
		banner.Image = req.Image
	}
	if req.Index != 0 {
		banner.Index = req.Index
	}

	dbres = global.DB.Save(&banner)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	return &emptypb.Empty{}, nil
}
