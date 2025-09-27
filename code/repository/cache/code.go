package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"time"
)

type CodeCache interface {
	Delete(ctx context.Context, addr string, subject string) error
	Get(ctx context.Context, addr string, subject string) (string, error)
	Set(ctx context.Context, addr string, subject string, code string) error
}

type RedisCodeCache struct {
	cmd        *redis.Client
	expiration time.Duration
}

func NewRedisCodeCache(cmd *redis.Client, expiration time.Duration) *RedisCodeCache {
	return &RedisCodeCache{cmd: cmd, expiration: expiration}
}

func (r *RedisCodeCache) Delete(ctx context.Context, addr string, subject string) error {
	return r.cmd.Del(ctx, r.key(addr, subject)).Err()
}

func (r *RedisCodeCache) Get(ctx context.Context, addr string, subject string) (string, error) {
	key := r.key(addr, subject)
	data, err := r.cmd.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return data, nil
}

func (r *RedisCodeCache) Set(ctx context.Context, addr string, subject string, code string, expiration int) error {
	key := r.key(addr, subject)
	zap.L().Info(key)
	return r.cmd.Set(ctx, key, code, time.Second*time.Duration(expiration)).Err()
}

func (r *RedisCodeCache) key(addr string, subject string) string {
	return fmt.Sprintf("code:%s:%s", subject, addr)
}
