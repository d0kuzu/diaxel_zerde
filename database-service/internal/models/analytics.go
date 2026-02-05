package models

import "github.com/google/uuid"

type Analytics struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	AssistantID    uuid.UUID `gorm:"not null"`
	TotalChats     int
	ActiveUsers    int
	EngagementRate float64
}
