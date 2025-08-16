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

func NewMorningCallUsecase(morningCallRepo repository.MorningCallRepository, userRepo repository.UserRepository) MorningCallUsecase {
	return &morningCallUsecase{
		morningCallRepo: morningCallRepo,
		userRepo:        userRepo,
	}
}

func (rcv *morningCallUsecase) SaveFriendMorningCall(ctx context.Context, userID, friendID domain.UserID, morningCall *domain.MorningCall) error {
	friend, err := rcv.userRepo.FindByID(ctx, friendID)
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

func (rcv *morningCallUsecase) GetFriendMorningCall(ctx context.Context, userID, friendID domain.UserID, morningCallID domain.MorningCallID) (*domain.MorningCall, error) {
	// モーニングコールを取得
	morningCall, err := rcv.morningCallRepo.FindByID(ctx, morningCallID)
	if err != nil {
		return nil, err
	}

	// アクセス権限チェック（送信者または受信者のみアクセス可能）
	if morningCall.SenderID != userID && morningCall.ReceiverID != userID {
		return nil, fmt.Errorf("アクセス権限がありません")
	}

	// フレンド関係チェック
	if morningCall.SenderID == userID && morningCall.ReceiverID != friendID {
		return nil, fmt.Errorf("指定されたフレンドのモーニングコールではありません")
	}
	if morningCall.ReceiverID == userID && morningCall.SenderID != friendID {
		return nil, fmt.Errorf("指定されたフレンドのモーニングコールではありません")
	}

	return morningCall, nil
}

func (rcv *morningCallUsecase) ListMorningCalls(ctx context.Context, userID domain.UserID) ([]*domain.MorningCall, error) {
	// 送信者として送ったモーニングコール
	sentCalls, err := rcv.morningCallRepo.ListBySenderID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 受信者として受け取ったモーニングコール
	receivedCalls, err := rcv.morningCallRepo.ListByReceiverID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 両方のリストを結合
	allCalls := append(sentCalls, receivedCalls...)

	return allCalls, nil
}

func (rcv *morningCallUsecase) UpdateMorningCall(ctx context.Context, userID domain.UserID, morningCall *domain.MorningCall) error {
	// 既存のモーニングコールを取得
	existingCall, err := rcv.morningCallRepo.FindByID(ctx, morningCall.ID)
	if err != nil {
		return err
	}

	// 更新権限チェック（送信者のみ更新可能）
	if ng := existingCall.CanUpdate(userID); ng.IsNG() {
		return fmt.Errorf("モーニングコールを更新できません: %s", ng.String())
	}

	// 時刻の妥当性チェック
	if ng := morningCall.ValidateScheduledTime(); ng.IsNG() {
		return fmt.Errorf("無効な時刻設定: %s", ng.String())
	}

	// 更新実行
	return rcv.morningCallRepo.Update(ctx, morningCall)
}

func (rcv *morningCallUsecase) DeleteMorningCall(ctx context.Context, userID domain.UserID, morningCallID domain.MorningCallID) error {
	// モーニングコールを取得
	morningCall, err := rcv.morningCallRepo.FindByID(ctx, morningCallID)
	if err != nil {
		return err
	}

	// 削除権限チェック（送信者のみ削除可能）
	if ng := morningCall.CanDelete(userID); ng.IsNG() {
		return fmt.Errorf("モーニングコールを削除できません: %s", ng.String())
	}

	// 削除実行
	return rcv.morningCallRepo.Delete(ctx, morningCallID)
}
