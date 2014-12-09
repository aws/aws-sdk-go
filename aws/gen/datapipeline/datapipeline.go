// Package datapipeline provides a client for AWS Data Pipeline.
package datapipeline

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// DataPipeline is a client for AWS Data Pipeline.
type DataPipeline struct {
	client *aws.JSONClient
}

// New returns a new DataPipeline client.
func New(key, secret, region string, client *http.Client) *DataPipeline {
	if client == nil {
		client = http.DefaultClient
	}

	service := "datapipeline"
	endpoint, service, region := endpoints.Lookup("datapipeline", region)

	return &DataPipeline{
		client: &aws.JSONClient{
			Signer: &aws.V4Signer{
				Key:     key,
				Secret:  secret,
				Service: service,
				Region:  region,
				IncludeXAmzContentSha256: true,
			},
			Client:       client,
			Endpoint:     endpoint,
			JSONVersion:  "1.1",
			TargetPrefix: "DataPipeline",
		},
	}
}

// ActivatePipeline validates a pipeline and initiates processing. If the
// pipeline does not pass validation, activation fails. You cannot perform
// this operation on pipelines and attempting to do so will return an
// InvalidRequestException. Call this action to start processing pipeline
// tasks of a pipeline you've created using the CreatePipeline and
// PutPipelineDefinition actions. A pipeline cannot be modified after it
// has been successfully activated.
func (c *DataPipeline) ActivatePipeline(req ActivatePipelineInput) (resp *ActivatePipelineOutput, err error) {
	resp = &ActivatePipelineOutput{}
	err = c.client.Do("ActivatePipeline", "POST", "/", req, resp)
	return
}

// CreatePipeline creates a new empty pipeline. When this action succeeds,
// you can then use the PutPipelineDefinition action to populate the
// pipeline.
func (c *DataPipeline) CreatePipeline(req CreatePipelineInput) (resp *CreatePipelineOutput, err error) {
	resp = &CreatePipelineOutput{}
	err = c.client.Do("CreatePipeline", "POST", "/", req, resp)
	return
}

// DeletePipeline permanently deletes a pipeline, its pipeline definition
// and its run history. You cannot query or restore a deleted pipeline. AWS
// Data Pipeline will attempt to cancel instances associated with the
// pipeline that are currently being processed by task runners. Deleting a
// pipeline cannot be undone. To temporarily pause a pipeline instead of
// deleting it, call SetStatus with the status set to Pause on individual
// components. Components that are paused by SetStatus can be resumed.
func (c *DataPipeline) DeletePipeline(req DeletePipelineInput) (err error) {
	// NRE
	err = c.client.Do("DeletePipeline", "POST", "/", req, nil)
	return
}

// DescribeObjects returns the object definitions for a set of objects
// associated with the pipeline. Object definitions are composed of a set
// of fields that define the properties of the object.
func (c *DataPipeline) DescribeObjects(req DescribeObjectsInput) (resp *DescribeObjectsOutput, err error) {
	resp = &DescribeObjectsOutput{}
	err = c.client.Do("DescribeObjects", "POST", "/", req, resp)
	return
}

// DescribePipelines retrieve metadata about one or more pipelines. The
// information retrieved includes the name of the pipeline, the pipeline
// identifier, its current state, and the user account that owns the
// pipeline. Using account credentials, you can retrieve metadata about
// pipelines that you or your IAM users have created. If you are using an
// IAM user account, you can retrieve metadata about only those pipelines
// you have read permission for. To retrieve the full pipeline definition
// instead of metadata about the pipeline, call the GetPipelineDefinition
// action.
func (c *DataPipeline) DescribePipelines(req DescribePipelinesInput) (resp *DescribePipelinesOutput, err error) {
	resp = &DescribePipelinesOutput{}
	err = c.client.Do("DescribePipelines", "POST", "/", req, resp)
	return
}

// EvaluateExpression evaluates a string in the context of a specified
// object. A task runner can use this action to evaluate SQL queries stored
// in Amazon S3.
func (c *DataPipeline) EvaluateExpression(req EvaluateExpressionInput) (resp *EvaluateExpressionOutput, err error) {
	resp = &EvaluateExpressionOutput{}
	err = c.client.Do("EvaluateExpression", "POST", "/", req, resp)
	return
}

