package redis

import (
	"github.com/go-redis/redis/v7"
	"time"
)

var R *redis.Client

type RedisConfig struct {
	Addr         string
	Password     string
	Db           int
	PoolSize     int
	PoolTimeout  time.Duration
	MinIdleConns int
}

func InitRedis(cfg *RedisConfig) error {
	R = redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password, // no password set
		DB:           cfg.Db,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
		PoolTimeout:  time.Second * cfg.PoolTimeout,
	})
	return nil
}
