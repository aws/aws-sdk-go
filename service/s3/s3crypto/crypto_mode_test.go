package s3crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws/session"
)

func TestKeyProviderFactory(t *testing.T) {
	cfg := Config{
		KMSSession: session.New(),
	}
	env := Envelope{
		WrapAlg: "kms",
		MatDesc: `{"kms_cmk_id":""}`,
	}
	kp, err := keyProviderForEnvelope(&env, cfg)
	_, ok := kp.(*KMSKeyProvider)
	assert.Nil(t, err)
	assert.NotNil(t, kp)
	assert.True(t, ok)
}

func TestCEKFactory(t *testing.T) {
	env := Envelope{
		CEKAlg:  AESGCMNoPadding,
		MatDesc: `{"kms_cmk_id":""}`,
	}

	iv := make([]byte, 16)
	key := make([]byte, 32)
	cek, err := cekForEnvelope(&env, &SymmetricKeyProvider{key: key, iv: iv})
	_, ok := cek.(*AESGCM)
	assert.Nil(t, err)
	assert.NotNil(t, cek)
	assert.True(t, ok)
}
