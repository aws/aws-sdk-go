// Package config provides a client for AWS Config.
package config

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// Config is a client for AWS Config.
type Config struct {
	client *aws.JSONClient
}

// New returns a new Config client.
func New(creds aws.Credentials, region string, client *http.Client) *Config {
	if client == nil {
		client = http.DefaultClient
	}

	service := "config"
	endpoint, service, region := endpoints.Lookup("config", region)

	return &Config{
		client: &aws.JSONClient{
			Credentials:  creds,
			Service:      service,
			Region:       region,
			Client:       client,
			Endpoint:     endpoint,
			JSONVersion:  "1.1",
			TargetPrefix: "StarlingDoveService",
		},
	}
}

// DeleteDeliveryChannel deletes the specified delivery channel. The
// delivery channel cannot be deleted if it is the only delivery channel
// and the configuration recorder is still running. To delete the delivery
// channel, stop the running configuration recorder using the
// StopConfigurationRecorder action.
func (c *Config) DeleteDeliveryChannel(req DeleteDeliveryChannelRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteDeliveryChannel", "POST", "/", req, nil)
	return
}

// DeliverConfigSnapshot schedules delivery of a configuration snapshot to
// the Amazon S3 bucket in the specified delivery channel. After the
// delivery has started, AWS Config sends following notifications using an
// Amazon SNS topic that you have specified. Notification of starting the
// delivery. Notification of delivery completed, if the delivery was
// successfully completed. Notification of delivery failure, if the
// delivery failed to complete.
func (c *Config) DeliverConfigSnapshot(req DeliverConfigSnapshotRequest) (resp *DeliverConfigSnapshotResponse, err error) {
	resp = &DeliverConfigSnapshotResponse{}
	err = c.client.Do("DeliverConfigSnapshot", "POST", "/", req, resp)
	return
}

// DescribeConfigurationRecorderStatus returns the current status of the
// specified configuration recorder. If a configuration recorder is not
// specified, this action returns the status of all configuration recorder
// associated with the account.
func (c *Config) DescribeConfigurationRecorderStatus(req DescribeConfigurationRecorderStatusRequest) (resp *DescribeConfigurationRecorderStatusResponse, err error) {
	resp = &DescribeConfigurationRecorderStatusResponse{}
	err = c.client.Do("DescribeConfigurationRecorderStatus", "POST", "/", req, resp)
	return
}

// DescribeConfigurationRecorders returns the name of one or more specified
// configuration recorders. If the recorder name is not specified, this
// action returns the names of all the configuration recorders associated
// with the account.
func (c *Config) DescribeConfigurationRecorders(req DescribeConfigurationRecordersRequest) (resp *DescribeConfigurationRecordersResponse, err error) {
	resp = &DescribeConfigurationRecordersResponse{}
	err = c.client.Do("DescribeConfigurationRecorders", "POST", "/", req, resp)
	return
}

// DescribeDeliveryChannelStatus returns the current status of the
// specified delivery channel. If a delivery channel is not specified, this
// action returns the current status of all delivery channels associated
// with the account.
func (c *Config) DescribeDeliveryChannelStatus(req DescribeDeliveryChannelStatusRequest) (resp *DescribeDeliveryChannelStatusResponse, err error) {
	resp = &DescribeDeliveryChannelStatusResponse{}
	err = c.client.Do("DescribeDeliveryChannelStatus", "POST", "/", req, resp)
	return
}

// DescribeDeliveryChannels returns details about the specified delivery
// channel. If a delivery channel is not specified, this action returns the
// details of all delivery channels associated with the account.
func (c *Config) DescribeDeliveryChannels(req DescribeDeliveryChannelsRequest) (resp *DescribeDeliveryChannelsResponse, err error) {
	resp = &DescribeDeliveryChannelsResponse{}
	err = c.client.Do("DescribeDeliveryChannels", "POST", "/", req, resp)
	return
}

// GetResourceConfigHistory returns a list of configuration items for the
// specified resource. The list contains details about each state of the
// resource during the specified time interval. You can specify a limit on
// the number of results returned on the page. If a limit is specified, a
// nextToken is returned as part of the result that you can use to continue
// this request.
func (c *Config) GetResourceConfigHistory(req GetResourceConfigHistoryRequest) (resp *GetResourceConfigHistoryResponse, err error) {
	resp = &GetResourceConfigHistoryResponse{}
	err = c.client.Do("GetResourceConfigHistory", "POST", "/", req, resp)
	return
}

// PutConfigurationRecorder creates a new configuration recorder to record
// the resource configurations. You can use this action to change the role
// roleARN ) of an existing recorder. To change the role, call the action
// on the existing configuration recorder and specify a role.
func (c *Config) PutConfigurationRecorder(req PutConfigurationRecorderRequest) (err error) {
	// NRE
	err = c.client.Do("PutConfigurationRecorder", "POST", "/", req, nil)
	return
}

