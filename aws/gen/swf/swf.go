// Package swf provides a client for Amazon Simple Workflow Service.
package swf

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// SWF is a client for Amazon Simple Workflow Service.
type SWF struct {
	client aws.Client
}

// New returns a new SWF client.
func New(key, secret, region string, client *http.Client) *SWF {
	if client == nil {
		client = http.DefaultClient
	}

	return &SWF{
		client: &aws.JSONClient{
			Client: client,
			Auth: aws.Auth{
				Key:     key,
				Secret:  secret,
				Service: "swf",
				Region:  region,
			},
			Endpoint:     endpoints.Lookup("swf", region),
			JSONVersion:  "1.0",
			TargetPrefix: "SimpleWorkflowService",
		},
	}
}

// CountClosedWorkflowExecutions returns the number of closed workflow
// executions within the given domain that meet the specified filtering
// criteria. You can use IAM policies to control this action's access to
// Amazon SWF resources as follows: Use a Resource element with the domain
// name to limit the action to only specified domains. Use an Action
// element to allow or deny permission to call this action. Constrain the
// following parameters by using a Condition element with the appropriate
// keys. tagFilter.tag : String constraint. The key is swf:tagFilter.tag
// typeFilter.name : String constraint. The key is swf:typeFilter.name
// typeFilter.version : String constraint. The key is
// swf:typeFilter.version If the caller does not have sufficient
// permissions to invoke the action, or the parameter values fall outside
// the specified constraints, the action fails by throwing
// OperationNotPermitted . For details and example IAM policies, see Using
// IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) CountClosedWorkflowExecutions(req CountClosedWorkflowExecutionsInput) (resp *WorkflowExecutionCount, err error) {
	resp = &WorkflowExecutionCount{}
	err = c.client.Do("CountClosedWorkflowExecutions", "POST", "/", req, resp)
	return
}

// CountOpenWorkflowExecutions returns the number of open workflow
// executions within the given domain that meet the specified filtering
// criteria. You can use IAM policies to control this action's access to
// Amazon SWF resources as follows: Use a Resource element with the domain
// name to limit the action to only specified domains. Use an Action
// element to allow or deny permission to call this action. Constrain the
// following parameters by using a Condition element with the appropriate
// keys. tagFilter.tag : String constraint. The key is swf:tagFilter.tag
// typeFilter.name : String constraint. The key is swf:typeFilter.name
// typeFilter.version : String constraint. The key is
// swf:typeFilter.version If the caller does not have sufficient
// permissions to invoke the action, or the parameter values fall outside
// the specified constraints, the action fails by throwing
// OperationNotPermitted . For details and example IAM policies, see Using
// IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) CountOpenWorkflowExecutions(req CountOpenWorkflowExecutionsInput) (resp *WorkflowExecutionCount, err error) {
	resp = &WorkflowExecutionCount{}
	err = c.client.Do("CountOpenWorkflowExecutions", "POST", "/", req, resp)
	return
}

// CountPendingActivityTasks returns the estimated number of activity tasks
// in the specified task list. The count returned is an approximation and
// is not guaranteed to be exact. If you specify a task list that no
// activity task was ever scheduled in then 0 will be returned. You can use
// IAM policies to control this action's access to Amazon SWF resources as
// follows: Use a Resource element with the domain name to limit the action
// to only specified domains. Use an Action element to allow or deny
// permission to call this action. Constrain the taskList.name parameter by
// using a Condition element with the swf:taskList.name key to allow the
// action to access only certain task lists. If the caller does not have
// sufficient permissions to invoke the action, or the parameter values
// fall outside the specified constraints, the action fails by throwing
// OperationNotPermitted . For details and example IAM policies, see Using
// IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) CountPendingActivityTasks(req CountPendingActivityTasksInput) (resp *PendingTaskCount, err error) {
	resp = &PendingTaskCount{}
	err = c.client.Do("CountPendingActivityTasks", "POST", "/", req, resp)
	return
}

// CountPendingDecisionTasks returns the estimated number of decision tasks
// in the specified task list. The count returned is an approximation and
// is not guaranteed to be exact. If you specify a task list that no
// decision task was ever scheduled in then 0 will be returned. You can use
// IAM policies to control this action's access to Amazon SWF resources as
// follows: Use a Resource element with the domain name to limit the action
// to only specified domains. Use an Action element to allow or deny
// permission to call this action. Constrain the taskList.name parameter by
// using a Condition element with the swf:taskList.name key to allow the
// action to access only certain task lists. If the caller does not have
// sufficient permissions to invoke the action, or the parameter values
// fall outside the specified constraints, the action fails by throwing
// OperationNotPermitted . For details and example IAM policies, see Using
// IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) CountPendingDecisionTasks(req CountPendingDecisionTasksInput) (resp *PendingTaskCount, err error) {
	resp = &PendingTaskCount{}
	err = c.client.Do("CountPendingDecisionTasks", "POST", "/", req, resp)
	return
}

// DeprecateActivityType deprecates the specified activity type . After an
// activity type has been deprecated, you cannot create new tasks of that
// activity type. Tasks of this type that were scheduled before the type
// was deprecated will continue to run. You can use IAM policies to control
// this action's access to Amazon SWF resources as follows: Use a Resource
// element with the domain name to limit the action to only specified
// domains. Use an Action element to allow or deny permission to call this
// action. Constrain the following parameters by using a Condition element
// with the appropriate keys. activityType.name : String constraint. The
// key is swf:activityType.name activityType.version : String constraint.
// The key is swf:activityType.version If the caller does not have
// sufficient permissions to invoke the action, or the parameter values
// fall outside the specified constraints, the action fails by throwing
// OperationNotPermitted . For details and example IAM policies, see Using
// IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) DeprecateActivityType(req DeprecateActivityTypeInput) (err error) {
	// NRE
	err = c.client.Do("DeprecateActivityType", "POST", "/", req, nil)
	return
}

// DeprecateDomain deprecates the specified domain. After a domain has been
// deprecated it cannot be used to create new workflow executions or
// register new types. However, you can still use visibility actions on
// this domain. Deprecating a domain also deprecates all activity and
// workflow types registered in the domain. Executions that were started
// before the domain was deprecated will continue to run. You can use IAM
// policies to control this action's access to Amazon SWF resources as
// follows: Use a Resource element with the domain name to limit the action
// to only specified domains. Use an Action element to allow or deny
// permission to call this action. You cannot use an IAM policy to
// constrain this action's parameters. If the caller does not have
// sufficient permissions to invoke the action, or the parameter values
// fall outside the specified constraints, the action fails by throwing
// OperationNotPermitted . For details and example IAM policies, see Using
// IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) DeprecateDomain(req DeprecateDomainInput) (err error) {
	// NRE
	err = c.client.Do("DeprecateDomain", "POST", "/", req, nil)
	return
}

// DeprecateWorkflowType deprecates the specified workflow type . After a
// workflow type has been deprecated, you cannot create new executions of
// that type. Executions that were started before the type was deprecated
// will continue to run. A deprecated workflow type may still be used when
// calling visibility actions. You can use IAM policies to control this
// action's access to Amazon SWF resources as follows: Use a Resource
// element with the domain name to limit the action to only specified
// domains. Use an Action element to allow or deny permission to call this
// action. Constrain the following parameters by using a Condition element
// with the appropriate keys. workflowType.name : String constraint. The
// key is swf:workflowType.name workflowType.version : String constraint.
// The key is swf:workflowType.version If the caller does not have
// sufficient permissions to invoke the action, or the parameter values
// fall outside the specified constraints, the action fails by throwing
// OperationNotPermitted . For details and example IAM policies, see Using
// IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) DeprecateWorkflowType(req DeprecateWorkflowTypeInput) (err error) {
	// NRE
	err = c.client.Do("DeprecateWorkflowType", "POST", "/", req, nil)
	return
}

// DescribeActivityType returns information about the specified activity
// type. This includes configuration settings provided at registration time
// as well as other general information about the type. You can use IAM
// policies to control this action's access to Amazon SWF resources as
// follows: Use a Resource element with the domain name to limit the action
// to only specified domains. Use an Action element to allow or deny
// permission to call this action. Constrain the following parameters by
// using a Condition element with the appropriate keys. activityType.name :
// String constraint. The key is swf:activityType.name activityType.version
// : String constraint. The key is swf:activityType.version If the caller
// does not have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action
// fails by throwing OperationNotPermitted . For details and example IAM
// policies, see Using IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) DescribeActivityType(req DescribeActivityTypeInput) (resp *ActivityTypeDetail, err error) {
	resp = &ActivityTypeDetail{}
	err = c.client.Do("DescribeActivityType", "POST", "/", req, resp)
	return
}

// DescribeDomain returns information about the specified domain including
// description and status. You can use IAM policies to control this
// action's access to Amazon SWF resources as follows: Use a Resource
// element with the domain name to limit the action to only specified
// domains. Use an Action element to allow or deny permission to call this
// action. You cannot use an IAM policy to constrain this action's
// parameters. If the caller does not have sufficient permissions to invoke
// the action, or the parameter values fall outside the specified
// constraints, the action fails by throwing OperationNotPermitted . For
// details and example IAM policies, see Using IAM to Manage Access to
// Amazon SWF Workflows
func (c *SWF) DescribeDomain(req DescribeDomainInput) (resp *DomainDetail, err error) {
	resp = &DomainDetail{}
	err = c.client.Do("DescribeDomain", "POST", "/", req, resp)
	return
}

// DescribeWorkflowExecution returns information about the specified
// workflow execution including its type and some statistics. You can use
// IAM policies to control this action's access to Amazon SWF resources as
// follows: Use a Resource element with the domain name to limit the action
// to only specified domains. Use an Action element to allow or deny
// permission to call this action. You cannot use an IAM policy to
// constrain this action's parameters. If the caller does not have
// sufficient permissions to invoke the action, or the parameter values
// fall outside the specified constraints, the action fails by throwing
// OperationNotPermitted . For details and example IAM policies, see Using
// IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) DescribeWorkflowExecution(req DescribeWorkflowExecutionInput) (resp *WorkflowExecutionDetail, err error) {
	resp = &WorkflowExecutionDetail{}
	err = c.client.Do("DescribeWorkflowExecution", "POST", "/", req, resp)
	return
}

