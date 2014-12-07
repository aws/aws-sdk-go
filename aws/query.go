package aws

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/crowdmob/goamz/aws"
)

type QueryClient struct {
	Client     *http.Client
	Region     string
	Endpoint   string
	Prefix     string
	Key        string
	Secret     string
	APIVersion string
}

func (c *QueryClient) Do(op, method, uri string, req, resp interface{}) error {
	v := url.Values{"Action": {op}, "Version": {c.APIVersion}}
	if err := loadValues(v, req); err != nil {
		return err
	}

	u, err := url.Parse(c.Endpoint + uri + "?" + v.Encode())
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return err
	}

	signer := aws.NewV4Signer(aws.Auth{
		AccessKey: c.Key,
		SecretKey: c.Secret,
	}, c.Prefix, aws.Region{Name: c.Region})
	signer.Sign(httpReq)

	httpResp, err := c.Client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	b, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	if httpResp.StatusCode != 200 {
		return errors.New(string(b))
	}

	return xml.Unmarshal(b, resp)
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
			if len(casted) != 0 {
				for i, val := range casted {
					v.Set(fmt.Sprintf("%s.member.%d", name, i+1), val)
				}
			}
		}
	}
	return nil
}
