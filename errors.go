package oneclaw

import (
	"errors"
	"fmt"
)

// Common SDK errors. Use errors.Is to check.
var (
	ErrUnauthorized   = errors.New("1claw: unauthorized")
	ErrValidation     = errors.New("1claw: validation error")
	ErrNotFound       = errors.New("1claw: not found")
	ErrRateLimited    = errors.New("1claw: rate limited")
	ErrPaymentRequired = errors.New("1claw: payment required")
	ErrServerError    = errors.New("1claw: server error")
)

// AuthError represents an authentication or authorization failure.
type AuthError struct {
	StatusCode int
	Message    string
	Detail     string
	Body       []byte
}

func (e *AuthError) Error() string {
	if e.Detail != "" {
		return fmt.Sprintf("%s: %s (%s)", ErrUnauthorized, e.Message, e.Detail)
	}
	return fmt.Sprintf("%s: %s", ErrUnauthorized, e.Message)
}

func (e *AuthError) Is(target error) bool {
	return target == ErrUnauthorized
}

// APIError wraps API response errors with status and body.
type APIError struct {
	StatusCode int
	Message    string
	Detail     string
	Body       []byte
}

func (e *APIError) Error() string {
	msg := e.Message
	if e.Detail != "" {
		msg += ": " + e.Detail
	}
	return fmt.Sprintf("1claw API error (status %d): %s", e.StatusCode, msg)
}
