package route53

import (
	"regexp"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol/restxml"
)

func init() {
	initClient = func(c *client.Client) {
		c.Handlers.Build.PushBack(sanitizeURL)
	}

	initRequest = func(r *request.Request) {
		switch r.Operation.Name {
		case opChangeResourceRecordSets:
			r.Handlers.UnmarshalError.Remove(restxml.UnmarshalErrorHandler)
			r.Handlers.UnmarshalError.PushBack(unmarshalChangeResourceRecordSetsError)
		}
	}
}

var reSanitizeURL = regexp.MustCompile(`\/%2F\w+%2F`)

func sanitizeURL(r *request.Request) {
	r.HTTPRequest.URL.RawPath = reSanitizeURL.ReplaceAllString(
		r.HTTPRequest.URL.RawPath, "/")

	// Save off the raw path so that it can be reset to the new URL. Parse
	// will not set RawPath.
	rawPath := r.HTTPRequest.URL.RawPath

	// Update Path so that it reflects the cleaned RawPath
	u, err := r.HTTPRequest.URL.Parse(rawPath)
	if err != nil {
		r.Error = awserr.New("SerializationError", "failed to clean Route53 URL", err)
		return
	}

	r.HTTPRequest.URL = u
	r.HTTPRequest.URL.RawPath = rawPath
}
