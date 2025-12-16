package cache

import (
	"context"
	"fmt"
	"strings"

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
	// key format: "userid:todoid"
	parts := strings.Split(key, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid key format, expected 'userid:todoid', got: %s", key)
	}

	userId := parts[0]
	todoId := parts[1]
	return c.client.GetUserTodo(ctx, userId, todoId)
}

func (c *cacheRepo) Set(ctx context.Context, key string, todo *model.Todo) error {
	// key format: "userid:todoid"
	parts := strings.Split(key, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid key format, expected 'userid:todoid', got: %s", key)
	}

	userId := parts[0]
	return c.client.SetUserTodo(ctx, userId, todo)
}

func (c *cacheRepo) GetTodos(ctx context.Context, key string) ([]*model.Todo, error) {
	return c.client.GetUserTodos(ctx, key)
}

func (c *cacheRepo) SetTodos(ctx context.Context, key string, todos []*model.Todo) error {
	return c.client.SetUserTodos(ctx, key, todos)
}

func (c *cacheRepo) InvalidateUserTodos(ctx context.Context, userId string) error {
	return c.client.InvalidateUserTodos(ctx, userId)
}
