package aws

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// EC2Client is the underlying client for EC2 APIs.
type EC2Client struct {
	Context    Context
	Client     *http.Client
	Endpoint   string
	APIVersion string
}

// Do sends an HTTP request and returns an HTTP response, following policy
// (e.g. redirects, cookies, auth) as configured on the client.
func (c *EC2Client) Do(op, method, uri string, req, resp interface{}) error {
	body := url.Values{"Action": {op}, "Version": {c.APIVersion}}
	if err := c.loadValues(body, req, ""); err != nil {
		return err
	}

	httpReq, err := http.NewRequest(method, c.Endpoint+uri, strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	httpReq.Header.Set("User-Agent", "aws-go")
	if err := c.Context.sign(httpReq); err != nil {
		return err
	}

	httpResp, err := c.Client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		var err ec2ErrorResponse
		if err := xml.NewDecoder(httpResp.Body).Decode(&err); err != nil {
			return err
		}
		return err.Err()
	}

	if resp != nil {
		return xml.NewDecoder(httpResp.Body).Decode(resp)
	}
	return nil
}

type ec2ErrorResponse struct {
	XMLName   xml.Name `xml:"Response"`
	Type      string   `xml:"Errors>Error>Type"`
	Code      string   `xml:"Errors>Error>Code"`
	Message   string   `xml:"Errors>Error>Message"`
	RequestID string   `xml:"RequestID"`
}

func (e ec2ErrorResponse) Err() error {
	return APIError{
		Type:      e.Type,
		Code:      e.Code,
		Message:   e.Message,
		RequestID: e.RequestID,
	}
}

func (c *EC2Client) loadValues(v url.Values, i interface{}, prefix string) error {
	value := reflect.ValueOf(i)

	// follow any pointers
	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return c.loadStruct(v, value, prefix)
}

func (c *EC2Client) loadStruct(v url.Values, value reflect.Value, prefix string) error {
	if !value.IsValid() {
		return nil
	}

	t := value.Type()
	for i := 0; i < value.NumField(); i++ {
		value := value.Field(i)
		name := t.Field(i).Tag.Get("ec2")

		if name == "" {
			name = t.Field(i).Name
		}
		if prefix != "" {
			name = prefix + "." + name
		}
		switch casted := value.Interface().(type) {
		case StringValue:
			if casted != nil {
				v.Set(name, *casted)
			}
		case BooleanValue:
			if casted != nil {
				v.Set(name, strconv.FormatBool(*casted))
			}
		case LongValue:
			if casted != nil {
				v.Set(name, strconv.FormatInt(*casted, 10))
			}
		case IntegerValue:
			if casted != nil {
				v.Set(name, strconv.Itoa(*casted))
			}
		case DoubleValue:
			if casted != nil {
				v.Set(name, strconv.FormatFloat(*casted, 'f', -1, 64))
			}
		case FloatValue:
			if casted != nil {
				v.Set(name, strconv.FormatFloat(float64(*casted), 'f', -1, 32))
			}
		case []string:
			if len(casted) != 0 {
				for i, val := range casted {
					v.Set(fmt.Sprintf("%s.%d", name, i+1), val)
				}
			}
		default:
			if err := c.loadValues(v, value.Interface(), name); err != nil {
				return err
			}
		}
	}
	return nil
}
