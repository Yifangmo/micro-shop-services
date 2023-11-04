package initialize

import (
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/Yifangmo/micro-shop-services/order/global"
	"github.com/Yifangmo/micro-shop-services/order/proto"
)

func InitRPC() {
	consulConfig := global.ServerConfig.ConsulConfig
	goodsSrvConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulConfig.Host, consulConfig.Port, global.ServerConfig.GoodsSrvName),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		//grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		zap.S().Panic("init GoodsSrv connection failed: %v", err)
	}
	global.GoodsSrvClient = proto.NewGoodsClient(goodsSrvConn)

	inventorySrvConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulConfig.Host, consulConfig.Port, global.ServerConfig.InventorySrvName),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		//grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		zap.S().Panicf("init InventorySrv connection failed: %v", err)
	}
	global.InventorySrvClient = proto.NewInventoryClient(inventorySrvConn)
}
