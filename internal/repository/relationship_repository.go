package repository

import (
	"context"

	"morning-call/internal/domain"
	"morning-call/internal/domain/entity"
	"morning-call/internal/domain/valueobject"
)

// RelationshipRepository defines the interface for relationship persistence
type RelationshipRepository interface {
	// Create creates a new relationship
	Create(ctx context.Context, relationship *entity.Relationship) error

	// FindByID finds a relationship by its ID
	FindByID(ctx context.Context, id valueobject.RelationshipID) (*entity.Relationship, error)

	// FindByUsers finds a relationship between two users (unidirectional)
	FindByUsers(ctx context.Context, requesterID, receiverID domain.UserID) (*entity.Relationship, error)

	// FindByUserWithStatus finds all relationships for a user with a specific status
	FindByUserWithStatus(ctx context.Context, userID domain.UserID, status valueobject.RelationshipStatus) ([]*entity.Relationship, error)

	// FindByUser finds all relationships involving a user
	FindByUser(ctx context.Context, userID domain.UserID) ([]*entity.Relationship, error)

	// Update updates an existing relationship
	Update(ctx context.Context, relationship *entity.Relationship) error

	// Delete deletes a relationship by its ID
	Delete(ctx context.Context, id valueobject.RelationshipID) error

	// DeleteByUsers deletes a relationship between two users
	DeleteByUsers(ctx context.Context, requesterID, receiverID domain.UserID) error

	// ExistsByUsers checks if a relationship exists between two users
	ExistsByUsers(ctx context.Context, requesterID, receiverID domain.UserID) (bool, error)
}
