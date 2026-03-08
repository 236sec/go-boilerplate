package redis

import (
	"time"

	"github.com/redis/go-redis/v9"
	"goboilerplate.com/config"
)

// Returns new redis client
func NewRedisClient(cfg *config.RedisConfig) *redis.Client {
	redisHost := cfg.RedisAddr

	if redisHost == "" {
		redisHost = ":6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:         redisHost,
		MinIdleConns: cfg.MinIdleConns,
		PoolSize:     cfg.PoolSize,
		PoolTimeout:  time.Duration(cfg.PoolTimeout) * time.Second,
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
	})

	return client
}