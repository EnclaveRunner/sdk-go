package enclave

import (
	"context"
	"io"

	"github.com/EnclaveRunner/sdk-go/client"
)

func (c *Client) GetCurrentUser(ctx context.Context) (User, error) {
	resp, err := c.raw.GetV1UserMeWithResponse(ctx)
	if err != nil {
		return User{}, err
	}
	if resp.JSON200 != nil {
		return mapUser(*resp.JSON200), nil
	}
	return User{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON413))
}

func (c *Client) DeleteCurrentUser(ctx context.Context) (User, error) {
	resp, err := c.raw.DeleteV1UserMeWithResponse(ctx)
	if err != nil {
		return User{}, err
	}
	if resp.JSON200 != nil {
		return mapUser(*resp.JSON200), nil
	}
	return User{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON413))
}

func (c *Client) PatchCurrentUser(ctx context.Context, req PatchCurrentUserRequest) (User, error) {
	resp, err := c.raw.PatchV1UserMeWithResponse(ctx, toClientPatchMe(req))
	if err != nil {
		return User{}, err
	}
	if resp.JSON200 != nil {
		return mapUser(*resp.JSON200), nil
	}
	return User{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON413))
}

func (c *Client) GetUser(ctx context.Context, username string) (User, error) {
	resp, err := c.raw.GetV1UserUsernameWithResponse(ctx, username)
	if err != nil {
		return User{}, err
	}
	if resp.JSON200 != nil {
		return mapUser(*resp.JSON200), nil
	}
	return User{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON404), valueOrError(resp.JSON413))
}

func (c *Client) DeleteUser(ctx context.Context, username string) (User, error) {
	resp, err := c.raw.DeleteV1UserUsernameWithResponse(ctx, username)
	if err != nil {
		return User{}, err
	}
	if resp.JSON200 != nil {
		return mapUser(*resp.JSON200), nil
	}
	return User{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON404), valueOrError(resp.JSON413))
}

func (c *Client) PatchUser(ctx context.Context, username string, req PatchUserRequest) (User, error) {
	resp, err := c.raw.PatchV1UserUsernameWithResponse(ctx, username, toClientPatchUser(req))
	if err != nil {
		return User{}, err
	}
	if resp.JSON200 != nil {
		return mapUser(*resp.JSON200), nil
	}
	return User{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON404), valueOrError(resp.JSON409), valueOrError(resp.JSON413))
}

func (c *Client) CreateUser(ctx context.Context, username string, req PutUserRequest) (User, error) {
	resp, err := c.raw.PutV1UserUsernameWithResponse(ctx, username, toClientPutUser(req))
	if err != nil {
		return User{}, err
	}
	if resp.JSON201 != nil {
		return mapUser(*resp.JSON201), nil
	}
	return User{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON409), valueOrError(resp.JSON413))
}

func (c *Client) CreateTask(ctx context.Context, req CreateTaskRequest) (Task, error) {
	resp, err := c.raw.PostV1TaskWithResponse(ctx, toClientCreateTask(req))
	if err != nil {
		return Task{}, err
	}
	if resp.JSON201 != nil {
		return mapTask(*resp.JSON201), nil
	}
	return Task{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON413))
}

func (c *Client) GetTask(ctx context.Context, id string) (Task, error) {
	resp, err := c.raw.GetV1TaskIdWithResponse(ctx, id)
	if err != nil {
		return Task{}, err
	}
	if resp.JSON200 != nil {
		return mapTask(*resp.JSON200), nil
	}
	return Task{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON404))
}

func (c *Client) GetTaskLogs(ctx context.Context, id string, req GetTaskLogsRequest) ([]TaskLog, error) {
	resp, err := c.raw.GetV1TaskIdLogsWithResponse(ctx, id, toClientGetTaskLogsParams(req))
	if err != nil {
		return nil, err
	}
	if resp.JSON200 != nil {
		return mapTaskLogs(*resp.JSON200), nil
	}
	return nil, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON404))
}

