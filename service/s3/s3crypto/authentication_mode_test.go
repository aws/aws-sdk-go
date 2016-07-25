package s3crypto_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/service/s3/s3crypto"
)

func TestAuthenticationModeCipherName(t *testing.T) {
	mode := s3crypto.Authentication(&s3crypto.SymmetricKeyProvider{})
	assert.Equal(t, s3crypto.AESGCMNoPadding, mode.GetCipherName())
}
