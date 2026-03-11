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
	GetChatsByUserID(ctx context.Context, userID, assistantID string, limit, offset int32) ([]*models.Chat, error)
	UpdateChat(ctx context.Context, id, assistantID, customerID string) (*models.Chat, error)
	DeleteChat(ctx context.Context, id string) error
	GetChatPagesCount(ctx context.Context, assistantID string, chatsPerPage int32) (int32, error)
	GetChatPage(ctx context.Context, assistantID string, page, chatsPerPage int32) ([]*models.Chat, error)
	GetChatPagesCountByUserID(ctx context.Context, userID string, assistantIDs []string, chatsPerPage int32) (int32, error)
	GetChatPageByUserID(ctx context.Context, userID string, assistantIDs []string, page, chatsPerPage int32) ([]*models.Chat, error)
	SearchChatsByCustomer(ctx context.Context, assistantIDs []string, search, userID string) ([]*models.Chat, int32, error)
	GetLatestChatByCustomer(ctx context.Context, assistantID, customerID string) (*models.Chat, error)
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
		UserID:      uuid.New().String(),
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

func (r *chatRepository) GetChatPagesCountByUserID(ctx context.Context, userID string, assistantIDs []string, chatsPerPage int32) (int32, error) {
	var count int64
	query := r.db.WithContext(ctx).Model(&models.Chat{}).Where("user_id = ?", userID)
	if len(assistantIDs) > 0 {
		query = query.Where("assistant_id IN ?", assistantIDs)
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count chats by user id: %w", err)
	}

	if chatsPerPage <= 0 {
		chatsPerPage = 10
	}

	pagesCount := int32((count + int64(chatsPerPage) - 1) / int64(chatsPerPage))
	return pagesCount, nil
}

func (r *chatRepository) GetChatPageByUserID(ctx context.Context, userID string, assistantIDs []string, page, chatsPerPage int32) ([]*models.Chat, error) {
	if page <= 0 {
		page = 1
	}
	if chatsPerPage <= 0 {
		chatsPerPage = 10
	}

	offset := (page - 1) * chatsPerPage

	var chats []*models.Chat
	query := r.db.WithContext(ctx).Where("user_id = ?", userID)
	if len(assistantIDs) > 0 {
		query = query.Where("assistant_id IN ?", assistantIDs)
	}

	if err := query.Order("created_at DESC").Limit(int(chatsPerPage)).Offset(int(offset)).Find(&chats).Error; err != nil {
		return nil, fmt.Errorf("failed to get chat page by user id: %w", err)
	}

	return chats, nil
}

func (r *chatRepository) SearchChatsByCustomer(ctx context.Context, assistantIDs []string, search, userID string) ([]*models.Chat, int32, error) {
	var chats []*models.Chat
	var count int64

	query := r.db.WithContext(ctx).Model(&models.Chat{})

	if len(assistantIDs) > 0 {
		query = query.Where("assistant_id IN ?", assistantIDs)
	}

	// If search term is provided, search by customer_id
	if search != "" {
		query = query.Where("customer_id LIKE ?", "%"+search+"%")
	}

	if userID != "" {
		query = query.Where("user_id = ?", userID)
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

func (r *chatRepository) GetLatestChatByCustomer(ctx context.Context, assistantID, customerID string) (*models.Chat, error) {
	var chat models.Chat
	result := r.db.WithContext(ctx).
		Where("assistant_id = ? AND customer_id = ?", assistantID, customerID).
		Order("started_at DESC").
		First(&chat)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Return nil if no chat found
		}
		return nil, fmt.Errorf("failed to get latest chat: %w", result.Error)
	}

	return &chat, nil
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

func (r *chatRepository) GetChatsByUserID(ctx context.Context, userID, assistantID string, limit, offset int32) ([]*models.Chat, error) {
	var chats []*models.Chat
	query := r.db.WithContext(ctx).Where("user_id = ?", userID)
	if assistantID != "" {
		query = query.Where("assistant_id = ?", assistantID)
	}
	if limit > 0 {
		query = query.Limit(int(limit))
	}
	if offset > 0 {
		query = query.Offset(int(offset))
	}
	if err := query.Order("created_at DESC").Find(&chats).Error; err != nil {
		return nil, fmt.Errorf("failed to get chats by user: %w", err)
	}

	return chats, nil
}

func (r *chatRepository) UpdateChat(ctx context.Context, id, assistantID, customerID string) (*models.Chat, error) {
	var chat models.Chat
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&chat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("chat not found")
		}
		return nil, fmt.Errorf("failed to get chat: %w", err)
	}

	if assistantID != "" {
		chat.AssistantID = assistantID
	}
	if customerID != "" {
		chat.CustomerID = &customerID
	}
	chat.UpdatedAt = time.Now()

	if err := r.db.WithContext(ctx).Save(&chat).Error; err != nil {
		return nil, fmt.Errorf("failed to update chat: %w", err)
	}

	return &chat, nil
}

func (r *chatRepository) DeleteChat(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Chat{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete chat: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("chat not found")
	}

	return nil
}
