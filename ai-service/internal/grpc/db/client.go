package db

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	dbpb "diaxel/proto/db"
)

type Client struct {
	conn *grpc.ClientConn
	DB   dbpb.DatabaseServiceClient
}

func New(address string) (*Client, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания gRPC клиента: %w", err)
	}

	dbClient := dbpb.NewDatabaseServiceClient(conn)

	return &Client{
		conn: conn,
		DB:   dbClient,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) CreateAssistant(name, token, userId string) (*dbpb.AssistantResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.CreateAssistantRequest{
		Name:     name,
		ApiToken: token,
		UserId:   userId,
	}

	resp, err := c.DB.CreateAssistant(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetAssistant(id string) (*dbpb.AssistantResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.GetAssistantRequest{
		Id: id,
	}

	resp, err := c.DB.GetAssistant(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetAssistantByAPIToken(apiToken string) (*dbpb.AssistantResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.GetAssistantByAPITokenRequest{
		ApiToken: apiToken,
	}

	resp, err := c.DB.GetAssistantByAPIToken(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) CreateChat(assistantID, customerID, platform string) (*dbpb.ChatResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.CreateChatRequest{
		AssistantId: assistantID,
		CustomerId:  customerID,
		Platform:    platform,
	}

	return c.DB.CreateChat(ctx, req)
}

func (c *Client) SaveMessage(chatID, role, content, platform string) (*dbpb.MessageResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.SaveMessageRequest{
		ChatId:   chatID,
		Role:     role,
		Content:  content,
		Platform: platform,
	}

	return c.DB.SaveMessage(ctx, req)
}

func (c *Client) GetAllChatMessages(chatID string) ([]*dbpb.MessageResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.GetAllChatMessagesRequest{
		ChatId: chatID,
	}

	resp, err := c.DB.GetAllChatMessages(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Messages, nil
}

func (c *Client) GetLatestChatByCustomer(assistantID, customerID string) (*dbpb.ChatResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.GetLatestChatByCustomerRequest{
		AssistantId: assistantID,
		CustomerId:  customerID,
	}

	return c.DB.GetLatestChatByCustomer(ctx, req)
}

func (c *Client) GetChatMessages(chatID string, limit, offset int32) ([]*dbpb.MessageResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.GetChatMessagesRequest{
		ChatId: chatID,
		Limit:  limit,
		Offset: offset,
	}

	resp, err := c.DB.GetChatMessages(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Messages, nil
}

func (c *Client) GetChatPage(assistantID string, page, chatsPerPage int32) ([]*dbpb.ChatResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.GetChatPageRequest{
		AssistantId:  assistantID,
		Page:         page,
		ChatsPerPage: chatsPerPage,
	}

	resp, err := c.DB.GetChatPage(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Chats, nil
}

func (c *Client) GetChatPagesCount(assistantID string, chatsPerPage int32) (int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.GetChatPagesCountRequest{
		AssistantId:  assistantID,
		ChatsPerPage: chatsPerPage,
	}

	resp, err := c.DB.GetChatPagesCount(ctx, req)
	if err != nil {
		return 0, err
	}

	return resp.PagesCount, nil
}

func (c *Client) SearchChatsByCustomer(assistantID, search string) ([]*dbpb.ChatResponse, int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.SearchChatsByCustomerRequest{
		AssistantId: assistantID,
		Search:      search,
	}

	resp, err := c.DB.SearchChatsByCustomer(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	return resp.Chats, resp.TotalCount, nil
}
