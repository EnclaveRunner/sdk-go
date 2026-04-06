package enclave

import (
	"context"
	"fmt"
	"net/http"

	"github.com/EnclaveRunner/sdk-go/client"
)

type Client struct {
	raw             *client.ClientWithResponses
	defaultPageSize int
}

func NewClient(server, username, password string, opts ...Option) (*Client, error) {
	if username == "" {
		return nil, fmt.Errorf("username is required")
	}
	if password == "" {
		return nil, fmt.Errorf("password is required")
	}

	cfg := newDefaultConfig()
	for _, opt := range opts {
		if err := opt(&cfg); err != nil {
			return nil, err
		}
	}

	basicAuthEditor := func(_ context.Context, req *http.Request) error {
		req.SetBasicAuth(username, password)
		return nil
	}
	cfg.requestEditors = append([]client.RequestEditorFn{basicAuthEditor}, cfg.requestEditors...)

	rawOpts := make([]client.ClientOption, 0, len(cfg.requestEditors)+1)
	if cfg.httpClient != nil {
		rawOpts = append(rawOpts, client.WithHTTPClient(cfg.httpClient))
	}
	for _, editor := range cfg.requestEditors {
		rawOpts = append(rawOpts, client.WithRequestEditorFn(editor))
	}

	rawClient, err := client.NewClientWithResponses(server, rawOpts...)
	if err != nil {
		return nil, err
	}

	return &Client{raw: rawClient, defaultPageSize: cfg.defaultPageSize}, nil
}
