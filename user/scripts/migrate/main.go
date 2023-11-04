package main

import (
	"fmt"

	"github.com/Yifangmo/micro-shop-services/user/global"
	"github.com/Yifangmo/micro-shop-services/user/initialize"
	"github.com/Yifangmo/micro-shop-services/user/models"
	"github.com/Yifangmo/micro-shop-services/user/utils"

	"github.com/opentracing/opentracing-go/log"
)

func init() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()
}

func main() {
	err := global.DB.AutoMigrate(&models.User{}, &models.LeavingMessage{}, &models.UserFav{}, models.Address{})
	if err != nil {
		panic(err)
	}
	for i := 0; i < 5; i++ {
		dbres := global.DB.Create(&models.User{
			Nickname: fmt.Sprintf("admin%d", i),
			Mobile:   fmt.Sprintf("1323212222%d", i),
			Password: utils.GenStorePassword("admin"),
			Role:     models.ADMIN,
		})
		if dbres.Error != nil {
			log.Error(dbres.Error)
		}
	}
}
