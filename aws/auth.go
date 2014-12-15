package aws

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"

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

func (c *staticCreds) AccessKeyID() string {
	return c.id
}

func (c *staticCreds) SecretAccessKey() string {
	return c.secret
}

func (c *staticCreds) SecurityToken() string {
	return c.token
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
func (c *instanceRoleCredentials) obtainCredentialsLazily() {
	// TODO: Do we need to refresh?

	// TODO: Need to loop over entries at /services/../, or pick the first line...
	r, err := http.Get("http://169.254.169.254/latest/meta-data/iam/security-credentials/services")
	if err != nil {
		// Nowhere else to put it right now.
		panic(err)
	}
	defer r.Close()
	decoder := json.Decoder(r.Body)
	err = decoder.Decode(c)
	if err != nil {
		panic(err)
	}
	if c.apiResponseCode != "success" {
		log.Panicln("Error fetching code:", c.apiResponseCode)
	}
	log.Println("Got a key:", c.AccessKeyID())
	// decoder
}

func (c *instanceRoleCredentials) AccessKeyID() string {
	c.obtainCredentialsLazily()
	return c.id
}

func (c *instanceRoleCredentials) SecretAccessKey() string {
	c.obtainCredentialsLazily()
	return c.secret
}

func (c *instanceRoleCredentials) SecurityToken() string {
	c.obtainCredentialsLazily()
	return c.token
}
