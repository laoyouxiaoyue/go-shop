package svc

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"shop/goods_gozero/goodsclient"
	"shop/inventory/inventoryclient"
	"shop/order/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
	Redis  *redis.Client
	Goods  goodsclient.Goods
	Inv    inventoryclient.Inventory
}

func NewServiceContext(c config.Config, db *gorm.DB, redis *redis.Client, goods goodsclient.Goods, inv inventoryclient.Inventory) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Db:     db,
		Redis:  redis,
		Goods:  goods,
		Inv:    inv,
	}
}
