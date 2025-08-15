package repository

import (
	"context"
	"morning-call/internal/domain"
)

type MorningCallRepository interface {
	FindByID(ctx context.Context, id domain.MorningCallID) (*domain.MorningCall, error)
	Save(ctx context.Context, morningCall *domain.MorningCall) error
	Update(ctx context.Context, morningCall *domain.MorningCall) error
	Delete(ctx context.Context, id domain.MorningCallID) error
	ListBySenderID(ctx context.Context, senderID domain.UserID) ([]*domain.MorningCall, error)
	ListByReceiverID(ctx context.Context, receiverID domain.UserID) ([]*domain.MorningCall, error)
}
