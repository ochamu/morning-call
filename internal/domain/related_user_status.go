package domain

type RelatedUserStatus string

const (
	RelatedUserStatusApproved RelatedUserStatus = "approved"
	RelatedUserStatusPending  RelatedUserStatus = "pending"
	RelatedUserStatusRejected RelatedUserStatus = "rejected"
	RelatedUserStatusBlocked  RelatedUserStatus = "blocked"
)

// CanTransitionTo checks if the status can transition to the target status
func (rcv RelatedUserStatus) CanTransitionTo(target RelatedUserStatus) NGReason {
	// 同じステータスへの遷移は無効
	if rcv == target {
		return NGReasonInvalidStatus
	}

	switch rcv {
	case RelatedUserStatusPending:
		// Pending → Approved, Rejected, Blocked
		switch target {
		case RelatedUserStatusApproved, RelatedUserStatusRejected, RelatedUserStatusBlocked:
			return ""
		default:
			return NGReasonInvalidStatus
		}

	case RelatedUserStatusApproved:
		// Approved → Blocked のみ可能
		if target == RelatedUserStatusBlocked {
			return ""
		}
		return NGReasonInvalidStatus

	case RelatedUserStatusRejected:
		// Rejected → Pending (再申請), Blocked
		switch target {
		case RelatedUserStatusPending, RelatedUserStatusBlocked:
			return ""
		default:
			return NGReasonInvalidStatus
		}

	case RelatedUserStatusBlocked:
		// Blocked → Pending (ブロック解除後の再申請) のみ可能
		if target == RelatedUserStatusPending {
			return ""
		}
		return NGReasonInvalidStatus

	default:
		return NGReasonInvalidStatus
	}
}

// IsActive checks if the status represents an active relationship
func (rcv RelatedUserStatus) IsActive() bool {
	return rcv == RelatedUserStatusApproved
}

// IsPending checks if the status represents a pending relationship
func (rcv RelatedUserStatus) IsPending() bool {
	return rcv == RelatedUserStatusPending
}

// IsBlocked checks if the status represents a blocked relationship
func (rcv RelatedUserStatus) IsBlocked() bool {
	return rcv == RelatedUserStatusBlocked
}
