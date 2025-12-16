package handler

import (
	"context"

	"github.com/AmiyoKm/todo-micro/internal/model"
	"github.com/google/uuid"
)

type Service interface {
	GetByID(ctx context.Context, id uuid.UUID, userId uuid.UUID) (*model.Todo, error)
	Create(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	GetTodos(ctx context.Context, userId uuid.UUID) ([]*model.Todo, error)
}
