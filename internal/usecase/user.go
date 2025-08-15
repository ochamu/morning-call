package usecase

import (
	"context"

	"morning-call/internal/domain"
	"morning-call/internal/repository"

	"github.com/google/uuid"
)

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) Register(ctx context.Context, username, email string) (*domain.User, error) {
	// TODO: 万が一被った時の対応策を考える
	// - 被らないようにulidで引く
	// - 被った時に再生成する
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	// TODO: constructor作成して中身を知らなくても良いようにする
	user := &domain.User{
		ID:       domain.UserID(newUUID.String()),
		Username: username,
		Email:    email,
	}

	err = u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Login(ctx context.Context, userID domain.UserID) error {
	// todo: implement
	return nil
}

func (u *userUsecase) ListFriends(ctx context.Context, userID domain.UserID) ([]domain.RelatedUser, error) {
	// TODO: implement
	return nil, nil
}

func (u *userUsecase) ReactFriendApply(ctx context.Context, userID, applyingUserID domain.UserID) (resultRelatedUserStatus domain.RelatedUserStatus, err error) {
	// TODO: implement
	return "", nil
}

func (u *userUsecase) BlockFriend(ctx context.Context, userID, blockUserID domain.UserID) error {
	// TODO: implement
	return nil
}
