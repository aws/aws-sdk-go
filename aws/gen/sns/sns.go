// Package sns provides a client for Amazon Simple Notification Service.
package sns

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// SNS is a client for Amazon Simple Notification Service.
type SNS struct {
	client *aws.QueryClient
}

// New returns a new SNS client.
func New(creds aws.Credentials, region string, client *http.Client) *SNS {
	if client == nil {
		client = http.DefaultClient
	}

	service := "sns"
	endpoint, service, region := endpoints.Lookup("sns", region)

	return &SNS{
		client: &aws.QueryClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2010-03-31",
		},
	}
}

// AddPermission adds a statement to a topic's access control policy,
// granting access for the specified AWS accounts to the specified actions.
func (c *SNS) AddPermission(req AddPermissionInput) (err error) {
	// NRE
	err = c.client.Do("AddPermission", "POST", "/", req, nil)
	return
}

// ConfirmSubscription verifies an endpoint owner's intent to receive
// messages by validating the token sent to the endpoint by an earlier
// Subscribe action. If the token is valid, the action creates a new
// subscription and returns its Amazon Resource Name This call requires an
// AWS signature only when the AuthenticateOnUnsubscribe flag is set to
// "true".
func (c *SNS) ConfirmSubscription(req ConfirmSubscriptionInput) (resp *ConfirmSubscriptionResult, err error) {
	resp = &ConfirmSubscriptionResult{}
	err = c.client.Do("ConfirmSubscription", "POST", "/", req, resp)
	return
}

// CreatePlatformApplication creates a platform application object for one
// of the supported push notification services, such as and to which
// devices and mobile apps may register. You must specify PlatformPrincipal
// and PlatformCredential attributes when using the
// CreatePlatformApplication action. The PlatformPrincipal is received from
// the notification service. For PlatformPrincipal is certificate". For
// PlatformPrincipal is not applicable. For PlatformPrincipal is "client
// id". The PlatformCredential is also received from the notification
// service. For PlatformCredential is "private key". For PlatformCredential
// is key". For PlatformCredential is "client secret". The
// PlatformApplicationArn that is returned when using
// CreatePlatformApplication is then used as an attribute for the
// CreatePlatformEndpoint action. For more information, see Using Amazon
// SNS Mobile Push Notifications .
func (c *SNS) CreatePlatformApplication(req CreatePlatformApplicationInput) (resp *CreatePlatformApplicationResult, err error) {
	resp = &CreatePlatformApplicationResult{}
	err = c.client.Do("CreatePlatformApplication", "POST", "/", req, resp)
	return
}

// CreatePlatformEndpoint creates an endpoint for a device and mobile app
// on one of the supported push notification services, such as GCM and
// CreatePlatformEndpoint requires the PlatformApplicationArn that is
// returned from CreatePlatformApplication . The EndpointArn that is
// returned when using CreatePlatformEndpoint can then be used by the
// Publish action to send a message to a mobile app or by the Subscribe
// action for subscription to a topic. The CreatePlatformEndpoint action is
// idempotent, so if the requester already owns an endpoint with the same
// device token and attributes, that endpoint's ARN is returned without
// creating a new endpoint. For more information, see Using Amazon SNS
// Mobile Push Notifications . When using CreatePlatformEndpoint with
// Baidu, two attributes must be provided: ChannelId and UserId. The token
// field must also contain the ChannelId. For more information, see
// Creating an Amazon SNS Endpoint for Baidu .
func (c *SNS) CreatePlatformEndpoint(req CreatePlatformEndpointInput) (resp *CreatePlatformEndpointResult, err error) {
	resp = &CreatePlatformEndpointResult{}
	err = c.client.Do("CreatePlatformEndpoint", "POST", "/", req, resp)
	return
}

// CreateTopic creates a topic to which notifications can be published.
// Users can create at most 3000 topics. For more information, see
// http://aws.amazon.com/sns . This action is idempotent, so if the
// requester already owns a topic with the specified name, that topic's ARN
// is returned without creating a new topic.
func (c *SNS) CreateTopic(req CreateTopicInput) (resp *CreateTopicResult, err error) {
	resp = &CreateTopicResult{}
	err = c.client.Do("CreateTopic", "POST", "/", req, resp)
	return
}

