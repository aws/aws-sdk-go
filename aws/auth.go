package aws

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"

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
	id               string `json:"AccessKeyId"`
	secret           string `json:"SecretAccessKey"`
	token            string `json:"Token"`
	expirationString string `json:"Expiration"`
	apiResponseCode  string `json:"Code"`
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

// Retrieve credentials from the EC2 Metadata endpoint
func (c *instanceRoleCredentials) obtainCredentialsLazily() error {
	// TODO: Do we need to refresh?

	// TODO: Need to loop over entries at /services/../, or pick the first line...
	r, err := http.Get("http://169.254.169.254/latest/meta-data/iam/security-credentials/services")
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(c)
	if err != nil {
		return err
	}
	if c.apiResponseCode != "success" {
		log.Panicln("Error fetching code:", c.apiResponseCode)
	}
	log.Println("Got credentials:", c.id)
	return nil
}

func (c *instanceRoleCredentials) Fetch() (RequestCredentials, error) {
	err := c.obtainCredentialsLazily()
	return RequestCredentials{c.id, c.secret, c.token}, err
}
