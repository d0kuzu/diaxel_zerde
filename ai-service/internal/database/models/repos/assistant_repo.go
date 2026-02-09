package repos

import (
	"diaxel/internal/database"
	"diaxel/internal/database/models"
	"errors"
	"gorm.io/gorm"
)

func CreateAssistant(name, token, userID string) (*models.Assistant, error) {
	db := database.GetDB()
	
	assistant := &models.Assistant{
		Name:     name,
		BotToken: token,
		UserID:   userID,
	}
	
	if err := db.Create(assistant).Error; err != nil {
		return nil, err
	}
	
	return assistant, nil
}

func GetAssistant(assistantID string) (*models.Assistant, error) {
	db := database.GetDB()
	
	var assistant models.Assistant
	if err := db.Where("id = ?", assistantID).First(&assistant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("assistant not found")
		}
		return nil, err
	}
	
	return &assistant, nil
}

func GetAssistantByToken(token string) (*models.Assistant, error) {
	db := database.GetDB()
	
	var assistant models.Assistant
	if err := db.Where("bot_token = ?", token).First(&assistant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("assistant not found")
		}
		return nil, err
	}
	
	return &assistant, nil
}
