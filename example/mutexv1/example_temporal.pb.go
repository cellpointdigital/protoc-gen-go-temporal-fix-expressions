// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 0.7.6-next (aff23bc1dabb6ac2b3a06abfc734d7a9d33ff13c)
//	go go1.20.4
//	protoc (unknown)
//
// source: example.proto
package mutexv1

import (
	"context"
	"errors"
	"fmt"
	expression "github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	v1 "go.temporal.io/api/enums/v1"
	activity "go.temporal.io/sdk/activity"
	client "go.temporal.io/sdk/client"
	worker "go.temporal.io/sdk/worker"
	workflow "go.temporal.io/sdk/workflow"
)

// MutexTaskQueue is the default task-queue for a Mutex worker
const MutexTaskQueue = "mutex-v1"

// Mutex workflow names
const (
	MutexWorkflowName                   = "mycompany.mutex.v1.Mutex.MutexWorkflow"
	SampleWorkflowWithMutexWorkflowName = "mycompany.mutex.v1.Mutex.SampleWorkflowWithMutexWorkflow"
)

// Mutex id expressions
var (
	MutexIDExpression                   = expression.MustParseExpression("mutex/${!resource}")
	SampleWorkflowWithMutexIDExpression = expression.MustParseExpression("sample-workflow-with-mutex/${!resource}/${!uuid_v4()}")
)

// Mutex signal names
const (
	AcquireLeaseSignalName  = "mycompany.mutex.v1.Mutex.AcquireLeaseSignal"
	LeaseAcquiredSignalName = "mycompany.mutex.v1.Mutex.LeaseAcquiredSignal"
	RenewLeaseSignalName    = "mycompany.mutex.v1.Mutex.RenewLeaseSignal"
	RevokeLeaseSignalName   = "mycompany.mutex.v1.Mutex.RevokeLeaseSignal"
)

// Mutex activity names
const (
	MutexActivityName = "mycompany.mutex.v1.Mutex.MutexActivity"
)

// Client describes a client for a Mutex worker
type Client interface {
	// Mutex provides a mutex over a shared resource
	Mutex(ctx context.Context, opts *client.StartWorkflowOptions, req *MutexRequest) error
	// ExecuteMutex executes a Mutex workflow
	ExecuteMutex(ctx context.Context, opts *client.StartWorkflowOptions, req *MutexRequest) (MutexRun, error)
	// GetMutex retrieves a Mutex workflow execution
	GetMutex(ctx context.Context, workflowID string, runID string) (MutexRun, error)
	// StartMutexWithAcquireLease sends a AcquireLease signal to a Mutex workflow, starting it if not present
	StartMutexWithAcquireLease(ctx context.Context, opts *client.StartWorkflowOptions, req *MutexRequest, signal *AcquireLeaseRequest) (MutexRun, error)
	// SampleWorkflowWithMutex provides an example of a running workflow that uses
	// a Mutex workflow to prevent concurrent access to a shared resource
	SampleWorkflowWithMutex(ctx context.Context, opts *client.StartWorkflowOptions, req *SampleWorkflowWithMutexRequest) (*SampleWorkflowWithMutexResponse, error)
	// ExecuteSampleWorkflowWithMutex executes a SampleWorkflowWithMutex workflow
	ExecuteSampleWorkflowWithMutex(ctx context.Context, opts *client.StartWorkflowOptions, req *SampleWorkflowWithMutexRequest) (SampleWorkflowWithMutexRun, error)
	// GetSampleWorkflowWithMutex retrieves a SampleWorkflowWithMutex workflow execution
	GetSampleWorkflowWithMutex(ctx context.Context, workflowID string, runID string) (SampleWorkflowWithMutexRun, error)
	// SignalAcquireLease sends a AcquireLease signal to an existing workflow
	SignalAcquireLease(ctx context.Context, workflowID string, runID string, signal *AcquireLeaseRequest) error
	// SignalLeaseAcquired sends a LeaseAcquired signal to an existing workflow
	SignalLeaseAcquired(ctx context.Context, workflowID string, runID string, signal *LeaseAcquiredRequest) error
	// SignalRenewLease sends a RenewLease signal to an existing workflow
	SignalRenewLease(ctx context.Context, workflowID string, runID string, signal *RenewLeaseRequest) error
	// SignalRevokeLease sends a RevokeLease signal to an existing workflow
	SignalRevokeLease(ctx context.Context, workflowID string, runID string, signal *RevokeLeaseRequest) error
}