// PutDeliveryChannel creates a new delivery channel object to deliver the
// configuration information to an Amazon S3 bucket, and to an Amazon SNS
// topic. You can use this action to change the Amazon S3 bucket or an
// Amazon SNS topic of the existing delivery channel. To change the Amazon
// S3 bucket or an Amazon SNS topic, call this action and specify the
// changed values for the S3 bucket and the SNS topic. If you specify a
// different value for either the S3 bucket or the SNS topic, this action
// will keep the existing value for the parameter that is not changed.
func (c *Config) PutDeliveryChannel(req PutDeliveryChannelRequest) (err error) {
	// NRE
	err = c.client.Do("PutDeliveryChannel", "POST", "/", req, nil)
	return
}

// StartConfigurationRecorder starts recording configurations of all the
// resources associated with the account. You must have created at least
// one delivery channel to successfully start the configuration recorder.
func (c *Config) StartConfigurationRecorder(req StartConfigurationRecorderRequest) (err error) {
	// NRE
	err = c.client.Do("StartConfigurationRecorder", "POST", "/", req, nil)
	return
}

// StopConfigurationRecorder stops recording configurations of all the
// resources associated with the account.
func (c *Config) StopConfigurationRecorder(req StopConfigurationRecorderRequest) (err error) {
	// NRE
	err = c.client.Do("StopConfigurationRecorder", "POST", "/", req, nil)
	return
}

// ConfigExportDeliveryInfo is undocumented.
type ConfigExportDeliveryInfo struct {
	LastAttemptTime    time.Time `json:"lastAttemptTime,omitempty"`
	LastErrorCode      string    `json:"lastErrorCode,omitempty"`
	LastErrorMessage   string    `json:"lastErrorMessage,omitempty"`
	LastStatus         string    `json:"lastStatus,omitempty"`
	LastSuccessfulTime time.Time `json:"lastSuccessfulTime,omitempty"`
}

// ConfigStreamDeliveryInfo is undocumented.
type ConfigStreamDeliveryInfo struct {
	LastErrorCode        string    `json:"lastErrorCode,omitempty"`
	LastErrorMessage     string    `json:"lastErrorMessage,omitempty"`
	LastStatus           string    `json:"lastStatus,omitempty"`
	LastStatusChangeTime time.Time `json:"lastStatusChangeTime,omitempty"`
}

// ConfigurationItem is undocumented.
type ConfigurationItem struct {
	AccountID                    string            `json:"accountId,omitempty"`
	ARN                          string            `json:"arn,omitempty"`
	AvailabilityZone             string            `json:"availabilityZone,omitempty"`
	Configuration                string            `json:"configuration,omitempty"`
	ConfigurationItemCaptureTime time.Time         `json:"configurationItemCaptureTime,omitempty"`
	ConfigurationItemMD5Hash     string            `json:"configurationItemMD5Hash,omitempty"`
	ConfigurationItemStatus      string            `json:"configurationItemStatus,omitempty"`
	ConfigurationStateID         string            `json:"configurationStateId,omitempty"`
	RelatedEvents                []string          `json:"relatedEvents,omitempty"`
	Relationships                []Relationship    `json:"relationships,omitempty"`
	ResourceCreationTime         time.Time         `json:"resourceCreationTime,omitempty"`
	ResourceID                   string            `json:"resourceId,omitempty"`
	ResourceType                 string            `json:"resourceType,omitempty"`
	Tags                         map[string]string `json:"tags,omitempty"`
	Version                      string            `json:"version,omitempty"`
}

// ConfigurationRecorder is undocumented.
type ConfigurationRecorder struct {
	Name    string `json:"name,omitempty"`
	RoleARN string `json:"roleARN,omitempty"`
}

// ConfigurationRecorderStatus is undocumented.
type ConfigurationRecorderStatus struct {
	LastErrorCode        string    `json:"lastErrorCode,omitempty"`
	LastErrorMessage     string    `json:"lastErrorMessage,omitempty"`
	LastStartTime        time.Time `json:"lastStartTime,omitempty"`
	LastStatus           string    `json:"lastStatus,omitempty"`
	LastStatusChangeTime time.Time `json:"lastStatusChangeTime,omitempty"`
	LastStopTime         time.Time `json:"lastStopTime,omitempty"`
	Name                 string    `json:"name,omitempty"`
	Recording            bool      `json:"recording,omitempty"`
}

// DeleteDeliveryChannelRequest is undocumented.
type DeleteDeliveryChannelRequest struct {
	DeliveryChannelName string `json:"DeliveryChannelName"`
}

// DeliverConfigSnapshotRequest is undocumented.
type DeliverConfigSnapshotRequest struct {
	DeliveryChannelName string `json:"deliveryChannelName"`
}

