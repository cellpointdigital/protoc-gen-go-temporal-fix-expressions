// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 1.13.3-next (63174f1c48e0790b0a0d555228ee3a9390371a3b)
//	go go1.22.2
//	protoc (unknown)
//
// source: test/option/v1/option.proto
package optionv1xns

import (
	"context"
	"errors"
	"fmt"
	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	xnsv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/xns/v1"
	v1 "github.com/cludden/protoc-gen-go-temporal/gen/test/option/v1"
	expression "github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	xns "github.com/cludden/protoc-gen-go-temporal/pkg/xns"
	uuid "github.com/google/uuid"
	enumsv1 "go.temporal.io/api/enums/v1"
	updatev1 "go.temporal.io/api/update/v1"
	activity "go.temporal.io/sdk/activity"
	client "go.temporal.io/sdk/client"
	temporal "go.temporal.io/sdk/temporal"
	worker "go.temporal.io/sdk/worker"
	workflow "go.temporal.io/sdk/workflow"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	"time"
)

// TestOptions is used to configure test.option.v1.Test xns activity registration
type TestOptions struct {
	// errorConverter is used to customize error
	errorConverter func(error) error
	// filter is used to filter xns activity registrations. It receives as
	// input the original activity name, and should return one of the following:
	// 1. the original activity name, for no changes
	// 2. a modified activity name, to override the original activity name
	// 3. an empty string, to skip registration
	filter func(string) string
}

// NewTestOptions initializes a new TestOptions value
func NewTestOptions() *TestOptions {
	return &TestOptions{}
}

// WithErrorConverter overrides the default error converter applied to xns activity errors
func (opts *TestOptions) WithErrorConverter(errorConverter func(error) error) *TestOptions {
	opts.errorConverter = errorConverter
	return opts
}

// Filter is used to filter registered xns activities or customize their name
func (opts *TestOptions) WithFilter(filter func(string) string) *TestOptions {
	opts.filter = filter
	return opts
}

// convertError is applied to all xns activity errors
func (opts *TestOptions) convertError(err error) error {
	if err == nil {
		return nil
	}
	if opts != nil && opts.errorConverter != nil {
		return opts.errorConverter(err)
	}
	return xns.ErrorToApplicationError(err)
}

// filterActivity is used to filter xns activity registrations
func (opts *TestOptions) filterActivity(name string) string {
	if opts == nil || opts.filter == nil {
		return name
	}
	return opts.filter(name)
}

// testOptions is a reference to the TestOptions initialized at registration
var testOptions *TestOptions

// RegisterTestActivities registers test.option.v1.Test cross-namespace activities
func RegisterTestActivities(r worker.ActivityRegistry, c v1.TestClient, options ...*TestOptions) {
	if testOptions == nil && len(options) > 0 && options[0] != nil {
		testOptions = options[0]
	}
	a := &testActivities{c}
	if name := testOptions.filterActivity("test.option.v1.Test.CancelWorkflow"); name != "" {
		r.RegisterActivityWithOptions(a.CancelWorkflow, activity.RegisterOptions{Name: name})
	}
	if name := testOptions.filterActivity(v1.WorkflowWithInputWorkflowName); name != "" {
		r.RegisterActivityWithOptions(a.WorkflowWithInput, activity.RegisterOptions{Name: name})
	}
	if name := testOptions.filterActivity(v1.UpdateWithInputUpdateName); name != "" {
		r.RegisterActivityWithOptions(a.UpdateWithInput, activity.RegisterOptions{Name: name})
	}
}

// WorkflowWithInputWorkflowOptions are used to configure a(n) test.option.v1.Test.WorkflowWithInput workflow execution
type WorkflowWithInputWorkflowOptions struct {
	ActivityOptions      *workflow.ActivityOptions
	Detached             bool
	HeartbeatInterval    time.Duration
	ParentClosePolicy    enumsv1.ParentClosePolicy
	StartWorkflowOptions *client.StartWorkflowOptions
}

// NewWorkflowWithInputWorkflowOptions initializes a new WorkflowWithInputWorkflowOptions value
func NewWorkflowWithInputWorkflowOptions() *WorkflowWithInputWorkflowOptions {
	return &WorkflowWithInputWorkflowOptions{}
}

// WithActivityOptions can be used to customize the activity options
func (opts *WorkflowWithInputWorkflowOptions) WithActivityOptions(ao workflow.ActivityOptions) *WorkflowWithInputWorkflowOptions {
	opts.ActivityOptions = &ao
	return opts
}