// Compile-time check that workflowClient satisfies Client
var _ Client = &workflowClient{}

// workflowClient implements a temporal client for a Mutex service
type workflowClient struct {
	client client.Client
}

// NewClient initializes a new Mutex client
func NewClient(c client.Client) Client {
	return &workflowClient{client: c}
}

// NewClientWithOptions initializes a new Mutex client with the given options
func NewClientWithOptions(c client.Client, opts client.Options) (Client, error) {
	var err error
	c, err = client.NewClientFromExisting(c, opts)
	if err != nil {
		return nil, fmt.Errorf("error initializing client with options: %w", err)
	}
	return &workflowClient{client: c}, nil
}

// Mutex provides a mutex over a shared resource
func (c *workflowClient) Mutex(ctx context.Context, opts *client.StartWorkflowOptions, req *MutexRequest) error {
	run, err := c.ExecuteMutex(ctx, opts, req)
	if err != nil {
		return err
	}
	return run.Get(ctx)
}

// ExecuteMutex starts a Mutex workflow
func (c *workflowClient) ExecuteMutex(ctx context.Context, opts *client.StartWorkflowOptions, req *MutexRequest) (MutexRun, error) {
	if opts == nil {
		opts = &client.StartWorkflowOptions{}
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "mutex-v1"
	}
	if opts.ID == "" {
		id, err := expression.EvalExpression(MutexIDExpression, req.ProtoReflect())
		if err != nil {
			return nil, err
		}
		opts.ID = id
	}
	if opts.WorkflowIDReusePolicy == v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v1.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE
	}
	if opts.WorkflowExecutionTimeout == 0 {
		opts.WorkflowRunTimeout = 3600000000000 // 1h0m0s
	}
	run, err := c.client.ExecuteWorkflow(ctx, *opts, MutexWorkflowName, req)
	if err != nil {
		return nil, err
	}
	if run == nil {
		return nil, errors.New("execute workflow returned nil run")
	}
	return &mutexRun{
		client: c,
		run:    run,
	}, nil
}

// GetMutex fetches an existing Mutex execution
func (c *workflowClient) GetMutex(ctx context.Context, workflowID string, runID string) (MutexRun, error) {
	return &mutexRun{
		client: c,
		run:    c.client.GetWorkflow(ctx, workflowID, runID),
	}, nil
}

// StartMutexWithAcquireLease starts a Mutex workflow and sends a AcquireLease signal in a transaction
func (c *workflowClient) StartMutexWithAcquireLease(ctx context.Context, opts *client.StartWorkflowOptions, req *MutexRequest, signal *AcquireLeaseRequest) (MutexRun, error) {
	if opts == nil {
		opts = &client.StartWorkflowOptions{}
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "mutex-v1"
	}
	if opts.ID == "" {
		id, err := expression.EvalExpression(MutexIDExpression, req.ProtoReflect())
		if err != nil {
			return nil, err
		}
		opts.ID = id
	}
	if opts.WorkflowIDReusePolicy == v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v1.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE
	}
	if opts.WorkflowExecutionTimeout == 0 {
		opts.WorkflowRunTimeout = 3600000000000 // 1h0m0s
	}
	run, err := c.client.SignalWithStartWorkflow(ctx, opts.ID, AcquireLeaseSignalName, signal, *opts, MutexWorkflowName, req)
	if run == nil || err != nil {
		return nil, err
	}
	return &mutexRun{
		client: c,
		run:    run,
	}, nil
}

