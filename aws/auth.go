package aws

import (
	"os"

	"github.com/juju/errors"
)

type Credentials interface {
	AccessKeyID() string
	SecretAccessKey() string
	SecurityToken() string
}

var (
	ErrAccessKeyIDNotFound     = errors.NotFoundf("AWS_ACCESS_KEY_ID or AWS_ACCESS_KEY not found in environment")
	ErrSecretAccessKeyNotFound = errors.NotFoundf("AWS_SECRET_ACCESS_KEY or AWS_SECRET_KEY not found in environment")
)

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

	return StaticCreds(id, secret, ""), nil
}

func StaticCreds(accessKeyID, secretAccessKey, securityToken string) Credentials {
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
