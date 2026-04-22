package enclave

import (
	"context"
	"fmt"
	"io"
	"iter"
	"net/http"

	"github.com/EnclaveRunner/sdk-go/client"
)

// ListArtifactNamespaces returns a paginated iterator
// over all artifact namespaces.
func (c *Client) ListArtifactNamespaces(
	ctx context.Context,
) iter.Seq2[Artifact, error] {
	return paginate(
		func(
			limit int,
			offset int,
		) ([]Artifact, error) {
			resp, err := c.api.
				GetV1ArtifactWithResponse(
					ctx,
					&client.GetV1ArtifactParams{
						Limit:  &limit,
						Offset: &offset,
					},
				)
			if err != nil {
				return nil, fmt.Errorf(
					"listing artifact namespaces: %w",
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

			return artifactsFromGen(
				*resp.JSON200,
			), nil
		},
	)
}

// ListArtifacts returns a paginated iterator over all
// artifacts in a namespace.
func (c *Client) ListArtifacts(
	ctx context.Context,
	namespace string,
) iter.Seq2[Artifact, error] {
	return paginate(
		func(
			limit int,
			offset int,
		) ([]Artifact, error) {
			resp, err := c.api.
				GetV1ArtifactNamespaceWithResponse(
					ctx,
					namespace,
					&client.GetV1ArtifactNamespaceParams{
						Limit:  &limit,
						Offset: &offset,
					},
				)
			if err != nil {
				return nil, fmt.Errorf(
					"listing artifacts: %w",
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

			return artifactsFromGen(
				*resp.JSON200,
			), nil
		},
	)
}

// ListArtifactVersions returns a paginated iterator
// over all versions of an artifact.
func (c *Client) ListArtifactVersions(
	ctx context.Context,
	namespace string,
	name string,
) iter.Seq2[Artifact, error] {
	return paginate(
		func(
			limit int,
			offset int,
		) ([]Artifact, error) {
			params := &client.GetV1ArtifactNamespaceNameParams{
				Limit:  &limit,
				Offset: &offset,
			}

			resp, err := c.api.
				GetV1ArtifactNamespaceNameWithResponse(
					ctx,
					namespace,
					name,
					params,
				)
			if err != nil {
				return nil, fmt.Errorf(
					"listing artifact versions: %w",
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

			return artifactsFromGen(
				*resp.JSON200,
			), nil
		},
	)
}

// UploadArtifact uploads an artifact and returns the
// resulting version hash.
func (c *Client) UploadArtifact(
	ctx context.Context,
	namespace string,
	name string,
	body io.Reader,
) (UploadResult, error) {
	contentType := "application/octet-stream"

	resp, err := c.api.
		PostV1ArtifactRawNamespaceNameWithBodyWithResponse(
			ctx,
			namespace,
			name,
			contentType,
			body,
		)
	if err != nil {
		return UploadResult{}, fmt.Errorf(
			"uploading artifact: %w",
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
		return UploadResult{}, apiErr
	}

	return UploadResult{
		VersionHash: resp.JSON201.VersionHash,
	}, nil
}

// GetArtifactByTag retrieves artifact metadata by tag.
func (c *Client) GetArtifactByTag(
	ctx context.Context,
	namespace string,
	name string,
	tag string,
) (Artifact, error) {
	resp, err := c.api.
		GetV1ArtifactNamespaceNameTagTagWithResponse(
			ctx,
			namespace,
			name,
			tag,
		)
	if err != nil {
		return Artifact{}, fmt.Errorf(
			"getting artifact by tag: %w",
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
		return Artifact{}, apiErr
	}

	return artifactFromGen(resp.JSON200), nil
}

// GetArtifactByHash retrieves artifact metadata by
// version hash.
func (c *Client) GetArtifactByHash(
	ctx context.Context,
	namespace string,
	name string,
	hash string,
) (Artifact, error) {
	resp, err := c.api.
		GetV1ArtifactNamespaceNameHashHashWithResponse(
			ctx,
			namespace,
			name,
			hash,
		)
	if err != nil {
		return Artifact{}, fmt.Errorf(
			"getting artifact by hash: %w",
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
		return Artifact{}, apiErr
	}

	return artifactFromGen(resp.JSON200), nil
}

// UpdateArtifactTagsByTag replaces the tags on the
// artifact version identified by tag.
func (c *Client) UpdateArtifactTagsByTag(
	ctx context.Context,
	namespace string,
	name string,
	tag string,
	tags []string,
) (Artifact, error) {
	body := client.PatchArtifact{Tags: &tags}

	resp, err := c.api.
		PatchV1ArtifactNamespaceNameTagTagWithResponse(
			ctx,
			namespace,
			name,
			tag,
			body,
		)
	if err != nil {
		return Artifact{}, fmt.Errorf(
			"updating artifact tags by tag: %w",
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
		return Artifact{}, apiErr
	}

	return artifactFromGen(resp.JSON200), nil
}

// UpdateArtifactTagsByHash replaces the tags on the
// artifact version identified by hash.
func (c *Client) UpdateArtifactTagsByHash(
	ctx context.Context,
	namespace string,
	name string,
	hash string,
	tags []string,
) (Artifact, error) {
	body := client.PatchArtifact{Tags: &tags}

	resp, err := c.api.
		PatchV1ArtifactNamespaceNameHashHashWithResponse(
			ctx,
			namespace,
			name,
			hash,
			body,
		)
	if err != nil {
		return Artifact{}, fmt.Errorf(
			"updating artifact tags by hash: %w",
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
		return Artifact{}, apiErr
	}

	return artifactFromGen(resp.JSON200), nil
}

// DeleteArtifactByTag deletes an artifact version by
// tag and returns the deleted artifact.
func (c *Client) DeleteArtifactByTag(
	ctx context.Context,
	namespace string,
	name string,
	tag string,
) (Artifact, error) {
	resp, err := c.api.
		DeleteV1ArtifactNamespaceNameTagTagWithResponse(
			ctx,
			namespace,
			name,
			tag,
		)
	if err != nil {
		return Artifact{}, fmt.Errorf(
			"deleting artifact by tag: %w",
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
		return Artifact{}, apiErr
	}

	return artifactFromGen(resp.JSON200), nil
}

// DeleteArtifactByHash deletes an artifact version by
// hash and returns the deleted artifact.
func (c *Client) DeleteArtifactByHash(
	ctx context.Context,
	namespace string,
	name string,
	hash string,
) (Artifact, error) {
	resp, err := c.api.
		DeleteV1ArtifactNamespaceNameHashHashWithResponse(
			ctx,
			namespace,
			name,
			hash,
		)
	if err != nil {
		return Artifact{}, fmt.Errorf(
			"deleting artifact by hash: %w",
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
		return Artifact{}, apiErr
	}

	return artifactFromGen(resp.JSON200), nil
}

// DownloadArtifactByTag downloads the raw artifact
// content by tag. The caller must close the returned
// io.ReadCloser.
func (c *Client) DownloadArtifactByTag(
	ctx context.Context,
	namespace string,
	name string,
	tag string,
) (io.ReadCloser, error) {
	resp, err := c.api.
		GetV1ArtifactRawNamespaceNameTagTag(
			ctx,
			namespace,
			name,
			tag,
		)
	if err != nil {
		return nil, fmt.Errorf(
			"downloading artifact by tag: %w",
			err,
		)
	}

	return handleDownload(resp)
}

// DownloadArtifactByHash downloads the raw artifact
// content by hash. The caller must close the returned
// io.ReadCloser.
func (c *Client) DownloadArtifactByHash(
	ctx context.Context,
	namespace string,
	name string,
	hash string,
) (io.ReadCloser, error) {
	resp, err := c.api.
		GetV1ArtifactRawNamespaceNameHashHash(
			ctx,
			namespace,
			name,
			hash,
		)
	if err != nil {
		return nil, fmt.Errorf(
			"downloading artifact by hash: %w",
			err,
		)
	}

	return handleDownload(resp)
}

// handleDownload checks the HTTP response status code
// and returns the body for success or an error for
// non-2xx responses.
func handleDownload(
	resp *http.Response,
) (io.ReadCloser, error) {
	if resp.StatusCode >= http.StatusOK &&
		resp.StatusCode < http.StatusMultipleChoices {
		return resp.Body, nil
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf(
			"reading error response: %w",
			err,
		)
	}

	msg := extractErrMessage(body, "")

	return nil, mapHTTPError(
		resp.StatusCode,
		msg,
	)
}