// SampleWorkflowWithMutex provides an example of a running workflow that uses
// a Mutex workflow to prevent concurrent access to a shared resource
func (c *workflowClient) SampleWorkflowWithMutex(ctx context.Context, opts *client.StartWorkflowOptions, req *SampleWorkflowWithMutexRequest) (*SampleWorkflowWithMutexResponse, error) {
	run, err := c.ExecuteSampleWorkflowWithMutex(ctx, opts, req)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// ExecuteSampleWorkflowWithMutex starts a SampleWorkflowWithMutex workflow
func (c *workflowClient) ExecuteSampleWorkflowWithMutex(ctx context.Context, opts *client.StartWorkflowOptions, req *SampleWorkflowWithMutexRequest) (SampleWorkflowWithMutexRun, error) {
	if opts == nil {
		opts = &client.StartWorkflowOptions{}
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "mutex-v1"
	}
	if opts.ID == "" {
		id, err := expression.EvalExpression(SampleWorkflowWithMutexIDExpression, req.ProtoReflect())
		if err != nil {
			return nil, err
		}
		opts.ID = id
	}
	if opts.WorkflowIDReusePolicy == v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v1.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY
	}
	if opts.WorkflowExecutionTimeout == 0 {
		opts.WorkflowRunTimeout = 3600000000000 // 1h0m0s
	}
	run, err := c.client.ExecuteWorkflow(ctx, *opts, SampleWorkflowWithMutexWorkflowName, req)
	if err != nil {
		return nil, err
	}
	if run == nil {
		return nil, errors.New("execute workflow returned nil run")
	}
	return &sampleWorkflowWithMutexRun{
		client: c,
		run:    run,
	}, nil
}

// GetSampleWorkflowWithMutex fetches an existing SampleWorkflowWithMutex execution
func (c *workflowClient) GetSampleWorkflowWithMutex(ctx context.Context, workflowID string, runID string) (SampleWorkflowWithMutexRun, error) {
	return &sampleWorkflowWithMutexRun{
		client: c,
		run:    c.client.GetWorkflow(ctx, workflowID, runID),
	}, nil
}

// SignalAcquireLease sends a AcquireLease signal to an existing workflow
func (c *workflowClient) SignalAcquireLease(ctx context.Context, workflowID string, runID string, signal *AcquireLeaseRequest) error {
	return c.client.SignalWorkflow(ctx, workflowID, runID, AcquireLeaseSignalName, signal)
}

// SignalLeaseAcquired sends a LeaseAcquired signal to an existing workflow
func (c *workflowClient) SignalLeaseAcquired(ctx context.Context, workflowID string, runID string, signal *LeaseAcquiredRequest) error {
	return c.client.SignalWorkflow(ctx, workflowID, runID, LeaseAcquiredSignalName, signal)
}

// SignalRenewLease sends a RenewLease signal to an existing workflow
func (c *workflowClient) SignalRenewLease(ctx context.Context, workflowID string, runID string, signal *RenewLeaseRequest) error {
	return c.client.SignalWorkflow(ctx, workflowID, runID, RenewLeaseSignalName, signal)
}

// SignalRevokeLease sends a RevokeLease signal to an existing workflow
func (c *workflowClient) SignalRevokeLease(ctx context.Context, workflowID string, runID string, signal *RevokeLeaseRequest) error {
	return c.client.SignalWorkflow(ctx, workflowID, runID, RevokeLeaseSignalName, signal)
}

// MutexRun describes a Mutex workflow run
type MutexRun interface {
	// ID returns the workflow ID
	ID() string
	// RunID returns the workflow instance ID
	RunID() string
	// Get blocks until the workflow is complete and returns the result
	Get(ctx context.Context) error
	// AcquireLease sends a AcquireLease signal to the workflow
	AcquireLease(ctx context.Context, req *AcquireLeaseRequest) error
	// RenewLease sends a RenewLease signal to the workflow
	RenewLease(ctx context.Context, req *RenewLeaseRequest) error
	// RevokeLease sends a RevokeLease signal to the workflow
	RevokeLease(ctx context.Context, req *RevokeLeaseRequest) error
}

// mutexRun provides an internal implementation of a MutexRun
type mutexRun struct {
	client *workflowClient
	run    client.WorkflowRun
}

// ID returns the workflow ID
func (r *mutexRun) ID() string {
	return r.run.GetID()
}

// RunID returns the execution ID
func (r *mutexRun) RunID() string {
	return r.run.GetRunID()
}

// Get blocks until the workflow is complete, returning the result if applicable
func (r *mutexRun) Get(ctx context.Context) error {
	return r.run.Get(ctx, nil)
}

// AcquireLease sends a AcquireLease signal to the workflow
func (r *mutexRun) AcquireLease(ctx context.Context, req *AcquireLeaseRequest) error {
	return r.client.SignalAcquireLease(ctx, r.ID(), "", req)
}

// RenewLease sends a RenewLease signal to the workflow
func (r *mutexRun) RenewLease(ctx context.Context, req *RenewLeaseRequest) error {
	return r.client.SignalRenewLease(ctx, r.ID(), "", req)
}

