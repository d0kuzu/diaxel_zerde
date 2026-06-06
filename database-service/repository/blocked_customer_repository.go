package repository

import (
	"context"
	"fmt"

	"diaxel_zerde/database-service/models"

	"gorm.io/gorm"
)

type BlockedCustomerRepository interface {
	IsBlocked(ctx context.Context, userID string) (bool, error)
}

type blockedCustomerRepository struct {
	db *gorm.DB
}

func NewBlockedCustomerRepository(db *gorm.DB) BlockedCustomerRepository {
	return &blockedCustomerRepository{db: db}
}

func (r *blockedCustomerRepository) IsBlocked(ctx context.Context, userID string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.BlockedCustomer{}).Where("user_id = ?", userID).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("failed to check blocked status: %w", err)
	}
	return count > 0, nil
}
