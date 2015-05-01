// Package provides Credential retrieval and management
//
//
package credentials

import (
	"sync"
	"time"
)

// A Value is the AWS credentials value for individual credential fields.
type Value struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

// A Provider is the interface for any component which will provide credential value.
type Provider interface {
	// Refresh returns nil if it successfully retrieved the value.
	// Error is returned if the value were not obtainable.
	Retrieve() (Value, error)

	// IsExpired returns if the credentials are no longer valid, and need
	// to be retrieved.
	IsExpired() bool
}

// A Credentials provides synchronous safe retrieval of AWS credential values.
// Credentials will cache the credentials value until they expire. Once the value
// expires the next Get will attempt to retrieve valid credentials.
type Credentials struct {
	creds        Value
	forceRefresh bool
	m            sync.Mutex

	provider Provider
}

// NewCredentials returns a pointer to a new Credentials object with the
// provider set.
func NewCredentials(provider Provider) *Credentials {
	return &Credentials{provider: provider}
}

// Get returns the credentials value, or error if the credential failed
// to be retrieved.
func (c *Credentials) Get() (Value, error) {
	c.m.Lock()
	defer c.m.Unlock()

	if c.isExpired() {
		creds, err := c.provider.Retrieve()
		if err != nil {
			return Value{}, err
		}
		c.creds = creds
		c.forceRefresh = false
	}

	return c.creds, nil
}

// Expire expires the credentials and forces them to be retrieved on the
// next call to Get.
func (c *Credentials) Expire() {
	c.m.Lock()
	defer c.m.Unlock()

	c.forceRefresh = true
}

// IsExpired returns if the credentials are no longer valid, and need
// to be retrieved.
func (c *Credentials) IsExpired() bool {
	c.m.Lock()
	defer c.m.Unlock()

	return c.isExpired()
}

// isExpired helper method wrapping the definition of expired credentials.
func (c *Credentials) isExpired() bool {
	return c.forceRefresh || c.provider.IsExpired()
}

// Provide a stub-able time.Now for unit tests so expiry can be tested.
var currentTime = time.Now
