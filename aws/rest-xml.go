package aws

import "net/http"

type RestXMLClient struct {
	Signer     *V4Signer
	Client     *http.Client
	Endpoint   string
	APIVersion string
}

func (c *RestXMLClient) Do(req *http.Request) (*http.Response, error) {
	c.Signer.sign(req)
	return c.Client.Do(req)
}
