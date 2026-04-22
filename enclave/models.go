package enclave

import "time"

// User is the response model for a user.
type User struct {
	Name        string
	DisplayName string
	Roles       []string
}

// Artifact is metadata for an artifact version.
type Artifact struct {
	Namespace   string
	Name        string
	VersionHash string
	CreatedAt   time.Time
	Pulls       int
	Tags        []string
}

// UploadResult is returned after uploading an
// artifact.
type UploadResult struct {
	VersionHash string
}

// EnvironmentVariable is a key-value pair supplied to
// a task.
type EnvironmentVariable struct {
	Key   string
	Value string
}

// TaskLog is a single log entry for a task.
type TaskLog struct {
	Timestamp time.Time
	Level     string
	Issuer    string
	Message   string
}

// TaskStatus holds the execution state of a task.
type TaskStatus struct {
	State         string
	Retries       int
	CompletedAt   time.Time
	LastFailedAt  time.Time
	NextProcessAt time.Time
	LastError     string
	ResultPayload string
}

// Task is the response model for a task.
type Task struct {
	ID        string
	Source    string
	Status    TaskStatus
	Args      []string
	Callback  string
	Env       []EnvironmentVariable
	Params    []any
	Retention string
	Retries   int
}

// PolicyMethod is the HTTP method in an RBAC policy.
type PolicyMethod string

// PolicyMethod constants.
const (
	PolicyMethodGet    PolicyMethod = "GET"
	PolicyMethodPost   PolicyMethod = "POST"
	PolicyMethodPut    PolicyMethod = "PUT"
	PolicyMethodPatch  PolicyMethod = "PATCH"
	PolicyMethodDelete PolicyMethod = "DELETE"
	PolicyMethodHead   PolicyMethod = "HEAD"
	PolicyMethodAll    PolicyMethod = "*"
)

// Policy is an RBAC policy binding a role, resource
// group, and HTTP method.
type Policy struct {
	Role          string
	ResourceGroup string
	Method        PolicyMethod
}

// Role is a named role with its assigned users.
type Role struct {
	Name  string
	Users []string
}

// ResourceGroup is a named resource group with its
// assigned API endpoints.
type ResourceGroup struct {
	Name      string
	Endpoints []string
}

// --- Functional option types ---

// CreateUserOption configures optional fields when
// creating a user.
type CreateUserOption func(*createUserConfig)

type createUserConfig struct {
	roles []string
}

// WithRoles sets the roles for a new user.
func WithRoles(roles ...string) CreateUserOption {
	return func(c *createUserConfig) {
		c.roles = roles
	}
}

// UpdateUserOption configures optional fields when
// updating a user or the current user.
type UpdateUserOption func(*updateUserConfig)

type updateUserConfig struct {
	displayName *string
	password    *string
	roles       *[]string
}

// WithDisplayName sets the display name.
func WithDisplayName(
	name string,
) UpdateUserOption {
	return func(c *updateUserConfig) {
		c.displayName = &name
	}
}

// WithPassword sets the password.
func WithPassword(pw string) UpdateUserOption {
	return func(c *updateUserConfig) {
		c.password = &pw
	}
}

// WithUserRoles sets the roles.
func WithUserRoles(
	roles ...string,
) UpdateUserOption {
	return func(c *updateUserConfig) {
		r := roles
		c.roles = &r
	}
}

// CreateTaskOption configures optional fields when
// creating a task.
type CreateTaskOption func(*createTaskConfig)

type createTaskConfig struct {
	args      []string
	callback  *string
	env       []EnvironmentVariable
	params    []any
	retention *string
	retries   *int
}

// WithArgs sets the argument list for a task.
func WithArgs(args ...string) CreateTaskOption {
	return func(c *createTaskConfig) {
		c.args = args
	}
}

// WithCallback sets the callback URL.
func WithCallback(url string) CreateTaskOption {
	return func(c *createTaskConfig) {
		c.callback = &url
	}
}

