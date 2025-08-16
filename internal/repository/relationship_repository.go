package repository

import (
	"context"
	"morning-call/internal/domain"
)

// RelationshipRepository defines the interface for relationship data access
type RelationshipRepository interface {
	// Create creates a new relationship
	Create(ctx context.Context, relationship *domain.Relationship) error

	// FindByID finds a relationship by its ID
	FindByID(ctx context.Context, id domain.RelationshipID) (*domain.Relationship, error)

	// FindByUsers finds a relationship between two users (unidirectional)
	// Returns the relationship where user1 is the requester and user2 is the receiver
	FindByUsers(ctx context.Context, requesterID, receiverID domain.UserID) (*domain.Relationship, error)

	// FindAllByUsers finds all relationships between two users (bidirectional)
	// Returns relationships in both directions (user1->user2 and user2->user1)
	FindAllByUsers(ctx context.Context, user1ID, user2ID domain.UserID) ([]*domain.Relationship, error)

	// FindByUser finds all relationships involving a specific user
	// Returns relationships where the user is either requester or receiver
	FindByUser(ctx context.Context, userID domain.UserID) ([]*domain.Relationship, error)

	// FindByUserWithStatus finds all relationships for a user with a specific status
	// Returns relationships where the user is involved and has the specified status
	FindByUserWithStatus(ctx context.Context, userID domain.UserID, status domain.RelationshipStatus) ([]*domain.Relationship, error)

	// FindByRequester finds all relationships initiated by a specific user
	FindByRequester(ctx context.Context, requesterID domain.UserID) ([]*domain.Relationship, error)

	// FindByReceiver finds all relationships received by a specific user
	FindByReceiver(ctx context.Context, receiverID domain.UserID) ([]*domain.Relationship, error)

	// Update updates an existing relationship
	Update(ctx context.Context, relationship *domain.Relationship) error

	// Delete deletes a relationship
	Delete(ctx context.Context, relationship *domain.Relationship) error

	// DeleteByID deletes a relationship by its ID
	DeleteByID(ctx context.Context, id domain.RelationshipID) error

	// ExistsByUsers checks if a relationship exists between two users (unidirectional)
	ExistsByUsers(ctx context.Context, requesterID, receiverID domain.UserID) (bool, error)
}
