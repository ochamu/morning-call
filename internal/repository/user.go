package repository

import (
	"context"

	"morning-call/internal/domain"
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
}