// DeleteEndpoint deletes the endpoint from Amazon This action is
// idempotent. For more information, see Using Amazon SNS Mobile Push
// Notifications .
func (c *SNS) DeleteEndpoint(req DeleteEndpointInput) (err error) {
	// NRE
	err = c.client.Do("DeleteEndpoint", "POST", "/", req, nil)
	return
}

// DeletePlatformApplication deletes a platform application object for one
// of the supported push notification services, such as and For more
// information, see Using Amazon SNS Mobile Push Notifications .
func (c *SNS) DeletePlatformApplication(req DeletePlatformApplicationInput) (err error) {
	// NRE
	err = c.client.Do("DeletePlatformApplication", "POST", "/", req, nil)
	return
}

// DeleteTopic deletes a topic and all its subscriptions. Deleting a topic
// might prevent some messages previously sent to the topic from being
// delivered to subscribers. This action is idempotent, so deleting a topic
// that does not exist does not result in an error.
func (c *SNS) DeleteTopic(req DeleteTopicInput) (err error) {
	// NRE
	err = c.client.Do("DeleteTopic", "POST", "/", req, nil)
	return
}

// GetEndpointAttributes retrieves the endpoint attributes for a device on
// one of the supported push notification services, such as GCM and For
// more information, see Using Amazon SNS Mobile Push Notifications .
func (c *SNS) GetEndpointAttributes(req GetEndpointAttributesInput) (resp *GetEndpointAttributesResult, err error) {
	resp = &GetEndpointAttributesResult{}
	err = c.client.Do("GetEndpointAttributes", "POST", "/", req, resp)
	return
}

// GetPlatformApplicationAttributes retrieves the attributes of the
// platform application object for the supported push notification
// services, such as and For more information, see Using Amazon SNS Mobile
// Push Notifications .
func (c *SNS) GetPlatformApplicationAttributes(req GetPlatformApplicationAttributesInput) (resp *GetPlatformApplicationAttributesResult, err error) {
	resp = &GetPlatformApplicationAttributesResult{}
	err = c.client.Do("GetPlatformApplicationAttributes", "POST", "/", req, resp)
	return
}

// GetSubscriptionAttributes is undocumented.
func (c *SNS) GetSubscriptionAttributes(req GetSubscriptionAttributesInput) (resp *GetSubscriptionAttributesResult, err error) {
	resp = &GetSubscriptionAttributesResult{}
	err = c.client.Do("GetSubscriptionAttributes", "POST", "/", req, resp)
	return
}

// GetTopicAttributes returns all of the properties of a topic. Topic
// properties returned might differ based on the authorization of the user.
func (c *SNS) GetTopicAttributes(req GetTopicAttributesInput) (resp *GetTopicAttributesResult, err error) {
	resp = &GetTopicAttributesResult{}
	err = c.client.Do("GetTopicAttributes", "POST", "/", req, resp)
	return
}

// ListEndpointsByPlatformApplication lists the endpoints and endpoint
// attributes for devices in a supported push notification service, such as
// GCM and The results for ListEndpointsByPlatformApplication are paginated
// and return a limited list of endpoints, up to 100. If additional records
// are available after the first page results, then a NextToken string will
// be returned. To receive the next page, you call
// ListEndpointsByPlatformApplication again using the NextToken string
// received from the previous call. When there are no more records to
// return, NextToken will be null. For more information, see Using Amazon
// SNS Mobile Push Notifications .
func (c *SNS) ListEndpointsByPlatformApplication(req ListEndpointsByPlatformApplicationInput) (resp *ListEndpointsByPlatformApplicationResult, err error) {
	resp = &ListEndpointsByPlatformApplicationResult{}
	err = c.client.Do("ListEndpointsByPlatformApplication", "POST", "/", req, resp)
	return
}

