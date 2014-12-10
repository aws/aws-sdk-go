// Package emr provides a client for Amazon Elastic MapReduce.
package emr

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// EMR is a client for Amazon Elastic MapReduce.
type EMR struct {
	client *aws.JSONClient
}

// New returns a new EMR client.
func New(creds aws.Credentials, region string, client *http.Client) *EMR {
	if client == nil {
		client = http.DefaultClient
	}

	service := "elasticmapreduce"
	endpoint, service, region := endpoints.Lookup("elasticmapreduce", region)

	return &EMR{
		client: &aws.JSONClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			}, Client: client,
			Endpoint:     endpoint,
			JSONVersion:  "1.1",
			TargetPrefix: "ElasticMapReduce",
		},
	}
}

// AddInstanceGroups addInstanceGroups adds an instance group to a running
// cluster.
func (c *EMR) AddInstanceGroups(req AddInstanceGroupsInput) (resp *AddInstanceGroupsOutput, err error) {
	resp = &AddInstanceGroupsOutput{}
	err = c.client.Do("AddInstanceGroups", "POST", "/", req, resp)
	return
}

// AddJobFlowSteps addJobFlowSteps adds new steps to a running job flow. A
// maximum of 256 steps are allowed in each job flow. If your job flow is
// long-running (such as a Hive data warehouse) or complex, you may require
// more than 256 steps to process your data. You can bypass the 256-step
// limitation in various ways, including using the SSH shell to connect to
// the master node and submitting queries directly to the software running
// on the master node, such as Hive and Hadoop. For more information on how
// to do this, go to Add More than 256 Steps to a Job Flow in the Amazon
// Elastic MapReduce Developer's Guide A step specifies the location of a
// JAR file stored either on the master node of the job flow or in Amazon
// S3. Each step is performed by the main function of the main class of the
// JAR file. The main class can be specified either in the manifest of the
// JAR or by using the MainFunction parameter of the step. Elastic
// MapReduce executes each step in the order listed. For a step to be
// considered complete, the main function must exit with a zero exit code
// and all Hadoop jobs started while the step was running must have
// completed and run successfully. You can only add steps to a job flow
// that is in one of the following states: or
func (c *EMR) AddJobFlowSteps(req AddJobFlowStepsInput) (resp *AddJobFlowStepsOutput, err error) {
	resp = &AddJobFlowStepsOutput{}
	err = c.client.Do("AddJobFlowSteps", "POST", "/", req, resp)
	return
}

// AddTags adds tags to an Amazon EMR resource. Tags make it easier to
// associate clusters in various ways, such as grouping clusters to track
// your Amazon EMR resource allocation costs. For more information, see
// Tagging Amazon EMR Resources .
func (c *EMR) AddTags(req AddTagsInput) (resp *AddTagsOutput, err error) {
	resp = &AddTagsOutput{}
	err = c.client.Do("AddTags", "POST", "/", req, resp)
	return
}

// DescribeCluster provides cluster-level details including status,
// hardware and software configuration, VPC settings, and so on. For
// information about the cluster steps, see ListSteps
func (c *EMR) DescribeCluster(req DescribeClusterInput) (resp *DescribeClusterOutput, err error) {
	resp = &DescribeClusterOutput{}
	err = c.client.Do("DescribeCluster", "POST", "/", req, resp)
	return
}

// DescribeJobFlows this API is deprecated and will eventually be removed.
// We recommend you use ListClusters , DescribeCluster , ListSteps ,
// ListInstanceGroups and ListBootstrapActions instead. DescribeJobFlows
// returns a list of job flows that match all of the supplied parameters.
// The parameters can include a list of job flow IDs, job flow states, and
// restrictions on job flow creation date and time. Regardless of supplied
// parameters, only job flows created within the last two months are
// returned. If no parameters are supplied, then job flows matching either
// of the following criteria are returned: Job flows created and completed
// in the last two weeks Job flows created within the last two months that
// are in one of the following states: , , , Amazon Elastic MapReduce can
// return a maximum of 512 job flow descriptions.
func (c *EMR) DescribeJobFlows(req DescribeJobFlowsInput) (resp *DescribeJobFlowsOutput, err error) {
	resp = &DescribeJobFlowsOutput{}
	err = c.client.Do("DescribeJobFlows", "POST", "/", req, resp)
	return
}

