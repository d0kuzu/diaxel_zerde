package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID       uint      `gorm:"column:id;primaryKey;autoIncrement"`
	ChatID   uuid.UUID `gorm:"column:chat_id;index;not null"`
	Role     string    `gorm:"column:role;type:text;not null"`
	Content  string    `gorm:"column:message;type:text;not null"`
	Time     time.Time `gorm:"column:time;autoCreateTime"`
	Platform string    `gorm:"column:platform;type:text;not null;default:'web'"`
}