// ListPlatformApplications lists the platform application objects for the
// supported push notification services, such as and The results for
// ListPlatformApplications are paginated and return a limited list of
// applications, up to 100. If additional records are available after the
// first page results, then a NextToken string will be returned. To receive
// the next page, you call ListPlatformApplications using the NextToken
// string received from the previous call. When there are no more records
// to return, NextToken will be null. For more information, see Using
// Amazon SNS Mobile Push Notifications .
func (c *SNS) ListPlatformApplications(req ListPlatformApplicationsInput) (resp *ListPlatformApplicationsResult, err error) {
	resp = &ListPlatformApplicationsResult{}
	err = c.client.Do("ListPlatformApplications", "POST", "/", req, resp)
	return
}

// ListSubscriptions returns a list of the requester's subscriptions. Each
// call returns a limited list of subscriptions, up to 100. If there are
// more subscriptions, a NextToken is also returned. Use the NextToken
// parameter in a new ListSubscriptions call to get further results.
func (c *SNS) ListSubscriptions(req ListSubscriptionsInput) (resp *ListSubscriptionsResult, err error) {
	resp = &ListSubscriptionsResult{}
	err = c.client.Do("ListSubscriptions", "POST", "/", req, resp)
	return
}

// ListSubscriptionsByTopic returns a list of the subscriptions to a
// specific topic. Each call returns a limited list of subscriptions, up to
// 100. If there are more subscriptions, a NextToken is also returned. Use
// the NextToken parameter in a new ListSubscriptionsByTopic call to get
// further results.
func (c *SNS) ListSubscriptionsByTopic(req ListSubscriptionsByTopicInput) (resp *ListSubscriptionsByTopicResult, err error) {
	resp = &ListSubscriptionsByTopicResult{}
	err = c.client.Do("ListSubscriptionsByTopic", "POST", "/", req, resp)
	return
}

// ListTopics returns a list of the requester's topics. Each call returns a
// limited list of topics, up to 100. If there are more topics, a NextToken
// is also returned. Use the NextToken parameter in a new ListTopics call
// to get further results.
func (c *SNS) ListTopics(req ListTopicsInput) (resp *ListTopicsResult, err error) {
	resp = &ListTopicsResult{}
	err = c.client.Do("ListTopics", "POST", "/", req, resp)
	return
}

// Publish sends a message to all of a topic's subscribed endpoints. When a
// messageId is returned, the message has been saved and Amazon SNS will
// attempt to deliver it to the topic's subscribers shortly. The format of
// the outgoing message to each subscribed endpoint depends on the
// notification protocol selected. To use the Publish action for sending a
// message to a mobile endpoint, such as an app on a Kindle device or
// mobile phone, you must specify the EndpointArn. The EndpointArn is
// returned when making a call with the CreatePlatformEndpoint action. The
// second example below shows a request and response for publishing to a
// mobile endpoint.
func (c *SNS) Publish(req PublishInput) (resp *PublishResult, err error) {
	resp = &PublishResult{}
	err = c.client.Do("Publish", "POST", "/", req, resp)
	return
}

// RemovePermission removes a statement from a topic's access control
// policy.
func (c *SNS) RemovePermission(req RemovePermissionInput) (err error) {
	// NRE
	err = c.client.Do("RemovePermission", "POST", "/", req, nil)
	return
}

// SetEndpointAttributes sets the attributes for an endpoint for a device
// on one of the supported push notification services, such as GCM and For
// more information, see Using Amazon SNS Mobile Push Notifications .
func (c *SNS) SetEndpointAttributes(req SetEndpointAttributesInput) (err error) {
	// NRE
	err = c.client.Do("SetEndpointAttributes", "POST", "/", req, nil)
	return
}

// SetPlatformApplicationAttributes sets the attributes of the platform
// application object for the supported push notification services, such as
// and For more information, see Using Amazon SNS Mobile Push Notifications
// .
func (c *SNS) SetPlatformApplicationAttributes(req SetPlatformApplicationAttributesInput) (err error) {
	// NRE
	err = c.client.Do("SetPlatformApplicationAttributes", "POST", "/", req, nil)
	return
}

