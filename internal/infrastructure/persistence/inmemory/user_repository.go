package inmemory

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
	users map[domain.UserID]*domain.User
}

// NewInMemoryUserRepository は新しい inMemoryUserRepository を生成します
func NewInMemoryUserRepository() repository.UserRepository {
	return &inMemoryUserRepository{
		users: make(map[domain.UserID]*domain.User),
	}
}

func (r *inMemoryUserRepository) FindByID(ctx context.Context, id domain.UserID) (*domain.User, error) {
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

func (r *inMemoryUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found with email: %s", email)
}

func (r *inMemoryUserRepository) Update(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[user.ID]; !ok {
		return fmt.Errorf("user not found: %s", user.ID)
	}
	r.users[user.ID] = user
	return nil
}

func (r *inMemoryUserRepository) UpdateRelatedUsers(ctx context.Context, userID domain.UserID, relatedUsers []domain.RelatedUser) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, ok := r.users[userID]
	if !ok {
		return fmt.Errorf("user not found: %s", userID)
	}
	user.RelatedUsers = relatedUsers
	return nil
}
