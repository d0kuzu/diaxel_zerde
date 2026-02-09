package models

import "time"

type Assistant struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string    `gorm:"column:name;type:text;not null"`
	BotToken  string    `gorm:"column:bot_token;type:text;not null;unique"`
	UserID    string    `gorm:"column:user_id;type:text;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Assistant) TableName() string {
	return "assistants"
}
