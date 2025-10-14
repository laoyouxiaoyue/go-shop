package svc

import (
	"gorm.io/gorm"
	"shop/goods_gozero/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config, db *gorm.DB) *ServiceContext {
	return &ServiceContext{
		DB:     db,
		Config: c,
	}
}
