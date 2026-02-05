package models

import "github.com/google/uuid"

type Assistant struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name          string    `gorm:"not null"`
	Configuration string
	UserID        uuid.UUID `gorm:"type:uuid;not null"`
}
