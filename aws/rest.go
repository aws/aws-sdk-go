package aws

import "net/http"

type RestClient struct {
	Credentials Credentials
	Client      *http.Client
	Service     string
	Region      string
	Endpoint    string
	APIVersion  string
}

func (c *RestClient) Do(req *http.Request) (*http.Response, error) {
	sign(c.Service, c.Region, c.Credentials, req)
	return c.Client.Do(req)
}
