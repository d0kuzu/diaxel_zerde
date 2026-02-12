package server

import (
	"context"
	"fmt"
	"time"

	"diaxel_zerde/database-service/proto"
	"diaxel_zerde/database-service/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DatabaseServer struct {
	proto.UnimplementedDatabaseServiceServer
	userRepo         repository.UserRepository
	refreshTokenRepo repository.RefreshTokenRepository
	chatRepo         repository.ChatRepository
	messageRepo      repository.MessageRepository
	assistantRepo    repository.AssistantRepository
}

func NewDatabaseServer(userRepo repository.UserRepository, refreshTokenRepo repository.RefreshTokenRepository, chatRepo repository.ChatRepository, messageRepo repository.MessageRepository, assistantRepo repository.AssistantRepository) *DatabaseServer {
	return &DatabaseServer{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		chatRepo:         chatRepo,
		messageRepo:      messageRepo,
		assistantRepo:    assistantRepo,
	}
}

func (s *DatabaseServer) CreateAssistant(ctx context.Context, req *proto.CreateAssistantRequest) (*proto.AssistantResponse, error) {
	assistant, err := s.assistantRepo.CreateAssistant(ctx, req.Name, req.ApiToken, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create assistant: %v", err)
	}

	return &proto.AssistantResponse{
		Id:        assistant.ID,
		Name:      assistant.Name,
		ApiToken:  assistant.APIToken,
		UserId:    assistant.UserID,
		CreatedAt: assistant.CreatedAt.Format(time.RFC3339),
		UpdatedAt: assistant.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *DatabaseServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.UserResponse, error) {
	user, err := s.userRepo.CreateUser(ctx, req.Email, req.PasswordHash, req.Role)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	var email, passwordHash, role string
	if user.Email != nil {
		email = *user.Email
	}
	if user.PasswordHash != nil {
		passwordHash = *user.PasswordHash
	}
	if user.Role != nil {
		role = *user.Role
	}

	return &proto.UserResponse{
		Id:           user.ID,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
		CreatedAt:    user.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *DatabaseServer) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.UserResponse, error) {
	user, err := s.userRepo.GetUserByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	var email, passwordHash, role string
	if user.Email != nil {
		email = *user.Email
	}
	if user.PasswordHash != nil {
		passwordHash = *user.PasswordHash
	}
	if user.Role != nil {
		role = *user.Role
	}

	return &proto.UserResponse{
		Id:           user.ID,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
		CreatedAt:    user.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *DatabaseServer) GetUserByEmail(ctx context.Context, req *proto.GetUserByEmailRequest) (*proto.UserResponse, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	var email, passwordHash, role string
	if user.Email != nil {
		email = *user.Email
	}
	if user.PasswordHash != nil {
		passwordHash = *user.PasswordHash
	}
	if user.Role != nil {
		role = *user.Role
	}

	return &proto.UserResponse{
		Id:           user.ID,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
		CreatedAt:    user.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *DatabaseServer) SaveRefreshToken(ctx context.Context, req *proto.SaveRefreshTokenRequest) (*proto.SaveRefreshTokenResponse, error) {
	expiresAt, err := time.Parse(time.RFC3339, req.ExpiresAt)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid expires_at format: %v", err)
	}

	err = s.refreshTokenRepo.SaveRefreshToken(ctx, req.UserId, req.TokenHash, expiresAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to save refresh token: %v", err)
	}

	return &proto.SaveRefreshTokenResponse{
		Success: true,
	}, nil
}

func (s *DatabaseServer) GetRefreshToken(ctx context.Context, req *proto.GetRefreshTokenRequest) (*proto.RefreshTokenResponse, error) {
	token, err := s.refreshTokenRepo.GetRefreshToken(ctx, req.TokenHash)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "refresh token not found: %v", err)
	}

	return &proto.RefreshTokenResponse{
		Id:        token.ID,
		UserId:    token.UserID,
		TokenHash: token.TokenHash,
		ExpiresAt: token.ExpiresAt.Format(time.RFC3339),
		CreatedAt: token.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *DatabaseServer) DeleteRefreshToken(ctx context.Context, req *proto.DeleteRefreshTokenRequest) (*proto.DeleteRefreshTokenResponse, error) {
	err := s.refreshTokenRepo.DeleteRefreshToken(ctx, req.TokenHash)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete refresh token: %v", err)
	}

	return &proto.DeleteRefreshTokenResponse{
		Success: true,
	}, nil
}

