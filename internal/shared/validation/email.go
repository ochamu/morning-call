package validation

import (
	"regexp"
	"strings"
)

// Email validation regex pattern
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// ValidateEmail validates email format
func ValidateEmail(email string) bool {
	email = strings.TrimSpace(email)
	if email == "" {
		return false
	}
	return emailRegex.MatchString(email)
}

// NormalizeEmail normalizes email address (lowercase and trim)
func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
