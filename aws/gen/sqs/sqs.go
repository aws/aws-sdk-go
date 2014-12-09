// Package sqs provides a client for Amazon Simple Queue Service.
package sqs

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// SQS is a client for Amazon Simple Queue Service.
type SQS struct {
	client *aws.QueryClient
}

// New returns a new SQS client.
func New(key, secret, region string, client *http.Client) *SQS {
	if client == nil {
		client = http.DefaultClient
	}

	return &SQS{
		client: &aws.QueryClient{
			Signer: &aws.V4Signer{
				Key:     key,
				Secret:  secret,
				Service: "sqs",
				Region:  region,
				IncludeXAmzContentSha256: true,
			},
			Client:     client,
			Endpoint:   endpoints.Lookup("sqs", region),
			APIVersion: "2012-11-05",
		},
	}
}

// AddPermission adds a permission to a queue for a specific principal .
// This allows for sharing access to the queue. When you create a queue,
// you have full control access rights for the queue. Only you (as owner of
// the queue) can grant or deny permissions to the queue. For more
// information about these permissions, see Shared Queues in the Amazon SQS
// Developer Guide
func (c *SQS) AddPermission(req AddPermissionRequest) (err error) {
	// NRE
	err = c.client.Do("AddPermission", "POST", "/", req, nil)
	return
}

// ChangeMessageVisibility changes the visibility timeout of a specified
// message in a queue to a new value. The maximum allowed timeout value you
// can set the value to is 12 hours. This means you can't extend the
// timeout of a message in an existing queue to more than a total
// visibility timeout of 12 hours. (For more information visibility
// timeout, see Visibility Timeout in the Amazon SQS Developer Guide For
// example, let's say you have a message and its default message visibility
// timeout is 30 minutes. You could call ChangeMessageVisiblity with a
// value of two hours and the effective timeout would be two hours and 30
// minutes. When that time comes near you could again extend the time out
// by calling ChangeMessageVisiblity, but this time the maximum allowed
// timeout would be 9 hours and 30 minutes. If you attempt to set the
// VisibilityTimeout to an amount more than the maximum time left, Amazon
// SQS returns an error. It will not automatically recalculate and increase
// the timeout to the maximum time remaining. Unlike with a queue, when you
// change the visibility timeout for a specific message, that timeout value
// is applied immediately but is not saved in memory for that message. If
// you don't delete a message after it is received, the visibility timeout
// for the message the next time it is received reverts to the original
// timeout value, not the value you set with the ChangeMessageVisibility
// action.
func (c *SQS) ChangeMessageVisibility(req ChangeMessageVisibilityRequest) (err error) {
	// NRE
	err = c.client.Do("ChangeMessageVisibility", "POST", "/", req, nil)
	return
}

// ChangeMessageVisibilityBatch changes the visibility timeout of multiple
// messages. This is a batch version of ChangeMessageVisibility . The
// result of the action on each message is reported individually in the
// response. You can send up to 10 ChangeMessageVisibility requests with
// each ChangeMessageVisibilityBatch action. Because the batch request can
// result in a combination of successful and unsuccessful actions, you
// should check for batch errors even when the call returns an status code
// of 200.
func (c *SQS) ChangeMessageVisibilityBatch(req ChangeMessageVisibilityBatchRequest) (resp *ChangeMessageVisibilityBatchResult, err error) {
	resp = &ChangeMessageVisibilityBatchResult{}
	err = c.client.Do("ChangeMessageVisibilityBatch", "POST", "/", req, resp)
	return
}

// CreateQueue creates a new queue, or returns the URL of an existing one.
// When you request CreateQueue , you provide a name for the queue. To
// successfully create a new queue, you must provide a name that is unique
// within the scope of your own queues. You may pass one or more attributes
// in the request. If you do not provide a value for any attribute, the
// queue will have the default value for that attribute. Permitted
// attributes are the same that can be set using SetQueueAttributes If you
// provide the name of an existing queue, along with the exact names and
// values of all the queue's attributes, CreateQueue returns the queue URL
// for the existing queue. If the queue name, attribute names, or attribute
// values do not match an existing queue, CreateQueue returns an error.
func (c *SQS) CreateQueue(req CreateQueueRequest) (resp *CreateQueueResult, err error) {
	resp = &CreateQueueResult{}
	err = c.client.Do("CreateQueue", "POST", "/", req, resp)
	return
}

