package cache

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AmiyoKm/todo-micro/configs"
	"github.com/AmiyoKm/todo-micro/internal/model"
	"github.com/redis/go-redis/v9"
)

const (
	TodoPrefix = "todos:"
)

type RedisClient struct {
	client     *redis.Client
	expiration time.Duration
}

func New(cfg *configs.RedisConfig) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &RedisClient{client: client, expiration: cfg.Expiration}, nil
}

// InvalidateAll flushes the entire Redis database (use with caution)
func (r *RedisClient) InvalidateAll(ctx context.Context) error {
	return r.client.FlushDB(ctx).Err()
}

// GetUserTodo retrieves a single todo from user's hash
func (r *RedisClient) GetUserTodo(ctx context.Context, userId, todoId string) (*model.Todo, error) {
	key := fmt.Sprintf("%s%s", TodoPrefix, userId)
	value, err := r.client.HGet(ctx, key, todoId).Result()
	if err != nil {
		return nil, err
	}

	var todo model.Todo
	if err := json.Unmarshal([]byte(value), &todo); err != nil {
		return nil, err
	}

	return &todo, nil
}

// GetUserTodos retrieves all todos for a user using Redis Hash
func (r *RedisClient) GetUserTodos(ctx context.Context, userId string) ([]*model.Todo, error) {
	key := fmt.Sprintf("%s%s", TodoPrefix, userId)
	result, err := r.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, redis.Nil
	}

	todos := make([]*model.Todo, 0, len(result))
	for _, value := range result {
		var todo model.Todo
		if err := json.Unmarshal([]byte(value), &todo); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}

	// Refresh TTL on successful read
	r.client.Expire(ctx, key, r.expiration)

	return todos, nil
}

func (r *RedisClient) SetUserTodo(ctx context.Context, userId string, todo *model.Todo) error {
	if todo == nil {
		return fmt.Errorf("todo cannot be nil")
	}

	key := fmt.Sprintf("%s%s", TodoPrefix, userId)
	valueBytes, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	pipe := r.client.Pipeline()
	pipe.HSet(ctx, key, todo.ID.String(), string(valueBytes))
	pipe.Expire(ctx, key, r.expiration)

	_, err = pipe.Exec(ctx)
	return err
}

func (r *RedisClient) SetUserTodos(ctx context.Context, userId string, todos []*model.Todo) error {
	if len(todos) == 0 {
		return nil
	}

	key := fmt.Sprintf("%s%s", TodoPrefix, userId)
	pipe := r.client.Pipeline()

	// Delete existing hash
	pipe.Del(ctx, key)

	// Store all todos in hash
	for _, todo := range todos {
		valueBytes, err := json.Marshal(todo)
		if err != nil {
			return err
		}
		pipe.HSet(ctx, key, todo.ID.String(), string(valueBytes))
	}

	pipe.Expire(ctx, key, r.expiration)

	_, err := pipe.Exec(ctx)
	return err
}

func (r *RedisClient) InvalidateUserTodos(ctx context.Context, userId string) error {
	key := fmt.Sprintf("%s%s", TodoPrefix, userId)
	return r.client.Del(ctx, key).Err()
}

func (r *RedisClient) Close() error {
	return r.client.Close()
}
