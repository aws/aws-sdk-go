package aws

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/crowdmob/goamz/aws"
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

	if httpResp.StatusCode != 200 {
		b, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(b))
	}

	return json.NewDecoder(httpResp.Body).Decode(resp)
}
