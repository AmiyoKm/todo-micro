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
	query := `INSERT INTO todos (user_id, title, description, done) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, todo.UserId, todo.Title, todo.Description, todo.Done).Scan(&todo.ID)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *repo) GetByID(ctx context.Context, id uuid.UUID, userId uuid.UUID) (*model.Todo, error) {
	query := `SELECT id, user_id, title, description, done FROM todos WHERE id = $1 AND user_id = $2`
	var todo model.Todo
	err := r.db.QueryRowContext(ctx, query, id, userId).Scan(&todo.ID, &todo.UserId, &todo.Title, &todo.Description, &todo.Done)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *repo) GetTodos(ctx context.Context, userId uuid.UUID) ([]*model.Todo, error) {
	query := `SELECT id, user_id, title, description, done FROM todos WHERE user_id = $1`
	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*model.Todo
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.ID, &todo.UserId, &todo.Title, &todo.Description, &todo.Done); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}
