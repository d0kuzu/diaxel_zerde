package repository

import (
	"context"
	"fmt"
	"time"

	"diaxel_zerde/database-service/models"

	"github.com/jmoiron/sqlx"
)

type MessageRepository interface {
	SaveMessage(ctx context.Context, chatUserID, role, content, platform string) (*models.Message, error)
	GetMessagesByChatUserID(ctx context.Context, chatUserID string, limit, offset int32) ([]*models.Message, error)
}

type messageRepository struct {
	db *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) SaveMessage(ctx context.Context, chatUserID, role, content, platform string) (*models.Message, error) {
	now := time.Now()

	query := `
		INSERT INTO messages (chat_user_id, role, content, time)
		VALUES ($1, $2, $3, $4)
		RETURNING id, chat_user_id, role, content, time
	`

	var message models.Message
	err := r.db.QueryRowContext(ctx, query, chatUserID, role, content, now).Scan(
		&message.ID,
		&message.ChatUserID,
		&message.Role,
		&message.Content,
		&message.Time,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to save message: %w", err)
	}

	return &message, nil
}

func (r *messageRepository) GetMessagesByChatUserID(ctx context.Context, chatUserID string, limit, offset int32) ([]*models.Message, error) {
	query := `
		SELECT id, chat_user_id, role, content, time
		FROM messages
		WHERE chat_user_id = $1
		ORDER BY time ASC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, chatUserID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages: %w", err)
	}
	defer rows.Close()

	var messages []*models.Message
	for rows.Next() {
		var message models.Message
		err := rows.Scan(
			&message.ID,
			&message.ChatUserID,
			&message.Role,
			&message.Content,
			&message.Time,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan message: %w", err)
		}
		messages = append(messages, &message)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating messages: %w", err)
	}

	return messages, nil
}
