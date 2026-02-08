package analytics

import (
	"context"
	"fmt"
	"log"
	"time"

	"diaxel/database/models"
	"gorm.io/gorm"
)

type AnalyticsService struct {
	db *gorm.DB
}

type AnalyticsResponse struct {
	AssistantID    string  `json:"assistant_id"`
	TotalChats     int     `json:"total_chats"`
	ActiveUsers    int     `json:"active_users"`
	EngagementRate float64 `json:"engagement_rate"`
}

type AnalyticsFilter struct {
	AssistantID string    `json:"assistant_id,omitempty"`
	Platform    string    `json:"platform,omitempty"`
	StartDate   time.Time `json:"start_date,omitempty"`
	EndDate     time.Time `json:"end_date,omitempty"`
}

func NewAnalyticsService(db *gorm.DB) *AnalyticsService {
	return &AnalyticsService{db: db}
}

func (s *AnalyticsService) GetAnalytics(ctx context.Context, filter AnalyticsFilter) (*AnalyticsResponse, error) {
	log.Printf("Getting analytics with filter: %+v", filter)

	query := s.db.WithContext(ctx).Model(&models.Message{})

	if filter.Platform != "" {
		query = query.Where("platform = ?", filter.Platform)
	}

	if filter.StartDate.IsZero() {
		filter.StartDate = time.Now().AddDate(0, 0, -7)
	}

	if filter.EndDate.IsZero() {
		filter.EndDate = time.Now()
	}

	query = query.Where("time BETWEEN ? AND ?", filter.StartDate, filter.EndDate)

	var totalChats int64
	err := query.Select("COUNT(DISTINCT chat_user_id)").Scan(&totalChats).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get total chats: %w", err)
	}

	var activeUsers int64
	err = query.Where("role = ?", "user").Select("COUNT(DISTINCT chat_user_id)").Scan(&activeUsers).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get active users: %w", err)
	}

	var totalMessages int64
	err = query.Select("COUNT(*)").Scan(&totalMessages).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get total messages: %w", err)
	}

	var userMessages int64
	err = query.Where("role = ?", "user").Select("COUNT(*)").Scan(&userMessages).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user messages: %w", err)
	}

	var engagementRate float64
	if totalMessages > 0 {
		engagementRate = float64(userMessages) / float64(totalMessages)
	}

	response := &AnalyticsResponse{
		AssistantID:    filter.AssistantID,
		TotalChats:     int(totalChats),
		ActiveUsers:    int(activeUsers),
		EngagementRate: engagementRate,
	}

	log.Printf("Analytics result: %+v", response)
	return response, nil
}

func (s *AnalyticsService) GetAnalyticsByAssistant(ctx context.Context, assistantID string, filter AnalyticsFilter) (*AnalyticsResponse, error) {
	filter.AssistantID = assistantID
	return s.GetAnalytics(ctx, filter)
}
