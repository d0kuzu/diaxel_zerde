package repository

import (
	"context"

	"github.com/tr1ki/diaxel_zerde_master/database-service/internal/models"

	"github.com/google/uuid"
)

type RefreshTokenRepository interface {
	Save(ctx context.Context, token *models.RefreshToken) error
	FindByToken(ctx context.Context, token string) (*models.RefreshToken, error)
	DeleteByUserID(ctx context.Context, userID uuid.UUID) error
	DeleteByToken(ctx context.Context, token string) error
}
