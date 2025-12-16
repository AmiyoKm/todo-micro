package handler

import (
	"context"

	"github.com/AmiyoKm/user-micro/gen/userpb"
	"github.com/AmiyoKm/user-micro/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	userpb.UnimplementedUserServiceServer
	svc       Service
	jwtSecret string
}

func NewGRPCHandler(server *grpc.Server, svc Service, jwtSecret string) *Handler {
	handler := &Handler{
		svc:       svc,
		jwtSecret: jwtSecret,
	}
	userpb.RegisterUserServiceServer(server, handler)
	return handler
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	user, err := h.svc.CreateUser(ctx, req.Email, req.Name, req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &userpb.CreateUserResponse{
		User: user.ToProto(),
	}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	if req.Jwt == "" {
		return nil, status.Error(codes.InvalidArgument, "jwt token is required")
	}

	// Extract user ID from JWT token
	userID, err := utils.ExtractUserID(req.Jwt, h.jwtSecret)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid or expired token")
	}

	user, err := h.svc.GetUser(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &userpb.GetUserResponse{
		User: user.ToProto(),
	}, nil
}

func (h *Handler) Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	token, err := h.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &userpb.LoginResponse{
		Jwt: token,
	}, nil
}
