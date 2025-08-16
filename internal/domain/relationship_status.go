package domain

// RelationshipStatus represents the status of a relationship between two users
type RelationshipStatus string

const (
	// RelationshipStatusPending indicates a pending friend request
	RelationshipStatusPending RelationshipStatus = "pending"

	// RelationshipStatusApproved indicates an approved friendship
	RelationshipStatusApproved RelationshipStatus = "approved"

	// RelationshipStatusRejected indicates a rejected friend request
	RelationshipStatusRejected RelationshipStatus = "rejected"

	// RelationshipStatusBlocked indicates a blocked relationship
	RelationshipStatusBlocked RelationshipStatus = "blocked"
)

// IsValid checks if the relationship status is valid
func (s RelationshipStatus) IsValid() bool {
	switch s {
	case RelationshipStatusPending,
		RelationshipStatusApproved,
		RelationshipStatusRejected,
		RelationshipStatusBlocked:
		return true
	default:
		return false
	}
}

// String returns the string representation of the relationship status
func (s RelationshipStatus) String() string {
	return string(s)
}
