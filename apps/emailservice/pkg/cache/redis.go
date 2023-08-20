package cache

import (
	"context"

	redis "github.com/redis/go-redis/v9"
)

type Redis struct {
	Cache

	client *redis.Client
}

func NewRedis(url, password string) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       0,
	})

	return &Redis{
		client,
	}
}

func (r *Redis) Put(ctx context.Context, key string, value interface{}) error {
	err := r.client.Set(ctx, key, value, 0).Err()
	return err
}

func (r *Redis) Get(ctx context.Context, key string) (interface{}, error) {
	val, err := r.client.Get(ctx, key).Result()
	return val, err
}
