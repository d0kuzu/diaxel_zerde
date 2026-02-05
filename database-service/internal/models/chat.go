package models

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	AssistantID uuid.UUID `gorm:"type:uuid;not null"`
	CustomerID  string
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	StartedAt   time.Time
}
