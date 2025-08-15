package domain

import (
	"time"
)

// アラーム情報
type MorningCall struct {
	ID         MorningCallID
	SenderID   string
	ReceiverID string
	Time       time.Time
	Message    string
	Status     MorningCallStatus
}