package usecase

import (
	"context"
	"fmt"

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
	// メールアドレスの重複チェック
	existingUser, _ := u.userRepo.FindByEmail(ctx, email)
	if existingUser != nil {
		return nil, fmt.Errorf("email already registered")
	}

	// ユーザー作成
	user, err := u.createUser(username, email)
	if err != nil {
		return nil, err
	}

	err = u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// createUser はドメインユーザーオブジェクトを生成するヘルパー関数
func (u *userUsecase) createUser(username, email string) (*domain.User, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate user ID: %w", err)
	}

	return &domain.User{
		ID:           domain.UserID(newUUID.String()),
		Username:     username,
		Email:        email,
		MorningCalls: []domain.MorningCall{},
		RelatedUsers: []domain.RelatedUser{},
	}, nil
}

func (u *userUsecase) Login(ctx context.Context, email string) (*domain.User, error) {
	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) ListFriends(ctx context.Context, userID domain.UserID) ([]domain.RelatedUser, error) {
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// フレンド状態のRelatedUserのみを返す
	var friends []domain.RelatedUser
	for _, relatedUser := range user.RelatedUsers {
		if relatedUser.Status == domain.RelatedUserStatusApproved {
			friends = append(friends, relatedUser)
		}
	}

	return friends, nil
}

func (u *userUsecase) ApplyFriend(ctx context.Context, userID, targetUserID domain.UserID) error {
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	targetUser, err := u.userRepo.FindByID(ctx, targetUserID)
	if err != nil {
		return err
	}

	// 申請可能かチェック
	if ng := user.CanAddFriend(targetUserID); ng.IsNG() {
		return fmt.Errorf("friend application failed: %s", ng.String())
	}

	// 申請者側にpending状態で追加
	user.RelatedUsers = append(user.RelatedUsers, domain.RelatedUser{
		ID:     targetUserID,
		Status: domain.RelatedUserStatusPending,
	})

	// 被申請者側にpending状態で追加
	targetUser.RelatedUsers = append(targetUser.RelatedUsers, domain.RelatedUser{
		ID:     userID,
		Status: domain.RelatedUserStatusPending,
	})

	// 両方のユーザーを更新
	if err := u.userRepo.Update(ctx, user); err != nil {
		return err
	}

	if err := u.userRepo.Update(ctx, targetUser); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) ReactFriendApply(ctx context.Context, userID, applyingUserID domain.UserID, approve bool) (domain.RelatedUserStatus, error) {
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return "", err
	}

	applyingUser, err := u.userRepo.FindByID(ctx, applyingUserID)
	if err != nil {
		return "", err
	}

	// 承認可能かチェック
	if ng := user.CanApproveFriend(applyingUserID); ng.IsNG() {
		return "", fmt.Errorf("friend approval failed: %s", ng.String())
	}

	var newStatus domain.RelatedUserStatus
	if approve {
		newStatus = domain.RelatedUserStatusApproved
	} else {
		newStatus = domain.RelatedUserStatusRejected
	}

	// 両方のユーザーのRelatedUsersを更新
	for i, relatedUser := range user.RelatedUsers {
		if relatedUser.ID == applyingUserID {
			if approve {
				user.RelatedUsers[i].Status = domain.RelatedUserStatusApproved
			} else {
				// 拒否の場合はステータスを変更
				user.RelatedUsers[i].Status = domain.RelatedUserStatusRejected
			}
			break
		}
	}

	for i, relatedUser := range applyingUser.RelatedUsers {
		if relatedUser.ID == userID {
			if approve {
				applyingUser.RelatedUsers[i].Status = domain.RelatedUserStatusApproved
			} else {
				// 拒否の場合はステータスを変更
				applyingUser.RelatedUsers[i].Status = domain.RelatedUserStatusRejected
			}
			break
		}
	}

	// 両方のユーザーを更新
	if err := u.userRepo.Update(ctx, user); err != nil {
		return "", err
	}

	if err := u.userRepo.Update(ctx, applyingUser); err != nil {
		return "", err
	}

	return newStatus, nil
}

func (u *userUsecase) BlockFriend(ctx context.Context, userID, blockUserID domain.UserID) error {
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	blockUser, err := u.userRepo.FindByID(ctx, blockUserID)
	if err != nil {
		return err
	}

	// ブロック可能かチェック
	if ng := user.CanBlockUser(blockUserID); ng.IsNG() {
		return fmt.Errorf("block user failed: %s", ng.String())
	}

	// ユーザーのRelatedUsersを更新（ブロック状態に変更または追加）
	found := false
	for i, relatedUser := range user.RelatedUsers {
		if relatedUser.ID == blockUserID {
			user.RelatedUsers[i].Status = domain.RelatedUserStatusBlocked
			found = true
			break
		}
	}

	if !found {
		user.RelatedUsers = append(user.RelatedUsers, domain.RelatedUser{
			ID:     blockUserID,
			Status: domain.RelatedUserStatusBlocked,
		})
	}

	// ブロックされた側のRelatedUsersから削除
	for i, relatedUser := range blockUser.RelatedUsers {
		if relatedUser.ID == userID {
			blockUser.RelatedUsers = append(blockUser.RelatedUsers[:i], blockUser.RelatedUsers[i+1:]...)
			break
		}
	}

	// 両方のユーザーを更新
	if err := u.userRepo.Update(ctx, user); err != nil {
		return err
	}

	if err := u.userRepo.Update(ctx, blockUser); err != nil {
		return err
	}

	return nil
}