// WithDetached can be used to start a workflow execution and exit immediately
func (opts *WorkflowWithInputWorkflowOptions) WithDetached(d bool) *WorkflowWithInputWorkflowOptions {
	opts.Detached = d
	return opts
}

// WithHeartbeatInterval can be used to customize the activity heartbeat interval
func (opts *WorkflowWithInputWorkflowOptions) WithHeartbeatInterval(d time.Duration) *WorkflowWithInputWorkflowOptions {
	opts.HeartbeatInterval = d
	return opts
}

// WithParentClosePolicy can be used to customize the cancellation propagation behavior
func (opts *WorkflowWithInputWorkflowOptions) WithParentClosePolicy(policy enumsv1.ParentClosePolicy) *WorkflowWithInputWorkflowOptions {
	opts.ParentClosePolicy = policy
	return opts
}

// WithStartWorkflowOptions can be used to customize the start workflow options
func (opts *WorkflowWithInputWorkflowOptions) WithStartWorkflow(swo client.StartWorkflowOptions) *WorkflowWithInputWorkflowOptions {
	opts.StartWorkflowOptions = &swo
	return opts
}

// WorkflowWithInputRun provides a handle to a test.option.v1.Test.WorkflowWithInput workflow execution
type WorkflowWithInputRun interface {
	// Cancel cancels the workflow
	Cancel(workflow.Context) error

	// Future returns the inner workflow.Future
	Future() workflow.Future

	// Get returns the inner workflow.Future
	Get(workflow.Context) error

	// ID returns the workflow id
	ID() string

	// UpdateWithInput executes a(n) test.option.v1.Test.UpdateWithInput update and blocks until completion
	UpdateWithInput(workflow.Context, *v1.UpdateWithInputRequest, ...*UpdateWithInputUpdateOptions) error

	// UpdateWithInputAsync executes a(n) test.option.v1.Test.UpdateWithInput update and returns a handle to the underlying activity
	UpdateWithInputAsync(workflow.Context, *v1.UpdateWithInputRequest, ...*UpdateWithInputUpdateOptions) (UpdateWithInputHandle, error)
}

// workflowWithInputRun provides a(n) WorkflowWithInputRun implementation
type workflowWithInputRun struct {
	cancel func()
	future workflow.Future
	id     string
}

// Cancel the underlying workflow execution
func (r *workflowWithInputRun) Cancel(ctx workflow.Context) error {
	if r.cancel != nil {
		r.cancel()
		if err := r.Get(ctx); err != nil && !errors.Is(err, workflow.ErrCanceled) {
			return err
		}
		return nil
	}
	return CancelTestWorkflow(ctx, r.id, "")
}

// Future returns the underlying activity future
func (r *workflowWithInputRun) Future() workflow.Future {
	return r.future
}

// Get blocks on activity completion and returns the underlying workflow result
func (r *workflowWithInputRun) Get(ctx workflow.Context) error {
	if err := r.future.Get(ctx, nil); err != nil {
		return err
	}
	return nil
}

// ID returns the underlying workflow id
func (r *workflowWithInputRun) ID() string {
	return r.id
}

// UpdateWithInput executes a(n) test.option.v1.Test.UpdateWithInput update and blocks until completion
func (r *workflowWithInputRun) UpdateWithInput(ctx workflow.Context, req *v1.UpdateWithInputRequest, opts ...*UpdateWithInputUpdateOptions) error {
	return UpdateWithInput(ctx, r.ID(), "", req, opts...)
}

// UpdateWithInputAsync executes a(n) test.option.v1.Test.UpdateWithInput update and returns a handle to the underlying activity
func (r *workflowWithInputRun) UpdateWithInputAsync(ctx workflow.Context, req *v1.UpdateWithInputRequest, opts ...*UpdateWithInputUpdateOptions) (UpdateWithInputHandle, error) {
	return UpdateWithInputAsync(ctx, r.ID(), "", req, opts...)
}

// WorkflowWithInput executes a(n) test.option.v1.Test.WorkflowWithInput workflow and blocks until error or response is received
func WorkflowWithInput(ctx workflow.Context, req *v1.WorkflowWithInputRequest, opts ...*WorkflowWithInputWorkflowOptions) error {
	run, err := WorkflowWithInputAsync(ctx, req, opts...)
	if err != nil {
		return err
	}
	return run.Get(ctx)
}