func (s *DatabaseServer) GetAnalytics(ctx context.Context, req *proto.AnalyticsRequest) (*proto.AnalyticsResponse, error) {
	// TODO: Implement analytics logic
	return &proto.AnalyticsResponse{
		AssistantId:    req.AssistantId,
		TotalChats:     0,
		ActiveUsers:    0,
		EngagementRate: 0.0,
	}, nil
}

func (s *DatabaseServer) GetAnalyticsByAssistant(ctx context.Context, req *proto.AnalyticsByAssistantRequest) (*proto.AnalyticsResponse, error) {
	// TODO: Implement analytics by assistant logic
	return &proto.AnalyticsResponse{
		AssistantId:    req.AssistantId,
		TotalChats:     0,
		ActiveUsers:    0,
		EngagementRate: 0.0,
	}, nil
}

func (s *DatabaseServer) CreateChat(ctx context.Context, req *proto.CreateChatRequest) (*proto.ChatResponse, error) {
	// Verify that the assistant exists before creating a chat
	_, err := s.assistantRepo.GetAssistantByID(ctx, req.AssistantId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "assistant not found: %v", err)
	}

	chat, err := s.chatRepo.CreateChat(ctx, req.AssistantId, req.CustomerId, req.Platform)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create chat: %v", err)
	}

	return &proto.ChatResponse{
		Id:          chat.ID,
		AssistantId: chat.AssistantID,
		CustomerId: func() string {
			if chat.CustomerID != nil {
				return *chat.CustomerID
			}
			return ""
		}(),
		Platform:  req.Platform,
		CreatedAt: chat.StartedAt.Format(time.RFC3339),
		UpdatedAt: chat.StartedAt.Format(time.RFC3339),
	}, nil
}

func (s *DatabaseServer) SaveMessage(ctx context.Context, req *proto.SaveMessageRequest) (*proto.MessageResponse, error) {
	message, err := s.messageRepo.SaveMessage(ctx, req.GetChatId(), req.GetRole(), req.GetContent(), req.GetPlatform())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to save message: %v", err)
	}

	// Update message count for the chat
	if err := s.chatRepo.UpdateMessageCount(ctx, req.GetChatId()); err != nil {
		// Log error but don't fail the message save
		// TODO: Add proper logging
	}

	return &proto.MessageResponse{
		Id:        fmt.Sprintf("%d", message.ID),
		ChatId:    message.ChatID,
		Role:      message.Role,
		Content:   message.Content,
		Platform:  req.GetPlatform(),
		CreatedAt: message.Time.Format(time.RFC3339),
	}, nil
}

func (s *DatabaseServer) GetChatMessages(ctx context.Context, req *proto.GetChatMessagesRequest) (*proto.MessagesResponse, error) {
	messages, err := s.messageRepo.GetMessagesByChatID(ctx, req.GetChatId(), req.GetLimit(), req.GetOffset())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get messages: %v", err)
	}

	var protoMessages []*proto.MessageResponse
	for _, msg := range messages {
		protoMessages = append(protoMessages, &proto.MessageResponse{
			Id:        fmt.Sprintf("%d", msg.ID),
			ChatId:    msg.ChatID,
			Role:      msg.Role,
			Content:   msg.Content,
			Platform:  "", // Platform not stored in message model
			CreatedAt: msg.Time.Format(time.RFC3339),
		})
	}

	return &proto.MessagesResponse{
		Messages: protoMessages,
	}, nil
}

func (s *DatabaseServer) GetAllChatMessages(ctx context.Context, req *proto.GetAllChatMessagesRequest) (*proto.MessagesResponse, error) {
	messages, err := s.messageRepo.GetAllChatMessages(ctx, req.ChatId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all messages: %v", err)
	}

	var protoMessages []*proto.MessageResponse
	for _, msg := range messages {
		protoMessages = append(protoMessages, &proto.MessageResponse{
			Id:        fmt.Sprintf("%d", msg.ID),
			ChatId:    msg.ChatID,
			Role:      msg.Role,
			Content:   msg.Content,
			Platform:  "", // Platform not stored in message model
			CreatedAt: msg.Time.Format(time.RFC3339),
		})
	}

	return &proto.MessagesResponse{
		Messages: protoMessages,
	}, nil
}

