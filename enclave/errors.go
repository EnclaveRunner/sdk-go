package enclave

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidRequest = errors.New("invalid request")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrNotFound       = errors.New("not found")
	ErrConflict       = errors.New("conflict")
	ErrTooLarge       = errors.New("payload too large")
	ErrServer         = errors.New("server error")
)

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	if e.Message == "" {
		return fmt.Sprintf("api request failed with status %d", e.StatusCode)
	}
	return fmt.Sprintf("api request failed with status %d: %s", e.StatusCode, e.Message)
}

func classifyStatus(statusCode int) error {
	switch statusCode {
	case 400:
		return ErrInvalidRequest
	case 401:
		return ErrUnauthorized
	case 403:
		return ErrForbidden
	case 404:
		return ErrNotFound
	case 409:
		return ErrConflict
	case 413:
		return ErrTooLarge
	default:
		if statusCode >= 500 {
			return ErrServer
		}
		return nil
	}
}

func newAPIError(statusCode int, message string) error {
	base := classifyStatus(statusCode)
	if base == nil {
		return &APIError{StatusCode: statusCode, Message: message}
	}
	return fmt.Errorf("%w: %w", base, &APIError{StatusCode: statusCode, Message: message})
}