// DescribeWorkflowType returns information about the specified workflow
// type . This includes configuration settings specified when the type was
// registered and other information such as creation date, current status,
// etc. You can use IAM policies to control this action's access to Amazon
// SWF resources as follows: Use a Resource element with the domain name to
// limit the action to only specified domains. Use an Action element to
// allow or deny permission to call this action. Constrain the following
// parameters by using a Condition element with the appropriate keys.
// workflowType.name : String constraint. The key is swf:workflowType.name
// workflowType.version : String constraint. The key is
// swf:workflowType.version If the caller does not have sufficient
// permissions to invoke the action, or the parameter values fall outside
// the specified constraints, the action fails by throwing
// OperationNotPermitted . For details and example IAM policies, see Using
// IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) DescribeWorkflowType(req DescribeWorkflowTypeInput) (resp *WorkflowTypeDetail, err error) {
	resp = &WorkflowTypeDetail{}
	err = c.client.Do("DescribeWorkflowType", "POST", "/", req, resp)
	return
}

// GetWorkflowExecutionHistory returns the history of the specified
// workflow execution. The results may be split into multiple pages. To
// retrieve subsequent pages, make the call again using the nextPageToken
// returned by the initial call. You can use IAM policies to control this
// action's access to Amazon SWF resources as follows: Use a Resource
// element with the domain name to limit the action to only specified
// domains. Use an Action element to allow or deny permission to call this
// action. You cannot use an IAM policy to constrain this action's
// parameters. If the caller does not have sufficient permissions to invoke
// the action, or the parameter values fall outside the specified
// constraints, the action fails by throwing OperationNotPermitted . For
// details and example IAM policies, see Using IAM to Manage Access to
// Amazon SWF Workflows
func (c *SWF) GetWorkflowExecutionHistory(req GetWorkflowExecutionHistoryInput) (resp *History, err error) {
	resp = &History{}
	err = c.client.Do("GetWorkflowExecutionHistory", "POST", "/", req, resp)
	return
}

// ListActivityTypes returns information about all activities registered in
// the specified domain that match the specified name and registration
// status. The result includes information like creation date, current
// status of the activity, etc. The results may be split into multiple
// pages. To retrieve subsequent pages, make the call again using the
// nextPageToken returned by the initial call. You can use IAM policies to
// control this action's access to Amazon SWF resources as follows: Use a
// Resource element with the domain name to limit the action to only
// specified domains. Use an Action element to allow or deny permission to
// call this action. You cannot use an IAM policy to constrain this
// action's parameters. If the caller does not have sufficient permissions
// to invoke the action, or the parameter values fall outside the specified
// constraints, the action fails by throwing OperationNotPermitted . For
// details and example IAM policies, see Using IAM to Manage Access to
// Amazon SWF Workflows
func (c *SWF) ListActivityTypes(req ListActivityTypesInput) (resp *ActivityTypeInfos, err error) {
	resp = &ActivityTypeInfos{}
	err = c.client.Do("ListActivityTypes", "POST", "/", req, resp)
	return
}

// ListClosedWorkflowExecutions returns a list of closed workflow
// executions in the specified domain that meet the filtering criteria. The
// results may be split into multiple pages. To retrieve subsequent pages,
// make the call again using the nextPageToken returned by the initial
// call. You can use IAM policies to control this action's access to Amazon
// SWF resources as follows: Use a Resource element with the domain name to
// limit the action to only specified domains. Use an Action element to
// allow or deny permission to call this action. Constrain the following
// parameters by using a Condition element with the appropriate keys.
// tagFilter.tag : String constraint. The key is swf:tagFilter.tag
// typeFilter.name : String constraint. The key is swf:typeFilter.name
// typeFilter.version : String constraint. The key is
// swf:typeFilter.version If the caller does not have sufficient
// permissions to invoke the action, or the parameter values fall outside
// the specified constraints, the action fails by throwing
// OperationNotPermitted . For details and example IAM policies, see Using
// IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) ListClosedWorkflowExecutions(req ListClosedWorkflowExecutionsInput) (resp *WorkflowExecutionInfos, err error) {
	resp = &WorkflowExecutionInfos{}
	err = c.client.Do("ListClosedWorkflowExecutions", "POST", "/", req, resp)
	return
}

// ListDomains returns the list of domains registered in the account. The
// results may be split into multiple pages. To retrieve subsequent pages,
// make the call again using the nextPageToken returned by the initial
// call. You can use IAM policies to control this action's access to Amazon
// SWF resources as follows: Use a Resource element with the domain name to
// limit the action to only specified domains. The element must be set to
// arn:aws:swf::AccountID:domain/*" , where "AccountID" is the account ID,
// with no dashes. Use an Action element to allow or deny permission to
// call this action. You cannot use an IAM policy to constrain this
// action's parameters. If the caller does not have sufficient permissions
// to invoke the action, or the parameter values fall outside the specified
// constraints, the action fails by throwing OperationNotPermitted . For
// details and example IAM policies, see Using IAM to Manage Access to
// Amazon SWF Workflows
func (c *SWF) ListDomains(req ListDomainsInput) (resp *DomainInfos, err error) {
	resp = &DomainInfos{}
	err = c.client.Do("ListDomains", "POST", "/", req, resp)
	return
}

// ListOpenWorkflowExecutions returns a list of open workflow executions in
// the specified domain that meet the filtering criteria. The results may
// be split into multiple pages. To retrieve subsequent pages, make the
// call again using the nextPageToken returned by the initial call. You can
// use IAM policies to control this action's access to Amazon SWF resources
// as follows: Use a Resource element with the domain name to limit the
// action to only specified domains. Use an Action element to allow or deny
// permission to call this action. Constrain the following parameters by
// using a Condition element with the appropriate keys. tagFilter.tag :
// String constraint. The key is swf:tagFilter.tag typeFilter.name : String
// constraint. The key is swf:typeFilter.name typeFilter.version : String
// constraint. The key is swf:typeFilter.version If the caller does not
// have sufficient permissions to invoke the action, or the parameter
// values fall outside the specified constraints, the action fails by
// throwing OperationNotPermitted . For details and example IAM policies,
// see Using IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) ListOpenWorkflowExecutions(req ListOpenWorkflowExecutionsInput) (resp *WorkflowExecutionInfos, err error) {
	resp = &WorkflowExecutionInfos{}
	err = c.client.Do("ListOpenWorkflowExecutions", "POST", "/", req, resp)
	return
}

// ListWorkflowTypes returns information about workflow types in the
// specified domain. The results may be split into multiple pages that can
// be retrieved by making the call repeatedly. You can use IAM policies to
// control this action's access to Amazon SWF resources as follows: Use a
// Resource element with the domain name to limit the action to only
// specified domains. Use an Action element to allow or deny permission to
// call this action. You cannot use an IAM policy to constrain this
// action's parameters. If the caller does not have sufficient permissions
// to invoke the action, or the parameter values fall outside the specified
// constraints, the action fails by throwing OperationNotPermitted . For
// details and example IAM policies, see Using IAM to Manage Access to
// Amazon SWF Workflows
func (c *SWF) ListWorkflowTypes(req ListWorkflowTypesInput) (resp *WorkflowTypeInfos, err error) {
	resp = &WorkflowTypeInfos{}
	err = c.client.Do("ListWorkflowTypes", "POST", "/", req, resp)
	return
}

// PollForActivityTask used by workers to get an ActivityTask from the
// specified activity taskList . This initiates a long poll, where the
// service holds the connection open and responds as soon as a task becomes
// available. The maximum time the service holds on to the request before
// responding is 60 seconds. If no task is available within 60 seconds, the
// poll will return an empty result. An empty result, in this context,
// means that an ActivityTask is returned, but that the value of taskToken
// is an empty string. If a task is returned, the worker should use its
// type to identify and process it correctly. Workers should set their
// client side socket timeout to at least 70 seconds (10 seconds higher
// than the maximum time service may hold the poll request). You can use
// IAM policies to control this action's access to Amazon SWF resources as
// follows: Use a Resource element with the domain name to limit the action
// to only specified domains. Use an Action element to allow or deny
// permission to call this action. Constrain the taskList.name parameter by
// using a Condition element with the swf:taskList.name key to allow the
// action to access only certain task lists. If the caller does not have
// sufficient permissions to invoke the action, or the parameter values
// fall outside the specified constraints, the action fails by throwing
// OperationNotPermitted . For details and example IAM policies, see Using
// IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) PollForActivityTask(req PollForActivityTaskInput) (resp *ActivityTask, err error) {
	resp = &ActivityTask{}
	err = c.client.Do("PollForActivityTask", "POST", "/", req, resp)
	return
}

// PollForDecisionTask used by deciders to get a DecisionTask from the
// specified decision taskList . A decision task may be returned for any
// open workflow execution that is using the specified task list. The task
// includes a paginated view of the history of the workflow execution. The
// decider should use the workflow type and the history to determine how to
// properly handle the task. This action initiates a long poll, where the
// service holds the connection open and responds as soon a task becomes
// available. If no decision task is available in the specified task list
// before the timeout of 60 seconds expires, an empty result is returned.
// An empty result, in this context, means that a DecisionTask is returned,
// but that the value of taskToken is an empty string. Deciders should set
// their client side socket timeout to at least 70 seconds (10 seconds
// higher than the timeout). Because the number of workflow history events
// for a single workflow execution might be very large, the result returned
// might be split up across a number of pages. To retrieve subsequent
// pages, make additional calls to PollForDecisionTask using the
// nextPageToken returned by the initial call. Note that you do not call
// GetWorkflowExecutionHistory with this nextPageToken . Instead, call
// PollForDecisionTask again. You can use IAM policies to control this
// action's access to Amazon SWF resources as follows: Use a Resource
// element with the domain name to limit the action to only specified
// domains. Use an Action element to allow or deny permission to call this
// action. Constrain the taskList.name parameter by using a Condition
// element with the swf:taskList.name key to allow the action to access
// only certain task lists. If the caller does not have sufficient
// permissions to invoke the action, or the parameter values fall outside
// the specified constraints, the action fails by throwing
// OperationNotPermitted . For details and example IAM policies, see Using
// IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) PollForDecisionTask(req PollForDecisionTaskInput) (resp *DecisionTask, err error) {
	resp = &DecisionTask{}
	err = c.client.Do("PollForDecisionTask", "POST", "/", req, resp)
	return
}

