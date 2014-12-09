// Package ses provides a client for Amazon Simple Email Service.
package ses

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// SES is a client for Amazon Simple Email Service.
type SES struct {
	client aws.Client
}

// New returns a new SES client.
func New(key, secret, region string, client *http.Client) *SES {
	if client == nil {
		client = http.DefaultClient
	}

	return &SES{
		client: &aws.QueryClient{
			Signer: &aws.V4Signer{
				Key:     key,
				Secret:  secret,
				Service: "email",
				Region:  region,
				IncludeXAmzContentSha256: true,
			},
			Client:     client,
			Endpoint:   endpoints.Lookup("email", region),
			APIVersion: "2010-12-01",
		},
	}
}

// DeleteIdentity deletes the specified identity (email address or domain)
// from the list of verified identities. This action is throttled at one
// request per second.
func (c *SES) DeleteIdentity(req DeleteIdentityRequest) (resp *DeleteIdentityResult, err error) {
	resp = &DeleteIdentityResult{}
	err = c.client.Do("DeleteIdentity", "POST", "/", req, resp)
	return
}

// DeleteVerifiedEmailAddress deletes the specified email address from the
// list of verified addresses. The DeleteVerifiedEmailAddress action is
// deprecated as of the May 15, 2012 release of Domain Verification. The
// DeleteIdentity action is now preferred. This action is throttled at one
// request per second.
func (c *SES) DeleteVerifiedEmailAddress(req DeleteVerifiedEmailAddressRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteVerifiedEmailAddress", "POST", "/", req, nil)
	return
}

// GetIdentityDkimAttributes returns the current status of Easy signing for
// an entity. For domain name identities, this action also returns the
// tokens that are required for Easy signing, and whether Amazon SES has
// successfully verified that these tokens have been published. This action
// takes a list of identities as input and returns the following
// information for each: Whether Easy signing is enabled or disabled. A set
// of tokens that represent the identity. If the identity is an email
// address, the tokens represent the domain of that address. Whether Amazon
// SES has successfully verified the tokens published in the domain's This
// information is only returned for domain name identities, not for email
// addresses. This action is throttled at one request per second. For more
// information about creating DNS records using tokens, go to the Amazon
// SES Developer Guide
func (c *SES) GetIdentityDkimAttributes(req GetIdentityDkimAttributesRequest) (resp *GetIdentityDkimAttributesResult, err error) {
	resp = &GetIdentityDkimAttributesResult{}
	err = c.client.Do("GetIdentityDkimAttributes", "POST", "/", req, resp)
	return
}

// GetIdentityNotificationAttributes given a list of verified identities
// (email addresses and/or domains), returns a structure describing
// identity notification attributes. This action is throttled at one
// request per second. For more information about using notifications with
// Amazon see the Amazon SES Developer Guide
func (c *SES) GetIdentityNotificationAttributes(req GetIdentityNotificationAttributesRequest) (resp *GetIdentityNotificationAttributesResult, err error) {
	resp = &GetIdentityNotificationAttributesResult{}
	err = c.client.Do("GetIdentityNotificationAttributes", "POST", "/", req, resp)
	return
}

// GetIdentityVerificationAttributes given a list of identities (email
// addresses and/or domains), returns the verification status and (for
// domain identities) the verification token for each identity. This action
// is throttled at one request per second.
func (c *SES) GetIdentityVerificationAttributes(req GetIdentityVerificationAttributesRequest) (resp *GetIdentityVerificationAttributesResult, err error) {
	resp = &GetIdentityVerificationAttributesResult{}
	err = c.client.Do("GetIdentityVerificationAttributes", "POST", "/", req, resp)
	return
}

// GetSendQuota returns the user's current sending limits. This action is
// throttled at one request per second.
func (c *SES) GetSendQuota() (resp *GetSendQuotaResult, err error) {
	resp = &GetSendQuotaResult{}
	err = c.client.Do("GetSendQuota", "POST", "/", nil, resp)
	return
}

// GetSendStatistics returns the user's sending statistics. The result is a
// list of data points, representing the last two weeks of sending
// activity. Each data point in the list contains statistics for a
// 15-minute interval. This action is throttled at one request per second.
func (c *SES) GetSendStatistics() (resp *GetSendStatisticsResult, err error) {
	resp = &GetSendStatisticsResult{}
	err = c.client.Do("GetSendStatistics", "POST", "/", nil, resp)
	return
}

