package enclave

import "context"

func (c *Client) IterateUsers(ctx context.Context, req ListUsersRequest, opts ...ListOption) (*Iterator[User], error) {
	return c.UsersIterator(ctx, req, opts...)
}

func (c *Client) EnumerateUsers(ctx context.Context, req ListUsersRequest, opts ...ListOption) (<-chan ItemResult[User], error) {
	return c.UsersChannel(ctx, req, opts...)
}

func (c *Client) ListAllUsers(ctx context.Context, req ListUsersRequest, opts ...ListOption) ([]User, error) {
	return c.ListUsers(ctx, req, opts...)
}

func (c *Client) IterateTasks(ctx context.Context, req ListTasksRequest, opts ...ListOption) (*Iterator[Task], error) {
	return c.TasksIterator(ctx, req, opts...)
}

func (c *Client) EnumerateTasks(ctx context.Context, req ListTasksRequest, opts ...ListOption) (<-chan ItemResult[Task], error) {
	return c.TasksChannel(ctx, req, opts...)
}

func (c *Client) ListAllTasks(ctx context.Context, req ListTasksRequest, opts ...ListOption) ([]Task, error) {
	return c.ListTasks(ctx, req, opts...)
}

func (c *Client) IterateRoles(ctx context.Context, req ListRolesRequest, opts ...ListOption) (*Iterator[Role], error) {
	return c.RolesIterator(ctx, req, opts...)
}

func (c *Client) EnumerateRoles(ctx context.Context, req ListRolesRequest, opts ...ListOption) (<-chan ItemResult[Role], error) {
	return c.RolesChannel(ctx, req, opts...)
}

func (c *Client) ListAllRoles(ctx context.Context, req ListRolesRequest, opts ...ListOption) ([]Role, error) {
	return c.ListRoles(ctx, req, opts...)
}

func (c *Client) IterateResourceGroups(ctx context.Context, req ListResourceGroupsRequest, opts ...ListOption) (*Iterator[ResourceGroup], error) {
	return c.ResourceGroupsIterator(ctx, req, opts...)
}

func (c *Client) EnumerateResourceGroups(ctx context.Context, req ListResourceGroupsRequest, opts ...ListOption) (<-chan ItemResult[ResourceGroup], error) {
	return c.ResourceGroupsChannel(ctx, req, opts...)
}

func (c *Client) ListAllResourceGroups(ctx context.Context, req ListResourceGroupsRequest, opts ...ListOption) ([]ResourceGroup, error) {
	return c.ListResourceGroups(ctx, req, opts...)
}

func (c *Client) IteratePolicies(ctx context.Context, req ListPoliciesRequest, opts ...ListOption) (*Iterator[RBACPolicy], error) {
	return c.PoliciesIterator(ctx, req, opts...)
}

func (c *Client) EnumeratePolicies(ctx context.Context, req ListPoliciesRequest, opts ...ListOption) (<-chan ItemResult[RBACPolicy], error) {
	return c.PoliciesChannel(ctx, req, opts...)
}

func (c *Client) ListAllPolicies(ctx context.Context, req ListPoliciesRequest, opts ...ListOption) ([]RBACPolicy, error) {
	return c.ListPolicies(ctx, req, opts...)
}

func (c *Client) IterateArtifacts(ctx context.Context, opts ...ListOption) (*Iterator[Artifact], error) {
	return c.ArtifactsIterator(ctx, opts...)
}

func (c *Client) EnumerateArtifacts(ctx context.Context, opts ...ListOption) (<-chan ItemResult[Artifact], error) {
	return c.ArtifactsChannel(ctx, opts...)
}

func (c *Client) ListAllArtifacts(ctx context.Context, opts ...ListOption) ([]Artifact, error) {
	return c.ListArtifacts(ctx, opts...)
}

func (c *Client) IterateNamespaceArtifacts(ctx context.Context, namespace string, opts ...ListOption) (*Iterator[Artifact], error) {
	return c.NamespaceArtifactsIterator(ctx, namespace, opts...)
}

func (c *Client) EnumerateNamespaceArtifacts(ctx context.Context, namespace string, opts ...ListOption) (<-chan ItemResult[Artifact], error) {
	return c.NamespaceArtifactsChannel(ctx, namespace, opts...)
}

func (c *Client) ListAllNamespaceArtifacts(ctx context.Context, namespace string, opts ...ListOption) ([]Artifact, error) {
	return c.ListNamespaceArtifacts(ctx, namespace, opts...)
}

func (c *Client) IterateArtifactVersions(ctx context.Context, namespace, name string, opts ...ListOption) (*Iterator[Artifact], error) {
	return c.ArtifactVersionsIterator(ctx, namespace, name, opts...)
}

func (c *Client) EnumerateArtifactVersions(ctx context.Context, namespace, name string, opts ...ListOption) (<-chan ItemResult[Artifact], error) {
	return c.ArtifactVersionsChannel(ctx, namespace, name, opts...)
}

func (c *Client) ListAllArtifactVersions(ctx context.Context, namespace, name string, opts ...ListOption) ([]Artifact, error) {
	return c.ListArtifactVersions(ctx, namespace, name, opts...)
}