// DeleteMessage deletes the specified message from the specified queue.
// You specify the message by using the message's receipt handle and not
// the message you received when you sent the message. Even if the message
// is locked by another reader due to the visibility timeout setting, it is
// still deleted from the queue. If you leave a message in the queue for
// longer than the queue's configured retention period, Amazon SQS
// automatically deletes it. It is possible you will receive a message even
// after you have deleted it. This might happen on rare occasions if one of
// the servers storing a copy of the message is unavailable when you
// request to delete the message. The copy remains on the server and might
// be returned to you again on a subsequent receive request. You should
// create your system to be idempotent so that receiving a particular
// message more than once is not a problem.
func (c *SQS) DeleteMessage(req DeleteMessageRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteMessage", "POST", "/", req, nil)
	return
}

// DeleteMessageBatch deletes multiple messages. This is a batch version of
// DeleteMessage . The result of the delete action on each message is
// reported individually in the response. Because the batch request can
// result in a combination of successful and unsuccessful actions, you
// should check for batch errors even when the call returns an status code
// of 200.
func (c *SQS) DeleteMessageBatch(req DeleteMessageBatchRequest) (resp *DeleteMessageBatchResult, err error) {
	resp = &DeleteMessageBatchResult{}
	err = c.client.Do("DeleteMessageBatch", "POST", "/", req, resp)
	return
}

// DeleteQueue deletes the queue specified by the queue , regardless of
// whether the queue is empty. If the specified queue does not exist,
// Amazon SQS returns a successful response. Use DeleteQueue with care;
// once you delete your queue, any messages in the queue are no longer
// available. When you delete a queue, the deletion process takes up to 60
// seconds. Requests you send involving that queue during the 60 seconds
// might succeed. For example, a SendMessage request might succeed, but
// after the 60 seconds, the queue and that message you sent no longer
// exist. Also, when you delete a queue, you must wait at least 60 seconds
// before creating a queue with the same name. We reserve the right to
// delete queues that have had no activity for more than 30 days. For more
// information, see How Amazon SQS Queues Work in the Amazon SQS Developer
// Guide .
func (c *SQS) DeleteQueue(req DeleteQueueRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteQueue", "POST", "/", req, nil)
	return
}

// GetQueueAttributes gets attributes for the specified queue. The
// following attributes are supported: All - returns all values.
// ApproximateNumberOfMessages - returns the approximate number of visible
// messages in a queue. For more information, see Resources Required to
// Process Messages in the Amazon SQS Developer Guide
// ApproximateNumberOfMessagesNotVisible - returns the approximate number
// of messages that are not timed-out and not deleted. For more
// information, see Resources Required to Process Messages in the Amazon
// SQS Developer Guide VisibilityTimeout - returns the visibility timeout
// for the queue. For more information about visibility timeout, see
// Visibility Timeout in the Amazon SQS Developer Guide CreatedTimestamp -
// returns the time when the queue was created (epoch time in seconds).
// LastModifiedTimestamp - returns the time when the queue was last changed
// (epoch time in seconds). Policy - returns the queue's policy.
// MaximumMessageSize - returns the limit of how many bytes a message can
// contain before Amazon SQS rejects it. MessageRetentionPeriod - returns
// the number of seconds Amazon SQS retains a message. QueueArn - returns
// the queue's Amazon resource name ApproximateNumberOfMessagesDelayed -
// returns the approximate number of messages that are pending to be added
// to the queue. DelaySeconds - returns the default delay on the queue in
// seconds. ReceiveMessageWaitTimeSeconds - returns the time for which a
// ReceiveMessage call will wait for a message to arrive. RedrivePolicy -
// returns the parameters for dead letter queue functionality of the source
// queue. For more information about RedrivePolicy and dead letter queues,
// see Using Amazon SQS Dead Letter Queues in the Amazon SQS Developer
// Guide
func (c *SQS) GetQueueAttributes(req GetQueueAttributesRequest) (resp *GetQueueAttributesResult, err error) {
	resp = &GetQueueAttributesResult{}
	err = c.client.Do("GetQueueAttributes", "POST", "/", req, resp)
	return
}