// ListIdentities returns a list containing all of the identities (email
// addresses and domains) for a specific AWS Account, regardless of
// verification status. This action is throttled at one request per second.
func (c *SES) ListIdentities(req ListIdentitiesRequest) (resp *ListIdentitiesResult, err error) {
	resp = &ListIdentitiesResult{}
	err = c.client.Do("ListIdentities", "POST", "/", req, resp)
	return
}

// ListVerifiedEmailAddresses returns a list containing all of the email
// addresses that have been verified. The ListVerifiedEmailAddresses action
// is deprecated as of the May 15, 2012 release of Domain Verification. The
// ListIdentities action is now preferred. This action is throttled at one
// request per second.
func (c *SES) ListVerifiedEmailAddresses() (resp *ListVerifiedEmailAddressesResult, err error) {
	resp = &ListVerifiedEmailAddressesResult{}
	err = c.client.Do("ListVerifiedEmailAddresses", "POST", "/", nil, resp)
	return
}

// SendEmail composes an email message based on input data, and then
// immediately queues the message for sending. You can only send email from
// verified email addresses and domains. If you have not requested
// production access to Amazon you must also verify every recipient email
// address except for the recipients provided by the Amazon SES mailbox
// simulator. For more information, go to the Amazon SES Developer Guide .
// The total size of the message cannot exceed 10 Amazon SES has a limit on
// the total number of recipients per message: The combined number of To:,
// CC: and email addresses cannot exceed 50. If you need to send an email
// message to a larger audience, you can divide your recipient list into
// groups of 50 or fewer, and then call Amazon SES repeatedly to send the
// message to each group. For every message that you send, the total number
// of recipients (To:, CC: and is counted against your sending quota - the
// maximum number of emails you can send in a 24-hour period. For
// information about your sending quota, go to the Amazon SES Developer
// Guide .
func (c *SES) SendEmail(req SendEmailRequest) (resp *SendEmailResult, err error) {
	resp = &SendEmailResult{}
	err = c.client.Do("SendEmail", "POST", "/", req, resp)
	return
}

// SendRawEmail sends an email message, with header and content specified
// by the client. The SendRawEmail action is useful for sending multipart
// emails. The raw text of the message must comply with Internet email
// standards; otherwise, the message cannot be sent. You can only send
// email from verified email addresses and domains. If you have not
// requested production access to Amazon you must also verify every
// recipient email address except for the recipients provided by the Amazon
// SES mailbox simulator. For more information, go to the Amazon SES
// Developer Guide . The total size of the message cannot exceed 10 MB.
// This includes any attachments that are part of the message. Amazon SES
// has a limit on the total number of recipients per message: The combined
// number of To:, CC: and email addresses cannot exceed 50. If you need to
// send an email message to a larger audience, you can divide your
// recipient list into groups of 50 or fewer, and then call Amazon SES
// repeatedly to send the message to each group. The To:, and headers in
// the raw message can contain a group list. Note that each recipient in a
// group list counts towards the 50-recipient limit. For every message that
// you send, the total number of recipients (To:, CC: and is counted
// against your sending quota - the maximum number of emails you can send
// in a 24-hour period. For information about your sending quota, go to the
// Amazon SES Developer Guide .
func (c *SES) SendRawEmail(req SendRawEmailRequest) (resp *SendRawEmailResult, err error) {
	resp = &SendRawEmailResult{}
	err = c.client.Do("SendRawEmail", "POST", "/", req, resp)
	return
}

// SetIdentityDkimEnabled enables or disables Easy signing of email sent
// from an identity: If Easy signing is enabled for a domain name identity
// (e.g., example.com ), then Amazon SES will DKIM-sign all email sent by
// addresses under that domain name (e.g., user@example.com If Easy signing
// is enabled for an email address, then Amazon SES will DKIM-sign all
// email sent by that email address. For email addresses (e.g.,
// user@example.com ), you can only enable Easy signing if the
// corresponding domain (e.g., example.com ) has been set up for Easy using
// the AWS Console or the VerifyDomainDkim action. This action is throttled
// at one request per second. For more information about Easy signing, go
// to the Amazon SES Developer Guide
func (c *SES) SetIdentityDkimEnabled(req SetIdentityDkimEnabledRequest) (resp *SetIdentityDkimEnabledResult, err error) {
	resp = &SetIdentityDkimEnabledResult{}
	err = c.client.Do("SetIdentityDkimEnabled", "POST", "/", req, resp)
	return
}

