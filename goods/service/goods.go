package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Yifangmo/micro-shop-services/goods/global"
	"github.com/Yifangmo/micro-shop-services/goods/models"
	"github.com/Yifangmo/micro-shop-services/goods/proto"
)

type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

func (s *GoodsServer) GoodsListQuery(ctx context.Context, req *proto.GoodsListQueryRequest) (*proto.GoodsListResponse, error) {
	resp := &proto.GoodsListResponse{}
	esQuery := elastic.NewBoolQuery()

	if req.KeyWord != nil {
		esQuery = esQuery.Must(elastic.NewMultiMatchQuery(req.KeyWord.Value, "name", "brief"))
	}
	if req.PriceMin != nil {
		esQuery = esQuery.Filter(elastic.NewRangeQuery("shop_price").Gte(req.PriceMin.Value))
	}
	if req.PriceMax != nil {
		esQuery = esQuery.Filter(elastic.NewRangeQuery("shop_price").Lte(req.PriceMax.Value))
	}
	if req.IsHot != nil {
		esQuery = esQuery.Filter(elastic.NewTermQuery("is_hot", req.IsHot.Value))
	}
	if req.IsNew != nil {
		esQuery = esQuery.Filter(elastic.NewTermQuery("is_new", req.IsNew.Value))
	}
	if req.Brand != nil {
		esQuery = esQuery.Filter(elastic.NewTermQuery("brand_id", req.Brand.Value))
	}

	var subQuery string
	categoryIds := make([]interface{}, 0)
	if req.TopCategory != nil {
		var category models.Category
		if dbres := global.DB.First(&category, req.TopCategory); dbres.RowsAffected == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "the category does not exist")
		}

		if category.Level == 1 {
			subQuery = fmt.Sprintf("select id from category where parent_category_id in (select id from category WHERE parent_category_id=%d)", req.TopCategory.Value)
		} else if category.Level == 2 {
			subQuery = fmt.Sprintf("select id from category WHERE parent_category_id=%d", req.TopCategory.Value)
		} else if category.Level == 3 {
			subQuery = fmt.Sprintf("select id from category WHERE id=%d", req.TopCategory.Value)
		}

		var ids []int32
		global.DB.Model(models.Category{}).Raw(subQuery).Scan(&ids)
		for _, id := range ids {
			categoryIds = append(categoryIds, id)
		}
		esQuery = esQuery.Filter(elastic.NewTermsQuery("category_id", categoryIds...))
	}

	if req.PageInfo.PageNumber == 0 {
		req.PageInfo.PageNumber = 1
	}

	if req.PageInfo.PageSize > 100 {
		req.PageInfo.PageSize = 100
	} else if req.PageInfo.PageSize <= 0 {
		req.PageInfo.PageSize = 10
	}

	esres, err := global.ESClient.Search().
		Index(models.ESGoods{}.GetIndexName()).
		Source("id").
		Query(esQuery).
		From(int(req.PageInfo.PageNumber)).
		Size(int(req.PageInfo.PageSize)).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	resp.Total = esres.Hits.TotalHits.Value
	goodsID := make([]int32, resp.Total)
	for i, value := range esres.Hits.Hits {
		if err := json.Unmarshal(value.Source, &goodsID[i]); err != nil {
			return nil, status.Errorf(codes.Internal, "Unmarshal error: %v", err)
		}
	}

	var goods []models.Goods
	dbres := global.DB.Model(models.Goods{}).Preload("Category").Preload("Brands").Find(&goods, goodsID)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}

	for _, good := range goods {
		resp.Data = append(resp.Data, good.ToProto())
	}

	return resp, nil
}

func (s *GoodsServer) GetGoodsByIDs(ctx context.Context, req *proto.GoodsIDsRequest) (*proto.GoodsMapResponse, error) {
	resp := &proto.GoodsMapResponse{
		GoodsMap: make(map[int32]*proto.GoodsInfoResponse),
	}
	var goods []models.Goods
	dbres := global.DB.Where("id in (?)", req.Ids).Find(&goods)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v",  dbres.Error)
	}
	for _, item := range goods {
		resp.GoodsMap[item.ID] = item.ToProto()
	}
	resp.Total = dbres.RowsAffected
	return resp, nil
}

func (s *GoodsServer) GetGoodsByID(ctx context.Context, req *proto.GoodsIDRequest) (*proto.GoodsInfoResponse, error) {
	var goods models.Goods

	if dbres := global.DB.Preload("Category").Preload("Brands").First(&goods, req.Id); dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "the goods does not exist")
	}
	return goods.ToProto(), nil
}

func (s *GoodsServer) CreateGoods(ctx context.Context, req *proto.GoodsInfoRequest) (*proto.GoodsInfoResponse, error) {
	var category models.Category
	if dbres := global.DB.First(&category, req.CategoryId); dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "the category does not exist")
	}

	var brand models.Brand
	if dbres := global.DB.First(&brand, req.BrandId); dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "the brand does not exist")
	}
	goods := models.Goods{
		Brand:          brand,
		BrandID:        brand.ID,
		Category:       category,
		CategoryID:     category.ID,
		Name:           req.Name,
		SN:             req.Sn,
		MarketPrice:    req.MarketPrice,
		ShopPrice:      req.ShopPrice,
		Brief:          req.Brief,
		IsFreeShipping: req.IsFreeShipping,
		Images:         req.Images,
		DescImages:     req.DescImages,
		PreviewImage:   req.PreviewImage,
		IsNew:          req.IsNew,
		IsHot:          req.IsHot,
		IsOnSale:       req.IsOnSale,
	}

	dbres := global.DB.Save(&goods)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}

	return &proto.GoodsInfoResponse{
		Id: goods.ID,
	}, nil
}

func (s *GoodsServer) DeleteGoods(ctx context.Context, req *proto.GoodsIDRequest) (*emptypb.Empty, error) {
	if dbres := global.DB.Delete(&models.Goods{BaseModel: models.BaseModel{ID: req.Id}}, req.Id); dbres.Error != nil {
		return nil, status.Errorf(codes.NotFound, "the goods does not exist")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateGoods(ctx context.Context, req *proto.GoodsInfoRequest) (*emptypb.Empty, error) {
	var goods models.Goods
	dbres := global.DB.Take(&goods, req.Id)
	if dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "the goods does not exist")
	}

	var category models.Category
	dbres = global.DB.Take(&category, req.CategoryId)
	if dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "the category does not exist")
	}

	var brand models.Brand
	dbres = global.DB.Take(&brand, req.BrandId)
	if dbres.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "the brand does not exist")
	}

	goods.Brand = brand
	goods.BrandID = brand.ID
	goods.Category = category
	goods.CategoryID = category.ID
	goods.Name = req.Name
	goods.Brief = req.Brief
	goods.SN = req.Sn
	goods.PreviewImage = req.PreviewImage
	goods.Images = req.Images
	goods.DescImages = req.DescImages
	goods.MarketPrice = req.MarketPrice
	goods.ShopPrice = req.ShopPrice
	goods.IsFreeShipping = req.IsFreeShipping
	goods.IsNew = req.IsNew
	goods.IsHot = req.IsHot
	goods.IsOnSale = req.IsOnSale

	dbres = global.DB.Save(&goods)
	if dbres.Error != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", dbres.Error)
	}
	return &emptypb.Empty{}, nil
}
