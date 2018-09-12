package crr

import (
	"sort"
	"strings"
	"time"
)

// Endpoint represents an endpoint used in endpoint discovery.
type Endpoint struct {
	Key       string
	Expired   int64
	Addresses WeightedAddresses
}

// WeightedAddresses represents a list of WeightedAddress.
type WeightedAddresses []WeightedAddress

// GetAddress will return an address from the list of addresses. If
// there are no addresses in the list, then false will be returned.
func (w WeightedAddresses) GetAddress() (string, bool) {
	if len(w) == 0 {
		return "", false
	}

	return w[0].Address, true
}

// WeightedAddress represents an address with a given weight.
// Currently the weight of addresses is set to 1.0.
type WeightedAddress struct {
	Address string
}

// Add will add a given WeightedAddress to the address list of Endpoint.
func (e *Endpoint) Add(addr WeightedAddress) {
	e.Addresses = append(e.Addresses, addr)
}

// HasExpired will return whether or not the endpoint has expired.
func (e Endpoint) HasExpired() bool {
	return e.Expired < time.Now().UnixNano()
}

// Discoverer is an interface used to discovery which endpoint hit. This
// allows for specifics about what parameters need to be used to be contained
// in the Discoverer implementor.
type Discoverer interface {
	Discover() (Endpoint, error)
}

// BuildEndpointKey will sort the keys in alphabetical order and then retrieve
// the values in that order. Those values are then concatenated together to form
// the endpoint key.
func BuildEndpointKey(params map[string]string) string {
	keys := make([]string, len(params))
	i := 0

	for k := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	values := make([]string, len(params))
	for i, k := range keys {
		values[i] = params[k]
	}

	return strings.Join(values, ".")
}
