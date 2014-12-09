// Package codedeploy provides a client for AWS CodeDeploy.
package codedeploy

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/bmizerany/aws4"
	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// CodeDeploy is a client for AWS CodeDeploy.
type CodeDeploy struct {
	client aws.Client
}

// New returns a new CodeDeploy client.
func New(key, secret, region string, client *http.Client) *CodeDeploy {
	if client == nil {
		client = http.DefaultClient
	}

	return &CodeDeploy{
		client: &aws.JSONClient{
			Client: &aws4.Client{
				Keys: &aws4.Keys{
					AccessKey: key,
					SecretKey: secret,
				},
				Client: client,
			},
			Endpoint:     endpoints.Lookup("codedeploy", region),
			JSONVersion:  "1.1",
			TargetPrefix: "CodeDeploy_20141006",
		},
	}
}

// BatchGetApplications is undocumented.
func (c *CodeDeploy) BatchGetApplications(req BatchGetApplicationsInput) (resp *BatchGetApplicationsOutput, err error) {
	resp = &BatchGetApplicationsOutput{}
	err = c.client.Do("BatchGetApplications", "POST", "/", req, resp)
	return
}

// BatchGetDeployments is undocumented.
func (c *CodeDeploy) BatchGetDeployments(req BatchGetDeploymentsInput) (resp *BatchGetDeploymentsOutput, err error) {
	resp = &BatchGetDeploymentsOutput{}
	err = c.client.Do("BatchGetDeployments", "POST", "/", req, resp)
	return
}

// CreateApplication is undocumented.
func (c *CodeDeploy) CreateApplication(req CreateApplicationInput) (resp *CreateApplicationOutput, err error) {
	resp = &CreateApplicationOutput{}
	err = c.client.Do("CreateApplication", "POST", "/", req, resp)
	return
}

// CreateDeployment deploys an application revision to the specified
// deployment group.
func (c *CodeDeploy) CreateDeployment(req CreateDeploymentInput) (resp *CreateDeploymentOutput, err error) {
	resp = &CreateDeploymentOutput{}
	err = c.client.Do("CreateDeployment", "POST", "/", req, resp)
	return
}

// CreateDeploymentConfig is undocumented.
func (c *CodeDeploy) CreateDeploymentConfig(req CreateDeploymentConfigInput) (resp *CreateDeploymentConfigOutput, err error) {
	resp = &CreateDeploymentConfigOutput{}
	err = c.client.Do("CreateDeploymentConfig", "POST", "/", req, resp)
	return
}

// CreateDeploymentGroup creates a new deployment group for application
// revisions to be deployed to.
func (c *CodeDeploy) CreateDeploymentGroup(req CreateDeploymentGroupInput) (resp *CreateDeploymentGroupOutput, err error) {
	resp = &CreateDeploymentGroupOutput{}
	err = c.client.Do("CreateDeploymentGroup", "POST", "/", req, resp)
	return
}

// DeleteApplication is undocumented.
func (c *CodeDeploy) DeleteApplication(req DeleteApplicationInput) (err error) {
	// NRE
	err = c.client.Do("DeleteApplication", "POST", "/", req, nil)
	return
}

// DeleteDeploymentConfig deletes a deployment configuration. A deployment
// configuration cannot be deleted if it is currently in use. Also,
// predefined configurations cannot be deleted.
func (c *CodeDeploy) DeleteDeploymentConfig(req DeleteDeploymentConfigInput) (err error) {
	// NRE
	err = c.client.Do("DeleteDeploymentConfig", "POST", "/", req, nil)
	return
}

// DeleteDeploymentGroup is undocumented.
func (c *CodeDeploy) DeleteDeploymentGroup(req DeleteDeploymentGroupInput) (resp *DeleteDeploymentGroupOutput, err error) {
	resp = &DeleteDeploymentGroupOutput{}
	err = c.client.Do("DeleteDeploymentGroup", "POST", "/", req, resp)
	return
}

// GetApplication is undocumented.
func (c *CodeDeploy) GetApplication(req GetApplicationInput) (resp *GetApplicationOutput, err error) {
	resp = &GetApplicationOutput{}
	err = c.client.Do("GetApplication", "POST", "/", req, resp)
	return
}

// GetApplicationRevision is undocumented.
func (c *CodeDeploy) GetApplicationRevision(req GetApplicationRevisionInput) (resp *GetApplicationRevisionOutput, err error) {
	resp = &GetApplicationRevisionOutput{}
	err = c.client.Do("GetApplicationRevision", "POST", "/", req, resp)
	return
}

