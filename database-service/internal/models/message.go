package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	ChatID  uuid.UUID `gorm:"not null"`
	Role    string    // user / assistant / system
	Content string
	Time    time.Time
}
