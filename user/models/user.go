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

	Mobile   string `gorm:"type:varchar(20);not null;unique"`
	Password string `gorm:"type:varchar(100);not null"`
	Nickname string `gorm:"type:varchar(255)"`
	Birthday *time.Time
	Gender   Gender `gorm:"type:tinyint;not null;default:0;comment:0(未知),1(男),2(女)"`
	Role     Role   `gorm:"type:tinyint;not null;default:1;comment:1(普通用户),2(管理员)"`
}

type LeavingMessage struct {
	BaseModel

	User        int32       `gorm:"type:int;index"`
	MessageType MessageType `gorm:"type:tinyint;not null;default:0;comment:留言类型,1(留言),2(投诉),3(询问),4(售后),5(求购)"`
	Subject     string      `gorm:"type:varchar(100)"`
	Message     string      `gorm:"type:varchar(255)"`
	File        string      `gorm:"type:varchar(200)"`
}

type Address struct {
	BaseModel

	UserID           int32  `gorm:"index"`
	Province         string `gorm:"type:varchar(64);not null"`
	City             string `gorm:"type:varchar(64);not null"`
	District         string `gorm:"type:varchar(64);not null"`
	ConsigneeAddress string `gorm:"type:varchar(255);not null"`
	ConsigneeName    string `gorm:"type:varchar(255);not null"`
	ConsigneeMobile  string `gorm:"type:varchar(20);not null"`
}

type UserFav struct {
	BaseModel

	UserID  int32 `gorm:"unique"`
	GoodsID int32 `gorm:"unique"`
}
