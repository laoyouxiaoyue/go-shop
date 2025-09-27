package repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"shop/code/errs"
	"shop/code/repository/cache"
)

type CodeRepository interface {
	Create(ctx context.Context, addr string, subject string, code string, expiration int) error
	Get(ctx context.Context, addr string, subject string) (string, error)
	Delete(ctx context.Context, addr string, subject string) error
}

type CacheCodeRepository struct {
	cache *cache.RedisCodeCache
}

func NewCodeRepository(cache *cache.RedisCodeCache) *CacheCodeRepository {
	return &CacheCodeRepository{cache: cache}
}

func (c *CacheCodeRepository) Create(ctx context.Context, addr string, subject string, code string, expiration int) error {
	_, err := c.cache.Get(ctx, addr, subject)
	//if err != nil {
	//	zap.Error(err)
	//}
	//zap.L().Info("<UNK>", zap.String("addr", addr), zap.String("subject", subject), zap.String("code", code))
	if err == nil || !errors.Is(err, redis.Nil) {
		return errs.ErrTooManyRequest
	}
	err = c.cache.Set(ctx, addr, subject, code, expiration)
	if err != nil {
		return err
	}
	return nil
}

func (c *CacheCodeRepository) Get(ctx context.Context, addr string, subject string) (string, error) {
	return c.cache.Get(ctx, addr, subject)
}
func (c *CacheCodeRepository) Delete(ctx context.Context, addr string, subject string) error {
	err := c.cache.Delete(ctx, addr, subject)
	if err != nil {
		return err
	}
	return nil
}