// RecordActivityTaskHeartbeat used by activity workers to report to the
// service that the ActivityTask represented by the specified taskToken is
// still making progress. The worker can also (optionally) specify details
// of the progress, for example percent complete, using the details
// parameter. This action can also be used by the worker as a mechanism to
// check if cancellation is being requested for the activity task. If a
// cancellation is being attempted for the specified task, then the boolean
// cancelRequested flag returned by the service is set to true . This
// action resets the taskHeartbeatTimeout clock. The taskHeartbeatTimeout
// is specified in RegisterActivityType . This action does not in itself
// create an event in the workflow execution history. However, if the task
// times out, the workflow execution history will contain a
// ActivityTaskTimedOut event that contains the information from the last
// heartbeat generated by the activity worker. If the cancelRequested flag
// returns true , a cancellation is being attempted. If the worker can
// cancel the activity, it should respond with RespondActivityTaskCanceled
// . Otherwise, it should ignore the cancellation request. You can use IAM
// policies to control this action's access to Amazon SWF resources as
// follows: Use a Resource element with the domain name to limit the action
// to only specified domains. Use an Action element to allow or deny
// permission to call this action. You cannot use an IAM policy to
// constrain this action's parameters. If the caller does not have
// sufficient permissions to invoke the action, or the parameter values
// fall outside the specified constraints, the action fails by throwing
// OperationNotPermitted . For details and example IAM policies, see Using
// IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) RecordActivityTaskHeartbeat(req RecordActivityTaskHeartbeatInput) (resp *ActivityTaskStatus, err error) {
	resp = &ActivityTaskStatus{}
	err = c.client.Do("RecordActivityTaskHeartbeat", "POST", "/", req, resp)
	return
}

// RegisterActivityType registers a new activity type along with its
// configuration settings in the specified domain. A TypeAlreadyExists
// fault is returned if the type already exists in the domain. You cannot
// change any configuration settings of the type after its registration,
// and it must be registered as a new version. You can use IAM policies to
// control this action's access to Amazon SWF resources as follows: Use a
// Resource element with the domain name to limit the action to only
// specified domains. Use an Action element to allow or deny permission to
// call this action. Constrain the following parameters by using a
// Condition element with the appropriate keys. defaultTaskList.name :
// String constraint. The key is swf:defaultTaskList.name name : String
// constraint. The key is swf:name version : String constraint. The key is
// swf:version If the caller does not have sufficient permissions to invoke
// the action, or the parameter values fall outside the specified
// constraints, the action fails by throwing OperationNotPermitted . For
// details and example IAM policies, see Using IAM to Manage Access to
// Amazon SWF Workflows
func (c *SWF) RegisterActivityType(req RegisterActivityTypeInput) (err error) {
	// NRE
	err = c.client.Do("RegisterActivityType", "POST", "/", req, nil)
	return
}

// RegisterDomain you can use IAM policies to control this action's access
// to Amazon SWF resources as follows: You cannot use an IAM policy to
// control domain access for this action. The name of the domain being
// registered is available as the resource of this action. Use an Action
// element to allow or deny permission to call this action. You cannot use
// an IAM policy to constrain this action's parameters. If the caller does
// not have sufficient permissions to invoke the action, or the parameter
// values fall outside the specified constraints, the action fails by
// throwing OperationNotPermitted . For details and example IAM policies,
// see Using IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) RegisterDomain(req RegisterDomainInput) (err error) {
	// NRE
	err = c.client.Do("RegisterDomain", "POST", "/", req, nil)
	return
}

// RegisterWorkflowType registers a new workflow type and its configuration
// settings in the specified domain. The retention period for the workflow
// history is set by the RegisterDomain action. If the type already exists,
// then a TypeAlreadyExists fault is returned. You cannot change the
// configuration settings of a workflow type once it is registered and it
// must be registered as a new version. You can use IAM policies to control
// this action's access to Amazon SWF resources as follows: Use a Resource
// element with the domain name to limit the action to only specified
// domains. Use an Action element to allow or deny permission to call this
// action. Constrain the following parameters by using a Condition element
// with the appropriate keys. defaultTaskList.name : String constraint. The
// key is swf:defaultTaskList.name name : String constraint. The key is
// swf:name version : String constraint. The key is swf:version If the
// caller does not have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action
// fails by throwing OperationNotPermitted . For details and example IAM
// policies, see Using IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) RegisterWorkflowType(req RegisterWorkflowTypeInput) (err error) {
	// NRE
	err = c.client.Do("RegisterWorkflowType", "POST", "/", req, nil)
	return
}

// RequestCancelWorkflowExecution records a
// WorkflowExecutionCancelRequested event in the currently running workflow
// execution identified by the given domain, workflowId, and runId. This
// logically requests the cancellation of the workflow execution as a
// whole. It is up to the decider to take appropriate actions when it
// receives an execution history with this event. You can use IAM policies
// to control this action's access to Amazon SWF resources as follows: Use
// a Resource element with the domain name to limit the action to only
// specified domains. Use an Action element to allow or deny permission to
// call this action. You cannot use an IAM policy to constrain this
// action's parameters. If the caller does not have sufficient permissions
// to invoke the action, or the parameter values fall outside the specified
// constraints, the action fails by throwing OperationNotPermitted . For
// details and example IAM policies, see Using IAM to Manage Access to
// Amazon SWF Workflows
func (c *SWF) RequestCancelWorkflowExecution(req RequestCancelWorkflowExecutionInput) (err error) {
	// NRE
	err = c.client.Do("RequestCancelWorkflowExecution", "POST", "/", req, nil)
	return
}

// RespondActivityTaskCanceled used by workers to tell the service that the
// ActivityTask identified by the taskToken was successfully canceled.
// Additional details can be optionally provided using the details
// argument. These details (if provided) appear in the ActivityTaskCanceled
// event added to the workflow history. Only use this operation if the
// canceled flag of a RecordActivityTaskHeartbeat request returns true and
// if the activity can be safely undone or abandoned. A task is considered
// open from the time that it is scheduled until it is closed. Therefore a
// task is reported as open while a worker is processing it. A task is
// closed after it has been specified in a call to
// RespondActivityTaskCompleted , RespondActivityTaskCanceled,
// RespondActivityTaskFailed , or the task has timed out . You can use IAM
// policies to control this action's access to Amazon SWF resources as
// follows: Use a Resource element with the domain name to limit the action
// to only specified domains. Use an Action element to allow or deny
// permission to call this action. You cannot use an IAM policy to
// constrain this action's parameters. If the caller does not have
// sufficient permissions to invoke the action, or the parameter values
// fall outside the specified constraints, the action fails by throwing
// OperationNotPermitted . For details and example IAM policies, see Using
// IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) RespondActivityTaskCanceled(req RespondActivityTaskCanceledInput) (err error) {
	// NRE
	err = c.client.Do("RespondActivityTaskCanceled", "POST", "/", req, nil)
	return
}

// RespondActivityTaskCompleted used by workers to tell the service that
// the ActivityTask identified by the taskToken completed successfully with
// a result (if provided). The result appears in the ActivityTaskCompleted
// event in the workflow history. If the requested task does not complete
// successfully, use RespondActivityTaskFailed instead. If the worker finds
// that the task is canceled through the canceled flag returned by
// RecordActivityTaskHeartbeat , it should cancel the task, clean up and
// then call RespondActivityTaskCanceled . A task is considered open from
// the time that it is scheduled until it is closed. Therefore a task is
// reported as open while a worker is processing it. A task is closed after
// it has been specified in a call to RespondActivityTaskCompleted,
// RespondActivityTaskCanceled , RespondActivityTaskFailed , or the task
// has timed out . You can use IAM policies to control this action's access
// to Amazon SWF resources as follows: Use a Resource element with the
// domain name to limit the action to only specified domains. Use an Action
// element to allow or deny permission to call this action. You cannot use
// an IAM policy to constrain this action's parameters. If the caller does
// not have sufficient permissions to invoke the action, or the parameter
// values fall outside the specified constraints, the action fails by
// throwing OperationNotPermitted . For details and example IAM policies,
// see Using IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) RespondActivityTaskCompleted(req RespondActivityTaskCompletedInput) (err error) {
	// NRE
	err = c.client.Do("RespondActivityTaskCompleted", "POST", "/", req, nil)
	return
}

// RespondActivityTaskFailed used by workers to tell the service that the
// ActivityTask identified by the taskToken has failed with reason (if
// specified). The reason and details appear in the ActivityTaskFailed
// event added to the workflow history. A task is considered open from the
// time that it is scheduled until it is closed. Therefore a task is
// reported as open while a worker is processing it. A task is closed after
// it has been specified in a call to RespondActivityTaskCompleted ,
// RespondActivityTaskCanceled , RespondActivityTaskFailed, or the task has
// timed out . You can use IAM policies to control this action's access to
// Amazon SWF resources as follows: Use a Resource element with the domain
// name to limit the action to only specified domains. Use an Action
// element to allow or deny permission to call this action. You cannot use
// an IAM policy to constrain this action's parameters. If the caller does
// not have sufficient permissions to invoke the action, or the parameter
// values fall outside the specified constraints, the action fails by
// throwing OperationNotPermitted . For details and example IAM policies,
// see Using IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) RespondActivityTaskFailed(req RespondActivityTaskFailedInput) (err error) {
	// NRE
	err = c.client.Do("RespondActivityTaskFailed", "POST", "/", req, nil)
	return
}

// RespondDecisionTaskCompleted used by deciders to tell the service that
// the DecisionTask identified by the taskToken has successfully completed.
// The decisions argument specifies the list of decisions made while
// processing the task. A DecisionTaskCompleted event is added to the
// workflow history. The executionContext specified is attached to the
// event in the workflow execution history. If an IAM policy grants
// permission to use RespondDecisionTaskCompleted , it can express
// permissions for the list of decisions in the decisions parameter. Each
// of the decisions has one or more parameters, much like a regular API
// call. To allow for policies to be as readable as possible, you can
// express permissions on decisions as if they were actual API calls,
// including applying conditions to some parameters. For more information,
// see Using IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) RespondDecisionTaskCompleted(req RespondDecisionTaskCompletedInput) (err error) {
	// NRE
	err = c.client.Do("RespondDecisionTaskCompleted", "POST", "/", req, nil)
	return
}

// SignalWorkflowExecution records a WorkflowExecutionSignaled event in the
// workflow execution history and creates a decision task for the workflow
// execution identified by the given domain, workflowId and runId. The
// event is recorded with the specified user defined signalName and input
// (if provided). You can use IAM policies to control this action's access
// to Amazon SWF resources as follows: Use a Resource element with the
// domain name to limit the action to only specified domains. Use an Action
// element to allow or deny permission to call this action. You cannot use
// an IAM policy to constrain this action's parameters. If the caller does
// not have sufficient permissions to invoke the action, or the parameter
// values fall outside the specified constraints, the action fails by
// throwing OperationNotPermitted . For details and example IAM policies,
// see Using IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) SignalWorkflowExecution(req SignalWorkflowExecutionInput) (err error) {
	// NRE
	err = c.client.Do("SignalWorkflowExecution", "POST", "/", req, nil)
	return
}

