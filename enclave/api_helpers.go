package enclave

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/EnclaveRunner/sdk-go/client"
)

func newAPIErrorFromResponse(statusCode int, body []byte, candidates ...string) error {
	message := ""
	for _, candidate := range candidates {
		if strings.TrimSpace(candidate) != "" {
			message = candidate
			break
		}
	}
	if message == "" {
		message = strings.TrimSpace(string(bytes.TrimSpace(body)))
	}
	if message == "" {
		message = fmt.Sprintf("unexpected response status %d", statusCode)
	}
	return newAPIError(statusCode, message)
}

func fieldErrorMessage(in *client.FieldError) string {
	if in == nil || in.Errors == nil || len(*in.Errors) == 0 {
		return ""
	}
	first := (*in.Errors)[0]
	if first.Field != "" && first.Error != "" {
		return first.Field + ": " + first.Error
	}
	if first.Error != "" {
		return first.Error
	}
	return ""
}
