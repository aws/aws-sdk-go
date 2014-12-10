package aws

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// QueryClient is the underlying client for json APIs.
type QueryClient struct {
	Credentials Credentials
	Client      *http.Client
	Service     string
	Region      string
	Endpoint    string
	APIVersion  string
}

// Do sends an HTTP request and returns an HTTP response, following policy
// (e.g. redirects, cookies, auth) as configured on the client.
func (c *QueryClient) Do(op, method, uri string, req, resp interface{}) error {
	body := url.Values{"Action": {op}, "Version": {c.APIVersion}}
	if err := loadValues(body, req); err != nil {
		return err
	}

	httpReq, err := http.NewRequest(method, c.Endpoint+uri, strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	sign(c.Service, c.Region, c.Credentials, httpReq)

	httpResp, err := c.Client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != 200 {
		var err queryErrorResponse
		if err := xml.NewDecoder(httpResp.Body).Decode(&err); err != nil {
			return err
		}
		return err.Err()
	}

	return xml.NewDecoder(httpResp.Body).Decode(resp)
}

type queryErrorResponse struct {
	XMLName   xml.Name `xml:"ErrorResponse"`
	Type      string   `xml:"Error>Type"`
	Code      string   `xml:"Error>Code"`
	Message   string   `xml:"Error>Message"`
	RequestID string   `xml:"RequestId"`
}

func (e queryErrorResponse) Err() error {
	return APIError{
		Type:      e.Type,
		Code:      e.Code,
		Message:   e.Message,
		RequestID: e.RequestID,
	}
}

func loadValues(v url.Values, i interface{}) error {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	t := value.Type()
	for i := 0; i < value.NumField(); i++ {
		value := value.Field(i)
		name := t.Field(i).Tag.Get("xml")
		switch casted := value.Interface().(type) {
		case string:
			if casted != "" {
				v.Set(name, casted)
			}
		case bool:
			if casted {
				v.Set(name, "true")
			}
		case int64:
			if casted != 0 {
				v.Set(name, fmt.Sprintf("%d", casted))
			}
		case int:
			if casted != 0 {
				v.Set(name, fmt.Sprintf("%d", casted))
			}
		case []string:
			name = strings.Replace(name, ">member", "", -1)
			if len(casted) != 0 {
				for i, val := range casted {
					v.Set(fmt.Sprintf("%s.member.%d", name, i+1), val)
				}
			}
		}
	}
	return nil
}
