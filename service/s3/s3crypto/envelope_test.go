package s3crypto

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
)

func TestGetV1Envelope(t *testing.T) {
	handler, err := NewKMSKeyProvider(session.New(), "", MaterialDescription{})
	assert.NoError(t, err)
	c := New(nil, AESGCMContentCipherBuilder(handler), func(c *Client) { c.Config.S3Session = session.New() })
	env, err := c.getEnvelope(nil, &request.Request{
		HTTPResponse: &http.Response{
			Header: http.Header{
				"X-Amz-Meta-X-Amz-Key": []string{"9adc8fbd506e032af7fa20cf5343719de6d1288c158c63d6878aaf64ce26ca85"},
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, env.version)
	assert.Equal(t, "9adc8fbd506e032af7fa20cf5343719de6d1288c158c63d6878aaf64ce26ca85", env.CipherKey)
}

func TestGetV2Envelope(t *testing.T) {
	handler, err := NewKMSKeyProvider(session.New(), "", MaterialDescription{})
	assert.NoError(t, err)
	c := New(nil, AESGCMContentCipherBuilder(handler), func(c *Client) { c.Config.S3Session = session.New() })
	env, err := c.getEnvelope(nil, &request.Request{
		HTTPResponse: &http.Response{
			Header: http.Header{
				"X-Amz-Meta-X-Amz-Key-V2": []string{"9adc8fbd506e032af7fa20cf5343719de6d1288c158c63d6878aaf64ce26ca85"},
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, 2, env.version)
	assert.Equal(t, "9adc8fbd506e032af7fa20cf5343719de6d1288c158c63d6878aaf64ce26ca85", env.CipherKey)
}