// GetQueueURL returns the URL of an existing queue. This action provides a
// simple way to retrieve the URL of an Amazon SQS queue. To access a queue
// that belongs to another AWS account, use the QueueOwnerAWSAccountId
// parameter to specify the account ID of the queue's owner. The queue's
// owner must grant you permission to access the queue. For more
// information about shared queue access, see AddPermission or go to Shared
// Queues in the Amazon SQS Developer Guide .
func (c *SQS) GetQueueURL(req GetQueueURLRequest) (resp *GetQueueURLResult, err error) {
	resp = &GetQueueURLResult{}
	err = c.client.Do("GetQueueUrl", "POST", "/", req, resp)
	return
}

// ListDeadLetterSourceQueues returns a list of your queues that have the
// RedrivePolicy queue attribute configured with a dead letter queue. For
// more information about using dead letter queues, see Using Amazon SQS
// Dead Letter Queues
func (c *SQS) ListDeadLetterSourceQueues(req ListDeadLetterSourceQueuesRequest) (resp *ListDeadLetterSourceQueuesResult, err error) {
	resp = &ListDeadLetterSourceQueuesResult{}
	err = c.client.Do("ListDeadLetterSourceQueues", "POST", "/", req, resp)
	return
}

// ListQueues returns a list of your queues. The maximum number of queues
// that can be returned is 1000. If you specify a value for the optional
// QueueNamePrefix parameter, only queues with a name beginning with the
// specified value are returned.
func (c *SQS) ListQueues(req ListQueuesRequest) (resp *ListQueuesResult, err error) {
	resp = &ListQueuesResult{}
	err = c.client.Do("ListQueues", "POST", "/", req, resp)
	return
}

// ReceiveMessage retrieves one or more messages, with a maximum limit of
// 10 messages, from the specified queue. Long poll support is enabled by
// using the WaitTimeSeconds parameter. For more information, see Amazon
// SQS Long Poll in the Amazon SQS Developer Guide . Short poll is the
// default behavior where a weighted random set of machines is sampled on a
// ReceiveMessage call. This means only the messages on the sampled
// machines are returned. If the number of messages in the queue is small
// (less than 1000), it is likely you will get fewer messages than you
// requested per ReceiveMessage call. If the number of messages in the
// queue is extremely small, you might not receive any messages in a
// particular ReceiveMessage response; in which case you should repeat the
// request. For each message returned, the response includes the following:
// MD5 digest of the message body. For information about MD5, go to
// http://www.faqs.org/rfcs/rfc1321.html . Message ID you received when you
// sent the message to the queue. Receipt handle. MD5 digest of the message
// attributes. The receipt handle is the identifier you must provide when
// deleting the message. For more information, see Queue and Message
// Identifiers in the Amazon SQS Developer Guide . You can provide the
// VisibilityTimeout parameter in your request, which will be applied to
// the messages that Amazon SQS returns in the response. If you do not
// include the parameter, the overall visibility timeout for the queue is
// used for the returned messages. For more information, see Visibility
// Timeout in the Amazon SQS Developer Guide .
func (c *SQS) ReceiveMessage(req ReceiveMessageRequest) (resp *ReceiveMessageResult, err error) {
	resp = &ReceiveMessageResult{}
	err = c.client.Do("ReceiveMessage", "POST", "/", req, resp)
	return
}

// RemovePermission revokes any permissions in the queue policy that
// matches the specified Label parameter. Only the owner of the queue can
// remove permissions.
func (c *SQS) RemovePermission(req RemovePermissionRequest) (err error) {
	// NRE
	err = c.client.Do("RemovePermission", "POST", "/", req, nil)
	return
}

