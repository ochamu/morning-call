package usecase

import (
	"context"
	"morning-call/internal/domain"
)

// UserUsecase defines the interface for user-related use cases
type UserUsecase interface {
	Register(ctx context.Context, username, email string) (*domain.User, error)
}

// MorningCallUsecase defines the interface for morning call-related use cases
type MorningCallUsecase interface {
	SaveFriendMorningCall(ctx context.Context, userID, friendID domain.UserID, morningCall *domain.MorningCall) error
	GetFriendMorningCall(ctx context.Context, userID, friendID domain.UserID, morningCallID domain.MorningCallID) (*domain.MorningCall, error)
}