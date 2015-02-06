package query

import (
	"encoding/xml"
	"io"

	"github.com/awslabs/aws-sdk-go/aws"
)

type queryErrorResponse struct {
	XMLName   xml.Name `xml:"ErrorResponse"`
	Code      string   `xml:"Error>Code"`
	Message   string   `xml:"Error>Message"`
	RequestID string   `xml:"RequestId"`
}

func UnmarshalError(r *aws.Request) {
	defer r.HTTPResponse.Body.Close()

	resp := &queryErrorResponse{}
	err := xml.NewDecoder(r.HTTPResponse.Body).Decode(resp)
	if err != nil && err != io.EOF {
		r.Error = err
	} else {
		apiErr := r.Error.(aws.APIError)
		apiErr.Code = resp.Code
		apiErr.Message = resp.Message
		r.Error = apiErr
	}
}
