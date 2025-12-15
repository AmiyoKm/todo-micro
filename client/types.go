package main

import "github.com/AmiyoKm/todo-micro/api-gateway/gen/todopb"

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (r *CreateTodoRequest) toProto() *todopb.CreateTodoRequest {
	return &todopb.CreateTodoRequest{
		Title:       r.Title,
		Description: r.Description,
	}
}

type GetTodoRequest struct {
	Id string `json:"id"`
}

func (r *GetTodoRequest) toProto() *todopb.GetTodoRequest {
	return &todopb.GetTodoRequest{
		Id: r.Id,
	}
}

type Todo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
