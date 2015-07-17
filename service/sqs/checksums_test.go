package sqs_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsconv"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/internal/test/unit"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"
)

var _ = unit.Imported

var svc = func() *sqs.SQS {
	s := sqs.New(&aws.Config{
		DisableParamValidation: awsconv.Bool(true),
	})
	s.Handlers.Send.Clear()
	return s
}()

func TestSendMessageChecksum(t *testing.T) {
	req, _ := svc.SendMessageRequest(&sqs.SendMessageInput{
		MessageBody: awsconv.String("test"),
	})
	req.Handlers.Send.PushBack(func(r *aws.Request) {
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.SendMessageOutput{
			MD5OfMessageBody: awsconv.String("098f6bcd4621d373cade4e832627b4f6"),
			MessageID:        awsconv.String("12345"),
		}
	})
	err := req.Send()
	assert.NoError(t, err)
}

func TestSendMessageChecksumInvalid(t *testing.T) {
	req, _ := svc.SendMessageRequest(&sqs.SendMessageInput{
		MessageBody: awsconv.String("test"),
	})
	req.Handlers.Send.PushBack(func(r *aws.Request) {
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.SendMessageOutput{
			MD5OfMessageBody: awsconv.String("000"),
			MessageID:        awsconv.String("12345"),
		}
	})
	err := req.Send()
	assert.Error(t, err)

	assert.Equal(t, "InvalidChecksum", err.(awserr.Error).Code())
	assert.Contains(t, err.(awserr.Error).Message(), "expected MD5 checksum '000', got '098f6bcd4621d373cade4e832627b4f6'")
}

func TestSendMessageChecksumInvalidNoValidation(t *testing.T) {
	s := sqs.New(&aws.Config{
		DisableParamValidation:  awsconv.Bool(true),
		DisableComputeChecksums: awsconv.Bool(true),
	})
	s.Handlers.Send.Clear()

	req, _ := s.SendMessageRequest(&sqs.SendMessageInput{
		MessageBody: awsconv.String("test"),
	})
	req.Handlers.Send.PushBack(func(r *aws.Request) {
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.SendMessageOutput{
			MD5OfMessageBody: awsconv.String("000"),
			MessageID:        awsconv.String("12345"),
		}
	})
	err := req.Send()
	assert.NoError(t, err)
}

func TestSendMessageChecksumNoInput(t *testing.T) {
	req, _ := svc.SendMessageRequest(&sqs.SendMessageInput{})
	req.Handlers.Send.PushBack(func(r *aws.Request) {
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.SendMessageOutput{}
	})
	err := req.Send()
	assert.Error(t, err)

	assert.Equal(t, "InvalidChecksum", err.(awserr.Error).Code())
	assert.Contains(t, err.(awserr.Error).Message(), "cannot compute checksum. missing body")
}

func TestSendMessageChecksumNoOutput(t *testing.T) {
	req, _ := svc.SendMessageRequest(&sqs.SendMessageInput{
		MessageBody: awsconv.String("test"),
	})
	req.Handlers.Send.PushBack(func(r *aws.Request) {
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.SendMessageOutput{}
	})
	err := req.Send()
	assert.Error(t, err)

	assert.Equal(t, "InvalidChecksum", err.(awserr.Error).Code())
	assert.Contains(t, err.(awserr.Error).Message(), "cannot verify checksum. missing response MD5")
}

func TestRecieveMessageChecksum(t *testing.T) {
	req, _ := svc.ReceiveMessageRequest(&sqs.ReceiveMessageInput{})
	req.Handlers.Send.PushBack(func(r *aws.Request) {
		md5 := "098f6bcd4621d373cade4e832627b4f6"
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.ReceiveMessageOutput{
			Messages: []*sqs.Message{
				{Body: awsconv.String("test"), MD5OfBody: &md5},
				{Body: awsconv.String("test"), MD5OfBody: &md5},
				{Body: awsconv.String("test"), MD5OfBody: &md5},
				{Body: awsconv.String("test"), MD5OfBody: &md5},
			},
		}
	})
	err := req.Send()
	assert.NoError(t, err)
}

