package initialize

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"go.uber.org/zap"

	"github.com/Yifangmo/micro-shop-services/order/global"
	"github.com/Yifangmo/micro-shop-services/order/service"
)

func InitMQ() {
	var err error
	mqconfig := global.ServerConfig.RocketMQConfig

	// 启动订单延时消息消费者，用于修改超时的订单状态并归还库存
	global.OrderTimeoutMQConsumer, err = rocketmq.NewPushConsumer(
		consumer.WithNameServer(mqconfig.NameServer),
		consumer.WithGroupName(mqconfig.ConsumerGroup),
	)
	if err != nil {
		zap.S().Panicf("NewPushConsumer error: %v", err)
	}
	if err := global.OrderTimeoutMQConsumer.Subscribe(global.DelayOrderMsgTopic, consumer.MessageSelector{}, service.HandleDelayOrderMessage); err != nil {
		zap.S().Panicf("Subscribe %s topic failed: %v", global.DelayOrderMsgTopic, err)
	}
	err = global.OrderTimeoutMQConsumer.Start()
	if err != nil {
		zap.S().Errorf("Consumer start error: %v", err)
	} else {
		zap.S().Info("PushConsumer[%s] for topic[%s] starts...", mqconfig.ConsumerGroup, global.DelayOrderMsgTopic)
	}

	// 启动全局使用的生产者
	global.MQProducer, err = rocketmq.NewProducer(producer.WithNameServer(mqconfig.NameServer))
	if err != nil {
		zap.S().Panicf("NewProducer error: %v", err)
	}
	if err = global.MQProducer.Start(); err != nil {
		zap.S().Panicf("Start producer error: %v", err)
	}
	zap.S().Info("MQ Producer starts...")

	// 异步监听关闭
	global.WG.Add(1)
	go func() {
		defer global.WG.Done()
		<-global.ServerClosing
		err := global.OrderTimeoutMQConsumer.Shutdown()
		if err != nil {
			zap.S().Errorf("PushConsumer[%s] for topic[%s] shutdown error: %v", "micro-shop-inventory", mqconfig.GiveBackInventoryTopic, err)
		}
		err = global.MQProducer.Shutdown()
		if err != nil {
			zap.S().Errorf("MQProducer shutdown error: %v", err)
		}
	}()
}
