package enclave

import (
	"fmt"

	"github.com/EnclaveRunner/sdk-go/client"
)

const defaultPageSize = 50

type config struct {
	httpClient      client.HttpRequestDoer
	requestEditors  []client.RequestEditorFn
	defaultPageSize int
}

type Option func(*config) error

type ListOption func(*listConfig) error

type listConfig struct {
	pageSize int
}

func newDefaultConfig() config {
	return config{defaultPageSize: defaultPageSize}
}

func WithHTTPClient(httpClient client.HttpRequestDoer) Option {
	return func(cfg *config) error {
		cfg.httpClient = httpClient
		return nil
	}
}

func WithRequestEditorFn(fn client.RequestEditorFn) Option {
	return func(cfg *config) error {
		cfg.requestEditors = append(cfg.requestEditors, fn)
		return nil
	}
}

func WithDefaultPageSize(pageSize int) Option {
	return func(cfg *config) error {
		if pageSize < 1 {
			return fmt.Errorf("default page size must be >= 1, got %d", pageSize)
		}
		cfg.defaultPageSize = pageSize
		return nil
	}
}

func WithPageSize(pageSize int) ListOption {
	return func(cfg *listConfig) error {
		if pageSize < 1 {
			return fmt.Errorf("page size must be >= 1, got %d", pageSize)
		}
		cfg.pageSize = pageSize
		return nil
	}
}

func resolvePageSize(defaultSize int, opts ...ListOption) (int, error) {
	cfg := listConfig{pageSize: defaultSize}
	for _, opt := range opts {
		if err := opt(&cfg); err != nil {
			return 0, err
		}
	}
	if cfg.pageSize < 1 {
		return 0, fmt.Errorf("page size must be >= 1, got %d", cfg.pageSize)
	}
	return cfg.pageSize, nil
}
