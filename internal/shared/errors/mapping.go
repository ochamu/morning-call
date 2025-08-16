package errors

import (
	"net/http"
)

// HTTPStatusFromError maps a domain error to an HTTP status code
func HTTPStatusFromError(err error) int {
	if domainErr, ok := err.(*DomainError); ok {
		switch domainErr.Type {
		case ErrorTypeNotFound:
			return http.StatusNotFound
		case ErrorTypeValidation:
			return http.StatusBadRequest
		case ErrorTypeAuthorization:
			return http.StatusForbidden
		case ErrorTypeConflict:
			return http.StatusConflict
		case ErrorTypeInternal:
			return http.StatusInternalServerError
		case ErrorTypeBadRequest:
			return http.StatusBadRequest
		default:
			return http.StatusInternalServerError
		}
	}

	// Default to internal server error for unknown errors
	return http.StatusInternalServerError
}

// ErrorResponse represents the API error response structure
type ErrorResponse struct {
	Error   string                 `json:"error"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// ToErrorResponse converts a domain error to an API error response
func ToErrorResponse(err error) *ErrorResponse {
	if domainErr, ok := err.(*DomainError); ok {
		return &ErrorResponse{
			Error:   string(domainErr.Type),
			Message: domainErr.Message,
			Details: domainErr.Details,
		}
	}

	// For non-domain errors, create a generic error response
	return &ErrorResponse{
		Error:   string(ErrorTypeInternal),
		Message: err.Error(),
		Details: nil,
	}
}

// IsNotFoundError checks if an error is a not found error
func IsNotFoundError(err error) bool {
	if domainErr, ok := err.(*DomainError); ok {
		return domainErr.Type == ErrorTypeNotFound
	}
	return false
}

// IsValidationError checks if an error is a validation error
func IsValidationError(err error) bool {
	if domainErr, ok := err.(*DomainError); ok {
		return domainErr.Type == ErrorTypeValidation
	}
	return false
}

// IsAuthorizationError checks if an error is an authorization error
func IsAuthorizationError(err error) bool {
	if domainErr, ok := err.(*DomainError); ok {
		return domainErr.Type == ErrorTypeAuthorization
	}
	return false
}

// IsConflictError checks if an error is a conflict error
func IsConflictError(err error) bool {
	if domainErr, ok := err.(*DomainError); ok {
		return domainErr.Type == ErrorTypeConflict
	}
	return false
}
