package aws

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/dynport/gocloud/aws"
)

type JSONClient struct {
	Client       *http.Client
	Region       string
	Endpoint     string
	Prefix       string
	Key          string
	Secret       string
	APIVersion   string
	JSONVersion  string
	TargetPrefix string
}

func (c *JSONClient) Do(op, method, uri string, req, resp interface{}) error {
	u, err := url.Parse(c.Endpoint + uri)
	if err != nil {
		return err
	}

	b, err := json.Marshal(req)
	if err != nil {
		return err
	}

	r := aws.RequestV4{
		Key:     c.Key,
		Secret:  c.Secret,
		Method:  method,
		URL:     u,
		Payload: b,
		Region:  c.Region,
		Service: c.Prefix,
		Time:    time.Now(),
	}
	r.SetHeader("X-Amz-Target", c.TargetPrefix+"."+op)
	r.SetHeader("Content-Type", "application/x-amz-json-"+c.JSONVersion)

	httpReq, err := r.Request()
	if err != nil {
		return err
	}

	httpResp, err := c.Client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != 200 {
		b, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(b))
	}

	return json.NewDecoder(httpResp.Body).Decode(resp)
}
