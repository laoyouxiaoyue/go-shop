package model

import (
	"database/sql/driver"
	"encoding/json"
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

type GormList []string

func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}
func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}
