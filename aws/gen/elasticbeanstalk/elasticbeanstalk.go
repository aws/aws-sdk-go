// Package elasticbeanstalk provides a client for AWS Elastic Beanstalk.
package elasticbeanstalk

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// ElasticBeanstalk is a client for AWS Elastic Beanstalk.
type ElasticBeanstalk struct {
	client *aws.QueryClient
}

// New returns a new ElasticBeanstalk client.
func New(key, secret, region string, client *http.Client) *ElasticBeanstalk {
	if client == nil {
		client = http.DefaultClient
	}

	return &ElasticBeanstalk{
		client: &aws.QueryClient{
			Signer: &aws.V4Signer{
				Key:     key,
				Secret:  secret,
				Service: "elasticbeanstalk",
				Region:  region,
				IncludeXAmzContentSha256: true,
			},
			Client:     client,
			Endpoint:   endpoints.Lookup("elasticbeanstalk", region),
			APIVersion: "2010-12-01",
		},
	}
}

// CheckDNSAvailability is undocumented.
func (c *ElasticBeanstalk) CheckDNSAvailability(req CheckDNSAvailabilityMessage) (resp *CheckDNSAvailabilityResult, err error) {
	resp = &CheckDNSAvailabilityResult{}
	err = c.client.Do("CheckDNSAvailability", "POST", "/", req, resp)
	return
}

// CreateApplication creates an application that has one configuration
// template named default and no application versions.
func (c *ElasticBeanstalk) CreateApplication(req CreateApplicationMessage) (resp *CreateApplicationResult, err error) {
	resp = &CreateApplicationResult{}
	err = c.client.Do("CreateApplication", "POST", "/", req, resp)
	return
}

// CreateApplicationVersion creates an application version for the
// specified application.
func (c *ElasticBeanstalk) CreateApplicationVersion(req CreateApplicationVersionMessage) (resp *CreateApplicationVersionResult, err error) {
	resp = &CreateApplicationVersionResult{}
	err = c.client.Do("CreateApplicationVersion", "POST", "/", req, resp)
	return
}

// CreateConfigurationTemplate creates a configuration template. Templates
// are associated with a specific application and are used to deploy
// different versions of the application with the same configuration
// settings.
func (c *ElasticBeanstalk) CreateConfigurationTemplate(req CreateConfigurationTemplateMessage) (resp *CreateConfigurationTemplateResult, err error) {
	resp = &CreateConfigurationTemplateResult{}
	err = c.client.Do("CreateConfigurationTemplate", "POST", "/", req, resp)
	return
}

// CreateEnvironment launches an environment for the specified application
// using the specified configuration.
func (c *ElasticBeanstalk) CreateEnvironment(req CreateEnvironmentMessage) (resp *CreateEnvironmentResult, err error) {
	resp = &CreateEnvironmentResult{}
	err = c.client.Do("CreateEnvironment", "POST", "/", req, resp)
	return
}

// CreateStorageLocation creates the Amazon S3 storage location for the
// account. This location is used to store user log files.
func (c *ElasticBeanstalk) CreateStorageLocation() (resp *CreateStorageLocationResult, err error) {
	resp = &CreateStorageLocationResult{}
	err = c.client.Do("CreateStorageLocation", "POST", "/", nil, resp)
	return
}

// DeleteApplication deletes the specified application along with all
// associated versions and configurations. The application versions will
// not be deleted from your Amazon S3 bucket.
func (c *ElasticBeanstalk) DeleteApplication(req DeleteApplicationMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteApplication", "POST", "/", req, nil)
	return
}

// DeleteApplicationVersion deletes the specified version from the
// specified application.
func (c *ElasticBeanstalk) DeleteApplicationVersion(req DeleteApplicationVersionMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteApplicationVersion", "POST", "/", req, nil)
	return
}

// DeleteConfigurationTemplate is undocumented.
func (c *ElasticBeanstalk) DeleteConfigurationTemplate(req DeleteConfigurationTemplateMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteConfigurationTemplate", "POST", "/", req, nil)
	return
}

// DeleteEnvironmentConfiguration deletes the draft configuration
// associated with the running environment. Updating a running environment
// with any configuration changes creates a draft configuration set. You
// can get the draft configuration using DescribeConfigurationSettings
// while the update is in progress or if the update fails. The
// DeploymentStatus for the draft configuration indicates whether the
// deployment is in process or has failed. The draft configuration remains
// in existence until it is deleted with this action.
func (c *ElasticBeanstalk) DeleteEnvironmentConfiguration(req DeleteEnvironmentConfigurationMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteEnvironmentConfiguration", "POST", "/", req, nil)
	return
}

// DescribeApplicationVersions returns descriptions for existing
// application versions.
func (c *ElasticBeanstalk) DescribeApplicationVersions(req DescribeApplicationVersionsMessage) (resp *DescribeApplicationVersionsResult, err error) {
	resp = &DescribeApplicationVersionsResult{}
	err = c.client.Do("DescribeApplicationVersions", "POST", "/", req, resp)
	return
}

// DescribeApplications is undocumented.
func (c *ElasticBeanstalk) DescribeApplications(req DescribeApplicationsMessage) (resp *DescribeApplicationsResult, err error) {
	resp = &DescribeApplicationsResult{}
	err = c.client.Do("DescribeApplications", "POST", "/", req, resp)
	return
}

// DescribeConfigurationOptions describes the configuration options that
// are used in a particular configuration template or environment, or that
// a specified solution stack defines. The description includes the values
// the options, their default values, and an indication of the required
// action on a running environment if an option value is changed.
func (c *ElasticBeanstalk) DescribeConfigurationOptions(req DescribeConfigurationOptionsMessage) (resp *DescribeConfigurationOptionsResult, err error) {
	resp = &DescribeConfigurationOptionsResult{}
	err = c.client.Do("DescribeConfigurationOptions", "POST", "/", req, resp)
	return
}