// SetIdentityFeedbackForwardingEnabled given an identity (email address or
// domain), enables or disables whether Amazon SES forwards bounce and
// complaint notifications as email. Feedback forwarding can only be
// disabled when Amazon Simple Notification Service (Amazon topics are
// specified for both bounces and complaints. This action is throttled at
// one request per second. For more information about using notifications
// with Amazon see the Amazon SES Developer Guide
func (c *SES) SetIdentityFeedbackForwardingEnabled(req SetIdentityFeedbackForwardingEnabledRequest) (resp *SetIdentityFeedbackForwardingEnabledResult, err error) {
	resp = &SetIdentityFeedbackForwardingEnabledResult{}
	err = c.client.Do("SetIdentityFeedbackForwardingEnabled", "POST", "/", req, resp)
	return
}

// SetIdentityNotificationTopic given an identity (email address or
// domain), sets the Amazon Simple Notification Service (Amazon topic to
// which Amazon SES will publish bounce, complaint, and/or delivery
// notifications for emails sent with that identity as the Source This
// action is throttled at one request per second. For more information
// about feedback notification, see the Amazon SES Developer Guide
func (c *SES) SetIdentityNotificationTopic(req SetIdentityNotificationTopicRequest) (resp *SetIdentityNotificationTopicResult, err error) {
	resp = &SetIdentityNotificationTopicResult{}
	err = c.client.Do("SetIdentityNotificationTopic", "POST", "/", req, resp)
	return
}

// VerifyDomainDkim returns a set of tokens for a domain. tokens are
// character strings that represent your domain's identity. Using these
// tokens, you will need to create DNS records that point to public keys
// hosted by Amazon Amazon Web Services will eventually detect that you
// have updated your DNS records; this detection process may take up to 72
// hours. Upon successful detection, Amazon SES will be able to DKIM-sign
// email originating from that domain. This action is throttled at one
// request per second. To enable or disable Easy signing for a domain, use
// the SetIdentityDkimEnabled action. For more information about creating
// DNS records using tokens, go to the Amazon SES Developer Guide
func (c *SES) VerifyDomainDkim(req VerifyDomainDkimRequest) (resp *VerifyDomainDkimResult, err error) {
	resp = &VerifyDomainDkimResult{}
	err = c.client.Do("VerifyDomainDkim", "POST", "/", req, resp)
	return
}

// VerifyDomainIdentity verifies a domain. This action is throttled at one
// request per second.
func (c *SES) VerifyDomainIdentity(req VerifyDomainIdentityRequest) (resp *VerifyDomainIdentityResult, err error) {
	resp = &VerifyDomainIdentityResult{}
	err = c.client.Do("VerifyDomainIdentity", "POST", "/", req, resp)
	return
}

// VerifyEmailAddress verifies an email address. This action causes a
// confirmation email message to be sent to the specified address. The
// VerifyEmailAddress action is deprecated as of the May 15, 2012 release
// of Domain Verification. The VerifyEmailIdentity action is now preferred.
// This action is throttled at one request per second.
func (c *SES) VerifyEmailAddress(req VerifyEmailAddressRequest) (err error) {
	// NRE
	err = c.client.Do("VerifyEmailAddress", "POST", "/", req, nil)
	return
}

// VerifyEmailIdentity verifies an email address. This action causes a
// confirmation email message to be sent to the specified address. This
// action is throttled at one request per second.
func (c *SES) VerifyEmailIdentity(req VerifyEmailIdentityRequest) (resp *VerifyEmailIdentityResult, err error) {
	resp = &VerifyEmailIdentityResult{}
	err = c.client.Do("VerifyEmailIdentity", "POST", "/", req, resp)
	return
}

// Body is undocumented.
type Body struct {
	Html Content `xml:"Html"`
	Text Content `xml:"Text"`
}

// Content is undocumented.
type Content struct {
	Charset string `xml:"Charset"`
	Data    string `xml:"Data"`
}

// DeleteIdentityRequest is undocumented.
type DeleteIdentityRequest struct {
	Identity string `xml:"Identity"`
}