// DescribeStep is undocumented.
func (c *EMR) DescribeStep(req DescribeStepInput) (resp *DescribeStepOutput, err error) {
	resp = &DescribeStepOutput{}
	err = c.client.Do("DescribeStep", "POST", "/", req, resp)
	return
}

// ListBootstrapActions provides information about the bootstrap actions
// associated with a cluster.
func (c *EMR) ListBootstrapActions(req ListBootstrapActionsInput) (resp *ListBootstrapActionsOutput, err error) {
	resp = &ListBootstrapActionsOutput{}
	err = c.client.Do("ListBootstrapActions", "POST", "/", req, resp)
	return
}

// ListClusters provides the status of all clusters visible to this AWS
// account. Allows you to filter the list of clusters based on certain
// criteria; for example, filtering by cluster creation date and time or by
// status. This call returns a maximum of 50 clusters per call, but returns
// a marker to track the paging of the cluster list across multiple
// ListClusters calls.
func (c *EMR) ListClusters(req ListClustersInput) (resp *ListClustersOutput, err error) {
	resp = &ListClustersOutput{}
	err = c.client.Do("ListClusters", "POST", "/", req, resp)
	return
}

// ListInstanceGroups provides all available details about the instance
// groups in a cluster.
func (c *EMR) ListInstanceGroups(req ListInstanceGroupsInput) (resp *ListInstanceGroupsOutput, err error) {
	resp = &ListInstanceGroupsOutput{}
	err = c.client.Do("ListInstanceGroups", "POST", "/", req, resp)
	return
}

// ListInstances provides information about the cluster instances that
// Amazon EMR provisions on behalf of a user when it creates the cluster.
// For example, this operation indicates when the EC2 instances reach the
// Ready state, when instances become available to Amazon EMR to use for
// jobs, and the IP addresses for cluster instances, etc.
func (c *EMR) ListInstances(req ListInstancesInput) (resp *ListInstancesOutput, err error) {
	resp = &ListInstancesOutput{}
	err = c.client.Do("ListInstances", "POST", "/", req, resp)
	return
}

// ListSteps is undocumented.
func (c *EMR) ListSteps(req ListStepsInput) (resp *ListStepsOutput, err error) {
	resp = &ListStepsOutput{}
	err = c.client.Do("ListSteps", "POST", "/", req, resp)
	return
}

// ModifyInstanceGroups modifyInstanceGroups modifies the number of nodes
// and configuration settings of an instance group. The input parameters
// include the new target instance count for the group and the instance
// group ID. The call will either succeed or fail atomically.
func (c *EMR) ModifyInstanceGroups(req ModifyInstanceGroupsInput) (err error) {
	// NRE
	err = c.client.Do("ModifyInstanceGroups", "POST", "/", req, nil)
	return
}

// RemoveTags removes tags from an Amazon EMR resource. Tags make it easier
// to associate clusters in various ways, such as grouping clusters to
// track your Amazon EMR resource allocation costs. For more information,
// see Tagging Amazon EMR Resources . The following example removes the
// stack tag with value Prod from a cluster:
func (c *EMR) RemoveTags(req RemoveTagsInput) (resp *RemoveTagsOutput, err error) {
	resp = &RemoveTagsOutput{}
	err = c.client.Do("RemoveTags", "POST", "/", req, resp)
	return
}

// RunJobFlow runJobFlow creates and starts running a new job flow. The job
// flow will run the steps specified. Once the job flow completes, the
// cluster is stopped and the partition is lost. To prevent loss of data,
// configure the last step of the job flow to store results in Amazon S3.
// If the JobFlowInstancesConfig KeepJobFlowAliveWhenNoSteps parameter is
// set to , the job flow will transition to the state rather than shutting
// down once the steps have completed. For additional protection, you can
// set the JobFlowInstancesConfig TerminationProtected parameter to to lock
// the job flow and prevent it from being terminated by API call, user
// intervention, or in the event of a job flow error. A maximum of 256
// steps are allowed in each job flow. If your job flow is long-running
// (such as a Hive data warehouse) or complex, you may require more than
// 256 steps to process your data. You can bypass the 256-step limitation
// in various ways, including using the SSH shell to connect to the master
// node and submitting queries directly to the software running on the
// master node, such as Hive and Hadoop. For more information on how to do
// this, go to Add More than 256 Steps to a Job Flow in the Amazon Elastic
// MapReduce Developer's Guide For long running job flows, we recommend
// that you periodically store your results.
func (c *EMR) RunJobFlow(req RunJobFlowInput) (resp *RunJobFlowOutput, err error) {
	resp = &RunJobFlowOutput{}
	err = c.client.Do("RunJobFlow", "POST", "/", req, resp)
	return
}