// DescribeConfigurationSettings returns a description of the settings for
// the specified configuration set, that is, either a configuration
// template or the configuration set associated with a running environment.
// When describing the settings for the configuration set associated with a
// running environment, it is possible to receive two sets of setting
// descriptions. One is the deployed configuration set, and the other is a
// draft configuration of an environment that is either in the process of
// deployment or that failed to deploy.
func (c *ElasticBeanstalk) DescribeConfigurationSettings(req DescribeConfigurationSettingsMessage) (resp *DescribeConfigurationSettingsResult, err error) {
	resp = &DescribeConfigurationSettingsResult{}
	err = c.client.Do("DescribeConfigurationSettings", "POST", "/", req, resp)
	return
}

// DescribeEnvironmentResources is undocumented.
func (c *ElasticBeanstalk) DescribeEnvironmentResources(req DescribeEnvironmentResourcesMessage) (resp *DescribeEnvironmentResourcesResult, err error) {
	resp = &DescribeEnvironmentResourcesResult{}
	err = c.client.Do("DescribeEnvironmentResources", "POST", "/", req, resp)
	return
}

// DescribeEnvironments is undocumented.
func (c *ElasticBeanstalk) DescribeEnvironments(req DescribeEnvironmentsMessage) (resp *DescribeEnvironmentsResult, err error) {
	resp = &DescribeEnvironmentsResult{}
	err = c.client.Do("DescribeEnvironments", "POST", "/", req, resp)
	return
}

// DescribeEvents returns list of event descriptions matching criteria up
// to the last 6 weeks.
func (c *ElasticBeanstalk) DescribeEvents(req DescribeEventsMessage) (resp *DescribeEventsResult, err error) {
	resp = &DescribeEventsResult{}
	err = c.client.Do("DescribeEvents", "POST", "/", req, resp)
	return
}

// ListAvailableSolutionStacks returns a list of the available solution
// stack names.
func (c *ElasticBeanstalk) ListAvailableSolutionStacks() (resp *ListAvailableSolutionStacksResult, err error) {
	resp = &ListAvailableSolutionStacksResult{}
	err = c.client.Do("ListAvailableSolutionStacks", "POST", "/", nil, resp)
	return
}

// RebuildEnvironment deletes and recreates all of the AWS resources (for
// example: the Auto Scaling group, load balancer, etc.) for a specified
// environment and forces a restart.
func (c *ElasticBeanstalk) RebuildEnvironment(req RebuildEnvironmentMessage) (err error) {
	// NRE
	err = c.client.Do("RebuildEnvironment", "POST", "/", req, nil)
	return
}

// RequestEnvironmentInfo initiates a request to compile the specified type
// of information of the deployed environment. Setting the InfoType to tail
// compiles the last lines from the application server log files of every
// Amazon EC2 instance in your environment. Use RetrieveEnvironmentInfo to
// access the compiled information.
func (c *ElasticBeanstalk) RequestEnvironmentInfo(req RequestEnvironmentInfoMessage) (err error) {
	// NRE
	err = c.client.Do("RequestEnvironmentInfo", "POST", "/", req, nil)
	return
}

// RestartAppServer causes the environment to restart the application
// container server running on each Amazon EC2 instance.
func (c *ElasticBeanstalk) RestartAppServer(req RestartAppServerMessage) (err error) {
	// NRE
	err = c.client.Do("RestartAppServer", "POST", "/", req, nil)
	return
}

// RetrieveEnvironmentInfo retrieves the compiled information from a
// RequestEnvironmentInfo request.
func (c *ElasticBeanstalk) RetrieveEnvironmentInfo(req RetrieveEnvironmentInfoMessage) (resp *RetrieveEnvironmentInfoResult, err error) {
	resp = &RetrieveEnvironmentInfoResult{}
	err = c.client.Do("RetrieveEnvironmentInfo", "POST", "/", req, resp)
	return
}

// SwapEnvironmentCNAMEs is undocumented.
func (c *ElasticBeanstalk) SwapEnvironmentCNAMEs(req SwapEnvironmentCNAMEsMessage) (err error) {
	// NRE
	err = c.client.Do("SwapEnvironmentCNAMEs", "POST", "/", req, nil)
	return
}

// TerminateEnvironment is undocumented.
func (c *ElasticBeanstalk) TerminateEnvironment(req TerminateEnvironmentMessage) (resp *TerminateEnvironmentResult, err error) {
	resp = &TerminateEnvironmentResult{}
	err = c.client.Do("TerminateEnvironment", "POST", "/", req, resp)
	return
}

// UpdateApplication updates the specified application to have the
// specified properties.
func (c *ElasticBeanstalk) UpdateApplication(req UpdateApplicationMessage) (resp *UpdateApplicationResult, err error) {
	resp = &UpdateApplicationResult{}
	err = c.client.Do("UpdateApplication", "POST", "/", req, resp)
	return
}

// UpdateApplicationVersion updates the specified application version to
// have the specified properties.
func (c *ElasticBeanstalk) UpdateApplicationVersion(req UpdateApplicationVersionMessage) (resp *UpdateApplicationVersionResult, err error) {
	resp = &UpdateApplicationVersionResult{}
	err = c.client.Do("UpdateApplicationVersion", "POST", "/", req, resp)
	return
}

// UpdateConfigurationTemplate updates the specified configuration template
// to have the specified properties or configuration option values.
func (c *ElasticBeanstalk) UpdateConfigurationTemplate(req UpdateConfigurationTemplateMessage) (resp *UpdateConfigurationTemplateResult, err error) {
	resp = &UpdateConfigurationTemplateResult{}
	err = c.client.Do("UpdateConfigurationTemplate", "POST", "/", req, resp)
	return
}

