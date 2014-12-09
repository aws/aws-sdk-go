package aws

import "net/http"

type RestClient struct {
	Signer     *V4Signer
	Client     *http.Client
	Endpoint   string
	APIVersion string
}

func (c *RestClient) Do(req *http.Request) (*http.Response, error) {
	c.Signer.sign(req)
	return c.Client.Do(req)
}
