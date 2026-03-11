package db

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	dbpb "api-gateway/proto/db"
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

func (c *Client) GetAssistantByAPIToken(token string) (*dbpb.AssistantResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.GetAssistantByAPITokenRequest{
		ApiToken: token,
	}

	fmt.Printf("[API-GATEWAY] Sending gRPC request GetAssistantByAPIToken to DB for token: '%s'\n", token)
	resp, err := c.DB.GetAssistantByAPIToken(ctx, req)

	if err != nil {
		fmt.Printf("[API-GATEWAY] Error from DB gRPC: %v\n", err)
		return nil, err
	}

	fmt.Printf("[API-GATEWAY] DB returned assistant successfully: %s\n", resp.Id)
	return resp, nil
}