// UpdateEnvironment updates the environment description, deploys a new
// application version, updates the configuration settings to an entirely
// new configuration template, or updates select configuration option
// values in the running environment. Attempting to update both the release
// and configuration is not allowed and AWS Elastic Beanstalk returns an
// InvalidParameterCombination error. When updating the configuration
// settings to a new template or individual settings, a draft configuration
// is created and DescribeConfigurationSettings for this environment
// returns two setting descriptions with different DeploymentStatus values.
func (c *ElasticBeanstalk) UpdateEnvironment(req UpdateEnvironmentMessage) (resp *UpdateEnvironmentResult, err error) {
	resp = &UpdateEnvironmentResult{}
	err = c.client.Do("UpdateEnvironment", "POST", "/", req, resp)
	return
}

// ValidateConfigurationSettings takes a set of configuration settings and
// either a configuration template or environment, and determines whether
// those values are valid. This action returns a list of messages
// indicating any errors or warnings associated with the selection of
// option values.
func (c *ElasticBeanstalk) ValidateConfigurationSettings(req ValidateConfigurationSettingsMessage) (resp *ValidateConfigurationSettingsResult, err error) {
	resp = &ValidateConfigurationSettingsResult{}
	err = c.client.Do("ValidateConfigurationSettings", "POST", "/", req, resp)
	return
}

// ApplicationDescription is undocumented.
type ApplicationDescription struct {
	ApplicationName        string    `xml:"ApplicationName"`
	ConfigurationTemplates []string  `xml:"ConfigurationTemplates>member"`
	DateCreated            time.Time `xml:"DateCreated"`
	DateUpdated            time.Time `xml:"DateUpdated"`
	Description            string    `xml:"Description"`
	Versions               []string  `xml:"Versions>member"`
}

// ApplicationDescriptionMessage is undocumented.
type ApplicationDescriptionMessage struct {
	Application ApplicationDescription `xml:"Application"`
}

// ApplicationDescriptionsMessage is undocumented.
type ApplicationDescriptionsMessage struct {
	Applications []ApplicationDescription `xml:"DescribeApplicationsResult>Applications>member"`
}

// ApplicationVersionDescription is undocumented.
type ApplicationVersionDescription struct {
	ApplicationName string     `xml:"ApplicationName"`
	DateCreated     time.Time  `xml:"DateCreated"`
	DateUpdated     time.Time  `xml:"DateUpdated"`
	Description     string     `xml:"Description"`
	SourceBundle    S3Location `xml:"SourceBundle"`
	VersionLabel    string     `xml:"VersionLabel"`
}

// ApplicationVersionDescriptionMessage is undocumented.
type ApplicationVersionDescriptionMessage struct {
	ApplicationVersion ApplicationVersionDescription `xml:"ApplicationVersion"`
}

// ApplicationVersionDescriptionsMessage is undocumented.
type ApplicationVersionDescriptionsMessage struct {
	ApplicationVersions []ApplicationVersionDescription `xml:"DescribeApplicationVersionsResult>ApplicationVersions>member"`
}

// AutoScalingGroup is undocumented.
type AutoScalingGroup struct {
	Name string `xml:"Name"`
}

// CheckDNSAvailabilityMessage is undocumented.
type CheckDNSAvailabilityMessage struct {
	CNAMEPrefix string `xml:"CNAMEPrefix"`
}

// CheckDNSAvailabilityResultMessage is undocumented.
type CheckDNSAvailabilityResultMessage struct {
	Available           bool   `xml:"CheckDNSAvailabilityResult>Available"`
	FullyQualifiedCNAME string `xml:"CheckDNSAvailabilityResult>FullyQualifiedCNAME"`
}

// ConfigurationOptionDescription is undocumented.
type ConfigurationOptionDescription struct {
	ChangeSeverity string                 `xml:"ChangeSeverity"`
	DefaultValue   string                 `xml:"DefaultValue"`
	MaxLength      int                    `xml:"MaxLength"`
	MaxValue       int                    `xml:"MaxValue"`
	MinValue       int                    `xml:"MinValue"`
	Name           string                 `xml:"Name"`
	Namespace      string                 `xml:"Namespace"`
	Regex          OptionRestrictionRegex `xml:"Regex"`
	UserDefined    bool                   `xml:"UserDefined"`
	ValueOptions   []string               `xml:"ValueOptions>member"`
	ValueType      string                 `xml:"ValueType"`
}

// ConfigurationOptionSetting is undocumented.
type ConfigurationOptionSetting struct {
	Namespace  string `xml:"Namespace"`
	OptionName string `xml:"OptionName"`
	Value      string `xml:"Value"`
}

// ConfigurationOptionsDescription is undocumented.
type ConfigurationOptionsDescription struct {
	Options           []ConfigurationOptionDescription `xml:"DescribeConfigurationOptionsResult>Options>member"`
	SolutionStackName string                           `xml:"DescribeConfigurationOptionsResult>SolutionStackName"`
}

// ConfigurationSettingsDescription is undocumented.
type ConfigurationSettingsDescription struct {
	ApplicationName   string                       `xml:"ApplicationName"`
	DateCreated       time.Time                    `xml:"DateCreated"`
	DateUpdated       time.Time                    `xml:"DateUpdated"`
	DeploymentStatus  string                       `xml:"DeploymentStatus"`
	Description       string                       `xml:"Description"`
	EnvironmentName   string                       `xml:"EnvironmentName"`
	OptionSettings    []ConfigurationOptionSetting `xml:"OptionSettings>member"`
	SolutionStackName string                       `xml:"SolutionStackName"`
	TemplateName      string                       `xml:"TemplateName"`
}

