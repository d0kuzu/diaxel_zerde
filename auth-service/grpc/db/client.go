package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	dbpb "auth-service/proto/db"
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

func (c *Client) GetUserByEmail(email string) (*dbpb.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &dbpb.GetUserByEmailRequest{
		Email: email,
	}

	resp, err := c.DB.GetUserByEmail(ctx, req)

	if err != nil {
		log.Fatalln("Ошибка:", err)
		return nil, err
	}

	return resp, nil
}
