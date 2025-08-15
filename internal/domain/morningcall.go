package domain

import (
	"time"
)

// アラーム情報
type MorningCall struct {
	ID         MorningCallID
	SenderID   UserID
	ReceiverID UserID
	Time       time.Time
	Message    string
	Status     MorningCallStatus
}
