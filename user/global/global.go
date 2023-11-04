package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/Yifangmo/micro-shop-services/user/configs"
)

var (
	IsDebug             bool
	DB                  *gorm.DB
	ServerConfig        configs.ServerConfig
	NacosConfig         configs.NacosConfig
	NacosConfigFileName string
)

func init() {
	viper.AutomaticEnv()
	IsDebug = viper.GetBool("MICRO_SHOP_DEBUG")
}
