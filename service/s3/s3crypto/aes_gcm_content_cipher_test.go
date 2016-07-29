package s3crypto_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/service/s3/s3crypto"
)

type mockHandler struct{}

func (handler mockHandler) GenerateCipherData(keySize, ivSize int) (s3crypto.CipherData, error) {
	return s3crypto.CipherData{
		Key: make([]byte, keySize),
		IV:  make([]byte, ivSize),
	}, nil
}

func (handler mockHandler) EncryptKey(key []byte) ([]byte, error) {
	return key, nil
}

func (handler mockHandler) DecryptKey(key []byte) ([]byte, error) {
	return key, nil
}

func TestAESGCMContentCipherBuilder(t *testing.T) {
	handler := mockHandler{}
	builder := s3crypto.AESGCMContentCipherBuilder(handler)
	assert.NotNil(t, builder)
}

func TestAESGCMContentCipherNewEncryptor(t *testing.T) {
	handler := mockHandler{}
	builder := s3crypto.AESGCMContentCipherBuilder(handler)
	cipher, err := builder.NewEncryptor()
	assert.NoError(t, err)
	assert.NotNil(t, cipher)
}

func TestAESGCMContentCipherGetHandler(t *testing.T) {
	handler := mockHandler{}
	builder := s3crypto.AESGCMContentCipherBuilder(handler)
	cipher, err := builder.NewEncryptor()
	assert.NoError(t, err)
	h := cipher.GetHandler()
	assert.Equal(t, handler, h)
}
