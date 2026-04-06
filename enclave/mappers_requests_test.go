package enclave

import (
	"reflect"
	"testing"
	"time"
)

func TestToClientCreateTaskMapsOptionalFields(t *testing.T) {
	req := CreateTaskRequest{
		Source:    "ns:name/fn@tag",
		Params:    []any{"arg", 10},
		Args:      []string{"--verbose"},
		Env:       []EnvironmentVariable{{Key: "A", Value: "B"}},
		Callback:  "https://callback",
		Retention: "24h",
		Retries:   2,
	}

	mapped := toClientCreateTask(req)
	if mapped.Source != req.Source {
		t.Fatalf("source mismatch: got %q", mapped.Source)
	}
	if mapped.Params == nil || len(*mapped.Params) != 2 {
		t.Fatalf("expected mapped params")
	}
	if mapped.Args == nil || len(*mapped.Args) != 1 {
		t.Fatalf("expected mapped args")
	}
	if mapped.Env == nil || len(*mapped.Env) != 1 {
		t.Fatalf("expected mapped env")
	}
	if mapped.Callback == nil || *mapped.Callback != req.Callback {
		t.Fatalf("expected callback mapping")
	}
	if mapped.Retention == nil || *mapped.Retention != req.Retention {
		t.Fatalf("expected retention mapping")
	}
	if mapped.Retries == nil || *mapped.Retries != req.Retries {
		t.Fatalf("expected retries mapping")
	}
}

func TestToClientPatchHelpersMapOnlyProvidedFields(t *testing.T) {
	me := toClientPatchMe(PatchCurrentUserRequest{DisplayName: "Alice"})
	if me.DisplayName == nil || *me.DisplayName != "Alice" {
		t.Fatalf("expected display name mapping")
	}
	if me.Password != nil {
		t.Fatalf("did not expect password to be set")
	}
	if me.Roles != nil {
		t.Fatalf("did not expect roles to be set")
	}

	user := toClientPatchUser(PatchUserRequest{Roles: []string{"admin"}})
	if user.Roles == nil || !reflect.DeepEqual(*user.Roles, []string{"admin"}) {
		t.Fatalf("expected roles mapping")
	}
}

func TestToClientGetTaskLogsParamsMapsTimes(t *testing.T) {
	from := time.Date(2026, 4, 6, 10, 0, 0, 0, time.UTC)
	to := from.Add(2 * time.Hour)
	mapped := toClientGetTaskLogsParams(GetTaskLogsRequest{
		Level:         "info",
		Issuer:        "runner",
		TimeRangeFrom: from,
		TimeRangeTo:   to,
	})

	if mapped.Level == nil || *mapped.Level != "info" {
		t.Fatalf("expected level mapping")
	}
	if mapped.Issuer == nil || *mapped.Issuer != "runner" {
		t.Fatalf("expected issuer mapping")
	}
	if mapped.TimeRangeFrom == nil || !mapped.TimeRangeFrom.Equal(from) {
		t.Fatalf("expected from mapping")
	}
	if mapped.TimeRangeTo == nil || !mapped.TimeRangeTo.Equal(to) {
		t.Fatalf("expected to mapping")
	}
}
