package v4_test

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/aws/credentials"
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