package validation

import (
	"time"
)

// ValidateScheduledTime validates if the scheduled time is valid for a morning call
func ValidateScheduledTime(scheduledTime time.Time) bool {
	now := time.Now()

	// Cannot schedule in the past
	if scheduledTime.Before(now) {
		return false
	}

	// Cannot schedule more than 1 year in advance
	oneYearFromNow := now.AddDate(1, 0, 0)
	if scheduledTime.After(oneYearFromNow) {
		return false
	}

	// Check if time is within reasonable morning call hours (4:00 AM - 12:00 PM)
	hour := scheduledTime.Hour()
	if hour < 4 || hour >= 12 {
		return false
	}

	return true
}

// ValidateDuration validates if a duration is within acceptable range
func ValidateDuration(duration time.Duration) bool {
	// Minimum 1 minute, maximum 1 hour
	if duration < time.Minute || duration > time.Hour {
		return false
	}
	return true
}

// IsBusinessHours checks if the given time is within business hours
func IsBusinessHours(t time.Time) bool {
	hour := t.Hour()
	// Business hours: 9:00 AM - 6:00 PM
	return hour >= 9 && hour < 18
}

// IsMorningCallTime checks if the given time is appropriate for morning calls
func IsMorningCallTime(t time.Time) bool {
	hour := t.Hour()
	// Morning call hours: 4:00 AM - 12:00 PM
	return hour >= 4 && hour < 12
}
