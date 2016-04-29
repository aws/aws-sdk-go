package v4

import (
	"net/http"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

type Signer struct {
	Creds *credentials.Credentials
}

func NewSigner(creds *credentials.Credentials) (*Signer) {
	return &Signer{
		Creds: creds,
	}
}

func (s Signer) Sign(r *http.Request) error {
	if s.Creds == nil {
		return awserr.New("NilCredentials", "Credentials can't be nil", nil)
	}
	if r == nil {
		return awserr.New("NilRequest", "Request can't be nil", nil)
	}

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