// WorkflowWithInputAsync executes a(n) test.option.v1.Test.WorkflowWithInput workflow and returns a handle to the underlying activity
func WorkflowWithInputAsync(ctx workflow.Context, req *v1.WorkflowWithInputRequest, opts ...*WorkflowWithInputWorkflowOptions) (WorkflowWithInputRun, error) {
	activityName := testOptions.filterActivity(v1.WorkflowWithInputWorkflowName)
	if activityName == "" {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("no activity registered for %s", v1.WorkflowWithInputWorkflowName),
			"Unimplemented",
			nil,
		)
	}

	opt := &WorkflowWithInputWorkflowOptions{}
	if len(opts) > 0 && opts[0] != nil {
		opt = opts[0]
	}
	if opt.HeartbeatInterval == 0 {
		opt.HeartbeatInterval = time.Second * 30
	}

	// configure activity options
	ao := workflow.GetActivityOptions(ctx)
	if opt.ActivityOptions != nil {
		ao = *opt.ActivityOptions
	}
	if ao.HeartbeatTimeout == 0 {
		ao.HeartbeatTimeout = opt.HeartbeatInterval * 2
	}
	// WaitForCancellation must be set otherwise the underlying workflow is not guaranteed to be canceled
	ao.WaitForCancellation = true

	if ao.StartToCloseTimeout == 0 && ao.ScheduleToCloseTimeout == 0 {
		ao.ScheduleToCloseTimeout = 600000000000 // 10 minutes
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// configure start workflow options
	wo := client.StartWorkflowOptions{}
	if opt.StartWorkflowOptions != nil {
		wo = *opt.StartWorkflowOptions
	}
	if wo.ID == "" {
		if err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
			id, err := expression.EvalExpression(v1.WorkflowWithInputIdexpression, req.ProtoReflect())
			if err != nil {
				workflow.GetLogger(ctx).Error("error evaluating id expression for \"test.option.v1.Test.WorkflowWithInput\" workflow", "error", err)
				return nil
			}
			return id
		}).Get(&wo.ID); err != nil {
			return nil, err
		}
	}
	if wo.ID == "" {
		if err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
			id, err := uuid.NewRandom()
			if err != nil {
				workflow.GetLogger(ctx).Error("error generating workflow id", "error", err)
				return nil
			}
			return id
		}).Get(&wo.ID); err != nil {
			return nil, err
		}
	}
	if wo.ID == "" {
		return nil, temporal.NewNonRetryableApplicationError("workflow id is required", "InvalidArgument", nil)
	}

	// marshal start workflow options protobuf message
	swo, err := xns.MarshalStartWorkflowOptions(wo)
	if err != nil {
		return nil, fmt.Errorf("error marshalling start workflow options: %w", err)
	}

	// marshal workflow request protobuf message
	wreq, err := anypb.New(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling workflow request: %w", err)
	}

	parentClosePolicy := temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL
	switch opt.ParentClosePolicy {
	case enumsv1.PARENT_CLOSE_POLICY_ABANDON:
		parentClosePolicy = temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_ABANDON
	case enumsv1.PARENT_CLOSE_POLICY_REQUEST_CANCEL:
		parentClosePolicy = temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL
	case enumsv1.PARENT_CLOSE_POLICY_TERMINATE:
		parentClosePolicy = temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE
	}

	ctx, cancel := workflow.WithCancel(ctx)
	return &workflowWithInputRun{
		cancel: cancel,
		id:     wo.ID,
		future: workflow.ExecuteActivity(ctx, activityName, &xnsv1.WorkflowRequest{
			Detached:             opt.Detached,
			HeartbeatInterval:    durationpb.New(opt.HeartbeatInterval),
			ParentClosePolicy:    parentClosePolicy,
			Request:              wreq,
			StartWorkflowOptions: swo,
		}),
	}, nil
}

// UpdateWithInputUpdateOptions are used to configure a(n) test.option.v1.Test.UpdateWithInput update execution
type UpdateWithInputUpdateOptions struct {
	ActivityOptions       *workflow.ActivityOptions
	HeartbeatInterval     time.Duration
	UpdateWorkflowOptions *client.UpdateWorkflowWithOptionsRequest
}

// NewUpdateWithInputUpdateOptions initializes a new UpdateWithInputUpdateOptions value
func NewUpdateWithInputUpdateOptions() *UpdateWithInputUpdateOptions {
	return &UpdateWithInputUpdateOptions{}
}

