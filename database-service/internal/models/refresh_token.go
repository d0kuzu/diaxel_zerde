package models

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    string    `gorm:"not null"`
	Token     string    `gorm:"not null"`
	ExpiresAt time.Time
	CreatedAt time.Time
}
