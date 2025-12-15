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

func (s *service) GetByID(ctx context.Context, id uuid.UUID) (*model.Todo, error) {
	cachedTodo, err := s.cache.Get(ctx, id.String())
	if err != nil && err != redis.Nil {
		log.Printf("Cache get error: %v", err)
	}

	if cachedTodo != nil {
		return cachedTodo, nil
	}

	todo, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := s.cache.Set(ctx, todo.ID.String(), todo); err != nil {
		log.Printf("Cache set error: %v", err)
	}

	return todo, nil
}

func (s *service) Create(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	newTodo, err := s.repo.Create(ctx, todo)
	if err != nil {
		return nil, err
	}

	if err := s.cache.Set(ctx, newTodo.ID.String(), newTodo); err != nil {
		log.Printf("Cache set error: %v", err)
	}

	return newTodo, nil
}
