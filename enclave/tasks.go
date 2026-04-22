package enclave

import (
	"context"
	"fmt"
	"iter"
	"net/http"

	"github.com/EnclaveRunner/sdk-go/client"
)

// ListTasks returns a paginated iterator over all
// tasks. Use [FilterByState] to narrow results.
func (c *Client) ListTasks(
	ctx context.Context,
	opts ...ListTasksOption,
) iter.Seq2[Task, error] {
	var cfg listTasksConfig
	for _, o := range opts {
		o(&cfg)
	}

	return paginate(
		func(
			limit int,
			offset int,
		) ([]Task, error) {
			resp, err := c.api.
				GetV1TaskWithResponse(
					ctx,
					&client.GetV1TaskParams{
						Limit:  &limit,
						Offset: &offset,
						State:  cfg.state,
					},
				)
			if err != nil {
				return nil, fmt.Errorf(
					"listing tasks: %w",
					err,
				)
			}

			msg := extractErrMessage(
				resp.Body,
				firstErrMessage(
					resp.JSON400,
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

			return tasksFromGen(
				*resp.JSON200,
			), nil
		},
	)
}

// CreateTask creates a new task from the given
// source. Use [WithArgs], [WithCallback], [WithEnv],
// [WithParams], [WithRetention], or [WithRetries]
// to configure optional fields.
func (c *Client) CreateTask(
	ctx context.Context,
	source string,
	opts ...CreateTaskOption,
) (Task, error) {
	var cfg createTaskConfig
	for _, o := range opts {
		o(&cfg)
	}

	req := client.CreateTaskRequest{
		Source: source,
	}

	if cfg.args != nil {
		req.Args = &cfg.args
	}

	if cfg.callback != nil {
		req.Callback = cfg.callback
	}

	if cfg.env != nil {
		genEnvs := envsToGen(cfg.env)
		req.Env = &genEnvs
	}

	if cfg.params != nil {
		req.Params = &cfg.params
	}

	if cfg.retention != nil {
		req.Retention = cfg.retention
	}

	if cfg.retries != nil {
		req.Retries = cfg.retries
	}

	resp, err := c.api.
		PostV1TaskWithResponse(ctx, req)
	if err != nil {
		return Task{}, fmt.Errorf(
			"creating task: %w",
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
		return Task{}, err
	}

	if resp.StatusCode() != http.StatusCreated {
		return Task{}, mapHTTPError(
			http.StatusInternalServerError,
			"unexpected status",
		)
	}

	return taskFromGen(resp.JSON201), nil
}

// GetTask retrieves a single task by ID.
func (c *Client) GetTask(
	ctx context.Context,
	id string,
) (Task, error) {
	resp, err := c.api.
		GetV1TaskIdWithResponse(ctx, id)
	if err != nil {
		return Task{}, fmt.Errorf(
			"getting task: %w",
			err,
		)
	}

	msg := extractErrMessage(
		resp.Body,
		firstErrMessage(
			resp.JSON400,
			resp.JSON404,
		),
	)

	if err := mapHTTPError(
		resp.StatusCode(),
		msg,
	); err != nil {
		return Task{}, err
	}

	return taskFromGen(resp.JSON200), nil
}

// GetTaskLogs retrieves the log entries for a task.
// Use [FilterLogByLevel], [FilterLogByIssuer], or
// [FilterLogByTimeRange] to narrow results.
func (c *Client) GetTaskLogs(
	ctx context.Context,
	id string,
	opts ...TaskLogOption,
) ([]TaskLog, error) {
	var cfg taskLogConfig
	for _, o := range opts {
		o(&cfg)
	}

	resp, err := c.api.
		GetV1TaskIdLogsWithResponse(
			ctx,
			id,
			&client.GetV1TaskIdLogsParams{
				Level:         cfg.level,
				Issuer:        cfg.issuer,
				TimeRangeFrom: cfg.timeRangeFrom,
				TimeRangeTo:   cfg.timeRangeTo,
			},
		)
	if err != nil {
		return nil, fmt.Errorf(
			"getting task logs: %w",
			err,
		)
	}

	msg := extractErrMessage(
		resp.Body,
		firstErrMessage(
			resp.JSON400,
			resp.JSON404,
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

	return taskLogsFromGen(*resp.JSON200), nil
}
