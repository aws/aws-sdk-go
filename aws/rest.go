package aws

import "net/http"

// RestClient is the underlying client for rest-json and rest-xml APIs.
type RestClient struct {
	Credentials Credentials
	Client      *http.Client
	Service     string
	Region      string
	Endpoint    string
	APIVersion  string
}

// Do sends an HTTP request and returns an HTTP response, following policy
// (e.g. redirects, cookies, auth) as configured on the client.
func (c *RestClient) Do(req *http.Request) (*http.Response, error) {
	sign(c.Service, c.Region, c.Credentials, req)
	return c.Client.Do(req)
}
