package repository

import (
	"context"
	"database/sql"

	"github.com/AmiyoKm/user-micro/internal/model"
	"github.com/AmiyoKm/user-micro/internal/service"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) service.Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, user *model.User) (*model.User, error) {
	query := `INSERT INTO users (email, name, password_hash) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, user.Email, user.Name, user.PasswordHash).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repo) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	query := `SELECT id, email, name, password_hash FROM users WHERE id = $1`
	var user model.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Email, &user.Name, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT id, email, name, password_hash FROM users WHERE email = $1`
	var user model.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Name, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