// StartWorkflowExecution starts an execution of the workflow type in the
// specified domain using the provided workflowId and input data. This
// action returns the newly started workflow execution. You can use IAM
// policies to control this action's access to Amazon SWF resources as
// follows: Use a Resource element with the domain name to limit the action
// to only specified domains. Use an Action element to allow or deny
// permission to call this action. Constrain the following parameters by
// using a Condition element with the appropriate keys. tagList.member.0 :
// The key is swf:tagList.member.0 tagList.member.1 : The key is
// swf:tagList.member.1 tagList.member.2 : The key is swf:tagList.member.2
// tagList.member.3 : The key is swf:tagList.member.3 tagList.member.4 :
// The key is swf:tagList.member.4 taskList : String constraint. The key is
// swf:taskList.name name : String constraint. The key is
// swf:workflowType.name version : String constraint. The key is
// swf:workflowType.version If the caller does not have sufficient
// permissions to invoke the action, or the parameter values fall outside
// the specified constraints, the action fails by throwing
// OperationNotPermitted . For details and example IAM policies, see Using
// IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) StartWorkflowExecution(req StartWorkflowExecutionInput) (resp *Run, err error) {
	resp = &Run{}
	err = c.client.Do("StartWorkflowExecution", "POST", "/", req, resp)
	return
}

// TerminateWorkflowExecution records a WorkflowExecutionTerminated event
// and forces closure of the workflow execution identified by the given
// domain, runId, and workflowId. The child policy, registered with the
// workflow type or specified when starting this execution, is applied to
// any open child workflow executions of this workflow execution. If the
// identified workflow execution was in progress, it is terminated
// immediately. You can use IAM policies to control this action's access to
// Amazon SWF resources as follows: Use a Resource element with the domain
// name to limit the action to only specified domains. Use an Action
// element to allow or deny permission to call this action. You cannot use
// an IAM policy to constrain this action's parameters. If the caller does
// not have sufficient permissions to invoke the action, or the parameter
// values fall outside the specified constraints, the action fails by
// throwing OperationNotPermitted . For details and example IAM policies,
// see Using IAM to Manage Access to Amazon SWF Workflows
func (c *SWF) TerminateWorkflowExecution(req TerminateWorkflowExecutionInput) (err error) {
	// NRE
	err = c.client.Do("TerminateWorkflowExecution", "POST", "/", req, nil)
	return
}

// ActivityTask is undocumented.
type ActivityTask struct {
	ActivityID        string            `json:"activityId"`
	ActivityType      ActivityType      `json:"activityType"`
	Input             string            `json:"input,omitempty"`
	StartedEventID    int               `json:"startedEventId"`
	TaskToken         string            `json:"taskToken"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
}

// ActivityTaskCancelRequestedEventAttributes is undocumented.
type ActivityTaskCancelRequestedEventAttributes struct {
	ActivityID                   string `json:"activityId"`
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
}

// ActivityTaskCanceledEventAttributes is undocumented.
type ActivityTaskCanceledEventAttributes struct {
	Details                      string `json:"details,omitempty"`
	LatestCancelRequestedEventID int    `json:"latestCancelRequestedEventId,omitempty"`
	ScheduledEventID             int    `json:"scheduledEventId"`
	StartedEventID               int    `json:"startedEventId"`
}

// ActivityTaskCompletedEventAttributes is undocumented.
type ActivityTaskCompletedEventAttributes struct {
	Result           string `json:"result,omitempty"`
	ScheduledEventID int    `json:"scheduledEventId"`
	StartedEventID   int    `json:"startedEventId"`
}

// ActivityTaskFailedEventAttributes is undocumented.
type ActivityTaskFailedEventAttributes struct {
	Details          string `json:"details,omitempty"`
	Reason           string `json:"reason,omitempty"`
	ScheduledEventID int    `json:"scheduledEventId"`
	StartedEventID   int    `json:"startedEventId"`
}

// ActivityTaskScheduledEventAttributes is undocumented.
type ActivityTaskScheduledEventAttributes struct {
	ActivityID                   string       `json:"activityId"`
	ActivityType                 ActivityType `json:"activityType"`
	Control                      string       `json:"control,omitempty"`
	DecisionTaskCompletedEventID int          `json:"decisionTaskCompletedEventId"`
	HeartbeatTimeout             string       `json:"heartbeatTimeout,omitempty"`
	Input                        string       `json:"input,omitempty"`
	ScheduleToCloseTimeout       string       `json:"scheduleToCloseTimeout,omitempty"`
	ScheduleToStartTimeout       string       `json:"scheduleToStartTimeout,omitempty"`
	StartToCloseTimeout          string       `json:"startToCloseTimeout,omitempty"`
	TaskList                     TaskList     `json:"taskList"`
}

// ActivityTaskStartedEventAttributes is undocumented.
type ActivityTaskStartedEventAttributes struct {
	Identity         string `json:"identity,omitempty"`
	ScheduledEventID int    `json:"scheduledEventId"`
}

// ActivityTaskStatus is undocumented.
type ActivityTaskStatus struct {
	CancelRequested bool `json:"cancelRequested"`
}

// ActivityTaskTimedOutEventAttributes is undocumented.
type ActivityTaskTimedOutEventAttributes struct {
	Details          string `json:"details,omitempty"`
	ScheduledEventID int    `json:"scheduledEventId"`
	StartedEventID   int    `json:"startedEventId"`
	TimeoutType      string `json:"timeoutType"`
}

// ActivityType is undocumented.
type ActivityType struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ActivityTypeConfiguration is undocumented.
type ActivityTypeConfiguration struct {
	DefaultTaskHeartbeatTimeout       string   `json:"defaultTaskHeartbeatTimeout,omitempty"`
	DefaultTaskList                   TaskList `json:"defaultTaskList,omitempty"`
	DefaultTaskScheduleToCloseTimeout string   `json:"defaultTaskScheduleToCloseTimeout,omitempty"`
	DefaultTaskScheduleToStartTimeout string   `json:"defaultTaskScheduleToStartTimeout,omitempty"`
	DefaultTaskStartToCloseTimeout    string   `json:"defaultTaskStartToCloseTimeout,omitempty"`
}

// ActivityTypeDetail is undocumented.
type ActivityTypeDetail struct {
	Configuration ActivityTypeConfiguration `json:"configuration"`
	TypeInfo      ActivityTypeInfo          `json:"typeInfo"`
}

// ActivityTypeInfo is undocumented.
type ActivityTypeInfo struct {
	ActivityType    ActivityType `json:"activityType"`
	CreationDate    time.Time    `json:"creationDate"`
	DeprecationDate time.Time    `json:"deprecationDate,omitempty"`
	Description     string       `json:"description,omitempty"`
	Status          string       `json:"status"`
}

// ActivityTypeInfos is undocumented.
type ActivityTypeInfos struct {
	NextPageToken string             `json:"nextPageToken,omitempty"`
	TypeInfos     []ActivityTypeInfo `json:"typeInfos"`
}

// CancelTimerDecisionAttributes is undocumented.
type CancelTimerDecisionAttributes struct {
	TimerID string `json:"timerId"`
}

// CancelTimerFailedEventAttributes is undocumented.
type CancelTimerFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
	TimerID                      string `json:"timerId"`
}

// CancelWorkflowExecutionDecisionAttributes is undocumented.
type CancelWorkflowExecutionDecisionAttributes struct {
	Details string `json:"details,omitempty"`
}

// CancelWorkflowExecutionFailedEventAttributes is undocumented.
type CancelWorkflowExecutionFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
}

// ChildWorkflowExecutionCanceledEventAttributes is undocumented.
type ChildWorkflowExecutionCanceledEventAttributes struct {
	Details           string            `json:"details,omitempty"`
	InitiatedEventID  int               `json:"initiatedEventId"`
	StartedEventID    int               `json:"startedEventId"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
	WorkflowType      WorkflowType      `json:"workflowType"`
}

// ChildWorkflowExecutionCompletedEventAttributes is undocumented.
type ChildWorkflowExecutionCompletedEventAttributes struct {
	InitiatedEventID  int               `json:"initiatedEventId"`
	Result            string            `json:"result,omitempty"`
	StartedEventID    int               `json:"startedEventId"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
	WorkflowType      WorkflowType      `json:"workflowType"`
}

// ChildWorkflowExecutionFailedEventAttributes is undocumented.
type ChildWorkflowExecutionFailedEventAttributes struct {
	Details           string            `json:"details,omitempty"`
	InitiatedEventID  int               `json:"initiatedEventId"`
	Reason            string            `json:"reason,omitempty"`
	StartedEventID    int               `json:"startedEventId"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
	WorkflowType      WorkflowType      `json:"workflowType"`
}

// ChildWorkflowExecutionStartedEventAttributes is undocumented.
type ChildWorkflowExecutionStartedEventAttributes struct {
	InitiatedEventID  int               `json:"initiatedEventId"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
	WorkflowType      WorkflowType      `json:"workflowType"`
}

// ChildWorkflowExecutionTerminatedEventAttributes is undocumented.
type ChildWorkflowExecutionTerminatedEventAttributes struct {
	InitiatedEventID  int               `json:"initiatedEventId"`
	StartedEventID    int               `json:"startedEventId"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
	WorkflowType      WorkflowType      `json:"workflowType"`
}

