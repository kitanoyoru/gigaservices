package cache

import "context"

type Cache interface {
	Put(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (interface{}, error)
}
