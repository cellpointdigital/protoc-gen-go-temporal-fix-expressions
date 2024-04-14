// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 1.10.4 (38d49e8722013d965532492a3d6b9318a9a33971)
//	go go1.21.9
//	protoc (unknown)
//
// source: example/searchattributes/v1/searchattributes.proto
package searchattributesv1xns

import (
	"context"
	"errors"
	"fmt"
	v1 "github.com/cludden/protoc-gen-go-temporal/gen/example/searchattributes/v1"
	v11 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/xns/v1"
	expression "github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	xns "github.com/cludden/protoc-gen-go-temporal/pkg/xns"
	uuid "github.com/google/uuid"
	activity "go.temporal.io/sdk/activity"
	client "go.temporal.io/sdk/client"
	temporal "go.temporal.io/sdk/temporal"
	worker "go.temporal.io/sdk/worker"
	workflow "go.temporal.io/sdk/workflow"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	"time"
)

// ExampleOptions is used to configure example.searchattributes.v1.Example xns activity registration
type ExampleOptions struct {
	// errorConverter is used to customize error
	errorConverter func(error) error
	// filter is used to filter xns activity registrations. It receives as
	// input the original activity name, and should return one of the following:
	// 1. the original activity name, for no changes
	// 2. a modified activity name, to override the original activity name
	// 3. an empty string, to skip registration
	filter func(string) string
}

// NewExampleOptions initializes a new ExampleOptions value
func NewExampleOptions() *ExampleOptions {
	return &ExampleOptions{}
}

// WithErrorConverter overrides the default error converter applied to xns activity errors
func (opts *ExampleOptions) WithErrorConverter(errorConverter func(error) error) *ExampleOptions {
	opts.errorConverter = errorConverter
	return opts
}

// Filter is used to filter registered xns activities or customize their name
func (opts *ExampleOptions) WithFilter(filter func(string) string) *ExampleOptions {
	opts.filter = filter
	return opts
}

// convertError is applied to all xns activity errors
func (opts *ExampleOptions) convertError(err error) error {
	if err == nil {
		return nil
	}
	if opts != nil && opts.errorConverter != nil {
		return opts.errorConverter(err)
	}
	return xns.ErrorToApplicationError(err)
}

// filterActivity is used to filter xns activity registrations
func (opts *ExampleOptions) filterActivity(name string) string {
	if opts == nil || opts.filter == nil {
		return name
	}
	return opts.filter(name)
}

// exampleOptions is a reference to the ExampleOptions initialized at registration
var exampleOptions *ExampleOptions

// RegisterExampleActivities registers example.searchattributes.v1.Example cross-namespace activities
func RegisterExampleActivities(r worker.ActivityRegistry, c v1.ExampleClient, options ...*ExampleOptions) {
	if exampleOptions == nil && len(options) > 0 && options[0] != nil {
		exampleOptions = options[0]
	}
	a := &exampleActivities{c}
	if name := exampleOptions.filterActivity("example.searchattributes.v1.Example.CancelWorkflow"); name != "" {
		r.RegisterActivityWithOptions(a.CancelWorkflow, activity.RegisterOptions{Name: name})
	}
	if name := exampleOptions.filterActivity(v1.SearchAttributesWorkflowName); name != "" {
		r.RegisterActivityWithOptions(a.SearchAttributes, activity.RegisterOptions{Name: name})
	}
}

// SearchAttributesWorkflowOptions are used to configure a(n) example.searchattributes.v1.Example.SearchAttributes workflow execution
type SearchAttributesWorkflowOptions struct {
	ActivityOptions      *workflow.ActivityOptions
	Detached             bool
	HeartbeatInterval    time.Duration
	StartWorkflowOptions *client.StartWorkflowOptions
}

// NewSearchAttributesWorkflowOptions initializes a new SearchAttributesWorkflowOptions value
func NewSearchAttributesWorkflowOptions() *SearchAttributesWorkflowOptions {
	return &SearchAttributesWorkflowOptions{}
}

// WithActivityOptions can be used to customize the activity options
func (opts *SearchAttributesWorkflowOptions) WithActivityOptions(ao workflow.ActivityOptions) *SearchAttributesWorkflowOptions {
	opts.ActivityOptions = &ao
	return opts
}

// WithDetached can be used to start a workflow execution and exit immediately
func (opts *SearchAttributesWorkflowOptions) WithDetached(d bool) *SearchAttributesWorkflowOptions {
	opts.Detached = d
	return opts
}

// WithHeartbeatInterval can be used to customize the activity heartbeat interval
func (opts *SearchAttributesWorkflowOptions) WithHeartbeatInterval(d time.Duration) *SearchAttributesWorkflowOptions {
	opts.HeartbeatInterval = d
	return opts
}

// WithStartWorkflowOptions can be used to customize the start workflow options
func (opts *SearchAttributesWorkflowOptions) WithStartWorkflow(swo client.StartWorkflowOptions) *SearchAttributesWorkflowOptions {
	opts.StartWorkflowOptions = &swo
	return opts
}

// SearchAttributesRun provides a handle to a example.searchattributes.v1.Example.SearchAttributes workflow execution
type SearchAttributesRun interface {
	// Cancel cancels the workflow
	Cancel(workflow.Context) error

	// Future returns the inner workflow.Future
	Future() workflow.Future

	// Get returns the inner workflow.Future
	Get(workflow.Context) error

	// ID returns the workflow id
	ID() string
}

// searchAttributesRun provides a(n) SearchAttributesRun implementation
type searchAttributesRun struct {
	cancel func()
	future workflow.Future
	id     string
}

