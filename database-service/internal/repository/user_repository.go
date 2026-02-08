package repository

import (
	"context"

	"github.com/tr1ki/diaxel_zerde_master/database-service/internal/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}
