package aws

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type JSONClient struct {
	Credentials  Credentials
	Client       *http.Client
	Service      string
	Region       string
	Endpoint     string
	TargetPrefix string
	JSONVersion  string
}

func (c *JSONClient) Do(op, method, uri string, req, resp interface{}) error {
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest(method, c.Endpoint+uri, bytes.NewReader(b))
	if err != nil {
		return err
	}
	httpReq.Header.Set("X-Amz-Target", c.TargetPrefix+"."+op)
	httpReq.Header.Set("Content-Type", "application/x-amz-json-"+c.JSONVersion)
	sign(c.Service, c.Region, c.Credentials, httpReq)

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