// ConfigurationSettingsDescriptions is undocumented.
type ConfigurationSettingsDescriptions struct {
	ConfigurationSettings []ConfigurationSettingsDescription `xml:"DescribeConfigurationSettingsResult>ConfigurationSettings>member"`
}

// ConfigurationSettingsValidationMessages is undocumented.
type ConfigurationSettingsValidationMessages struct {
	Messages []ValidationMessage `xml:"ValidateConfigurationSettingsResult>Messages>member"`
}

// CreateApplicationMessage is undocumented.
type CreateApplicationMessage struct {
	ApplicationName string `xml:"ApplicationName"`
	Description     string `xml:"Description"`
}

// CreateApplicationVersionMessage is undocumented.
type CreateApplicationVersionMessage struct {
	ApplicationName       string     `xml:"ApplicationName"`
	AutoCreateApplication bool       `xml:"AutoCreateApplication"`
	Description           string     `xml:"Description"`
	SourceBundle          S3Location `xml:"SourceBundle"`
	VersionLabel          string     `xml:"VersionLabel"`
}

// CreateConfigurationTemplateMessage is undocumented.
type CreateConfigurationTemplateMessage struct {
	ApplicationName     string                       `xml:"ApplicationName"`
	Description         string                       `xml:"Description"`
	EnvironmentID       string                       `xml:"EnvironmentId"`
	OptionSettings      []ConfigurationOptionSetting `xml:"OptionSettings>member"`
	SolutionStackName   string                       `xml:"SolutionStackName"`
	SourceConfiguration SourceConfiguration          `xml:"SourceConfiguration"`
	TemplateName        string                       `xml:"TemplateName"`
}

// CreateEnvironmentMessage is undocumented.
type CreateEnvironmentMessage struct {
	ApplicationName   string                       `xml:"ApplicationName"`
	CNAMEPrefix       string                       `xml:"CNAMEPrefix"`
	Description       string                       `xml:"Description"`
	EnvironmentName   string                       `xml:"EnvironmentName"`
	OptionSettings    []ConfigurationOptionSetting `xml:"OptionSettings>member"`
	OptionsToRemove   []OptionSpecification        `xml:"OptionsToRemove>member"`
	SolutionStackName string                       `xml:"SolutionStackName"`
	Tags              []Tag                        `xml:"Tags>member"`
	TemplateName      string                       `xml:"TemplateName"`
	Tier              EnvironmentTier              `xml:"Tier"`
	VersionLabel      string                       `xml:"VersionLabel"`
}

// CreateStorageLocationResultMessage is undocumented.
type CreateStorageLocationResultMessage struct {
	S3Bucket string `xml:"CreateStorageLocationResult>S3Bucket"`
}

// DeleteApplicationMessage is undocumented.
type DeleteApplicationMessage struct {
	ApplicationName     string `xml:"ApplicationName"`
	TerminateEnvByForce bool   `xml:"TerminateEnvByForce"`
}

// DeleteApplicationVersionMessage is undocumented.
type DeleteApplicationVersionMessage struct {
	ApplicationName    string `xml:"ApplicationName"`
	DeleteSourceBundle bool   `xml:"DeleteSourceBundle"`
	VersionLabel       string `xml:"VersionLabel"`
}

// DeleteConfigurationTemplateMessage is undocumented.
type DeleteConfigurationTemplateMessage struct {
	ApplicationName string `xml:"ApplicationName"`
	TemplateName    string `xml:"TemplateName"`
}

// DeleteEnvironmentConfigurationMessage is undocumented.
type DeleteEnvironmentConfigurationMessage struct {
	ApplicationName string `xml:"ApplicationName"`
	EnvironmentName string `xml:"EnvironmentName"`
}

// DescribeApplicationVersionsMessage is undocumented.
type DescribeApplicationVersionsMessage struct {
	ApplicationName string   `xml:"ApplicationName"`
	VersionLabels   []string `xml:"VersionLabels>member"`
}

// DescribeApplicationsMessage is undocumented.
type DescribeApplicationsMessage struct {
	ApplicationNames []string `xml:"ApplicationNames>member"`
}

// DescribeConfigurationOptionsMessage is undocumented.
type DescribeConfigurationOptionsMessage struct {
	ApplicationName   string                `xml:"ApplicationName"`
	EnvironmentName   string                `xml:"EnvironmentName"`
	Options           []OptionSpecification `xml:"Options>member"`
	SolutionStackName string                `xml:"SolutionStackName"`
	TemplateName      string                `xml:"TemplateName"`
}

// DescribeConfigurationSettingsMessage is undocumented.
type DescribeConfigurationSettingsMessage struct {
	ApplicationName string `xml:"ApplicationName"`
	EnvironmentName string `xml:"EnvironmentName"`
	TemplateName    string `xml:"TemplateName"`
}

// DescribeEnvironmentResourcesMessage is undocumented.
type DescribeEnvironmentResourcesMessage struct {
	EnvironmentID   string `xml:"EnvironmentId"`
	EnvironmentName string `xml:"EnvironmentName"`
}

// DescribeEnvironmentsMessage is undocumented.
type DescribeEnvironmentsMessage struct {
	ApplicationName       string    `xml:"ApplicationName"`
	EnvironmentIds        []string  `xml:"EnvironmentIds>member"`
	EnvironmentNames      []string  `xml:"EnvironmentNames>member"`
	IncludeDeleted        bool      `xml:"IncludeDeleted"`
	IncludedDeletedBackTo time.Time `xml:"IncludedDeletedBackTo"`
	VersionLabel          string    `xml:"VersionLabel"`
}

