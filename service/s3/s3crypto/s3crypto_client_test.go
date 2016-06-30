package s3crypto

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func TestDefaultConfigValues(t *testing.T) {
	mkey, _ := hex.DecodeString("2b7e151628aed2a6abf7158809cf4f3c")
	cipher, err := NewAESECB([]byte(mkey))
	assert.Nil(t, err)
	c := New(EncryptionOnly(NewSymmetricKeyProvider(cipher)), func(c *Client) { c.Config.S3Session = session.New() })

	assert.NotNil(t, c)
	assert.NotNil(t, c.Config.Mode)
	assert.NotNil(t, c.Config.SaveStrategy)
}

func TestPutObject(t *testing.T) {
	mkey, _ := hex.DecodeString("2b7e151628aed2a6abf7158809cf4f3c")
	cipher, err := NewAESECB([]byte(mkey))
	assert.Nil(t, err)
	c := New(EncryptionOnly(NewSymmetricKeyProvider(cipher)), func(c *Client) { c.Config.S3Session = session.New() })

	key := "test-key"
	body := "test body"
	input := &s3.PutObjectInput{
		Bucket: aws.String("test-bucket"),
		Key:    &key,
		Body:   strings.NewReader(body),
	}

	req, out := c.PutObjectRequest(input)

	assert.NotNil(t, out)
	assert.NotNil(t, req)

	req.Handlers.Send.Clear()
	req.Handlers.Unmarshal.Clear()
	req.Handlers.UnmarshalMeta.Clear()
	req.Handlers.UnmarshalError.Clear()
	req.HTTPResponse = &http.Response{
		StatusCode: 200,
	}
	req.Handlers.Send.PushBack(func(r *request.Request) {
		assert.NotEmpty(t, *input.Metadata["X-Amz-Key-V2"])
		assert.NotEmpty(t, *input.Metadata["X-Amz-Unencrypted-Content-Md5"])
		assert.Equal(t, *input.Metadata["X-Amz-Unencrypted-Content-Length"], fmt.Sprintf("%d", len(body)))
		b, err := ioutil.ReadAll(r.HTTPRequest.Body)
		assert.Nil(t, err)
		assert.NotEqual(t, string(b), body)
		assert.NotEmpty(t, b)
	})
	err = req.Send()
	assert.Nil(t, err)
}

func TestGetObject(t *testing.T) {
}
