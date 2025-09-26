package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int32     `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(255);not null" json:"mobile"`
	Password string     `gorm:"type:varchar(255);not null" json:"password"`
	NickName string     `gorm:"type:varchar(255);not null" json:"nick_name"`
	Birthday *time.Time `gorm:"type:varchar(255);" json:"birthday"`
	Gender   string     `gorm:"type:varchar(255);default:male;type:varchar(6);" json:"gender"`
	Role     int        `gorm:"column:role;default: 1; type: int comment '1表示普通用户 2 表示管理员用户'" json:"role"`
}

func (u *User) TableName() string {
	return "old_user"
}
