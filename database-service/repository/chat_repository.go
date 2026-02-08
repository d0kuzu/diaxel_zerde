package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"diaxel_zerde/database-service/models"

	"github.com/jmoiron/sqlx"
)

type ChatRepository interface {
	CreateChat(ctx context.Context, assistantID, customerID, platform string) (*models.Chat, error)
	GetChatByID(ctx context.Context, id string) (*models.Chat, error)
	GetChatsByCustomerID(ctx context.Context, customerID string) ([]*models.Chat, error)
}

type chatRepository struct {
	db *sqlx.DB
}

func NewChatRepository(db *sqlx.DB) ChatRepository {
	return &chatRepository{db: db}
}

func (r *chatRepository) CreateChat(ctx context.Context, assistantID, customerID, platform string) (*models.Chat, error) {
	now := time.Now()

	query := `
		INSERT INTO chats (assistant_id, customer_id, platform, started_at)
		VALUES ($1, $2, $3, $4)
		RETURNING assistant_id, customer_id, started_at
	`

	var chat models.Chat
	err := r.db.QueryRowContext(ctx, query, assistantID, customerID, platform, now).Scan(
		&chat.AssistantID,
		&chat.CustomerID,
		&chat.StartedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create chat: %w", err)
	}

	return &chat, nil
}

func (r *chatRepository) GetChatByID(ctx context.Context, id string) (*models.Chat, error) {
	query := `
		SELECT assistant_id, customer_id, started_at
		FROM chats
		WHERE id = $1
	`

	var chat models.Chat
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&chat.AssistantID,
		&chat.CustomerID,
		&chat.StartedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("chat not found")
		}
		return nil, fmt.Errorf("failed to get chat: %w", err)
	}

	return &chat, nil
}

func (r *chatRepository) GetChatsByCustomerID(ctx context.Context, customerID string) ([]*models.Chat, error) {
	query := `
		SELECT assistant_id, customer_id, started_at
		FROM chats
		WHERE customer_id = $1
		ORDER BY started_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get chats: %w", err)
	}
	defer rows.Close()

	var chats []*models.Chat
	for rows.Next() {
		var chat models.Chat
		err := rows.Scan(
			&chat.AssistantID,
			&chat.CustomerID,
			&chat.StartedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan chat: %w", err)
		}
		chats = append(chats, &chat)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating chats: %w", err)
	}

	return chats, nil
}
