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

type User struct {
	BaseModel

	Mobile   string `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Nickname string `gorm:"type:varchar(64)"`
	Birthday *time.Time
	Gender   Gender `gorm:"type:tinyint;not null;default:0;comment:0(男), 1(女)"`
	Role     Role   `gorm:"type:tinyint;not null;default:0;comment:0(普通用户), 1(管理员)"`
}

type LeavingMessage struct {
	BaseModel

	User        int32       `gorm:"type:int;index"`
	MessageType MessageType `gorm:"type:tinyint;not null;default:0;comment:留言类型,0(留言),1(投诉),2(询问),3(售后),4(求购)"`
	Subject     string      `gorm:"type:varchar(100)"`
	Message     string      `gorm:"type:varchar(255)"`
	File        string      `gorm:"type:varchar(200)"`
}

type Address struct {
	BaseModel

	UserID           int32  `gorm:"type:int;index"`
	Province         string `gorm:"type:varchar(64)"`
	City             string `gorm:"type:varchar(64)"`
	District         string `gorm:"type:varchar(64)"`
	ConsigneeAddress string `gorm:"type:varchar(100)"`
	ConsigneeName    string `gorm:"type:varchar(64)"`
	ConsigneeMobile  string `gorm:"type:varchar(20)"`
}

type UserFav struct {
	BaseModel

	UserID  int32 `gorm:"type:int;index:idx_user_goods,unique"`
	GoodsID int32 `gorm:"type:int;index:idx_user_goods,unique"`
}
