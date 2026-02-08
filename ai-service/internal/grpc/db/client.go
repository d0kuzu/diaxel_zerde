package db

import (
	"context"
	"fmt"
	"log"
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

func (c *Client) GetStats() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.AnalyticsRequest{
		AssistantId: "asst_555",
		Platform:    "telegram",
		StartDate:   "2023-01-01",
		EndDate:     "2023-02-01",
	}

	resp, err := c.DB.GetAnalytics(ctx, req)

	if err != nil {
		log.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Всего чатов: %d\n", resp.GetTotalChats())
}
