// Package config provides a client for AWS Config.
package config

import (
	"fmt"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
)

// Config is a client for AWS Config.
type Config struct {
	client aws.Client
}

// New returns a new Config client.
func New(key, secret, region string, client *http.Client) *Config {
	if client == nil {
		client = http.DefaultClient
	}

	return &Config{
		client: &aws.JSONClient{
			Client:       client,
			Region:       region,
			Endpoint:     fmt.Sprintf("https://config.%s.amazonaws.com", region),
			Prefix:       "config",
			Key:          key,
			Secret:       secret,
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

type ConfigExportDeliveryInfo struct {
	LastAttemptTime    time.Time `json:"lastAttemptTime,omitempty"`
	LastErrorCode      string    `json:"lastErrorCode,omitempty"`
	LastErrorMessage   string    `json:"lastErrorMessage,omitempty"`
	LastStatus         string    `json:"lastStatus,omitempty"`
	LastSuccessfulTime time.Time `json:"lastSuccessfulTime,omitempty"`
}

type ConfigStreamDeliveryInfo struct {
	LastErrorCode        string    `json:"lastErrorCode,omitempty"`
	LastErrorMessage     string    `json:"lastErrorMessage,omitempty"`
	LastStatus           string    `json:"lastStatus,omitempty"`
	LastStatusChangeTime time.Time `json:"lastStatusChangeTime,omitempty"`
}

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

type ConfigurationRecorder struct {
	Name    string `json:"name,omitempty"`
	RoleARN string `json:"roleARN,omitempty"`
}

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

type DeleteDeliveryChannelRequest struct {
	DeliveryChannelName string `json:"DeliveryChannelName"`
}

type DeliverConfigSnapshotRequest struct {
	DeliveryChannelName string `json:"deliveryChannelName"`
}

type DeliverConfigSnapshotResponse struct {
	ConfigSnapshotID string `json:"configSnapshotId,omitempty"`
}

type DeliveryChannel struct {
	Name         string `json:"name,omitempty"`
	S3BucketName string `json:"s3BucketName,omitempty"`
	S3KeyPrefix  string `json:"s3KeyPrefix,omitempty"`
	SnsTopicARN  string `json:"snsTopicARN,omitempty"`
}

type DeliveryChannelStatus struct {
	ConfigHistoryDeliveryInfo  ConfigExportDeliveryInfo `json:"configHistoryDeliveryInfo,omitempty"`
	ConfigSnapshotDeliveryInfo ConfigExportDeliveryInfo `json:"configSnapshotDeliveryInfo,omitempty"`
	ConfigStreamDeliveryInfo   ConfigStreamDeliveryInfo `json:"configStreamDeliveryInfo,omitempty"`
	Name                       string                   `json:"name,omitempty"`
}

type DescribeConfigurationRecorderStatusRequest struct {
	ConfigurationRecorderNames []string `json:"ConfigurationRecorderNames,omitempty"`
}

type DescribeConfigurationRecorderStatusResponse struct {
	ConfigurationRecordersStatus []ConfigurationRecorderStatus `json:"ConfigurationRecordersStatus,omitempty"`
}

type DescribeConfigurationRecordersRequest struct {
	ConfigurationRecorderNames []string `json:"ConfigurationRecorderNames,omitempty"`
}

type DescribeConfigurationRecordersResponse struct {
	ConfigurationRecorders []ConfigurationRecorder `json:"ConfigurationRecorders,omitempty"`
}

type DescribeDeliveryChannelStatusRequest struct {
	DeliveryChannelNames []string `json:"DeliveryChannelNames,omitempty"`
}

type DescribeDeliveryChannelStatusResponse struct {
	DeliveryChannelsStatus []DeliveryChannelStatus `json:"DeliveryChannelsStatus,omitempty"`
}

type DescribeDeliveryChannelsRequest struct {
	DeliveryChannelNames []string `json:"DeliveryChannelNames,omitempty"`
}

type DescribeDeliveryChannelsResponse struct {
	DeliveryChannels []DeliveryChannel `json:"DeliveryChannels,omitempty"`
}

type GetResourceConfigHistoryRequest struct {
	ChronologicalOrder string    `json:"chronologicalOrder,omitempty"`
	EarlierTime        time.Time `json:"earlierTime,omitempty"`
	LaterTime          time.Time `json:"laterTime,omitempty"`
	Limit              int       `json:"limit,omitempty"`
	NextToken          string    `json:"nextToken,omitempty"`
	ResourceID         string    `json:"resourceId"`
	ResourceType       string    `json:"resourceType"`
}

type GetResourceConfigHistoryResponse struct {
	ConfigurationItems []ConfigurationItem `json:"configurationItems,omitempty"`
	NextToken          string              `json:"nextToken,omitempty"`
}

type PutConfigurationRecorderRequest struct {
	ConfigurationRecorder ConfigurationRecorder `json:"ConfigurationRecorder"`
}

type PutDeliveryChannelRequest struct {
	DeliveryChannel DeliveryChannel `json:"DeliveryChannel"`
}

type Relationship struct {
	RelationshipName string `json:"relationshipName,omitempty"`
	ResourceID       string `json:"resourceId,omitempty"`
	ResourceType     string `json:"resourceType,omitempty"`
}

type StartConfigurationRecorderRequest struct {
	ConfigurationRecorderName string `json:"ConfigurationRecorderName"`
}

type StopConfigurationRecorderRequest struct {
	ConfigurationRecorderName string `json:"ConfigurationRecorderName"`
}

var _ time.Time // to avoid errors if the time package isn't referenced
