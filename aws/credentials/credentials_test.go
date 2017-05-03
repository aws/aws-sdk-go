package credentials

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/stretchr/testify/assert"
)

type stubProvider struct {
	creds   Value
	expired bool
	err     error
}

func (s *stubProvider) Retrieve() (Value, error) {
	s.expired = false
	s.creds.ProviderName = "stubProvider"
	return s.creds, s.err
}
func (s *stubProvider) IsExpired() bool {
	return s.expired
}

func TestCredentialsGet(t *testing.T) {
	c := NewCredentials(&stubProvider{
		creds: Value{
			AccessKeyID:     "AKID",
			SecretAccessKey: "SECRET",
			SessionToken:    "",
		},
		expired: true,
	})

	creds, err := c.Get()
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, "AKID", creds.AccessKeyID, "Expect access key ID to match")
	assert.Equal(t, "SECRET", creds.SecretAccessKey, "Expect secret access key to match")
	assert.Empty(t, creds.SessionToken, "Expect session token to be empty")
}

func TestCredentialsGetWithError(t *testing.T) {
	c := NewCredentials(&stubProvider{err: awserr.New("provider error", "", nil), expired: true})

	_, err := c.Get()
	assert.Equal(t, "provider error", err.(awserr.Error).Code(), "Expected provider error")
}

func TestCredentialsExpire(t *testing.T) {
	stub := &stubProvider{}
	c := NewCredentials(stub)

	stub.expired = false
	assert.True(t, c.IsExpired(), "Expected to start out expired")
	c.Expire()
	assert.True(t, c.IsExpired(), "Expected to be expired")

	c.forceRefresh = false
	assert.False(t, c.IsExpired(), "Expected not to be expired")

	stub.expired = true
	assert.True(t, c.IsExpired(), "Expected to be expired")
}

func TestCredentialsGetWithProviderName(t *testing.T) {
	stub := &stubProvider{}

	c := NewCredentials(stub)

	creds, err := c.Get()
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, creds.ProviderName, "stubProvider", "Expected provider name to match")
}

func TestRetrieve(t *testing.T) {
	assert := assert.New(t)
	e := EnvProvider{}

	// No envs set
	_, err := e.Retrieve()
	assert.Equal(ErrAccessKeyIDNotFound, err)

	// Setting the access key only
	defer func(v string) { os.Setenv("AWS_ACCESS_KEY", v) }(os.Getenv("AWS_ACCESS_KEY"))
	os.Setenv("AWS_ACCESS_KEY", "1")

	_, err = e.Retrieve()
	assert.Equal(ErrSecretAccessKeyNotFound, err)

	// Setting all of them
	defer func(v string) { os.Setenv("AWS_SECRET_KEY", v) }(os.Getenv("AWS_SECURITY_TOKEN"))
	os.Setenv("AWS_SECRET_KEY", "2")
	defer func(v string) { os.Setenv("AWS_SECURITY_TOKEN", v) }(os.Getenv("AWS_SECURITY_TOKEN"))
	os.Setenv("AWS_SECURITY_TOKEN", "3")
	v, err := e.Retrieve()
	assert.NoError(err)

	assert.Equal("1", v.AccessKeyID)
	assert.Equal("2", v.SecretAccessKey)
	assert.Equal("3", v.SessionToken)
}