// DescribeEventsMessage is undocumented.
type DescribeEventsMessage struct {
	ApplicationName string    `xml:"ApplicationName"`
	EndTime         time.Time `xml:"EndTime"`
	EnvironmentID   string    `xml:"EnvironmentId"`
	EnvironmentName string    `xml:"EnvironmentName"`
	MaxRecords      int       `xml:"MaxRecords"`
	NextToken       string    `xml:"NextToken"`
	RequestID       string    `xml:"RequestId"`
	Severity        string    `xml:"Severity"`
	StartTime       time.Time `xml:"StartTime"`
	TemplateName    string    `xml:"TemplateName"`
	VersionLabel    string    `xml:"VersionLabel"`
}

// EnvironmentDescription is undocumented.
type EnvironmentDescription struct {
	ApplicationName   string                          `xml:"ApplicationName"`
	CNAME             string                          `xml:"CNAME"`
	DateCreated       time.Time                       `xml:"DateCreated"`
	DateUpdated       time.Time                       `xml:"DateUpdated"`
	Description       string                          `xml:"Description"`
	EndpointURL       string                          `xml:"EndpointURL"`
	EnvironmentID     string                          `xml:"EnvironmentId"`
	EnvironmentName   string                          `xml:"EnvironmentName"`
	Health            string                          `xml:"Health"`
	Resources         EnvironmentResourcesDescription `xml:"Resources"`
	SolutionStackName string                          `xml:"SolutionStackName"`
	Status            string                          `xml:"Status"`
	TemplateName      string                          `xml:"TemplateName"`
	Tier              EnvironmentTier                 `xml:"Tier"`
	VersionLabel      string                          `xml:"VersionLabel"`
}

// EnvironmentDescriptionsMessage is undocumented.
type EnvironmentDescriptionsMessage struct {
	Environments []EnvironmentDescription `xml:"DescribeEnvironmentsResult>Environments>member"`
}

// EnvironmentInfoDescription is undocumented.
type EnvironmentInfoDescription struct {
	Ec2InstanceID   string    `xml:"Ec2InstanceId"`
	InfoType        string    `xml:"InfoType"`
	Message         string    `xml:"Message"`
	SampleTimestamp time.Time `xml:"SampleTimestamp"`
}

// EnvironmentResourceDescription is undocumented.
type EnvironmentResourceDescription struct {
	AutoScalingGroups    []AutoScalingGroup    `xml:"AutoScalingGroups>member"`
	EnvironmentName      string                `xml:"EnvironmentName"`
	Instances            []Instance            `xml:"Instances>member"`
	LaunchConfigurations []LaunchConfiguration `xml:"LaunchConfigurations>member"`
	LoadBalancers        []LoadBalancer        `xml:"LoadBalancers>member"`
	Queues               []Queue               `xml:"Queues>member"`
	Triggers             []Trigger             `xml:"Triggers>member"`
}

// EnvironmentResourceDescriptionsMessage is undocumented.
type EnvironmentResourceDescriptionsMessage struct {
	EnvironmentResources EnvironmentResourceDescription `xml:"DescribeEnvironmentResourcesResult>EnvironmentResources"`
}

// EnvironmentResourcesDescription is undocumented.
type EnvironmentResourcesDescription struct {
	LoadBalancer LoadBalancerDescription `xml:"LoadBalancer"`
}

// EnvironmentTier is undocumented.
type EnvironmentTier struct {
	Name    string `xml:"Name"`
	Type    string `xml:"Type"`
	Version string `xml:"Version"`
}

// EventDescription is undocumented.
type EventDescription struct {
	ApplicationName string    `xml:"ApplicationName"`
	EnvironmentName string    `xml:"EnvironmentName"`
	EventDate       time.Time `xml:"EventDate"`
	Message         string    `xml:"Message"`
	RequestID       string    `xml:"RequestId"`
	Severity        string    `xml:"Severity"`
	TemplateName    string    `xml:"TemplateName"`
	VersionLabel    string    `xml:"VersionLabel"`
}

// EventDescriptionsMessage is undocumented.
type EventDescriptionsMessage struct {
	Events    []EventDescription `xml:"DescribeEventsResult>Events>member"`
	NextToken string             `xml:"DescribeEventsResult>NextToken"`
}

// Instance is undocumented.
type Instance struct {
	ID string `xml:"Id"`
}

// LaunchConfiguration is undocumented.
type LaunchConfiguration struct {
	Name string `xml:"Name"`
}

// ListAvailableSolutionStacksResultMessage is undocumented.
type ListAvailableSolutionStacksResultMessage struct {
	SolutionStackDetails []SolutionStackDescription `xml:"ListAvailableSolutionStacksResult>SolutionStackDetails>member"`
	SolutionStacks       []string                   `xml:"ListAvailableSolutionStacksResult>SolutionStacks>member"`
}

// Listener is undocumented.
type Listener struct {
	Port     int    `xml:"Port"`
	Protocol string `xml:"Protocol"`
}

// LoadBalancer is undocumented.
type LoadBalancer struct {
	Name string `xml:"Name"`
}

// LoadBalancerDescription is undocumented.
type LoadBalancerDescription struct {
	Domain           string     `xml:"Domain"`
	Listeners        []Listener `xml:"Listeners>member"`
	LoadBalancerName string     `xml:"LoadBalancerName"`
}

// OptionRestrictionRegex is undocumented.
type OptionRestrictionRegex struct {
	Label   string `xml:"Label"`
	Pattern string `xml:"Pattern"`
}

// OptionSpecification is undocumented.
type OptionSpecification struct {
	Namespace  string `xml:"Namespace"`
	OptionName string `xml:"OptionName"`
}

// Queue is undocumented.
type Queue struct {
	Name string `xml:"Name"`
	URL  string `xml:"URL"`
}

// RebuildEnvironmentMessage is undocumented.
type RebuildEnvironmentMessage struct {
	EnvironmentID   string `xml:"EnvironmentId"`
	EnvironmentName string `xml:"EnvironmentName"`
}

