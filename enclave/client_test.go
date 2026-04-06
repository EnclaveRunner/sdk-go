package enclave

import (
	"context"
	"errors"
	"net/http"
	"testing"
)

type captureAuthTransport struct {
	req   *http.Request
	calls int
}

func (t *captureAuthTransport) Do(req *http.Request) (*http.Response, error) {
	t.calls++
	t.req = req
	return nil, errors.New("stop")
}

func TestNewClientRequiresUsernameAndPassword(t *testing.T) {
	if _, err := NewClient("http://localhost:8080", "", "secret"); err == nil {
		t.Fatalf("expected error for empty username")
	}
	if _, err := NewClient("http://localhost:8080", "user", ""); err == nil {
		t.Fatalf("expected error for empty password")
	}
}

func TestNewClientAppliesBasicAuthEditor(t *testing.T) {
	transport := &captureAuthTransport{}
	c, err := NewClient("http://localhost:8080", "demo", "secret", WithHTTPClient(transport))
	if err != nil {
		t.Fatalf("unexpected NewClient error: %v", err)
	}

	_, err = c.raw.GetV1UserWithResponse(context.Background(), nil)
	if err == nil {
		t.Fatalf("expected injected transport error")
	}
	if transport.calls != 1 {
		t.Fatalf("expected one transport call, got %d", transport.calls)
	}
	if transport.req == nil {
		t.Fatalf("expected captured request")
	}
	username, password, ok := transport.req.BasicAuth()
	if !ok {
		t.Fatalf("expected basic auth to be set")
	}
	if username != "demo" || password != "secret" {
		t.Fatalf("unexpected basic auth credentials: %q / %q", username, password)
	}
}
