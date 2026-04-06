package enclave

import "testing"

func TestBuildGetUsersParamsInjectsPaginationAndFilters(t *testing.T) {
	params := buildGetUsersParams(ListUsersRequest{Name: "alice", DisplayName: "Alice"}, 25, 75)

	if params.Limit == nil || *params.Limit != 25 {
		t.Fatalf("limit not set correctly: %#v", params.Limit)
	}
	if params.Offset == nil || *params.Offset != 75 {
		t.Fatalf("offset not set correctly: %#v", params.Offset)
	}
	if params.Name == nil || *params.Name != "alice" {
		t.Fatalf("name filter not set correctly: %#v", params.Name)
	}
	if params.DisplayName == nil || *params.DisplayName != "Alice" {
		t.Fatalf("display_name filter not set correctly: %#v", params.DisplayName)
	}
}

func TestBuildGetPoliciesParamsPreservesFilters(t *testing.T) {
	req := ListPoliciesRequest{Role: "admin", ResourceGroup: "artifact", Method: "GET"}
	params := buildGetPoliciesParams(req, 10, 30)

	if params.Limit == nil || *params.Limit != 10 {
		t.Fatalf("limit not set correctly")
	}
	if params.Offset == nil || *params.Offset != 30 {
		t.Fatalf("offset not set correctly")
	}
	if params.Role == nil || *params.Role != "admin" {
		t.Fatalf("role filter not preserved")
	}
	if params.ResourceGroup == nil || *params.ResourceGroup != "artifact" {
		t.Fatalf("resource-group filter not preserved")
	}
	if params.Method == nil || *params.Method != "GET" {
		t.Fatalf("method filter not preserved")
	}
}

func TestResolvePageSizeUsesDefaultAndOverrides(t *testing.T) {
	pageSize, err := resolvePageSize(50)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pageSize != 50 {
		t.Fatalf("expected default page size 50, got %d", pageSize)
	}

	pageSize, err = resolvePageSize(50, WithPageSize(20))
	if err != nil {
		t.Fatalf("unexpected override error: %v", err)
	}
	if pageSize != 20 {
		t.Fatalf("expected overridden page size 20, got %d", pageSize)
	}
}

func TestWithPageSizeRejectsInvalidValue(t *testing.T) {
	_, err := resolvePageSize(50, WithPageSize(0))
	if err == nil {
		t.Fatalf("expected error for invalid page size")
	}
}
