package jsonrpc

//go:generate go run ../../fixtures/protocol/generate.go ../../fixtures/protocol/input/json.json build_test.go
//go:generate go run ../../fixtures/protocol/generate.go ../../fixtures/protocol/output/json.json unmarshal_test.go

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/json/jsonutil"
)

var emptyJSON = []byte("{}")

func Build(req *aws.Request) {
	var buf []byte
	var err error
	if req.ParamsFilled() {
		buf, err = jsonutil.BuildJSON(req.Params)
		if err != nil {
			req.Error = err
			return
		}
	} else {
		buf = emptyJSON
	}

	if req.Service.TargetPrefix != "" || string(buf) != "{}" {
		req.SetBufferBody(buf)
	}

	if req.Service.TargetPrefix != "" {
		target := req.Service.TargetPrefix + "." + req.Operation.Name
		req.HTTPRequest.Header.Add("X-Amz-Target", target)
	}
	if req.Service.JSONVersion != "" {
		jsonVersion := req.Service.JSONVersion
		req.HTTPRequest.Header.Add("Content-Type", "application/x-amz-json-"+jsonVersion)
	}
}

func Unmarshal(req *aws.Request) {
	if req.DataFilled() {
		err := jsonutil.UnmarshalJSON(req.Data, req.HTTPResponse.Body)
		if err != nil {
			req.Error = err
		}
	}
	return
}

func UnmarshalMeta(req *aws.Request) {
	req.RequestID = req.HTTPResponse.Header.Get("x-amzn-requestid")
}

func UnmarshalError(req *aws.Request) {
	bodyBytes, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil {
		req.Error = err
		return
	}
	if len(bodyBytes) == 0 {
		req.Error = aws.APIError{
			StatusCode: req.HTTPResponse.StatusCode,
			Message:    req.HTTPResponse.Status,
		}
		return
	}
	var jsonErr jsonErrorResponse
	if err := json.Unmarshal(bodyBytes, &jsonErr); err != nil {
		req.Error = err
		return
	}

	codes := strings.SplitN(jsonErr.Code, "#", 2)
	req.Error = aws.APIError{
		StatusCode: req.HTTPResponse.StatusCode,
		Code:       codes[len(codes)-1],
		Message:    jsonErr.Message,
	}
}

type jsonErrorResponse struct {
	Code    string `json:"__type"`
	Message string `json:"message"`
}