// SendMessage delivers a message to the specified queue. With Amazon you
// now have the ability to send large payload messages that are up to 256KB
// (262,144 bytes) in size. To send large payloads, you must use an AWS SDK
// that supports SigV4 signing. To verify whether SigV4 is supported for an
// AWS check the SDK release notes. The following list shows the characters
// (in Unicode) allowed in your message, according to the W3C XML
// specification. For more information, go to
// http://www.w3.org/TR/REC-xml/#charsets If you send any characters not
// included in the list, your request will be rejected. #x9 | #xA | #xD |
// [#x20 to #xD7FF] | [#xE000 to #xFFFD] | [#x10000 to #x10FFFF]
func (c *SQS) SendMessage(req SendMessageRequest) (resp *SendMessageResult, err error) {
	resp = &SendMessageResult{}
	err = c.client.Do("SendMessage", "POST", "/", req, resp)
	return
}

// SendMessageBatch delivers up to ten messages to the specified queue.
// This is a batch version of SendMessage . The result of the send action
// on each message is reported individually in the response. The maximum
// allowed individual message size is 256 KB (262,144 bytes). The maximum
// total payload size (i.e., the sum of all a batch's individual message
// lengths) is also 256 KB (262,144 bytes). If the DelaySeconds parameter
// is not specified for an entry, the default for the queue is used. The
// following list shows the characters (in Unicode) that are allowed in
// your message, according to the W3C XML specification. For more
// information, go to http://www.faqs.org/rfcs/rfc1321.html . If you send
// any characters that are not included in the list, your request will be
// rejected. #x9 | #xA | #xD | [#x20 to #xD7FF] | [#xE000 to #xFFFD] |
// [#x10000 to #x10FFFF] Because the batch request can result in a
// combination of successful and unsuccessful actions, you should check for
// batch errors even when the call returns an status code of 200.
func (c *SQS) SendMessageBatch(req SendMessageBatchRequest) (resp *SendMessageBatchResult, err error) {
	resp = &SendMessageBatchResult{}
	err = c.client.Do("SendMessageBatch", "POST", "/", req, resp)
	return
}

// SetQueueAttributes sets the value of one or more queue attributes. When
// you change a queue's attributes, the change can take up to 60 seconds
// for most of the attributes to propagate throughout the SQS system.
// Changes made to the MessageRetentionPeriod attribute can take up to 15
// minutes.
func (c *SQS) SetQueueAttributes(req SetQueueAttributesRequest) (err error) {
	// NRE
	err = c.client.Do("SetQueueAttributes", "POST", "/", req, nil)
	return
}

// AddPermissionRequest is undocumented.
type AddPermissionRequest struct {
	AWSAccountIds []string `xml:"AWSAccountIds>AWSAccountId"`
	Actions       []string `xml:"Actions>ActionName"`
	Label         string   `xml:"Label"`
	QueueURL      string   `xml:"QueueUrl"`
}

// BatchResultErrorEntry is undocumented.
type BatchResultErrorEntry struct {
	Code        string `xml:"Code"`
	ID          string `xml:"Id"`
	Message     string `xml:"Message"`
	SenderFault bool   `xml:"SenderFault"`
}

// ChangeMessageVisibilityBatchRequest is undocumented.
type ChangeMessageVisibilityBatchRequest struct {
	Entries  []ChangeMessageVisibilityBatchRequestEntry `xml:"Entries>ChangeMessageVisibilityBatchRequestEntry"`
	QueueURL string                                     `xml:"QueueUrl"`
}

// ChangeMessageVisibilityBatchRequestEntry is undocumented.
type ChangeMessageVisibilityBatchRequestEntry struct {
	ID                string `xml:"Id"`
	ReceiptHandle     string `xml:"ReceiptHandle"`
	VisibilityTimeout int    `xml:"VisibilityTimeout"`
}