// WithActivityOptions can be used to customize the activity options
func (opts *UpdateWithInputUpdateOptions) WithActivityOptions(ao workflow.ActivityOptions) *UpdateWithInputUpdateOptions {
	opts.ActivityOptions = &ao
	return opts
}

// WithHeartbeatInterval can be used to customize the activity heartbeat interval
func (opts *UpdateWithInputUpdateOptions) WithHeartbeatInterval(d time.Duration) *UpdateWithInputUpdateOptions {
	opts.HeartbeatInterval = d
	return opts
}

// WithUpdateWorkflowOptions can be used to customize the update workflow options
func (opts *UpdateWithInputUpdateOptions) WithUpdateWorkflowOptions(uwo client.UpdateWorkflowWithOptionsRequest) *UpdateWithInputUpdateOptions {
	opts.UpdateWorkflowOptions = &uwo
	return opts
}

// UpdateWithInputHandle provides a handle to a test.option.v1.Test.UpdateWithInput workflow update
type UpdateWithInputHandle interface {
	// Cancel cancels the update activity
	Cancel(workflow.Context) error

	// Future returns the inner workflow.Future
	Future() workflow.Future

	// Get blocks on update completion and returns the result
	Get(workflow.Context) error

	// ID returns the update id
	ID() string
}

// updateWithInputHandle provides a(n) UpdateWithInputHandle implementation
type updateWithInputHandle struct {
	cancel func()
	future workflow.Future
	id     string
}

// Cancel the underlying workflow update
func (r *updateWithInputHandle) Cancel(ctx workflow.Context) error {
	r.cancel()
	if err := r.Get(ctx); err != nil && !errors.Is(err, workflow.ErrCanceled) {
		return err
	}
	return nil
}

// Future returns the underlying activity future
func (r *updateWithInputHandle) Future() workflow.Future {
	return r.future
}

// Get blocks on activity completion and returns the underlying update result
func (r *updateWithInputHandle) Get(ctx workflow.Context) error {
	if err := r.future.Get(ctx, nil); err != nil {
		return err
	}
	return nil
}

// ID returns the underlying workflow id
func (r *updateWithInputHandle) ID() string {
	return r.id
}

// UpdateWithInput executes a(n) test.option.v1.Test.UpdateWithInput update and blocks until error or response received
func UpdateWithInput(ctx workflow.Context, workflowID string, runID string, req *v1.UpdateWithInputRequest, opts ...*UpdateWithInputUpdateOptions) error {
	run, err := UpdateWithInputAsync(ctx, workflowID, runID, req, opts...)
	if err != nil {
		return err
	}
	return run.Get(ctx)
}

// UpdateWithInputAsync executes a(n) test.option.v1.Test.UpdateWithInput update and blocks until error or response received
func UpdateWithInputAsync(ctx workflow.Context, workflowID string, runID string, req *v1.UpdateWithInputRequest, opts ...*UpdateWithInputUpdateOptions) (UpdateWithInputHandle, error) {
	activityName := testOptions.filterActivity(v1.UpdateWithInputUpdateName)
	if activityName == "" {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("no activity registered for %s", v1.UpdateWithInputUpdateName),
			"Unimplemented",
			nil,
		)
	}

	opt := &UpdateWithInputUpdateOptions{}
	if len(opts) > 0 && opts[0] != nil {
		opt = opts[0]
	}

	if opt.HeartbeatInterval == 0 {
		opt.HeartbeatInterval = time.Second * 30
	}

	// configure activity options
	ao := workflow.GetActivityOptions(ctx)
	if opt.ActivityOptions != nil {
		ao = *opt.ActivityOptions
	}
	if ao.HeartbeatTimeout == 0 {
		ao.HeartbeatTimeout = opt.HeartbeatInterval * 2
	}
	// WaitForCancellation must be set otherwise the underlying workflow is not guaranteed to be canceled
	ao.WaitForCancellation = true

	if ao.StartToCloseTimeout == 0 && ao.ScheduleToCloseTimeout == 0 {
		ao.ScheduleToCloseTimeout = 300000000000 // 5 minutes
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	uo := client.UpdateWorkflowWithOptionsRequest{}
	if opt.UpdateWorkflowOptions != nil {
		uo = *opt.UpdateWorkflowOptions
	}
	uo.WorkflowID = workflowID
	uo.RunID = runID
	if uo.UpdateID == "" {
		if err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
			id, err := expression.EvalExpression(v1.UpdateWithInputIdexpression, req.ProtoReflect())
			if err != nil {
				workflow.GetLogger(ctx).Error("error evaluating id expression for \"test.option.v1.Test.UpdateWithInput\" update", "error", err)
				return nil
			}
			return id
		}).Get(&uo.UpdateID); err != nil {
			return nil, err
		}
	}
	if uo.UpdateID == "" {
		if err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
			id, err := uuid.NewRandom()
			if err != nil {
				workflow.GetLogger(ctx).Error("error generating update id", "error", err)
				return nil
			}
			return id
		}).Get(&uo.UpdateID); err != nil {
			return nil, err
		}
	}
	if uo.UpdateID == "" {
		return nil, temporal.NewNonRetryableApplicationError("update id is required", "InvalidArgument", nil)
	}

	uopb, err := xns.MarshalUpdateWorkflowOptions(uo)
	if err != nil {
		return nil, fmt.Errorf("error marshalling update workflow options: %w", err)
	}

	wreq, err := anypb.New(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling update request: %w", err)
	}

	ctx, cancel := workflow.WithCancel(ctx)
	return &updateWithInputHandle{
		cancel: cancel,
		id:     uo.UpdateID,
		future: workflow.ExecuteActivity(ctx, activityName, &xnsv1.UpdateRequest{
			HeartbeatInterval:     durationpb.New(opt.HeartbeatInterval),
			Request:               wreq,
			UpdateWorkflowOptions: uopb,
		}),
	}, nil
}