// DeleteIdentityResponse is undocumented.
type DeleteIdentityResponse struct {
}

// DeleteVerifiedEmailAddressRequest is undocumented.
type DeleteVerifiedEmailAddressRequest struct {
	EmailAddress string `xml:"EmailAddress"`
}

// Destination is undocumented.
type Destination struct {
	BccAddresses []string `xml:"BccAddresses>member"`
	CcAddresses  []string `xml:"CcAddresses>member"`
	ToAddresses  []string `xml:"ToAddresses>member"`
}

// GetIdentityDkimAttributesRequest is undocumented.
type GetIdentityDkimAttributesRequest struct {
	Identities []string `xml:"Identities>member"`
}

// GetIdentityDkimAttributesResponse is undocumented.
type GetIdentityDkimAttributesResponse struct {
	DkimAttributes map[string]IdentityDkimAttributes `xml:"GetIdentityDkimAttributesResult>DkimAttributes"`
}

// GetIdentityNotificationAttributesRequest is undocumented.
type GetIdentityNotificationAttributesRequest struct {
	Identities []string `xml:"Identities>member"`
}

// GetIdentityNotificationAttributesResponse is undocumented.
type GetIdentityNotificationAttributesResponse struct {
	NotificationAttributes map[string]IdentityNotificationAttributes `xml:"GetIdentityNotificationAttributesResult>NotificationAttributes"`
}

// GetIdentityVerificationAttributesRequest is undocumented.
type GetIdentityVerificationAttributesRequest struct {
	Identities []string `xml:"Identities>member"`
}

// GetIdentityVerificationAttributesResponse is undocumented.
type GetIdentityVerificationAttributesResponse struct {
	VerificationAttributes map[string]IdentityVerificationAttributes `xml:"GetIdentityVerificationAttributesResult>VerificationAttributes"`
}

// GetSendQuotaResponse is undocumented.
type GetSendQuotaResponse struct {
	Max24HourSend   float64 `xml:"GetSendQuotaResult>Max24HourSend"`
	MaxSendRate     float64 `xml:"GetSendQuotaResult>MaxSendRate"`
	SentLast24Hours float64 `xml:"GetSendQuotaResult>SentLast24Hours"`
}

// GetSendStatisticsResponse is undocumented.
type GetSendStatisticsResponse struct {
	SendDataPoints []SendDataPoint `xml:"GetSendStatisticsResult>SendDataPoints>member"`
}

// IdentityDkimAttributes is undocumented.
type IdentityDkimAttributes struct {
	DkimEnabled            bool     `xml:"DkimEnabled"`
	DkimTokens             []string `xml:"DkimTokens>member"`
	DkimVerificationStatus string   `xml:"DkimVerificationStatus"`
}

// IdentityNotificationAttributes is undocumented.
type IdentityNotificationAttributes struct {
	BounceTopic       string `xml:"BounceTopic"`
	ComplaintTopic    string `xml:"ComplaintTopic"`
	DeliveryTopic     string `xml:"DeliveryTopic"`
	ForwardingEnabled bool   `xml:"ForwardingEnabled"`
}

// IdentityVerificationAttributes is undocumented.
type IdentityVerificationAttributes struct {
	VerificationStatus string `xml:"VerificationStatus"`
	VerificationToken  string `xml:"VerificationToken"`
}

// ListIdentitiesRequest is undocumented.
type ListIdentitiesRequest struct {
	IdentityType string `xml:"IdentityType"`
	MaxItems     int    `xml:"MaxItems"`
	NextToken    string `xml:"NextToken"`
}

// ListIdentitiesResponse is undocumented.
type ListIdentitiesResponse struct {
	Identities []string `xml:"ListIdentitiesResult>Identities>member"`
	NextToken  string   `xml:"ListIdentitiesResult>NextToken"`
}

// ListVerifiedEmailAddressesResponse is undocumented.
type ListVerifiedEmailAddressesResponse struct {
	VerifiedEmailAddresses []string `xml:"ListVerifiedEmailAddressesResult>VerifiedEmailAddresses>member"`
}

// Message is undocumented.
type Message struct {
	Body    Body    `xml:"Body"`
	Subject Content `xml:"Subject"`
}

// RawMessage is undocumented.
type RawMessage struct {
	Data []byte `xml:"Data"`
}