// DeliverConfigSnapshotResponse is undocumented.
type DeliverConfigSnapshotResponse struct {
	ConfigSnapshotID string `json:"configSnapshotId,omitempty"`
}

// DeliveryChannel is undocumented.
type DeliveryChannel struct {
	Name         string `json:"name,omitempty"`
	S3BucketName string `json:"s3BucketName,omitempty"`
	S3KeyPrefix  string `json:"s3KeyPrefix,omitempty"`
	SnsTopicARN  string `json:"snsTopicARN,omitempty"`
}

// DeliveryChannelStatus is undocumented.
type DeliveryChannelStatus struct {
	ConfigHistoryDeliveryInfo  ConfigExportDeliveryInfo `json:"configHistoryDeliveryInfo,omitempty"`
	ConfigSnapshotDeliveryInfo ConfigExportDeliveryInfo `json:"configSnapshotDeliveryInfo,omitempty"`
	ConfigStreamDeliveryInfo   ConfigStreamDeliveryInfo `json:"configStreamDeliveryInfo,omitempty"`
	Name                       string                   `json:"name,omitempty"`
}

// DescribeConfigurationRecorderStatusRequest is undocumented.
type DescribeConfigurationRecorderStatusRequest struct {
	ConfigurationRecorderNames []string `json:"ConfigurationRecorderNames,omitempty"`
}

// DescribeConfigurationRecorderStatusResponse is undocumented.
type DescribeConfigurationRecorderStatusResponse struct {
	ConfigurationRecordersStatus []ConfigurationRecorderStatus `json:"ConfigurationRecordersStatus,omitempty"`
}

// DescribeConfigurationRecordersRequest is undocumented.
type DescribeConfigurationRecordersRequest struct {
	ConfigurationRecorderNames []string `json:"ConfigurationRecorderNames,omitempty"`
}

// DescribeConfigurationRecordersResponse is undocumented.
type DescribeConfigurationRecordersResponse struct {
	ConfigurationRecorders []ConfigurationRecorder `json:"ConfigurationRecorders,omitempty"`
}

// DescribeDeliveryChannelStatusRequest is undocumented.
type DescribeDeliveryChannelStatusRequest struct {
	DeliveryChannelNames []string `json:"DeliveryChannelNames,omitempty"`
}

// DescribeDeliveryChannelStatusResponse is undocumented.
type DescribeDeliveryChannelStatusResponse struct {
	DeliveryChannelsStatus []DeliveryChannelStatus `json:"DeliveryChannelsStatus,omitempty"`
}

// DescribeDeliveryChannelsRequest is undocumented.
type DescribeDeliveryChannelsRequest struct {
	DeliveryChannelNames []string `json:"DeliveryChannelNames,omitempty"`
}

// DescribeDeliveryChannelsResponse is undocumented.
type DescribeDeliveryChannelsResponse struct {
	DeliveryChannels []DeliveryChannel `json:"DeliveryChannels,omitempty"`
}

// GetResourceConfigHistoryRequest is undocumented.
type GetResourceConfigHistoryRequest struct {
	ChronologicalOrder string    `json:"chronologicalOrder,omitempty"`
	EarlierTime        time.Time `json:"earlierTime,omitempty"`
	LaterTime          time.Time `json:"laterTime,omitempty"`
	Limit              int       `json:"limit,omitempty"`
	NextToken          string    `json:"nextToken,omitempty"`
	ResourceID         string    `json:"resourceId"`
	ResourceType       string    `json:"resourceType"`
}

// GetResourceConfigHistoryResponse is undocumented.
type GetResourceConfigHistoryResponse struct {
	ConfigurationItems []ConfigurationItem `json:"configurationItems,omitempty"`
	NextToken          string              `json:"nextToken,omitempty"`
}

// PutConfigurationRecorderRequest is undocumented.
type PutConfigurationRecorderRequest struct {
	ConfigurationRecorder ConfigurationRecorder `json:"ConfigurationRecorder"`
}

// PutDeliveryChannelRequest is undocumented.
type PutDeliveryChannelRequest struct {
	DeliveryChannel DeliveryChannel `json:"DeliveryChannel"`
}

// Relationship is undocumented.
type Relationship struct {
	RelationshipName string `json:"relationshipName,omitempty"`
	ResourceID       string `json:"resourceId,omitempty"`
	ResourceType     string `json:"resourceType,omitempty"`
}

// StartConfigurationRecorderRequest is undocumented.
type StartConfigurationRecorderRequest struct {
	ConfigurationRecorderName string `json:"ConfigurationRecorderName"`
}

// StopConfigurationRecorderRequest is undocumented.
type StopConfigurationRecorderRequest struct {
	ConfigurationRecorderName string `json:"ConfigurationRecorderName"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
