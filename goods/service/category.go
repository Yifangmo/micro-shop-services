package service

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Yifangmo/micro-shop-services/goods/global"
	"github.com/Yifangmo/micro-shop-services/goods/models"
	"github.com/Yifangmo/micro-shop-services/goods/proto"
)

func (s *GoodsServer) AllCategoryList(context.Context, *emptypb.Empty) (*proto.CategoryListResponse, error) {
	var categories []models.Category
	global.DB.Where(&models.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categories)
	b, _ := json.Marshal(&categories)
	return &proto.CategoryListResponse{JsonData: string(b)}, nil
}

func (s *GoodsServer) GetSubCategory(ctx context.Context, req *proto.GetSubCategoryRequest) (*proto.SubCategoryListResponse, error) {
	resp := proto.SubCategoryListResponse{}

	var category models.Category
	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "the category does not exist")
	}

	resp.Info = &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		Level:          category.Level,
		IsTab:          category.IsTab,
		ParentCategory: category.ParentCategoryID,
	}

	var subCategorys []models.Category
	var subCategoryResponse []*proto.CategoryInfoResponse
	global.DB.Where(&models.Category{ParentCategoryID: req.Id}).Find(&subCategorys)
	for _, subCategory := range subCategorys {
		subCategoryResponse = append(subCategoryResponse, &proto.CategoryInfoResponse{
			Id:             subCategory.ID,
			Name:           subCategory.Name,
			Level:          subCategory.Level,
			IsTab:          subCategory.IsTab,
			ParentCategory: subCategory.ParentCategoryID,
		})
	}

	resp.Subcategory = subCategoryResponse
	return &resp, nil
}

func (s *GoodsServer) CreateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	category := models.Category{}
	cMap := map[string]interface{}{}
	cMap["name"] = req.Name
	cMap["level"] = req.Level
	cMap["is_tab"] = req.IsTab
	if req.Level != 1 {
		cMap["parent_category_id"] = req.ParentCategory
	}
	global.DB.Model(&models.Category{}).Create(cMap)
	return &proto.CategoryInfoResponse{Id: int32(category.ID)}, nil
}

func (s *GoodsServer) DeleteCategory(ctx context.Context, req *proto.CategoryIDRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&models.Category{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "the category does not exist")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	var category models.Category

	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "the category does not exist")
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.ParentCategory != 0 {
		category.ParentCategoryID = req.ParentCategory
	}
	if req.Level != 0 {
		category.Level = req.Level
	}
	if req.IsTab {
		category.IsTab = req.IsTab
	}

	global.DB.Save(&category)

	return &emptypb.Empty{}, nil
}
