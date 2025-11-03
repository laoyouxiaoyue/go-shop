package svc

import (
    "github.com/go-redsync/redsync/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"shop/inventory/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
	Redis  *redis.Client
    Lock   *redsync.Redsync
}

func NewServiceContext(c config.Config, db *gorm.DB, redis *redis.Client, lock *redsync.Redsync) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Db:     db,
        Redis:  redis,
        Lock:   lock,
	}
}
