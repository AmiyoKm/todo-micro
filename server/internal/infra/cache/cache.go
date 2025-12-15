package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AmiyoKm/todo-micro/configs"
	"github.com/AmiyoKm/todo-micro/internal/model"
	"github.com/redis/go-redis/v9"
)

const (
	TodoPrefix = "todo:"
	Expiration = time.Second * 5
)

type RedisClient struct {
	client *redis.Client
}

func New(cfg *configs.RedisConfig) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &RedisClient{client: client}, nil
}

func (r *RedisClient) Get(ctx context.Context, key string) (*model.Todo, error) {
	withPrefix := fmt.Sprintf("%s%s", TodoPrefix, key)
	value, err := r.client.Get(ctx, withPrefix).Result()

	if err != nil {
		return nil, err
	}
	var todo model.Todo
	err = json.Unmarshal([]byte(value), &todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *RedisClient) Set(ctx context.Context, key string, value *model.Todo) error {
	if value == nil {
		return fmt.Errorf("value cannot be nil")
	}
	valueBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	withPrefix := fmt.Sprintf("%s%s", TodoPrefix, key)
	return r.client.Set(ctx, withPrefix, string(valueBytes), Expiration).Err()
}

func (r *RedisClient) Invalidate(ctx context.Context, key string) error {
	withPrefix := fmt.Sprintf("%s%s", TodoPrefix, key)
	return r.client.Del(ctx, withPrefix).Err()
}

func (r *RedisClient) InvalidateAll(ctx context.Context) error {
	return r.client.FlushDB(ctx).Err()
}

func (r *RedisClient) Close() error {
	return r.client.Close()
}
