package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Yifangmo/micro-shop-services/common"
	"github.com/Yifangmo/micro-shop-services/goods/global"
	"github.com/Yifangmo/micro-shop-services/goods/models"
	"github.com/Yifangmo/micro-shop-services/goods/proto"
	"github.com/Yifangmo/micro-shop-services/goods/utils"
)

func (s *GoodsServer) CategoryBrandList(ctx context.Context, req *common.PageInfo) (*proto.CategoryBrandListResponse, error) {
	var categoryBrands []models.GoodsCategoryBrand
	resp := proto.CategoryBrandListResponse{}

	var total int64
	global.DB.Model(&models.GoodsCategoryBrand{}).Count(&total)
	resp.Total = total

	global.DB.Preload("Category").Preload("Brands").Scopes(utils.Paginate(int(req.PageNumber), int(req.PageSize))).Find(&categoryBrands)

	var categoryResponses []*proto.CategoryBrandResponse
	for _, categoryBrand := range categoryBrands {
		categoryResponses = append(categoryResponses, &proto.CategoryBrandResponse{
			Category: &proto.CategoryInfoResponse{
				Id:             categoryBrand.Category.ID,
				Name:           categoryBrand.Category.Name,
				Level:          categoryBrand.Category.Level,
				IsTab:          categoryBrand.Category.IsTab,
				ParentCategory: categoryBrand.Category.ParentCategoryID,
			},
			Brand: &proto.BrandInfoResponse{
				Id:   categoryBrand.Brand.ID,
				Name: categoryBrand.Brand.Name,
				Logo: categoryBrand.Brand.Logo,
			},
		})
	}

	resp.Data = categoryResponses
	return &resp, nil
}

func (s *GoodsServer) GetCategoryBrandList(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.BrandListResponse, error) {
	resp := proto.BrandListResponse{}

	var category models.Category
	if dbres := global.DB.Find(&category, req.Id).First(&category); dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "the category does not exist")
	}

	var categoryBrands []models.GoodsCategoryBrand
	if dbres := global.DB.Preload("Brands").Where(&models.GoodsCategoryBrand{CategoryID: req.Id}).Find(&categoryBrands); dbres.RowsAffected > 0 {
		resp.Total = dbres.RowsAffected
	}

	var brandInfoResponses []*proto.BrandInfoResponse
	for _, categoryBrand := range categoryBrands {
		brandInfoResponses = append(brandInfoResponses, &proto.BrandInfoResponse{
			Id:   categoryBrand.Brand.ID,
			Name: categoryBrand.Brand.Name,
			Logo: categoryBrand.Brand.Logo,
		})
	}

	resp.Data = brandInfoResponses

	return &resp, nil
}

func (s *GoodsServer) CreateCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*proto.CategoryBrandResponse, error) {
	var category models.Category
	if dbres := global.DB.First(&category, req.CategoryId); dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "the category does not exist")
	}

	var brand models.Brand
	if dbres := global.DB.First(&brand, req.BrandId); dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	categoryBrand := models.GoodsCategoryBrand{
		CategoryID: req.CategoryId,
		BrandID:    req.BrandId,
	}

	global.DB.Save(&categoryBrand)
	return &proto.CategoryBrandResponse{Id: categoryBrand.ID}, nil
}

func (s *GoodsServer) DeleteCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
	if dbres := global.DB.Delete(&models.GoodsCategoryBrand{}, req.Id); dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "the category does not exist")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
	var categoryBrand models.GoodsCategoryBrand

	if dbres := global.DB.First(&categoryBrand, req.Id); dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "the category does not exist")
	}

	var category models.Category
	if dbres := global.DB.First(&category, req.CategoryId); dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "the category does not exist")
	}

	var brand models.Brand
	if dbres := global.DB.First(&brand, req.BrandId); dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "the goods does not exist")
	}
	categoryBrand.CategoryID = req.CategoryId
	categoryBrand.BrandID = req.BrandId

	global.DB.Save(&categoryBrand)

	return &emptypb.Empty{}, nil
}
