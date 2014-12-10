package aws

import (
	"os"

	"github.com/juju/errors"
)

// Credentials are used to authenticate and authorize calls that you make to
// AWS.
type Credentials interface {
	AccessKeyID() string
	SecretAccessKey() string
	SecurityToken() string
}

var (
	// ErrAccessKeyIDNotFound is returned when the AWS Access Key ID can't be
	// found in the process's environment.
	ErrAccessKeyIDNotFound = errors.NotFoundf("AWS_ACCESS_KEY_ID or AWS_ACCESS_KEY not found in environment")
	// ErrSecretAccessKeyNotFound is returned when the AWS Secret Access Key
	// can't be found in the process's environment.
	ErrSecretAccessKeyNotFound = errors.NotFoundf("AWS_SECRET_ACCESS_KEY or AWS_SECRET_KEY not found in environment")
)

// EnvCreds returns the AWS credentials from the process's environment, or an
// error if none are found.
func EnvCreds() (Credentials, error) {
	id := os.Getenv("AWS_ACCESS_KEY_ID")
	if id == "" {
		id = os.Getenv("AWS_ACCESS_KEY")
	}

	if id == "" {
		return nil, ErrAccessKeyIDNotFound
	}

	secret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if secret == "" {
		secret = os.Getenv("AWS_SECRET_KEY")
	}

	if secret == "" {
		return nil, ErrSecretAccessKeyNotFound
	}

	return Creds(id, secret, ""), nil
}

// Creds returns a static set of credentials.
func Creds(accessKeyID, secretAccessKey, securityToken string) Credentials {
	return &staticCreds{
		id:     accessKeyID,
		secret: secretAccessKey,
		token:  securityToken,
	}
}

type staticCreds struct {
	id     string
	secret string
	token  string
}

func (c *staticCreds) AccessKeyID() string {
	return c.id
}

func (c *staticCreds) SecretAccessKey() string {
	return c.secret
}

func (c *staticCreds) SecurityToken() string {
	return c.token
}
