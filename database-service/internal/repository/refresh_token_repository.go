package repository

import (
	"context"

	"diaxel/database/app/models"

	"github.com/google/uuid"
)

type RefreshTokenRepository interface {
	Save(ctx context.Context, token *models.RefreshToken) error
	FindByToken(ctx context.Context, token string) (*models.RefreshToken, error)
	DeleteByUserID(ctx context.Context, userID uuid.UUID) error
	DeleteByToken(ctx context.Context, token string) error
}
