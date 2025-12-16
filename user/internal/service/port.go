package service

import (
	"context"

	"github.com/AmiyoKm/user-micro/internal/handler"
	"github.com/AmiyoKm/user-micro/internal/model"
	"github.com/google/uuid"
)

type Repository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
}

type Service interface {
	handler.Service
}
