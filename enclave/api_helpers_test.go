package enclave

import (
	"errors"
	"testing"

	"github.com/EnclaveRunner/sdk-go/client"
)

func TestFieldErrorMessage(t *testing.T) {
	errDetail := []client.ErrField{{Field: "role", Error: "unknown role"}}
	msg := fieldErrorMessage(&client.FieldError{Errors: &errDetail})
	if msg != "role: unknown role" {
		t.Fatalf("unexpected field error message: %q", msg)
	}
}

func TestNewAPIErrorFromResponsePrefersCandidate(t *testing.T) {
	err := newAPIErrorFromResponse(404, []byte("{\"error\":\"not found\"}"), "missing")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
	if err.Error() == "" {
		t.Fatalf("expected non-empty error message")
	}
}

func TestToClientGetTaskLogsParams(t *testing.T) {
	req := GetTaskLogsRequest{Level: "error", Issuer: "runner"}
	params := toClientGetTaskLogsParams(req)
	if params.Level == nil || *params.Level != "error" {
		t.Fatalf("expected level filter")
	}
	if params.Issuer == nil || *params.Issuer != "runner" {
		t.Fatalf("expected issuer filter")
	}
}
