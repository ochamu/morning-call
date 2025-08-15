package domain

import (
	"time"
)

// IsSender checks if the specified user is the sender of this morning call
func (rcv *MorningCall) IsSender(userID string) bool {
	return rcv.SenderID == userID
}

// IsReceiver checks if the specified user is the receiver of this morning call
func (rcv *MorningCall) IsReceiver(userID string) bool {
	return rcv.ReceiverID == userID
}

// CanUpdate checks if the morning call can be updated
func (rcv *MorningCall) CanUpdate(userID string) NGReason {
	// 送信者のみが更新可能
	if !rcv.IsSender(userID) {
		return NGReasonNotSender
	}

	// ステータスチェック
	switch rcv.Status {
	case MorningCallStatusCompleted:
		return NGReasonAlreadyCompleted
	case MorningCallStatusDeleted:
		return NGReasonAlreadyDeleted
	case MorningCallStatusScheduled:
		return ""
	default:
		return NGReasonInvalidStatus
	}
}

// CanDelete checks if the morning call can be deleted
func (rcv *MorningCall) CanDelete(userID string) NGReason {
	// 送信者または受信者が削除可能
	if !rcv.IsSender(userID) && !rcv.IsReceiver(userID) {
		return NGReasonNoPermission
	}

	// ステータスチェック
	switch rcv.Status {
	case MorningCallStatusCompleted:
		return NGReasonAlreadyCompleted
	case MorningCallStatusDeleted:
		return NGReasonAlreadyDeleted
	case MorningCallStatusScheduled:
		return ""
	default:
		return NGReasonInvalidStatus
	}
}

// ValidateScheduledTime validates if the scheduled time is valid
func (rcv *MorningCall) ValidateScheduledTime() NGReason {
	now := time.Now()

	// 過去の時刻チェック
	if rcv.Time.Before(now) {
		return NGReasonPastTime
	}

	// 設定可能期間のチェック（例：30日以内）
	maxFuture := now.AddDate(0, 0, 30)
	if rcv.Time.After(maxFuture) {
		return NGReasonTooFarInFuture
	}

	return ""
}

// CanComplete checks if the morning call can be marked as complete
func (rcv *MorningCall) CanComplete() NGReason {
	switch rcv.Status {
	case MorningCallStatusScheduled:
		// 予定時刻を過ぎているかチェック
		if time.Now().Before(rcv.Time) {
			return NGReasonInvalidTime
		}
		return ""
	case MorningCallStatusCompleted:
		return NGReasonAlreadyCompleted
	case MorningCallStatusDeleted:
		return NGReasonAlreadyDeleted
	default:
		return NGReasonInvalidStatus
	}
}
