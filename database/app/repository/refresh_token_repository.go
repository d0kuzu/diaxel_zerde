package repository

import (
	"context"

	"github.com/google/uuid"
	"DIAXEL-ZERDE-MASTER/database/app/models"
)

type RefreshTokenRepository interface {
	Save(ctx context.Context, token *models.RefreshToken) error
	FindByToken(ctx context.Context, token string) (*models.RefreshToken, error)
	DeleteByUserID(ctx context.Context, userID uuid.UUID) error
	DeleteByToken(ctx context.Context, token string) error
}
