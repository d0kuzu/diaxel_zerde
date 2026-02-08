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
	GetChatPagesCount(ctx context.Context, assistantID string, chatsPerPage int32) (int32, error)
	GetChatPage(ctx context.Context, assistantID string, page, chatsPerPage int32) ([]*models.Chat, error)
	SearchChatsByUser(ctx context.Context, assistantID, search string) ([]*models.Chat, int32, error)
	UpdateMessageCount(ctx context.Context, chatID string) error
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

func (r *chatRepository) GetChatPagesCount(ctx context.Context, assistantID string, chatsPerPage int32) (int32, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.Chat{}).Where("assistant_id = ?", assistantID).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count chats: %w", err)
	}

	if chatsPerPage <= 0 {
		chatsPerPage = 10 // default value
	}

	pagesCount := int32((count + int64(chatsPerPage) - 1) / int64(chatsPerPage))
	return pagesCount, nil
}

func (r *chatRepository) GetChatPage(ctx context.Context, assistantID string, page, chatsPerPage int32) ([]*models.Chat, error) {
	if page <= 0 {
		page = 1
	}
	if chatsPerPage <= 0 {
		chatsPerPage = 10
	}

	offset := (page - 1) * chatsPerPage

	var chats []*models.Chat
	if err := r.db.WithContext(ctx).
		Where("assistant_id = ?", assistantID).
		Order("created_at DESC").
		Limit(int(chatsPerPage)).
		Offset(int(offset)).
		Find(&chats).Error; err != nil {
		return nil, fmt.Errorf("failed to get chat page: %w", err)
	}

	return chats, nil
}

func (r *chatRepository) SearchChatsByUser(ctx context.Context, assistantID, search string) ([]*models.Chat, int32, error) {
	var chats []*models.Chat
	var count int64

	query := r.db.WithContext(ctx).Model(&models.Chat{}).Where("assistant_id = ?", assistantID)

	// If search term is provided, search by user_id
	if search != "" {
		query = query.Where("user_id LIKE ?", "%"+search+"%")
	}

	// Get total count
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count search results: %w", err)
	}

	// Get the actual chats
	if err := query.Order("created_at DESC").Find(&chats).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to search chats: %w", err)
	}

	return chats, int32(count), nil
}

func (r *chatRepository) UpdateMessageCount(ctx context.Context, chatID string) error {
	// Count messages for this chat
	var messageCount int64
	if err := r.db.WithContext(ctx).Model(&models.Message{}).Where("chat_id = ?", chatID).Count(&messageCount).Error; err != nil {
		return fmt.Errorf("failed to count messages: %w", err)
	}

	// Update the message count in the chat
	if err := r.db.WithContext(ctx).Model(&models.Chat{}).Where("id = ?", chatID).Update("message_count", messageCount).Error; err != nil {
		return fmt.Errorf("failed to update message count: %w", err)
	}

	return nil
}
