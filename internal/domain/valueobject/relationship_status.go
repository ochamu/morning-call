package valueobject

// RelationshipStatus represents the status of a relationship between users
type RelationshipStatus string

const (
	// RelationshipStatusPending indicates a pending friend request
	RelationshipStatusPending RelationshipStatus = "pending"

	// RelationshipStatusApproved indicates an approved friend relationship
	RelationshipStatusApproved RelationshipStatus = "approved"

	// RelationshipStatusRejected indicates a rejected friend request
	RelationshipStatusRejected RelationshipStatus = "rejected"

	// RelationshipStatusBlocked indicates a blocked user relationship
	RelationshipStatusBlocked RelationshipStatus = "blocked"
)

// IsValid checks if the relationship status is valid
func (s RelationshipStatus) IsValid() bool {
	switch s {
	case RelationshipStatusPending, RelationshipStatusApproved,
		RelationshipStatusRejected, RelationshipStatusBlocked:
		return true
	default:
		return false
	}
}

// CanTransitionTo checks if the status can transition to the target status
func (s RelationshipStatus) CanTransitionTo(target RelationshipStatus) bool {
	switch s {
	case RelationshipStatusPending:
		return target == RelationshipStatusApproved ||
			target == RelationshipStatusRejected
	case RelationshipStatusRejected:
		return target == RelationshipStatusPending // Can re-request after rejection
	case RelationshipStatusApproved:
		return false // Once approved, must delete to change
	case RelationshipStatusBlocked:
		return false // Once blocked, must delete to change
	default:
		return false
	}
}

// String returns the string representation of the status
func (s RelationshipStatus) String() string {
	return string(s)
}
