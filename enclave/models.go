package enclave

import "time"

type EnvironmentVariable struct {
	Key   string
	Value string
}

type TaskStatus struct {
	CompletedAt   time.Time
	LastError     string
	LastFailedAt  time.Time
	NextProcessAt time.Time
	ResultPayload string
	Retries       int
	State         string
}

type Task struct {
	ID        string
	Source    string
	Params    []any
	Args      []string
	Env       []EnvironmentVariable
	Callback  string
	Retention string
	Retries   int
	Status    TaskStatus
}

type TaskLog struct {
	Issuer    string
	Level     string
	Message   string
	Timestamp time.Time
}

type CreateTaskRequest struct {
	Source    string
	Params    []any
	Args      []string
	Env       []EnvironmentVariable
	Callback  string
	Retention string
	Retries   int
}

type GetTaskLogsRequest struct {
	Level         string
	Issuer        string
	TimeRangeFrom time.Time
	TimeRangeTo   time.Time
}

type Artifact struct {
	Namespace   string
	Name        string
	VersionHash string
	CreatedAt   time.Time
	Pulls       int
	Tags        []string
}

type User struct {
	Name        string
	DisplayName string
	Roles       []string
}

type Role struct {
	Name  string
	Users []string
}

type ResourceGroup struct {
	Name      string
	Endpoints []string
}

type RBACPolicy struct {
	Role          string
	ResourceGroup string
	Method        string
}

type PutRoleRequest struct {
	Users []string
}

type PutResourceGroupRequest struct {
	Endpoints []string
}

type PatchArtifactRequest struct {
	Tags []string
}

type UploadArtifactResult struct {
	VersionHash string
}

type PatchCurrentUserRequest struct {
	DisplayName string
	Password    string
	Roles       []string
}

type PatchUserRequest struct {
	DisplayName string
	Password    string
	Roles       []string
}

type PutUserRequest struct {
	DisplayName string
	Password    string
	Roles       []string
}

type ListUsersRequest struct {
	Name        string
	DisplayName string
}

type ListRolesRequest struct {
	Role string
}

type ListResourceGroupsRequest struct {
	Role string
}

type ListPoliciesRequest struct {
	Role          string
	ResourceGroup string
	Method        string
}

type ListTasksRequest struct {
	State string
}