// RequestEnvironmentInfoMessage is undocumented.
type RequestEnvironmentInfoMessage struct {
	EnvironmentID   string `xml:"EnvironmentId"`
	EnvironmentName string `xml:"EnvironmentName"`
	InfoType        string `xml:"InfoType"`
}

// RestartAppServerMessage is undocumented.
type RestartAppServerMessage struct {
	EnvironmentID   string `xml:"EnvironmentId"`
	EnvironmentName string `xml:"EnvironmentName"`
}

// RetrieveEnvironmentInfoMessage is undocumented.
type RetrieveEnvironmentInfoMessage struct {
	EnvironmentID   string `xml:"EnvironmentId"`
	EnvironmentName string `xml:"EnvironmentName"`
	InfoType        string `xml:"InfoType"`
}

// RetrieveEnvironmentInfoResultMessage is undocumented.
type RetrieveEnvironmentInfoResultMessage struct {
	EnvironmentInfo []EnvironmentInfoDescription `xml:"RetrieveEnvironmentInfoResult>EnvironmentInfo>member"`
}

// S3Location is undocumented.
type S3Location struct {
	S3Bucket string `xml:"S3Bucket"`
	S3Key    string `xml:"S3Key"`
}

// SolutionStackDescription is undocumented.
type SolutionStackDescription struct {
	PermittedFileTypes []string `xml:"PermittedFileTypes>member"`
	SolutionStackName  string   `xml:"SolutionStackName"`
}

// SourceConfiguration is undocumented.
type SourceConfiguration struct {
	ApplicationName string `xml:"ApplicationName"`
	TemplateName    string `xml:"TemplateName"`
}

// SwapEnvironmentCNAMEsMessage is undocumented.
type SwapEnvironmentCNAMEsMessage struct {
	DestinationEnvironmentID   string `xml:"DestinationEnvironmentId"`
	DestinationEnvironmentName string `xml:"DestinationEnvironmentName"`
	SourceEnvironmentID        string `xml:"SourceEnvironmentId"`
	SourceEnvironmentName      string `xml:"SourceEnvironmentName"`
}

// Tag is undocumented.
type Tag struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

// TerminateEnvironmentMessage is undocumented.
type TerminateEnvironmentMessage struct {
	EnvironmentID      string `xml:"EnvironmentId"`
	EnvironmentName    string `xml:"EnvironmentName"`
	TerminateResources bool   `xml:"TerminateResources"`
}

// Trigger is undocumented.
type Trigger struct {
	Name string `xml:"Name"`
}

// UpdateApplicationMessage is undocumented.
type UpdateApplicationMessage struct {
	ApplicationName string `xml:"ApplicationName"`
	Description     string `xml:"Description"`
}

// UpdateApplicationVersionMessage is undocumented.
type UpdateApplicationVersionMessage struct {
	ApplicationName string `xml:"ApplicationName"`
	Description     string `xml:"Description"`
	VersionLabel    string `xml:"VersionLabel"`
}

// UpdateConfigurationTemplateMessage is undocumented.
type UpdateConfigurationTemplateMessage struct {
	ApplicationName string                       `xml:"ApplicationName"`
	Description     string                       `xml:"Description"`
	OptionSettings  []ConfigurationOptionSetting `xml:"OptionSettings>member"`
	OptionsToRemove []OptionSpecification        `xml:"OptionsToRemove>member"`
	TemplateName    string                       `xml:"TemplateName"`
}

// UpdateEnvironmentMessage is undocumented.
type UpdateEnvironmentMessage struct {
	Description     string                       `xml:"Description"`
	EnvironmentID   string                       `xml:"EnvironmentId"`
	EnvironmentName string                       `xml:"EnvironmentName"`
	OptionSettings  []ConfigurationOptionSetting `xml:"OptionSettings>member"`
	OptionsToRemove []OptionSpecification        `xml:"OptionsToRemove>member"`
	TemplateName    string                       `xml:"TemplateName"`
	Tier            EnvironmentTier              `xml:"Tier"`
	VersionLabel    string                       `xml:"VersionLabel"`
}

// ValidateConfigurationSettingsMessage is undocumented.
type ValidateConfigurationSettingsMessage struct {
	ApplicationName string                       `xml:"ApplicationName"`
	EnvironmentName string                       `xml:"EnvironmentName"`
	OptionSettings  []ConfigurationOptionSetting `xml:"OptionSettings>member"`
	TemplateName    string                       `xml:"TemplateName"`
}

// ValidationMessage is undocumented.
type ValidationMessage struct {
	Message    string `xml:"Message"`
	Namespace  string `xml:"Namespace"`
	OptionName string `xml:"OptionName"`
	Severity   string `xml:"Severity"`
}

// CheckDNSAvailabilityResult is a wrapper for CheckDNSAvailabilityResultMessage.
type CheckDNSAvailabilityResult struct {
	XMLName xml.Name `xml:"CheckDNSAvailabilityResponse"`

	Available           bool   `xml:"CheckDNSAvailabilityResult>Available"`
	FullyQualifiedCNAME string `xml:"CheckDNSAvailabilityResult>FullyQualifiedCNAME"`
}

// CreateApplicationResult is a wrapper for ApplicationDescriptionMessage.
type CreateApplicationResult struct {
	XMLName xml.Name `xml:"Response"`

	Application ApplicationDescription `xml:"CreateApplicationResult>Application"`
}

// CreateApplicationVersionResult is a wrapper for ApplicationVersionDescriptionMessage.
type CreateApplicationVersionResult struct {
	XMLName xml.Name `xml:"Response"`

	ApplicationVersion ApplicationVersionDescription `xml:"CreateApplicationVersionResult>ApplicationVersion"`
}

