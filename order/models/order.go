package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int32          `json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type ShoppingCart struct {
	BaseModel
	UserID     int32 `gorm:"not null;index"`
	GoodsID    int32 `gorm:"not null;index"`
	GoodNumber int32 `gorm:"not null;comment:商品数量"`
	Checked    bool  `gorm:"not null;default:0;comment:是否选中以结算"`
}

type Order struct {
	BaseModel

	UserID  int32       `gorm:"not null;index"`
	OrderSN string      `gorm:"not null;type:varchar(30);index;comment:平台订单号"`
	TradeSN string      `gorm:"not null;default:'';type:varchar(100);comment:第三方平台交易流水号"`
	Amount  float64     `gorm:"not null;type:decimal(12,2);comment:订单金额"`
	PayType PayType     `gorm:"not null;default:0;comment:0(未支付),1(支付宝),2(微信)"`
	Status  OrderStatus `gorm:"not null;default:1;comment:1(待支付),2(交易成功),3(主动取消),4(超时未支付),5(交易关闭)"`
	PayTime *time.Time  `gorm:"comment:支付时间"`

	ConsigneeAddress string `gorm:"type:varchar(255);not null;comment:收货人地址"`
	ConsigneeName    string `gorm:"type:varchar(50);not null;comment:收货人姓名"`
	ConsigneeMobile  string `gorm:"type:varchar(20);not null;comment:收货人手机号"`
	Remark           string `gorm:"type:varchar(255);not null;default:'';comment:备注"`
}

type OrderGoods struct {
	BaseModel

	OrderID     int32   `gorm:"index"`
	GoodsID     int32   `gorm:"index"`
	GoodsName   string  `gorm:"type:varchar(100);not null;index;comment:商品名称"`
	GoodsImage  string  `gorm:"type:varchar(255);comment:商品图片"`
	GoodsPrice  float64 `gorm:"type:decimal(12,2);not null;comment:商品价格"`
	GoodsNumber int32   `gorm:"not null;comment:商品数量"`
}
