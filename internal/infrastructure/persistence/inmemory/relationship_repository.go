package inmemory

import (
	"context"
	"errors"
	"sync"

	"morning-call/internal/domain"
	"morning-call/internal/domain/entity"
	"morning-call/internal/domain/valueobject"
	"morning-call/internal/repository"
)

var (
	ErrRelationshipNotFound = errors.New("relationship not found")
	ErrRelationshipExists   = errors.New("relationship already exists")
)

// inMemoryRelationshipRepository is an in-memory implementation of RelationshipRepository
type inMemoryRelationshipRepository struct {
	mu            sync.RWMutex
	relationships map[valueobject.RelationshipID]*entity.Relationship
	// Indexes for faster lookups
	userIndex map[domain.UserID][]valueobject.RelationshipID // userID -> []relationshipID
	pairIndex map[string]valueobject.RelationshipID          // "requesterID:receiverID" -> relationshipID
}

// NewInMemoryRelationshipRepository creates a new in-memory relationship repository
func NewInMemoryRelationshipRepository() repository.RelationshipRepository {
	return &inMemoryRelationshipRepository{
		relationships: make(map[valueobject.RelationshipID]*entity.Relationship),
		userIndex:     make(map[domain.UserID][]valueobject.RelationshipID),
		pairIndex:     make(map[string]valueobject.RelationshipID),
	}
}

// Create creates a new relationship
func (r *inMemoryRelationshipRepository) Create(ctx context.Context, relationship *entity.Relationship) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if relationship already exists
	pairKey := r.makePairKey(relationship.RequesterID, relationship.ReceiverID)
	if _, exists := r.pairIndex[pairKey]; exists {
		return ErrRelationshipExists
	}

	// Store relationship
	r.relationships[relationship.ID] = relationship

	// Update indexes
	r.pairIndex[pairKey] = relationship.ID
	r.userIndex[relationship.RequesterID] = append(r.userIndex[relationship.RequesterID], relationship.ID)
	r.userIndex[relationship.ReceiverID] = append(r.userIndex[relationship.ReceiverID], relationship.ID)

	return nil
}

// FindByID finds a relationship by its ID
func (r *inMemoryRelationshipRepository) FindByID(ctx context.Context, id valueobject.RelationshipID) (*entity.Relationship, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	relationship, exists := r.relationships[id]
	if !exists {
		return nil, ErrRelationshipNotFound
	}

	// Return a copy to prevent external modifications
	return r.copyRelationship(relationship), nil
}

// FindByUsers finds a relationship between two users
func (r *inMemoryRelationshipRepository) FindByUsers(ctx context.Context, requesterID, receiverID domain.UserID) (*entity.Relationship, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	pairKey := r.makePairKey(requesterID, receiverID)
	relationshipID, exists := r.pairIndex[pairKey]
	if !exists {
		return nil, ErrRelationshipNotFound
	}

	relationship := r.relationships[relationshipID]
	return r.copyRelationship(relationship), nil
}

// FindByUserWithStatus finds all relationships for a user with a specific status
func (r *inMemoryRelationshipRepository) FindByUserWithStatus(ctx context.Context, userID domain.UserID, status valueobject.RelationshipStatus) ([]*entity.Relationship, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*entity.Relationship
	relationshipIDs := r.userIndex[userID]

	for _, relID := range relationshipIDs {
		rel := r.relationships[relID]
		if rel.Status == status {
			result = append(result, r.copyRelationship(rel))
		}
	}

	return result, nil
}

// FindByUser finds all relationships involving a user
func (r *inMemoryRelationshipRepository) FindByUser(ctx context.Context, userID domain.UserID) ([]*entity.Relationship, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*entity.Relationship
	relationshipIDs := r.userIndex[userID]

	for _, relID := range relationshipIDs {
		rel := r.relationships[relID]
		result = append(result, r.copyRelationship(rel))
	}

	return result, nil
}

// Update updates an existing relationship
func (r *inMemoryRelationshipRepository) Update(ctx context.Context, relationship *entity.Relationship) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.relationships[relationship.ID]; !exists {
		return ErrRelationshipNotFound
	}

	r.relationships[relationship.ID] = r.copyRelationship(relationship)
	return nil
}

// Delete deletes a relationship by its ID
func (r *inMemoryRelationshipRepository) Delete(ctx context.Context, id valueobject.RelationshipID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	relationship, exists := r.relationships[id]
	if !exists {
		return ErrRelationshipNotFound
	}

	// Remove from main storage
	delete(r.relationships, id)

	// Remove from indexes
	pairKey := r.makePairKey(relationship.RequesterID, relationship.ReceiverID)
	delete(r.pairIndex, pairKey)

	// Remove from user indexes
	r.removeFromUserIndex(relationship.RequesterID, id)
	r.removeFromUserIndex(relationship.ReceiverID, id)

	return nil
}

// DeleteByUsers deletes a relationship between two users
func (r *inMemoryRelationshipRepository) DeleteByUsers(ctx context.Context, requesterID, receiverID domain.UserID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	pairKey := r.makePairKey(requesterID, receiverID)
	relationshipID, exists := r.pairIndex[pairKey]
	if !exists {
		return ErrRelationshipNotFound
	}

	relationship := r.relationships[relationshipID]

	// Remove from main storage
	delete(r.relationships, relationshipID)

	// Remove from indexes
	delete(r.pairIndex, pairKey)

	// Remove from user indexes
	r.removeFromUserIndex(relationship.RequesterID, relationshipID)
	r.removeFromUserIndex(relationship.ReceiverID, relationshipID)

	return nil
}

// ExistsByUsers checks if a relationship exists between two users
func (r *inMemoryRelationshipRepository) ExistsByUsers(ctx context.Context, requesterID, receiverID domain.UserID) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	pairKey := r.makePairKey(requesterID, receiverID)
	_, exists := r.pairIndex[pairKey]
	return exists, nil
}

// Helper methods

func (r *inMemoryRelationshipRepository) makePairKey(requesterID, receiverID domain.UserID) string {
	return string(requesterID) + ":" + string(receiverID)
}

func (r *inMemoryRelationshipRepository) copyRelationship(rel *entity.Relationship) *entity.Relationship {
	return &entity.Relationship{
		ID:          rel.ID,
		RequesterID: rel.RequesterID,
		ReceiverID:  rel.ReceiverID,
		Status:      rel.Status,
		CreatedAt:   rel.CreatedAt,
		UpdatedAt:   rel.UpdatedAt,
	}
}

func (r *inMemoryRelationshipRepository) removeFromUserIndex(userID domain.UserID, relationshipID valueobject.RelationshipID) {
	ids := r.userIndex[userID]
	for i, id := range ids {
		if id == relationshipID {
			r.userIndex[userID] = append(ids[:i], ids[i+1:]...)
			break
		}
	}
	if len(r.userIndex[userID]) == 0 {
		delete(r.userIndex, userID)
	}
}