// GetDeployment is undocumented.
func (c *CodeDeploy) GetDeployment(req GetDeploymentInput) (resp *GetDeploymentOutput, err error) {
	resp = &GetDeploymentOutput{}
	err = c.client.Do("GetDeployment", "POST", "/", req, resp)
	return
}

// GetDeploymentConfig is undocumented.
func (c *CodeDeploy) GetDeploymentConfig(req GetDeploymentConfigInput) (resp *GetDeploymentConfigOutput, err error) {
	resp = &GetDeploymentConfigOutput{}
	err = c.client.Do("GetDeploymentConfig", "POST", "/", req, resp)
	return
}

// GetDeploymentGroup is undocumented.
func (c *CodeDeploy) GetDeploymentGroup(req GetDeploymentGroupInput) (resp *GetDeploymentGroupOutput, err error) {
	resp = &GetDeploymentGroupOutput{}
	err = c.client.Do("GetDeploymentGroup", "POST", "/", req, resp)
	return
}

// GetDeploymentInstance gets information about an Amazon EC2 instance as
// part of a deployment.
func (c *CodeDeploy) GetDeploymentInstance(req GetDeploymentInstanceInput) (resp *GetDeploymentInstanceOutput, err error) {
	resp = &GetDeploymentInstanceOutput{}
	err = c.client.Do("GetDeploymentInstance", "POST", "/", req, resp)
	return
}

// ListApplicationRevisions lists information about revisions for an
// application.
func (c *CodeDeploy) ListApplicationRevisions(req ListApplicationRevisionsInput) (resp *ListApplicationRevisionsOutput, err error) {
	resp = &ListApplicationRevisionsOutput{}
	err = c.client.Do("ListApplicationRevisions", "POST", "/", req, resp)
	return
}

// ListApplications lists the applications registered within the AWS user
// account.
func (c *CodeDeploy) ListApplications(req ListApplicationsInput) (resp *ListApplicationsOutput, err error) {
	resp = &ListApplicationsOutput{}
	err = c.client.Do("ListApplications", "POST", "/", req, resp)
	return
}

// ListDeploymentConfigs lists the deployment configurations within the AWS
// user account.
func (c *CodeDeploy) ListDeploymentConfigs(req ListDeploymentConfigsInput) (resp *ListDeploymentConfigsOutput, err error) {
	resp = &ListDeploymentConfigsOutput{}
	err = c.client.Do("ListDeploymentConfigs", "POST", "/", req, resp)
	return
}

// ListDeploymentGroups lists the deployment groups for an application
// registered within the AWS user account.
func (c *CodeDeploy) ListDeploymentGroups(req ListDeploymentGroupsInput) (resp *ListDeploymentGroupsOutput, err error) {
	resp = &ListDeploymentGroupsOutput{}
	err = c.client.Do("ListDeploymentGroups", "POST", "/", req, resp)
	return
}

// ListDeploymentInstances lists the Amazon EC2 instances for a deployment
// within the AWS user account.
func (c *CodeDeploy) ListDeploymentInstances(req ListDeploymentInstancesInput) (resp *ListDeploymentInstancesOutput, err error) {
	resp = &ListDeploymentInstancesOutput{}
	err = c.client.Do("ListDeploymentInstances", "POST", "/", req, resp)
	return
}

// ListDeployments lists the deployments under a deployment group for an
// application registered within the AWS user account.
func (c *CodeDeploy) ListDeployments(req ListDeploymentsInput) (resp *ListDeploymentsOutput, err error) {
	resp = &ListDeploymentsOutput{}
	err = c.client.Do("ListDeployments", "POST", "/", req, resp)
	return
}

// RegisterApplicationRevision registers with AWS CodeDeploy a revision for
// the specified application.
func (c *CodeDeploy) RegisterApplicationRevision(req RegisterApplicationRevisionInput) (err error) {
	// NRE
	err = c.client.Do("RegisterApplicationRevision", "POST", "/", req, nil)
	return
}

// StopDeployment is undocumented.
func (c *CodeDeploy) StopDeployment(req StopDeploymentInput) (resp *StopDeploymentOutput, err error) {
	resp = &StopDeploymentOutput{}
	err = c.client.Do("StopDeployment", "POST", "/", req, resp)
	return
}

