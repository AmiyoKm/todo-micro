package service

import (
	"context"
	"fmt"

	"github.com/AmiyoKm/user-micro/internal/model"
	"github.com/AmiyoKm/user-micro/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo            Repository
	jwtSecret       string
	jwtExpirationHr int
}

func NewService(repo Repository, jwtSecret string, jwtExpirationHr int) Service {
	return &service{
		repo:            repo,
		jwtSecret:       jwtSecret,
		jwtExpirationHr: jwtExpirationHr,
	}
}

func (s *service) CreateUser(ctx context.Context, email, name, password string) (*model.User, error) {
	// Check if user already exists
	existingUser, err := s.repo.GetByEmail(ctx, email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", email)
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &model.User{
		Email:        email,
		Name:         name,
		PasswordHash: string(hashedPassword),
	}

	newUser, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return newUser, nil
}

func (s *service) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (s *service) Login(ctx context.Context, email, password string) (string, error) {
	// Get user by email
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	// Compare password with hash
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Email, s.jwtSecret, s.jwtExpirationHr)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}