// CancelTestWorkflow cancels an existing workflow
func CancelTestWorkflow(ctx workflow.Context, workflowID string, runID string) error {
	return CancelTestWorkflowAsync(ctx, workflowID, runID).Get(ctx, nil)
}

// CancelTestWorkflowAsync cancels an existing workflow
func CancelTestWorkflowAsync(ctx workflow.Context, workflowID string, runID string) workflow.Future {
	activityName := testOptions.filterActivity("test.option.v1.Test.CancelWorkflow")
	if activityName == "" {
		f, s := workflow.NewFuture(ctx)
		s.SetError(temporal.NewNonRetryableApplicationError(
			"no activity registered for test.option.v1.Test.CancelWorkflow",
			"Unimplemented",
			nil,
		))
		return f
	}
	ao := workflow.GetActivityOptions(ctx)
	if ao.StartToCloseTimeout == 0 && ao.ScheduleToCloseTimeout == 0 {
		ao.StartToCloseTimeout = time.Minute
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	return workflow.ExecuteActivity(ctx, activityName, workflowID, runID)
}

// testActivities provides activities that can be used to interact with a(n) Test service's workflow, queries, signals, and updates across namespaces
type testActivities struct {
	client v1.TestClient
}

// CancelWorkflow cancels an existing workflow execution
func (a *testActivities) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	return a.client.CancelWorkflow(ctx, workflowID, runID)
}

