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
	CreateAssistant(ctx context.Context, name, apiToken, userID, telegramBotToken, assistantType string) (*models.Assistant, error)
	GetAssistantByID(ctx context.Context, id string) (*models.Assistant, error)
	GetAssistantByAPIToken(ctx context.Context, apiToken string) (*models.Assistant, error)
	GetAssistantsByUserID(ctx context.Context, userID string) ([]models.Assistant, error)
	UpdateAssistant(ctx context.Context, id, name, configuration, apiToken, telegramBotToken, assistantType string) (*models.Assistant, error)
	DeleteAssistant(ctx context.Context, id string) error
}

type assistantRepository struct {
	db *gorm.DB
}

func NewAssistantRepository(db *gorm.DB) AssistantRepository {
	return &assistantRepository{db: db}
}

func (r *assistantRepository) CreateAssistant(ctx context.Context, name, apiToken, userID, telegramBotToken, assistantType string) (*models.Assistant, error) {
	if assistantType == "" {
		assistantType = "api"
	}
	assistant := models.Assistant{
		ID:               uuid.New().String(),
		Name:             name,
		Configuration:    "",
		APIToken:         apiToken,
		TelegramBotToken: telegramBotToken,
		Type:             assistantType,
		UserID:           userID,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
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
	fmt.Printf("--- START GetAssistantByAPIToken ---\n")
	fmt.Printf("Searching for apiToken: '%s'\n", apiToken)
	
	var assistant models.Assistant
	if err := r.db.WithContext(ctx).Where("api_token = ?", apiToken).First(&assistant).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("Result: Assistant NOT FOUND for token: '%s'\n", apiToken)
			fmt.Printf("--- END GetAssistantByAPIToken ---\n")
			return nil, fmt.Errorf("assistant not found")
		}
		fmt.Printf("Result: DB Error finding assistant: %v\n", err)
		fmt.Printf("--- END GetAssistantByAPIToken ---\n")
		return nil, fmt.Errorf("failed to get assistant: %w", err)
	}

	fmt.Printf("Result: FOUND Assistant! ID: '%s', Name: '%s'\n", assistant.ID, assistant.Name)
	fmt.Printf("--- END GetAssistantByAPIToken ---\n")
	return &assistant, nil
}

func (r *assistantRepository) UpdateAssistant(ctx context.Context, id, name, configuration, apiToken, telegramBotToken, assistantType string) (*models.Assistant, error) {
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
	assistant.TelegramBotToken = telegramBotToken
	if assistantType != "" {
		assistant.Type = assistantType
	}
	assistant.UpdatedAt = time.Now()

	if err := r.db.WithContext(ctx).Save(&assistant).Error; err != nil {
		return nil, fmt.Errorf("failed to update assistant: %w", err)
	}

	return &assistant, nil
}

func (r *assistantRepository) GetAssistantsByUserID(ctx context.Context, userID string) ([]models.Assistant, error) {
	var assistants []models.Assistant
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&assistants).Error; err != nil {
		return nil, fmt.Errorf("failed to get assistants by user id: %w", err)
	}

	return assistants, nil
}

func (r *assistantRepository) DeleteAssistant(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Assistant{}).Error; err != nil {
		return fmt.Errorf("failed to delete assistant: %w", err)
	}

	return nil
}
