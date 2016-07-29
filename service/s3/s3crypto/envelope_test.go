package s3crypto

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func TestGetV1Envelope(t *testing.T) {
	handler, err := NewKMSEncryptHandler(session.New(), "", MaterialDescription{})
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
	handler, err := NewKMSEncryptHandler(session.New(), "", MaterialDescription{})
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

func TestHeaderSaveStrategy(t *testing.T) {
	env := Envelope{
		CipherKey:             "123",
		IV:                    "456",
		MatDesc:               "matdesc",
		WrapAlg:               "wrap",
		CEKAlg:                "cek",
		TagLen:                "7",
		UnencryptedMD5:        "1234567",
		UnencryptedContentLen: "8",
	}
	strat := headerSaveStrategy{}
	input := &s3.PutObjectInput{}
	expected := map[string]*string{
		http.CanonicalHeaderKey(keyV2Header):                    aws.String("123"),
		http.CanonicalHeaderKey(ivHeader):                       aws.String("456"),
		http.CanonicalHeaderKey(matDescHeader):                  aws.String("matdesc"),
		http.CanonicalHeaderKey(wrapAlgorithmHeader):            aws.String("wrap"),
		http.CanonicalHeaderKey(cekAlgorithmHeader):             aws.String("cek"),
		http.CanonicalHeaderKey(tagLengthHeader):                aws.String("7"),
		http.CanonicalHeaderKey(unencryptedMD5Header):           aws.String("1234567"),
		http.CanonicalHeaderKey(unencryptedContentLengthHeader): aws.String("8"),
	}
	err := strat.Save(env, input)
	assert.NoError(t, err)
	assert.Equal(t, expected, input.Metadata)
}