func (c *Client) CreateRole(ctx context.Context, role string, req PutRoleRequest) (Role, error) {
	resp, err := c.raw.PutV1RbacRoleRoleWithResponse(ctx, role, toClientPutRole(req))
	if err != nil {
		return Role{}, err
	}
	if resp.JSON201 != nil {
		return mapRole(*resp.JSON201), nil
	}
	return Role{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON409), valueOrError(resp.JSON413))
}

func (c *Client) GetRole(ctx context.Context, role string) (Role, error) {
	resp, err := c.raw.GetV1RbacRoleRoleWithResponse(ctx, role)
	if err != nil {
		return Role{}, err
	}
	if resp.JSON200 != nil {
		return mapRole(*resp.JSON200), nil
	}
	return Role{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON404), valueOrError(resp.JSON413))
}

func (c *Client) DeleteRole(ctx context.Context, role string) (Role, error) {
	resp, err := c.raw.DeleteV1RbacRoleRoleWithResponse(ctx, role)
	if err != nil {
		return Role{}, err
	}
	if resp.JSON200 != nil {
		return mapRole(*resp.JSON200), nil
	}
	return Role{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON404), valueOrError(resp.JSON409), valueOrError(resp.JSON413))
}

func (c *Client) CreateResourceGroup(ctx context.Context, resourceGroup string, req PutResourceGroupRequest) (ResourceGroup, error) {
	resp, err := c.raw.PutV1RbacResourceGroupResourceGroupWithResponse(ctx, resourceGroup, toClientPutResourceGroup(req))
	if err != nil {
		return ResourceGroup{}, err
	}
	if resp.JSON201 != nil {
		return mapResourceGroup(*resp.JSON201), nil
	}
	return ResourceGroup{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON409), valueOrError(resp.JSON413))
}

func (c *Client) GetResourceGroup(ctx context.Context, resourceGroup string) (ResourceGroup, error) {
	resp, err := c.raw.GetV1RbacResourceGroupResourceGroupWithResponse(ctx, resourceGroup)
	if err != nil {
		return ResourceGroup{}, err
	}
	if resp.JSON200 != nil {
		return mapResourceGroup(*resp.JSON200), nil
	}
	return ResourceGroup{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON404), valueOrError(resp.JSON413))
}

func (c *Client) DeleteResourceGroup(ctx context.Context, resourceGroup string) (ResourceGroup, error) {
	resp, err := c.raw.DeleteV1RbacResourceGroupResourceGroupWithResponse(ctx, resourceGroup)
	if err != nil {
		return ResourceGroup{}, err
	}
	if resp.JSON200 != nil {
		return mapResourceGroup(*resp.JSON200), nil
	}
	return ResourceGroup{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON404), valueOrError(resp.JSON413))
}

func (c *Client) CreatePolicy(ctx context.Context, policy RBACPolicy) error {
	resp, err := c.raw.PutV1RbacPolicyWithResponse(ctx, client.RBACPolicy{
		Role:          policy.Role,
		ResourceGroup: policy.ResourceGroup,
		Method:        client.RBACPolicyMethod(policy.Method),
	})
	if err != nil {
		return err
	}
	if resp.StatusCode() == 200 {
		return nil
	}
	return newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), fieldErrorMessage(resp.JSON404), valueOrError(resp.JSON413))
}

func (c *Client) DeletePolicy(ctx context.Context, policy RBACPolicy) error {
	resp, err := c.raw.DeleteV1RbacPolicyWithResponse(ctx, client.RBACPolicy{
		Role:          policy.Role,
		ResourceGroup: policy.ResourceGroup,
		Method:        client.RBACPolicyMethod(policy.Method),
	})
	if err != nil {
		return err
	}
	if resp.StatusCode() == 200 {
		return nil
	}
	return newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON409), valueOrError(resp.JSON413))
}

func (c *Client) UploadArtifactRaw(ctx context.Context, namespace, name, contentType string, body io.Reader) (UploadArtifactResult, error) {
	resp, err := c.raw.PostV1ArtifactRawNamespaceNameWithBodyWithResponse(ctx, namespace, name, contentType, body)
	if err != nil {
		return UploadArtifactResult{}, err
	}
	if resp.JSON201 != nil {
		return UploadArtifactResult{VersionHash: resp.JSON201.VersionHash}, nil
	}
	return UploadArtifactResult{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON409), valueOrError(resp.JSON413))
}

