package repository

import (
	"context"
	"fmt"
	"time"

	"diaxel_zerde/database-service/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AssistantRepository interface {
	CreateAssistant(ctx context.Context, name, apiToken, userID string) (*models.Assistant, error)
	GetAssistantByID(ctx context.Context, id string) (*models.Assistant, error)
	GetAssistantByAPIToken(ctx context.Context, apiToken string) (*models.Assistant, error)
	UpdateAssistant(ctx context.Context, id, name, configuration, apiToken string) (*models.Assistant, error)
	DeleteAssistant(ctx context.Context, id string) error
}

type assistantRepository struct {
	db *gorm.DB
}

func NewAssistantRepository(db *gorm.DB) AssistantRepository {
	return &assistantRepository{db: db}
}

func (r *assistantRepository) CreateAssistant(ctx context.Context, name, apiToken, userID string) (*models.Assistant, error) {
	assistant := models.Assistant{
		ID:            uuid.New().String(),
		Name:          name,
		Configuration: "",
		APIToken:      apiToken,
		UserID:        userID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := r.db.WithContext(ctx).Create(&assistant).Error; err != nil {
		return nil, fmt.Errorf("failed to create assistant: %w", err)
	}

	return &assistant, nil
}

func (r *assistantRepository) GetAssistantByID(ctx context.Context, id string) (*models.Assistant, error) {
	var assistant models.Assistant
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&assistant).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("assistant not found")
		}
		return nil, fmt.Errorf("failed to get assistant: %w", err)
	}

	return &assistant, nil
}

func (r *assistantRepository) GetAssistantByAPIToken(ctx context.Context, apiToken string) (*models.Assistant, error) {
	var assistant models.Assistant
	if err := r.db.WithContext(ctx).Where("api_token = ?", apiToken).First(&assistant).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("assistant not found")
		}
		return nil, fmt.Errorf("failed to get assistant: %w", err)
	}

	return &assistant, nil
}

func (r *assistantRepository) UpdateAssistant(ctx context.Context, id, name, configuration, apiToken string) (*models.Assistant, error) {
	var assistant models.Assistant
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&assistant).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("assistant not found")
		}
		return nil, fmt.Errorf("failed to get assistant: %w", err)
	}

	assistant.Name = name
	assistant.Configuration = configuration
	assistant.APIToken = apiToken
	assistant.UpdatedAt = time.Now()

	if err := r.db.WithContext(ctx).Save(&assistant).Error; err != nil {
		return nil, fmt.Errorf("failed to update assistant: %w", err)
	}

	return &assistant, nil
}

func (r *assistantRepository) DeleteAssistant(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Assistant{}).Error; err != nil {
		return fmt.Errorf("failed to delete assistant: %w", err)
	}

	return nil
}
