package model

import (
	"github.com/AmiyoKm/todo-micro/gen/todopb"
	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
}

func (t *Todo) ToProto() *todopb.Todo {
	return &todopb.Todo{
		Id:          t.ID.String(),
		UserId:      t.UserId.String(),
		Title:       t.Title,
		Description: t.Description,
		Done:        t.Done,
	}
}
