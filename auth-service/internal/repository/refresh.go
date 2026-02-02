package repository

import "sync"

type RefreshRepo struct {
	mu     sync.Mutex
	tokens map[string]string
}

func NewRefreshRepo() *RefreshRepo {
	return &RefreshRepo{
		tokens: make(map[string]string),
	}
}

func (r *RefreshRepo) Save(token, userID string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tokens[token] = userID
}

func (r *RefreshRepo) Delete(token string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.tokens, token)
}

func (r *RefreshRepo) Get(token string) (string, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	u, ok := r.tokens[token]
	return u, ok
}
