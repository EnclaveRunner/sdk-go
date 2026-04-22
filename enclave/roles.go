package enclave

import (
	"context"
	"fmt"
	"iter"
	"net/http"

	"github.com/EnclaveRunner/sdk-go/client"
)

// ListRoles returns an iterator over all roles,
// fetching pages transparently. Use FilterByRole to
// narrow results.
func (c *Client) ListRoles(
	ctx context.Context,
	opts ...ListRolesOption,
) iter.Seq2[Role, error] {
	var cfg listRolesConfig
	for _, o := range opts {
		o(&cfg)
	}

	return paginate(
		func(limit, offset int) ([]Role, error) {
			resp, err := c.api.
				GetV1RbacRoleWithResponse(
					ctx,
					&client.GetV1RbacRoleParams{
						Limit:  &limit,
						Offset: &offset,
						Role:   cfg.role,
					},
				)
			if err != nil {
				return nil, fmt.Errorf(
					"listing roles: %w",
					err,
				)
			}

			msg := extractErrMessage(
				resp.Body,
				firstErrMessage(
					resp.JSON413,
				),
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

			return rolesFromGen(
				*resp.JSON200,
			), nil
		},
	)
}

// CreateRole creates a new role with the given name
// and user assignments.
func (c *Client) CreateRole(
	ctx context.Context,
	name string,
	users []string,
) (Role, error) {
	resp, err := c.api.
		PutV1RbacRoleRoleWithResponse(
			ctx,
			name,
			client.PutRoleRequest{
				Users: users,
			},
		)
	if err != nil {
		return Role{}, fmt.Errorf(
			"creating role: %w",
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
		return Role{}, err
	}

	if resp.JSON201 == nil {
		return Role{}, nil
	}

	return roleFromGen(*resp.JSON201), nil
}

// GetRole retrieves a single role by name.
func (c *Client) GetRole(
	ctx context.Context,
	name string,
) (Role, error) {
	resp, err := c.api.
		GetV1RbacRoleRoleWithResponse(
			ctx,
			name,
		)
	if err != nil {
		return Role{}, fmt.Errorf(
			"getting role: %w",
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
		return Role{}, err
	}

	if resp.JSON200 == nil {
		return Role{}, nil
	}

	return roleFromGen(*resp.JSON200), nil
}

// RoleExists checks whether a role with the given
// name exists. Returns true when the API responds
// with 200 OK, false for 404.
func (c *Client) RoleExists(
	ctx context.Context,
	name string,
) (bool, error) {
	resp, err := c.api.
		HeadV1RbacRoleRoleWithResponse(
			ctx,
			name,
		)
	if err != nil {
		return false, fmt.Errorf(
			"checking role existence: %w",
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

// DeleteRole deletes a role by name and returns the
// deleted role.
func (c *Client) DeleteRole(
	ctx context.Context,
	name string,
) (Role, error) {
	resp, err := c.api.
		DeleteV1RbacRoleRoleWithResponse(
			ctx,
			name,
		)
	if err != nil {
		return Role{}, fmt.Errorf(
			"deleting role: %w",
			err,
		)
	}

	msg := extractErrMessage(
		resp.Body,
		firstErrMessage(
			resp.JSON400,
			resp.JSON404,
			resp.JSON409,
			resp.JSON413,
		),
	)

	if err := mapHTTPError(
		resp.StatusCode(),
		msg,
	); err != nil {
		return Role{}, err
	}

	if resp.JSON200 == nil {
		return Role{}, nil
	}

	return roleFromGen(*resp.JSON200), nil
}
