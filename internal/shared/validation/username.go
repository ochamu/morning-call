package validation

import (
	"strings"
	"unicode/utf8"
)

// ValidateUsername validates username format and length
func ValidateUsername(username string) bool {
	username = strings.TrimSpace(username)

	// Check if empty
	if username == "" {
		return false
	}

	// Check length (min 2, max 30 characters)
	length := utf8.RuneCountInString(username)
	if length < 2 || length > 30 {
		return false
	}

	// Check for invalid characters (allow alphanumeric, Japanese, underscore, hyphen)
	for _, r := range username {
		if !isValidUsernameRune(r) {
			return false
		}
	}

	return true
}

// isValidUsernameRune checks if a rune is valid for username
func isValidUsernameRune(r rune) bool {
	// Allow alphanumeric
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
		return true
	}

	// Allow underscore and hyphen
	if r == '_' || r == '-' {
		return true
	}

	// Allow Japanese characters (Hiragana, Katakana, Kanji)
	if r >= 0x3040 && r <= 0x309F { // Hiragana
		return true
	}
	if r >= 0x30A0 && r <= 0x30FF { // Katakana
		return true
	}
	if r >= 0x4E00 && r <= 0x9FAF { // CJK Unified Ideographs
		return true
	}

	return false
}

// NormalizeUsername normalizes username (trim spaces)
func NormalizeUsername(username string) string {
	return strings.TrimSpace(username)
}