// ChildWorkflowExecutionTimedOutEventAttributes is undocumented.
type ChildWorkflowExecutionTimedOutEventAttributes struct {
	InitiatedEventID  int               `json:"initiatedEventId"`
	StartedEventID    int               `json:"startedEventId"`
	TimeoutType       string            `json:"timeoutType"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
	WorkflowType      WorkflowType      `json:"workflowType"`
}

// CloseStatusFilter is undocumented.
type CloseStatusFilter struct {
	Status string `json:"status"`
}

// CompleteWorkflowExecutionDecisionAttributes is undocumented.
type CompleteWorkflowExecutionDecisionAttributes struct {
	Result string `json:"result,omitempty"`
}

// CompleteWorkflowExecutionFailedEventAttributes is undocumented.
type CompleteWorkflowExecutionFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
}

// ContinueAsNewWorkflowExecutionDecisionAttributes is undocumented.
type ContinueAsNewWorkflowExecutionDecisionAttributes struct {
	ChildPolicy                  string   `json:"childPolicy,omitempty"`
	ExecutionStartToCloseTimeout string   `json:"executionStartToCloseTimeout,omitempty"`
	Input                        string   `json:"input,omitempty"`
	TagList                      []string `json:"tagList,omitempty"`
	TaskList                     TaskList `json:"taskList,omitempty"`
	TaskStartToCloseTimeout      string   `json:"taskStartToCloseTimeout,omitempty"`
	WorkflowTypeVersion          string   `json:"workflowTypeVersion,omitempty"`
}

// ContinueAsNewWorkflowExecutionFailedEventAttributes is undocumented.
type ContinueAsNewWorkflowExecutionFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
}

// CountClosedWorkflowExecutionsInput is undocumented.
type CountClosedWorkflowExecutionsInput struct {
	CloseStatusFilter CloseStatusFilter       `json:"closeStatusFilter,omitempty"`
	CloseTimeFilter   ExecutionTimeFilter     `json:"closeTimeFilter,omitempty"`
	Domain            string                  `json:"domain"`
	ExecutionFilter   WorkflowExecutionFilter `json:"executionFilter,omitempty"`
	StartTimeFilter   ExecutionTimeFilter     `json:"startTimeFilter,omitempty"`
	TagFilter         TagFilter               `json:"tagFilter,omitempty"`
	TypeFilter        WorkflowTypeFilter      `json:"typeFilter,omitempty"`
}

// CountOpenWorkflowExecutionsInput is undocumented.
type CountOpenWorkflowExecutionsInput struct {
	Domain          string                  `json:"domain"`
	ExecutionFilter WorkflowExecutionFilter `json:"executionFilter,omitempty"`
	StartTimeFilter ExecutionTimeFilter     `json:"startTimeFilter"`
	TagFilter       TagFilter               `json:"tagFilter,omitempty"`
	TypeFilter      WorkflowTypeFilter      `json:"typeFilter,omitempty"`
}

// CountPendingActivityTasksInput is undocumented.
type CountPendingActivityTasksInput struct {
	Domain   string   `json:"domain"`
	TaskList TaskList `json:"taskList"`
}

// CountPendingDecisionTasksInput is undocumented.
type CountPendingDecisionTasksInput struct {
	Domain   string   `json:"domain"`
	TaskList TaskList `json:"taskList"`
}

// Decision is undocumented.
type Decision struct {
	CancelTimerDecisionAttributes                            CancelTimerDecisionAttributes                            `json:"cancelTimerDecisionAttributes,omitempty"`
	CancelWorkflowExecutionDecisionAttributes                CancelWorkflowExecutionDecisionAttributes                `json:"cancelWorkflowExecutionDecisionAttributes,omitempty"`
	CompleteWorkflowExecutionDecisionAttributes              CompleteWorkflowExecutionDecisionAttributes              `json:"completeWorkflowExecutionDecisionAttributes,omitempty"`
	ContinueAsNewWorkflowExecutionDecisionAttributes         ContinueAsNewWorkflowExecutionDecisionAttributes         `json:"continueAsNewWorkflowExecutionDecisionAttributes,omitempty"`
	DecisionType                                             string                                                   `json:"decisionType"`
	FailWorkflowExecutionDecisionAttributes                  FailWorkflowExecutionDecisionAttributes                  `json:"failWorkflowExecutionDecisionAttributes,omitempty"`
	RecordMarkerDecisionAttributes                           RecordMarkerDecisionAttributes                           `json:"recordMarkerDecisionAttributes,omitempty"`
	RequestCancelActivityTaskDecisionAttributes              RequestCancelActivityTaskDecisionAttributes              `json:"requestCancelActivityTaskDecisionAttributes,omitempty"`
	RequestCancelExternalWorkflowExecutionDecisionAttributes RequestCancelExternalWorkflowExecutionDecisionAttributes `json:"requestCancelExternalWorkflowExecutionDecisionAttributes,omitempty"`
	ScheduleActivityTaskDecisionAttributes                   ScheduleActivityTaskDecisionAttributes                   `json:"scheduleActivityTaskDecisionAttributes,omitempty"`
	SignalExternalWorkflowExecutionDecisionAttributes        SignalExternalWorkflowExecutionDecisionAttributes        `json:"signalExternalWorkflowExecutionDecisionAttributes,omitempty"`
	StartChildWorkflowExecutionDecisionAttributes            StartChildWorkflowExecutionDecisionAttributes            `json:"startChildWorkflowExecutionDecisionAttributes,omitempty"`
	StartTimerDecisionAttributes                             StartTimerDecisionAttributes                             `json:"startTimerDecisionAttributes,omitempty"`
}

// DecisionTask is undocumented.
type DecisionTask struct {
	Events                 []HistoryEvent    `json:"events"`
	NextPageToken          string            `json:"nextPageToken,omitempty"`
	PreviousStartedEventID int               `json:"previousStartedEventId,omitempty"`
	StartedEventID         int               `json:"startedEventId"`
	TaskToken              string            `json:"taskToken"`
	WorkflowExecution      WorkflowExecution `json:"workflowExecution"`
	WorkflowType           WorkflowType      `json:"workflowType"`
}

// DecisionTaskCompletedEventAttributes is undocumented.
type DecisionTaskCompletedEventAttributes struct {
	ExecutionContext string `json:"executionContext,omitempty"`
	ScheduledEventID int    `json:"scheduledEventId"`
	StartedEventID   int    `json:"startedEventId"`
}

// DecisionTaskScheduledEventAttributes is undocumented.
type DecisionTaskScheduledEventAttributes struct {
	StartToCloseTimeout string   `json:"startToCloseTimeout,omitempty"`
	TaskList            TaskList `json:"taskList"`
}

// DecisionTaskStartedEventAttributes is undocumented.
type DecisionTaskStartedEventAttributes struct {
	Identity         string `json:"identity,omitempty"`
	ScheduledEventID int    `json:"scheduledEventId"`
}

// DecisionTaskTimedOutEventAttributes is undocumented.
type DecisionTaskTimedOutEventAttributes struct {
	ScheduledEventID int    `json:"scheduledEventId"`
	StartedEventID   int    `json:"startedEventId"`
	TimeoutType      string `json:"timeoutType"`
}

// DeprecateActivityTypeInput is undocumented.
type DeprecateActivityTypeInput struct {
	ActivityType ActivityType `json:"activityType"`
	Domain       string       `json:"domain"`
}

// DeprecateDomainInput is undocumented.
type DeprecateDomainInput struct {
	Name string `json:"name"`
}

// DeprecateWorkflowTypeInput is undocumented.
type DeprecateWorkflowTypeInput struct {
	Domain       string       `json:"domain"`
	WorkflowType WorkflowType `json:"workflowType"`
}

// DescribeActivityTypeInput is undocumented.
type DescribeActivityTypeInput struct {
	ActivityType ActivityType `json:"activityType"`
	Domain       string       `json:"domain"`
}

// DescribeDomainInput is undocumented.
type DescribeDomainInput struct {
	Name string `json:"name"`
}

// DescribeWorkflowExecutionInput is undocumented.
type DescribeWorkflowExecutionInput struct {
	Domain    string            `json:"domain"`
	Execution WorkflowExecution `json:"execution"`
}

// DescribeWorkflowTypeInput is undocumented.
type DescribeWorkflowTypeInput struct {
	Domain       string       `json:"domain"`
	WorkflowType WorkflowType `json:"workflowType"`
}

// DomainConfiguration is undocumented.
type DomainConfiguration struct {
	WorkflowExecutionRetentionPeriodInDays string `json:"workflowExecutionRetentionPeriodInDays"`
}

// DomainDetail is undocumented.
type DomainDetail struct {
	Configuration DomainConfiguration `json:"configuration"`
	DomainInfo    DomainInfo          `json:"domainInfo"`
}

// DomainInfo is undocumented.
type DomainInfo struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	Status      string `json:"status"`
}

// DomainInfos is undocumented.
type DomainInfos struct {
	DomainInfos   []DomainInfo `json:"domainInfos"`
	NextPageToken string       `json:"nextPageToken,omitempty"`
}

// ExecutionTimeFilter is undocumented.
type ExecutionTimeFilter struct {
	LatestDate time.Time `json:"latestDate,omitempty"`
	OldestDate time.Time `json:"oldestDate"`
}

// ExternalWorkflowExecutionCancelRequestedEventAttributes is undocumented.
type ExternalWorkflowExecutionCancelRequestedEventAttributes struct {
	InitiatedEventID  int               `json:"initiatedEventId"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
}

// ExternalWorkflowExecutionSignaledEventAttributes is undocumented.
type ExternalWorkflowExecutionSignaledEventAttributes struct {
	InitiatedEventID  int               `json:"initiatedEventId"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
}

// FailWorkflowExecutionDecisionAttributes is undocumented.
type FailWorkflowExecutionDecisionAttributes struct {
	Details string `json:"details,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

// FailWorkflowExecutionFailedEventAttributes is undocumented.
type FailWorkflowExecutionFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
}

// GetWorkflowExecutionHistoryInput is undocumented.
type GetWorkflowExecutionHistoryInput struct {
	Domain          string            `json:"domain"`
	Execution       WorkflowExecution `json:"execution"`
	MaximumPageSize int               `json:"maximumPageSize,omitempty"`
	NextPageToken   string            `json:"nextPageToken,omitempty"`
	ReverseOrder    bool              `json:"reverseOrder,omitempty"`
}

// History is undocumented.
type History struct {
	Events        []HistoryEvent `json:"events"`
	NextPageToken string         `json:"nextPageToken,omitempty"`
}

