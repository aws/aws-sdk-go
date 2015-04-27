package credentials

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const metadataCredentialsEndpoint = "http://169.254.169.254/latest/meta-data/iam/security-credentials/"

// A EC2RoleProvider retrieves credentials from the EC2 service, and keeps track if
// those credentials are expired.
type EC2RoleProvider struct {
	// Endpoint must be fully quantified URL
	Endpoint string

	// HTTP client to use when connecting to EC2 service
	Client *http.Client

	// ExpiryWindow will allow the credentials to trigger refreshing prior to
	// the credentials actually expiring. This is beneficial so race conditions
	// with expiring credentials do not cause request to fail unexpectedly
	// due to ExpiredTokenException exceptions.
	//
	// So a ExpiryWindow of 10s would cause calls to IsExpired() to return true
	// 10 seconds before the credentials are actually expired.
	ExpiryWindow time.Duration

	// The date/time at which the credentials expire.
	expiresOn time.Time
}

// NewMetadataServiceCredentials returns a pointer to a new Credentials object
// wrapping the MetadataService provider.
func NewMetadataServiceCredentials(client *http.Client, endpoint string, window time.Duration) *Credentials {
	return NewCredentials(&EC2RoleProvider{
		Endpoint:     endpoint,
		Client:       client,
		ExpiryWindow: window,
	})
}

// Retrieve retrieves credentials from the EC2 service.
// Error will be returned if the request fails, or unable to extract
// the desired credentials.
func (m *EC2RoleProvider) Retrieve() (Value, error) {
	if m.Client == nil {
		m.Client = http.DefaultClient
	}
	if m.Endpoint == "" {
		m.Endpoint = metadataCredentialsEndpoint
	}

	credsList, err := requestCredList(m.Client, m.Endpoint)
	if err != nil {
		return Value{}, err
	}

	if len(credsList) == 0 {
		return Value{}, fmt.Errorf("empty MetadataService credentials list")
	}
	credsName := credsList[0]

	roleCreds, err := requestCred(m.Client, m.Endpoint, credsName)
	if err != nil {
		return Value{}, err
	}

	m.expiresOn = roleCreds.Expiration
	return Value{
		AccessKeyID:     roleCreds.AccessKeyID,
		SecretAccessKey: roleCreds.SecretAccessKey,
		SessionToken:    roleCreds.Token,
	}, nil
}

// IsExpired returns if the credentials are expired.
func (m *EC2RoleProvider) IsExpired() bool {
	return m.expiresOn.Before(currentTime())
}

// A ec2RoleCredRespBody provides the shape for deserializing credential
// request responses.
type ec2RoleCredRespBody struct {
	Expiration      time.Time
	AccessKeyID     string
	SecretAccessKey string
	Token           string
}

// requestCredList requests a list of credentials from the EC2 service.
// If there are no credentials, or there is an error making or receiving the request
func requestCredList(client *http.Client, endpoint string) ([]string, error) {
	resp, err := client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("%s listing MetadataService credentials", err)
	}
	defer resp.Body.Close()

	credsList := []string{}
	s := bufio.NewScanner(resp.Body)
	for s.Scan() {
		credsList = append(credsList, s.Text())
	}

	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("%s reading list of MetadataService credentials", err)
	}

	return credsList, nil
}

// requestCred requests the credentials for a specific credentials from the EC2 service.
//
// If the credentials cannot be found, or there is an error reading the response
// and error will be returned.
func requestCred(client *http.Client, endpoint, credsName string) (*ec2RoleCredRespBody, error) {
	resp, err := client.Get(endpoint + credsName)
	if err != nil {
		return nil, fmt.Errorf("getting %s MetadataService credentials", credsName)
	}
	defer resp.Body.Close()

	respCreds := &ec2RoleCredRespBody{}
	if err := json.NewDecoder(resp.Body).Decode(respCreds); err != nil {
		return nil, fmt.Errorf("decoding %s MetadataService credentials", credsName)
	}

	return respCreds, nil
}
