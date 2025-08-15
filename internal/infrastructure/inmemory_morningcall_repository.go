package infrastructure

import (
	"context"
	"fmt"
	"morning-call/internal/domain"
	"sync"
)

type inMemoryMorningCallRepository struct {
	mu           sync.RWMutex
	morningCalls map[string]*domain.MorningCall
}

func NewInMemoryMorningCallRepository() *inMemoryMorningCallRepository {
	return &inMemoryMorningCallRepository{
		morningCalls: make(map[string]*domain.MorningCall),
	}
}

func (r *inMemoryMorningCallRepository) FindByID(ctx context.Context, id string) (*domain.MorningCall, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	morningCall, ok := r.morningCalls[id]
	if !ok {
		return nil, fmt.Errorf("morning call not found: %s", id)
	}
	return morningCall, nil
}

func (r *inMemoryMorningCallRepository) Save(ctx context.Context, morningCall *domain.MorningCall) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if morningCall.ID == "" {
		return fmt.Errorf("morning call ID is required")
	}

	r.morningCalls[string(morningCall.ID)] = morningCall
	return nil
}

func (r *inMemoryMorningCallRepository) Update(ctx context.Context, morningCall *domain.MorningCall) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.morningCalls[string(morningCall.ID)]; !ok {
		return fmt.Errorf("morning call not found: %s", morningCall.ID)
	}

	r.morningCalls[string(morningCall.ID)] = morningCall
	return nil
}

func (r *inMemoryMorningCallRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.morningCalls[id]; !ok {
		return fmt.Errorf("morning call not found: %s", id)
	}

	delete(r.morningCalls, id)
	return nil
}

func (r *inMemoryMorningCallRepository) ListBySender(ctx context.Context, senderID string) ([]*domain.MorningCall, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*domain.MorningCall
	for _, mc := range r.morningCalls {
		if mc.SenderID == senderID {
			result = append(result, mc)
		}
	}
	return result, nil
}

func (r *inMemoryMorningCallRepository) ListByReceiver(ctx context.Context, receiverID string) ([]*domain.MorningCall, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*domain.MorningCall
	for _, mc := range r.morningCalls {
		if mc.ReceiverID == receiverID {
			result = append(result, mc)
		}
	}
	return result, nil
}