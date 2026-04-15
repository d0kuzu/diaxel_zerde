package models

import "time"

// Assistant - соответствует таблице assistants в БД
type Assistant struct {
	ID               string    `json:"id" db:"id"`
	Name             string    `json:"name" db:"name"`
	Configuration    string    `json:"configuration" db:"configuration"`
	APIToken         string    `json:"api_token" db:"api_token"`
	TelegramBotToken string    `json:"telegram_bot_token" db:"telegram_bot_token"`
	Type             string    `json:"type" db:"type" gorm:"type:varchar(20);default:'api'"`
	UserID           string    `json:"user_id" db:"user_id"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

// TableName returns the table name for Assistant model
func (Assistant) TableName() string {
	return "assistants"
}
