package jsonrpc

import (
	"encoding/json"
	"io/ioutil"

	"github.com/stripe/aws-go/aws"
)

var emptyJSON = []byte("{}")

func Build(req *aws.Request) {
	var buf []byte
	var err error
	if req.Params != nil {
		buf, err = json.Marshal(req.Params)
		if err != nil {
			req.Error = err
			return
		}

		if string(buf) == "null" {
			buf = emptyJSON
		}
	} else {
		buf = emptyJSON
	}
	req.SetBufferBody(buf)

	target := req.Service.TargetPrefix + "." + req.Operation.Name
	jsonVersion := req.Service.JSONVersion
	req.HTTPRequest.Header.Add("X-Amz-Target", target)
	req.HTTPRequest.Header.Add("Content-Type", "application/x-amz-json-"+jsonVersion)
}

func Unmarshal(req *aws.Request) {
	defer req.HTTPResponse.Body.Close()

	req.RequestID = req.HTTPResponse.Header.Get("x-amzn-requestid")

	if req.HTTPResponse.StatusCode >= 300 {
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
		req.Error = jsonErr.Err(req.HTTPResponse.StatusCode)
	}

	if req.Data != nil {
		json.NewDecoder(req.HTTPResponse.Body).Decode(req.Data)
	}
	return
}

type jsonErrorResponse struct {
	Type    string `json:"__type"`
	Message string `json:"message"`
}

func (e jsonErrorResponse) Err(StatusCode int) error {
	return aws.APIError{
		StatusCode: StatusCode,
		Type:       e.Type,
		Message:    e.Message,
	}
}