// CreateConfigurationTemplateResult is a wrapper for ConfigurationSettingsDescription.
type CreateConfigurationTemplateResult struct {
	XMLName xml.Name `xml:"Response"`

	ApplicationName   string                       `xml:"CreateConfigurationTemplateResult>ApplicationName"`
	DateCreated       time.Time                    `xml:"CreateConfigurationTemplateResult>DateCreated"`
	DateUpdated       time.Time                    `xml:"CreateConfigurationTemplateResult>DateUpdated"`
	DeploymentStatus  string                       `xml:"CreateConfigurationTemplateResult>DeploymentStatus"`
	Description       string                       `xml:"CreateConfigurationTemplateResult>Description"`
	EnvironmentName   string                       `xml:"CreateConfigurationTemplateResult>EnvironmentName"`
	OptionSettings    []ConfigurationOptionSetting `xml:"CreateConfigurationTemplateResult>OptionSettings>member"`
	SolutionStackName string                       `xml:"CreateConfigurationTemplateResult>SolutionStackName"`
	TemplateName      string                       `xml:"CreateConfigurationTemplateResult>TemplateName"`
}

// CreateEnvironmentResult is a wrapper for EnvironmentDescription.
type CreateEnvironmentResult struct {
	XMLName xml.Name `xml:"Response"`

	ApplicationName   string                          `xml:"CreateEnvironmentResult>ApplicationName"`
	CNAME             string                          `xml:"CreateEnvironmentResult>CNAME"`
	DateCreated       time.Time                       `xml:"CreateEnvironmentResult>DateCreated"`
	DateUpdated       time.Time                       `xml:"CreateEnvironmentResult>DateUpdated"`
	Description       string                          `xml:"CreateEnvironmentResult>Description"`
	EndpointURL       string                          `xml:"CreateEnvironmentResult>EndpointURL"`
	EnvironmentID     string                          `xml:"CreateEnvironmentResult>EnvironmentId"`
	EnvironmentName   string                          `xml:"CreateEnvironmentResult>EnvironmentName"`
	Health            string                          `xml:"CreateEnvironmentResult>Health"`
	Resources         EnvironmentResourcesDescription `xml:"CreateEnvironmentResult>Resources"`
	SolutionStackName string                          `xml:"CreateEnvironmentResult>SolutionStackName"`
	Status            string                          `xml:"CreateEnvironmentResult>Status"`
	TemplateName      string                          `xml:"CreateEnvironmentResult>TemplateName"`
	Tier              EnvironmentTier                 `xml:"CreateEnvironmentResult>Tier"`
	VersionLabel      string                          `xml:"CreateEnvironmentResult>VersionLabel"`
}

// CreateStorageLocationResult is a wrapper for CreateStorageLocationResultMessage.
type CreateStorageLocationResult struct {
	XMLName xml.Name `xml:"CreateStorageLocationResponse"`

	S3Bucket string `xml:"CreateStorageLocationResult>S3Bucket"`
}

// DescribeApplicationVersionsResult is a wrapper for ApplicationVersionDescriptionsMessage.
type DescribeApplicationVersionsResult struct {
	XMLName xml.Name `xml:"DescribeApplicationVersionsResponse"`

	ApplicationVersions []ApplicationVersionDescription `xml:"DescribeApplicationVersionsResult>ApplicationVersions>member"`
}

// DescribeApplicationsResult is a wrapper for ApplicationDescriptionsMessage.
type DescribeApplicationsResult struct {
	XMLName xml.Name `xml:"DescribeApplicationsResponse"`

	Applications []ApplicationDescription `xml:"DescribeApplicationsResult>Applications>member"`
}

// DescribeConfigurationOptionsResult is a wrapper for ConfigurationOptionsDescription.
type DescribeConfigurationOptionsResult struct {
	XMLName xml.Name `xml:"DescribeConfigurationOptionsResponse"`

	Options           []ConfigurationOptionDescription `xml:"DescribeConfigurationOptionsResult>Options>member"`
	SolutionStackName string                           `xml:"DescribeConfigurationOptionsResult>SolutionStackName"`
}

// DescribeConfigurationSettingsResult is a wrapper for ConfigurationSettingsDescriptions.
type DescribeConfigurationSettingsResult struct {
	XMLName xml.Name `xml:"DescribeConfigurationSettingsResponse"`

	ConfigurationSettings []ConfigurationSettingsDescription `xml:"DescribeConfigurationSettingsResult>ConfigurationSettings>member"`
}

// DescribeEnvironmentResourcesResult is a wrapper for EnvironmentResourceDescriptionsMessage.
type DescribeEnvironmentResourcesResult struct {
	XMLName xml.Name `xml:"DescribeEnvironmentResourcesResponse"`

	EnvironmentResources EnvironmentResourceDescription `xml:"DescribeEnvironmentResourcesResult>EnvironmentResources"`
}

// DescribeEnvironmentsResult is a wrapper for EnvironmentDescriptionsMessage.
type DescribeEnvironmentsResult struct {
	XMLName xml.Name `xml:"DescribeEnvironmentsResponse"`

	Environments []EnvironmentDescription `xml:"DescribeEnvironmentsResult>Environments>member"`
}

// DescribeEventsResult is a wrapper for EventDescriptionsMessage.
type DescribeEventsResult struct {
	XMLName xml.Name `xml:"DescribeEventsResponse"`

	Events    []EventDescription `xml:"DescribeEventsResult>Events>member"`
	NextToken string             `xml:"DescribeEventsResult>NextToken"`
}

// ListAvailableSolutionStacksResult is a wrapper for ListAvailableSolutionStacksResultMessage.
type ListAvailableSolutionStacksResult struct {
	XMLName xml.Name `xml:"ListAvailableSolutionStacksResponse"`

	SolutionStackDetails []SolutionStackDescription `xml:"ListAvailableSolutionStacksResult>SolutionStackDetails>member"`
	SolutionStacks       []string                   `xml:"ListAvailableSolutionStacksResult>SolutionStacks>member"`
}