// WithEnv sets environment variables for a task.
func WithEnv(
	env ...EnvironmentVariable,
) CreateTaskOption {
	return func(c *createTaskConfig) {
		c.env = env
	}
}

// WithParams sets the parameters for a task.
func WithParams(params ...any) CreateTaskOption {
	return func(c *createTaskConfig) {
		c.params = params
	}
}

// WithRetention sets the retention duration.
func WithRetention(d string) CreateTaskOption {
	return func(c *createTaskConfig) {
		c.retention = &d
	}
}

// WithRetries sets the maximum number of retries.
func WithRetries(n int) CreateTaskOption {
	return func(c *createTaskConfig) {
		c.retries = &n
	}
}

// ListUsersOption configures filters for ListUsers.
type ListUsersOption func(*listUsersConfig)

type listUsersConfig struct {
	name        *string
	displayName *string
}

// FilterByName filters users by exact username.
func FilterByName(name string) ListUsersOption {
	return func(c *listUsersConfig) {
		c.name = &name
	}
}

// FilterByDisplayName filters users by display name.
func FilterByDisplayName(
	name string,
) ListUsersOption {
	return func(c *listUsersConfig) {
		c.displayName = &name
	}
}

// ListRolesOption configures filters for ListRoles.
type ListRolesOption func(*listRolesConfig)

type listRolesConfig struct {
	role *string
}

// FilterByRole filters roles by name.
func FilterByRole(role string) ListRolesOption {
	return func(c *listRolesConfig) {
		c.role = &role
	}
}

// ListResourceGroupsOption configures filters for
// ListResourceGroups.
type ListResourceGroupsOption func(
	*listResourceGroupsConfig,
)

type listResourceGroupsConfig struct {
	role *string
}

// FilterResourceGroupsByRole filters resource groups
// by associated role.
func FilterResourceGroupsByRole(
	role string,
) ListResourceGroupsOption {
	return func(c *listResourceGroupsConfig) {
		c.role = &role
	}
}

// ListPoliciesOption configures filters for
// ListPolicies.
type ListPoliciesOption func(*listPoliciesConfig)

type listPoliciesConfig struct {
	role          *string
	resourceGroup *string
	method        *string
}

// FilterPolicyByRole filters policies by role.
func FilterPolicyByRole(
	role string,
) ListPoliciesOption {
	return func(c *listPoliciesConfig) {
		c.role = &role
	}
}

// FilterPolicyByResourceGroup filters policies by
// resource group.
func FilterPolicyByResourceGroup(
	rg string,
) ListPoliciesOption {
	return func(c *listPoliciesConfig) {
		c.resourceGroup = &rg
	}
}

// FilterPolicyByMethod filters policies by HTTP
// method.
func FilterPolicyByMethod(
	method string,
) ListPoliciesOption {
	return func(c *listPoliciesConfig) {
		c.method = &method
	}
}

// ListTasksOption configures filters for ListTasks.
type ListTasksOption func(*listTasksConfig)

type listTasksConfig struct {
	state *string
}

// FilterByState filters tasks by state.
func FilterByState(state string) ListTasksOption {
	return func(c *listTasksConfig) {
		c.state = &state
	}
}

// TaskLogOption configures filters for GetTaskLogs.
type TaskLogOption func(*taskLogConfig)

type taskLogConfig struct {
	level         *string
	issuer        *string
	timeRangeFrom *time.Time
	timeRangeTo   *time.Time
}

// FilterLogByLevel filters logs by level.
func FilterLogByLevel(level string) TaskLogOption {
	return func(c *taskLogConfig) {
		c.level = &level
	}
}

// FilterLogByIssuer filters logs by issuer.
func FilterLogByIssuer(issuer string) TaskLogOption {
	return func(c *taskLogConfig) {
		c.issuer = &issuer
	}
}

// FilterLogByTimeRange filters logs by time range.
func FilterLogByTimeRange(
	from time.Time,
	to time.Time,
) TaskLogOption {
	return func(c *taskLogConfig) {
		c.timeRangeFrom = &from
		c.timeRangeTo = &to
	}
}
