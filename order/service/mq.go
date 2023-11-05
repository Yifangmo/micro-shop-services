package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"

	"github.com/Yifangmo/micro-shop-services/order/global"
	"github.com/Yifangmo/micro-shop-services/order/models"
	"github.com/Yifangmo/micro-shop-services/order/proto"
)

type CreateOrderListener struct {
	Ctx         context.Context
	OrderID     int32
	OrderAmount float64
	Code        codes.Code
	ErrMsg      string
}

func (o *CreateOrderListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	var (
		order       models.Order
		goodsIds    []int32
		shopCarts   []models.ShoppingCart
		goodsNumMap = make(map[int32]int32)
	)
	if err := json.Unmarshal(msg.Body, &order); err != nil {
		o.Code = codes.Internal
		o.ErrMsg = fmt.Sprintf("Unmarshal msg to Order error: %v", err)
		return primitive.RollbackMessageState
	}

	parentSpan := opentracing.SpanFromContext(o.Ctx)

	// 查询购物车需要结算的商品
	tracerSpan := opentracing.GlobalTracer().StartSpan("query_shopcart", opentracing.ChildOf(parentSpan.Context()))
	if dbres := global.DB.Where(&models.ShoppingCart{UserID: order.UserID, Checked: true}).Find(&shopCarts); dbres.RowsAffected == 0 {
		o.Code = codes.InvalidArgument
		o.ErrMsg = "no goods checked in shopping cart"
		return primitive.RollbackMessageState
	}
	tracerSpan.Finish()

	for _, shopCart := range shopCarts {
		goodsIds = append(goodsIds, shopCart.GoodsID)
		goodsNumMap[shopCart.GoodsID] = shopCart.GoodNumber
	}

	// 调用商品微服务查询商品信息
	tracerSpan = opentracing.GlobalTracer().StartSpan("query_goods", opentracing.ChildOf(parentSpan.Context()))
	goods, err := global.GoodsSrvClient.GetGoodsByIDs(context.Background(), &proto.GoodsIDsRequest{Ids: goodsIds})
	if err != nil {
		o.Code = codes.Internal
		o.ErrMsg = fmt.Sprintf("[GoodsSrv] query goods details error: %v", err)
		return primitive.RollbackMessageState
	}
	tracerSpan.Finish()

	var (
		orderAmount    float64
		orderGoods     []*models.OrderGoods
		goodsInventory []*proto.GoodsInventory // 每种商品需要扣减的库存
	)
	for _, good := range goods.GoodsMap {
		orderAmount += good.ShopPrice * float64(goodsNumMap[good.Id])
		orderGoods = append(orderGoods, &models.OrderGoods{
			GoodsID:     good.Id,
			GoodsName:   good.Name,
			GoodsImage:  good.PreviewImage,
			GoodsPrice:  good.ShopPrice,
			GoodsNumber: goodsNumMap[good.Id],
		})
		goodsInventory = append(goodsInventory, &proto.GoodsInventory{
			GoodsId: good.Id,
			Num:     goodsNumMap[good.Id],
		})
	}
	order.Amount = orderAmount
	o.OrderAmount = orderAmount

	tracerSpan = opentracing.GlobalTracer().StartSpan("inventory_rpc_sell", opentracing.ChildOf(parentSpan.Context()))
	if _, err = global.InventorySrvClient.Sell(context.Background(), &proto.OrderInfo{OrderSn: order.OrderSN, GoodsNumDelta: goodsInventory}); err != nil {
		o.Code = codes.ResourceExhausted
		o.ErrMsg = fmt.Sprintf("call inventory rpc for sell error: %v", err)
		return primitive.RollbackMessageState
	}
	tracerSpan.Finish()

	tx := global.DB.Begin()
	// 插入新订单记录
	tracerSpan = opentracing.GlobalTracer().StartSpan("create_order", opentracing.ChildOf(parentSpan.Context()))
	if dbres := tx.Save(&order); dbres.Error != nil {
		tx.Rollback()
		o.Code = codes.Internal
		o.ErrMsg = fmt.Sprintf("insert into Order db error: %v", dbres.Error)
		return primitive.CommitMessageState
	}
	tracerSpan.Finish()

	o.OrderID = order.ID
	// 填充外键
	for _, orderGood := range orderGoods {
		orderGood.OrderID = order.ID
	}

	// 插入订单商品信息记录
	tracerSpan = opentracing.GlobalTracer().StartSpan("create_order_goods", opentracing.ChildOf(parentSpan.Context()))
	if dbres := tx.CreateInBatches(orderGoods, 100); dbres.Error != nil {
		tx.Rollback()
		o.Code = codes.Internal
		o.ErrMsg = fmt.Sprintf("insert into OrderGoods db error: %v", dbres.Error)
		return primitive.CommitMessageState
	}
	tracerSpan.Finish()

	// 删除用户购物车记录
	tracerSpan = opentracing.GlobalTracer().StartSpan("delete_shopcart", opentracing.ChildOf(parentSpan.Context()))
	if dbres := tx.Where(&models.ShoppingCart{UserID: order.UserID, Checked: true}).Delete(&models.ShoppingCart{}); dbres.Error != nil {
		tx.Rollback()
		o.Code = codes.Internal
		o.ErrMsg = fmt.Sprintf("delete shopping cart records error: %v", dbres.Error)
		return primitive.CommitMessageState
	}
	tracerSpan.Finish()

	// 发送订单延时消息，在消费端处理超时订单
	msg = primitive.NewMessage(global.DelayOrderMsgTopic, msg.Body)
	msg.WithDelayTimeLevel(3)
	tracerSpan = opentracing.GlobalTracer().StartSpan("send_delay_order_msg", opentracing.ChildOf(parentSpan.Context()))
	_, err = global.MQProducer.SendSync(context.Background(), msg)
	if err != nil {
		tx.Rollback()
		o.Code = codes.Internal
		o.ErrMsg = fmt.Sprintf("send order_timeout message error: %v", err)
		return primitive.CommitMessageState
	}
	tracerSpan.Finish()

	tx.Commit()
	o.Code = codes.OK
	return primitive.RollbackMessageState
}

