package service

import (
	"context"

	"github.com/AmiyoKm/todo-micro/internal/handler"
	"github.com/AmiyoKm/todo-micro/internal/model"
	"github.com/google/uuid"
)

type Repository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.Todo, error)
	Create(ctx context.Context, todo *model.Todo) (*model.Todo, error)
}

type CacheRepository interface {
	Get(ctx context.Context, key string) (*model.Todo, error)
	Set(ctx context.Context, key string, todo *model.Todo) error
}

type Service interface {
	handler.Service
}
