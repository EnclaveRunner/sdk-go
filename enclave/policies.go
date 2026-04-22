package enclave

import (
	"context"
	"fmt"
	"iter"
	"net/http"

	"github.com/EnclaveRunner/sdk-go/client"
)

// ListPolicies returns an iterator over all RBAC
// policies, fetching pages transparently. Use
// [FilterPolicyByRole], [FilterPolicyByResourceGroup],
// or [FilterPolicyByMethod] to narrow results.
func (c *Client) ListPolicies(
	ctx context.Context,
	opts ...ListPoliciesOption,
) iter.Seq2[Policy, error] {
	var cfg listPoliciesConfig
	for _, o := range opts {
		o(&cfg)
	}

	return paginate(
		func(
			limit int,
			offset int,
		) ([]Policy, error) {
			resp, err := c.api.
				GetV1RbacPolicyWithResponse(
					ctx,
					&client.GetV1RbacPolicyParams{
						Limit:         &limit,
						Offset:        &offset,
						Role:          cfg.role,
						ResourceGroup: cfg.resourceGroup,
						Method:        cfg.method,
					},
				)
			if err != nil {
				return nil, fmt.Errorf(
					"listing policies: %w",
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

			return policiesFromGen(
				*resp.JSON200,
			), nil
		},
	)
}

// CreatePolicy creates a new RBAC policy binding a
// role, resource group, and HTTP method. Success is
// indicated by a 201 status with no response body.
func (c *Client) CreatePolicy(
	ctx context.Context,
	policy Policy,
) error {
	genPolicy := client.RBACPolicy{
		Role:          policy.Role,
		ResourceGroup: policy.ResourceGroup,
		Method: client.RBACPolicyMethod(
			policy.Method,
		),
	}

	resp, err := c.api.
		PutV1RbacPolicyWithResponse(
			ctx,
			genPolicy,
		)
	if err != nil {
		return fmt.Errorf(
			"creating policy: %w",
			err,
		)
	}

	msg := extractErrMessage(
		resp.Body,
		firstErrMessage(
			resp.JSON400,
			resp.JSON413,
		),
	)

	if err := mapHTTPError(
		resp.StatusCode(),
		msg,
	); err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusCreated {
		return mapHTTPError(
			http.StatusInternalServerError,
			"unexpected status",
		)
	}

	return nil
}

// DeletePolicy deletes an RBAC policy. Success is
// indicated by a 200 status with no response body.
func (c *Client) DeletePolicy(
	ctx context.Context,
	policy Policy,
) error {
	genPolicy := client.RBACPolicy{
		Role:          policy.Role,
		ResourceGroup: policy.ResourceGroup,
		Method: client.RBACPolicyMethod(
			policy.Method,
		),
	}

	resp, err := c.api.
		DeleteV1RbacPolicyWithResponse(
			ctx,
			genPolicy,
		)
	if err != nil {
		return fmt.Errorf(
			"deleting policy: %w",
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
		return err
	}

	return nil
}
