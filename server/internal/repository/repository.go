package repository

import (
	"context"
	"database/sql"

	"github.com/AmiyoKm/todo-micro/internal/model"
	"github.com/AmiyoKm/todo-micro/internal/service"
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

func (r *repo) Create(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	query := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, todo.Title, todo.Description, todo.Done).Scan(&todo.ID)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *repo) GetByID(ctx context.Context, id uuid.UUID) (*model.Todo, error) {
	query := `SELECT id, title, description, done FROM todos WHERE id = $1`
	var todo model.Todo
	err := r.db.QueryRowContext(ctx, query, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}