// SendDataPoint is undocumented.
type SendDataPoint struct {
	Bounces          int64     `xml:"Bounces"`
	Complaints       int64     `xml:"Complaints"`
	DeliveryAttempts int64     `xml:"DeliveryAttempts"`
	Rejects          int64     `xml:"Rejects"`
	Timestamp        time.Time `xml:"Timestamp"`
}

// SendEmailRequest is undocumented.
type SendEmailRequest struct {
	Destination      Destination `xml:"Destination"`
	Message          Message     `xml:"Message"`
	ReplyToAddresses []string    `xml:"ReplyToAddresses>member"`
	ReturnPath       string      `xml:"ReturnPath"`
	Source           string      `xml:"Source"`
}

// SendEmailResponse is undocumented.
type SendEmailResponse struct {
	MessageID string `xml:"SendEmailResult>MessageId"`
}

// SendRawEmailRequest is undocumented.
type SendRawEmailRequest struct {
	Destinations []string   `xml:"Destinations>member"`
	RawMessage   RawMessage `xml:"RawMessage"`
	Source       string     `xml:"Source"`
}

// SendRawEmailResponse is undocumented.
type SendRawEmailResponse struct {
	MessageID string `xml:"SendRawEmailResult>MessageId"`
}

// SetIdentityDkimEnabledRequest is undocumented.
type SetIdentityDkimEnabledRequest struct {
	DkimEnabled bool   `xml:"DkimEnabled"`
	Identity    string `xml:"Identity"`
}

// SetIdentityDkimEnabledResponse is undocumented.
type SetIdentityDkimEnabledResponse struct {
}

// SetIdentityFeedbackForwardingEnabledRequest is undocumented.
type SetIdentityFeedbackForwardingEnabledRequest struct {
	ForwardingEnabled bool   `xml:"ForwardingEnabled"`
	Identity          string `xml:"Identity"`
}

// SetIdentityFeedbackForwardingEnabledResponse is undocumented.
type SetIdentityFeedbackForwardingEnabledResponse struct {
}

// SetIdentityNotificationTopicRequest is undocumented.
type SetIdentityNotificationTopicRequest struct {
	Identity         string `xml:"Identity"`
	NotificationType string `xml:"NotificationType"`
	SnsTopic         string `xml:"SnsTopic"`
}

// SetIdentityNotificationTopicResponse is undocumented.
type SetIdentityNotificationTopicResponse struct {
}

// VerifyDomainDkimRequest is undocumented.
type VerifyDomainDkimRequest struct {
	Domain string `xml:"Domain"`
}

// VerifyDomainDkimResponse is undocumented.
type VerifyDomainDkimResponse struct {
	DkimTokens []string `xml:"VerifyDomainDkimResult>DkimTokens>member"`
}

// VerifyDomainIdentityRequest is undocumented.
type VerifyDomainIdentityRequest struct {
	Domain string `xml:"Domain"`
}

// VerifyDomainIdentityResponse is undocumented.
type VerifyDomainIdentityResponse struct {
	VerificationToken string `xml:"VerifyDomainIdentityResult>VerificationToken"`
}

// VerifyEmailAddressRequest is undocumented.
type VerifyEmailAddressRequest struct {
	EmailAddress string `xml:"EmailAddress"`
}

// VerifyEmailIdentityRequest is undocumented.
type VerifyEmailIdentityRequest struct {
	EmailAddress string `xml:"EmailAddress"`
}

// VerifyEmailIdentityResponse is undocumented.
type VerifyEmailIdentityResponse struct {
}

// DeleteIdentityResult is a wrapper for DeleteIdentityResponse.
type DeleteIdentityResult struct {
	XMLName xml.Name `xml:"DeleteIdentityResponse"`
}

// GetIdentityDkimAttributesResult is a wrapper for GetIdentityDkimAttributesResponse.
type GetIdentityDkimAttributesResult struct {
	XMLName xml.Name `xml:"GetIdentityDkimAttributesResponse"`

	DkimAttributes map[string]IdentityDkimAttributes `xml:"GetIdentityDkimAttributesResult>DkimAttributes"`
}

