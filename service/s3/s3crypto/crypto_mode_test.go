package s3crypto

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws/session"
)

func TestKeyProviderFactory(t *testing.T) {
	cfg := Config{
		MasterKey:  []byte("00000000000000000000000000000000"),
		KMSSession: session.New(),
	}
	env := Envelope{
		WrapAlg: "kms",
		MatDesc: "{\"kms_cmk_id\":\"\"}",
	}
	kp, err := keyProviderFactory(&env, cfg)
	_, ok := kp.(*KMSKeyProvider)
	assert.Nil(t, err)
	assert.NotNil(t, kp)
	assert.True(t, ok)
}

func TestCEKFactory(t *testing.T) {
	env := Envelope{
		CEKAlg:  "AES/GCM/NoPadding",
		MatDesc: "{\"kms_cmk_id\":\"\"}",
	}
	iv, _ := hex.DecodeString("00000000000000000000000000000000")
	key, _ := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
	cek, err := cekFactory(&env, &SymmetricKeyProvider{key: key, iv: iv})
	_, ok := cek.(*AESGCM)
	assert.Nil(t, err)
	assert.NotNil(t, cek)
	assert.True(t, ok)
}
