package errors

import (
	"fmt"
)

// ErrorType represents the type of domain error
type ErrorType string

const (
	// ErrorTypeNotFound indicates a resource was not found
	ErrorTypeNotFound ErrorType = "NOT_FOUND"

	// ErrorTypeValidation indicates a validation error
	ErrorTypeValidation ErrorType = "VALIDATION"

	// ErrorTypeAuthorization indicates an authorization error
	ErrorTypeAuthorization ErrorType = "AUTHORIZATION"

	// ErrorTypeConflict indicates a conflict error (e.g., duplicate)
	ErrorTypeConflict ErrorType = "CONFLICT"

	// ErrorTypeInternal indicates an internal server error
	ErrorTypeInternal ErrorType = "INTERNAL"

	// ErrorTypeBadRequest indicates a bad request error
	ErrorTypeBadRequest ErrorType = "BAD_REQUEST"
)

// DomainError represents a domain-specific error
type DomainError struct {
	Type    ErrorType
	Message string
	Details map[string]interface{}
}

// Error implements the error interface
func (e *DomainError) Error() string {
	return e.Message
}

// NewDomainError creates a new domain error
func NewDomainError(errType ErrorType, message string) *DomainError {
	return &DomainError{
		Type:    errType,
		Message: message,
		Details: make(map[string]interface{}),
	}
}

// WithDetails adds details to the domain error
func (e *DomainError) WithDetails(key string, value interface{}) *DomainError {
	if e.Details == nil {
		e.Details = make(map[string]interface{})
	}
	e.Details[key] = value
	return e
}

// NotFoundError creates a not found error
func NotFoundError(resource string) *DomainError {
	return NewDomainError(
		ErrorTypeNotFound,
		fmt.Sprintf("%s not found", resource),
	).WithDetails("resource", resource)
}

// ValidationError creates a validation error
func ValidationError(field, reason string) *DomainError {
	return NewDomainError(
		ErrorTypeValidation,
		fmt.Sprintf("validation failed for %s: %s", field, reason),
	).WithDetails("field", field).WithDetails("reason", reason)
}

// AuthorizationError creates an authorization error
func AuthorizationError(action string) *DomainError {
	return NewDomainError(
		ErrorTypeAuthorization,
		fmt.Sprintf("not authorized to %s", action),
	).WithDetails("action", action)
}

// ConflictError creates a conflict error
func ConflictError(resource string) *DomainError {
	return NewDomainError(
		ErrorTypeConflict,
		fmt.Sprintf("%s already exists", resource),
	).WithDetails("resource", resource)
}

// InternalError creates an internal error
func InternalError(message string) *DomainError {
	return NewDomainError(ErrorTypeInternal, message)
}

// BadRequestError creates a bad request error
func BadRequestError(message string) *DomainError {
	return NewDomainError(ErrorTypeBadRequest, message)
}
