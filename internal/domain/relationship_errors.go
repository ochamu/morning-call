package domain

import "errors"

// Relationship-related errors
var (
	// ErrInvalidRelationshipStatus indicates an invalid status transition
	ErrInvalidRelationshipStatus = errors.New("invalid relationship status transition")

	// ErrRelationshipNotFound indicates the relationship was not found
	ErrRelationshipNotFound = errors.New("relationship not found")

	// ErrAlreadyFriends indicates users are already friends
	ErrAlreadyFriends = errors.New("users are already friends")

	// ErrAlreadyRequested indicates a friend request already exists
	ErrAlreadyRequested = errors.New("friend request already sent")

	// ErrCannotRequestSelf indicates a user cannot send a friend request to themselves
	ErrCannotRequestSelf = errors.New("cannot send friend request to yourself")

	// ErrUserBlocked indicates the user is blocked
	ErrUserBlocked = errors.New("user is blocked")

	// ErrNotReceiver indicates the user is not the receiver of the relationship
	ErrNotReceiver = errors.New("user is not the receiver of this relationship")

	// ErrNotFriends indicates users are not friends
	ErrNotFriends = errors.New("users are not friends")

	// ErrCannotBlockSelf indicates a user cannot block themselves
	ErrCannotBlockSelf = errors.New("cannot block yourself")
)
