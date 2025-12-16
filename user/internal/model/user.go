package model

import (
	"github.com/AmiyoKm/user-micro/gen/userpb"
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	PasswordHash string    `json:"-"`
}

func (u *User) ToProto() *userpb.User {
	return &userpb.User{
		Id:    u.ID.String(),
		Email: u.Email,
		Name:  u.Name,
	}
}
