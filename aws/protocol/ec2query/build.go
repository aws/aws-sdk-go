package ec2query

import (
	"bytes"
	"io/ioutil"
	"net/url"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/query/queryutil"
)

func Build(r *aws.Request) {
	body := url.Values{
		"Action":  {r.Operation.Name},
		"Version": {r.Service.APIVersion},
	}
	if err := queryutil.Parse(body, r.Params, true); err != nil {
		r.Error = err
		return
	}

	r.HTTPRequest.Method = "POST"
	r.HTTPRequest.Body = ioutil.NopCloser(bytes.NewReader([]byte(body.Encode())))
}
