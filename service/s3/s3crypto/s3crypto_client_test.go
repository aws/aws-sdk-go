package s3crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws/session"
)

func TestDefaultConfigValues(t *testing.T) {
	sess := session.New()
	handler, err := NewKMSDecryptHandler(sess, "{\"kms_cmk_id\":\"\"}")
	assert.Nil(t, err)

	c := New(sess, AESGCMContentCipherBuilder(handler))

	assert.NotNil(t, c)
	assert.NotNil(t, c.Config.ContentCipherBuilder)
	assert.NotNil(t, c.Config.SaveStrategy)
}