// SetSubscriptionAttributes allows a subscription owner to set an
// attribute of the topic to a new value.
func (c *SNS) SetSubscriptionAttributes(req SetSubscriptionAttributesInput) (err error) {
	// NRE
	err = c.client.Do("SetSubscriptionAttributes", "POST", "/", req, nil)
	return
}

// SetTopicAttributes allows a topic owner to set an attribute of the topic
// to a new value.
func (c *SNS) SetTopicAttributes(req SetTopicAttributesInput) (err error) {
	// NRE
	err = c.client.Do("SetTopicAttributes", "POST", "/", req, nil)
	return
}

// Subscribe prepares to subscribe an endpoint by sending the endpoint a
// confirmation message. To actually create a subscription, the endpoint
// owner must call the ConfirmSubscription action with the token from the
// confirmation message. Confirmation tokens are valid for three days.
func (c *SNS) Subscribe(req SubscribeInput) (resp *SubscribeResult, err error) {
	resp = &SubscribeResult{}
	err = c.client.Do("Subscribe", "POST", "/", req, resp)
	return
}

// Unsubscribe deletes a subscription. If the subscription requires
// authentication for deletion, only the owner of the subscription or the
// topic's owner can unsubscribe, and an AWS signature is required. If the
// Unsubscribe call does not require authentication and the requester is
// not the subscription owner, a final cancellation message is delivered to
// the endpoint, so that the endpoint owner can easily resubscribe to the
// topic if the Unsubscribe request was unintended.
func (c *SNS) Unsubscribe(req UnsubscribeInput) (err error) {
	// NRE
	err = c.client.Do("Unsubscribe", "POST", "/", req, nil)
	return
}

// AddPermissionInput is undocumented.
type AddPermissionInput struct {
	AWSAccountID []string `xml:"AWSAccountId>member"`
	ActionName   []string `xml:"ActionName>member"`
	Label        string   `xml:"Label"`
	TopicARN     string   `xml:"TopicArn"`
}

// ConfirmSubscriptionInput is undocumented.
type ConfirmSubscriptionInput struct {
	AuthenticateOnUnsubscribe string `xml:"AuthenticateOnUnsubscribe"`
	Token                     string `xml:"Token"`
	TopicARN                  string `xml:"TopicArn"`
}

// ConfirmSubscriptionResponse is undocumented.
type ConfirmSubscriptionResponse struct {
	SubscriptionARN string `xml:"ConfirmSubscriptionResult>SubscriptionArn"`
}

// CreateEndpointResponse is undocumented.
type CreateEndpointResponse struct {
	EndpointARN string `xml:"CreatePlatformEndpointResult>EndpointArn"`
}

// CreatePlatformApplicationInput is undocumented.
type CreatePlatformApplicationInput struct {
	Attributes map[string]string `xml:"Attributes"`
	Name       string            `xml:"Name"`
	Platform   string            `xml:"Platform"`
}

// CreatePlatformApplicationResponse is undocumented.
type CreatePlatformApplicationResponse struct {
	PlatformApplicationARN string `xml:"CreatePlatformApplicationResult>PlatformApplicationArn"`
}

// CreatePlatformEndpointInput is undocumented.
type CreatePlatformEndpointInput struct {
	Attributes             map[string]string `xml:"Attributes"`
	CustomUserData         string            `xml:"CustomUserData"`
	PlatformApplicationARN string            `xml:"PlatformApplicationArn"`
	Token                  string            `xml:"Token"`
}

// CreateTopicInput is undocumented.
type CreateTopicInput struct {
	Name string `xml:"Name"`
}

// CreateTopicResponse is undocumented.
type CreateTopicResponse struct {
	TopicARN string `xml:"CreateTopicResult>TopicArn"`
}

// DeleteEndpointInput is undocumented.
type DeleteEndpointInput struct {
	EndpointARN string `xml:"EndpointArn"`
}

// DeletePlatformApplicationInput is undocumented.
type DeletePlatformApplicationInput struct {
	PlatformApplicationARN string `xml:"PlatformApplicationArn"`
}