// RevokeLease sends a RevokeLease signal to the workflow
func (r *mutexRun) RevokeLease(ctx context.Context, req *RevokeLeaseRequest) error {
	return r.client.SignalRevokeLease(ctx, r.ID(), "", req)
}

// SampleWorkflowWithMutexRun describes a SampleWorkflowWithMutex workflow run
type SampleWorkflowWithMutexRun interface {
	// ID returns the workflow ID
	ID() string
	// RunID returns the workflow instance ID
	RunID() string
	// Get blocks until the workflow is complete and returns the result
	Get(ctx context.Context) (*SampleWorkflowWithMutexResponse, error)
	// LeaseAcquired sends a LeaseAcquired signal to the workflow
	LeaseAcquired(ctx context.Context, req *LeaseAcquiredRequest) error
}

// sampleWorkflowWithMutexRun provides an internal implementation of a SampleWorkflowWithMutexRun
type sampleWorkflowWithMutexRun struct {
	client *workflowClient
	run    client.WorkflowRun
}

// ID returns the workflow ID
func (r *sampleWorkflowWithMutexRun) ID() string {
	return r.run.GetID()
}

// RunID returns the execution ID
func (r *sampleWorkflowWithMutexRun) RunID() string {
	return r.run.GetRunID()
}

// Get blocks until the workflow is complete, returning the result if applicable
func (r *sampleWorkflowWithMutexRun) Get(ctx context.Context) (*SampleWorkflowWithMutexResponse, error) {
	var resp SampleWorkflowWithMutexResponse
	if err := r.run.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// LeaseAcquired sends a LeaseAcquired signal to the workflow
func (r *sampleWorkflowWithMutexRun) LeaseAcquired(ctx context.Context, req *LeaseAcquiredRequest) error {
	return r.client.SignalLeaseAcquired(ctx, r.ID(), "", req)
}

// Workflows provides methods for initializing new Mutex workflow values
type Workflows interface {
	// Mutex initializes a new MutexWorkflow value
	Mutex(ctx workflow.Context, input *MutexInput) (MutexWorkflow, error)
	// SampleWorkflowWithMutex initializes a new SampleWorkflowWithMutexWorkflow value
	SampleWorkflowWithMutex(ctx workflow.Context, input *SampleWorkflowWithMutexInput) (SampleWorkflowWithMutexWorkflow, error)
}

// RegisterWorkflows registers Mutex workflows with the given worker
func RegisterWorkflows(r worker.Registry, workflows Workflows) {
	RegisterMutexWorkflow(r, workflows.Mutex)
	RegisterSampleWorkflowWithMutexWorkflow(r, workflows.SampleWorkflowWithMutex)
}

// RegisterMutexWorkflow registers a Mutex workflow with the given worker
func RegisterMutexWorkflow(r worker.Registry, wf func(workflow.Context, *MutexInput) (MutexWorkflow, error)) {
	r.RegisterWorkflowWithOptions(buildMutex(wf), workflow.RegisterOptions{Name: MutexWorkflowName})
}

// buildMutex converts a Mutex workflow struct into a valid workflow function
func buildMutex(wf func(workflow.Context, *MutexInput) (MutexWorkflow, error)) func(workflow.Context, *MutexRequest) error {
	return (&mutex{wf}).Mutex
}

// mutex provides an Mutex method for calling the user's implementation
type mutex struct {
	ctor func(workflow.Context, *MutexInput) (MutexWorkflow, error)
}

// Mutex constructs a new Mutex value and executes it
func (w *mutex) Mutex(ctx workflow.Context, req *MutexRequest) error {
	input := &MutexInput{
		Req: req,
		AcquireLease: &AcquireLeaseSignal{
			Channel: workflow.GetSignalChannel(ctx, AcquireLeaseSignalName),
		},
		RenewLease: &RenewLeaseSignal{
			Channel: workflow.GetSignalChannel(ctx, RenewLeaseSignalName),
		},
		RevokeLease: &RevokeLeaseSignal{
			Channel: workflow.GetSignalChannel(ctx, RevokeLeaseSignalName),
		},
	}
	wf, err := w.ctor(ctx, input)
	if err != nil {
		return err
	}
	return wf.Execute(ctx)
}

// MutexInput describes the input to a Mutex workflow constructor
type MutexInput struct {
	Req          *MutexRequest
	AcquireLease *AcquireLeaseSignal
	RenewLease   *RenewLeaseSignal
	RevokeLease  *RevokeLeaseSignal
}

// Mutex provides a mutex over a shared resource
type MutexWorkflow interface {
	// Execute a Mutex workflow
	Execute(ctx workflow.Context) error
}

// MutexChild executes a child Mutex workflow
func MutexChild(ctx workflow.Context, opts *workflow.ChildWorkflowOptions, req *MutexRequest) *MutexChildRun {
	if opts == nil {
		childOpts := workflow.GetChildWorkflowOptions(ctx)
		opts = &childOpts
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "mutex-v1"
	}
	if opts.WorkflowID == "" {
		id, err := expression.EvalExpression(MutexIDExpression, req.ProtoReflect())
		if err != nil {
			panic(err)
		}
		opts.WorkflowID = id
	}
	if opts.WorkflowIDReusePolicy == v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v1.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE
	}
	if opts.WorkflowExecutionTimeout == 0 {
		opts.WorkflowRunTimeout = 3600000000000 // 1h0m0s
	}
	ctx = workflow.WithChildOptions(ctx, *opts)
	return &MutexChildRun{Future: workflow.ExecuteChildWorkflow(ctx, MutexWorkflowName, req)}
}

