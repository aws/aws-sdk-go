// Package v4 implements signing for AWS V4 signer
package v4

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
)

// Sign requests with signature version 4.
//
// Will sign the requests with the service config's Credentials object
// Signing is skipped if the credentials is the credentials.AnonymousCredentials
// object.
func Sign(req *request.Request) {
	v4.SignSDKRequest(req)
}
