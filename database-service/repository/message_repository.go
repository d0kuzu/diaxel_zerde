package repository

import (
	"context"
	"fmt"
	"time"

	"diaxel_zerde/database-service/models"

	"gorm.io/gorm"
)

type MessageRepository interface {
	SaveMessage(ctx context.Context, chatID, role, content, platform string) (*models.Message, error)
	GetMessagesByChatID(ctx context.Context, chatID string, limit, offset int32) ([]*models.Message, error)
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) SaveMessage(ctx context.Context, chatID, role, content, platform string) (*models.Message, error) {
	message := models.Message{
		ChatID:  chatID,
		Role:    role,
		Content: content,
		Time:    time.Now(),
	}

	if err := r.db.WithContext(ctx).Create(&message).Error; err != nil {
		return nil, fmt.Errorf("failed to save message: %w", err)
	}

	return &message, nil
}

func (r *messageRepository) GetMessagesByChatID(ctx context.Context, chatID string, limit, offset int32) ([]*models.Message, error) {
	var messages []*models.Message
	if err := r.db.WithContext(ctx).Where("chat_id = ?", chatID).Order("time ASC").Limit(int(limit)).Offset(int(offset)).Find(&messages).Error; err != nil {
		return nil, fmt.Errorf("failed to get messages: %w", err)
	}

	return messages, nil
}