// SetTerminationProtection setTerminationProtection locks a job flow so
// the Amazon EC2 instances in the cluster cannot be terminated by user
// intervention, an API call, or in the event of a job-flow error. The
// cluster still terminates upon successful completion of the job flow.
// Calling SetTerminationProtection on a job flow is analogous to calling
// the Amazon EC2 DisableAPITermination API on all of the EC2 instances in
// a cluster. SetTerminationProtection is used to prevent accidental
// termination of a job flow and to ensure that in the event of an error,
// the instances will persist so you can recover any data stored in their
// ephemeral instance storage. To terminate a job flow that has been locked
// by setting SetTerminationProtection to true , you must first unlock the
// job flow by a subsequent call to SetTerminationProtection in which you
// set the value to false . For more information, go to Protecting a Job
// Flow from Termination in the Amazon Elastic MapReduce Developer's Guide.
func (c *EMR) SetTerminationProtection(req SetTerminationProtectionInput) (err error) {
	// NRE
	err = c.client.Do("SetTerminationProtection", "POST", "/", req, nil)
	return
}

// SetVisibleToAllUsers sets whether all AWS Identity and Access Management
// users under your account can access the specified job flows. This action
// works on running job flows. You can also set the visibility of a job
// flow when you launch it using the VisibleToAllUsers parameter of
// RunJobFlow . The SetVisibleToAllUsers action can be called only by an
// IAM user who created the job flow or the AWS account that owns the job
// flow.
func (c *EMR) SetVisibleToAllUsers(req SetVisibleToAllUsersInput) (err error) {
	// NRE
	err = c.client.Do("SetVisibleToAllUsers", "POST", "/", req, nil)
	return
}

// TerminateJobFlows terminateJobFlows shuts a list of job flows down. When
// a job flow is shut down, any step not yet completed is canceled and the
// EC2 instances on which the job flow is running are stopped. Any log
// files not already saved are uploaded to Amazon S3 if a LogUri was
// specified when the job flow was created. The call to TerminateJobFlows
// is asynchronous. Depending on the configuration of the job flow, it may
// take up to 5-20 minutes for the job flow to completely terminate and
// release allocated resources, such as Amazon EC2 instances.
func (c *EMR) TerminateJobFlows(req TerminateJobFlowsInput) (err error) {
	// NRE
	err = c.client.Do("TerminateJobFlows", "POST", "/", req, nil)
	return
}

// AddInstanceGroupsInput is undocumented.
type AddInstanceGroupsInput struct {
	InstanceGroups []InstanceGroupConfig `json:"InstanceGroups"`
	JobFlowID      string                `json:"JobFlowId"`
}

// AddInstanceGroupsOutput is undocumented.
type AddInstanceGroupsOutput struct {
	InstanceGroupIds []string `json:"InstanceGroupIds,omitempty"`
	JobFlowID        string   `json:"JobFlowId,omitempty"`
}

// AddJobFlowStepsInput is undocumented.
type AddJobFlowStepsInput struct {
	JobFlowID string       `json:"JobFlowId"`
	Steps     []StepConfig `json:"Steps"`
}

// AddJobFlowStepsOutput is undocumented.
type AddJobFlowStepsOutput struct {
	StepIds []string `json:"StepIds,omitempty"`
}

// AddTagsInput is undocumented.
type AddTagsInput struct {
	ResourceID string `json:"ResourceId"`
	Tags       []Tag  `json:"Tags"`
}

// AddTagsOutput is undocumented.
type AddTagsOutput struct {
}

// Application is undocumented.
type Application struct {
	AdditionalInfo map[string]string `json:"AdditionalInfo,omitempty"`
	Args           []string          `json:"Args,omitempty"`
	Name           string            `json:"Name,omitempty"`
	Version        string            `json:"Version,omitempty"`
}

