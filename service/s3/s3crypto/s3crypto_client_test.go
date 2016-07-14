package s3crypto

import (
	"bytes"
	"encoding/base64"
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

	t.Log("WEEEEEE", []byte(body))

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

func TestGetObject_V1_WRAP_ECB_CONTENT_CBC(t *testing.T) {
	mkey, _ := base64.StdEncoding.DecodeString("w1WLio3agRWRTSJK/Ouh8NHoqRQ6fn5WbSXDTHjXMSo=")
	ciphertext, _ := hex.DecodeString("bb6d801fa7bc7ed756db8d69f9db17ee406af3f32e8800fc39f10291e682509e781641cd03b9d8bd77332080fad72857e3ddbdd88c70862e6f41b46f5e2920d249fe2ae911a50fe609a1833beaa0ba9a")
	cipher, err := NewAESECB([]byte(mkey))
	assert.Nil(t, err)
	c := New(EncryptionOnly(NewSymmetricKeyProvider(cipher)), func(c *Client) {
		c.Config.S3Session = session.New()
		c.Config.MasterKey = mkey
	})

	key := "test-key"
	input := &s3.GetObjectInput{
		Bucket: aws.String("test-bucket"),
		Key:    &key,
	}

	// master key w1WLio3agRWRTSJK/Ouh8NHoqRQ6fn5WbSXDTHjXMSo=
	req, out := c.GetObjectRequest(input)
	assert.NotNil(t, out)
	assert.NotNil(t, req)
	body := ioutil.NopCloser(bytes.NewBuffer(ciphertext))
	req.HTTPResponse = &http.Response{
		StatusCode: 200,
		Header: http.Header{
			"X-Amz-Meta-X-Amz-Iv":      []string{"qxKNPKvYnj28sgP0OQ6ItQ=="},
			"Accept-Ranges":            []string{"bytes"},
			"X-Amz-Id-2":               []string{"JGEqnzI5B2bK2lJHJb33BWmzLsGriR3dUbXGHCsLCNkZXa50gLJeswAA2SXtc/R0D4ju+kA2pjU="},
			"X-Amz-Request-Id":         []string{"C3B8D582056B4420"},
			"Date":                     []string{"Fri, 01 Jul 2016 17:31:14 GMT"},
			"X-Amz-Meta-X-Amz-Key":     []string{"QCwoHJ/cOGmhQeNZ0GAeep+ysKWpqOY7w63kijvBCv+mCQMmX+H4u8HtGLdU3LFj"},
			"X-Amz-Meta-X-Amz-Matdesc": []string{"{}"},
			"Last-Modified":            []string{"Fri, 01 Jul 2016 17:31:14 GMT"},
			"Etag":                     []string{"5f257b03bf10e2a74ab08dde2e66c59b"},
			"Content-Type":             []string{"application/octet-stream"},
			"Content-Length":           []string{"80"},
		},
		Body: body,
	}

	req.Handlers.Send.Clear()

	err = req.Send()
	assert.Nil(t, err)
	plaintext, err := ioutil.ReadAll(out.Body)
	assert.Nil(t, err)
	expectedPlaintext, _ := hex.DecodeString("34c9e4da626670368a1ca2d371309b7d1bc5bbe32f66cad0bad61bf3f12e7d0cae732165bff9acadfa8ad68a2d0249498108f6488477ac0836b4c2f3db0d982a")
	fmt.Printf("DATA\n%x\n%x\n", expectedPlaintext, plaintext)
	assert.Equal(t, len(expectedPlaintext), len(plaintext))
	assert.Equal(t, expectedPlaintext, plaintext)
}
