package models

type InventoryNew struct {
	Goods   int32 `gorm:"type:int;index"`
	Stocks  int32 `gorm:"type:int"`
	Freeze  int32 `gorm:"type:int"` //冻结库存
}