package infrastructure

import (
	"context"
	"fmt"
	"sync"

	"morning-call/internal/domain"
	"morning-call/internal/repository"
)

// inMemoryUserRepository は UserRepository のインメモリ実装です
type inMemoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]*domain.User
}

// NewInMemoryUserRepository は新しい inMemoryUserRepository を生成します
func NewInMemoryUserRepository() repository.UserRepository {
	return &inMemoryUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *inMemoryUserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user not found: %s", id)
	}
	return user, nil
}

func (r *inMemoryUserRepository) Create(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[user.ID]; ok {
		return fmt.Errorf("user already exists: %s", user.ID)
	}
	r.users[user.ID] = user
	return nil
}
