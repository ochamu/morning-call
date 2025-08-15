package repository

import (
	"context"

	"morning-call/internal/domain"
)

type UserRepository interface {
	FindByID(ctx context.Context, id domain.UserID) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	UpdateRelatedUsers(ctx context.Context, userID domain.UserID, relatedUsers []domain.RelatedUser) error
}
