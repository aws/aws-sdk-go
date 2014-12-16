package aws

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/juju/errors"
)

// Credentials are used to authenticate and authorize calls that you make to
// AWS.
type Credentials struct {
	AccessKeyID     string
	SecretAccessKey string
	SecurityToken   string
}

// A CredentialsProvider is a provider of credentials.
type CredentialsProvider interface {
	// Credentials returns a set of credentials (or an error if no credentials
	// could be provided).
	Credentials() (*Credentials, error)
}

var (
	// ErrAccessKeyIDNotFound is returned when the AWS Access Key ID can't be
	// found in the process's environment.
	ErrAccessKeyIDNotFound = errors.NotFoundf("AWS_ACCESS_KEY_ID or AWS_ACCESS_KEY not found in environment")
	// ErrSecretAccessKeyNotFound is returned when the AWS Secret Access Key
	// can't be found in the process's environment.
	ErrSecretAccessKeyNotFound = errors.NotFoundf("AWS_SECRET_ACCESS_KEY or AWS_SECRET_KEY not found in environment")
)

// EnvCreds returns a static provider of AWS credentials from the process's
// environment, or an error if none are found.
func EnvCreds() (CredentialsProvider, error) {
	id := os.Getenv("AWS_ACCESS_KEY_ID")
	if id == "" {
		id = os.Getenv("AWS_ACCESS_KEY")
	}

	secret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if secret == "" {
		secret = os.Getenv("AWS_SECRET_KEY")
	}

	if id == "" {
		return nil, ErrAccessKeyIDNotFound
	}

	if secret == "" {
		return nil, ErrSecretAccessKeyNotFound
	}

	return Creds(id, secret, os.Getenv("AWS_SESSION_TOKEN")), nil
}

// Creds returns a static provider of credentials.
func Creds(accessKeyID, secretAccessKey, securityToken string) CredentialsProvider {
	return staticCredentialsProvider{
		creds: Credentials{
			AccessKeyID:     accessKeyID,
			SecretAccessKey: secretAccessKey,
			SecurityToken:   securityToken,
		},
	}
}

// IAMCreds returns a provider which pulls credentials from the local EC2
// instance's IAM roles.
func IAMCreds() CredentialsProvider {
	return &iamProvider{}
}

type iamProvider struct {
	creds      Credentials
	m          sync.Mutex
	expiration time.Time
}

var metadataCredentialsEndpoint = "http://169.254.169.254/latest/meta-data/iam/security-credentials/"

func (p *iamProvider) Credentials() (*Credentials, error) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.expiration.After(currentTime()) {
		return &p.creds, nil
	}

	var body struct {
		Expiration      time.Time
		AccessKeyID     string
		SecretAccessKey string
		Token           string
	}

	resp, err := http.Get(metadataCredentialsEndpoint)
	if err != nil {
		return nil, errors.Annotate(err, "listing IAM credentials")
	}
	defer resp.Body.Close()

	// Take the first line of the body of the metadata endpoint
	s := bufio.NewScanner(resp.Body)
	if !s.Scan() {
		return nil, errors.NotFoundf("unable to find default IAM credentials")
	} else if s.Err() != nil {
		return nil, errors.Annotate(s.Err(), "listing IAM credentials")
	}

	resp, err = http.Get(metadataCredentialsEndpoint + s.Text())
	if err != nil {
		return nil, errors.Annotatef(err, "getting %s IAM credentials", s.Text())
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, errors.Annotatef(err, "decoding %s IAM credentials", s.Text())
	}

	p.creds = Credentials{
		AccessKeyID:     body.AccessKeyID,
		SecretAccessKey: body.SecretAccessKey,
		SecurityToken:   body.Token,
	}
	p.expiration = body.Expiration

	return &p.creds, nil
}

type staticCredentialsProvider struct {
	creds Credentials
}

func (p staticCredentialsProvider) Credentials() (*Credentials, error) {
	return &p.creds, nil
}
