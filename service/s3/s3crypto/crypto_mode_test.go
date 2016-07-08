package s3crypto

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModeFactory(t *testing.T) {
	masterkey, _ := base64.StdEncoding.DecodeString("w1WLio3agRWRTSJK/Ouh8NHoqRQ6fn5WbSXDTHjXMSo=")
	cfg := Config{
		MasterKey: masterkey,
	}
	env := Envelope{
		WrapAlg:   "ecb",
		CipherKey: "QCwoHJ/cOGmhQeNZ0GAeep+ysKWpqOY7w63kijvBCv+mCQMmX+H4u8HtGLdU3LFj",
		IV:        "qxKNPKvYnj28sgP0OQ6ItQ==",
	}

	mode, err := modeFactory(&env, cfg)
	assert.Nil(t, err)

	_, ok := mode.(*decryptionMode)
	assert.True(t, ok)
}

func TestKeyProviderFactory(t *testing.T) {
	cfg := Config{
		MasterKey: []byte("00000000000000000000000000000000"),
	}
	env := Envelope{
		WrapAlg: "ecb",
	}
	kp, err := keyProviderFactory(&env, cfg)
	_, ok := kp.(*SymmetricKeyProvider)
	assert.Nil(t, err)
	assert.NotNil(t, kp)
	assert.True(t, ok)
}

func TestCEKFactory(t *testing.T) {
	env := Envelope{
		CEKAlg: "AES/CBC/PKCS5Padding",
	}
	iv, _ := hex.DecodeString("00000000000000000000000000000000")
	key, _ := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
	cek, err := cekFactory(&env, &SymmetricKeyProvider{key: key, iv: iv})
	_, ok := cek.(*AESCBC)
	assert.Nil(t, err)
	assert.NotNil(t, cek)
	assert.True(t, ok)
}

func TestEncodeMeta(t *testing.T) {
	t.Log("TODO")
}

func TestDecodeMetaV1(t *testing.T) {
	masterkey, _ := base64.StdEncoding.DecodeString("w1WLio3agRWRTSJK/Ouh8NHoqRQ6fn5WbSXDTHjXMSo=")
	cipher, err := NewAESECB(masterkey)
	assert.Nil(t, err)
	kp := NewSymmetricKeyProvider(cipher)
	env := Envelope{
		CipherKey: "QCwoHJ/cOGmhQeNZ0GAeep+ysKWpqOY7w63kijvBCv+mCQMmX+H4u8HtGLdU3LFj",
		IV:        "qxKNPKvYnj28sgP0OQ6ItQ==",
		MatDesc:   "{}",
		version:   1,
	}
	assert.Nil(t, err)

	err = DecodeMeta(&env, kp)
	assert.Nil(t, err)

	expectedKey, _ := hex.DecodeString("783b935dca5f2328294dd56243b0b4f516014129b5ff6f145d4bb3f59c7abcae")
	expectedIV, _ := base64.StdEncoding.DecodeString("qxKNPKvYnj28sgP0OQ6ItQ==")
	assert.Nil(t, err)
	assert.Equal(t, len(expectedKey), len([]byte(env.CipherKey)))
	assert.Equal(t, expectedKey, []byte(env.CipherKey))

	assert.Equal(t, len(expectedIV), len([]byte(env.IV)))
	assert.Equal(t, expectedIV, []byte(env.IV))
}