// HistoryEvent is undocumented.
type HistoryEvent struct {
	ActivityTaskCancelRequestedEventAttributes                     ActivityTaskCancelRequestedEventAttributes                     `json:"activityTaskCancelRequestedEventAttributes,omitempty"`
	ActivityTaskCanceledEventAttributes                            ActivityTaskCanceledEventAttributes                            `json:"activityTaskCanceledEventAttributes,omitempty"`
	ActivityTaskCompletedEventAttributes                           ActivityTaskCompletedEventAttributes                           `json:"activityTaskCompletedEventAttributes,omitempty"`
	ActivityTaskFailedEventAttributes                              ActivityTaskFailedEventAttributes                              `json:"activityTaskFailedEventAttributes,omitempty"`
	ActivityTaskScheduledEventAttributes                           ActivityTaskScheduledEventAttributes                           `json:"activityTaskScheduledEventAttributes,omitempty"`
	ActivityTaskStartedEventAttributes                             ActivityTaskStartedEventAttributes                             `json:"activityTaskStartedEventAttributes,omitempty"`
	ActivityTaskTimedOutEventAttributes                            ActivityTaskTimedOutEventAttributes                            `json:"activityTaskTimedOutEventAttributes,omitempty"`
	CancelTimerFailedEventAttributes                               CancelTimerFailedEventAttributes                               `json:"cancelTimerFailedEventAttributes,omitempty"`
	CancelWorkflowExecutionFailedEventAttributes                   CancelWorkflowExecutionFailedEventAttributes                   `json:"cancelWorkflowExecutionFailedEventAttributes,omitempty"`
	ChildWorkflowExecutionCanceledEventAttributes                  ChildWorkflowExecutionCanceledEventAttributes                  `json:"childWorkflowExecutionCanceledEventAttributes,omitempty"`
	ChildWorkflowExecutionCompletedEventAttributes                 ChildWorkflowExecutionCompletedEventAttributes                 `json:"childWorkflowExecutionCompletedEventAttributes,omitempty"`
	ChildWorkflowExecutionFailedEventAttributes                    ChildWorkflowExecutionFailedEventAttributes                    `json:"childWorkflowExecutionFailedEventAttributes,omitempty"`
	ChildWorkflowExecutionStartedEventAttributes                   ChildWorkflowExecutionStartedEventAttributes                   `json:"childWorkflowExecutionStartedEventAttributes,omitempty"`
	ChildWorkflowExecutionTerminatedEventAttributes                ChildWorkflowExecutionTerminatedEventAttributes                `json:"childWorkflowExecutionTerminatedEventAttributes,omitempty"`
	ChildWorkflowExecutionTimedOutEventAttributes                  ChildWorkflowExecutionTimedOutEventAttributes                  `json:"childWorkflowExecutionTimedOutEventAttributes,omitempty"`
	CompleteWorkflowExecutionFailedEventAttributes                 CompleteWorkflowExecutionFailedEventAttributes                 `json:"completeWorkflowExecutionFailedEventAttributes,omitempty"`
	ContinueAsNewWorkflowExecutionFailedEventAttributes            ContinueAsNewWorkflowExecutionFailedEventAttributes            `json:"continueAsNewWorkflowExecutionFailedEventAttributes,omitempty"`
	DecisionTaskCompletedEventAttributes                           DecisionTaskCompletedEventAttributes                           `json:"decisionTaskCompletedEventAttributes,omitempty"`
	DecisionTaskScheduledEventAttributes                           DecisionTaskScheduledEventAttributes                           `json:"decisionTaskScheduledEventAttributes,omitempty"`
	DecisionTaskStartedEventAttributes                             DecisionTaskStartedEventAttributes                             `json:"decisionTaskStartedEventAttributes,omitempty"`
	DecisionTaskTimedOutEventAttributes                            DecisionTaskTimedOutEventAttributes                            `json:"decisionTaskTimedOutEventAttributes,omitempty"`
	EventID                                                        int                                                            `json:"eventId"`
	EventTimestamp                                                 time.Time                                                      `json:"eventTimestamp"`
	EventType                                                      string                                                         `json:"eventType"`
	ExternalWorkflowExecutionCancelRequestedEventAttributes        ExternalWorkflowExecutionCancelRequestedEventAttributes        `json:"externalWorkflowExecutionCancelRequestedEventAttributes,omitempty"`
	ExternalWorkflowExecutionSignaledEventAttributes               ExternalWorkflowExecutionSignaledEventAttributes               `json:"externalWorkflowExecutionSignaledEventAttributes,omitempty"`
	FailWorkflowExecutionFailedEventAttributes                     FailWorkflowExecutionFailedEventAttributes                     `json:"failWorkflowExecutionFailedEventAttributes,omitempty"`
	MarkerRecordedEventAttributes                                  MarkerRecordedEventAttributes                                  `json:"markerRecordedEventAttributes,omitempty"`
	RecordMarkerFailedEventAttributes                              RecordMarkerFailedEventAttributes                              `json:"recordMarkerFailedEventAttributes,omitempty"`
	RequestCancelActivityTaskFailedEventAttributes                 RequestCancelActivityTaskFailedEventAttributes                 `json:"requestCancelActivityTaskFailedEventAttributes,omitempty"`
	RequestCancelExternalWorkflowExecutionFailedEventAttributes    RequestCancelExternalWorkflowExecutionFailedEventAttributes    `json:"requestCancelExternalWorkflowExecutionFailedEventAttributes,omitempty"`
	RequestCancelExternalWorkflowExecutionInitiatedEventAttributes RequestCancelExternalWorkflowExecutionInitiatedEventAttributes `json:"requestCancelExternalWorkflowExecutionInitiatedEventAttributes,omitempty"`
	ScheduleActivityTaskFailedEventAttributes                      ScheduleActivityTaskFailedEventAttributes                      `json:"scheduleActivityTaskFailedEventAttributes,omitempty"`
	SignalExternalWorkflowExecutionFailedEventAttributes           SignalExternalWorkflowExecutionFailedEventAttributes           `json:"signalExternalWorkflowExecutionFailedEventAttributes,omitempty"`
	SignalExternalWorkflowExecutionInitiatedEventAttributes        SignalExternalWorkflowExecutionInitiatedEventAttributes        `json:"signalExternalWorkflowExecutionInitiatedEventAttributes,omitempty"`
	StartChildWorkflowExecutionFailedEventAttributes               StartChildWorkflowExecutionFailedEventAttributes               `json:"startChildWorkflowExecutionFailedEventAttributes,omitempty"`
	StartChildWorkflowExecutionInitiatedEventAttributes            StartChildWorkflowExecutionInitiatedEventAttributes            `json:"startChildWorkflowExecutionInitiatedEventAttributes,omitempty"`
	StartTimerFailedEventAttributes                                StartTimerFailedEventAttributes                                `json:"startTimerFailedEventAttributes,omitempty"`
	TimerCanceledEventAttributes                                   TimerCanceledEventAttributes                                   `json:"timerCanceledEventAttributes,omitempty"`
	TimerFiredEventAttributes                                      TimerFiredEventAttributes                                      `json:"timerFiredEventAttributes,omitempty"`
	TimerStartedEventAttributes                                    TimerStartedEventAttributes                                    `json:"timerStartedEventAttributes,omitempty"`
	WorkflowExecutionCancelRequestedEventAttributes                WorkflowExecutionCancelRequestedEventAttributes                `json:"workflowExecutionCancelRequestedEventAttributes,omitempty"`
	WorkflowExecutionCanceledEventAttributes                       WorkflowExecutionCanceledEventAttributes                       `json:"workflowExecutionCanceledEventAttributes,omitempty"`
	WorkflowExecutionCompletedEventAttributes                      WorkflowExecutionCompletedEventAttributes                      `json:"workflowExecutionCompletedEventAttributes,omitempty"`
	WorkflowExecutionContinuedAsNewEventAttributes                 WorkflowExecutionContinuedAsNewEventAttributes                 `json:"workflowExecutionContinuedAsNewEventAttributes,omitempty"`
	WorkflowExecutionFailedEventAttributes                         WorkflowExecutionFailedEventAttributes                         `json:"workflowExecutionFailedEventAttributes,omitempty"`
	WorkflowExecutionSignaledEventAttributes                       WorkflowExecutionSignaledEventAttributes                       `json:"workflowExecutionSignaledEventAttributes,omitempty"`
	WorkflowExecutionStartedEventAttributes                        WorkflowExecutionStartedEventAttributes                        `json:"workflowExecutionStartedEventAttributes,omitempty"`
	WorkflowExecutionTerminatedEventAttributes                     WorkflowExecutionTerminatedEventAttributes                     `json:"workflowExecutionTerminatedEventAttributes,omitempty"`
	WorkflowExecutionTimedOutEventAttributes                       WorkflowExecutionTimedOutEventAttributes                       `json:"workflowExecutionTimedOutEventAttributes,omitempty"`
}

// ListActivityTypesInput is undocumented.
type ListActivityTypesInput struct {
	Domain             string `json:"domain"`
	MaximumPageSize    int    `json:"maximumPageSize,omitempty"`
	Name               string `json:"name,omitempty"`
	NextPageToken      string `json:"nextPageToken,omitempty"`
	RegistrationStatus string `json:"registrationStatus"`
	ReverseOrder       bool   `json:"reverseOrder,omitempty"`
}

// ListClosedWorkflowExecutionsInput is undocumented.
type ListClosedWorkflowExecutionsInput struct {
	CloseStatusFilter CloseStatusFilter       `json:"closeStatusFilter,omitempty"`
	CloseTimeFilter   ExecutionTimeFilter     `json:"closeTimeFilter,omitempty"`
	Domain            string                  `json:"domain"`
	ExecutionFilter   WorkflowExecutionFilter `json:"executionFilter,omitempty"`
	MaximumPageSize   int                     `json:"maximumPageSize,omitempty"`
	NextPageToken     string                  `json:"nextPageToken,omitempty"`
	ReverseOrder      bool                    `json:"reverseOrder,omitempty"`
	StartTimeFilter   ExecutionTimeFilter     `json:"startTimeFilter,omitempty"`
	TagFilter         TagFilter               `json:"tagFilter,omitempty"`
	TypeFilter        WorkflowTypeFilter      `json:"typeFilter,omitempty"`
}

// ListDomainsInput is undocumented.
type ListDomainsInput struct {
	MaximumPageSize    int    `json:"maximumPageSize,omitempty"`
	NextPageToken      string `json:"nextPageToken,omitempty"`
	RegistrationStatus string `json:"registrationStatus"`
	ReverseOrder       bool   `json:"reverseOrder,omitempty"`
}

// ListOpenWorkflowExecutionsInput is undocumented.
type ListOpenWorkflowExecutionsInput struct {
	Domain          string                  `json:"domain"`
	ExecutionFilter WorkflowExecutionFilter `json:"executionFilter,omitempty"`
	MaximumPageSize int                     `json:"maximumPageSize,omitempty"`
	NextPageToken   string                  `json:"nextPageToken,omitempty"`
	ReverseOrder    bool                    `json:"reverseOrder,omitempty"`
	StartTimeFilter ExecutionTimeFilter     `json:"startTimeFilter"`
	TagFilter       TagFilter               `json:"tagFilter,omitempty"`
	TypeFilter      WorkflowTypeFilter      `json:"typeFilter,omitempty"`
}

// ListWorkflowTypesInput is undocumented.
type ListWorkflowTypesInput struct {
	Domain             string `json:"domain"`
	MaximumPageSize    int    `json:"maximumPageSize,omitempty"`
	Name               string `json:"name,omitempty"`
	NextPageToken      string `json:"nextPageToken,omitempty"`
	RegistrationStatus string `json:"registrationStatus"`
	ReverseOrder       bool   `json:"reverseOrder,omitempty"`
}

