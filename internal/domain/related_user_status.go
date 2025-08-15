package domain

type RelatedUserStatus string

const (
	RelatedUserStatusApproved RelatedUserStatus = "approved"
	RelatedUserStatusPending  RelatedUserStatus = "pending"
	RelatedUserStatusRejected RelatedUserStatus = "rejected"
	RelatedUserStatusBlocked  RelatedUserStatus = "blocked"
)