package credentials

import (
	"fmt"
)

// A ChainProvider will search for a provider which returns credentials
// and cache that provider until Retrieve is called again.
type ChainProvider struct {
	Providers []Provider
	curr      Provider
}

// NewMetadataServiceCredentials returns a pointer to a new Credentials object
// wrapping a chain of providers.
func NewChainCredentials(providers []Provider) *Credentials {
	return NewCredentials(&ChainProvider{
		Providers: append([]Provider{}, providers...),
	})
}

// Retrieve returns the credentials value or error if no provider returned
// without error.
//
// If a provider is found it will be cached and any calls to IsExpired
// will return the expired state of the cached provider.
func (c *ChainProvider) Retrieve() (Value, error) {
	for _, p := range c.Providers {
		if creds, err := p.Retrieve(); err == nil {
			c.curr = p
			return creds, nil
		}
	}
	c.curr = nil

	// TODO better error reporting. maybe report error for each failed retrieve?

	return Value{}, fmt.Errorf("no valid providers in chain")
}

// IsExpired will returned the expired state of the currently cached provider
// if there is one.  If there is no current provider, true will be returned.
func (c *ChainProvider) IsExpired() bool {
	if c.curr != nil {
		return c.curr.IsExpired()
	}

	return true
}
