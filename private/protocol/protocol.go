package protocol

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
)

type RequireHTTPMinProtocol struct {
	Major, Minor int
}

func (p RequireHTTPMinProtocol) Handler(r *request.Request) {
	if r.Error != nil || r.HTTPResponse == nil {
		return
	}

	if !strings.HasPrefix(r.HTTPResponse.Proto, "HTTP") {
		r.Error = newMinHTTPProtoError(p.Major, p.Minor, r)
	}

	if r.HTTPResponse.ProtoMajor < p.Major || r.HTTPResponse.ProtoMinor < p.Minor {
		r.Error = newMinHTTPProtoError(p.Major, p.Minor, r)
	}
}

const ErrCodeMinimumHTTPProtocolError = "MinimumHTTPProtocolError"

func newMinHTTPProtoError(major, minor int, r *request.Request) error {
	return awserr.NewRequestFailure(
		awserr.New("MinimumHTTPProtocolError",
			fmt.Sprintf(
				"operation requires minimum HTTP protocol of HTTP/%d.%d, but was %s",
				major, minor, r.HTTPResponse.Proto,
			),
			nil,
		),
		r.HTTPResponse.StatusCode, r.RequestID,
	)
}
