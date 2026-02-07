package memory

import (
	"context"
	"sync"

	"github.com/tr1ki/diaxel_zerde_master/database-service/internal/models"
)

type RefreshTokenRepository struct {
	mu     sync.RWMutex
	tokens map[string]*models.RefreshToken
}

func NewRefreshTokenRepository() *RefreshTokenRepository {
	return &RefreshTokenRepository{
		tokens: make(map[string]*models.RefreshToken),
	}
}

func (r *RefreshTokenRepository) Save(ctx context.Context, token *models.RefreshToken) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tokens[token.Token] = token
	return nil
}

func (r *RefreshTokenRepository) FindByToken(ctx context.Context, token string) (*models.RefreshToken, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	rt, ok := r.tokens[token]
	if !ok {
		return nil, nil
	}
	return rt, nil
}

func (r *RefreshTokenRepository) DeleteByUserID(ctx context.Context, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for k, v := range r.tokens {
		if v.UserID == userID {
			delete(r.tokens, k)
		}
	}
	return nil
}

func (r *RefreshTokenRepository) DeleteByToken(ctx context.Context, token string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.tokens, token)
	return nil
}
