package aws

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/juju/errors"
)

// A set of credentials which can be used for a single request.
// Some RequestCredentials providers may return different credentials over time.
type RequestCredentials struct {
	ID, Secret, Token string
}

// Credentials are used to authenticate and authorize calls that you make to
// AWS.
type Credentials interface {
	Fetch() (RequestCredentials, error)
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
	token := os.Getenv("AWS_SESSION_TOKEN")
	if token != "" {
		return Creds("", "", os.Getenv("AWS_SESSION_TOKEN")), nil
	}

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

	return Creds(id, secret, token), nil
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

func (c *staticCreds) Fetch() (RequestCredentials, error) {
	return RequestCredentials{c.id, c.secret, c.token}, nil
}

type instanceRoleCredentials struct {
	m                sync.Mutex
	expires          time.Time
	ID               string `json:"AccessKeyId"`
	Secret           string `json:"SecretAccessKey"`
	Token            string `json:"Token"`
	ExpirationString string `json:"Expiration"`
	APIResponseCode  string `json:"Code"`
}

func InstanceRoleCredentials() Credentials {
	return &instanceRoleCredentials{}
}

// {
//   "Code" : "Success",
//   "LastUpdated" : "2014-12-15T19:17:56Z",
//   "Type" : "AWS-HMAC",
//   "AccessKeyId" : "",
//   "SecretAccessKey" : "",
//   "Token" : "",
//   "Expiration" : "2014-12-16T01:51:37Z"
// }

var metadataCredentialsEndpoint = "http://169.254.169.254/latest/meta-data/iam/security-credentials/"

// Retrieve credentials from the EC2 Metadata endpoint
func (c *instanceRoleCredentials) obtainCredentialsLazily() error {
	zeroTime := time.Time{}
	if c.expires != zeroTime || -time.Since(c.expires) > 10*time.Second {
		// Reuse existing credentials
		return nil
	}

	c.m.Lock()
	defer c.m.Unlock()

	// Query the security-credentials/ endpoint
	r, err := http.Get(metadataCredentialsEndpoint)
	if err != nil {
		return err
	}

	s := bufio.NewScanner(r.Body)
	s.Scan()
	if s.Err() != nil {
		return s.Err()
	}
	firstLine := s.Text()

	// Query the role that it returns
	r, err = http.Get(metadataCredentialsEndpoint + firstLine)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(c)
	if err != nil {
		return err
	}
	if c.APIResponseCode != "Success" {
		return fmt.Errorf("Metadata endpoint did not succeed. Code: %#+v", c)
	}
	c.expires, err = time.Parse("2006-01-02T15:04:05Z", c.ExpirationString)
	if err != nil {
		return err
	}
	return nil
}

func (c *instanceRoleCredentials) Fetch() (RequestCredentials, error) {
	err := c.obtainCredentialsLazily()
	return RequestCredentials{c.ID, c.Secret, c.Token}, err
}
