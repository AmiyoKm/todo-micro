package service

import (
	"context"
	"log"

	"github.com/AmiyoKm/todo-micro/internal/model"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type service struct {
	repo  Repository
	cache CacheRepository
}

func NewService(repo Repository, cache CacheRepository) Service {
	return &service{
		repo:  repo,
		cache: cache,
	}
}

func (s *service) GetTodos(ctx context.Context, userId uuid.UUID) ([]*model.Todo, error) {
	// Try to get all todos from Redis Hash (key: todos:userid)
	cachedTodos, err := s.cache.GetTodos(ctx, userId.String())
	if err != nil && err != redis.Nil {
		log.Printf("Cache get error: %v", err)
	}

	if cachedTodos != nil {
		return cachedTodos, nil
	}

	todos, err := s.repo.GetTodos(ctx, userId)
	if err != nil {
		return nil, err
	}

	// Store all todos in Redis Hash using HSET for each todo
	go func() {
		if err := s.cache.SetTodos(ctx, userId.String(), todos); err != nil {
			log.Printf("Cache set error: %v", err)
		}
	}()

	return todos, nil
}

func (s *service) GetByID(ctx context.Context, id uuid.UUID, userId uuid.UUID) (*model.Todo, error) {
	// Try to get from user's Redis Hash using HGET (key: todos:userid, field: todoid)
	key := userId.String() + ":" + id.String()
	cachedTodo, err := s.cache.Get(ctx, key)
	if err != nil && err != redis.Nil {
		log.Printf("Cache get error: %v", err)
	}

	if cachedTodo != nil {
		return cachedTodo, nil
	}

	todo, err := s.repo.GetByID(ctx, id, userId)
	if err != nil {
		return nil, err
	}

	// Store in user's Redis Hash using HSET (refreshes TTL on the entire hash)
	go func() {
		if err := s.cache.Set(ctx, key, todo); err != nil {
			log.Printf("Cache set error: %v", err)
		}
	}()

	return todo, nil
}

func (s *service) Create(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	newTodo, err := s.repo.Create(ctx, todo)
	if err != nil {
		return nil, err
	}

	if err := s.cache.InvalidateUserTodos(ctx, newTodo.UserId.String()); err != nil {
		log.Printf("Cache invalidation error: %v", err)
	}

	return newTodo, nil
}
