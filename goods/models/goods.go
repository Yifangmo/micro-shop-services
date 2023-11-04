package models

import (
	"context"
	"strconv"

	"gorm.io/gorm"

	"github.com/Yifangmo/micro-shop-services/goods/global"
	"github.com/Yifangmo/micro-shop-services/goods/proto"
)

type Goods struct {
	BaseModel

	CategoryID int32 `gorm:"not null;comment:商品分类ID" json:"category_id"`
	Category   Category
	BrandID    int32 `gorm:"not null;comment:商品品牌ID" json:"brand_id"`
	Brand      Brand

	Name         string  `gorm:"type:varchar(100);not null;comment:商品名称" json:"name"`
	Brief        string  `gorm:"type:varchar(255);not null;comment:商品描述" json:"brief"`
	SN           string  `gorm:"type:varchar(50);not null;comment:商品SN码" json:"sn"`
	MarketPrice  float64 `gorm:"type:decimal(10,2);not null;comment:商品当前价格" json:"market_price"`
	ShopPrice    float64 `gorm:"type:decimal(10,2);not null;comment:商品原价格" json:"shop_price"`
	Images       Strings `gorm:"type:json;not null;comment:商品图片" json:"images"`
	DescImages   Strings `gorm:"type:json;not null;comment:商品详情图片" json:"desc_images"`
	PreviewImage string  `gorm:"type:varchar(255);not null;comment:商品预览图片" json:"preview_image"`

	ClickNum int32 `gorm:"default:0;not null;comment:点击数量" json:"click_num"`
	SoldNum  int32 `gorm:"default:0;not null;comment:销售数量" json:"sold_num"`
	FavNum   int32 `gorm:"default:0;not null;comment:收藏数量" json:"fav_num"`

	IsFreeShipping bool `gorm:"not null;default:false;comment:是否免运费" json:"is_free_shipping"`
	IsOnSale       bool `gorm:"not null;default:false;comment:是否在销售" json:"is_on_sale"`
	IsHot          bool `gorm:"not null;default:false;comment:是否热销" json:"is_hot"`
	IsNew          bool `gorm:"not null;default:false;comment:是否新品" json:"is_new"`
}

func (g *Goods) AfterCreate(tx *gorm.DB) (err error) {
	esGoods := ESGoods{
		ID:             g.ID,
		CategoryID:     g.CategoryID,
		BrandID:        g.BrandID,
		Name:           g.Name,
		Brief:          g.Brief,
		MarketPrice:    g.MarketPrice,
		ShopPrice:      g.ShopPrice,
		ClickNum:       g.ClickNum,
		SoldNum:        g.SoldNum,
		FavNum:         g.FavNum,
		IsOnSale:       g.IsOnSale,
		IsFreeShipping: g.IsFreeShipping,
		IsNew:          g.IsNew,
		IsHot:          g.IsHot,
	}

	_, err = global.ESClient.Index().Index(esGoods.GetIndexName()).BodyJson(esGoods).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
	return err
}

func (g *Goods) AfterUpdate(tx *gorm.DB) (err error) {
	esGoods := ESGoods{
		ID:             g.ID,
		CategoryID:     g.CategoryID,
		BrandID:        g.BrandID,
		Name:           g.Name,
		Brief:          g.Brief,
		MarketPrice:    g.MarketPrice,
		ShopPrice:      g.ShopPrice,
		ClickNum:       g.ClickNum,
		SoldNum:        g.SoldNum,
		FavNum:         g.FavNum,
		IsOnSale:       g.IsOnSale,
		IsFreeShipping: g.IsFreeShipping,
		IsNew:          g.IsNew,
		IsHot:          g.IsHot,
	}
	_, err = global.ESClient.Update().Index(esGoods.GetIndexName()).Doc(esGoods).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
	return err
}

func (g *Goods) AfterDelete(tx *gorm.DB) (err error) {
	_, err = global.ESClient.Delete().Index(ESGoods{}.GetIndexName()).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
	return err
}

func (g *Goods) ToProto() *proto.GoodsInfoResponse {
	return &proto.GoodsInfoResponse{
		Id:             g.ID,
		CategoryId:     g.CategoryID,
		Name:           g.Name,
		GoodsSn:        g.SN,
		Brief:          g.Brief,
		MarketPrice:    g.MarketPrice,
		ShopPrice:      g.ShopPrice,
		PreviewImage:   g.PreviewImage,
		Images:         g.Images,
		DescImages:     g.DescImages,
		ClickNum:       g.ClickNum,
		SoldNum:        g.SoldNum,
		FavNum:         g.FavNum,
		IsFreeShipping: g.IsFreeShipping,
		IsNew:          g.IsNew,
		IsHot:          g.IsHot,
		IsOnSale:       g.IsOnSale,
		Category: &proto.CategoryInfo{
			Id:   g.Category.ID,
			Name: g.Category.Name,
		},
		Brand: &proto.BrandInfoResponse{
			Id:   g.Brand.ID,
			Name: g.Brand.Name,
			Logo: g.Brand.Logo,
		},
	}
}

type Brand struct {
	BaseModel
	Name string `gorm:"type:varchar(50);not null;comment:品牌名称" json:"name"`
	Logo string `gorm:"type:varchar(255);default:'';not null;comment:品牌Logo" json:"logo"`
}

type Category struct {
	BaseModel
	Name             string      `gorm:"type:varchar(20);not null;comment:分类名称" json:"name"`
	ParentCategoryID int32       `gorm:"comment:父分类ID" json:"parent"`
	ParentCategory   *Category   `json:"-"`
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID;comment:子分类" json:"sub_category"`
	Level            int32       `gorm:"not null;default:1;comment:分类级别" json:"level"`
	IsTab            bool        `gorm:"default:false;not null;comment:是否展示" json:"is_tab"`
}

// 商品分类与品牌关系表，多对多
type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"index:idx_category_brand,unique;comment:分类ID" json:"category_id"`
	Category   Category

	BrandID int32 `gorm:"index:idx_category_brand,unique;comment:品牌ID" json:"brand_id"`
	Brand   Brand
}

// 轮播图
type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(255);not null;comment:轮播图图片" json:"image"`
	Url   string `gorm:"type:varchar(255);not null;comment:轮播图URL" json:"url"`
	Index int32  `gorm:"default:1;not null;comment:轮播图索引" json:"index"`
}
