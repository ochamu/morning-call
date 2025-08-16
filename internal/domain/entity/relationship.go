package entity

import (
	"time"

	"morning-call/internal/domain"
	"morning-call/internal/domain/valueobject"
)

// Relationship represents a relationship between two users
type Relationship struct {
	ID          valueobject.RelationshipID
	RequesterID domain.UserID
	ReceiverID  domain.UserID
	Status      valueobject.RelationshipStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewRelationship creates a new Relationship instance
func NewRelationship(requesterID, receiverID domain.UserID, status valueobject.RelationshipStatus) *Relationship {
	now := time.Now()
	return &Relationship{
		ID:          valueobject.NewRelationshipID(),
		RequesterID: requesterID,
		ReceiverID:  receiverID,
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// IsPending checks if the relationship is pending
func (r *Relationship) IsPending() bool {
	return r.Status == valueobject.RelationshipStatusPending
}

// IsApproved checks if the relationship is approved
func (r *Relationship) IsApproved() bool {
	return r.Status == valueobject.RelationshipStatusApproved
}

// IsRejected checks if the relationship is rejected
func (r *Relationship) IsRejected() bool {
	return r.Status == valueobject.RelationshipStatusRejected
}

// IsBlocked checks if the relationship is blocked
func (r *Relationship) IsBlocked() bool {
	return r.Status == valueobject.RelationshipStatusBlocked
}

// CanBeApprovedBy checks if the relationship can be approved by the given user
func (r *Relationship) CanBeApprovedBy(userID domain.UserID) bool {
	return r.ReceiverID == userID && r.IsPending()
}

// CanBeRejectedBy checks if the relationship can be rejected by the given user
func (r *Relationship) CanBeRejectedBy(userID domain.UserID) bool {
	return r.ReceiverID == userID && r.IsPending()
}

// InvolvesUser checks if the relationship involves the given user
func (r *Relationship) InvolvesUser(userID domain.UserID) bool {
	return r.RequesterID == userID || r.ReceiverID == userID
}

// GetOtherUserID returns the ID of the other user in the relationship
func (r *Relationship) GetOtherUserID(userID domain.UserID) domain.UserID {
	if r.RequesterID == userID {
		return r.ReceiverID
	}
	return r.RequesterID
}

// UpdateStatus updates the status of the relationship
func (r *Relationship) UpdateStatus(newStatus valueobject.RelationshipStatus) domain.NGReason {
	if !newStatus.IsValid() {
		return domain.NGReasonInvalidRelationshipStatus
	}

	if !r.Status.CanTransitionTo(newStatus) {
		return domain.NGReasonInvalidStatusTransition
	}

	r.Status = newStatus
	r.UpdatedAt = time.Now()
	return ""
}
