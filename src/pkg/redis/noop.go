package redis

import (
	"context"
	"errors"
	"time"
)

var ErrCacheDisabled = errors.New("cache disabled")

type NoopRedis struct{}

func NewNoopRedisClient() IRedis {
	return &NoopRedis{}
}

func (n *NoopRedis) Get(ctx context.Context, key string) (string, error) {
	return "", ErrCacheDisabled
}

func (n *NoopRedis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return nil
}

func (n *NoopRedis) Del(ctx context.Context, keys ...string) error {
	return nil
}

func (n *NoopRedis) Exists(ctx context.Context, keys ...string) (int64, error) {
	return 0, nil
}

func (n *NoopRedis) HSet(ctx context.Context, key string, values ...interface{}) error {
	return nil
}

func (n *NoopRedis) HGet(ctx context.Context, key, field string) (string, error) {
	return "", ErrCacheDisabled
}

func (n *NoopRedis) HDel(ctx context.Context, key string, fields ...string) error {
	return nil
}

func (n *NoopRedis) HExists(ctx context.Context, key, field string) (bool, error) {
	return false, nil
}

func (n *NoopRedis) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return map[string]string{}, nil
}