// DeleteTopicInput is undocumented.
type DeleteTopicInput struct {
	TopicARN string `xml:"TopicArn"`
}

// Endpoint is undocumented.
type Endpoint struct {
	Attributes  map[string]string `xml:"Attributes"`
	EndpointARN string            `xml:"EndpointArn"`
}

// GetEndpointAttributesInput is undocumented.
type GetEndpointAttributesInput struct {
	EndpointARN string `xml:"EndpointArn"`
}

// GetEndpointAttributesResponse is undocumented.
type GetEndpointAttributesResponse struct {
	Attributes map[string]string `xml:"GetEndpointAttributesResult>Attributes"`
}

// GetPlatformApplicationAttributesInput is undocumented.
type GetPlatformApplicationAttributesInput struct {
	PlatformApplicationARN string `xml:"PlatformApplicationArn"`
}

// GetPlatformApplicationAttributesResponse is undocumented.
type GetPlatformApplicationAttributesResponse struct {
	Attributes map[string]string `xml:"GetPlatformApplicationAttributesResult>Attributes"`
}

// GetSubscriptionAttributesInput is undocumented.
type GetSubscriptionAttributesInput struct {
	SubscriptionARN string `xml:"SubscriptionArn"`
}

// GetSubscriptionAttributesResponse is undocumented.
type GetSubscriptionAttributesResponse struct {
	Attributes map[string]string `xml:"GetSubscriptionAttributesResult>Attributes"`
}

// GetTopicAttributesInput is undocumented.
type GetTopicAttributesInput struct {
	TopicARN string `xml:"TopicArn"`
}

// GetTopicAttributesResponse is undocumented.
type GetTopicAttributesResponse struct {
	Attributes map[string]string `xml:"GetTopicAttributesResult>Attributes"`
}

// ListEndpointsByPlatformApplicationInput is undocumented.
type ListEndpointsByPlatformApplicationInput struct {
	NextToken              string `xml:"NextToken"`
	PlatformApplicationARN string `xml:"PlatformApplicationArn"`
}

// ListEndpointsByPlatformApplicationResponse is undocumented.
type ListEndpointsByPlatformApplicationResponse struct {
	Endpoints []Endpoint `xml:"ListEndpointsByPlatformApplicationResult>Endpoints>member"`
	NextToken string     `xml:"ListEndpointsByPlatformApplicationResult>NextToken"`
}

// ListPlatformApplicationsInput is undocumented.
type ListPlatformApplicationsInput struct {
	NextToken string `xml:"NextToken"`
}

// ListPlatformApplicationsResponse is undocumented.
type ListPlatformApplicationsResponse struct {
	NextToken            string                `xml:"ListPlatformApplicationsResult>NextToken"`
	PlatformApplications []PlatformApplication `xml:"ListPlatformApplicationsResult>PlatformApplications>member"`
}

// ListSubscriptionsByTopicInput is undocumented.
type ListSubscriptionsByTopicInput struct {
	NextToken string `xml:"NextToken"`
	TopicARN  string `xml:"TopicArn"`
}

// ListSubscriptionsByTopicResponse is undocumented.
type ListSubscriptionsByTopicResponse struct {
	NextToken     string         `xml:"ListSubscriptionsByTopicResult>NextToken"`
	Subscriptions []Subscription `xml:"ListSubscriptionsByTopicResult>Subscriptions>member"`
}

// ListSubscriptionsInput is undocumented.
type ListSubscriptionsInput struct {
	NextToken string `xml:"NextToken"`
}

// ListSubscriptionsResponse is undocumented.
type ListSubscriptionsResponse struct {
	NextToken     string         `xml:"ListSubscriptionsResult>NextToken"`
	Subscriptions []Subscription `xml:"ListSubscriptionsResult>Subscriptions>member"`
}

// ListTopicsInput is undocumented.
type ListTopicsInput struct {
	NextToken string `xml:"NextToken"`
}

// ListTopicsResponse is undocumented.
type ListTopicsResponse struct {
	NextToken string  `xml:"ListTopicsResult>NextToken"`
	Topics    []Topic `xml:"ListTopicsResult>Topics>member"`
}

