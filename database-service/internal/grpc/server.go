package grpc

import (
	"context"
	"log"
	"time"

	"github.com/tr1ki/diaxel_zerde_master/database-service/internal/models"
	pb "github.com/tr1ki/diaxel_zerde_master/database-service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedDatabaseServiceServer
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

func (s *Server) GetAnalytics(ctx context.Context, req *pb.AnalyticsRequest) (*pb.AnalyticsResponse, error) {
	log.Printf("GetAnalytics called with: %+v", req)

	query := s.db.WithContext(ctx).Model(&models.Message{})

	if req.Platform != "" {
		query = query.Where("platform = ?", req.Platform)
	}

	startDate := time.Now().AddDate(0, 0, -7)
	if req.StartDate != "" {
		if parsed, err := time.Parse("2006-01-02", req.StartDate); err == nil {
			startDate = parsed
		}
	}

	endDate := time.Now()
	if req.EndDate != "" {
		if parsed, err := time.Parse("2006-01-02", req.EndDate); err == nil {
			endDate = parsed
		}
	}

	query = query.Where("time BETWEEN ? AND ?", startDate, endDate)

	var totalChats int64
	err := query.Select("COUNT(DISTINCT chat_user_id)").Scan(&totalChats).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get total chats: %v", err)
	}

	var activeUsers int64
	err = query.Where("role = ?", "user").Select("COUNT(DISTINCT chat_user_id)").Scan(&activeUsers).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get active users: %v", err)
	}

	var totalMessages int64
	err = query.Select("COUNT(*)").Scan(&totalMessages).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get total messages: %v", err)
	}

	var userMessages int64
	err = query.Where("role = ?", "user").Select("COUNT(*)").Scan(&userMessages).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get user messages: %v", err)
	}

	var engagementRate float64
	if totalMessages > 0 {
		engagementRate = float64(userMessages) / float64(totalMessages)
	}

	response := &pb.AnalyticsResponse{
		AssistantId:    req.AssistantId,
		TotalChats:     int32(totalChats),
		ActiveUsers:    int32(activeUsers),
		EngagementRate: engagementRate,
	}

	log.Printf("GetAnalytics response: %+v", response)
	return response, nil
}

func (s *Server) GetAnalyticsByAssistant(ctx context.Context, req *pb.AnalyticsByAssistantRequest) (*pb.AnalyticsResponse, error) {
	log.Printf("GetAnalyticsByAssistant called with: %+v", req)

	analyticsReq := &pb.AnalyticsRequest{
		AssistantId: req.AssistantId,
		Platform:    req.Platform,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}

	return s.GetAnalytics(ctx, analyticsReq)
}
