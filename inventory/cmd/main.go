package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/Yifangmo/micro-shop-services/inventory/global"
	"github.com/Yifangmo/micro-shop-services/inventory/initialize"
	"github.com/Yifangmo/micro-shop-services/inventory/proto"
	"github.com/Yifangmo/micro-shop-services/inventory/service"
	"github.com/Yifangmo/micro-shop-services/inventory/utils"
)

func main() {
	// 初始化全局变量
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()
	zap.S().Infof("Server config: %#v", global.ServerConfig)

	server := grpc.NewServer()
	proto.RegisterInventoryServer(server, &service.InventoryServer{})
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		zap.S().Panic(err)
	}
	zap.S().Info("Listen to ", listener.Addr())
	port := listener.Addr().(*net.TCPAddr).Port
	// 注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	done := make(chan struct{})
	// 启动服务
	go func() {
		err = server.Serve(listener)
		if err != nil {
			zap.S().Error(err)
		}
		close(done)
	}()

	// 监听库存归还消息队列
	mqconfig := global.ServerConfig.RocketMQConfig
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer(mqconfig.NameServer),
		consumer.WithGroupName(mqconfig.ConsumerGroup),
	)
	if err != nil {
		zap.S().Panicf("NewPushConsumer error: %v", err)
	}
	if err := c.Subscribe(mqconfig.Topic, consumer.MessageSelector{}, service.GiveBackInventory); err != nil {
		zap.S().Panicf("Subscribe %s topic failed: %v", mqconfig.Topic, err)
	}
	err = c.Start()
	if err != nil {
		zap.S().Errorf("Consumer start error: %v", err)
	} else {
		zap.S().Info("PushConsumer[%s] for topic[%s] starts...", mqconfig.ConsumerGroup, mqconfig.Topic)
	}

	// consul 注册服务
	serviceID := uuid.NewV4().String()
	registryClient := utils.NewRegistryClient(global.ServerConfig.ConsulConfig.Host, global.ServerConfig.ConsulConfig.Port)
	err = registryClient.Register(global.ServerConfig.Host, port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceID)
	if err != nil {
		zap.S().Panic(err)
	}
	zap.S().Info("Service register success, service id: ", serviceID)
	// 退出前注销服务
	defer func() {
		if err = registryClient.Deregister(serviceID); err != nil {
			zap.S().Errorf("Service[%s] deregister failed: %v", serviceID, err)
		} else {
			zap.S().Infof("Service[%s] deregister success", serviceID)
		}
		server.GracefulStop()
		err := c.Shutdown()
		if err != nil {
			zap.S().Errorf("PushConsumer[%s] for topic[%s] shutdown error: %v", "micro-shop-inventory", mqconfig.Topic, err)
		}
		<-done
	}()

	// 捕获退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
	case <-done:
	}
}
