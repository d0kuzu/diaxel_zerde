package models

import "github.com/google/uuid"

type Chat struct {
	ID       uuid.UUID `gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID   uuid.UUID `gorm:"column:user_id;index;not null"`
	Messages []Message `gorm:"foreignKey:ChatID;references:ID"`
	IsClient bool      `gorm:"column:is_client;default:false"`
}
