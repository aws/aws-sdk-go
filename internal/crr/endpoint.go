package crr

import (
	"time"
)

// Endpoint represents an endpoint used in endpoint discovery.
type Endpoint struct {
	Key     string
	Expired int64
	Address string
}

// HasExpired will return whether or not the endpoint has expired.
func (e Endpoint) HasExpired() bool {
	return time.Now().UnixNano() < e.Expired
}

// Discoverer is an interface used to discovery which endpoint hit. This
// allows for specifics about what parameters need to be used to be contained
// in the Discoverer implementor.
type Discoverer interface {
	Discover() (Endpoint, error)
}