// GetPipelineDefinition returns the definition of the specified pipeline.
// You can call GetPipelineDefinition to retrieve the pipeline definition
// you provided using PutPipelineDefinition
func (c *DataPipeline) GetPipelineDefinition(req GetPipelineDefinitionInput) (resp *GetPipelineDefinitionOutput, err error) {
	resp = &GetPipelineDefinitionOutput{}
	err = c.client.Do("GetPipelineDefinition", "POST", "/", req, resp)
	return
}

// ListPipelines returns a list of pipeline identifiers for all active
// pipelines. Identifiers are returned only for pipelines you have
// permission to access.
func (c *DataPipeline) ListPipelines(req ListPipelinesInput) (resp *ListPipelinesOutput, err error) {
	resp = &ListPipelinesOutput{}
	err = c.client.Do("ListPipelines", "POST", "/", req, resp)
	return
}

// PollForTask task runners call this action to receive a task to perform
// from AWS Data Pipeline. The task runner specifies which tasks it can
// perform by setting a value for the workerGroup parameter of the
// PollForTask call. The task returned by PollForTask may come from any of
// the pipelines that match the workerGroup value passed in by the task
// runner and that was launched using the IAM user credentials specified by
// the task runner. If tasks are ready in the work queue, PollForTask
// returns a response immediately. If no tasks are available in the queue,
// PollForTask uses long-polling and holds on to a poll connection for up
// to a 90 seconds during which time the first newly scheduled task is
// handed to the task runner. To accomodate this, set the socket timeout in
// your task runner to 90 seconds. The task runner should not call
// PollForTask again on the same workerGroup until it receives a response,
// and this may take up to 90 seconds.
func (c *DataPipeline) PollForTask(req PollForTaskInput) (resp *PollForTaskOutput, err error) {
	resp = &PollForTaskOutput{}
	err = c.client.Do("PollForTask", "POST", "/", req, resp)
	return
}

// PutPipelineDefinition adds tasks, schedules, and preconditions that
// control the behavior of the pipeline. You can use PutPipelineDefinition
// to populate a new pipeline. PutPipelineDefinition also validates the
// configuration as it adds it to the pipeline. Changes to the pipeline are
// saved unless one of the following three validation errors exists in the
// pipeline. An object is missing a name or identifier field. A string or
// reference field is empty. The number of objects in the pipeline exceeds
// the maximum allowed objects. The pipeline is in a state. Pipeline object
// definitions are passed to the PutPipelineDefinition action and returned
// by the GetPipelineDefinition action.
func (c *DataPipeline) PutPipelineDefinition(req PutPipelineDefinitionInput) (resp *PutPipelineDefinitionOutput, err error) {
	resp = &PutPipelineDefinitionOutput{}
	err = c.client.Do("PutPipelineDefinition", "POST", "/", req, resp)
	return
}

// QueryObjects queries a pipeline for the names of objects that match a
// specified set of conditions. The objects returned by QueryObjects are
// paginated and then filtered by the value you set for query. This means
// the action may return an empty result set with a value set for marker.
// If HasMoreResults is set to True , you should continue to call
// QueryObjects , passing in the returned value for marker, until
// HasMoreResults returns False
func (c *DataPipeline) QueryObjects(req QueryObjectsInput) (resp *QueryObjectsOutput, err error) {
	resp = &QueryObjectsOutput{}
	err = c.client.Do("QueryObjects", "POST", "/", req, resp)
	return
}