func (c *Client) GetArtifactRawByHash(ctx context.Context, namespace, name, hash string) ([]byte, error) {
	resp, err := c.raw.GetV1ArtifactRawNamespaceNameHashHashWithResponse(ctx, namespace, name, hash)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() == 200 {
		return resp.Body, nil
	}
	return nil, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON404), valueOrError(resp.JSON413))
}

func (c *Client) GetArtifactRawByTag(ctx context.Context, namespace, name, tag string) ([]byte, error) {
	resp, err := c.raw.GetV1ArtifactRawNamespaceNameTagTagWithResponse(ctx, namespace, name, tag)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() == 200 {
		return resp.Body, nil
	}
	return nil, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON404), valueOrError(resp.JSON413))
}

func (c *Client) GetArtifactByHash(ctx context.Context, namespace, name, hash string) (Artifact, error) {
	resp, err := c.raw.GetV1ArtifactNamespaceNameHashHashWithResponse(ctx, namespace, name, hash)
	if err != nil {
		return Artifact{}, err
	}
	if resp.JSON200 != nil {
		return mapArtifact(*resp.JSON200), nil
	}
	return Artifact{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON404), valueOrError(resp.JSON413))
}

func (c *Client) DeleteArtifactByHash(ctx context.Context, namespace, name, hash string) (Artifact, error) {
	resp, err := c.raw.DeleteV1ArtifactNamespaceNameHashHashWithResponse(ctx, namespace, name, hash)
	if err != nil {
		return Artifact{}, err
	}
	if resp.JSON200 != nil {
		return mapArtifact(*resp.JSON200), nil
	}
	return Artifact{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON404), valueOrError(resp.JSON413))
}

func (c *Client) PatchArtifactByHash(ctx context.Context, namespace, name, hash string, req PatchArtifactRequest) (Artifact, error) {
	resp, err := c.raw.PatchV1ArtifactNamespaceNameHashHashWithResponse(ctx, namespace, name, hash, toClientPatchArtifact(req))
	if err != nil {
		return Artifact{}, err
	}
	if resp.JSON200 != nil {
		return mapArtifact(*resp.JSON200), nil
	}
	return Artifact{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON404), valueOrError(resp.JSON413))
}

func (c *Client) GetArtifactByTag(ctx context.Context, namespace, name, tag string) (Artifact, error) {
	resp, err := c.raw.GetV1ArtifactNamespaceNameTagTagWithResponse(ctx, namespace, name, tag)
	if err != nil {
		return Artifact{}, err
	}
	if resp.JSON200 != nil {
		return mapArtifact(*resp.JSON200), nil
	}
	return Artifact{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON404), valueOrError(resp.JSON413))
}

func (c *Client) DeleteArtifactByTag(ctx context.Context, namespace, name, tag string) (Artifact, error) {
	resp, err := c.raw.DeleteV1ArtifactNamespaceNameTagTagWithResponse(ctx, namespace, name, tag)
	if err != nil {
		return Artifact{}, err
	}
	if resp.JSON200 != nil {
		return mapArtifact(*resp.JSON200), nil
	}
	return Artifact{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON404), valueOrError(resp.JSON413))
}

func (c *Client) PatchArtifactByTag(ctx context.Context, namespace, name, tag string, req PatchArtifactRequest) (Artifact, error) {
	resp, err := c.raw.PatchV1ArtifactNamespaceNameTagTagWithResponse(ctx, namespace, name, tag, toClientPatchArtifact(req))
	if err != nil {
		return Artifact{}, err
	}
	if resp.JSON200 != nil {
		return mapArtifact(*resp.JSON200), nil
	}
	return Artifact{}, newAPIErrorFromResponse(resp.StatusCode(), resp.Body, valueOrError(resp.JSON400), valueOrError(resp.JSON404), valueOrError(resp.JSON413))
}

func valueOrError(g *client.ErrGeneric) string {
	if g == nil {
		return ""
	}
	return g.Error
}
