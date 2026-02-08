package repository

import (
	"context"
	"fmt"
	"time"

	"diaxel_zerde/database-service/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChatRepository interface {
	CreateChat(ctx context.Context, assistantID, customerID, platform string) (*models.Chat, error)
	GetChatByID(ctx context.Context, id string) (*models.Chat, error)
	GetChatsByCustomerID(ctx context.Context, customerID string) ([]*models.Chat, error)
}

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &chatRepository{db: db}
}

func (r *chatRepository) CreateChat(ctx context.Context, assistantID, customerID, platform string) (*models.Chat, error) {
	chat := models.Chat{
		ID:          uuid.New().String(),
		UserID:      uuid.New().String(), // TODO: получить реальный user_id из контекста
		AssistantID: assistantID,
		CustomerID:  &customerID,
		StartedAt:   time.Now(),
	}

	if err := r.db.WithContext(ctx).Create(&chat).Error; err != nil {
		return nil, fmt.Errorf("failed to create chat: %w", err)
	}

	return &chat, nil
}

func (r *chatRepository) GetChatByID(ctx context.Context, id string) (*models.Chat, error) {
	var chat models.Chat
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&chat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("chat not found")
		}
		return nil, fmt.Errorf("failed to get chat: %w", err)
	}

	return &chat, nil
}

func (r *chatRepository) GetChatsByCustomerID(ctx context.Context, customerID string) ([]*models.Chat, error) {
	var chats []*models.Chat
	if err := r.db.WithContext(ctx).Where("customer_id = ?", customerID).Order("started_at DESC").Find(&chats).Error; err != nil {
		return nil, fmt.Errorf("failed to get chats: %w", err)
	}

	return chats, nil
}
