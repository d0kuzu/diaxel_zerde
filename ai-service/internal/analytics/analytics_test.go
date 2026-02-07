package analytics

import (
	"context"
	"testing"
	"time"

	"diaxel/database/models"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupAnalyticsTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.Message{})
	assert.NoError(t, err)

	return db
}

func seedTestData(t *testing.T, db *gorm.DB) {
	now := time.Now()
	messages := []models.Message{
		{
			ChatUserID: "user1",
			Role:       "user",
			Content:    "Hello",
			Platform:   "telegram",
			Time:       now.Add(-2 * time.Hour),
		},
		{
			ChatUserID: "user1",
			Role:       "assistant",
			Content:    "Hi there!",
			Platform:   "telegram",
			Time:       now.Add(-2 * time.Hour),
		},
		{
			ChatUserID: "user2",
			Role:       "user",
			Content:    "Help me",
			Platform:   "web",
			Time:       now.Add(-1 * time.Hour),
		},
		{
			ChatUserID: "user2",
			Role:       "assistant",
			Content:    "How can I help?",
			Platform:   "web",
			Time:       now.Add(-1 * time.Hour),
		},
		{
			ChatUserID: "user3",
			Role:       "user",
			Content:    "Another question",
			Platform:   "telegram",
			Time:       now.Add(-30 * time.Minute),
		},
	}

	for _, msg := range messages {
		err := db.Create(&msg).Error
		assert.NoError(t, err)
	}
}

func TestAnalyticsService_GetAnalytics(t *testing.T) {
	db := setupAnalyticsTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	seedTestData(t, db)

	service := NewAnalyticsService(db)

	filter := AnalyticsFilter{
		StartDate: time.Now().Add(-24 * time.Hour),
		EndDate:   time.Now(),
	}

	result, err := service.GetAnalytics(context.Background(), filter)
	assert.NoError(t, err)

	assert.Equal(t, 3, result.TotalChats)
	assert.Equal(t, 3, result.ActiveUsers)
	assert.Equal(t, 0.6, result.EngagementRate)
}

func TestAnalyticsService_GetAnalyticsByPlatform(t *testing.T) {
	db := setupAnalyticsTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	seedTestData(t, db)

	service := NewAnalyticsService(db)

	filter := AnalyticsFilter{
		Platform:  "telegram",
		StartDate: time.Now().Add(-24 * time.Hour),
		EndDate:   time.Now(),
	}

	result, err := service.GetAnalytics(context.Background(), filter)
	assert.NoError(t, err)

	assert.Equal(t, 2, result.TotalChats)
	assert.Equal(t, 2, result.ActiveUsers)
	assert.Equal(t, 0.67, result.EngagementRate)
}

func TestAnalyticsService_GetAnalyticsByAssistant(t *testing.T) {
	db := setupAnalyticsTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	seedTestData(t, db)

	service := NewAnalyticsService(db)

	assistantID := "assistant-123"
	result, err := service.GetAnalyticsByAssistant(context.Background(), assistantID, AnalyticsFilter{
		StartDate: time.Now().Add(-24 * time.Hour),
		EndDate:   time.Now(),
	})

	assert.NoError(t, err)
	assert.Equal(t, assistantID, result.AssistantID)
	assert.Equal(t, 3, result.TotalChats)
	assert.Equal(t, 3, result.ActiveUsers)
}
