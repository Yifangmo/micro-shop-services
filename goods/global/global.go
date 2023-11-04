package global

import (
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	config "github.com/Yifangmo/micro-shop-services/goods/configs"
)

var (
	DB                  *gorm.DB
	ServerConfig        config.ServerConfig
	NacosConfig         config.NacosConfig
	IsDebug             bool
	NacosConfigFileName string
	ESClient            *elastic.Client
)

func init() {
	viper.AutomaticEnv()
	IsDebug = viper.GetBool("MICRO_SHOP_DEBUG")
}
