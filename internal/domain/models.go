package domain

import (
	"time"
)

// システムの利用者
type User struct {
	ID           UserID
	Username     string
	Email        string
	MorningCalls []MorningCall
	RelatedUsers []RelatedUser
}

// フレンド申請のステータス管理
type RelatedUser struct {
	ID       UserID
	Username string
	Email    string
	Status   RelatedUserStatus
}

// アラーム情報
type MorningCall struct {
	ID         MorningCallID
	SenderID   string
	ReceiverID string
	Time       time.Time
	Message    string
	Status     MorningCallStatus
}

type RelatedUserStatus string

const (
	RelatedUserStatusApproved RelatedUserStatus = "approved"
	RelatedUserStatusPending  RelatedUserStatus = "pending"
	RelatedUserStatusRejected RelatedUserStatus = "rejected"
	RelatedUserStatusBlocked  RelatedUserStatus = "blocked"
)

type MorningCallStatus string

const (
	MorningCallStatusScheduled MorningCallStatus = "scheduled"
	MorningCallStatusDeleted   MorningCallStatus = "deleted"
	MorningCallStatusCompleted MorningCallStatus = "completed"
	MorningCallStatusFailed    MorningCallStatus = "failed"
)

func (rcv *User) CanAcceptMorningCall(friendID UserID) NGReason {
	for _, ru := range rcv.RelatedUsers {
		if ru.ID == friendID {
			return ""
		}
	}
	return "フレンドではありません。"
}
