package usecase

import (
	"context"
	"morning-call/internal/domain"
)

// UserUsecase defines the interface for user-related use cases
type UserUsecase interface {
	Register(ctx context.Context, username, email string) (*domain.User, error)
	Login(ctx context.Context, email string) (*domain.User, error)
	ListFriends(ctx context.Context, userID domain.UserID) ([]domain.RelatedUser, error)
	ApplyFriend(ctx context.Context, userID, targetUserID domain.UserID) error
	ReactFriendApply(ctx context.Context, userID, applyingUserID domain.UserID, approve bool) (domain.RelatedUserStatus, error)
	BlockFriend(ctx context.Context, userID, blockUserID domain.UserID) error
}

// MorningCallUsecase defines the interface for morning call-related use cases
type MorningCallUsecase interface {
	SaveFriendMorningCall(ctx context.Context, userID, friendID domain.UserID, morningCall *domain.MorningCall) error
	GetFriendMorningCall(ctx context.Context, userID, friendID domain.UserID, morningCallID domain.MorningCallID) (*domain.MorningCall, error)
	ListMorningCalls(ctx context.Context, userID domain.UserID) ([]*domain.MorningCall, error)
	UpdateMorningCall(ctx context.Context, userID domain.UserID, morningCall *domain.MorningCall) error
	DeleteMorningCall(ctx context.Context, userID domain.UserID, morningCallID domain.MorningCallID) error
}
