package enclave

import (
	"time"

	"github.com/EnclaveRunner/sdk-go/client"
)

func mapTaskLogs(in []client.TaskLog) []TaskLog {
	out := make([]TaskLog, 0, len(in))
	for _, item := range in {
		out = append(out, TaskLog{
			Issuer:    item.Issuer,
			Level:     item.Level,
			Message:   item.Message,
			Timestamp: item.Timestamp,
		})
	}
	return out
}

func toClientPatchMe(in PatchCurrentUserRequest) client.PatchMe {
	out := client.PatchMe{}
	if in.DisplayName != "" {
		out.DisplayName = stringPtr(in.DisplayName)
	}
	if in.Password != "" {
		out.Password = stringPtr(in.Password)
	}
	if len(in.Roles) > 0 {
		roles := append([]string(nil), in.Roles...)
		out.Roles = &roles
	}
	return out
}

func toClientPatchUser(in PatchUserRequest) client.PatchUser {
	out := client.PatchUser{}
	if in.DisplayName != "" {
		out.DisplayName = stringPtr(in.DisplayName)
	}
	if in.Password != "" {
		out.Password = stringPtr(in.Password)
	}
	if len(in.Roles) > 0 {
		roles := append([]string(nil), in.Roles...)
		out.Roles = &roles
	}
	return out
}

func toClientPutUser(in PutUserRequest) client.PutUserRequest {
	out := client.PutUserRequest{
		DisplayName: in.DisplayName,
		Password:    in.Password,
	}
	if len(in.Roles) > 0 {
		roles := append([]string(nil), in.Roles...)
		out.Roles = &roles
	}
	return out
}

func toClientCreateTask(in CreateTaskRequest) client.CreateTaskRequest {
	out := client.CreateTaskRequest{Source: in.Source}
	if len(in.Params) > 0 {
		params := append([]interface{}(nil), in.Params...)
		out.Params = &params
	}
	if len(in.Args) > 0 {
		args := append([]string(nil), in.Args...)
		out.Args = &args
	}
	if len(in.Env) > 0 {
		env := make([]client.EnvironmentVariable, 0, len(in.Env))
		for _, entry := range in.Env {
			env = append(env, client.EnvironmentVariable{Key: entry.Key, Value: entry.Value})
		}
		out.Env = &env
	}
	if in.Callback != "" {
		out.Callback = stringPtr(in.Callback)
	}
	if in.Retention != "" {
		out.Retention = stringPtr(in.Retention)
	}
	if in.Retries != 0 {
		out.Retries = intPtr(in.Retries)
	}
	return out
}

func toClientPutRole(in PutRoleRequest) client.PutRoleRequest {
	return client.PutRoleRequest{Users: append([]string(nil), in.Users...)}
}

func toClientPutResourceGroup(in PutResourceGroupRequest) client.PutResourceGroupRequest {
	return client.PutResourceGroupRequest{Endpoints: append([]string(nil), in.Endpoints...)}
}

func toClientPatchArtifact(in PatchArtifactRequest) client.PatchArtifact {
	out := client.PatchArtifact{}
	if len(in.Tags) > 0 {
		tags := append([]string(nil), in.Tags...)
		out.Tags = &tags
	}
	return out
}

func toClientGetTaskLogsParams(in GetTaskLogsRequest) *client.GetV1TaskIdLogsParams {
	out := &client.GetV1TaskIdLogsParams{}
	if in.Level != "" {
		out.Level = stringPtr(in.Level)
	}
	if in.Issuer != "" {
		out.Issuer = stringPtr(in.Issuer)
	}
	if !in.TimeRangeFrom.IsZero() {
		from := in.TimeRangeFrom
		out.TimeRangeFrom = &from
	}
	if !in.TimeRangeTo.IsZero() {
		to := in.TimeRangeTo
		out.TimeRangeTo = &to
	}
	return out
}

func mapUser(in client.UserResponse) User {
	out := User{
		Name:        in.Name,
		DisplayName: in.DisplayName,
	}
	if in.Roles != nil {
		out.Roles = append([]string(nil), (*in.Roles)...)
	}
	return out
}

func mapUsers(in []client.UserResponse) []User {
	out := make([]User, 0, len(in))
	for _, item := range in {
		out = append(out, mapUser(item))
	}
	return out
}

func mapRole(in client.RoleResource) Role {
	return Role{Name: in.Name, Users: append([]string(nil), in.Users...)}
}

func mapRoles(in []client.RoleResource) []Role {
	out := make([]Role, 0, len(in))
	for _, item := range in {
		out = append(out, mapRole(item))
	}
	return out
}

func mapResourceGroup(in client.ResourceGroupResource) ResourceGroup {
	return ResourceGroup{Name: in.Name, Endpoints: append([]string(nil), in.Endpoints...)}
}

func mapResourceGroups(in []client.ResourceGroupResource) []ResourceGroup {
	out := make([]ResourceGroup, 0, len(in))
	for _, item := range in {
		out = append(out, mapResourceGroup(item))
	}
	return out
}

func mapPolicy(in client.RBACPolicy) RBACPolicy {
	return RBACPolicy{Role: in.Role, ResourceGroup: in.ResourceGroup, Method: string(in.Method)}
}

func mapPolicies(in []client.RBACPolicy) []RBACPolicy {
	out := make([]RBACPolicy, 0, len(in))
	for _, item := range in {
		out = append(out, mapPolicy(item))
	}
	return out
}

func mapArtifact(in client.Artifact) Artifact {
	return Artifact{
		Namespace:   in.Namespace,
		Name:        in.Name,
		VersionHash: in.VersionHash,
		CreatedAt:   in.CreatedAt,
		Pulls:       in.Pulls,
		Tags:        append([]string(nil), in.Tags...),
	}
}

func mapArtifacts(in []client.Artifact) []Artifact {
	out := make([]Artifact, 0, len(in))
	for _, item := range in {
		out = append(out, mapArtifact(item))
	}
	return out
}

func mapTaskStatus(in client.TaskStatus) TaskStatus {
	return TaskStatus{
		CompletedAt:   valueOrZeroTime(in.CompletedAt),
		LastError:     valueOrZeroString(in.LastError),
		LastFailedAt:  valueOrZeroTime(in.LastFailedAt),
		NextProcessAt: valueOrZeroTime(in.NextProcessAt),
		ResultPayload: valueOrZeroString(in.ResultPayload),
		Retries:       in.Retries,
		State:         in.State,
	}
}

func mapTask(in client.Task) Task {
	out := Task{
		ID:        in.Id,
		Source:    in.Source,
		Callback:  valueOrZeroString(in.Callback),
		Retention: valueOrZeroString(in.Retention),
		Retries:   valueOrZeroInt(in.Retries),
		Status:    mapTaskStatus(in.Status),
	}
	if in.Args != nil {
		out.Args = append([]string(nil), (*in.Args)...)
	}
	if in.Params != nil {
		out.Params = append([]any(nil), (*in.Params)...)
	}
	if in.Env != nil {
		out.Env = make([]EnvironmentVariable, 0, len(*in.Env))
		for _, v := range *in.Env {
			out.Env = append(out.Env, EnvironmentVariable{Key: v.Key, Value: v.Value})
		}
	}
	return out
}

func mapTasks(in []client.Task) []Task {
	out := make([]Task, 0, len(in))
	for _, item := range in {
		out = append(out, mapTask(item))
	}
	return out
}

func valueOrZeroString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func valueOrZeroInt(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func valueOrZeroTime(v *time.Time) time.Time {
	if v == nil {
		return time.Time{}
	}
	return *v
}
