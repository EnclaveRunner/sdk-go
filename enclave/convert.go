package enclave

import "github.com/EnclaveRunner/sdk-go/client"

// --- Generated -> SDK conversions ---

func userFromGen(u client.UserResponse) User {
	user := User{
		Name:        u.Name,
		DisplayName: u.DisplayName,
		Roles:       make([]string, 0),
	}

	if u.Roles != nil {
		user.Roles = *u.Roles
	}

	return user
}

func usersFromGen(us []client.UserResponse) []User {
	result := make([]User, len(us))
	for i, u := range us {
		result[i] = userFromGen(u)
	}

	return result
}

func artifactFromGen(a *client.Artifact) Artifact {
	return Artifact{
		Namespace:   a.Namespace,
		Name:        a.Name,
		VersionHash: a.VersionHash,
		CreatedAt:   a.CreatedAt,
		Pulls:       a.Pulls,
		Tags:        a.Tags,
	}
}

func artifactsFromGen(
	as []client.Artifact,
) []Artifact {
	result := make([]Artifact, len(as))
	for i := range as {
		result[i] = artifactFromGen(&as[i])
	}

	return result
}

func taskStatusFromGen(
	s client.TaskStatus,
) TaskStatus {
	ts := TaskStatus{
		State:   s.State,
		Retries: s.Retries,
	}

	if s.CompletedAt != nil {
		ts.CompletedAt = *s.CompletedAt
	}

	if s.LastFailedAt != nil {
		ts.LastFailedAt = *s.LastFailedAt
	}

	if s.NextProcessAt != nil {
		ts.NextProcessAt = *s.NextProcessAt
	}

	if s.LastError != nil {
		ts.LastError = *s.LastError
	}

	if s.ResultPayload != nil {
		ts.ResultPayload = *s.ResultPayload
	}

	return ts
}

func taskFromGen(t *client.Task) Task {
	task := Task{
		ID:     t.Id,
		Source: t.Source,
		Status: taskStatusFromGen(t.Status),
	}

	if t.Args != nil {
		task.Args = *t.Args
	}

	if t.Callback != nil {
		task.Callback = *t.Callback
	}

	if t.Env != nil {
		task.Env = envsFromGen(*t.Env)
	}

	if t.Params != nil {
		task.Params = *t.Params
	}

	if t.Retention != nil {
		task.Retention = *t.Retention
	}

	if t.Retries != nil {
		task.Retries = *t.Retries
	}

	return task
}

func tasksFromGen(ts []client.Task) []Task {
	result := make([]Task, len(ts))
	for i := range ts {
		result[i] = taskFromGen(&ts[i])
	}

	return result
}

func envFromGen(
	e client.EnvironmentVariable,
) EnvironmentVariable {
	return EnvironmentVariable{
		Key:   e.Key,
		Value: e.Value,
	}
}

func envsFromGen(
	es []client.EnvironmentVariable,
) []EnvironmentVariable {
	result := make([]EnvironmentVariable, len(es))
	for i, e := range es {
		result[i] = envFromGen(e)
	}

	return result
}

func roleFromGen(r client.RoleResource) Role {
	return Role{
		Name:  r.Name,
		Users: r.Users,
	}
}

func rolesFromGen(
	rs []client.RoleResource,
) []Role {
	result := make([]Role, len(rs))
	for i, r := range rs {
		result[i] = roleFromGen(r)
	}

	return result
}

func resourceGroupFromGen(
	rg client.ResourceGroupResource,
) ResourceGroup {
	return ResourceGroup{
		Name:      rg.Name,
		Endpoints: rg.Endpoints,
	}
}

func resourceGroupsFromGen(
	rgs []client.ResourceGroupResource,
) []ResourceGroup {
	result := make([]ResourceGroup, len(rgs))
	for i, rg := range rgs {
		result[i] = resourceGroupFromGen(rg)
	}

	return result
}

func policyFromGen(p client.RBACPolicy) Policy {
	return Policy{
		Role:          p.Role,
		ResourceGroup: p.ResourceGroup,
		Method:        PolicyMethod(p.Method),
	}
}

func policiesFromGen(
	ps []client.RBACPolicy,
) []Policy {
	result := make([]Policy, len(ps))
	for i, p := range ps {
		result[i] = policyFromGen(p)
	}

	return result
}

func taskLogFromGen(l client.TaskLog) TaskLog {
	return TaskLog{
		Timestamp: l.Timestamp,
		Level:     l.Level,
		Issuer:    l.Issuer,
		Message:   l.Message,
	}
}

func taskLogsFromGen(
	ls []client.TaskLog,
) []TaskLog {
	result := make([]TaskLog, len(ls))
	for i, l := range ls {
		result[i] = taskLogFromGen(l)
	}

	return result
}

// --- SDK -> Generated conversions (for requests) ---

func envsToGen(
	es []EnvironmentVariable,
) []client.EnvironmentVariable {
	result := make(
		[]client.EnvironmentVariable,
		len(es),
	)

	for i, e := range es {
		result[i] = client.EnvironmentVariable{
			Key:   e.Key,
			Value: e.Value,
		}
	}

	return result
}
