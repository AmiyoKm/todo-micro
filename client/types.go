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

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Jwt string `json:"jwt"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
