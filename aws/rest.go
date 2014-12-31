package aws

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// RestClient is the underlying client for REST-JSON and REST-XML APIs.
type RestClient struct {
	Context    Context
	Client     *http.Client
	Endpoint   string
	APIVersion string
}

// Do sends an HTTP request and returns an HTTP response, following policy
// (e.g. redirects, cookies, auth) as configured on the client.
func (c *RestClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", "aws-go")
	if err := c.Context.sign(req); err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return nil, err
		}
		if len(bodyBytes) == 0 {
			return nil, APIError{
				StatusCode: resp.StatusCode,
				Message:    resp.Status,
			}
		}
		var restErr restError
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			if err := json.Unmarshal(bodyBytes, &restErr); err != nil {
				return nil, err
			}
			return nil, restErr.Err(resp.StatusCode)
		case "application/xml", "text/xml":
			// AWS XML error documents can have a couple of different formats.
			// Try each before returning a decode error.
			var wrappedErr restErrorResponse
			if err := xml.Unmarshal(bodyBytes, &wrappedErr); err == nil {
				return nil, wrappedErr.Error.Err(resp.StatusCode)
			}
			if err := xml.Unmarshal(bodyBytes, &restErr); err != nil {
				return nil, err
			}
			return nil, restErr.Err(resp.StatusCode)
		default:
			return nil, APIError{
				StatusCode: resp.StatusCode,
				Message:    string(bodyBytes),
			}
		}
	}

	return resp, nil
}

type restErrorResponse struct {
	XMLName xml.Name `xml:"ErrorResponse",json:"-"`
	Error   restError
}

type restError struct {
	XMLName    xml.Name `xml:"Error",json:"-"`
	Code       string
	BucketName string
	Message    string
	RequestID  string
	HostID     string
}

func (e restError) Err(StatusCode int) error {
	return APIError{
		StatusCode: StatusCode,
		Code:       e.Code,
		Message:    e.Message,
		RequestID:  e.RequestID,
		HostID:     e.HostID,
		Specifics: map[string]string{
			"BucketName": e.BucketName,
		},
	}
}