// UpdateApplication is undocumented.
func (c *CodeDeploy) UpdateApplication(req UpdateApplicationInput) (err error) {
	// NRE
	err = c.client.Do("UpdateApplication", "POST", "/", req, nil)
	return
}

// UpdateDeploymentGroup changes information about an existing deployment
// group.
func (c *CodeDeploy) UpdateDeploymentGroup(req UpdateDeploymentGroupInput) (resp *UpdateDeploymentGroupOutput, err error) {
	resp = &UpdateDeploymentGroupOutput{}
	err = c.client.Do("UpdateDeploymentGroup", "POST", "/", req, resp)
	return
}

// ApplicationInfo is undocumented.
type ApplicationInfo struct {
	ApplicationID   string    `json:"applicationId,omitempty"`
	ApplicationName string    `json:"applicationName,omitempty"`
	CreateTime      time.Time `json:"createTime,omitempty"`
	LinkedToGitHub  bool      `json:"linkedToGitHub,omitempty"`
}

// AutoScalingGroup is undocumented.
type AutoScalingGroup struct {
	Hook string `json:"hook,omitempty"`
	Name string `json:"name,omitempty"`
}

// BatchGetApplicationsInput is undocumented.
type BatchGetApplicationsInput struct {
	ApplicationNames []string `json:"applicationNames,omitempty"`
}

// BatchGetApplicationsOutput is undocumented.
type BatchGetApplicationsOutput struct {
	ApplicationsInfo []ApplicationInfo `json:"applicationsInfo,omitempty"`
}

// BatchGetDeploymentsInput is undocumented.
type BatchGetDeploymentsInput struct {
	DeploymentIds []string `json:"deploymentIds,omitempty"`
}

// BatchGetDeploymentsOutput is undocumented.
type BatchGetDeploymentsOutput struct {
	DeploymentsInfo []DeploymentInfo `json:"deploymentsInfo,omitempty"`
}

// CreateApplicationInput is undocumented.
type CreateApplicationInput struct {
	ApplicationName string `json:"applicationName"`
}

// CreateApplicationOutput is undocumented.
type CreateApplicationOutput struct {
	ApplicationID string `json:"applicationId,omitempty"`
}

// CreateDeploymentConfigInput is undocumented.
type CreateDeploymentConfigInput struct {
	DeploymentConfigName string              `json:"deploymentConfigName"`
	MinimumHealthyHosts  MinimumHealthyHosts `json:"minimumHealthyHosts,omitempty"`
}

// CreateDeploymentConfigOutput is undocumented.
type CreateDeploymentConfigOutput struct {
	DeploymentConfigID string `json:"deploymentConfigId,omitempty"`
}

// CreateDeploymentGroupInput is undocumented.
type CreateDeploymentGroupInput struct {
	ApplicationName      string         `json:"applicationName"`
	AutoScalingGroups    []string       `json:"autoScalingGroups,omitempty"`
	DeploymentConfigName string         `json:"deploymentConfigName,omitempty"`
	DeploymentGroupName  string         `json:"deploymentGroupName"`
	Ec2TagFilters        []EC2TagFilter `json:"ec2TagFilters,omitempty"`
	ServiceRoleARN       string         `json:"serviceRoleArn,omitempty"`
}

// CreateDeploymentGroupOutput is undocumented.
type CreateDeploymentGroupOutput struct {
	DeploymentGroupID string `json:"deploymentGroupId,omitempty"`
}

// CreateDeploymentInput is undocumented.
type CreateDeploymentInput struct {
	ApplicationName               string           `json:"applicationName"`
	DeploymentConfigName          string           `json:"deploymentConfigName,omitempty"`
	DeploymentGroupName           string           `json:"deploymentGroupName,omitempty"`
	Description                   string           `json:"description,omitempty"`
	IgnoreApplicationStopFailures bool             `json:"ignoreApplicationStopFailures,omitempty"`
	Revision                      RevisionLocation `json:"revision,omitempty"`
}

// CreateDeploymentOutput is undocumented.
type CreateDeploymentOutput struct {
	DeploymentID string `json:"deploymentId,omitempty"`
}

// DeleteApplicationInput is undocumented.
type DeleteApplicationInput struct {
	ApplicationName string `json:"applicationName"`
}

// DeleteDeploymentConfigInput is undocumented.
type DeleteDeploymentConfigInput struct {
	DeploymentConfigName string `json:"deploymentConfigName"`
}

