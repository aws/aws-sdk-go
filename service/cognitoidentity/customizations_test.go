package cognitoidentity_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/stretchr/testify/assert"
)

var svc = cognitoidentity.New(&aws.Config{
	Region: "mock-region",
})

func TestUnsignedRequest_GetID(t *testing.T) {
	req, _ := svc.GetIDRequest(&cognitoidentity.GetIDInput{
		IdentityPoolID: aws.String("IdentityPoolId"),
	})

	err := req.Sign()
	assert.NoError(t, err)
	assert.Equal(t, "", req.HTTPRequest.Header.Get("Authorization"))
}

func TestUnsignedRequest_GetOpenIDToken(t *testing.T) {
	req, _ := svc.GetOpenIDTokenRequest(&cognitoidentity.GetOpenIDTokenInput{
		IdentityID: aws.String("IdentityId"),
	})

	err := req.Sign()
	assert.NoError(t, err)
	assert.Equal(t, "", req.HTTPRequest.Header.Get("Authorization"))
}

func TestUnsignedRequest_GetCredentialsForIdentity(t *testing.T) {
	req, _ := svc.GetCredentialsForIdentityRequest(&cognitoidentity.GetCredentialsForIdentityInput{
		IdentityID: aws.String("IdentityId"),
	})

	err := req.Sign()
	assert.NoError(t, err)
	assert.Equal(t, "", req.HTTPRequest.Header.Get("Authorization"))
}
