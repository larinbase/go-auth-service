package dto

import (
	"github.com/google/uuid"
	"time"
)

type UserResponse struct {
	ID        uuid.UUID
	Email     string
	RoleName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserResponse(ID uuid.UUID, email string, roleName string, createdAt time.Time, updatedAt time.Time) *UserResponse {
	return &UserResponse{ID: ID, Email: email, RoleName: roleName, CreatedAt: createdAt, UpdatedAt: updatedAt}
}