// BootstrapActionConfig is undocumented.
type BootstrapActionConfig struct {
	Name                  string                      `json:"Name"`
	ScriptBootstrapAction ScriptBootstrapActionConfig `json:"ScriptBootstrapAction"`
}

// BootstrapActionDetail is undocumented.
type BootstrapActionDetail struct {
	BootstrapActionConfig BootstrapActionConfig `json:"BootstrapActionConfig,omitempty"`
}

// Cluster is undocumented.
type Cluster struct {
	Applications          []Application         `json:"Applications,omitempty"`
	AutoTerminate         bool                  `json:"AutoTerminate,omitempty"`
	Ec2InstanceAttributes Ec2InstanceAttributes `json:"Ec2InstanceAttributes,omitempty"`
	ID                    string                `json:"Id,omitempty"`
	LogURI                string                `json:"LogUri,omitempty"`
	Name                  string                `json:"Name,omitempty"`
	RequestedAmiVersion   string                `json:"RequestedAmiVersion,omitempty"`
	RunningAmiVersion     string                `json:"RunningAmiVersion,omitempty"`
	ServiceRole           string                `json:"ServiceRole,omitempty"`
	Status                ClusterStatus         `json:"Status,omitempty"`
	Tags                  []Tag                 `json:"Tags,omitempty"`
	TerminationProtected  bool                  `json:"TerminationProtected,omitempty"`
	VisibleToAllUsers     bool                  `json:"VisibleToAllUsers,omitempty"`
}

// ClusterStateChangeReason is undocumented.
type ClusterStateChangeReason struct {
	Code    string `json:"Code,omitempty"`
	Message string `json:"Message,omitempty"`
}

// ClusterStatus is undocumented.
type ClusterStatus struct {
	State             string                   `json:"State,omitempty"`
	StateChangeReason ClusterStateChangeReason `json:"StateChangeReason,omitempty"`
	Timeline          ClusterTimeline          `json:"Timeline,omitempty"`
}

// ClusterSummary is undocumented.
type ClusterSummary struct {
	ID     string        `json:"Id,omitempty"`
	Name   string        `json:"Name,omitempty"`
	Status ClusterStatus `json:"Status,omitempty"`
}

// ClusterTimeline is undocumented.
type ClusterTimeline struct {
	CreationDateTime time.Time `json:"CreationDateTime,omitempty"`
	EndDateTime      time.Time `json:"EndDateTime,omitempty"`
	ReadyDateTime    time.Time `json:"ReadyDateTime,omitempty"`
}

// Command is undocumented.
type Command struct {
	Args       []string `json:"Args,omitempty"`
	Name       string   `json:"Name,omitempty"`
	ScriptPath string   `json:"ScriptPath,omitempty"`
}

// DescribeClusterInput is undocumented.
type DescribeClusterInput struct {
	ClusterID string `json:"ClusterId"`
}

// DescribeClusterOutput is undocumented.
type DescribeClusterOutput struct {
	Cluster Cluster `json:"Cluster,omitempty"`
}

// DescribeJobFlowsInput is undocumented.
type DescribeJobFlowsInput struct {
	CreatedAfter  time.Time `json:"CreatedAfter,omitempty"`
	CreatedBefore time.Time `json:"CreatedBefore,omitempty"`
	JobFlowIds    []string  `json:"JobFlowIds,omitempty"`
	JobFlowStates []string  `json:"JobFlowStates,omitempty"`
}

// DescribeJobFlowsOutput is undocumented.
type DescribeJobFlowsOutput struct {
	JobFlows []JobFlowDetail `json:"JobFlows,omitempty"`
}

// DescribeStepInput is undocumented.
type DescribeStepInput struct {
	ClusterID string `json:"ClusterId"`
	StepID    string `json:"StepId"`
}

// DescribeStepOutput is undocumented.
type DescribeStepOutput struct {
	Step Step `json:"Step,omitempty"`
}

