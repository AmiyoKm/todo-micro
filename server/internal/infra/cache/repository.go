package cache

import (
	"context"

	"github.com/AmiyoKm/todo-micro/internal/model"
	"github.com/AmiyoKm/todo-micro/internal/service"
)

type cacheRepo struct {
	client *RedisClient
}

func NewCacheRepository(client *RedisClient) service.CacheRepository {
	return &cacheRepo{client: client}
}

func (c *cacheRepo) Get(ctx context.Context, key string) (*model.Todo, error) {
	return c.client.Get(ctx, key)
}

func (c *cacheRepo) Set(ctx context.Context, key string, todo *model.Todo) error {
	return c.client.Set(ctx, key, todo)
}
