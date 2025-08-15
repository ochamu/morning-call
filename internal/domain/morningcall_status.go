package domain

type MorningCallStatus string

const (
	MorningCallStatusScheduled MorningCallStatus = "scheduled"
	MorningCallStatusDeleted   MorningCallStatus = "deleted"
	MorningCallStatusCompleted MorningCallStatus = "completed"
	MorningCallStatusFailed    MorningCallStatus = "failed"
)