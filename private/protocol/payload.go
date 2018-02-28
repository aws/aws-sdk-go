package protocol

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
)

type PayloadUnmarshaler interface {
	UnmarshalPayload(io.Reader, interface{}) error
}

type HandlerPayloadUnmarshal struct {
	Unmarshalers request.HandlerList
}

func (h HandlerPayloadUnmarshal) UnmarshalPayload(r io.Reader, v interface{}) error {
	req := &request.Request{
		HTTPRequest: &http.Request{},
		HTTPResponse: &http.Response{
			StatusCode: 200,
			Header:     http.Header{},
			Body:       ioutil.NopCloser(r),
		},
		Data: v,
	}

	h.Unmarshalers.Run(req)

	return req.Error
}

type PayloadMarshaler interface {
	MarshalPayload(io.Writer, interface{}) error
}

type HandlerPayloadMarshal struct {
	Marshalers request.HandlerList
}

func (h HandlerPayloadMarshal) MarshalPayload(w io.Writer, v interface{}) error {
	req := request.New(
		aws.Config{},
		metadata.ClientInfo{},
		request.Handlers{},
		nil,
		&request.Operation{HTTPMethod: "GET"},
		v,
		nil,
	)

	h.Marshalers.Run(req)

	if req.Error != nil {
		return req.Error
	}

	io.Copy(w, req.GetBody())

	return nil
}