// WorkflowWithInput executes a(n) test.option.v1.Test.WorkflowWithInput workflow via an activity
func (a *testActivities) WorkflowWithInput(ctx context.Context, input *xnsv1.WorkflowRequest) (err error) {
	// unmarshal workflow request
	var req v1.WorkflowWithInputRequest
	if err := input.Request.UnmarshalTo(&req); err != nil {
		return testOptions.convertError(temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("error unmarshalling workflow request of type %s as github.com/cludden/protoc-gen-go-temporal/gen/test/option/v1.WorkflowWithInputRequest", input.Request.GetTypeUrl()),
			"InvalidArgument",
			err,
		))
	}

	// initialize workflow execution
	var run v1.WorkflowWithInputRun
	run, err = a.client.WorkflowWithInputAsync(ctx, &req, v1.NewWorkflowWithInputOptions().WithStartWorkflowOptions(
		xns.UnmarshalStartWorkflowOptions(input.GetStartWorkflowOptions()),
	))
	if err != nil {
		return testOptions.convertError(err)
	}

	// exit early if detached enabled
	if input.GetDetached() {
		return nil
	}

	// otherwise, wait for execution to complete in child goroutine
	doneCh := make(chan struct{})
	go func() {
		err = run.Get(ctx)
		close(doneCh)
	}()

	heartbeatInterval := input.GetHeartbeatInterval().AsDuration()
	if heartbeatInterval == 0 {
		heartbeatInterval = time.Second * 30
	}

	// heartbeat activity while waiting for workflow execution to complete
	for {
		select {
		// send heartbeats periodically
		case <-time.After(heartbeatInterval):
			activity.RecordHeartbeat(ctx, run.ID())

		// return retryable error on worker close
		case <-activity.GetWorkerStopChannel(ctx):
			return temporal.NewApplicationError("worker is stopping", "WorkerStopped")

		// catch parent activity context cancellation. in most cases, this should indicate a
		// server-sent cancellation, but there's a non-zero possibility that this cancellation
		// is received due to the worker stopping, prior to detecting the closing of the worker
		// stop channel. to give us an opportunity to detect a cancellation stemming from the
		// worker closing, we again check to see if the worker stop channel is closed before
		// propagating the cancellation
		case <-ctx.Done():
			select {
			case <-activity.GetWorkerStopChannel(ctx):
				return temporal.NewApplicationError("worker is stopping", "WorkerStopped")
			default:
				parentClosePolicy := input.GetParentClosePolicy()
				if parentClosePolicy == temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL || parentClosePolicy == temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE {
					disconnectedCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
					defer cancel()
					if parentClosePolicy == temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL {
						err = run.Cancel(disconnectedCtx)
					} else {
						err = run.Terminate(disconnectedCtx, "xns activity cancellation received", "error", ctx.Err())
					}
					if err != nil {
						return testOptions.convertError(err)
					}
				}
				return testOptions.convertError(temporal.NewCanceledError(ctx.Err().Error()))
			}

		// handle workflow completion
		case <-doneCh:
			return testOptions.convertError(err)
		}
	}
}

// UpdateWithInput executes a(n) test.option.v1.Test.UpdateWithInput update via an activity
func (a *testActivities) UpdateWithInput(ctx context.Context, input *xnsv1.UpdateRequest) (err error) {
	var handle v1.UpdateWithInputHandle
	if activity.HasHeartbeatDetails(ctx) {
		// extract update id from heartbeat details
		var updateID string
		if err := activity.GetHeartbeatDetails(ctx, &updateID); err != nil {
			return testOptions.convertError(err)
		}

		// retrieve handle for existing update
		handle, err = a.client.GetUpdateWithInput(ctx, client.GetWorkflowUpdateHandleOptions{
			WorkflowID: input.GetUpdateWorkflowOptions().GetWorkflowId(),
			RunID:      input.GetUpdateWorkflowOptions().GetRunId(),
			UpdateID:   updateID,
		})
		if err != nil {
			return testOptions.convertError(err)
		}
	} else {
		// unmarshal update request
		var req v1.UpdateWithInputRequest
		if err := input.Request.UnmarshalTo(&req); err != nil {
			return testOptions.convertError(temporal.NewNonRetryableApplicationError(
				fmt.Sprintf("error unmarshalling update request of type %s as github.com/cludden/protoc-gen-go-temporal/gen/test/option/v1.UpdateWithInputRequest", input.Request.GetTypeUrl()),
				"InvalidArgument",
				err,
			))
		}

		uo := xns.UnmarshalUpdateWorkflowOptions(input.GetUpdateWorkflowOptions())
		uo.WaitPolicy = &updatev1.WaitPolicy{
			LifecycleStage: enumsv1.UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ACCEPTED,
		}

		// initialize update execution
		handle, err = a.client.UpdateWithInputAsync(
			ctx,
			input.GetUpdateWorkflowOptions().GetWorkflowId(),
			input.GetUpdateWorkflowOptions().GetRunId(),
			&req,
			v1.NewUpdateWithInputOptions().WithUpdateWorkflowOptions(uo),
		)
		if err != nil {
			return testOptions.convertError(err)
		}
		activity.RecordHeartbeat(ctx, handle.UpdateID())
	}

	// wait for update to complete in child goroutine
	doneCh := make(chan struct{})
	go func() {
		err = handle.Get(ctx)
		close(doneCh)
	}()

	heartbeatInterval := input.GetHeartbeatInterval().AsDuration()
	if heartbeatInterval == 0 {
		heartbeatInterval = time.Minute
	}

	// heartbeat activity while waiting for workflow update to complete
	for {
		select {
		case <-time.After(heartbeatInterval):
			activity.RecordHeartbeat(ctx, handle.UpdateID())
		case <-ctx.Done():
			return testOptions.convertError(ctx.Err())
		case <-doneCh:
			return testOptions.convertError(err)
		}
	}
}