// GetIdentityNotificationAttributesResult is a wrapper for GetIdentityNotificationAttributesResponse.
type GetIdentityNotificationAttributesResult struct {
	XMLName xml.Name `xml:"GetIdentityNotificationAttributesResponse"`

	NotificationAttributes map[string]IdentityNotificationAttributes `xml:"GetIdentityNotificationAttributesResult>NotificationAttributes"`
}

// GetIdentityVerificationAttributesResult is a wrapper for GetIdentityVerificationAttributesResponse.
type GetIdentityVerificationAttributesResult struct {
	XMLName xml.Name `xml:"GetIdentityVerificationAttributesResponse"`

	VerificationAttributes map[string]IdentityVerificationAttributes `xml:"GetIdentityVerificationAttributesResult>VerificationAttributes"`
}

// GetSendQuotaResult is a wrapper for GetSendQuotaResponse.
type GetSendQuotaResult struct {
	XMLName xml.Name `xml:"GetSendQuotaResponse"`

	Max24HourSend   float64 `xml:"GetSendQuotaResult>Max24HourSend"`
	MaxSendRate     float64 `xml:"GetSendQuotaResult>MaxSendRate"`
	SentLast24Hours float64 `xml:"GetSendQuotaResult>SentLast24Hours"`
}

// GetSendStatisticsResult is a wrapper for GetSendStatisticsResponse.
type GetSendStatisticsResult struct {
	XMLName xml.Name `xml:"GetSendStatisticsResponse"`

	SendDataPoints []SendDataPoint `xml:"GetSendStatisticsResult>SendDataPoints>member"`
}

// ListIdentitiesResult is a wrapper for ListIdentitiesResponse.
type ListIdentitiesResult struct {
	XMLName xml.Name `xml:"ListIdentitiesResponse"`

	Identities []string `xml:"ListIdentitiesResult>Identities>member"`
	NextToken  string   `xml:"ListIdentitiesResult>NextToken"`
}

// ListVerifiedEmailAddressesResult is a wrapper for ListVerifiedEmailAddressesResponse.
type ListVerifiedEmailAddressesResult struct {
	XMLName xml.Name `xml:"ListVerifiedEmailAddressesResponse"`

	VerifiedEmailAddresses []string `xml:"ListVerifiedEmailAddressesResult>VerifiedEmailAddresses>member"`
}

// SendEmailResult is a wrapper for SendEmailResponse.
type SendEmailResult struct {
	XMLName xml.Name `xml:"SendEmailResponse"`

	MessageID string `xml:"SendEmailResult>MessageId"`
}

// SendRawEmailResult is a wrapper for SendRawEmailResponse.
type SendRawEmailResult struct {
	XMLName xml.Name `xml:"SendRawEmailResponse"`

	MessageID string `xml:"SendRawEmailResult>MessageId"`
}

// SetIdentityDkimEnabledResult is a wrapper for SetIdentityDkimEnabledResponse.
type SetIdentityDkimEnabledResult struct {
	XMLName xml.Name `xml:"SetIdentityDkimEnabledResponse"`
}

// SetIdentityFeedbackForwardingEnabledResult is a wrapper for SetIdentityFeedbackForwardingEnabledResponse.
type SetIdentityFeedbackForwardingEnabledResult struct {
	XMLName xml.Name `xml:"SetIdentityFeedbackForwardingEnabledResponse"`
}

// SetIdentityNotificationTopicResult is a wrapper for SetIdentityNotificationTopicResponse.
type SetIdentityNotificationTopicResult struct {
	XMLName xml.Name `xml:"SetIdentityNotificationTopicResponse"`
}

// VerifyDomainDkimResult is a wrapper for VerifyDomainDkimResponse.
type VerifyDomainDkimResult struct {
	XMLName xml.Name `xml:"VerifyDomainDkimResponse"`

	DkimTokens []string `xml:"VerifyDomainDkimResult>DkimTokens>member"`
}

// VerifyDomainIdentityResult is a wrapper for VerifyDomainIdentityResponse.
type VerifyDomainIdentityResult struct {
	XMLName xml.Name `xml:"VerifyDomainIdentityResponse"`

	VerificationToken string `xml:"VerifyDomainIdentityResult>VerificationToken"`
}

// VerifyEmailIdentityResult is a wrapper for VerifyEmailIdentityResponse.
type VerifyEmailIdentityResult struct {
	XMLName xml.Name `xml:"VerifyEmailIdentityResponse"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
