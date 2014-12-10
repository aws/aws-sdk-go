package aws

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

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
	if err := sign(c.Service, c.Region, c.Credentials, req); err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		var err restErrorResponse
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			if err := json.NewDecoder(resp.Body).Decode(&err); err != nil {
				return nil, err
			}
			return nil, err.Err()
		case "application/xml":
			if err := xml.NewDecoder(resp.Body).Decode(&err); err != nil {
				return nil, err
			}
			return nil, err.Err()
		default:
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			return nil, errors.New(string(b))
		}
	}

	return resp, nil
}

type restErrorResponse struct {
	XMLName    xml.Name `xml:"Error",json:"-"`
	Code       string
	BucketName string
	Message    string
	RequestID  string
	HostID     string
}

func (e restErrorResponse) Err() error {
	return APIError{
		Code:      e.Code,
		Message:   e.Message,
		RequestID: e.RequestID,
		HostID:    e.HostID,
		Specifics: map[string]string{
			"BucketName": e.BucketName,
		},
	}
}
