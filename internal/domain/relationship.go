package domain

import "time"

// Relationship represents the relationship between two users
// It follows a unidirectional model: Requester -> Receiver
type Relationship struct {
	ID          RelationshipID
	RequesterID UserID // The user who initiated the relationship
	ReceiverID  UserID // The user who received the relationship request
	Status      RelationshipStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewRelationship creates a new relationship with pending status
func NewRelationship(requesterID, receiverID UserID) *Relationship {
	now := time.Now()
	return &Relationship{
		ID:          NewRelationshipID(),
		RequesterID: requesterID,
		ReceiverID:  receiverID,
		Status:      RelationshipStatusPending,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// CanApprove checks if the given user can approve this relationship
func (r *Relationship) CanApprove(userID UserID) bool {
	return r.ReceiverID == userID && r.Status == RelationshipStatusPending
}

// CanReject checks if the given user can reject this relationship
func (r *Relationship) CanReject(userID UserID) bool {
	return r.ReceiverID == userID && r.Status == RelationshipStatusPending
}

// Approve approves the relationship
func (r *Relationship) Approve() error {
	if r.Status != RelationshipStatusPending {
		return ErrInvalidRelationshipStatus
	}
	r.Status = RelationshipStatusApproved
	r.UpdatedAt = time.Now()
	return nil
}

// Reject rejects the relationship
func (r *Relationship) Reject() error {
	if r.Status != RelationshipStatusPending {
		return ErrInvalidRelationshipStatus
	}
	r.Status = RelationshipStatusRejected
	r.UpdatedAt = time.Now()
	return nil
}

// Block blocks the relationship
func (r *Relationship) Block() {
	r.Status = RelationshipStatusBlocked
	r.UpdatedAt = time.Now()
}

// IsFriend checks if the relationship represents a friendship
func (r *Relationship) IsFriend() bool {
	return r.Status == RelationshipStatusApproved
}

// IsBlocked checks if the relationship is blocked
func (r *Relationship) IsBlocked() bool {
	return r.Status == RelationshipStatusBlocked
}

// IsPending checks if the relationship is pending
func (r *Relationship) IsPending() bool {
	return r.Status == RelationshipStatusPending
}

// InvolvesUser checks if the relationship involves the given user
func (r *Relationship) InvolvesUser(userID UserID) bool {
	return r.RequesterID == userID || r.ReceiverID == userID
}

// GetOtherUserID returns the ID of the other user in the relationship
// Returns empty UserID if the provided user is not involved in the relationship
func (r *Relationship) GetOtherUserID(userID UserID) UserID {
	if r.RequesterID == userID {
		return r.ReceiverID
	}
	if r.ReceiverID == userID {
		return r.RequesterID
	}
	// User not involved in this relationship
	return ""
}