// Cancel the underlying workflow execution
func (r *searchAttributesRun) Cancel(ctx workflow.Context) error {
	if r.cancel != nil {
		r.cancel()
		if err := r.Get(ctx); err != nil && !errors.Is(err, workflow.ErrCanceled) {
			return err
		}
		return nil
	}
	return CancelExampleWorkflow(ctx, r.id, "")
}

// Future returns the underlying activity future
func (r *searchAttributesRun) Future() workflow.Future {
	return r.future
}

// Get blocks on activity completion and returns the underlying workflow result
func (r *searchAttributesRun) Get(ctx workflow.Context) error {
	if err := r.future.Get(ctx, nil); err != nil {
		return err
	}
	return nil
}

// ID returns the underlying workflow id
func (r *searchAttributesRun) ID() string {
	return r.id
}

// SearchAttributes executes a(n) example.searchattributes.v1.Example.SearchAttributes workflow and blocks until error or response received
func SearchAttributes(ctx workflow.Context, req *v1.SearchAttributesInput, opts ...*SearchAttributesWorkflowOptions) error {
	run, err := SearchAttributesAsync(ctx, req, opts...)
	if err != nil {
		return err
	}
	return run.Get(ctx)
}

// SearchAttributesAsync executes a(n) example.searchattributes.v1.Example.SearchAttributes workflow and blocks until error or response received
func SearchAttributesAsync(ctx workflow.Context, req *v1.SearchAttributesInput, opts ...*SearchAttributesWorkflowOptions) (SearchAttributesRun, error) {
	activityName := exampleOptions.filterActivity(v1.SearchAttributesWorkflowName)
	if activityName == "" {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("no activity registered for %s", v1.SearchAttributesWorkflowName),
			"Unimplemented",
			nil,
		)
	}

	opt := &SearchAttributesWorkflowOptions{}
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
	if ao.StartToCloseTimeout == 0 && ao.ScheduleToCloseTimeout == 0 {
		ao.ScheduleToCloseTimeout = 86400000000000 // 1 day
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// configure start workflow options
	wo := client.StartWorkflowOptions{}
	if opt.StartWorkflowOptions != nil {
		wo = *opt.StartWorkflowOptions
	}
	if wo.ID == "" {
		if err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
			id, err := expression.EvalExpression(v1.SearchAttributesIdexpression, req.ProtoReflect())
			if err != nil {
				workflow.GetLogger(ctx).Error("error evaluating id expression for \"example.searchattributes.v1.Example.SearchAttributes\" workflow", "error", err)
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

	ctx, cancel := workflow.WithCancel(ctx)
	return &searchAttributesRun{
		cancel: cancel,
		id:     wo.ID,
		future: workflow.ExecuteActivity(ctx, activityName, &v11.WorkflowRequest{
			Detached:             opt.Detached,
			HeartbeatInterval:    durationpb.New(opt.HeartbeatInterval),
			Request:              wreq,
			StartWorkflowOptions: swo,
		}),
	}, nil
}

// CancelExampleWorkflow cancels an existing workflow
func CancelExampleWorkflow(ctx workflow.Context, workflowID string, runID string) error {
	return CancelExampleWorkflowAsync(ctx, workflowID, runID).Get(ctx, nil)
}

// CancelExampleWorkflowAsync cancels an existing workflow
func CancelExampleWorkflowAsync(ctx workflow.Context, workflowID string, runID string) workflow.Future {
	activityName := exampleOptions.filterActivity("example.searchattributes.v1.Example.CancelWorkflow")
	if activityName == "" {
		f, s := workflow.NewFuture(ctx)
		s.SetError(temporal.NewNonRetryableApplicationError(
			"no activity registered for example.searchattributes.v1.Example.CancelWorkflow",
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

// exampleActivities provides activities that can be used to interact with a(n) Example service's workflow, queries, signals, and updates across namespaces
type exampleActivities struct {
	client v1.ExampleClient
}

// CancelWorkflow cancels an existing workflow execution
func (a *exampleActivities) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	return a.client.CancelWorkflow(ctx, workflowID, runID)
}

// SearchAttributes executes a(n) example.searchattributes.v1.Example.SearchAttributes workflow via an activity
func (a *exampleActivities) SearchAttributes(ctx context.Context, input *v11.WorkflowRequest) (err error) {
	// unmarshal workflow request
	var req v1.SearchAttributesInput
	if err := input.Request.UnmarshalTo(&req); err != nil {
		return exampleOptions.convertError(temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("error unmarshalling workflow request of type %s as github.com/cludden/protoc-gen-go-temporal/gen/example/searchattributes/v1.SearchAttributesInput", input.Request.GetTypeUrl()),
			"InvalidArgument",
			err,
		))
	}

	// initialize workflow execution
	var run v1.SearchAttributesRun
	run, err = a.client.SearchAttributesAsync(ctx, &req, v1.NewSearchAttributesOptions().WithStartWorkflowOptions(
		xns.UnmarshalStartWorkflowOptions(input.GetStartWorkflowOptions()),
	))
	if err != nil {
		return exampleOptions.convertError(err)
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
		heartbeatInterval = time.Minute
	}

	// heartbeat activity while waiting for workflow execution to complete
	for {
		select {
		case <-time.After(heartbeatInterval):
			activity.RecordHeartbeat(ctx, run.ID())
		case <-ctx.Done():
			if err := run.Cancel(ctx); err != nil {
				return exampleOptions.convertError(err)
			}
			return exampleOptions.convertError(workflow.ErrCanceled)
		case <-doneCh:
			return exampleOptions.convertError(err)
		}
	}
}
