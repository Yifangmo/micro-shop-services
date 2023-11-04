package main

import (
	"context"
	"strconv"

	"github.com/Yifangmo/micro-shop-services/goods/global"
	"github.com/Yifangmo/micro-shop-services/goods/initialize"
	"github.com/Yifangmo/micro-shop-services/goods/models"
)

func init() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()
	initialize.InitES()
}

func main() {
	err := global.DB.AutoMigrate(
		&models.Category{},
		&models.Brand{},
		&models.GoodsCategoryBrand{},
		&models.Banner{},
		&models.Goods{},
	)
	if err != nil {
		panic(err)
	}
	ESMigrate()
}

func ESMigrate() {
	var goods []models.Goods
	global.DB.Find(&goods)
	for _, g := range goods {
		esModel := models.ESGoods{
			ID:             g.ID,
			CategoryID:     g.CategoryID,
			BrandID:        g.BrandID,
			IsOnSale:       g.IsOnSale,
			IsFreeShipping: g.IsFreeShipping,
			IsNew:          g.IsNew,
			IsHot:          g.IsHot,
			Name:           g.Name,
			ClickNum:       g.ClickNum,
			SoldNum:        g.SoldNum,
			FavNum:         g.FavNum,
			MarketPrice:    g.MarketPrice,
			Brief:          g.Brief,
			ShopPrice:      g.ShopPrice,
		}

		_, err := global.ESClient.Index().Index(esModel.GetIndexName()).BodyJson(esModel).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
}
