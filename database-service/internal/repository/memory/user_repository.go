package memory

import (
	"context"
	"sync"

	"diaxel/database/app/models"

	"github.com/google/uuid"
)

type UserRepository struct {
	mu    sync.RWMutex
	users map[uuid.UUID]*models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[uuid.UUID]*models.User),
	}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return nil, nil
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, nil
}
