package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	dbpb "auth-service/proto"
)

type Client struct {
	conn *grpc.ClientConn
	DB   dbpb.DatabaseServiceClient
}

func New(address string) (*Client, error) {
	conn, err := grpc.DialContext(
		context.Background(),
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

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

func (c *Client) GetUserByEmail(email string) (*dbpb.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.GetUserByEmailRequest{
		Email: email,
	}

	resp, err := c.DB.GetUserByEmail(ctx, req)

	if err != nil {
		log.Println("Ошибка:", err)
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetUserByID(userID string) (*dbpb.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.GetUserRequest{
		Id: userID,
	}

	resp, err := c.DB.GetUser(ctx, req)
	if err != nil {
		log.Println("Ошибка получения пользователя:", err)
		return nil, err
	}

	return resp, nil
}

func (c *Client) CreateUser(email, passwordHash, role string) (*dbpb.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.CreateUserRequest{
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
	}

	resp, err := c.DB.CreateUser(ctx, req)
	if err != nil {
		log.Println("Ошибка создания пользователя:", err)
		return nil, err
	}

	return resp, nil
}

func (c *Client) SaveRefreshToken(token, userID string, expiresAt time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.SaveRefreshTokenRequest{
		TokenHash: token,
		UserId:    userID,
		ExpiresAt: expiresAt.Format(time.RFC3339),
	}

	_, err := c.DB.SaveRefreshToken(ctx, req)
	if err != nil {
		log.Println("Ошибка сохранения refresh токена:", err)
		return err
	}

	return nil
}

func (c *Client) GetRefreshToken(token string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.GetRefreshTokenRequest{
		TokenHash: token,
	}

	resp, err := c.DB.GetRefreshToken(ctx, req)
	if err != nil {
		log.Println("Ошибка получения refresh токена:", err)
		return "", err
	}

	return resp.UserId, nil
}

func (c *Client) DeleteRefreshToken(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.DeleteRefreshTokenRequest{
		TokenHash: token,
	}

	_, err := c.DB.DeleteRefreshToken(ctx, req)
	if err != nil {
		log.Println("Ошибка удаления refresh токена:", err)
		return err
	}

	return nil
}

func (c *Client) CreateAssistant(name, botToken, userID string) (*dbpb.AssistantResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.CreateAssistantRequest{
		Name:     name,
		BotToken: botToken,
		UserId:   userID,
	}

	resp, err := c.DB.CreateAssistant(ctx, req)
	if err != nil {
		log.Println("Ошибка создания ассистента:", err)
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetAssistant(assistantID string) (*dbpb.AssistantResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.GetAssistantRequest{
		Id: assistantID,
	}

	resp, err := c.DB.GetAssistant(ctx, req)
	if err != nil {
		log.Println("Ошибка получения ассистента:", err)
		return nil, err
	}

	return resp, nil
}

func (c *Client) UpdateAssistant(assistantID, name, botToken, userID string) (*dbpb.AssistantResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.UpdateAssistantRequest{
		Id:       assistantID,
		Name:     name,
		BotToken: botToken,
		UserId:   userID,
	}

	resp, err := c.DB.UpdateAssistant(ctx, req)
	if err != nil {
		log.Println("Ошибка обновления ассистента:", err)
		return nil, err
	}

	return resp, nil
}
