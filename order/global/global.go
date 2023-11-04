package global

import (
	"sync"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/Yifangmo/micro-shop-services/order/configs"
	"github.com/Yifangmo/micro-shop-services/order/proto"
)

var (
	IsDebug                bool
	ServerClosing          chan struct{}
	WG                     sync.WaitGroup
	DB                     *gorm.DB
	ServerConfig           configs.ServerConfig
	NacosConfig            configs.NacosConfig
	NacosConfigFileName    string
	GoodsSrvClient         proto.GoodsClient
	InventorySrvClient     proto.InventoryClient
	OrderTimeoutMQConsumer rocketmq.PushConsumer
	MQProducer             rocketmq.Producer
	DelayOrderMsgTopic     string
)

func init() {
	viper.AutomaticEnv()
	IsDebug = viper.GetBool("MICRO_SHOP_DEBUG")
	ServerClosing = make(chan struct{})
	DelayOrderMsgTopic = "delay_order_msg"
}
