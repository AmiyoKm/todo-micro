package handler

import (
	"context"

	"github.com/AmiyoKm/user-micro/internal/model"
	"github.com/google/uuid"
)

type Service interface {
	CreateUser(ctx context.Context, email, name, password string) (*model.User, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error)
	Login(ctx context.Context, email, password string) (string, error)
}
