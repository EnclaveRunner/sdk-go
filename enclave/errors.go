package enclave

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/EnclaveRunner/sdk-go/client"
)

// Sentinel errors for API response status codes.
var (
	ErrBadRequest = errors.New(
		"bad request",
	)
	ErrUnauthenticated = errors.New(
		"unauthenticated",
	)
	ErrForbidden = errors.New(
		"forbidden",
	)
	ErrNotFound = errors.New(
		"not found",
	)
	ErrConflict = errors.New(
		"conflict",
	)
	ErrTooLarge = errors.New(
		"request too large",
	)
	ErrInternal = errors.New(
		"internal server error",
	)
)

// APIError is a structured error returned by the
// Enclave API. Use errors.Is to check the underlying
// sentinel (e.g. errors.Is(err, enclave.ErrNotFound)).
type APIError struct {
	StatusCode int
	Message    string
	Sentinel   error
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf(
			"enclave: %s: %s",
			e.Sentinel,
			e.Message,
		)
	}

	return fmt.Sprintf("enclave: %s", e.Sentinel)
}

// Unwrap returns the sentinel error.
func (e *APIError) Unwrap() error {
	return e.Sentinel
}

// mapHTTPError returns an *APIError for non-2xx status
// codes, or nil for success codes.
func mapHTTPError(
	statusCode int,
	message string,
) error {
	sentinel := statusToSentinel(statusCode)
	if sentinel == nil {
		return nil
	}

	return &APIError{
		StatusCode: statusCode,
		Message:    message,
		Sentinel:   sentinel,
	}
}

// statusToSentinel maps an HTTP status code to the
// corresponding sentinel error. Returns nil for 2xx.
func statusToSentinel(code int) error {
	switch code {
	case http.StatusOK,
		http.StatusCreated,
		http.StatusNoContent:
		return nil
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusUnauthorized:
		return ErrUnauthenticated
	case http.StatusForbidden:
		return ErrForbidden
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusConflict:
		return ErrConflict
	case http.StatusRequestEntityTooLarge:
		return ErrTooLarge
	case http.StatusInternalServerError:
		return ErrInternal
	default:
		return ErrInternal
	}
}

// extractErrMessage returns the first non-empty error
// message from the known message, or attempts to parse
// the body as a generic error JSON object.
func extractErrMessage(
	body []byte,
	knownMsg string,
) string {
	if knownMsg != "" {
		return knownMsg
	}

	if len(body) == 0 {
		return ""
	}

	var generic struct {
		Error string `json:"error"`
	}

	if json.Unmarshal(body, &generic) == nil {
		return generic.Error
	}

	return ""
}

// firstErrMessage returns the error message from the
// first non-nil ErrGeneric value.
func firstErrMessage(
	errs ...*client.ErrGeneric,
) string {
	for _, e := range errs {
		if e != nil {
			return e.Error
		}
	}

	return ""
}
