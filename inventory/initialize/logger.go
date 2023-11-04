package initialize

import (
	"go.uber.org/zap"

	"github.com/Yifangmo/micro-shop-services/inventory/global"
)

func InitLogger() {
	logger, _ := zap.NewProduction()
	if global.IsDebug {
		logger, _ = zap.NewDevelopment()
	}
	zap.ReplaceGlobals(logger)
}
