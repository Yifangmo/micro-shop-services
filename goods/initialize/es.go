package initialize

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"

	"github.com/Yifangmo/micro-shop-services/goods/global"
	"github.com/Yifangmo/micro-shop-services/goods/models"
)

func InitES() {
	var err error
	global.ESClient, err = elastic.NewClient(
		elastic.SetURL(fmt.Sprintf("http://%s:%d", global.ServerConfig.ESConfig.Host, global.ServerConfig.ESConfig.Port)),
		elastic.SetSniff(false),
		elastic.SetTraceLog(zap.NewStdLog(zap.L())))
	if err != nil {
		panic(err)
	}

	exists, err := global.ESClient.IndexExists(models.ESGoods{}.GetIndexName()).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err = global.ESClient.CreateIndex(models.ESGoods{}.GetIndexName()).BodyString(models.ESGoods{}.GetMapping()).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
}