// MarkerRecordedEventAttributes is undocumented.
type MarkerRecordedEventAttributes struct {
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
	Details                      string `json:"details,omitempty"`
	MarkerName                   string `json:"markerName"`
}

// PendingTaskCount is undocumented.
type PendingTaskCount struct {
	Count     int  `json:"count"`
	Truncated bool `json:"truncated,omitempty"`
}

// PollForActivityTaskInput is undocumented.
type PollForActivityTaskInput struct {
	Domain   string   `json:"domain"`
	Identity string   `json:"identity,omitempty"`
	TaskList TaskList `json:"taskList"`
}

// PollForDecisionTaskInput is undocumented.
type PollForDecisionTaskInput struct {
	Domain          string   `json:"domain"`
	Identity        string   `json:"identity,omitempty"`
	MaximumPageSize int      `json:"maximumPageSize,omitempty"`
	NextPageToken   string   `json:"nextPageToken,omitempty"`
	ReverseOrder    bool     `json:"reverseOrder,omitempty"`
	TaskList        TaskList `json:"taskList"`
}

// RecordActivityTaskHeartbeatInput is undocumented.
type RecordActivityTaskHeartbeatInput struct {
	Details   string `json:"details,omitempty"`
	TaskToken string `json:"taskToken"`
}

// RecordMarkerDecisionAttributes is undocumented.
type RecordMarkerDecisionAttributes struct {
	Details    string `json:"details,omitempty"`
	MarkerName string `json:"markerName"`
}

// RecordMarkerFailedEventAttributes is undocumented.
type RecordMarkerFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
	MarkerName                   string `json:"markerName"`
}

// RegisterActivityTypeInput is undocumented.
type RegisterActivityTypeInput struct {
	DefaultTaskHeartbeatTimeout       string   `json:"defaultTaskHeartbeatTimeout,omitempty"`
	DefaultTaskList                   TaskList `json:"defaultTaskList,omitempty"`
	DefaultTaskScheduleToCloseTimeout string   `json:"defaultTaskScheduleToCloseTimeout,omitempty"`
	DefaultTaskScheduleToStartTimeout string   `json:"defaultTaskScheduleToStartTimeout,omitempty"`
	DefaultTaskStartToCloseTimeout    string   `json:"defaultTaskStartToCloseTimeout,omitempty"`
	Description                       string   `json:"description,omitempty"`
	Domain                            string   `json:"domain"`
	Name                              string   `json:"name"`
	Version                           string   `json:"version"`
}

// RegisterDomainInput is undocumented.
type RegisterDomainInput struct {
	Description                            string `json:"description,omitempty"`
	Name                                   string `json:"name"`
	WorkflowExecutionRetentionPeriodInDays string `json:"workflowExecutionRetentionPeriodInDays"`
}

// RegisterWorkflowTypeInput is undocumented.
type RegisterWorkflowTypeInput struct {
	DefaultChildPolicy                  string   `json:"defaultChildPolicy,omitempty"`
	DefaultExecutionStartToCloseTimeout string   `json:"defaultExecutionStartToCloseTimeout,omitempty"`
	DefaultTaskList                     TaskList `json:"defaultTaskList,omitempty"`
	DefaultTaskStartToCloseTimeout      string   `json:"defaultTaskStartToCloseTimeout,omitempty"`
	Description                         string   `json:"description,omitempty"`
	Domain                              string   `json:"domain"`
	Name                                string   `json:"name"`
	Version                             string   `json:"version"`
}

// RequestCancelActivityTaskDecisionAttributes is undocumented.
type RequestCancelActivityTaskDecisionAttributes struct {
	ActivityID string `json:"activityId"`
}

// RequestCancelActivityTaskFailedEventAttributes is undocumented.
type RequestCancelActivityTaskFailedEventAttributes struct {
	ActivityID                   string `json:"activityId"`
	Cause                        string `json:"cause"`
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
}

// RequestCancelExternalWorkflowExecutionDecisionAttributes is undocumented.
type RequestCancelExternalWorkflowExecutionDecisionAttributes struct {
	Control    string `json:"control,omitempty"`
	RunID      string `json:"runId,omitempty"`
	WorkflowID string `json:"workflowId"`
}

// RequestCancelExternalWorkflowExecutionFailedEventAttributes is undocumented.
type RequestCancelExternalWorkflowExecutionFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	Control                      string `json:"control,omitempty"`
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
	InitiatedEventID             int    `json:"initiatedEventId"`
	RunID                        string `json:"runId,omitempty"`
	WorkflowID                   string `json:"workflowId"`
}

// RequestCancelExternalWorkflowExecutionInitiatedEventAttributes is undocumented.
type RequestCancelExternalWorkflowExecutionInitiatedEventAttributes struct {
	Control                      string `json:"control,omitempty"`
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
	RunID                        string `json:"runId,omitempty"`
	WorkflowID                   string `json:"workflowId"`
}

// RequestCancelWorkflowExecutionInput is undocumented.
type RequestCancelWorkflowExecutionInput struct {
	Domain     string `json:"domain"`
	RunID      string `json:"runId,omitempty"`
	WorkflowID string `json:"workflowId"`
}

// RespondActivityTaskCanceledInput is undocumented.
type RespondActivityTaskCanceledInput struct {
	Details   string `json:"details,omitempty"`
	TaskToken string `json:"taskToken"`
}

// RespondActivityTaskCompletedInput is undocumented.
type RespondActivityTaskCompletedInput struct {
	Result    string `json:"result,omitempty"`
	TaskToken string `json:"taskToken"`
}

// RespondActivityTaskFailedInput is undocumented.
type RespondActivityTaskFailedInput struct {
	Details   string `json:"details,omitempty"`
	Reason    string `json:"reason,omitempty"`
	TaskToken string `json:"taskToken"`
}

// RespondDecisionTaskCompletedInput is undocumented.
type RespondDecisionTaskCompletedInput struct {
	Decisions        []Decision `json:"decisions,omitempty"`
	ExecutionContext string     `json:"executionContext,omitempty"`
	TaskToken        string     `json:"taskToken"`
}

// Run is undocumented.
type Run struct {
	RunID string `json:"runId,omitempty"`
}

// ScheduleActivityTaskDecisionAttributes is undocumented.
type ScheduleActivityTaskDecisionAttributes struct {
	ActivityID             string       `json:"activityId"`
	ActivityType           ActivityType `json:"activityType"`
	Control                string       `json:"control,omitempty"`
	HeartbeatTimeout       string       `json:"heartbeatTimeout,omitempty"`
	Input                  string       `json:"input,omitempty"`
	ScheduleToCloseTimeout string       `json:"scheduleToCloseTimeout,omitempty"`
	ScheduleToStartTimeout string       `json:"scheduleToStartTimeout,omitempty"`
	StartToCloseTimeout    string       `json:"startToCloseTimeout,omitempty"`
	TaskList               TaskList     `json:"taskList,omitempty"`
}

// ScheduleActivityTaskFailedEventAttributes is undocumented.
type ScheduleActivityTaskFailedEventAttributes struct {
	ActivityID                   string       `json:"activityId"`
	ActivityType                 ActivityType `json:"activityType"`
	Cause                        string       `json:"cause"`
	DecisionTaskCompletedEventID int          `json:"decisionTaskCompletedEventId"`
}

// SignalExternalWorkflowExecutionDecisionAttributes is undocumented.
type SignalExternalWorkflowExecutionDecisionAttributes struct {
	Control    string `json:"control,omitempty"`
	Input      string `json:"input,omitempty"`
	RunID      string `json:"runId,omitempty"`
	SignalName string `json:"signalName"`
	WorkflowID string `json:"workflowId"`
}

// SignalExternalWorkflowExecutionFailedEventAttributes is undocumented.
type SignalExternalWorkflowExecutionFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	Control                      string `json:"control,omitempty"`
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
	InitiatedEventID             int    `json:"initiatedEventId"`
	RunID                        string `json:"runId,omitempty"`
	WorkflowID                   string `json:"workflowId"`
}

// SignalExternalWorkflowExecutionInitiatedEventAttributes is undocumented.
type SignalExternalWorkflowExecutionInitiatedEventAttributes struct {
	Control                      string `json:"control,omitempty"`
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
	Input                        string `json:"input,omitempty"`
	RunID                        string `json:"runId,omitempty"`
	SignalName                   string `json:"signalName"`
	WorkflowID                   string `json:"workflowId"`
}

// SignalWorkflowExecutionInput is undocumented.
type SignalWorkflowExecutionInput struct {
	Domain     string `json:"domain"`
	Input      string `json:"input,omitempty"`
	RunID      string `json:"runId,omitempty"`
	SignalName string `json:"signalName"`
	WorkflowID string `json:"workflowId"`
}

// StartChildWorkflowExecutionDecisionAttributes is undocumented.
type StartChildWorkflowExecutionDecisionAttributes struct {
	ChildPolicy                  string       `json:"childPolicy,omitempty"`
	Control                      string       `json:"control,omitempty"`
	ExecutionStartToCloseTimeout string       `json:"executionStartToCloseTimeout,omitempty"`
	Input                        string       `json:"input,omitempty"`
	TagList                      []string     `json:"tagList,omitempty"`
	TaskList                     TaskList     `json:"taskList,omitempty"`
	TaskStartToCloseTimeout      string       `json:"taskStartToCloseTimeout,omitempty"`
	WorkflowID                   string       `json:"workflowId"`
	WorkflowType                 WorkflowType `json:"workflowType"`
}

// StartChildWorkflowExecutionFailedEventAttributes is undocumented.
type StartChildWorkflowExecutionFailedEventAttributes struct {
	Cause                        string       `json:"cause"`
	Control                      string       `json:"control,omitempty"`
	DecisionTaskCompletedEventID int          `json:"decisionTaskCompletedEventId"`
	InitiatedEventID             int          `json:"initiatedEventId"`
	WorkflowID                   string       `json:"workflowId"`
	WorkflowType                 WorkflowType `json:"workflowType"`
}

// StartChildWorkflowExecutionInitiatedEventAttributes is undocumented.
type StartChildWorkflowExecutionInitiatedEventAttributes struct {
	ChildPolicy                  string       `json:"childPolicy"`
	Control                      string       `json:"control,omitempty"`
	DecisionTaskCompletedEventID int          `json:"decisionTaskCompletedEventId"`
	ExecutionStartToCloseTimeout string       `json:"executionStartToCloseTimeout,omitempty"`
	Input                        string       `json:"input,omitempty"`
	TagList                      []string     `json:"tagList,omitempty"`
	TaskList                     TaskList     `json:"taskList"`
	TaskStartToCloseTimeout      string       `json:"taskStartToCloseTimeout,omitempty"`
	WorkflowID                   string       `json:"workflowId"`
	WorkflowType                 WorkflowType `json:"workflowType"`
}

