package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email        string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	Role         string    `gorm:"default:user"`
	CreatedAt    time.Time
}

// GetUserIDAsString возвращает ID пользователя как строку для совместимости с Chat моделью
func (u *User) GetUserIDAsString() string {
	return u.ID.String()
}