// Ec2InstanceAttributes is undocumented.
type Ec2InstanceAttributes struct {
	Ec2AvailabilityZone string `json:"Ec2AvailabilityZone,omitempty"`
	Ec2KeyName          string `json:"Ec2KeyName,omitempty"`
	Ec2SubnetID         string `json:"Ec2SubnetId,omitempty"`
	IamInstanceProfile  string `json:"IamInstanceProfile,omitempty"`
}

// HadoopJarStepConfig is undocumented.
type HadoopJarStepConfig struct {
	Args       []string   `json:"Args,omitempty"`
	Jar        string     `json:"Jar"`
	MainClass  string     `json:"MainClass,omitempty"`
	Properties []KeyValue `json:"Properties,omitempty"`
}

// HadoopStepConfig is undocumented.
type HadoopStepConfig struct {
	Args       []string          `json:"Args,omitempty"`
	Jar        string            `json:"Jar,omitempty"`
	MainClass  string            `json:"MainClass,omitempty"`
	Properties map[string]string `json:"Properties,omitempty"`
}

// Instance is undocumented.
type Instance struct {
	Ec2InstanceID    string         `json:"Ec2InstanceId,omitempty"`
	ID               string         `json:"Id,omitempty"`
	PrivateDNSName   string         `json:"PrivateDnsName,omitempty"`
	PrivateIPAddress string         `json:"PrivateIpAddress,omitempty"`
	PublicDNSName    string         `json:"PublicDnsName,omitempty"`
	PublicIPAddress  string         `json:"PublicIpAddress,omitempty"`
	Status           InstanceStatus `json:"Status,omitempty"`
}

// InstanceGroup is undocumented.
type InstanceGroup struct {
	BidPrice               string              `json:"BidPrice,omitempty"`
	ID                     string              `json:"Id,omitempty"`
	InstanceGroupType      string              `json:"InstanceGroupType,omitempty"`
	InstanceType           string              `json:"InstanceType,omitempty"`
	Market                 string              `json:"Market,omitempty"`
	Name                   string              `json:"Name,omitempty"`
	RequestedInstanceCount int                 `json:"RequestedInstanceCount,omitempty"`
	RunningInstanceCount   int                 `json:"RunningInstanceCount,omitempty"`
	Status                 InstanceGroupStatus `json:"Status,omitempty"`
}

// InstanceGroupConfig is undocumented.
type InstanceGroupConfig struct {
	BidPrice      string `json:"BidPrice,omitempty"`
	InstanceCount int    `json:"InstanceCount"`
	InstanceRole  string `json:"InstanceRole"`
	InstanceType  string `json:"InstanceType"`
	Market        string `json:"Market,omitempty"`
	Name          string `json:"Name,omitempty"`
}

// InstanceGroupDetail is undocumented.
type InstanceGroupDetail struct {
	BidPrice              string    `json:"BidPrice,omitempty"`
	CreationDateTime      time.Time `json:"CreationDateTime"`
	EndDateTime           time.Time `json:"EndDateTime,omitempty"`
	InstanceGroupID       string    `json:"InstanceGroupId,omitempty"`
	InstanceRequestCount  int       `json:"InstanceRequestCount"`
	InstanceRole          string    `json:"InstanceRole"`
	InstanceRunningCount  int       `json:"InstanceRunningCount"`
	InstanceType          string    `json:"InstanceType"`
	LastStateChangeReason string    `json:"LastStateChangeReason,omitempty"`
	Market                string    `json:"Market"`
	Name                  string    `json:"Name,omitempty"`
	ReadyDateTime         time.Time `json:"ReadyDateTime,omitempty"`
	StartDateTime         time.Time `json:"StartDateTime,omitempty"`
	State                 string    `json:"State"`
}

// InstanceGroupModifyConfig is undocumented.
type InstanceGroupModifyConfig struct {
	EC2InstanceIdsToTerminate []string `json:"EC2InstanceIdsToTerminate,omitempty"`
	InstanceCount             int      `json:"InstanceCount,omitempty"`
	InstanceGroupID           string   `json:"InstanceGroupId"`
}

// InstanceGroupStateChangeReason is undocumented.
type InstanceGroupStateChangeReason struct {
	Code    string `json:"Code,omitempty"`
	Message string `json:"Message,omitempty"`
}