// ChangeMessageVisibilityBatchResult is undocumented.
type ChangeMessageVisibilityBatchResult struct {
	Failed     []BatchResultErrorEntry                   `xml:"ChangeMessageVisibilityBatchResult>Failed>BatchResultErrorEntry"`
	Successful []ChangeMessageVisibilityBatchResultEntry `xml:"ChangeMessageVisibilityBatchResult>Successful>ChangeMessageVisibilityBatchResultEntry"`
}

// ChangeMessageVisibilityBatchResultEntry is undocumented.
type ChangeMessageVisibilityBatchResultEntry struct {
	ID string `xml:"Id"`
}

// ChangeMessageVisibilityRequest is undocumented.
type ChangeMessageVisibilityRequest struct {
	QueueURL          string `xml:"QueueUrl"`
	ReceiptHandle     string `xml:"ReceiptHandle"`
	VisibilityTimeout int    `xml:"VisibilityTimeout"`
}

// CreateQueueRequest is undocumented.
type CreateQueueRequest struct {
	Attributes map[string]string `xml:"Attribute"`
	QueueName  string            `xml:"QueueName"`
}

// CreateQueueResult is undocumented.
type CreateQueueResult struct {
	QueueURL string `xml:"CreateQueueResult>QueueUrl"`
}

// DeleteMessageBatchRequest is undocumented.
type DeleteMessageBatchRequest struct {
	Entries  []DeleteMessageBatchRequestEntry `xml:"Entries>DeleteMessageBatchRequestEntry"`
	QueueURL string                           `xml:"QueueUrl"`
}

// DeleteMessageBatchRequestEntry is undocumented.
type DeleteMessageBatchRequestEntry struct {
	ID            string `xml:"Id"`
	ReceiptHandle string `xml:"ReceiptHandle"`
}

// DeleteMessageBatchResult is undocumented.
type DeleteMessageBatchResult struct {
	Failed     []BatchResultErrorEntry         `xml:"DeleteMessageBatchResult>Failed>BatchResultErrorEntry"`
	Successful []DeleteMessageBatchResultEntry `xml:"DeleteMessageBatchResult>Successful>DeleteMessageBatchResultEntry"`
}

// DeleteMessageBatchResultEntry is undocumented.
type DeleteMessageBatchResultEntry struct {
	ID string `xml:"Id"`
}

// DeleteMessageRequest is undocumented.
type DeleteMessageRequest struct {
	QueueURL      string `xml:"QueueUrl"`
	ReceiptHandle string `xml:"ReceiptHandle"`
}

// DeleteQueueRequest is undocumented.
type DeleteQueueRequest struct {
	QueueURL string `xml:"QueueUrl"`
}

// GetQueueAttributesRequest is undocumented.
type GetQueueAttributesRequest struct {
	AttributeNames []string `xml:"AttributeNames>AttributeName"`
	QueueURL       string   `xml:"QueueUrl"`
}

// GetQueueAttributesResult is undocumented.
type GetQueueAttributesResult struct {
	Attributes map[string]string `xml:"GetQueueAttributesResult>Attribute"`
}

// GetQueueURLRequest is undocumented.
type GetQueueURLRequest struct {
	QueueName              string `xml:"QueueName"`
	QueueOwnerAWSAccountID string `xml:"QueueOwnerAWSAccountId"`
}

// GetQueueURLResult is undocumented.
type GetQueueURLResult struct {
	QueueURL string `xml:"GetQueueUrlResult>QueueUrl"`
}

// ListDeadLetterSourceQueuesRequest is undocumented.
type ListDeadLetterSourceQueuesRequest struct {
	QueueURL string `xml:"QueueUrl"`
}

// ListDeadLetterSourceQueuesResult is undocumented.
type ListDeadLetterSourceQueuesResult struct {
	QueueURLs []string `xml:"ListDeadLetterSourceQueuesResult>queueUrls>QueueUrl"`
}

// ListQueuesRequest is undocumented.
type ListQueuesRequest struct {
	QueueNamePrefix string `xml:"QueueNamePrefix"`
}

// ListQueuesResult is undocumented.
type ListQueuesResult struct {
	QueueURLs []string `xml:"ListQueuesResult>QueueUrls>QueueUrl"`
}

