package domain

import "time"

type RefreshSession struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	ExpiresIn time.Time `json:"expires_at" gorm:"not null"`
}

func (s *RefreshSession) SetExpiresIn(expiresIn time.Time) {
	s.ExpiresIn = expiresIn
}

func (s *RefreshSession) SetCreatedAt(createdAt time.Time) {
	s.CreatedAt = createdAt
}

func (s *RefreshSession) SetId(Id string) {
	s.ID = Id
}