// InstanceGroupStatus is undocumented.
type InstanceGroupStatus struct {
	State             string                         `json:"State,omitempty"`
	StateChangeReason InstanceGroupStateChangeReason `json:"StateChangeReason,omitempty"`
	Timeline          InstanceGroupTimeline          `json:"Timeline,omitempty"`
}

// InstanceGroupTimeline is undocumented.
type InstanceGroupTimeline struct {
	CreationDateTime time.Time `json:"CreationDateTime,omitempty"`
	EndDateTime      time.Time `json:"EndDateTime,omitempty"`
	ReadyDateTime    time.Time `json:"ReadyDateTime,omitempty"`
}

// InstanceStateChangeReason is undocumented.
type InstanceStateChangeReason struct {
	Code    string `json:"Code,omitempty"`
	Message string `json:"Message,omitempty"`
}

// InstanceStatus is undocumented.
type InstanceStatus struct {
	State             string                    `json:"State,omitempty"`
	StateChangeReason InstanceStateChangeReason `json:"StateChangeReason,omitempty"`
	Timeline          InstanceTimeline          `json:"Timeline,omitempty"`
}

// InstanceTimeline is undocumented.
type InstanceTimeline struct {
	CreationDateTime time.Time `json:"CreationDateTime,omitempty"`
	EndDateTime      time.Time `json:"EndDateTime,omitempty"`
	ReadyDateTime    time.Time `json:"ReadyDateTime,omitempty"`
}

// JobFlowDetail is undocumented.
type JobFlowDetail struct {
	AmiVersion            string                       `json:"AmiVersion,omitempty"`
	BootstrapActions      []BootstrapActionDetail      `json:"BootstrapActions,omitempty"`
	ExecutionStatusDetail JobFlowExecutionStatusDetail `json:"ExecutionStatusDetail"`
	Instances             JobFlowInstancesDetail       `json:"Instances"`
	JobFlowID             string                       `json:"JobFlowId"`
	JobFlowRole           string                       `json:"JobFlowRole,omitempty"`
	LogURI                string                       `json:"LogUri,omitempty"`
	Name                  string                       `json:"Name"`
	ServiceRole           string                       `json:"ServiceRole,omitempty"`
	Steps                 []StepDetail                 `json:"Steps,omitempty"`
	SupportedProducts     []string                     `json:"SupportedProducts,omitempty"`
	VisibleToAllUsers     bool                         `json:"VisibleToAllUsers,omitempty"`
}

// JobFlowExecutionStatusDetail is undocumented.
type JobFlowExecutionStatusDetail struct {
	CreationDateTime      time.Time `json:"CreationDateTime"`
	EndDateTime           time.Time `json:"EndDateTime,omitempty"`
	LastStateChangeReason string    `json:"LastStateChangeReason,omitempty"`
	ReadyDateTime         time.Time `json:"ReadyDateTime,omitempty"`
	StartDateTime         time.Time `json:"StartDateTime,omitempty"`
	State                 string    `json:"State"`
}

// JobFlowInstancesConfig is undocumented.
type JobFlowInstancesConfig struct {
	Ec2KeyName                  string                `json:"Ec2KeyName,omitempty"`
	Ec2SubnetID                 string                `json:"Ec2SubnetId,omitempty"`
	HadoopVersion               string                `json:"HadoopVersion,omitempty"`
	InstanceCount               int                   `json:"InstanceCount,omitempty"`
	InstanceGroups              []InstanceGroupConfig `json:"InstanceGroups,omitempty"`
	KeepJobFlowAliveWhenNoSteps bool                  `json:"KeepJobFlowAliveWhenNoSteps,omitempty"`
	MasterInstanceType          string                `json:"MasterInstanceType,omitempty"`
	Placement                   PlacementType         `json:"Placement,omitempty"`
	SlaveInstanceType           string                `json:"SlaveInstanceType,omitempty"`
	TerminationProtected        bool                  `json:"TerminationProtected,omitempty"`
}