// DeleteDeploymentGroupInput is undocumented.
type DeleteDeploymentGroupInput struct {
	ApplicationName     string `json:"applicationName"`
	DeploymentGroupName string `json:"deploymentGroupName"`
}

// DeleteDeploymentGroupOutput is undocumented.
type DeleteDeploymentGroupOutput struct {
	HooksNotCleanedUp []AutoScalingGroup `json:"hooksNotCleanedUp,omitempty"`
}

// DeploymentConfigInfo is undocumented.
type DeploymentConfigInfo struct {
	CreateTime           time.Time           `json:"createTime,omitempty"`
	DeploymentConfigID   string              `json:"deploymentConfigId,omitempty"`
	DeploymentConfigName string              `json:"deploymentConfigName,omitempty"`
	MinimumHealthyHosts  MinimumHealthyHosts `json:"minimumHealthyHosts,omitempty"`
}

// DeploymentGroupInfo is undocumented.
type DeploymentGroupInfo struct {
	ApplicationName      string             `json:"applicationName,omitempty"`
	AutoScalingGroups    []AutoScalingGroup `json:"autoScalingGroups,omitempty"`
	DeploymentConfigName string             `json:"deploymentConfigName,omitempty"`
	DeploymentGroupID    string             `json:"deploymentGroupId,omitempty"`
	DeploymentGroupName  string             `json:"deploymentGroupName,omitempty"`
	Ec2TagFilters        []EC2TagFilter     `json:"ec2TagFilters,omitempty"`
	ServiceRoleARN       string             `json:"serviceRoleArn,omitempty"`
	TargetRevision       RevisionLocation   `json:"targetRevision,omitempty"`
}

// DeploymentInfo is undocumented.
type DeploymentInfo struct {
	ApplicationName               string             `json:"applicationName,omitempty"`
	CompleteTime                  time.Time          `json:"completeTime,omitempty"`
	CreateTime                    time.Time          `json:"createTime,omitempty"`
	Creator                       string             `json:"creator,omitempty"`
	DeploymentConfigName          string             `json:"deploymentConfigName,omitempty"`
	DeploymentGroupName           string             `json:"deploymentGroupName,omitempty"`
	DeploymentID                  string             `json:"deploymentId,omitempty"`
	DeploymentOverview            DeploymentOverview `json:"deploymentOverview,omitempty"`
	Description                   string             `json:"description,omitempty"`
	ErrorInformation              ErrorInformation   `json:"errorInformation,omitempty"`
	IgnoreApplicationStopFailures bool               `json:"ignoreApplicationStopFailures,omitempty"`
	Revision                      RevisionLocation   `json:"revision,omitempty"`
	StartTime                     time.Time          `json:"startTime,omitempty"`
	Status                        string             `json:"status,omitempty"`
}

// DeploymentOverview is undocumented.
type DeploymentOverview struct {
	Failed     int64 `json:"Failed,omitempty"`
	InProgress int64 `json:"InProgress,omitempty"`
	Pending    int64 `json:"Pending,omitempty"`
	Skipped    int64 `json:"Skipped,omitempty"`
	Succeeded  int64 `json:"Succeeded,omitempty"`
}

// Diagnostics is undocumented.
type Diagnostics struct {
	ErrorCode  string `json:"errorCode,omitempty"`
	LogTail    string `json:"logTail,omitempty"`
	Message    string `json:"message,omitempty"`
	ScriptName string `json:"scriptName,omitempty"`
}

// EC2TagFilter is undocumented.
type EC2TagFilter struct {
	Key   string `json:"Key,omitempty"`
	Type  string `json:"Type,omitempty"`
	Value string `json:"Value,omitempty"`
}

// ErrorInformation is undocumented.
type ErrorInformation struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// GenericRevisionInfo is undocumented.
type GenericRevisionInfo struct {
	DeploymentGroups []string  `json:"deploymentGroups,omitempty"`
	Description      string    `json:"description,omitempty"`
	FirstUsedTime    time.Time `json:"firstUsedTime,omitempty"`
	LastUsedTime     time.Time `json:"lastUsedTime,omitempty"`
	RegisterTime     time.Time `json:"registerTime,omitempty"`
}

// GetApplicationInput is undocumented.
type GetApplicationInput struct {
	ApplicationName string `json:"applicationName"`
}

// GetApplicationOutput is undocumented.
type GetApplicationOutput struct {
	Application ApplicationInfo `json:"application,omitempty"`
}

// GetApplicationRevisionInput is undocumented.
type GetApplicationRevisionInput struct {
	ApplicationName string           `json:"applicationName"`
	Revision        RevisionLocation `json:"revision"`
}

