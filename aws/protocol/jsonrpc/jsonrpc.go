package jsonrpc

import (
	"encoding/json"
	"io/ioutil"

	"github.com/awslabs/aws-sdk-go/aws"
)

var emptyJSON = []byte("{}")

func Build(req *aws.Request) {
	var buf []byte
	var err error
	if req.ParamsFilled() {
		buf, err = json.Marshal(req.Params)
		if err != nil {
			req.Error = err
			return
		}
	} else {
		buf = emptyJSON
	}
	req.SetBufferBody(buf)

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
	defer req.HTTPResponse.Body.Close()
	if req.Data != nil {
		json.NewDecoder(req.HTTPResponse.Body).Decode(req.Data)
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
	req.Error = jsonErr.Err(req.HTTPResponse.StatusCode)
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
