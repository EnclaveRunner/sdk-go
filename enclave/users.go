package enclave

import (
	"context"
	"fmt"
	"iter"
	"net/http"

	"github.com/EnclaveRunner/sdk-go/client"
)

// ListUsers returns a paginated iterator over all
// users. Use [Collect] to drain the iterator into a
// slice.
func (c *Client) ListUsers(
	ctx context.Context,
	opts ...ListUsersOption,
) iter.Seq2[User, error] {
	cfg := listUsersConfig{}
	for _, o := range opts {
		o(&cfg)
	}

	return paginate(
		func(
			limit int,
			offset int,
		) ([]User, error) {
			resp, err := c.api.GetV1UserWithResponse(
				ctx,
				&client.GetV1UserParams{
					Limit:       &limit,
					Offset:      &offset,
					Name:        cfg.name,
					DisplayName: cfg.displayName,
				},
			)
			if err != nil {
				return nil, fmt.Errorf(
					"listing users: %w",
					err,
				)
			}

			if apiErr := mapHTTPError(
				resp.StatusCode(),
				extractErrMessage(
					resp.Body,
					firstErrMessage(
						resp.JSON400,
						resp.JSON413,
					),
				),
			); apiErr != nil {
				return nil, apiErr
			}

			if resp.JSON200 == nil {
				return nil, nil
			}

			return usersFromGen(
				*resp.JSON200,
			), nil
		},
	)
}

// CreateUser creates a new user with the given
// username, password, and display name. Use
// [WithRoles] to assign roles.
func (c *Client) CreateUser(
	ctx context.Context,
	username string,
	password string,
	displayName string,
	opts ...CreateUserOption,
) (User, error) {
	cfg := createUserConfig{}
	for _, o := range opts {
		o(&cfg)
	}

	body := client.PutUserRequest{
		DisplayName: displayName,
		Password:    password,
	}

	if cfg.roles != nil {
		body.Roles = &cfg.roles
	}

	resp, err := c.api.PutV1UserUsernameWithResponse(
		ctx,
		username,
		body,
	)
	if err != nil {
		return User{}, fmt.Errorf(
			"creating user: %w",
			err,
		)
	}

	if apiErr := mapHTTPError(
		resp.StatusCode(),
		extractErrMessage(
			resp.Body,
			firstErrMessage(
				resp.JSON400,
				resp.JSON409,
				resp.JSON413,
			),
		),
	); apiErr != nil {
		return User{}, apiErr
	}

	return userFromGen(*resp.JSON201), nil
}

// GetUser retrieves a single user by username.
func (c *Client) GetUser(
	ctx context.Context,
	username string,
) (User, error) {
	resp, err := c.api.GetV1UserUsernameWithResponse(
		ctx,
		username,
	)
	if err != nil {
		return User{}, fmt.Errorf(
			"getting user: %w",
			err,
		)
	}

	if apiErr := mapHTTPError(
		resp.StatusCode(),
		extractErrMessage(
			resp.Body,
			firstErrMessage(
				resp.JSON400,
				resp.JSON404,
				resp.JSON413,
			),
		),
	); apiErr != nil {
		return User{}, apiErr
	}

	return userFromGen(*resp.JSON200), nil
}

// UserExists checks whether a user with the given
// username exists. Returns (true, nil) for 200,
// (false, nil) for 404, and an error otherwise.
func (c *Client) UserExists(
	ctx context.Context,
	username string,
) (bool, error) {
	resp, err := c.api.HeadV1UserUsernameWithResponse(
		ctx,
		username,
	)
	if err != nil {
		return false, fmt.Errorf(
			"checking user existence: %w",
			err,
		)
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		return true, nil
	case http.StatusNotFound:
		return false, nil
	default:
		apiErr := mapHTTPError(
			resp.StatusCode(),
			extractErrMessage(
				resp.Body,
				firstErrMessage(
					resp.JSON413,
				),
			),
		)

		return false, apiErr
	}
}

