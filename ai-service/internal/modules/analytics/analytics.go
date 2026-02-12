package analytics

import (
	"context"
	"log"
	"time"

	"diaxel/internal/grpc/db"
)

type AnalyticsService struct {
	dbClient *db.Client
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

func NewAnalyticsService(dbClient *db.Client) *AnalyticsService {
	return &AnalyticsService{dbClient: dbClient}
}

func (s *AnalyticsService) GetAnalytics(ctx context.Context, filter AnalyticsFilter) (*AnalyticsResponse, error) {
	log.Printf("Getting analytics with filter: %+v", filter)

	response := &AnalyticsResponse{
		AssistantID:    filter.AssistantID,
		TotalChats:     0,
		ActiveUsers:    0,
		EngagementRate: 0.0,
	}

	log.Printf("Analytics result: %+v", response)
	return response, nil
}

func (s *AnalyticsService) GetAnalyticsByAssistant(ctx context.Context, assistantID string, filter AnalyticsFilter) (*AnalyticsResponse, error) {
	filter.AssistantID = assistantID
	return s.GetAnalytics(ctx, filter)
}