// ReportTaskProgress updates the AWS Data Pipeline service on the progress
// of the calling task runner. When the task runner is assigned a task, it
// should call ReportTaskProgress to acknowledge that it has the task
// within 2 minutes. If the web service does not recieve this
// acknowledgement within the 2 minute window, it will assign the task in a
// subsequent PollForTask call. After this initial acknowledgement, the
// task runner only needs to report progress every 15 minutes to maintain
// its ownership of the task. You can change this reporting time from 15
// minutes by specifying a reportProgressTimeout field in your pipeline. If
// a task runner does not report its status after 5 minutes, AWS Data
// Pipeline will assume that the task runner is unable to process the task
// and will reassign the task in a subsequent response to PollForTask .
// task runners should call ReportTaskProgress every 60 seconds.
func (c *DataPipeline) ReportTaskProgress(req ReportTaskProgressInput) (resp *ReportTaskProgressOutput, err error) {
	resp = &ReportTaskProgressOutput{}
	err = c.client.Do("ReportTaskProgress", "POST", "/", req, resp)
	return
}

// ReportTaskRunnerHeartbeat task runners call ReportTaskRunnerHeartbeat
// every 15 minutes to indicate that they are operational. In the case of
// AWS Data Pipeline Task Runner launched on a resource managed by AWS Data
// Pipeline, the web service can use this call to detect when the task
// runner application has failed and restart a new instance.
func (c *DataPipeline) ReportTaskRunnerHeartbeat(req ReportTaskRunnerHeartbeatInput) (resp *ReportTaskRunnerHeartbeatOutput, err error) {
	resp = &ReportTaskRunnerHeartbeatOutput{}
	err = c.client.Do("ReportTaskRunnerHeartbeat", "POST", "/", req, resp)
	return
}

// SetStatus requests that the status of an array of physical or logical
// pipeline objects be updated in the pipeline. This update may not occur
// immediately, but is eventually consistent. The status that can be set
// depends on the type of object, e.g. DataNode or Activity. You cannot
// perform this operation on pipelines and attempting to do so will return
// an InvalidRequestException.
func (c *DataPipeline) SetStatus(req SetStatusInput) (err error) {
	// NRE
	err = c.client.Do("SetStatus", "POST", "/", req, nil)
	return
}

// SetTaskStatus notifies AWS Data Pipeline that a task is completed and
// provides information about the final status. The task runner calls this
// action regardless of whether the task was sucessful. The task runner
// does not need to call SetTaskStatus for tasks that are canceled by the
// web service during a call to ReportTaskProgress .
func (c *DataPipeline) SetTaskStatus(req SetTaskStatusInput) (resp *SetTaskStatusOutput, err error) {
	resp = &SetTaskStatusOutput{}
	err = c.client.Do("SetTaskStatus", "POST", "/", req, resp)
	return
}

// ValidatePipelineDefinition tests the pipeline definition with a set of
// validation checks to ensure that it is well formed and can run without
// error.
func (c *DataPipeline) ValidatePipelineDefinition(req ValidatePipelineDefinitionInput) (resp *ValidatePipelineDefinitionOutput, err error) {
	resp = &ValidatePipelineDefinitionOutput{}
	err = c.client.Do("ValidatePipelineDefinition", "POST", "/", req, resp)
	return
}

// ActivatePipelineInput is undocumented.
type ActivatePipelineInput struct {
	ParameterValues []ParameterValue `json:"parameterValues,omitempty"`
	PipelineID      string           `json:"pipelineId"`
}

// ActivatePipelineOutput is undocumented.
type ActivatePipelineOutput struct {
}

// CreatePipelineInput is undocumented.
type CreatePipelineInput struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	UniqueID    string `json:"uniqueId"`
}

// CreatePipelineOutput is undocumented.
type CreatePipelineOutput struct {
	PipelineID string `json:"pipelineId"`
}

// DeletePipelineInput is undocumented.
type DeletePipelineInput struct {
	PipelineID string `json:"pipelineId"`
}

// DescribeObjectsInput is undocumented.
type DescribeObjectsInput struct {
	EvaluateExpressions bool     `json:"evaluateExpressions,omitempty"`
	Marker              string   `json:"marker,omitempty"`
	ObjectIds           []string `json:"objectIds"`
	PipelineID          string   `json:"pipelineId"`
}

// DescribeObjectsOutput is undocumented.
type DescribeObjectsOutput struct {
	HasMoreResults  bool             `json:"hasMoreResults,omitempty"`
	Marker          string           `json:"marker,omitempty"`
	PipelineObjects []PipelineObject `json:"pipelineObjects"`
}

