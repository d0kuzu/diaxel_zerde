package models

import "time"

// RefreshToken - соответствует таблице refresh_tokens в БД
type RefreshToken struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	TokenHash string    `json:"token_hash" db:"token_hash"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// TableName returns the table name for RefreshToken model
func (RefreshToken) TableName() string {
	return "refresh_tokens"
}