// MutexChildRun describes a child Mutex workflow run
type MutexChildRun struct {
	Future workflow.ChildWorkflowFuture
}

// Get blocks until the workflow is completed, returning the response value
func (r *MutexChildRun) Get(ctx workflow.Context) error {
	if err := r.Future.Get(ctx, nil); err != nil {
		return err
	}
	return nil
}

// Select adds this completion to the selector. Callback can be nil.
func (r *MutexChildRun) Select(sel workflow.Selector, fn func(MutexChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future, func(workflow.Future) {
		if fn != nil {
			fn(*r)
		}
	})
}

// SelectStart adds waiting for start to the selector. Callback can be nil.
func (r *MutexChildRun) SelectStart(sel workflow.Selector, fn func(MutexChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future.GetChildWorkflowExecution(), func(workflow.Future) {
		if fn != nil {
			fn(*r)
		}
	})
}

// WaitStart waits for the child workflow to start
func (r *MutexChildRun) WaitStart(ctx workflow.Context) (*workflow.Execution, error) {
	var exec workflow.Execution
	if err := r.Future.GetChildWorkflowExecution().Get(ctx, &exec); err != nil {
		return nil, err
	}
	return &exec, nil
}

// AcquireLease sends the corresponding signal request to the child workflow
func (r *MutexChildRun) AcquireLease(ctx workflow.Context, input *AcquireLeaseRequest) workflow.Future {
	return r.Future.SignalChildWorkflow(ctx, AcquireLeaseSignalName, input)
}

// RenewLease sends the corresponding signal request to the child workflow
func (r *MutexChildRun) RenewLease(ctx workflow.Context, input *RenewLeaseRequest) workflow.Future {
	return r.Future.SignalChildWorkflow(ctx, RenewLeaseSignalName, input)
}

// RevokeLease sends the corresponding signal request to the child workflow
func (r *MutexChildRun) RevokeLease(ctx workflow.Context, input *RevokeLeaseRequest) workflow.Future {
	return r.Future.SignalChildWorkflow(ctx, RevokeLeaseSignalName, input)
}

// RegisterSampleWorkflowWithMutexWorkflow registers a SampleWorkflowWithMutex workflow with the given worker
func RegisterSampleWorkflowWithMutexWorkflow(r worker.Registry, wf func(workflow.Context, *SampleWorkflowWithMutexInput) (SampleWorkflowWithMutexWorkflow, error)) {
	r.RegisterWorkflowWithOptions(buildSampleWorkflowWithMutex(wf), workflow.RegisterOptions{Name: SampleWorkflowWithMutexWorkflowName})
}

// buildSampleWorkflowWithMutex converts a SampleWorkflowWithMutex workflow struct into a valid workflow function
func buildSampleWorkflowWithMutex(wf func(workflow.Context, *SampleWorkflowWithMutexInput) (SampleWorkflowWithMutexWorkflow, error)) func(workflow.Context, *SampleWorkflowWithMutexRequest) (*SampleWorkflowWithMutexResponse, error) {
	return (&sampleWorkflowWithMutex{wf}).SampleWorkflowWithMutex
}