// RetrieveEnvironmentInfoResult is a wrapper for RetrieveEnvironmentInfoResultMessage.
type RetrieveEnvironmentInfoResult struct {
	XMLName xml.Name `xml:"RetrieveEnvironmentInfoResponse"`

	EnvironmentInfo []EnvironmentInfoDescription `xml:"RetrieveEnvironmentInfoResult>EnvironmentInfo>member"`
}

// TerminateEnvironmentResult is a wrapper for EnvironmentDescription.
type TerminateEnvironmentResult struct {
	XMLName xml.Name `xml:"Response"`

	ApplicationName   string                          `xml:"TerminateEnvironmentResult>ApplicationName"`
	CNAME             string                          `xml:"TerminateEnvironmentResult>CNAME"`
	DateCreated       time.Time                       `xml:"TerminateEnvironmentResult>DateCreated"`
	DateUpdated       time.Time                       `xml:"TerminateEnvironmentResult>DateUpdated"`
	Description       string                          `xml:"TerminateEnvironmentResult>Description"`
	EndpointURL       string                          `xml:"TerminateEnvironmentResult>EndpointURL"`
	EnvironmentID     string                          `xml:"TerminateEnvironmentResult>EnvironmentId"`
	EnvironmentName   string                          `xml:"TerminateEnvironmentResult>EnvironmentName"`
	Health            string                          `xml:"TerminateEnvironmentResult>Health"`
	Resources         EnvironmentResourcesDescription `xml:"TerminateEnvironmentResult>Resources"`
	SolutionStackName string                          `xml:"TerminateEnvironmentResult>SolutionStackName"`
	Status            string                          `xml:"TerminateEnvironmentResult>Status"`
	TemplateName      string                          `xml:"TerminateEnvironmentResult>TemplateName"`
	Tier              EnvironmentTier                 `xml:"TerminateEnvironmentResult>Tier"`
	VersionLabel      string                          `xml:"TerminateEnvironmentResult>VersionLabel"`
}

// UpdateApplicationResult is a wrapper for ApplicationDescriptionMessage.
type UpdateApplicationResult struct {
	XMLName xml.Name `xml:"Response"`

	Application ApplicationDescription `xml:"UpdateApplicationResult>Application"`
}

// UpdateApplicationVersionResult is a wrapper for ApplicationVersionDescriptionMessage.
type UpdateApplicationVersionResult struct {
	XMLName xml.Name `xml:"Response"`

	ApplicationVersion ApplicationVersionDescription `xml:"UpdateApplicationVersionResult>ApplicationVersion"`
}

// UpdateConfigurationTemplateResult is a wrapper for ConfigurationSettingsDescription.
type UpdateConfigurationTemplateResult struct {
	XMLName xml.Name `xml:"Response"`

	ApplicationName   string                       `xml:"UpdateConfigurationTemplateResult>ApplicationName"`
	DateCreated       time.Time                    `xml:"UpdateConfigurationTemplateResult>DateCreated"`
	DateUpdated       time.Time                    `xml:"UpdateConfigurationTemplateResult>DateUpdated"`
	DeploymentStatus  string                       `xml:"UpdateConfigurationTemplateResult>DeploymentStatus"`
	Description       string                       `xml:"UpdateConfigurationTemplateResult>Description"`
	EnvironmentName   string                       `xml:"UpdateConfigurationTemplateResult>EnvironmentName"`
	OptionSettings    []ConfigurationOptionSetting `xml:"UpdateConfigurationTemplateResult>OptionSettings>member"`
	SolutionStackName string                       `xml:"UpdateConfigurationTemplateResult>SolutionStackName"`
	TemplateName      string                       `xml:"UpdateConfigurationTemplateResult>TemplateName"`
}

// UpdateEnvironmentResult is a wrapper for EnvironmentDescription.
type UpdateEnvironmentResult struct {
	XMLName xml.Name `xml:"Response"`

	ApplicationName   string                          `xml:"UpdateEnvironmentResult>ApplicationName"`
	CNAME             string                          `xml:"UpdateEnvironmentResult>CNAME"`
	DateCreated       time.Time                       `xml:"UpdateEnvironmentResult>DateCreated"`
	DateUpdated       time.Time                       `xml:"UpdateEnvironmentResult>DateUpdated"`
	Description       string                          `xml:"UpdateEnvironmentResult>Description"`
	EndpointURL       string                          `xml:"UpdateEnvironmentResult>EndpointURL"`
	EnvironmentID     string                          `xml:"UpdateEnvironmentResult>EnvironmentId"`
	EnvironmentName   string                          `xml:"UpdateEnvironmentResult>EnvironmentName"`
	Health            string                          `xml:"UpdateEnvironmentResult>Health"`
	Resources         EnvironmentResourcesDescription `xml:"UpdateEnvironmentResult>Resources"`
	SolutionStackName string                          `xml:"UpdateEnvironmentResult>SolutionStackName"`
	Status            string                          `xml:"UpdateEnvironmentResult>Status"`
	TemplateName      string                          `xml:"UpdateEnvironmentResult>TemplateName"`
	Tier              EnvironmentTier                 `xml:"UpdateEnvironmentResult>Tier"`
	VersionLabel      string                          `xml:"UpdateEnvironmentResult>VersionLabel"`
}

// ValidateConfigurationSettingsResult is a wrapper for ConfigurationSettingsValidationMessages.
type ValidateConfigurationSettingsResult struct {
	XMLName xml.Name `xml:"ValidateConfigurationSettingsResponse"`

	Messages []ValidationMessage `xml:"ValidateConfigurationSettingsResult>Messages>member"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
