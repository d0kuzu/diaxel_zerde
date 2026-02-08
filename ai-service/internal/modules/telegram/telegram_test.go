package telegram

import (
	"context"
	"testing"
	"time"

	"diaxel/internal/config"
	"diaxel/internal/database/models"
	"diaxel/internal/modules/llm"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MockLLMClient struct {
	mock.Mock
}

func (m *MockLLMClient) Conversation(ctx interface{}, userID, message string) (string, error) {
	args := m.Called(ctx, userID, message)
	return args.String(0), args.Error(1)
}

func (m *MockLLMClient) toLLMClient() *llm.Client {
	return &llm.Client{}
}

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.Message{})
	assert.NoError(t, err)

	return db
}

func TestTelegramClient_HandleWebhook(t *testing.T) {
	db := setupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	mockLLM := new(MockLLMClient)
	mockLLM.On("Conversation", mock.Anything, mock.AnythingOfType("string"), "Hello").Return("Hi there!", nil)

	cfg := &config.Settings{
		TelegramBotToken:      "test_token",
		TelegramWebhookSecret: "test_secret",
	}

	client := NewClient(mockLLM.toLLMClient(), cfg)

	update := TelegramUpdate{
		UpdateID: 123,
		Message: struct {
			MessageID int `json:"message_id"`
			From      struct {
				ID        int64  `json:"id"`
				FirstName string `json:"first_name"`
				Username  string `json:"username"`
			} `json:"from"`
			Chat struct {
				ID        int64  `json:"id"`
				FirstName string `json:"first_name"`
				Username  string `json:"username"`
				Type      string `json:"type"`
			} `json:"chat"`
			Date int64  `json:"date"`
			Text string `json:"text"`
		}{
			MessageID: 456,
			From: struct {
				ID        int64  `json:"id"`
				FirstName string `json:"first_name"`
				Username  string `json:"username"`
			}{
				ID:        789,
				FirstName: "Test",
				Username:  "testuser",
			},
			Chat: struct {
				ID        int64  `json:"id"`
				FirstName string `json:"first_name"`
				Username  string `json:"username"`
				Type      string `json:"type"`
			}{
				ID:        789,
				FirstName: "Test",
				Username:  "testuser",
				Type:      "private",
			},
			Date: time.Now().Unix(),
			Text: "Hello",
		},
	}

	err := client.HandleWebhook(context.Background(), update)
	assert.NoError(t, err)

	var messages []models.Message
	err = db.Find(&messages).Error
	assert.NoError(t, err)
	assert.Len(t, messages, 2)

	assert.Equal(t, "user", messages[0].Role)
	assert.Equal(t, "Hello", messages[0].Content)
	assert.Equal(t, "telegram", messages[0].Platform)

	assert.Equal(t, "assistant", messages[1].Role)
	assert.Equal(t, "Hi there!", messages[1].Content)
	assert.Equal(t, "telegram", messages[1].Platform)

	mockLLM.AssertExpectations(t)
}

func TestTelegramClient_ValidateWebhookSecret(t *testing.T) {
	cfg := &config.Settings{
		TelegramWebhookSecret: "test_secret",
	}

	client := NewClient((&MockLLMClient{}).toLLMClient(), cfg)

	assert.True(t, client.ValidateWebhookSecret("test_secret"))
	assert.False(t, client.ValidateWebhookSecret("wrong_secret"))
	assert.False(t, client.ValidateWebhookSecret(""))
}
