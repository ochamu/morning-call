package usecase

import (
	"context"
	"fmt"
	"morning-call/internal/domain"
	"morning-call/internal/repository"
)

type morningCallUsecase struct {
	morningCallRepo repository.MorningCallRepository
	userRepo        repository.UserRepository
}

func NewMorningCallRepository(morningCallRepo repository.MorningCallRepository, userRepo repository.UserRepository) *morningCallUsecase {
	return &morningCallUsecase{
		morningCallRepo: morningCallRepo,
		userRepo:        userRepo,
	}
}

func (rcv *morningCallUsecase) SaveFriendMorningCall(ctx context.Context, userID, friendID domain.UserID, morningCall *domain.MorningCall) error {
	friend, err := rcv.userRepo.FindByID(ctx, string(friendID))
	if err != nil {
		return err
	}

	if ng := friend.CanAcceptMorningCall(userID); ng.IsNG() {
		return fmt.Errorf("次の理由によりモーニングコールを設定できません。 %s", ng.String())
	}

	if err := rcv.morningCallRepo.Save(ctx, morningCall); err != nil {
		return err
	}

	return nil
}

func (rcv *morningCallUsecase) GetFriendMorningCall(ctx context.Context, userID, firendID domain.UserID, morningCallID domain.MorningCallID) (*domain.MorningCall, error) {
	// todo
	return nil, nil
}

func (rcv *morningCallUsecase) ListMorningCalls(ctx context.Context, userID domain.UserID) ([]*domain.MorningCall, error) {
	// todo
	return nil, nil
}
