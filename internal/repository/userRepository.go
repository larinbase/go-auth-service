package repository

import (
	"errors"

	"auth-service/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	user.ID = uuid.New()
	result := r.db.Create(user)
	return result.Error
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	result := r.db.Where("email = ?", email).Preload("Role").First(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepository) GetByID(id uuid.UUID) (*domain.User, error) {
	user := &domain.User{}
	result := r.db.Where("id = ?", id).Preload("Role").First(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
