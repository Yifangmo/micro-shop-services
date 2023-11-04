package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Inventory struct {
	GoodsID   int32          `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Stock     int32          `gorm:"not null;default:0"`
}

type OrderDetail struct {
	OrderSn   string         `gorm:"primaryKey;type:varchar(255);"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Status    OrderStatus    `gorm:"default:0;not null;comment:1(已扣减),2(已归还)"`
	Detail    GoodsNumDelta  `gorm:"type:json;not null;comment:各商品扣减或归还的数量"`
}

type GoodsDetail struct {
	GoodsID int32 `json:"goods_id"`
	Num     int32 `json:"num"`
}

type GoodsNumDelta []GoodsDetail

func (g GoodsNumDelta) Value() (driver.Value, error) {
	return json.Marshal(g)
}

func (g *GoodsNumDelta) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

type OrderStatus uint8

const (
	SOLD OrderStatus = iota + 1
	GIVEN_BACK
)
