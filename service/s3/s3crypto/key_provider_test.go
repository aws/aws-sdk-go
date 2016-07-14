package s3crypto_test

import (
	"encoding/base64"
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

func TestSymmetricKeyDecryption(t *testing.T) {
	mkey, _ := base64.StdEncoding.DecodeString("w1WLio3agRWRTSJK/Ouh8NHoqRQ6fn5WbSXDTHjXMSo=")
	cipher, err := s3crypto.NewAESECB(mkey)
	assert.Nil(t, err)
	kp := &s3crypto.SymmetricKeyProvider{
		Wrap: cipher,
	}
	encryptedKey, _ := base64.StdEncoding.DecodeString("QCwoHJ/cOGmhQeNZ0GAeep+ysKWpqOY7w63kijvBCv+mCQMmX+H4u8HtGLdU3LFj")
	val, err := kp.GetDecryptedKey(encryptedKey)
	assert.Nil(t, err)
	iv, _ := base64.StdEncoding.DecodeString("qxKNPKvYnj28sgP0OQ6ItQ==")
	kp.SetKey(val)
	kp.SetIV(iv)
	_, err = s3crypto.NewAESCBC(kp)
	assert.Nil(t, err)
}
