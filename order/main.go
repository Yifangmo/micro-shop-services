package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/opentracing/opentracing-go"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/Yifangmo/micro-shop-services/order/global"
	"github.com/Yifangmo/micro-shop-services/order/initialize"
	"github.com/Yifangmo/micro-shop-services/order/proto"
	"github.com/Yifangmo/micro-shop-services/order/service"
	"github.com/Yifangmo/micro-shop-services/order/utils"
	"github.com/Yifangmo/micro-shop-services/order/utils/otgrpc"
)

func main() {
	// 初始化全局变量
	initialize.InitConfig()
	initialize.InitLogger()
	// 链路追踪
	initialize.InitTracer()
	initialize.InitDB()
	initialize.InitRPC()
	// 启动延时订单消息消费者和全局生产者
	initialize.InitMQ()
	zap.S().Infof("Server config: %#v", global.ServerConfig)

	server := grpc.NewServer(grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer())))
	proto.RegisterOrderServer(server, &service.OrderServer{})
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
			zap.S().Error("Service deregister failed: ", err)
		} else {
			zap.S().Info("Service deregister success")
		}
		server.GracefulStop()
		global.WG.Wait()
	}()
	// 捕获退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
	case <-done:
	}
	close(global.ServerClosing)
	zap.S().Info("Server is closing...")
}
