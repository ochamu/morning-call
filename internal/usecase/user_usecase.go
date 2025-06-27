package usecase

import (
	"context"

	"morning-call/internal/domain"
	"morning-call/internal/repository"

	"github.com/google/uuid"
)

type UserUsecase interface {
	Register(ctx context.Context, username, email string) (*domain.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) Register(ctx context.Context, username, email string) (*domain.User, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		ID:       newUUID.String(),
		Username: username,
		Email:    email,
	}

	err = u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
