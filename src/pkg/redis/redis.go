package redis

import (
	"time"

	"github.com/redis/go-redis/v9"
	"goboilerplate.com/config"
)

// Returns new redis client
func NewRedisClient(cfg *config.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.RedisAddr,
		MinIdleConns: cfg.MinIdleConns,
		PoolSize:     cfg.PoolSize,
		PoolTimeout:  time.Duration(cfg.PoolTimeout) * time.Second,
		Password:     cfg.RedisPassword,
		DB:           cfg.RedisDB,
	})

	return client
}