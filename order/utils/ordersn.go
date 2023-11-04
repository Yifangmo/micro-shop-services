package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// 订单号生成规则：年月日时分秒+用户id+2位随机数
func GenerateOrderSN(userId int32) string {
	now := time.Now()
	rand.New(rand.NewSource(now.UnixNano()))
	return fmt.Sprintf("%d%d%d%d%d%d%d%d",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Nanosecond(),
		userId, rand.Intn(90)+10,
	)
}
