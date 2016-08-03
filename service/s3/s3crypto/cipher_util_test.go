package s3crypto

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func TestWrapFactory(t *testing.T) {
	cfg := DecryptionConfig{
		KMSClient: kms.New(session.New()),
	}
	env := Envelope{
		WrapAlg: "kms",
		MatDesc: `{"kms_cmk_id":""}`,
	}
	wrap, err := wrapFromEnvelope(&env, cfg)
	_, ok := wrap.(*KMSKeyHandler)
	assert.Nil(t, err)
	assert.NotNil(t, wrap)
	assert.True(t, ok)
}

func TestCEKFactory(t *testing.T) {
	key, _ := hex.DecodeString("31bdadd96698c204aa9ce1448ea94ae1fb4a9a0b3c9d773b51bb1822666b8f22")
	keyB64 := base64.URLEncoding.EncodeToString(key)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, fmt.Sprintf("%s%s%s", `{"KeyId":"test-key-id","Plaintext":"`, keyB64, `"}`))
	}))
	defer ts.Close()
	sess := session.New(&aws.Config{
		MaxRetries:       aws.Int(0),
		Endpoint:         aws.String(ts.URL[7:]),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	})

	handler, err := NewKMSDecryptHandler(kms.New(sess), `{"kms_cmk_id": "test"}`)
	assert.NoError(t, err)

	iv, err := hex.DecodeString("0d18e06c7c725ac9e362e1ce")
	assert.NoError(t, err)
	ivB64 := base64.URLEncoding.EncodeToString(iv)

	cipherKey, err := hex.DecodeString("31bdadd96698c204aa9ce1448ea94ae1fb4a9a0b3c9d773b51bb1822666b8f22")
	assert.NoError(t, err)
	cipherKeyB64 := base64.URLEncoding.EncodeToString(cipherKey)

	env := Envelope{
		WrapAlg:   "kms",
		CEKAlg:    AESGCMNoPadding,
		CipherKey: cipherKeyB64,
		IV:        ivB64,
		MatDesc:   `{"kms_cmk_id":""}`,
	}
	cek, err := cekFromEnvelope(&env, handler)
	assert.NoError(t, err)
	assert.NotNil(t, cek)
}