// Message is undocumented.
type Message struct {
	Attributes             map[string]string                `xml:"Attribute"`
	Body                   string                           `xml:"Body"`
	MD5OfBody              string                           `xml:"MD5OfBody"`
	MD5OfMessageAttributes string                           `xml:"MD5OfMessageAttributes"`
	MessageAttributes      map[string]MessageAttributeValue `xml:"MessageAttribute"`
	MessageID              string                           `xml:"MessageId"`
	ReceiptHandle          string                           `xml:"ReceiptHandle"`
}

// MessageAttributeValue is undocumented.
type MessageAttributeValue struct {
	BinaryListValues [][]byte `xml:"BinaryListValue>BinaryListValue"`
	BinaryValue      []byte   `xml:"BinaryValue"`
	DataType         string   `xml:"DataType"`
	StringListValues []string `xml:"StringListValue>StringListValue"`
	StringValue      string   `xml:"StringValue"`
}

// ReceiveMessageRequest is undocumented.
type ReceiveMessageRequest struct {
	AttributeNames        []string `xml:"AttributeNames>AttributeName"`
	MaxNumberOfMessages   int      `xml:"MaxNumberOfMessages"`
	MessageAttributeNames []string `xml:"MessageAttributeNames>MessageAttributeName"`
	QueueURL              string   `xml:"QueueUrl"`
	VisibilityTimeout     int      `xml:"VisibilityTimeout"`
	WaitTimeSeconds       int      `xml:"WaitTimeSeconds"`
}

// ReceiveMessageResult is undocumented.
type ReceiveMessageResult struct {
	Messages []Message `xml:"ReceiveMessageResult>Messages>Message"`
}

// RemovePermissionRequest is undocumented.
type RemovePermissionRequest struct {
	Label    string `xml:"Label"`
	QueueURL string `xml:"QueueUrl"`
}

// SendMessageBatchRequest is undocumented.
type SendMessageBatchRequest struct {
	Entries  []SendMessageBatchRequestEntry `xml:"Entries>SendMessageBatchRequestEntry"`
	QueueURL string                         `xml:"QueueUrl"`
}

// SendMessageBatchRequestEntry is undocumented.
type SendMessageBatchRequestEntry struct {
	DelaySeconds      int                              `xml:"DelaySeconds"`
	ID                string                           `xml:"Id"`
	MessageAttributes map[string]MessageAttributeValue `xml:"MessageAttribute"`
	MessageBody       string                           `xml:"MessageBody"`
}

// SendMessageBatchResult is undocumented.
type SendMessageBatchResult struct {
	Failed     []BatchResultErrorEntry       `xml:"SendMessageBatchResult>Failed>BatchResultErrorEntry"`
	Successful []SendMessageBatchResultEntry `xml:"SendMessageBatchResult>Successful>SendMessageBatchResultEntry"`
}

// SendMessageBatchResultEntry is undocumented.
type SendMessageBatchResultEntry struct {
	ID                     string `xml:"Id"`
	MD5OfMessageAttributes string `xml:"MD5OfMessageAttributes"`
	MD5OfMessageBody       string `xml:"MD5OfMessageBody"`
	MessageID              string `xml:"MessageId"`
}

// SendMessageRequest is undocumented.
type SendMessageRequest struct {
	DelaySeconds      int                              `xml:"DelaySeconds"`
	MessageAttributes map[string]MessageAttributeValue `xml:"MessageAttribute"`
	MessageBody       string                           `xml:"MessageBody"`
	QueueURL          string                           `xml:"QueueUrl"`
}

// SendMessageResult is undocumented.
type SendMessageResult struct {
	MD5OfMessageAttributes string `xml:"SendMessageResult>MD5OfMessageAttributes"`
	MD5OfMessageBody       string `xml:"SendMessageResult>MD5OfMessageBody"`
	MessageID              string `xml:"SendMessageResult>MessageId"`
}

// SetQueueAttributesRequest is undocumented.
type SetQueueAttributesRequest struct {
	Attributes map[string]string `xml:"Attribute"`
	QueueURL   string            `xml:"QueueUrl"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
