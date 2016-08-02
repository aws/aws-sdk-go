package s3crypto_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3crypto"
)

func TestDefaultConfigValues(t *testing.T) {
	sess := session.New()
	svc := kms.New(sess)
	handler, err := s3crypto.NewKMSEncryptHandler(svc, "testid", s3crypto.MaterialDescription{})
	assert.Nil(t, err)

	c := s3crypto.NewEncryptionClient(sess, s3crypto.AESGCMContentCipherBuilder(handler))

	assert.NotNil(t, c)
	assert.NotNil(t, c.Config.ContentCipherBuilder)
	assert.NotNil(t, c.Config.SaveStrategy)
}

func TestPutObject(t *testing.T) {
	size := 1024 * 1024
	data := make([]byte, size)
	expected := bytes.Repeat([]byte{1}, size)
	generator := mockGenerator{}
	cb := mockCipherBuilder{generator}
	sess := session.New(&aws.Config{
		MaxRetries: aws.Int(0),
	})
	c := s3crypto.NewEncryptionClient(sess, cb)
	assert.NotNil(t, c)
	input := &s3.PutObjectInput{
		Key:    aws.String("test"),
		Bucket: aws.String("test"),
		Body:   bytes.NewReader(data),
	}
	req, _ := c.PutObjectRequest(input)
	req.Handlers.Send.Clear()
	req.Handlers.Send.PushBack(func(r *request.Request) {
		r.Error = errors.New("stop")
		r.HTTPResponse = &http.Response{
			StatusCode: 200,
		}
	})
	err := req.Send()
	assert.Equal(t, "stop", err.Error())
	b, err := ioutil.ReadAll(req.HTTPRequest.Body)
	assert.NoError(t, err)
	assert.Equal(t, expected, b)
}
