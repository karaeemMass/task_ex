package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() (*redis.Client, error) {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}

	password := os.Getenv("REDIS_PASSWORD")

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0, // use default DB
	})

	ctx := context.Background()

	err := rdb.Ping(ctx).Err()
	if err != nil {
		_ = rdb.Close()
		return nil, fmt.Errorf("connect redis at %s: %w", addr, err)
	}

	fmt.Printf("Redis connected at %s\n", addr)

	return rdb, nil
}
