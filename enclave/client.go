package enclave

import (
	"context"
	"fmt"
	"net/http"

	"github.com/EnclaveRunner/sdk-go/client"
)

// Client is the Enclave SDK client. Create one with
// New.
type Client struct {
	api *client.ClientWithResponses
}

// New creates a new Enclave SDK client.
// Both username and password are required for basic
// authentication. The serverURL is the base URL of the
// Enclave API (e.g. "https://enclave.example.com").
func New(
	serverURL string,
	username string,
	password string,
) (*Client, error) {
	basicAuth := func(
		_ context.Context,
		req *http.Request,
	) error {
		req.SetBasicAuth(username, password)

		return nil
	}

	api, err := client.NewClientWithResponses(
		serverURL,
		client.WithRequestEditorFn(basicAuth),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"creating enclave client: %w",
			err,
		)
	}

	return &Client{api: api}, nil
}
