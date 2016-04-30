package v4_test

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/aws/credentials"
"strings"
)

func TestSignNilCredentials(t *testing.T) {
	signer := v4.NewSigner(nil)
	request, err := http.NewRequest("", "", nil)
	assert.NoError(t, err)

	err = signer.Sign(request)

	assert.Error(t, err)
}

func TestSignNilRequest(t *testing.T) {
	creds := credentials.NewStaticCredentials("id", "secret", "token")
	signer := v4.NewSigner(creds)

	assert.NotNil(t, signer)

	err := signer.Sign(nil)

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

func TestSignPostRequest(t *testing.T) {
	creds := credentials.NewStaticCredentials("id", "secret", "token")

	serviceName := "dynamodb"
	region := "us-west-2"

	endpoint := "https://" + serviceName + "." + region + ".amazonaws.com"

	body := "body"
	reader := strings.NewReader(body)

	req, _ := http.NewRequest("POST", endpoint, reader)

	req.URL.Opaque = "//example.org/bucket/key-._~,!@#$%^&*()"
	req.Header.Add("X-Amz-Target", "prefix.Operation")
	req.Header.Add("Content-Type", "application/x-amz-json-1.0")
	req.Header.Add("Content-Length", string(len(body)))
	req.Header.Add("X-Amz-Meta-Other-Header", "some-value=!@#$%^&* (+)")
	req.Header.Add("X-Amz-Meta-Other-Header_With_Underscore", "some-value=!@#$%^&* (+)")
	req.Header.Add("X-amz-Meta-Other-Header_With_Underscore", "some-value=!@#$%^&* (+)")

	signer := v4.NewSigner(creds)
	err := signer.Sign(req)
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