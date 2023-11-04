package models

type (
	PayType     uint8
	OrderStatus uint8
)

const (
	PAY_TYPE_NO_PAY PayType = iota
	PAY_TYPE_ALIPAY
	PAY_TYPE_WECHAT
)

const (
	ORDER_STATUS_WAITING_PAY OrderStatus = iota + 1
	ORDER_STATUS_SUCCESS
	ORDER_STATUS_CANCELLED
	ORDER_STATUS_TIMEOUT
	ORDER_STATUS_CLOSED
)