func TestRecieveMessageChecksumInvalid(t *testing.T) {
	req, _ := svc.ReceiveMessageRequest(&sqs.ReceiveMessageInput{})
	req.Handlers.Send.PushBack(func(r *aws.Request) {
		md5 := "098f6bcd4621d373cade4e832627b4f6"
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.ReceiveMessageOutput{
			Messages: []*sqs.Message{
				{Body: awsconv.String("test"), MD5OfBody: &md5},
				{Body: awsconv.String("test"), MD5OfBody: awsconv.String("000"), MessageID: awsconv.String("123")},
				{Body: awsconv.String("test"), MD5OfBody: awsconv.String("000"), MessageID: awsconv.String("456")},
				{Body: awsconv.String("test"), MD5OfBody: &md5},
			},
		}
	})
	err := req.Send()
	assert.Error(t, err)

	assert.Equal(t, "InvalidChecksum", err.(awserr.Error).Code())
	assert.Contains(t, err.(awserr.Error).Message(), "invalid messages: 123, 456")
}

func TestSendMessageBatchChecksum(t *testing.T) {
	req, _ := svc.SendMessageBatchRequest(&sqs.SendMessageBatchInput{
		Entries: []*sqs.SendMessageBatchRequestEntry{
			{ID: awsconv.String("1"), MessageBody: awsconv.String("test")},
			{ID: awsconv.String("2"), MessageBody: awsconv.String("test")},
			{ID: awsconv.String("3"), MessageBody: awsconv.String("test")},
			{ID: awsconv.String("4"), MessageBody: awsconv.String("test")},
		},
	})
	req.Handlers.Send.PushBack(func(r *aws.Request) {
		md5 := "098f6bcd4621d373cade4e832627b4f6"
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.SendMessageBatchOutput{
			Successful: []*sqs.SendMessageBatchResultEntry{
				{MD5OfMessageBody: &md5, MessageID: awsconv.String("123"), ID: awsconv.String("1")},
				{MD5OfMessageBody: &md5, MessageID: awsconv.String("456"), ID: awsconv.String("2")},
				{MD5OfMessageBody: &md5, MessageID: awsconv.String("789"), ID: awsconv.String("3")},
				{MD5OfMessageBody: &md5, MessageID: awsconv.String("012"), ID: awsconv.String("4")},
			},
		}
	})
	err := req.Send()
	assert.NoError(t, err)
}

func TestSendMessageBatchChecksumInvalid(t *testing.T) {
	req, _ := svc.SendMessageBatchRequest(&sqs.SendMessageBatchInput{
		Entries: []*sqs.SendMessageBatchRequestEntry{
			{ID: awsconv.String("1"), MessageBody: awsconv.String("test")},
			{ID: awsconv.String("2"), MessageBody: awsconv.String("test")},
			{ID: awsconv.String("3"), MessageBody: awsconv.String("test")},
			{ID: awsconv.String("4"), MessageBody: awsconv.String("test")},
		},
	})
	req.Handlers.Send.PushBack(func(r *aws.Request) {
		md5 := "098f6bcd4621d373cade4e832627b4f6"
		body := ioutil.NopCloser(bytes.NewReader([]byte("")))
		r.HTTPResponse = &http.Response{StatusCode: 200, Body: body}
		r.Data = &sqs.SendMessageBatchOutput{
			Successful: []*sqs.SendMessageBatchResultEntry{
				{MD5OfMessageBody: &md5, MessageID: awsconv.String("123"), ID: awsconv.String("1")},
				{MD5OfMessageBody: awsconv.String("000"), MessageID: awsconv.String("456"), ID: awsconv.String("2")},
				{MD5OfMessageBody: awsconv.String("000"), MessageID: awsconv.String("789"), ID: awsconv.String("3")},
				{MD5OfMessageBody: &md5, MessageID: awsconv.String("012"), ID: awsconv.String("4")},
			},
		}
	})
	err := req.Send()
	assert.Error(t, err)

	assert.Equal(t, "InvalidChecksum", err.(awserr.Error).Code())
	assert.Contains(t, err.(awserr.Error).Message(), "invalid messages: 456, 789")
}