// sampleWorkflowWithMutex provides an SampleWorkflowWithMutex method for calling the user's implementation
type sampleWorkflowWithMutex struct {
	ctor func(workflow.Context, *SampleWorkflowWithMutexInput) (SampleWorkflowWithMutexWorkflow, error)
}

// SampleWorkflowWithMutex constructs a new SampleWorkflowWithMutex value and executes it
func (w *sampleWorkflowWithMutex) SampleWorkflowWithMutex(ctx workflow.Context, req *SampleWorkflowWithMutexRequest) (*SampleWorkflowWithMutexResponse, error) {
	input := &SampleWorkflowWithMutexInput{
		Req: req,
		LeaseAcquired: &LeaseAcquiredSignal{
			Channel: workflow.GetSignalChannel(ctx, LeaseAcquiredSignalName),
		},
	}
	wf, err := w.ctor(ctx, input)
	if err != nil {
		return nil, err
	}
	return wf.Execute(ctx)
}

// SampleWorkflowWithMutexInput describes the input to a SampleWorkflowWithMutex workflow constructor
type SampleWorkflowWithMutexInput struct {
	Req           *SampleWorkflowWithMutexRequest
	LeaseAcquired *LeaseAcquiredSignal
}

// SampleWorkflowWithMutex provides an example of a running workflow that uses
// a Mutex workflow to prevent concurrent access to a shared resource
type SampleWorkflowWithMutexWorkflow interface {
	// Execute a SampleWorkflowWithMutex workflow
	Execute(ctx workflow.Context) (*SampleWorkflowWithMutexResponse, error)
}

// SampleWorkflowWithMutexChild executes a child SampleWorkflowWithMutex workflow
func SampleWorkflowWithMutexChild(ctx workflow.Context, opts *workflow.ChildWorkflowOptions, req *SampleWorkflowWithMutexRequest) *SampleWorkflowWithMutexChildRun {
	if opts == nil {
		childOpts := workflow.GetChildWorkflowOptions(ctx)
		opts = &childOpts
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "mutex-v1"
	}
	if opts.WorkflowID == "" {
		id, err := expression.EvalExpression(SampleWorkflowWithMutexIDExpression, req.ProtoReflect())
		if err != nil {
			panic(err)
		}
		opts.WorkflowID = id
	}
	if opts.WorkflowIDReusePolicy == v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v1.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY
	}
	if opts.WorkflowExecutionTimeout == 0 {
		opts.WorkflowRunTimeout = 3600000000000 // 1h0m0s
	}
	ctx = workflow.WithChildOptions(ctx, *opts)
	return &SampleWorkflowWithMutexChildRun{Future: workflow.ExecuteChildWorkflow(ctx, SampleWorkflowWithMutexWorkflowName, req)}
}

// SampleWorkflowWithMutexChildRun describes a child SampleWorkflowWithMutex workflow run
type SampleWorkflowWithMutexChildRun struct {
	Future workflow.ChildWorkflowFuture
}

