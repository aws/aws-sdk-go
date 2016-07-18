package s3crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws/session"
)

func TestDefaultConfigValues(t *testing.T) {
	kp, err := NewKMSKeyProviderWithMatDesc(session.New(), "{\"kms_cmk_id\":\"\"}")
	assert.Nil(t, err)

	c := New(Authentication(kp), func(c *Client) { c.Config.KMSSession = session.New() })

	assert.NotNil(t, c)
	assert.NotNil(t, c.Config.Mode)
	assert.NotNil(t, c.Config.SaveStrategy)
}
