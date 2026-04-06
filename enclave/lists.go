package enclave

import (
	"context"
	"fmt"

	"github.com/EnclaveRunner/sdk-go/client"
)

func (c *Client) UsersIterator(ctx context.Context, req ListUsersRequest, opts ...ListOption) (*Iterator[User], error) {
	pageSize, err := resolvePageSize(c.defaultPageSize, opts...)
	if err != nil {
		return nil, err
	}
	fetch := func(ctx context.Context, limit int, offset int) ([]User, error) {
		resp, err := c.raw.GetV1UserWithResponse(ctx, buildGetUsersParams(req, limit, offset))
		if err != nil {
			return nil, err
		}
		if resp.JSON200 != nil {
			return mapUsers(*resp.JSON200), nil
		}
		return nil, mapCommonListError(resp.StatusCode(), resp.JSON400, resp.JSON413)
	}
	return newIterator(ctx, pageSize, fetch), nil
}

func (c *Client) UsersChannel(ctx context.Context, req ListUsersRequest, opts ...ListOption) (<-chan ItemResult[User], error) {
	it, err := c.UsersIterator(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return enumerate(ctx, it), nil
}

func (c *Client) ListUsers(ctx context.Context, req ListUsersRequest, opts ...ListOption) ([]User, error) {
	it, err := c.UsersIterator(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return collectAll(it)
}

func (c *Client) TasksIterator(ctx context.Context, req ListTasksRequest, opts ...ListOption) (*Iterator[Task], error) {
	pageSize, err := resolvePageSize(c.defaultPageSize, opts...)
	if err != nil {
		return nil, err
	}
	fetch := func(ctx context.Context, limit int, offset int) ([]Task, error) {
		resp, err := c.raw.GetV1TaskWithResponse(ctx, buildGetTasksParams(req, limit, offset))
		if err != nil {
			return nil, err
		}
		if resp.JSON200 != nil {
			return mapTasks(*resp.JSON200), nil
		}
		return nil, mapTaskListError(resp.StatusCode(), resp.JSON400)
	}
	return newIterator(ctx, pageSize, fetch), nil
}

func (c *Client) TasksChannel(ctx context.Context, req ListTasksRequest, opts ...ListOption) (<-chan ItemResult[Task], error) {
	it, err := c.TasksIterator(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return enumerate(ctx, it), nil
}

func (c *Client) ListTasks(ctx context.Context, req ListTasksRequest, opts ...ListOption) ([]Task, error) {
	it, err := c.TasksIterator(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return collectAll(it)
}

func (c *Client) RolesIterator(ctx context.Context, req ListRolesRequest, opts ...ListOption) (*Iterator[Role], error) {
	pageSize, err := resolvePageSize(c.defaultPageSize, opts...)
	if err != nil {
		return nil, err
	}
	fetch := func(ctx context.Context, limit int, offset int) ([]Role, error) {
		resp, err := c.raw.GetV1RbacRoleWithResponse(ctx, buildGetRolesParams(req, limit, offset))
		if err != nil {
			return nil, err
		}
		if resp.JSON200 != nil {
			return mapRoles(*resp.JSON200), nil
		}
		return nil, mapTooLargeOnlyError(resp.StatusCode(), resp.JSON413)
	}
	return newIterator(ctx, pageSize, fetch), nil
}

func (c *Client) RolesChannel(ctx context.Context, req ListRolesRequest, opts ...ListOption) (<-chan ItemResult[Role], error) {
	it, err := c.RolesIterator(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return enumerate(ctx, it), nil
}

func (c *Client) ListRoles(ctx context.Context, req ListRolesRequest, opts ...ListOption) ([]Role, error) {
	it, err := c.RolesIterator(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return collectAll(it)
}

func (c *Client) ResourceGroupsIterator(ctx context.Context, req ListResourceGroupsRequest, opts ...ListOption) (*Iterator[ResourceGroup], error) {
	pageSize, err := resolvePageSize(c.defaultPageSize, opts...)
	if err != nil {
		return nil, err
	}
	fetch := func(ctx context.Context, limit int, offset int) ([]ResourceGroup, error) {
		resp, err := c.raw.GetV1RbacResourceGroupWithResponse(ctx, buildGetResourceGroupsParams(req, limit, offset))
		if err != nil {
			return nil, err
		}
		if resp.JSON200 != nil {
			return mapResourceGroups(*resp.JSON200), nil
		}
		return nil, mapTooLargeOnlyError(resp.StatusCode(), resp.JSON413)
	}
	return newIterator(ctx, pageSize, fetch), nil
}

func (c *Client) ResourceGroupsChannel(ctx context.Context, req ListResourceGroupsRequest, opts ...ListOption) (<-chan ItemResult[ResourceGroup], error) {
	it, err := c.ResourceGroupsIterator(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return enumerate(ctx, it), nil
}

func (c *Client) ListResourceGroups(ctx context.Context, req ListResourceGroupsRequest, opts ...ListOption) ([]ResourceGroup, error) {
	it, err := c.ResourceGroupsIterator(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return collectAll(it)
}

func (c *Client) PoliciesIterator(ctx context.Context, req ListPoliciesRequest, opts ...ListOption) (*Iterator[RBACPolicy], error) {
	pageSize, err := resolvePageSize(c.defaultPageSize, opts...)
	if err != nil {
		return nil, err
	}
	fetch := func(ctx context.Context, limit int, offset int) ([]RBACPolicy, error) {
		resp, err := c.raw.GetV1RbacPolicyWithResponse(ctx, buildGetPoliciesParams(req, limit, offset))
		if err != nil {
			return nil, err
		}
		if resp.JSON200 != nil {
			return mapPolicies(*resp.JSON200), nil
		}
		return nil, mapTooLargeOnlyError(resp.StatusCode(), resp.JSON413)
	}
	return newIterator(ctx, pageSize, fetch), nil
}

func (c *Client) PoliciesChannel(ctx context.Context, req ListPoliciesRequest, opts ...ListOption) (<-chan ItemResult[RBACPolicy], error) {
	it, err := c.PoliciesIterator(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return enumerate(ctx, it), nil
}

func (c *Client) ListPolicies(ctx context.Context, req ListPoliciesRequest, opts ...ListOption) ([]RBACPolicy, error) {
	it, err := c.PoliciesIterator(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return collectAll(it)
}

func (c *Client) ArtifactsIterator(ctx context.Context, opts ...ListOption) (*Iterator[Artifact], error) {
	pageSize, err := resolvePageSize(c.defaultPageSize, opts...)
	if err != nil {
		return nil, err
	}
	fetch := func(ctx context.Context, limit int, offset int) ([]Artifact, error) {
		resp, err := c.raw.GetV1ArtifactWithResponse(ctx, buildGetArtifactsParams(limit, offset))
		if err != nil {
			return nil, err
		}
		if resp.JSON200 != nil {
			return mapArtifacts(*resp.JSON200), nil
		}
		return nil, mapCommonListError(resp.StatusCode(), resp.JSON400, resp.JSON413)
	}
	return newIterator(ctx, pageSize, fetch), nil
}

func (c *Client) ArtifactsChannel(ctx context.Context, opts ...ListOption) (<-chan ItemResult[Artifact], error) {
	it, err := c.ArtifactsIterator(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return enumerate(ctx, it), nil
}

func (c *Client) ListArtifacts(ctx context.Context, opts ...ListOption) ([]Artifact, error) {
	it, err := c.ArtifactsIterator(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return collectAll(it)
}

func (c *Client) NamespaceArtifactsIterator(ctx context.Context, namespace string, opts ...ListOption) (*Iterator[Artifact], error) {
	pageSize, err := resolvePageSize(c.defaultPageSize, opts...)
	if err != nil {
		return nil, err
	}
	fetch := func(ctx context.Context, limit int, offset int) ([]Artifact, error) {
		resp, err := c.raw.GetV1ArtifactNamespaceWithResponse(ctx, namespace, buildGetNamespaceArtifactsParams(limit, offset))
		if err != nil {
			return nil, err
		}
		if resp.JSON200 != nil {
			return mapArtifacts(*resp.JSON200), nil
		}
		return nil, mapCommonListError(resp.StatusCode(), resp.JSON400, resp.JSON413)
	}
	return newIterator(ctx, pageSize, fetch), nil
}

func (c *Client) NamespaceArtifactsChannel(ctx context.Context, namespace string, opts ...ListOption) (<-chan ItemResult[Artifact], error) {
	it, err := c.NamespaceArtifactsIterator(ctx, namespace, opts...)
	if err != nil {
		return nil, err
	}
	return enumerate(ctx, it), nil
}

func (c *Client) ListNamespaceArtifacts(ctx context.Context, namespace string, opts ...ListOption) ([]Artifact, error) {
	it, err := c.NamespaceArtifactsIterator(ctx, namespace, opts...)
	if err != nil {
		return nil, err
	}
	return collectAll(it)
}

func (c *Client) ArtifactVersionsIterator(ctx context.Context, namespace, name string, opts ...ListOption) (*Iterator[Artifact], error) {
	pageSize, err := resolvePageSize(c.defaultPageSize, opts...)
	if err != nil {
		return nil, err
	}
	fetch := func(ctx context.Context, limit int, offset int) ([]Artifact, error) {
		resp, err := c.raw.GetV1ArtifactNamespaceNameWithResponse(ctx, namespace, name, buildGetArtifactVersionsParams(limit, offset))
		if err != nil {
			return nil, err
		}
		if resp.JSON200 != nil {
			return mapArtifacts(*resp.JSON200), nil
		}
		return nil, mapCommonListError(resp.StatusCode(), resp.JSON400, resp.JSON413)
	}
	return newIterator(ctx, pageSize, fetch), nil
}

func (c *Client) ArtifactVersionsChannel(ctx context.Context, namespace, name string, opts ...ListOption) (<-chan ItemResult[Artifact], error) {
	it, err := c.ArtifactVersionsIterator(ctx, namespace, name, opts...)
	if err != nil {
		return nil, err
	}
	return enumerate(ctx, it), nil
}

func (c *Client) ListArtifactVersions(ctx context.Context, namespace, name string, opts ...ListOption) ([]Artifact, error) {
	it, err := c.ArtifactVersionsIterator(ctx, namespace, name, opts...)
	if err != nil {
		return nil, err
	}
	return collectAll(it)
}

func buildGetUsersParams(req ListUsersRequest, limit int, offset int) *client.GetV1UserParams {
	params := &client.GetV1UserParams{Limit: intPtr(limit), Offset: intPtr(offset)}
	if req.Name != "" {
		params.Name = stringPtr(req.Name)
	}
	if req.DisplayName != "" {
		params.DisplayName = stringPtr(req.DisplayName)
	}
	return params
}

func buildGetTasksParams(req ListTasksRequest, limit int, offset int) *client.GetV1TaskParams {
	params := &client.GetV1TaskParams{Limit: intPtr(limit), Offset: intPtr(offset)}
	if req.State != "" {
		params.State = stringPtr(req.State)
	}
	return params
}

func buildGetRolesParams(req ListRolesRequest, limit int, offset int) *client.GetV1RbacRoleParams {
	params := &client.GetV1RbacRoleParams{Limit: intPtr(limit), Offset: intPtr(offset)}
	if req.Role != "" {
		params.Role = stringPtr(req.Role)
	}
	return params
}

func buildGetResourceGroupsParams(req ListResourceGroupsRequest, limit int, offset int) *client.GetV1RbacResourceGroupParams {
	params := &client.GetV1RbacResourceGroupParams{Limit: intPtr(limit), Offset: intPtr(offset)}
	if req.Role != "" {
		params.Role = stringPtr(req.Role)
	}
	return params
}

func buildGetPoliciesParams(req ListPoliciesRequest, limit int, offset int) *client.GetV1RbacPolicyParams {
	params := &client.GetV1RbacPolicyParams{Limit: intPtr(limit), Offset: intPtr(offset)}
	if req.Role != "" {
		params.Role = stringPtr(req.Role)
	}
	if req.ResourceGroup != "" {
		params.ResourceGroup = stringPtr(req.ResourceGroup)
	}
	if req.Method != "" {
		params.Method = stringPtr(req.Method)
	}
	return params
}

func buildGetArtifactsParams(limit int, offset int) *client.GetV1ArtifactParams {
	return &client.GetV1ArtifactParams{Limit: intPtr(limit), Offset: intPtr(offset)}
}

func buildGetNamespaceArtifactsParams(limit int, offset int) *client.GetV1ArtifactNamespaceParams {
	return &client.GetV1ArtifactNamespaceParams{Limit: intPtr(limit), Offset: intPtr(offset)}
}

func buildGetArtifactVersionsParams(limit int, offset int) *client.GetV1ArtifactNamespaceNameParams {
	return &client.GetV1ArtifactNamespaceNameParams{Limit: intPtr(limit), Offset: intPtr(offset)}
}

func mapCommonListError(statusCode int, badRequest *client.GenericBadRequest, tooLarge *client.GenericTooLarge) error {
	if badRequest != nil {
		return newAPIError(statusCode, badRequest.Error)
	}
	if tooLarge != nil {
		return newAPIError(statusCode, tooLarge.Error)
	}
	return newAPIError(statusCode, fmt.Sprintf("unexpected response status %d", statusCode))
}

func mapTaskListError(statusCode int, badRequest *client.GenericBadRequest) error {
	if badRequest != nil {
		return newAPIError(statusCode, badRequest.Error)
	}
	return newAPIError(statusCode, fmt.Sprintf("unexpected response status %d", statusCode))
}

func mapTooLargeOnlyError(statusCode int, tooLarge *client.GenericTooLarge) error {
	if tooLarge != nil {
		return newAPIError(statusCode, tooLarge.Error)
	}
	return newAPIError(statusCode, fmt.Sprintf("unexpected response status %d", statusCode))
}

func stringPtr(v string) *string { return &v }
func intPtr(v int) *int { return &v }
