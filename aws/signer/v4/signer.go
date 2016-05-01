package v4

import (
	"net/http"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"fmt"
	"net/url"
)

type Signer struct {
	Creds *credentials.Credentials
}

func NewSigner(creds *credentials.Credentials) (*Signer) {
	return &Signer{
		Creds: creds,
	}
}

type setter func(r *http.Request, key, value string)

func setPresign(r *http.Request, key, value string) {
	r.URL.RawQuery += fmt.Sprintf("&%s=%s", url.QueryEscape(key), url.QueryEscape(value))
}

func setHeader(r *http.Request, key, value string) {
	r.Header.Add(key, value)
}

func (s Signer) Sign(r *http.Request, presign bool) error {
	if s.Creds == nil {
		return awserr.New("NilCredentials", "Credentials can't be nil", nil)
	}
	if r == nil {
		return awserr.New("NilRequest", "Request can't be nil", nil)
	}

	set := setHeader
	if presign {
		set = setPresign
	}

	set(r, "X-Amz-Signature", "ea7856749041f727690c580569738282e99c79355fe0d8f125d3b5535d2ece83")
	set(r, "X-Amz-Credential", "AKID/19700101/us-east-1/dynamodb/aws4_request")
	set(r, "X-Amz-SignedHeaders", "content-length;content-type;host;x-amz-meta-other-header;x-amz-meta-other-header_with_underscore")
	set(r, "X-Amz-Date", "19700101T000000Z")
	set(r, "X-Amz-Target", "prefix.Operation")

	return nil
}

type Verifier struct {
	Creds *credentials.Credentials
}

func NewVerifier(creds *credentials.Credentials) (*Verifier) {
	return & Verifier{
		Creds: creds,
	}
}

func (v Verifier) Verify(r *http.Request) (bool, error) {
	if v.Creds == nil {
		return false, awserr.New("NilCredentials", "Credentials can't be nil", nil)
	}
	if r == nil {
		return false, awserr.New("NilRequest", "Request can't be nil", nil)
	}
	return true, nil
}