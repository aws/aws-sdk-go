package restjson

//go:generate go run ../../fixtures/protocol/generate.go ../../fixtures/protocol/input/rest-json.json build_test.go
//go:generate go run ../../fixtures/protocol/generate.go ../../fixtures/protocol/output/rest-json.json unmarshal_test.go

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/internal/protocol/rest"
)

func Build(r *aws.Request) {
	rest.Build(r)

	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
}

func Unmarshal(r *aws.Request) {
	if t := rest.PayloadType(r.Data); t == "structure" || t == "" {
		jsonrpc.Unmarshal(r)
	}
}

func UnmarshalMeta(r *aws.Request) {
	rest.Unmarshal(r)
}

func UnmarshalError(r *aws.Request) {
	code := r.HTTPResponse.Header.Get("X-Amzn-Errortype")
	bodyBytes, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		r.Error = err
		return
	}
	if len(bodyBytes) == 0 {
		r.Error = aws.APIError{
			StatusCode: r.HTTPResponse.StatusCode,
			Message:    r.HTTPResponse.Status,
		}
		return
	}
	var jsonErr jsonErrorResponse
	if err := json.Unmarshal(bodyBytes, &jsonErr); err != nil {
		r.Error = err
		return
	}

	codes := strings.SplitN(code, ":", 2)
	r.Error = aws.APIError{
		StatusCode: r.HTTPResponse.StatusCode,
		Code:       codes[0],
		Message:    jsonErr.Message,
	}
}

type jsonErrorResponse struct {
	Message string `json:"message"`
}