// DescribePipelinesInput is undocumented.
type DescribePipelinesInput struct {
	PipelineIds []string `json:"pipelineIds"`
}

// DescribePipelinesOutput is undocumented.
type DescribePipelinesOutput struct {
	PipelineDescriptionList []PipelineDescription `json:"pipelineDescriptionList"`
}

// EvaluateExpressionInput is undocumented.
type EvaluateExpressionInput struct {
	Expression string `json:"expression"`
	ObjectID   string `json:"objectId"`
	PipelineID string `json:"pipelineId"`
}

// EvaluateExpressionOutput is undocumented.
type EvaluateExpressionOutput struct {
	EvaluatedExpression string `json:"evaluatedExpression"`
}

// Field is undocumented.
type Field struct {
	Key         string `json:"key"`
	RefValue    string `json:"refValue,omitempty"`
	StringValue string `json:"stringValue,omitempty"`
}

// GetPipelineDefinitionInput is undocumented.
type GetPipelineDefinitionInput struct {
	PipelineID string `json:"pipelineId"`
	Version    string `json:"version,omitempty"`
}

// GetPipelineDefinitionOutput is undocumented.
type GetPipelineDefinitionOutput struct {
	ParameterObjects []ParameterObject `json:"parameterObjects,omitempty"`
	ParameterValues  []ParameterValue  `json:"parameterValues,omitempty"`
	PipelineObjects  []PipelineObject  `json:"pipelineObjects,omitempty"`
}

// InstanceIdentity is undocumented.
type InstanceIdentity struct {
	Document  string `json:"document,omitempty"`
	Signature string `json:"signature,omitempty"`
}

// ListPipelinesInput is undocumented.
type ListPipelinesInput struct {
	Marker string `json:"marker,omitempty"`
}

// ListPipelinesOutput is undocumented.
type ListPipelinesOutput struct {
	HasMoreResults bool             `json:"hasMoreResults,omitempty"`
	Marker         string           `json:"marker,omitempty"`
	PipelineIDList []PipelineIDName `json:"pipelineIdList"`
}

// Operator is undocumented.
type Operator struct {
	Type   string   `json:"type,omitempty"`
	Values []string `json:"values,omitempty"`
}

// ParameterAttribute is undocumented.
type ParameterAttribute struct {
	Key         string `json:"key"`
	StringValue string `json:"stringValue"`
}

// ParameterObject is undocumented.
type ParameterObject struct {
	Attributes []ParameterAttribute `json:"attributes"`
	ID         string               `json:"id"`
}

// ParameterValue is undocumented.
type ParameterValue struct {
	ID          string `json:"id"`
	StringValue string `json:"stringValue"`
}

// PipelineDescription is undocumented.
type PipelineDescription struct {
	Description string  `json:"description,omitempty"`
	Fields      []Field `json:"fields"`
	Name        string  `json:"name"`
	PipelineID  string  `json:"pipelineId"`
}

// PipelineIDName is undocumented.
type PipelineIDName struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// PipelineObject is undocumented.
type PipelineObject struct {
	Fields []Field `json:"fields"`
	ID     string  `json:"id"`
	Name   string  `json:"name"`
}

// PollForTaskInput is undocumented.
type PollForTaskInput struct {
	Hostname         string           `json:"hostname,omitempty"`
	InstanceIdentity InstanceIdentity `json:"instanceIdentity,omitempty"`
	WorkerGroup      string           `json:"workerGroup"`
}

// PollForTaskOutput is undocumented.
type PollForTaskOutput struct {
	TaskObject TaskObject `json:"taskObject,omitempty"`
}

// PutPipelineDefinitionInput is undocumented.
type PutPipelineDefinitionInput struct {
	ParameterObjects []ParameterObject `json:"parameterObjects,omitempty"`
	ParameterValues  []ParameterValue  `json:"parameterValues,omitempty"`
	PipelineID       string            `json:"pipelineId"`
	PipelineObjects  []PipelineObject  `json:"pipelineObjects"`
}

