package s3crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws/session"
)

func TestWrapFactory(t *testing.T) {
	cfg := Config{
		KMSSession: session.New(),
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
