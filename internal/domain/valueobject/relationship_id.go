package valueobject

import "github.com/google/uuid"

// RelationshipID represents a unique identifier for a Relationship
type RelationshipID string

// NewRelationshipID creates a new RelationshipID
func NewRelationshipID() RelationshipID {
	return RelationshipID(uuid.New().String())
}

// String returns the string representation of RelationshipID
func (id RelationshipID) String() string {
	return string(id)
}

// IsEmpty checks if the RelationshipID is empty
func (id RelationshipID) IsEmpty() bool {
	return id == ""
}
