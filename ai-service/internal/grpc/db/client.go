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
		BotToken: token,
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