// GetApplicationRevisionOutput is undocumented.
type GetApplicationRevisionOutput struct {
	ApplicationName string              `json:"applicationName,omitempty"`
	Revision        RevisionLocation    `json:"revision,omitempty"`
	RevisionInfo    GenericRevisionInfo `json:"revisionInfo,omitempty"`
}

// GetDeploymentConfigInput is undocumented.
type GetDeploymentConfigInput struct {
	DeploymentConfigName string `json:"deploymentConfigName"`
}

// GetDeploymentConfigOutput is undocumented.
type GetDeploymentConfigOutput struct {
	DeploymentConfigInfo DeploymentConfigInfo `json:"deploymentConfigInfo,omitempty"`
}

// GetDeploymentGroupInput is undocumented.
type GetDeploymentGroupInput struct {
	ApplicationName     string `json:"applicationName"`
	DeploymentGroupName string `json:"deploymentGroupName"`
}

// GetDeploymentGroupOutput is undocumented.
type GetDeploymentGroupOutput struct {
	DeploymentGroupInfo DeploymentGroupInfo `json:"deploymentGroupInfo,omitempty"`
}

// GetDeploymentInput is undocumented.
type GetDeploymentInput struct {
	DeploymentID string `json:"deploymentId"`
}

// GetDeploymentInstanceInput is undocumented.
type GetDeploymentInstanceInput struct {
	DeploymentID string `json:"deploymentId"`
	InstanceID   string `json:"instanceId"`
}

// GetDeploymentInstanceOutput is undocumented.
type GetDeploymentInstanceOutput struct {
	InstanceSummary InstanceSummary `json:"instanceSummary,omitempty"`
}

// GetDeploymentOutput is undocumented.
type GetDeploymentOutput struct {
	DeploymentInfo DeploymentInfo `json:"deploymentInfo,omitempty"`
}

// GitHubLocation is undocumented.
type GitHubLocation struct {
	CommitID   string `json:"commitId,omitempty"`
	Repository string `json:"repository,omitempty"`
}

// InstanceSummary is undocumented.
type InstanceSummary struct {
	DeploymentID    string           `json:"deploymentId,omitempty"`
	InstanceID      string           `json:"instanceId,omitempty"`
	LastUpdatedAt   time.Time        `json:"lastUpdatedAt,omitempty"`
	LifecycleEvents []LifecycleEvent `json:"lifecycleEvents,omitempty"`
	Status          string           `json:"status,omitempty"`
}

// LifecycleEvent is undocumented.
type LifecycleEvent struct {
	Diagnostics        Diagnostics `json:"diagnostics,omitempty"`
	EndTime            time.Time   `json:"endTime,omitempty"`
	LifecycleEventName string      `json:"lifecycleEventName,omitempty"`
	StartTime          time.Time   `json:"startTime,omitempty"`
	Status             string      `json:"status,omitempty"`
}

// ListApplicationRevisionsInput is undocumented.
type ListApplicationRevisionsInput struct {
	ApplicationName string `json:"applicationName"`
	Deployed        string `json:"deployed,omitempty"`
	NextToken       string `json:"nextToken,omitempty"`
	S3Bucket        string `json:"s3Bucket,omitempty"`
	S3KeyPrefix     string `json:"s3KeyPrefix,omitempty"`
	SortBy          string `json:"sortBy,omitempty"`
	SortOrder       string `json:"sortOrder,omitempty"`
}

// ListApplicationRevisionsOutput is undocumented.
type ListApplicationRevisionsOutput struct {
	NextToken string             `json:"nextToken,omitempty"`
	Revisions []RevisionLocation `json:"revisions,omitempty"`
}

// ListApplicationsInput is undocumented.
type ListApplicationsInput struct {
	NextToken string `json:"nextToken,omitempty"`
}

// ListApplicationsOutput is undocumented.
type ListApplicationsOutput struct {
	Applications []string `json:"applications,omitempty"`
	NextToken    string   `json:"nextToken,omitempty"`
}

// ListDeploymentConfigsInput is undocumented.
type ListDeploymentConfigsInput struct {
	NextToken string `json:"nextToken,omitempty"`
}

// ListDeploymentConfigsOutput is undocumented.
type ListDeploymentConfigsOutput struct {
	DeploymentConfigsList []string `json:"deploymentConfigsList,omitempty"`
	NextToken             string   `json:"nextToken,omitempty"`
}