// MessageAttributeValue is undocumented.
type MessageAttributeValue struct {
	BinaryValue []byte `xml:"BinaryValue"`
	DataType    string `xml:"DataType"`
	StringValue string `xml:"StringValue"`
}

// PlatformApplication is undocumented.
type PlatformApplication struct {
	Attributes             map[string]string `xml:"Attributes"`
	PlatformApplicationARN string            `xml:"PlatformApplicationArn"`
}

// PublishInput is undocumented.
type PublishInput struct {
	Message           string                           `xml:"Message"`
	MessageAttributes map[string]MessageAttributeValue `xml:"MessageAttributes"`
	MessageStructure  string                           `xml:"MessageStructure"`
	Subject           string                           `xml:"Subject"`
	TargetARN         string                           `xml:"TargetArn"`
	TopicARN          string                           `xml:"TopicArn"`
}

// PublishResponse is undocumented.
type PublishResponse struct {
	MessageID string `xml:"PublishResult>MessageId"`
}

// RemovePermissionInput is undocumented.
type RemovePermissionInput struct {
	Label    string `xml:"Label"`
	TopicARN string `xml:"TopicArn"`
}

// SetEndpointAttributesInput is undocumented.
type SetEndpointAttributesInput struct {
	Attributes  map[string]string `xml:"Attributes"`
	EndpointARN string            `xml:"EndpointArn"`
}

// SetPlatformApplicationAttributesInput is undocumented.
type SetPlatformApplicationAttributesInput struct {
	Attributes             map[string]string `xml:"Attributes"`
	PlatformApplicationARN string            `xml:"PlatformApplicationArn"`
}

// SetSubscriptionAttributesInput is undocumented.
type SetSubscriptionAttributesInput struct {
	AttributeName   string `xml:"AttributeName"`
	AttributeValue  string `xml:"AttributeValue"`
	SubscriptionARN string `xml:"SubscriptionArn"`
}

// SetTopicAttributesInput is undocumented.
type SetTopicAttributesInput struct {
	AttributeName  string `xml:"AttributeName"`
	AttributeValue string `xml:"AttributeValue"`
	TopicARN       string `xml:"TopicArn"`
}

// SubscribeInput is undocumented.
type SubscribeInput struct {
	Endpoint string `xml:"Endpoint"`
	Protocol string `xml:"Protocol"`
	TopicARN string `xml:"TopicArn"`
}

// SubscribeResponse is undocumented.
type SubscribeResponse struct {
	SubscriptionARN string `xml:"SubscribeResult>SubscriptionArn"`
}

// Subscription is undocumented.
type Subscription struct {
	Endpoint        string `xml:"Endpoint"`
	Owner           string `xml:"Owner"`
	Protocol        string `xml:"Protocol"`
	SubscriptionARN string `xml:"SubscriptionArn"`
	TopicARN        string `xml:"TopicArn"`
}

// Topic is undocumented.
type Topic struct {
	TopicARN string `xml:"TopicArn"`
}

// UnsubscribeInput is undocumented.
type UnsubscribeInput struct {
	SubscriptionARN string `xml:"SubscriptionArn"`
}

// ConfirmSubscriptionResult is a wrapper for ConfirmSubscriptionResponse.
type ConfirmSubscriptionResult struct {
	XMLName xml.Name `xml:"ConfirmSubscriptionResponse"`

	SubscriptionARN string `xml:"ConfirmSubscriptionResult>SubscriptionArn"`
}

// CreatePlatformApplicationResult is a wrapper for CreatePlatformApplicationResponse.
type CreatePlatformApplicationResult struct {
	XMLName xml.Name `xml:"CreatePlatformApplicationResponse"`

	PlatformApplicationARN string `xml:"CreatePlatformApplicationResult>PlatformApplicationArn"`
}

// CreatePlatformEndpointResult is a wrapper for CreateEndpointResponse.
type CreatePlatformEndpointResult struct {
	XMLName xml.Name `xml:"CreatePlatformEndpointResponse"`

	EndpointARN string `xml:"CreatePlatformEndpointResult>EndpointArn"`
}

