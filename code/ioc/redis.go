package ioc

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"shop/code/config"
	"time"
)

func InitRedis() *redis.Client {
	// 这里演示读取特定的某个字段
	zap.S().Infof("redis配置为:", zap.String("IP:", config.Cf.Redis.Host), zap.Int("Port:", config.Cf.Redis.Port))
	cmd := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Cf.Redis.Host, config.Cf.Redis.Port),
		Password: config.Cf.Redis.Password,
		DB:       config.Cf.Redis.DB,
	})
	zap.S().Info("Redis链接成功")
	cmd.Set(context.Background(), "213", "test", time.Minute*10)
	return cmd
}