// ListDeploymentGroupsInput is undocumented.
type ListDeploymentGroupsInput struct {
	ApplicationName string `json:"applicationName"`
	NextToken       string `json:"nextToken,omitempty"`
}

// ListDeploymentGroupsOutput is undocumented.
type ListDeploymentGroupsOutput struct {
	ApplicationName  string   `json:"applicationName,omitempty"`
	DeploymentGroups []string `json:"deploymentGroups,omitempty"`
	NextToken        string   `json:"nextToken,omitempty"`
}

// ListDeploymentInstancesInput is undocumented.
type ListDeploymentInstancesInput struct {
	DeploymentID         string   `json:"deploymentId"`
	InstanceStatusFilter []string `json:"instanceStatusFilter,omitempty"`
	NextToken            string   `json:"nextToken,omitempty"`
}

// ListDeploymentInstancesOutput is undocumented.
type ListDeploymentInstancesOutput struct {
	InstancesList []string `json:"instancesList,omitempty"`
	NextToken     string   `json:"nextToken,omitempty"`
}

// ListDeploymentsInput is undocumented.
type ListDeploymentsInput struct {
	ApplicationName     string    `json:"applicationName,omitempty"`
	CreateTimeRange     TimeRange `json:"createTimeRange,omitempty"`
	DeploymentGroupName string    `json:"deploymentGroupName,omitempty"`
	IncludeOnlyStatuses []string  `json:"includeOnlyStatuses,omitempty"`
	NextToken           string    `json:"nextToken,omitempty"`
}

// ListDeploymentsOutput is undocumented.
type ListDeploymentsOutput struct {
	Deployments []string `json:"deployments,omitempty"`
	NextToken   string   `json:"nextToken,omitempty"`
}

// MinimumHealthyHosts is undocumented.
type MinimumHealthyHosts struct {
	Type  string `json:"type,omitempty"`
	Value int    `json:"value,omitempty"`
}

// RegisterApplicationRevisionInput is undocumented.
type RegisterApplicationRevisionInput struct {
	ApplicationName string           `json:"applicationName"`
	Description     string           `json:"description,omitempty"`
	Revision        RevisionLocation `json:"revision"`
}

// RevisionLocation is undocumented.
type RevisionLocation struct {
	GitHubLocation GitHubLocation `json:"gitHubLocation,omitempty"`
	RevisionType   string         `json:"revisionType,omitempty"`
	S3Location     S3Location     `json:"s3Location,omitempty"`
}

// S3Location is undocumented.
type S3Location struct {
	Bucket     string `json:"bucket,omitempty"`
	BundleType string `json:"bundleType,omitempty"`
	ETag       string `json:"eTag,omitempty"`
	Key        string `json:"key,omitempty"`
	Version    string `json:"version,omitempty"`
}

// StopDeploymentInput is undocumented.
type StopDeploymentInput struct {
	DeploymentID string `json:"deploymentId"`
}

// StopDeploymentOutput is undocumented.
type StopDeploymentOutput struct {
	Status        string `json:"status,omitempty"`
	StatusMessage string `json:"statusMessage,omitempty"`
}

// TimeRange is undocumented.
type TimeRange struct {
	End   time.Time `json:"end,omitempty"`
	Start time.Time `json:"start,omitempty"`
}

// UpdateApplicationInput is undocumented.
type UpdateApplicationInput struct {
	ApplicationName    string `json:"applicationName,omitempty"`
	NewApplicationName string `json:"newApplicationName,omitempty"`
}

// UpdateDeploymentGroupInput is undocumented.
type UpdateDeploymentGroupInput struct {
	ApplicationName            string         `json:"applicationName"`
	AutoScalingGroups          []string       `json:"autoScalingGroups,omitempty"`
	CurrentDeploymentGroupName string         `json:"currentDeploymentGroupName"`
	DeploymentConfigName       string         `json:"deploymentConfigName,omitempty"`
	Ec2TagFilters              []EC2TagFilter `json:"ec2TagFilters,omitempty"`
	NewDeploymentGroupName     string         `json:"newDeploymentGroupName,omitempty"`
	ServiceRoleARN             string         `json:"serviceRoleArn,omitempty"`
}

// UpdateDeploymentGroupOutput is undocumented.
type UpdateDeploymentGroupOutput struct {
	HooksNotCleanedUp []AutoScalingGroup `json:"hooksNotCleanedUp,omitempty"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