// CreateTopicResult is a wrapper for CreateTopicResponse.
type CreateTopicResult struct {
	XMLName xml.Name `xml:"CreateTopicResponse"`

	TopicARN string `xml:"CreateTopicResult>TopicArn"`
}

// GetEndpointAttributesResult is a wrapper for GetEndpointAttributesResponse.
type GetEndpointAttributesResult struct {
	XMLName xml.Name `xml:"GetEndpointAttributesResponse"`

	Attributes map[string]string `xml:"GetEndpointAttributesResult>Attributes"`
}

// GetPlatformApplicationAttributesResult is a wrapper for GetPlatformApplicationAttributesResponse.
type GetPlatformApplicationAttributesResult struct {
	XMLName xml.Name `xml:"GetPlatformApplicationAttributesResponse"`

	Attributes map[string]string `xml:"GetPlatformApplicationAttributesResult>Attributes"`
}

// GetSubscriptionAttributesResult is a wrapper for GetSubscriptionAttributesResponse.
type GetSubscriptionAttributesResult struct {
	XMLName xml.Name `xml:"GetSubscriptionAttributesResponse"`

	Attributes map[string]string `xml:"GetSubscriptionAttributesResult>Attributes"`
}

// GetTopicAttributesResult is a wrapper for GetTopicAttributesResponse.
type GetTopicAttributesResult struct {
	XMLName xml.Name `xml:"GetTopicAttributesResponse"`

	Attributes map[string]string `xml:"GetTopicAttributesResult>Attributes"`
}

// ListEndpointsByPlatformApplicationResult is a wrapper for ListEndpointsByPlatformApplicationResponse.
type ListEndpointsByPlatformApplicationResult struct {
	XMLName xml.Name `xml:"ListEndpointsByPlatformApplicationResponse"`

	Endpoints []Endpoint `xml:"ListEndpointsByPlatformApplicationResult>Endpoints>member"`
	NextToken string     `xml:"ListEndpointsByPlatformApplicationResult>NextToken"`
}

// ListPlatformApplicationsResult is a wrapper for ListPlatformApplicationsResponse.
type ListPlatformApplicationsResult struct {
	XMLName xml.Name `xml:"ListPlatformApplicationsResponse"`

	NextToken            string                `xml:"ListPlatformApplicationsResult>NextToken"`
	PlatformApplications []PlatformApplication `xml:"ListPlatformApplicationsResult>PlatformApplications>member"`
}

// ListSubscriptionsByTopicResult is a wrapper for ListSubscriptionsByTopicResponse.
type ListSubscriptionsByTopicResult struct {
	XMLName xml.Name `xml:"ListSubscriptionsByTopicResponse"`

	NextToken     string         `xml:"ListSubscriptionsByTopicResult>NextToken"`
	Subscriptions []Subscription `xml:"ListSubscriptionsByTopicResult>Subscriptions>member"`
}

// ListSubscriptionsResult is a wrapper for ListSubscriptionsResponse.
type ListSubscriptionsResult struct {
	XMLName xml.Name `xml:"ListSubscriptionsResponse"`

	NextToken     string         `xml:"ListSubscriptionsResult>NextToken"`
	Subscriptions []Subscription `xml:"ListSubscriptionsResult>Subscriptions>member"`
}

// ListTopicsResult is a wrapper for ListTopicsResponse.
type ListTopicsResult struct {
	XMLName xml.Name `xml:"ListTopicsResponse"`

	NextToken string  `xml:"ListTopicsResult>NextToken"`
	Topics    []Topic `xml:"ListTopicsResult>Topics>member"`
}

// PublishResult is a wrapper for PublishResponse.
type PublishResult struct {
	XMLName xml.Name `xml:"PublishResponse"`

	MessageID string `xml:"PublishResult>MessageId"`
}

// SubscribeResult is a wrapper for SubscribeResponse.
type SubscribeResult struct {
	XMLName xml.Name `xml:"SubscribeResponse"`

	SubscriptionARN string `xml:"SubscribeResult>SubscriptionArn"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
