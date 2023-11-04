package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/Yifangmo/micro-shop-services/inventory/global"
	"github.com/Yifangmo/micro-shop-services/inventory/models"
)

// 该函数实现从消息队列中获取要归还库存的消息并归还库存，消息只包含订单SN，
// 库存数据库 StockSellDetail 存放各订单SN对应的扣减库存的历史记录，以作为各商品归还数量的依据
// 为保证幂等性，订单归还库存后需要更新该订单SN对应的历史记录状态字段信息为 GIVEN_BACK（已归还）
func GiveBackInventory(ctx context.Context, message ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	type OrderInfo struct {
		OrderSn string
	}
	for i := range message {
		var orderInfo OrderInfo
		err := json.Unmarshal(message[i].Body, &orderInfo)
		if err != nil {
			zap.S().Errorf("Unmarshal message to orderSN error: %s\n", message[i].Body)
			return consumer.ConsumeSuccess, nil
		}
		tx := global.DB.Begin()
		// 先查询出该订单关于扣减库存的详情信息
		var orderDetail models.OrderDetail
		if dbres := tx.Model(&models.OrderDetail{}).Where(&models.OrderDetail{OrderSn: orderInfo.OrderSn, Status: models.SOLD}).Take(&orderDetail); dbres.Error != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				tx.Rollback()
				zap.S().Errorf("Update OrderDetail error: %v", dbres.Error)
				return consumer.ConsumeRetryLater, nil
			} else {
				tx.Commit()
				continue
			}
		}
		// 再逐个商品更新库存信息归还库存
		for _, orderGood := range orderDetail.Detail {
			if dbres := tx.Model(&models.Inventory{}).Where(&models.Inventory{GoodsID: orderGood.GoodsID}).Update("stock", gorm.Expr("stock+?", orderGood.Num)); dbres.Error != nil {
				tx.Rollback()
				zap.S().Errorf("Update Inventory error: %v", dbres.Error)
				return consumer.ConsumeRetryLater, nil
			}
		}
		// 最后更新该订单详情状态字段为已归还
		if dbres := tx.Model(&models.OrderDetail{}).Where(&models.OrderDetail{OrderSn: orderInfo.OrderSn}).Update("status", models.GIVEN_BACK); dbres.Error != nil {
			tx.Rollback()
			zap.S().Errorf("Update OrderDetail error: %v", dbres.Error)
			return consumer.ConsumeRetryLater, nil
		}
		tx.Commit()
	}
	return consumer.ConsumeSuccess, nil
}