// JobFlowInstancesDetail is undocumented.
type JobFlowInstancesDetail struct {
	Ec2KeyName                  string                `json:"Ec2KeyName,omitempty"`
	Ec2SubnetID                 string                `json:"Ec2SubnetId,omitempty"`
	HadoopVersion               string                `json:"HadoopVersion,omitempty"`
	InstanceCount               int                   `json:"InstanceCount"`
	InstanceGroups              []InstanceGroupDetail `json:"InstanceGroups,omitempty"`
	KeepJobFlowAliveWhenNoSteps bool                  `json:"KeepJobFlowAliveWhenNoSteps,omitempty"`
	MasterInstanceID            string                `json:"MasterInstanceId,omitempty"`
	MasterInstanceType          string                `json:"MasterInstanceType"`
	MasterPublicDNSName         string                `json:"MasterPublicDnsName,omitempty"`
	NormalizedInstanceHours     int                   `json:"NormalizedInstanceHours,omitempty"`
	Placement                   PlacementType         `json:"Placement,omitempty"`
	SlaveInstanceType           string                `json:"SlaveInstanceType"`
	TerminationProtected        bool                  `json:"TerminationProtected,omitempty"`
}

// KeyValue is undocumented.
type KeyValue struct {
	Key   string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
}

// ListBootstrapActionsInput is undocumented.
type ListBootstrapActionsInput struct {
	ClusterID string `json:"ClusterId"`
	Marker    string `json:"Marker,omitempty"`
}

// ListBootstrapActionsOutput is undocumented.
type ListBootstrapActionsOutput struct {
	BootstrapActions []Command `json:"BootstrapActions,omitempty"`
	Marker           string    `json:"Marker,omitempty"`
}

// ListClustersInput is undocumented.
type ListClustersInput struct {
	ClusterStates []string  `json:"ClusterStates,omitempty"`
	CreatedAfter  time.Time `json:"CreatedAfter,omitempty"`
	CreatedBefore time.Time `json:"CreatedBefore,omitempty"`
	Marker        string    `json:"Marker,omitempty"`
}

// ListClustersOutput is undocumented.
type ListClustersOutput struct {
	Clusters []ClusterSummary `json:"Clusters,omitempty"`
	Marker   string           `json:"Marker,omitempty"`
}

// ListInstanceGroupsInput is undocumented.
type ListInstanceGroupsInput struct {
	ClusterID string `json:"ClusterId"`
	Marker    string `json:"Marker,omitempty"`
}

// ListInstanceGroupsOutput is undocumented.
type ListInstanceGroupsOutput struct {
	InstanceGroups []InstanceGroup `json:"InstanceGroups,omitempty"`
	Marker         string          `json:"Marker,omitempty"`
}

// ListInstancesInput is undocumented.
type ListInstancesInput struct {
	ClusterID          string   `json:"ClusterId"`
	InstanceGroupID    string   `json:"InstanceGroupId,omitempty"`
	InstanceGroupTypes []string `json:"InstanceGroupTypes,omitempty"`
	Marker             string   `json:"Marker,omitempty"`
}

// ListInstancesOutput is undocumented.
type ListInstancesOutput struct {
	Instances []Instance `json:"Instances,omitempty"`
	Marker    string     `json:"Marker,omitempty"`
}

// ListStepsInput is undocumented.
type ListStepsInput struct {
	ClusterID  string   `json:"ClusterId"`
	Marker     string   `json:"Marker,omitempty"`
	StepStates []string `json:"StepStates,omitempty"`
}

// ListStepsOutput is undocumented.
type ListStepsOutput struct {
	Marker string        `json:"Marker,omitempty"`
	Steps  []StepSummary `json:"Steps,omitempty"`
}

// ModifyInstanceGroupsInput is undocumented.
type ModifyInstanceGroupsInput struct {
	InstanceGroups []InstanceGroupModifyConfig `json:"InstanceGroups,omitempty"`
}

// PlacementType is undocumented.
type PlacementType struct {
	AvailabilityZone string `json:"AvailabilityZone"`
}

// RemoveTagsInput is undocumented.
type RemoveTagsInput struct {
	ResourceID string   `json:"ResourceId"`
	TagKeys    []string `json:"TagKeys"`
}

// RemoveTagsOutput is undocumented.
type RemoveTagsOutput struct {
}

