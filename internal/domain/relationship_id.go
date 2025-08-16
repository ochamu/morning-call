package domain

import "github.com/google/uuid"

// RelationshipID represents a unique identifier for a relationship
type RelationshipID string

// NewRelationshipID generates a new unique relationship ID
func NewRelationshipID() RelationshipID {
	return RelationshipID(uuid.New().String())
}

// String returns the string representation of the relationship ID
func (id RelationshipID) String() string {
	return string(id)
}

// IsValid checks if the relationship ID is valid
func (id RelationshipID) IsValid() bool {
	if id == "" {
		return false
	}
	_, err := uuid.Parse(string(id))
	return err == nil
}
