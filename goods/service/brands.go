package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Yifangmo/micro-shop-services/goods/global"
	"github.com/Yifangmo/micro-shop-services/goods/models"
	"github.com/Yifangmo/micro-shop-services/goods/proto"
	"github.com/Yifangmo/micro-shop-services/goods/utils"
)

func (s *GoodsServer) BrandList(ctx context.Context, req *proto.PageInfo) (*proto.BrandListResponse, error) {
	resp := proto.BrandListResponse{}
	var brands []models.Brand
	dbres := global.DB.Scopes(utils.Paginate(int(req.PageNumber), int(req.PageSize))).Find(&brands)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	var total int64
	dbres = global.DB.Model(&models.Brand{}).Count(&total)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	resp.Total = total

	var brandsProto []*proto.BrandInfoResponse
	for _, brand := range brands {
		brandsProto = append(brandsProto, &proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		})
	}
	resp.Data = brandsProto
	return &resp, nil
}

func (s *GoodsServer) CreateBrand(ctx context.Context, req *proto.BrandInfoRequest) (*proto.BrandInfoResponse, error) {
	dbres := global.DB.Where("name=?", req.Name).First(&models.Brand{})
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	if dbres.RowsAffected > 0 {
		return nil, status.Error(codes.InvalidArgument, "the brand name does exist")
	}

	brand := &models.Brand{
		Name: req.Name,
		Logo: req.Logo,
	}
	global.DB.Save(brand)

	return &proto.BrandInfoResponse{Id: brand.ID}, nil
}

func (s *GoodsServer) DeleteBrand(ctx context.Context, req *proto.BrandIDRequest) (*emptypb.Empty, error) {
	dbres := global.DB.Delete(&models.Brand{}, req.Id)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	if dbres.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "the brand does not exist")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateBrand(ctx context.Context, req *proto.BrandInfoRequest) (*emptypb.Empty, error) {
	brands := models.Brand{
		BaseModel: models.BaseModel{ID: req.Id},
	}
	dbres := global.DB.Take(&brands)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	if dbres.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, "the brand does not exist")
	}
	if req.Name != "" {
		brands.Name = req.Name
	}
	if req.Logo != "" {
		brands.Logo = req.Logo
	}
	dbres = global.DB.Save(&brands)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	return &emptypb.Empty{}, nil
}
