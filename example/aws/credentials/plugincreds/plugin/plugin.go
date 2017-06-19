// +build example,go18

package main

// Example plugin
//
// Build with:
//   go build -tags example -o plugin.so -buildmode=plugin plugin.go
func main() {}

var myCredProvider provider

func init() {
	// Initialize a mock credential provider with stubs
	myCredProvider = provider{"a", "b", "c"}
}

// GetAWSSDKCredentialProvider is the symbol SDK will lookup and use to
// get the credential provider's retrieve and isExpired functions.
func GetAWSSDKCredentialProvider() (func() (key, secret, token string, err error), func() bool) {
	return myCredProvider.Retrieve, myCredProvider.IsExpired
}

// mock implementation of a type that returns retrieves credentials and
// returns if they have expired.
type provider struct {
	key, secret, token string
}

func (p provider) Retrieve() (key, secret, token string, err error) {
	return p.key, p.secret, p.token, nil
}

func (p *provider) IsExpired() bool {
	return false
}