// Get blocks until the workflow is completed, returning the response value
func (r *SampleWorkflowWithMutexChildRun) Get(ctx workflow.Context) (*SampleWorkflowWithMutexResponse, error) {
	var resp SampleWorkflowWithMutexResponse
	if err := r.Future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Select adds this completion to the selector. Callback can be nil.
func (r *SampleWorkflowWithMutexChildRun) Select(sel workflow.Selector, fn func(SampleWorkflowWithMutexChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future, func(workflow.Future) {
		if fn != nil {
			fn(*r)
		}
	})
}

// SelectStart adds waiting for start to the selector. Callback can be nil.
func (r *SampleWorkflowWithMutexChildRun) SelectStart(sel workflow.Selector, fn func(SampleWorkflowWithMutexChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future.GetChildWorkflowExecution(), func(workflow.Future) {
		if fn != nil {
			fn(*r)
		}
	})
}

// WaitStart waits for the child workflow to start
func (r *SampleWorkflowWithMutexChildRun) WaitStart(ctx workflow.Context) (*workflow.Execution, error) {
	var exec workflow.Execution
	if err := r.Future.GetChildWorkflowExecution().Get(ctx, &exec); err != nil {
		return nil, err
	}
	return &exec, nil
}

// LeaseAcquired sends the corresponding signal request to the child workflow
func (r *SampleWorkflowWithMutexChildRun) LeaseAcquired(ctx workflow.Context, input *LeaseAcquiredRequest) workflow.Future {
	return r.Future.SignalChildWorkflow(ctx, LeaseAcquiredSignalName, input)
}

// AcquireLeaseSignal describes a AcquireLease signal
type AcquireLeaseSignal struct {
	Channel workflow.ReceiveChannel
}

// Receive blocks until a AcquireLease signal is received
func (s *AcquireLeaseSignal) Receive(ctx workflow.Context) (*AcquireLeaseRequest, bool) {
	var resp AcquireLeaseRequest
	more := s.Channel.Receive(ctx, &resp)
	return &resp, more
}

// ReceiveAsync checks for a AcquireLease signal without blocking
func (s *AcquireLeaseSignal) ReceiveAsync() *AcquireLeaseRequest {
	var resp AcquireLeaseRequest
	if ok := s.Channel.ReceiveAsync(&resp); !ok {
		return nil
	}
	return &resp
}

// Select checks for a AcquireLease signal without blocking
func (s *AcquireLeaseSignal) Select(sel workflow.Selector, fn func(*AcquireLeaseRequest)) workflow.Selector {
	return sel.AddReceive(s.Channel, func(workflow.ReceiveChannel, bool) {
		req := s.ReceiveAsync()
		if fn != nil {
			fn(req)
		}
	})
}

// AcquireLeaseExternal sends a AcquireLease signal to an existing workflow
func AcquireLeaseExternal(ctx workflow.Context, workflowID string, runID string, req *AcquireLeaseRequest) workflow.Future {
	return workflow.SignalExternalWorkflow(ctx, workflowID, runID, AcquireLeaseSignalName, req)
}

// LeaseAcquiredSignal describes a LeaseAcquired signal
type LeaseAcquiredSignal struct {
	Channel workflow.ReceiveChannel
}

// Receive blocks until a LeaseAcquired signal is received
func (s *LeaseAcquiredSignal) Receive(ctx workflow.Context) (*LeaseAcquiredRequest, bool) {
	var resp LeaseAcquiredRequest
	more := s.Channel.Receive(ctx, &resp)
	return &resp, more
}

// ReceiveAsync checks for a LeaseAcquired signal without blocking
func (s *LeaseAcquiredSignal) ReceiveAsync() *LeaseAcquiredRequest {
	var resp LeaseAcquiredRequest
	if ok := s.Channel.ReceiveAsync(&resp); !ok {
		return nil
	}
	return &resp
}

// Select checks for a LeaseAcquired signal without blocking
func (s *LeaseAcquiredSignal) Select(sel workflow.Selector, fn func(*LeaseAcquiredRequest)) workflow.Selector {
	return sel.AddReceive(s.Channel, func(workflow.ReceiveChannel, bool) {
		req := s.ReceiveAsync()
		if fn != nil {
			fn(req)
		}
	})
}

// LeaseAcquiredExternal sends a LeaseAcquired signal to an existing workflow
func LeaseAcquiredExternal(ctx workflow.Context, workflowID string, runID string, req *LeaseAcquiredRequest) workflow.Future {
	return workflow.SignalExternalWorkflow(ctx, workflowID, runID, LeaseAcquiredSignalName, req)
}

// RenewLeaseSignal describes a RenewLease signal
type RenewLeaseSignal struct {
	Channel workflow.ReceiveChannel
}

// Receive blocks until a RenewLease signal is received
func (s *RenewLeaseSignal) Receive(ctx workflow.Context) (*RenewLeaseRequest, bool) {
	var resp RenewLeaseRequest
	more := s.Channel.Receive(ctx, &resp)
	return &resp, more
}

// ReceiveAsync checks for a RenewLease signal without blocking
func (s *RenewLeaseSignal) ReceiveAsync() *RenewLeaseRequest {
	var resp RenewLeaseRequest
	if ok := s.Channel.ReceiveAsync(&resp); !ok {
		return nil
	}
	return &resp
}

// Select checks for a RenewLease signal without blocking
func (s *RenewLeaseSignal) Select(sel workflow.Selector, fn func(*RenewLeaseRequest)) workflow.Selector {
	return sel.AddReceive(s.Channel, func(workflow.ReceiveChannel, bool) {
		req := s.ReceiveAsync()
		if fn != nil {
			fn(req)
		}
	})
}

// RenewLeaseExternal sends a RenewLease signal to an existing workflow
func RenewLeaseExternal(ctx workflow.Context, workflowID string, runID string, req *RenewLeaseRequest) workflow.Future {
	return workflow.SignalExternalWorkflow(ctx, workflowID, runID, RenewLeaseSignalName, req)
}

// RevokeLeaseSignal describes a RevokeLease signal
type RevokeLeaseSignal struct {
	Channel workflow.ReceiveChannel
}

// Receive blocks until a RevokeLease signal is received
func (s *RevokeLeaseSignal) Receive(ctx workflow.Context) (*RevokeLeaseRequest, bool) {
	var resp RevokeLeaseRequest
	more := s.Channel.Receive(ctx, &resp)
	return &resp, more
}

// ReceiveAsync checks for a RevokeLease signal without blocking
func (s *RevokeLeaseSignal) ReceiveAsync() *RevokeLeaseRequest {
	var resp RevokeLeaseRequest
	if ok := s.Channel.ReceiveAsync(&resp); !ok {
		return nil
	}
	return &resp
}

// Select checks for a RevokeLease signal without blocking
func (s *RevokeLeaseSignal) Select(sel workflow.Selector, fn func(*RevokeLeaseRequest)) workflow.Selector {
	return sel.AddReceive(s.Channel, func(workflow.ReceiveChannel, bool) {
		req := s.ReceiveAsync()
		if fn != nil {
			fn(req)
		}
	})
}

// RevokeLeaseExternal sends a RevokeLease signal to an existing workflow
func RevokeLeaseExternal(ctx workflow.Context, workflowID string, runID string, req *RevokeLeaseRequest) workflow.Future {
	return workflow.SignalExternalWorkflow(ctx, workflowID, runID, RevokeLeaseSignalName, req)
}

// Activities describes available worker activites
type Activities interface {
	// Mutex provides a mutex over a shared resource
	Mutex(ctx context.Context, req *MutexRequest) error
}

// RegisterActivities registers activities with a worker
func RegisterActivities(r worker.Registry, activities Activities) {
	RegisterMutexActivity(r, activities.Mutex)
}

// RegisterMutexActivity registers a Mutex activity
func RegisterMutexActivity(r worker.Registry, fn func(context.Context, *MutexRequest) error) {
	r.RegisterActivityWithOptions(fn, activity.RegisterOptions{
		Name: MutexActivityName,
	})
}

// MutexFuture describes a Mutex activity execution
type MutexFuture struct {
	Future workflow.Future
}

// Get blocks on a Mutex execution, returning the response
func (f *MutexFuture) Get(ctx workflow.Context) error {
	return f.Future.Get(ctx, nil)
}

// Select adds the Mutex completion to the selector, callback can be nil
func (f *MutexFuture) Select(sel workflow.Selector, fn func(*MutexFuture)) workflow.Selector {
	return sel.AddFuture(f.Future, func(workflow.Future) {
		if fn != nil {
			fn(f)
		}
	})
}

// Mutex provides a mutex over a shared resource
func Mutex(ctx workflow.Context, opts *workflow.ActivityOptions, req *MutexRequest) *MutexFuture {
	if opts == nil {
		activityOpts := workflow.GetActivityOptions(ctx)
		opts = &activityOpts
	}
	ctx = workflow.WithActivityOptions(ctx, *opts)
	return &MutexFuture{Future: workflow.ExecuteActivity(ctx, MutexActivityName, req)}
}

// Mutex provides a mutex over a shared resource
func MutexLocal(ctx workflow.Context, opts *workflow.LocalActivityOptions, fn func(context.Context, *MutexRequest) error, req *MutexRequest) *MutexFuture {
	if opts == nil {
		activityOpts := workflow.GetLocalActivityOptions(ctx)
		opts = &activityOpts
	}
	ctx = workflow.WithLocalActivityOptions(ctx, *opts)
	var activity any
	if fn == nil {
		activity = MutexActivityName
	} else {
		activity = fn
	}
	return &MutexFuture{Future: workflow.ExecuteLocalActivity(ctx, activity, req)}
}
