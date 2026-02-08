package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// User - соответствует таблице users в БД
type User struct {
	ID           string    `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"password_hash" db:"password_hash"`
	Role         string    `json:"role" db:"role"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// Assistant - соответствует таблице assistants в БД
type Assistant struct {
	ID            string    `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	Configuration string    `json:"configuration" db:"configuration"`
	UserID        string    `json:"user_id" db:"user_id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

// Chat - соответствует таблице chats в БД (из AI сервиса)
type Chat struct {
	UserID      string    `json:"user_id" db:"user_id"`
	AssistantID string    `json:"assistant_id" db:"assistant_id"`
	CustomerID  string    `json:"customer_id" db:"customer_id"`
	StartedAt   time.Time `json:"started_at" db:"started_at"`
}

// Message - соответствует таблице messages в БД (из AI сервиса)
type Message struct {
	ID         uint      `json:"id" db:"id"`
	ChatUserID string    `json:"chat_user_id" db:"chat_user_id"`
	Role       string    `json:"role" db:"role"`
	Content    string    `json:"content" db:"content"`
	Time       time.Time `json:"time" db:"time"`
}

// RefreshToken - соответствует таблице refresh_tokens в БД
type RefreshToken struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	TokenHash string    `json:"token_hash" db:"token_hash"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Analytics - соответствует таблице analytics в БД
type Analytics struct {
	ID             string  `json:"id" db:"id"`
	AssistantID    string  `json:"assistant_id" db:"assistant_id"`
	TotalChats     int32   `json:"total_chats" db:"total_chats"`
	ActiveUsers    int32   `json:"active_users" db:"active_users"`
	EngagementRate float64 `json:"engagement_rate" db:"engagement_rate"`
}

// Custom type for UUID to handle database/sql driver
type UUID string

func (u UUID) Value() (driver.Value, error) {
	if u == "" {
		return nil, nil
	}
	return string(u), nil
}

func (u *UUID) Scan(value interface{}) error {
	if value == nil {
		*u = ""
		return nil
	}

	switch v := value.(type) {
	case string:
		*u = UUID(v)
	case []byte:
		*u = UUID(v)
	default:
		return fmt.Errorf("cannot scan %T into UUID", value)
	}
	return nil
}

// TableName returns the table name for User model
func (User) TableName() string {
	return "users"
}

// TableName returns the table name for Assistant model
func (Assistant) TableName() string {
	return "assistants"
}

// TableName returns the table name for Chat model
func (Chat) TableName() string {
	return "chats"
}

// TableName returns the table name for Message model
func (Message) TableName() string {
	return "messages"
}

// TableName returns the table name for RefreshToken model
func (RefreshToken) TableName() string {
	return "refresh_tokens"
}

// TableName returns the table name for Analytics model
func (Analytics) TableName() string {
	return "analytics"
}
