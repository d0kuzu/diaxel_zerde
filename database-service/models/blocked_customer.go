package models

type BlockedCustomer struct {
	UserID string `gorm:"primaryKey;column:user_id"`
}
