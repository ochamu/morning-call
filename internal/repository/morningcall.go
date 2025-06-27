package repository

import (
	"context"
	"morning-call/internal/domain"
)

type MorningCallRepository interface {
	FindByID(ctx context.Context, id string) (*domain.MorningCall, error)
	Save(ctx context.Context, morningCall *domain.MorningCall) error
}