// PutPipelineDefinitionOutput is undocumented.
type PutPipelineDefinitionOutput struct {
	Errored            bool                `json:"errored"`
	ValidationErrors   []ValidationError   `json:"validationErrors,omitempty"`
	ValidationWarnings []ValidationWarning `json:"validationWarnings,omitempty"`
}

// Query is undocumented.
type Query struct {
	Selectors []Selector `json:"selectors,omitempty"`
}

// QueryObjectsInput is undocumented.
type QueryObjectsInput struct {
	Limit      int    `json:"limit,omitempty"`
	Marker     string `json:"marker,omitempty"`
	PipelineID string `json:"pipelineId"`
	Query      Query  `json:"query,omitempty"`
	Sphere     string `json:"sphere"`
}

// QueryObjectsOutput is undocumented.
type QueryObjectsOutput struct {
	HasMoreResults bool     `json:"hasMoreResults,omitempty"`
	Ids            []string `json:"ids,omitempty"`
	Marker         string   `json:"marker,omitempty"`
}

// ReportTaskProgressInput is undocumented.
type ReportTaskProgressInput struct {
	Fields []Field `json:"fields,omitempty"`
	TaskID string  `json:"taskId"`
}

// ReportTaskProgressOutput is undocumented.
type ReportTaskProgressOutput struct {
	Canceled bool `json:"canceled"`
}

// ReportTaskRunnerHeartbeatInput is undocumented.
type ReportTaskRunnerHeartbeatInput struct {
	Hostname     string `json:"hostname,omitempty"`
	TaskrunnerID string `json:"taskrunnerId"`
	WorkerGroup  string `json:"workerGroup,omitempty"`
}

// ReportTaskRunnerHeartbeatOutput is undocumented.
type ReportTaskRunnerHeartbeatOutput struct {
	Terminate bool `json:"terminate"`
}

// Selector is undocumented.
type Selector struct {
	FieldName string   `json:"fieldName,omitempty"`
	Operator  Operator `json:"operator,omitempty"`
}

// SetStatusInput is undocumented.
type SetStatusInput struct {
	ObjectIds  []string `json:"objectIds"`
	PipelineID string   `json:"pipelineId"`
	Status     string   `json:"status"`
}

// SetTaskStatusInput is undocumented.
type SetTaskStatusInput struct {
	ErrorID         string `json:"errorId,omitempty"`
	ErrorMessage    string `json:"errorMessage,omitempty"`
	ErrorStackTrace string `json:"errorStackTrace,omitempty"`
	TaskID          string `json:"taskId"`
	TaskStatus      string `json:"taskStatus"`
}

// SetTaskStatusOutput is undocumented.
type SetTaskStatusOutput struct {
}

// TaskObject is undocumented.
type TaskObject struct {
	AttemptID  string                    `json:"attemptId,omitempty"`
	Objects    map[string]PipelineObject `json:"objects,omitempty"`
	PipelineID string                    `json:"pipelineId,omitempty"`
	TaskID     string                    `json:"taskId,omitempty"`
}

// ValidatePipelineDefinitionInput is undocumented.
type ValidatePipelineDefinitionInput struct {
	ParameterObjects []ParameterObject `json:"parameterObjects,omitempty"`
	ParameterValues  []ParameterValue  `json:"parameterValues,omitempty"`
	PipelineID       string            `json:"pipelineId"`
	PipelineObjects  []PipelineObject  `json:"pipelineObjects"`
}

// ValidatePipelineDefinitionOutput is undocumented.
type ValidatePipelineDefinitionOutput struct {
	Errored            bool                `json:"errored"`
	ValidationErrors   []ValidationError   `json:"validationErrors,omitempty"`
	ValidationWarnings []ValidationWarning `json:"validationWarnings,omitempty"`
}

// ValidationError is undocumented.
type ValidationError struct {
	Errors []string `json:"errors,omitempty"`
	ID     string   `json:"id,omitempty"`
}

// ValidationWarning is undocumented.
type ValidationWarning struct {
	ID       string   `json:"id,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
