package inmemory

import (
	"context"
	"testing"
	"time"

	"morning-call/internal/domain"
	"morning-call/internal/domain/entity"
	"morning-call/internal/domain/valueobject"
)

func TestInMemoryRelationshipRepository(t *testing.T) {
	ctx := context.Background()
	repo := NewInMemoryRelationshipRepository()

	userID1 := domain.UserID("user-1")
	userID2 := domain.UserID("user-2")
	userID3 := domain.UserID("user-3")

	t.Run("Create and FindByID", func(t *testing.T) {
		rel := &entity.Relationship{
			ID:          valueobject.NewRelationshipID(),
			RequesterID: userID1,
			ReceiverID:  userID2,
			Status:      valueobject.RelationshipStatusPending,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		err := repo.Create(ctx, rel)
		if err != nil {
			t.Fatalf("Failed to create relationship: %v", err)
		}

		found, err := repo.FindByID(ctx, rel.ID)
		if err != nil {
			t.Fatalf("Failed to find relationship by ID: %v", err)
		}

		if found.ID != rel.ID {
			t.Errorf("Expected ID %s, got %s", rel.ID, found.ID)
		}
		if found.RequesterID != rel.RequesterID {
			t.Errorf("Expected RequesterID %s, got %s", rel.RequesterID, found.RequesterID)
		}
		if found.ReceiverID != rel.ReceiverID {
			t.Errorf("Expected ReceiverID %s, got %s", rel.ReceiverID, found.ReceiverID)
		}
		if found.Status != rel.Status {
			t.Errorf("Expected Status %s, got %s", rel.Status, found.Status)
		}
	})

	t.Run("FindByUsers", func(t *testing.T) {
		rel := &entity.Relationship{
			ID:          valueobject.NewRelationshipID(),
			RequesterID: userID2,
			ReceiverID:  userID3,
			Status:      valueobject.RelationshipStatusApproved,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		err := repo.Create(ctx, rel)
		if err != nil {
			t.Fatalf("Failed to create relationship: %v", err)
		}

		found, err := repo.FindByUsers(ctx, userID2, userID3)
		if err != nil {
			t.Fatalf("Failed to find relationship by users: %v", err)
		}

		if found.ID != rel.ID {
			t.Errorf("Expected ID %s, got %s", rel.ID, found.ID)
		}

		// Should not find with reversed order
		_, err = repo.FindByUsers(ctx, userID3, userID2)
		if err != ErrRelationshipNotFound {
			t.Errorf("Expected ErrRelationshipNotFound for reversed order, got %v", err)
		}
	})

	t.Run("FindByUserWithStatus", func(t *testing.T) {
		// Clear existing data
		repo = NewInMemoryRelationshipRepository()

		// Create multiple relationships
		rel1 := &entity.Relationship{
			ID:          valueobject.NewRelationshipID(),
			RequesterID: userID1,
			ReceiverID:  userID2,
			Status:      valueobject.RelationshipStatusApproved,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		rel2 := &entity.Relationship{
			ID:          valueobject.NewRelationshipID(),
			RequesterID: userID3,
			ReceiverID:  userID1,
			Status:      valueobject.RelationshipStatusApproved,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		rel3 := &entity.Relationship{
			ID:          valueobject.NewRelationshipID(),
			RequesterID: userID1,
			ReceiverID:  userID3,
			Status:      valueobject.RelationshipStatusPending,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		repo.Create(ctx, rel1)
		repo.Create(ctx, rel2)
		repo.Create(ctx, rel3)

		// Find approved relationships for userID1
		approved, err := repo.FindByUserWithStatus(ctx, userID1, valueobject.RelationshipStatusApproved)
		if err != nil {
			t.Fatalf("Failed to find relationships by user with status: %v", err)
		}

		if len(approved) != 2 {
			t.Errorf("Expected 2 approved relationships, got %d", len(approved))
		}

		// Find pending relationships for userID1
		pending, err := repo.FindByUserWithStatus(ctx, userID1, valueobject.RelationshipStatusPending)
		if err != nil {
			t.Fatalf("Failed to find relationships by user with status: %v", err)
		}

		if len(pending) != 1 {
			t.Errorf("Expected 1 pending relationship, got %d", len(pending))
		}
	})

	t.Run("Update", func(t *testing.T) {
		repo = NewInMemoryRelationshipRepository()

		rel := &entity.Relationship{
			ID:          valueobject.NewRelationshipID(),
			RequesterID: userID1,
			ReceiverID:  userID2,
			Status:      valueobject.RelationshipStatusPending,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		err := repo.Create(ctx, rel)
		if err != nil {
			t.Fatalf("Failed to create relationship: %v", err)
		}

		// Update status
		rel.Status = valueobject.RelationshipStatusApproved
		rel.UpdatedAt = time.Now()

		err = repo.Update(ctx, rel)
		if err != nil {
			t.Fatalf("Failed to update relationship: %v", err)
		}

		found, err := repo.FindByID(ctx, rel.ID)
		if err != nil {
			t.Fatalf("Failed to find updated relationship: %v", err)
		}

		if found.Status != valueobject.RelationshipStatusApproved {
			t.Errorf("Expected status to be approved, got %s", found.Status)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		repo = NewInMemoryRelationshipRepository()

		rel := &entity.Relationship{
			ID:          valueobject.NewRelationshipID(),
			RequesterID: userID1,
			ReceiverID:  userID2,
			Status:      valueobject.RelationshipStatusPending,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		err := repo.Create(ctx, rel)
		if err != nil {
			t.Fatalf("Failed to create relationship: %v", err)
		}

		err = repo.Delete(ctx, rel.ID)
		if err != nil {
			t.Fatalf("Failed to delete relationship: %v", err)
		}

		_, err = repo.FindByID(ctx, rel.ID)
		if err != ErrRelationshipNotFound {
			t.Errorf("Expected ErrRelationshipNotFound after deletion, got %v", err)
		}
	})

	t.Run("ExistsByUsers", func(t *testing.T) {
		repo = NewInMemoryRelationshipRepository()

		rel := &entity.Relationship{
			ID:          valueobject.NewRelationshipID(),
			RequesterID: userID1,
			ReceiverID:  userID2,
			Status:      valueobject.RelationshipStatusApproved,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		exists, err := repo.ExistsByUsers(ctx, userID1, userID2)
		if err != nil {
			t.Fatalf("Failed to check existence: %v", err)
		}
		if exists {
			t.Error("Expected relationship to not exist before creation")
		}

		err = repo.Create(ctx, rel)
		if err != nil {
			t.Fatalf("Failed to create relationship: %v", err)
		}

		exists, err = repo.ExistsByUsers(ctx, userID1, userID2)
		if err != nil {
			t.Fatalf("Failed to check existence: %v", err)
		}
		if !exists {
			t.Error("Expected relationship to exist after creation")
		}
	})
}