// 检查 OrderInfo 表里是否已经存在该 ordersn 的记录，
// 如果已存在则说明订单创建成功，回滚半消息来避免归还库存
// 如果不存在则说明订单未创建成功，需要提交半消息，但在订单服务中无法确定是否成功扣减库存（可能在扣减前出错，也可能在扣减后出错），需要在消费端即库存服务正确处理该消息
func (o *CreateOrderListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	var orderInfo models.Order
	if err := json.Unmarshal(msg.Body, &orderInfo); err != nil {
		o.Code = codes.Internal
		o.ErrMsg = fmt.Sprintf("Unmarshal msg to OrderInfo error: %v", err)
		return primitive.RollbackMessageState
	}

	if dbres := global.DB.Where(models.Order{OrderSN: orderInfo.OrderSN}).First(&orderInfo); dbres.RowsAffected == 0 {
		return primitive.CommitMessageState
	}

	return primitive.RollbackMessageState
}

// 所有订单信息都会被发送到延时消息队列来走这个逻辑，需要在逻辑中判断订单状态是否未支付，
// 若未支付说明由于未支付导致订单超时，需要修改订单状态为超时支付并归还库存
func HandleDelayOrderMessage(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for i := range msgs {
		var order models.Order
		json.Unmarshal(msgs[i].Body, &order)
		var queryOrder models.Order
		if dbres := global.DB.Model(models.Order{}).Where(models.Order{OrderSN: order.OrderSN}).First(&queryOrder); dbres.RowsAffected == 0 {
			return consumer.ConsumeSuccess, nil
		}
		if queryOrder.Status == models.ORDER_STATUS_WAITING_PAY { // 若未支付状态，则订单超时
			tx := global.DB.Begin()
			queryOrder.Status = models.ORDER_STATUS_TIMEOUT
			tx.Save(&queryOrder)
			// 发送归还库存消息
			_, err := global.MQProducer.SendSync(
				ctx,
				primitive.NewMessage(global.ServerConfig.RocketMQConfig.GiveBackInventoryTopic, msgs[i].Body),
			)
			if err != nil {
				tx.Rollback()
				return consumer.ConsumeRetryLater, nil
			}
			// 如果在这里服务挂了，需要在消费端来确保是否真的需要归还库存，具体需要查询订单状态是否已被修改为超时
			tx.Commit()
			return consumer.ConsumeSuccess, nil
		}
	}
	return consumer.ConsumeSuccess, nil
}