// RunJobFlowInput is undocumented.
type RunJobFlowInput struct {
	AdditionalInfo       string                   `json:"AdditionalInfo,omitempty"`
	AmiVersion           string                   `json:"AmiVersion,omitempty"`
	BootstrapActions     []BootstrapActionConfig  `json:"BootstrapActions,omitempty"`
	Instances            JobFlowInstancesConfig   `json:"Instances"`
	JobFlowRole          string                   `json:"JobFlowRole,omitempty"`
	LogURI               string                   `json:"LogUri,omitempty"`
	Name                 string                   `json:"Name"`
	NewSupportedProducts []SupportedProductConfig `json:"NewSupportedProducts,omitempty"`
	ServiceRole          string                   `json:"ServiceRole,omitempty"`
	Steps                []StepConfig             `json:"Steps,omitempty"`
	SupportedProducts    []string                 `json:"SupportedProducts,omitempty"`
	Tags                 []Tag                    `json:"Tags,omitempty"`
	VisibleToAllUsers    bool                     `json:"VisibleToAllUsers,omitempty"`
}

// RunJobFlowOutput is undocumented.
type RunJobFlowOutput struct {
	JobFlowID string `json:"JobFlowId,omitempty"`
}

// ScriptBootstrapActionConfig is undocumented.
type ScriptBootstrapActionConfig struct {
	Args []string `json:"Args,omitempty"`
	Path string   `json:"Path"`
}

// SetTerminationProtectionInput is undocumented.
type SetTerminationProtectionInput struct {
	JobFlowIds           []string `json:"JobFlowIds"`
	TerminationProtected bool     `json:"TerminationProtected"`
}

// SetVisibleToAllUsersInput is undocumented.
type SetVisibleToAllUsersInput struct {
	JobFlowIds        []string `json:"JobFlowIds"`
	VisibleToAllUsers bool     `json:"VisibleToAllUsers"`
}

// Step is undocumented.
type Step struct {
	ActionOnFailure string           `json:"ActionOnFailure,omitempty"`
	Config          HadoopStepConfig `json:"Config,omitempty"`
	ID              string           `json:"Id,omitempty"`
	Name            string           `json:"Name,omitempty"`
	Status          StepStatus       `json:"Status,omitempty"`
}

// StepConfig is undocumented.
type StepConfig struct {
	ActionOnFailure string              `json:"ActionOnFailure,omitempty"`
	HadoopJarStep   HadoopJarStepConfig `json:"HadoopJarStep"`
	Name            string              `json:"Name"`
}

// StepDetail is undocumented.
type StepDetail struct {
	ExecutionStatusDetail StepExecutionStatusDetail `json:"ExecutionStatusDetail"`
	StepConfig            StepConfig                `json:"StepConfig"`
}

// StepExecutionStatusDetail is undocumented.
type StepExecutionStatusDetail struct {
	CreationDateTime      time.Time `json:"CreationDateTime"`
	EndDateTime           time.Time `json:"EndDateTime,omitempty"`
	LastStateChangeReason string    `json:"LastStateChangeReason,omitempty"`
	StartDateTime         time.Time `json:"StartDateTime,omitempty"`
	State                 string    `json:"State"`
}

// StepStateChangeReason is undocumented.
type StepStateChangeReason struct {
	Code    string `json:"Code,omitempty"`
	Message string `json:"Message,omitempty"`
}

// StepStatus is undocumented.
type StepStatus struct {
	State             string                `json:"State,omitempty"`
	StateChangeReason StepStateChangeReason `json:"StateChangeReason,omitempty"`
	Timeline          StepTimeline          `json:"Timeline,omitempty"`
}

// StepSummary is undocumented.
type StepSummary struct {
	ID     string     `json:"Id,omitempty"`
	Name   string     `json:"Name,omitempty"`
	Status StepStatus `json:"Status,omitempty"`
}

// StepTimeline is undocumented.
type StepTimeline struct {
	CreationDateTime time.Time `json:"CreationDateTime,omitempty"`
	EndDateTime      time.Time `json:"EndDateTime,omitempty"`
	StartDateTime    time.Time `json:"StartDateTime,omitempty"`
}

// SupportedProductConfig is undocumented.
type SupportedProductConfig struct {
	Args []string `json:"Args,omitempty"`
	Name string   `json:"Name,omitempty"`
}

// Tag is undocumented.
type Tag struct {
	Key   string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
}

// TerminateJobFlowsInput is undocumented.
type TerminateJobFlowsInput struct {
	JobFlowIds []string `json:"JobFlowIds"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
