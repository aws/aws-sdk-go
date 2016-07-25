package s3crypto_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/service/s3/s3crypto"
)

func TestSymmetricKeyProviderGenerates(t *testing.T) {
	kp := &s3crypto.SymmetricKeyProvider{}
	key, err := kp.GenerateKey(32)
	assert.Nil(t, err)
	kp.SetKey(key)
	assert.Equal(t, len(key), len(kp.GetKey()))
	assert.Equal(t, 32, len(kp.GetKey()))
	iv, err := kp.GenerateIV(16)
	assert.Nil(t, err)
	kp.SetIV(iv)
	assert.Equal(t, len(iv), len(kp.GetIV()))
	assert.Equal(t, 16, len(kp.GetIV()))
}
