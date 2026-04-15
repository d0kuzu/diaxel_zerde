package models

import (
	"time"
)

// User - соответствует таблице users в БД
type User struct {
	ID           string    `json:"id" db:"id"`
	TelegramID   *string   `json:"telegram_id" db:"telegram_id"`
	Username     *string   `json:"username" db:"username"`
	FirstName    *string   `json:"first_name" db:"first_name"`
	LastName     *string   `json:"last_name" db:"last_name"`
	Email        *string   `json:"email" db:"email"`
	PasswordHash *string   `json:"password_hash" db:"password_hash"`
	Role         *string   `json:"role" db:"role"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// TableName returns the table name for User model
func (User) TableName() string {
	return "users"
}
