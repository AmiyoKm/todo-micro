package handler

import (
	"context"

	"github.com/AmiyoKm/todo-micro/gen/todopb"
	"github.com/AmiyoKm/todo-micro/internal/model"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type Handler struct {
	todopb.UnimplementedTodoServiceServer
	svc Service
}

func NewGRPCHandler(server *grpc.Server, svc Service) *Handler {
	handler := &Handler{
		svc: svc,
	}
	todopb.RegisterTodoServiceServer(server, handler)
	return handler
}

func (h *Handler) CreateTodo(ctx context.Context, req *todopb.CreateTodoRequest) (*todopb.TodoResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}

	todo := &model.Todo{
		Title:       req.Title,
		Description: req.Description,
		UserId:      userId,
	}

	newTodo, err := h.svc.Create(ctx, todo)
	if err != nil {
		return nil, err
	}

	return &todopb.TodoResponse{
		Todo: newTodo.ToProto(),
	}, nil
}

func (h *Handler) GetTodo(ctx context.Context, req *todopb.GetTodoRequest) (*todopb.TodoResponse, error) {
	paramID := req.GetId()

	id, err := uuid.Parse(paramID)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}

	todo, err := h.svc.GetByID(ctx, id, userId)
	if err != nil {
		return nil, err
	}

	return &todopb.TodoResponse{
		Todo: todo.ToProto(),
	}, nil
}
