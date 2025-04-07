package repository

import (
	"auth-service/internal/domain"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshSessionRepository struct {
	db *gorm.DB
}

func NewRefreshSessionRepository(db *gorm.DB) RefreshSessionRepository {
	return RefreshSessionRepository{db: db}
}

func (r *RefreshSessionRepository) Create(session *domain.RefreshSession) error {
	result := r.db.Create(session)
	return result.Error
}

func (r *RefreshSessionRepository) GetById(id uuid.UUID) (*domain.RefreshSession, error) {
	refreshSession := &domain.RefreshSession{}
	result := r.db.Where("id = ?", id).First(refreshSession)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("session not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return refreshSession, nil
}

func (r *RefreshSessionRepository) Update(session *domain.RefreshSession) error {
	result := r.db.Updates(session)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("session not found")
	}
	return result.Error
}
