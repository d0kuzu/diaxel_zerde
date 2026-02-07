package models

import (
	"time"

	"github.com/google/uuid"
)

type Assistant struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name          string    `gorm:"not null"`
	Configuration string
	UserID        string `gorm:"not null"`
	CreatedAt     time.Time
}