// StartTimerDecisionAttributes is undocumented.
type StartTimerDecisionAttributes struct {
	Control            string `json:"control,omitempty"`
	StartToFireTimeout string `json:"startToFireTimeout"`
	TimerID            string `json:"timerId"`
}

// StartTimerFailedEventAttributes is undocumented.
type StartTimerFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
	TimerID                      string `json:"timerId"`
}

// StartWorkflowExecutionInput is undocumented.
type StartWorkflowExecutionInput struct {
	ChildPolicy                  string       `json:"childPolicy,omitempty"`
	Domain                       string       `json:"domain"`
	ExecutionStartToCloseTimeout string       `json:"executionStartToCloseTimeout,omitempty"`
	Input                        string       `json:"input,omitempty"`
	TagList                      []string     `json:"tagList,omitempty"`
	TaskList                     TaskList     `json:"taskList,omitempty"`
	TaskStartToCloseTimeout      string       `json:"taskStartToCloseTimeout,omitempty"`
	WorkflowID                   string       `json:"workflowId"`
	WorkflowType                 WorkflowType `json:"workflowType"`
}

// TagFilter is undocumented.
type TagFilter struct {
	Tag string `json:"tag"`
}

// TaskList is undocumented.
type TaskList struct {
	Name string `json:"name"`
}

// TerminateWorkflowExecutionInput is undocumented.
type TerminateWorkflowExecutionInput struct {
	ChildPolicy string `json:"childPolicy,omitempty"`
	Details     string `json:"details,omitempty"`
	Domain      string `json:"domain"`
	Reason      string `json:"reason,omitempty"`
	RunID       string `json:"runId,omitempty"`
	WorkflowID  string `json:"workflowId"`
}

// TimerCanceledEventAttributes is undocumented.
type TimerCanceledEventAttributes struct {
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
	StartedEventID               int    `json:"startedEventId"`
	TimerID                      string `json:"timerId"`
}

// TimerFiredEventAttributes is undocumented.
type TimerFiredEventAttributes struct {
	StartedEventID int    `json:"startedEventId"`
	TimerID        string `json:"timerId"`
}

// TimerStartedEventAttributes is undocumented.
type TimerStartedEventAttributes struct {
	Control                      string `json:"control,omitempty"`
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
	StartToFireTimeout           string `json:"startToFireTimeout"`
	TimerID                      string `json:"timerId"`
}

// WorkflowExecution is undocumented.
type WorkflowExecution struct {
	RunID      string `json:"runId"`
	WorkflowID string `json:"workflowId"`
}

// WorkflowExecutionCancelRequestedEventAttributes is undocumented.
type WorkflowExecutionCancelRequestedEventAttributes struct {
	Cause                     string            `json:"cause,omitempty"`
	ExternalInitiatedEventID  int               `json:"externalInitiatedEventId,omitempty"`
	ExternalWorkflowExecution WorkflowExecution `json:"externalWorkflowExecution,omitempty"`
}

// WorkflowExecutionCanceledEventAttributes is undocumented.
type WorkflowExecutionCanceledEventAttributes struct {
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
	Details                      string `json:"details,omitempty"`
}

// WorkflowExecutionCompletedEventAttributes is undocumented.
type WorkflowExecutionCompletedEventAttributes struct {
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
	Result                       string `json:"result,omitempty"`
}

// WorkflowExecutionConfiguration is undocumented.
type WorkflowExecutionConfiguration struct {
	ChildPolicy                  string   `json:"childPolicy"`
	ExecutionStartToCloseTimeout string   `json:"executionStartToCloseTimeout"`
	TaskList                     TaskList `json:"taskList"`
	TaskStartToCloseTimeout      string   `json:"taskStartToCloseTimeout"`
}

// WorkflowExecutionContinuedAsNewEventAttributes is undocumented.
type WorkflowExecutionContinuedAsNewEventAttributes struct {
	ChildPolicy                  string       `json:"childPolicy"`
	DecisionTaskCompletedEventID int          `json:"decisionTaskCompletedEventId"`
	ExecutionStartToCloseTimeout string       `json:"executionStartToCloseTimeout,omitempty"`
	Input                        string       `json:"input,omitempty"`
	NewExecutionRunID            string       `json:"newExecutionRunId"`
	TagList                      []string     `json:"tagList,omitempty"`
	TaskList                     TaskList     `json:"taskList"`
	TaskStartToCloseTimeout      string       `json:"taskStartToCloseTimeout,omitempty"`
	WorkflowType                 WorkflowType `json:"workflowType"`
}

// WorkflowExecutionCount is undocumented.
type WorkflowExecutionCount struct {
	Count     int  `json:"count"`
	Truncated bool `json:"truncated,omitempty"`
}

// WorkflowExecutionDetail is undocumented.
type WorkflowExecutionDetail struct {
	ExecutionConfiguration      WorkflowExecutionConfiguration `json:"executionConfiguration"`
	ExecutionInfo               WorkflowExecutionInfo          `json:"executionInfo"`
	LatestActivityTaskTimestamp time.Time                      `json:"latestActivityTaskTimestamp,omitempty"`
	LatestExecutionContext      string                         `json:"latestExecutionContext,omitempty"`
	OpenCounts                  WorkflowExecutionOpenCounts    `json:"openCounts"`
}

// WorkflowExecutionFailedEventAttributes is undocumented.
type WorkflowExecutionFailedEventAttributes struct {
	DecisionTaskCompletedEventID int    `json:"decisionTaskCompletedEventId"`
	Details                      string `json:"details,omitempty"`
	Reason                       string `json:"reason,omitempty"`
}

// WorkflowExecutionFilter is undocumented.
type WorkflowExecutionFilter struct {
	WorkflowID string `json:"workflowId"`
}

// WorkflowExecutionInfo is undocumented.
type WorkflowExecutionInfo struct {
	CancelRequested bool              `json:"cancelRequested,omitempty"`
	CloseStatus     string            `json:"closeStatus,omitempty"`
	CloseTimestamp  time.Time         `json:"closeTimestamp,omitempty"`
	Execution       WorkflowExecution `json:"execution"`
	ExecutionStatus string            `json:"executionStatus"`
	Parent          WorkflowExecution `json:"parent,omitempty"`
	StartTimestamp  time.Time         `json:"startTimestamp"`
	TagList         []string          `json:"tagList,omitempty"`
	WorkflowType    WorkflowType      `json:"workflowType"`
}

// WorkflowExecutionInfos is undocumented.
type WorkflowExecutionInfos struct {
	ExecutionInfos []WorkflowExecutionInfo `json:"executionInfos"`
	NextPageToken  string                  `json:"nextPageToken,omitempty"`
}

// WorkflowExecutionOpenCounts is undocumented.
type WorkflowExecutionOpenCounts struct {
	OpenActivityTasks           int `json:"openActivityTasks"`
	OpenChildWorkflowExecutions int `json:"openChildWorkflowExecutions"`
	OpenDecisionTasks           int `json:"openDecisionTasks"`
	OpenTimers                  int `json:"openTimers"`
}

// WorkflowExecutionSignaledEventAttributes is undocumented.
type WorkflowExecutionSignaledEventAttributes struct {
	ExternalInitiatedEventID  int               `json:"externalInitiatedEventId,omitempty"`
	ExternalWorkflowExecution WorkflowExecution `json:"externalWorkflowExecution,omitempty"`
	Input                     string            `json:"input,omitempty"`
	SignalName                string            `json:"signalName"`
}

// WorkflowExecutionStartedEventAttributes is undocumented.
type WorkflowExecutionStartedEventAttributes struct {
	ChildPolicy                  string            `json:"childPolicy"`
	ContinuedExecutionRunID      string            `json:"continuedExecutionRunId,omitempty"`
	ExecutionStartToCloseTimeout string            `json:"executionStartToCloseTimeout,omitempty"`
	Input                        string            `json:"input,omitempty"`
	ParentInitiatedEventID       int               `json:"parentInitiatedEventId,omitempty"`
	ParentWorkflowExecution      WorkflowExecution `json:"parentWorkflowExecution,omitempty"`
	TagList                      []string          `json:"tagList,omitempty"`
	TaskList                     TaskList          `json:"taskList"`
	TaskStartToCloseTimeout      string            `json:"taskStartToCloseTimeout,omitempty"`
	WorkflowType                 WorkflowType      `json:"workflowType"`
}

// WorkflowExecutionTerminatedEventAttributes is undocumented.
type WorkflowExecutionTerminatedEventAttributes struct {
	Cause       string `json:"cause,omitempty"`
	ChildPolicy string `json:"childPolicy"`
	Details     string `json:"details,omitempty"`
	Reason      string `json:"reason,omitempty"`
}

// WorkflowExecutionTimedOutEventAttributes is undocumented.
type WorkflowExecutionTimedOutEventAttributes struct {
	ChildPolicy string `json:"childPolicy"`
	TimeoutType string `json:"timeoutType"`
}

// WorkflowType is undocumented.
type WorkflowType struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// WorkflowTypeConfiguration is undocumented.
type WorkflowTypeConfiguration struct {
	DefaultChildPolicy                  string   `json:"defaultChildPolicy,omitempty"`
	DefaultExecutionStartToCloseTimeout string   `json:"defaultExecutionStartToCloseTimeout,omitempty"`
	DefaultTaskList                     TaskList `json:"defaultTaskList,omitempty"`
	DefaultTaskStartToCloseTimeout      string   `json:"defaultTaskStartToCloseTimeout,omitempty"`
}

// WorkflowTypeDetail is undocumented.
type WorkflowTypeDetail struct {
	Configuration WorkflowTypeConfiguration `json:"configuration"`
	TypeInfo      WorkflowTypeInfo          `json:"typeInfo"`
}

// WorkflowTypeFilter is undocumented.
type WorkflowTypeFilter struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
}

// WorkflowTypeInfo is undocumented.
type WorkflowTypeInfo struct {
	CreationDate    time.Time    `json:"creationDate"`
	DeprecationDate time.Time    `json:"deprecationDate,omitempty"`
	Description     string       `json:"description,omitempty"`
	Status          string       `json:"status"`
	WorkflowType    WorkflowType `json:"workflowType"`
}

// WorkflowTypeInfos is undocumented.
type WorkflowTypeInfos struct {
	NextPageToken string             `json:"nextPageToken,omitempty"`
	TypeInfos     []WorkflowTypeInfo `json:"typeInfos"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
