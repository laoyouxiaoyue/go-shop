package ioc

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"shop/inventory/internal/config"
	"time"
)

func InitRedis(c config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%d", c.RedisConfig.Host, c.RedisConfig.Port),
		Password: c.RedisConfig.Password,
		DB:       c.RedisConfig.DB,
		PoolSize: c.RedisConfig.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		zap.S().Panicf("Redis 连接失败: %v", err)
		return nil
	}
    zap.L().Info("Redis 连接成功", zap.String("addr", fmt.Sprintf("%s:%d", c.RedisConfig.Host, c.RedisConfig.Port)))
	return rdb
}
