package s3crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws/session"
)

func TestDefaultConfigValues(t *testing.T) {
	sess := session.New()
	kp, err := NewKMSKeyProviderWithMatDesc(sess, "{\"kms_cmk_id\":\"\"}")
	assert.Nil(t, err)

	c := New(sess, Authentication(kp))

	assert.NotNil(t, c)
	assert.NotNil(t, c.Config.Mode)
	assert.NotNil(t, c.Config.SaveStrategy)
}
