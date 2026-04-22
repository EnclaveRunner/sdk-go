package enclave

import (
	"context"
	"fmt"
	"iter"
	"net/http"

	"github.com/EnclaveRunner/sdk-go/client"
)

// ListResourceGroups returns an iterator over all
// resource groups. Use FilterResourceGroupsByRole to
// narrow results.
func (c *Client) ListResourceGroups(
	ctx context.Context,
	opts ...ListResourceGroupsOption,
) iter.Seq2[ResourceGroup, error] {
	var cfg listResourceGroupsConfig
	for _, o := range opts {
		o(&cfg)
	}

	return paginate(
		func(
			limit int,
			offset int,
		) ([]ResourceGroup, error) {
			resp, err := c.api.
				GetV1RbacResourceGroupWithResponse(
					ctx,
					&client.GetV1RbacResourceGroupParams{
						Limit:  &limit,
						Offset: &offset,
						Role:   cfg.role,
					},
				)
			if err != nil {
				return nil, fmt.Errorf(
					"listing resource groups: %w",
					err,
				)
			}

			msg := firstErrMessage(
				resp.JSON413,
			)
			if err := mapHTTPError(
				resp.StatusCode(),
				msg,
			); err != nil {
				return nil, err
			}

			if resp.JSON200 == nil {
				return nil, nil
			}

			return resourceGroupsFromGen(
				*resp.JSON200,
			), nil
		},
	)
}

// CreateResourceGroup creates a new resource group
// with the given name and endpoint list.
func (c *Client) CreateResourceGroup(
	ctx context.Context,
	name string,
	endpoints []string,
) (ResourceGroup, error) {
	body := client.PutResourceGroupRequest{
		Endpoints: endpoints,
	}

	resp, err := c.api.
		PutV1RbacResourceGroupResourceGroupWithResponse(
			ctx,
			name,
			body,
		)
	if err != nil {
		return ResourceGroup{}, fmt.Errorf(
			"creating resource group: %w",
			err,
		)
	}

	msg := extractErrMessage(
		resp.Body,
		firstErrMessage(
			resp.JSON400,
			resp.JSON409,
			resp.JSON413,
		),
	)

	if err := mapHTTPError(
		resp.StatusCode(),
		msg,
	); err != nil {
		return ResourceGroup{}, err
	}

	if resp.StatusCode() != http.StatusCreated {
		return ResourceGroup{}, mapHTTPError(
			http.StatusInternalServerError,
			"unexpected status",
		)
	}

	return resourceGroupFromGen(
		*resp.JSON201,
	), nil
}

// GetResourceGroup retrieves a single resource group
// by name.
func (c *Client) GetResourceGroup(
	ctx context.Context,
	name string,
) (ResourceGroup, error) {
	resp, err := c.api.
		GetV1RbacResourceGroupResourceGroupWithResponse(
			ctx,
			name,
		)
	if err != nil {
		return ResourceGroup{}, fmt.Errorf(
			"getting resource group: %w",
			err,
		)
	}

	msg := extractErrMessage(
		resp.Body,
		firstErrMessage(
			resp.JSON400,
			resp.JSON404,
			resp.JSON413,
		),
	)

	if err := mapHTTPError(
		resp.StatusCode(),
		msg,
	); err != nil {
		return ResourceGroup{}, err
	}

	return resourceGroupFromGen(
		*resp.JSON200,
	), nil
}

// ResourceGroupExists checks whether a resource group
// with the given name exists.
func (c *Client) ResourceGroupExists(
	ctx context.Context,
	name string,
) (bool, error) {
	resp, err := c.api.
		HeadV1RbacResourceGroupResourceGroupWithResponse(
			ctx,
			name,
		)
	if err != nil {
		return false, fmt.Errorf(
			"checking resource group: %w",
			err,
		)
	}

	code := resp.StatusCode()

	if code == http.StatusOK {
		return true, nil
	}

	if code == http.StatusNotFound {
		return false, nil
	}

	msg := extractErrMessage(
		resp.Body,
		firstErrMessage(resp.JSON413),
	)

	return false, mapHTTPError(code, msg)
}

// DeleteResourceGroup deletes a resource group by
// name and returns the deleted resource.
func (c *Client) DeleteResourceGroup(
	ctx context.Context,
	name string,
) (ResourceGroup, error) {
	resp, err := c.api.
		DeleteV1RbacResourceGroupResourceGroupWithResponse(
			ctx,
			name,
		)
	if err != nil {
		return ResourceGroup{}, fmt.Errorf(
			"deleting resource group: %w",
			err,
		)
	}

	msg := extractErrMessage(
		resp.Body,
		firstErrMessage(
			resp.JSON400,
			resp.JSON404,
			resp.JSON413,
		),
	)

	if err := mapHTTPError(
		resp.StatusCode(),
		msg,
	); err != nil {
		return ResourceGroup{}, err
	}

	return resourceGroupFromGen(
		*resp.JSON200,
	), nil
}