func (s *DatabaseServer) GetChatPagesCount(ctx context.Context, req *proto.GetChatPagesCountRequest) (*proto.ChatPagesCountResponse, error) {
	pagesCount, err := s.chatRepo.GetChatPagesCount(ctx, req.AssistantId, req.ChatsPerPage)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get chat pages count: %v", err)
	}

	return &proto.ChatPagesCountResponse{
		PagesCount: pagesCount,
	}, nil
}

func (s *DatabaseServer) GetChatPage(ctx context.Context, req *proto.GetChatPageRequest) (*proto.ChatsResponse, error) {
	chats, err := s.chatRepo.GetChatPage(ctx, req.AssistantId, req.Page, req.ChatsPerPage)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get chat page: %v", err)
	}

	var protoChats []*proto.ChatResponse
	for _, chat := range chats {
		customerId := ""
		if chat.CustomerID != nil {
			customerId = *chat.CustomerID
		}

		protoChats = append(protoChats, &proto.ChatResponse{
			Id:           chat.ID,
			AssistantId:  chat.AssistantID,
			CustomerId:   customerId,
			Platform:     "", // Platform not stored in chat model
			CreatedAt:    chat.CreatedAt.Format(time.RFC3339),
			UpdatedAt:    chat.UpdatedAt.Format(time.RFC3339),
			MessageCount: chat.MessageCount,
		})
	}

	return &proto.ChatsResponse{
		Chats: protoChats,
	}, nil
}

func (s *DatabaseServer) SearchChatsByUser(ctx context.Context, req *proto.SearchChatsByUserRequest) (*proto.SearchChatsResponse, error) {
	chats, totalCount, err := s.chatRepo.SearchChatsByUser(ctx, req.AssistantId, req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to search chats: %v", err)
	}

	var protoChats []*proto.ChatResponse
	for _, chat := range chats {
		customerId := ""
		if chat.CustomerID != nil {
			customerId = *chat.CustomerID
		}

		protoChats = append(protoChats, &proto.ChatResponse{
			Id:           chat.ID,
			AssistantId:  chat.AssistantID,
			CustomerId:   customerId,
			Platform:     "", // Platform not stored in chat model
			CreatedAt:    chat.CreatedAt.Format(time.RFC3339),
			UpdatedAt:    chat.UpdatedAt.Format(time.RFC3339),
			MessageCount: chat.MessageCount,
		})
	}

	return &proto.SearchChatsResponse{
		Chats:      protoChats,
		TotalCount: totalCount,
	}, nil
}

func (s *DatabaseServer) GetAssistant(ctx context.Context, req *proto.GetAssistantRequest) (*proto.AssistantResponse, error) {
	assistant, err := s.assistantRepo.GetAssistantByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "assistant not found: %v", err)
	}

	return &proto.AssistantResponse{
		Id:        assistant.ID,
		Name:      assistant.Name,
		ApiToken:  assistant.APIToken,
		UserId:    assistant.UserID,
		CreatedAt: assistant.CreatedAt.Format(time.RFC3339),
		UpdatedAt: assistant.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *DatabaseServer) GetAssistantByAPIToken(ctx context.Context, req *proto.GetAssistantByAPITokenRequest) (*proto.AssistantResponse, error) {
	assistant, err := s.assistantRepo.GetAssistantByAPIToken(ctx, req.ApiToken)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "assistant not found: %v", err)
	}

	return &proto.AssistantResponse{
		Id:        assistant.ID,
		Name:      assistant.Name,
		ApiToken:  assistant.APIToken,
		UserId:    assistant.UserID,
		CreatedAt: assistant.CreatedAt.Format(time.RFC3339),
		UpdatedAt: assistant.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *DatabaseServer) UpdateAssistant(ctx context.Context, req *proto.UpdateAssistantRequest) (*proto.AssistantResponse, error) {
	assistant, err := s.assistantRepo.UpdateAssistant(ctx, req.Id, req.Name, "", req.ApiToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update assistant: %v", err)
	}

	return &proto.AssistantResponse{
		Id:        assistant.ID,
		Name:      assistant.Name,
		ApiToken:  assistant.APIToken,
		UserId:    assistant.UserID,
		CreatedAt: assistant.CreatedAt.Format(time.RFC3339),
		UpdatedAt: assistant.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *DatabaseServer) DeleteAssistant(ctx context.Context, req *proto.DeleteAssistantRequest) (*proto.DeleteAssistantResponse, error) {
	err := s.assistantRepo.DeleteAssistant(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete assistant: %v", err)
	}

	return &proto.DeleteAssistantResponse{
		Success: true,
	}, nil
}
