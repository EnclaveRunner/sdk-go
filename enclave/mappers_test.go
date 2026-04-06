package enclave

import (
	"reflect"
	"testing"
	"time"

	"github.com/EnclaveRunner/sdk-go/client"
)

func TestMapTaskPointerReduction(t *testing.T) {
	timeVal := time.Date(2026, 4, 6, 12, 0, 0, 0, time.UTC)
	callback := "https://callback"
	retention := "24h"
	retries := 3
	lastErr := "failure"
	payload := "result"

	in := client.Task{
		Id:        "task-1",
		Source:    "ns:name/function@tag",
		Callback:  &callback,
		Retention: &retention,
		Retries:   &retries,
		Status: client.TaskStatus{
			CompletedAt:   &timeVal,
			LastError:     &lastErr,
			ResultPayload: &payload,
			Retries:       5,
			State:         "DONE",
		},
	}

	mapped := mapTask(in)
	if mapped.Callback != callback {
		t.Fatalf("callback mismatch: got %q", mapped.Callback)
	}
	if mapped.Retention != retention {
		t.Fatalf("retention mismatch: got %q", mapped.Retention)
	}
	if mapped.Retries != retries {
		t.Fatalf("retries mismatch: got %d", mapped.Retries)
	}
	if mapped.Status.CompletedAt != timeVal {
		t.Fatalf("completed_at mismatch: got %v", mapped.Status.CompletedAt)
	}
	if mapped.Status.LastError != lastErr {
		t.Fatalf("last_error mismatch: got %q", mapped.Status.LastError)
	}
	if mapped.Status.ResultPayload != payload {
		t.Fatalf("result payload mismatch: got %q", mapped.Status.ResultPayload)
	}
}

func TestMapTaskNilPointersBecomeZeroValues(t *testing.T) {
	mapped := mapTask(client.Task{Id: "id", Source: "src"})

	if mapped.Callback != "" {
		t.Fatalf("expected empty callback, got %q", mapped.Callback)
	}
	if mapped.Retention != "" {
		t.Fatalf("expected empty retention, got %q", mapped.Retention)
	}
	if mapped.Retries != 0 {
		t.Fatalf("expected retries=0, got %d", mapped.Retries)
	}
	if !mapped.Status.CompletedAt.IsZero() {
		t.Fatalf("expected zero completed_at, got %v", mapped.Status.CompletedAt)
	}
}

func TestMapUserRolesPointerReduction(t *testing.T) {
	roles := []string{"admin", "ops"}
	mapped := mapUser(client.UserResponse{Name: "alice", DisplayName: "Alice", Roles: &roles})

	if !reflect.DeepEqual(mapped.Roles, roles) {
		t.Fatalf("roles mismatch: got %v", mapped.Roles)
	}

	roles[0] = "changed"
	if mapped.Roles[0] != "admin" {
		t.Fatalf("expected mapped roles to be copied")
	}
}