// UpdateUser updates an existing user by username.
// Use [WithDisplayName], [WithPassword], or
// [WithUserRoles] to set fields.
func (c *Client) UpdateUser(
	ctx context.Context,
	username string,
	opts ...UpdateUserOption,
) (User, error) {
	cfg := updateUserConfig{}
	for _, o := range opts {
		o(&cfg)
	}

	body := client.PatchUser{
		DisplayName: cfg.displayName,
		Password:    cfg.password,
		Roles:       cfg.roles,
	}

	resp, err := c.api.PatchV1UserUsernameWithResponse(
		ctx,
		username,
		body,
	)
	if err != nil {
		return User{}, fmt.Errorf(
			"updating user: %w",
			err,
		)
	}

	if apiErr := mapHTTPError(
		resp.StatusCode(),
		extractErrMessage(
			resp.Body,
			firstErrMessage(
				resp.JSON400,
				resp.JSON404,
				resp.JSON409,
				resp.JSON413,
			),
		),
	); apiErr != nil {
		return User{}, apiErr
	}

	return userFromGen(*resp.JSON200), nil
}

// DeleteUser deletes a user by username and returns
// the deleted user.
func (c *Client) DeleteUser(
	ctx context.Context,
	username string,
) (User, error) {
	resp, err := c.api.DeleteV1UserUsernameWithResponse(
		ctx,
		username,
	)
	if err != nil {
		return User{}, fmt.Errorf(
			"deleting user: %w",
			err,
		)
	}

	if apiErr := mapHTTPError(
		resp.StatusCode(),
		extractErrMessage(
			resp.Body,
			firstErrMessage(
				resp.JSON400,
				resp.JSON404,
				resp.JSON413,
			),
		),
	); apiErr != nil {
		return User{}, apiErr
	}

	return userFromGen(*resp.JSON200), nil
}

// GetMe retrieves the currently authenticated user.
func (c *Client) GetMe(
	ctx context.Context,
) (User, error) {
	resp, err := c.api.GetV1UserMeWithResponse(ctx)
	if err != nil {
		return User{}, fmt.Errorf(
			"getting current user: %w",
			err,
		)
	}

	if apiErr := mapHTTPError(
		resp.StatusCode(),
		extractErrMessage(
			resp.Body,
			firstErrMessage(
				resp.JSON413,
			),
		),
	); apiErr != nil {
		return User{}, apiErr
	}

	return userFromGen(*resp.JSON200), nil
}

// UpdateMe updates the currently authenticated user.
// Use [WithDisplayName], [WithPassword], or
// [WithUserRoles] to set fields.
func (c *Client) UpdateMe(
	ctx context.Context,
	opts ...UpdateUserOption,
) (User, error) {
	cfg := updateUserConfig{}
	for _, o := range opts {
		o(&cfg)
	}

	body := client.PatchMe{
		DisplayName: cfg.displayName,
		Password:    cfg.password,
		Roles:       cfg.roles,
	}

	resp, err := c.api.PatchV1UserMeWithResponse(
		ctx,
		body,
	)
	if err != nil {
		return User{}, fmt.Errorf(
			"updating current user: %w",
			err,
		)
	}

	if apiErr := mapHTTPError(
		resp.StatusCode(),
		extractErrMessage(
			resp.Body,
			firstErrMessage(
				resp.JSON400,
				resp.JSON409,
				resp.JSON413,
			),
		),
	); apiErr != nil {
		return User{}, apiErr
	}

	return userFromGen(*resp.JSON200), nil
}

// DeleteMe deletes the currently authenticated user
// and returns the deleted user.
func (c *Client) DeleteMe(
	ctx context.Context,
) (User, error) {
	resp, err := c.api.DeleteV1UserMeWithResponse(ctx)
	if err != nil {
		return User{}, fmt.Errorf(
			"deleting current user: %w",
			err,
		)
	}

	if apiErr := mapHTTPError(
		resp.StatusCode(),
		extractErrMessage(
			resp.Body,
			firstErrMessage(
				resp.JSON413,
			),
		),
	); apiErr != nil {
		return User{}, apiErr
	}

	return userFromGen(*resp.JSON200), nil
}
