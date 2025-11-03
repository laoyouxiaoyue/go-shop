package model

import (
	"gorm.io/gorm"
	"time"
)

// Inventory 库存
type BaseModel struct {
	ID        int32     `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}
type Inventory struct {
	BaseModel
	Goods   int32 `gorm:"type:int comment '商品id';index"`
	Stocks  int32 `gorm:"type:int comment '商品库存'"`
	Version int32 `gorm:"type:int"` //分布式乐观锁，用来判断并发情况下库存是否扣减一致
}

func (Inventory) TableName() string {
	return "inventory"
}
