package v4_test

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/aws/credentials"
"strings"
	"fmt"
)

func TestSignNilCredentials(t *testing.T) {
	signer := v4.NewSigner(nil)
	request, err := http.NewRequest("", "", nil)
	assert.NoError(t, err)

	err = signer.Sign(request, false)

	assert.Error(t, err)
}

func TestSignNilRequest(t *testing.T) {
	creds := credentials.NewStaticCredentials("id", "secret", "token")
	signer := v4.NewSigner(creds)

	assert.NotNil(t, signer)

	err := signer.Sign(nil, false)

	assert.Error(t, err)
}

func TestVerifyFailsNoCredentials(t *testing.T) {
	verifier := v4.NewVerifier(nil)
	request, err := http.NewRequest("", "", nil)
	assert.NoError(t, err)

	verified, err := verifier.Verify(request)

	assert.Error(t, err)
	assert.False(t, verified)
}

func TestVerifyFailsNoRequest(t *testing.T) {
	creds := credentials.NewStaticCredentials("id", "secret", "token")
	verifier := v4.NewVerifier(creds)
	verified, err := verifier.Verify(nil)

	assert.Error(t, err)
	assert.False(t, verified)
}

type Headers map[string]string

func request(method, service, region, body, url string, headers map[string]string) *http.Request {
	endpoint := fmt.Sprintf("https://%s.%s.amazonaws.com", service, region)
	req, _ := http.NewRequest(method, endpoint, strings.NewReader(body))
	req.URL.Opaque = url
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.Header.Add("Content-Length", string(len(body)))
	return req
}

func TestSignPostRequest(t *testing.T) {
	creds := credentials.NewStaticCredentials("AKID", "SECRET", "SESSION")

	headers := make(Headers)
	headers["X-Amz-Target"] = "prefix.Operation"
	headers["Content-Type"] = "application/x-amz-json-1.0"
	headers["X-Amz-Meta-Other-Header"] = "some-value=!@#$%^&* (+)"
	headers["X-Amz-Meta-Other-Header_With_Underscore"] = "some-value=!@#$%^&* (+)"
	headers["X-amz-Meta-Other-Header_With_Underscore"] = "some-value=!@#$%^&* (+)"

	req := request(
		"POST",
		"dynamodb",
		"us-east-1",
		"",
		"//example.org/bucket/key-._~,!@#$%^&*()",
		headers)

	signer := v4.NewSigner(creds)
	err := signer.Sign(req, true)
	assert.NoError(t, err)

	expectedDate := "19700101T000000Z"
	expectedHeaders := "content-length;content-type;host;x-amz-meta-other-header;x-amz-meta-other-header_with_underscore"
	expectedSig := "ea7856749041f727690c580569738282e99c79355fe0d8f125d3b5535d2ece83"
	expectedCred := "AKID/19700101/us-east-1/dynamodb/aws4_request"
	expectedTarget := "prefix.Operation"

	q := req.URL.Query()
	assert.Equal(t, expectedSig, q.Get("X-Amz-Signature"))
	assert.Equal(t, expectedCred, q.Get("X-Amz-Credential"))
	assert.Equal(t, expectedHeaders, q.Get("X-Amz-SignedHeaders"))
	assert.Equal(t, expectedDate, q.Get("X-Amz-Date"))
	assert.Empty(t, q.Get("X-Amz-Meta-Other-Header"))
	assert.Equal(t, expectedTarget, q.Get("X-Amz-Target"))
}

func TestSignOther(t *testing.T) {
	creds := credentials.NewStaticCredentials("AKID", "SECRET", "SESSION")

	req := request(
		"POST",
		"dynamodb",
		"us-east-1",
		"{}",
		"//example.org/bucket/key-._~,!@#$%^&*()",
		make(Headers))

	signer := v4.NewSigner(creds)
	signer.Sign(req, false)

	expectedDate := "19700101T000000Z"
	expectedSig := "AWS4-HMAC-SHA256 Credential=AKID/19700101/us-east-1/dynamodb/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date;x-amz-meta-other-header;x-amz-meta-other-header_with_underscore;x-amz-security-token;x-amz-target, Signature=ea766cabd2ec977d955a3c2bae1ae54f4515d70752f2207618396f20aa85bd21"

	q := req.Header
	assert.Equal(t, expectedSig, q.Get("Authorization"))
	assert.Equal(t, expectedDate, q.Get("X-Amz-Date"))
}

