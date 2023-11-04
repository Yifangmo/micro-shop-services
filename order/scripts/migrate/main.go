package main

import (
	"github.com/Yifangmo/micro-shop-services/inventory/global"
	"github.com/Yifangmo/micro-shop-services/inventory/initialize"
	"github.com/Yifangmo/micro-shop-services/order/models"
)

func init() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()
}

func main() {
	err := global.DB.AutoMigrate(&models.ShoppingCart{}, &models.Order{}, &models.OrderGoods{})
	if err != nil {
		panic(err)
	}
}
