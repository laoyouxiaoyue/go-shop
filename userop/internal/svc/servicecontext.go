package svc

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"shop/userop/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config, db *gorm.DB, redis *redis.Client) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Db:     db,
		Redis:  redis,
	}
}
