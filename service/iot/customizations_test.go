package iot_test

import (
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/service/iot"
)

func TestBackfillIotSigningName(t *testing.T) {
	svc := iot.New(unit.Session)
	if svc.SigningName != iot.ServiceName {
		t.Errorf("Expected backfilled signing name `" + iot.ServiceName + "`, but found: `" + svc.SigningName + "`")
	}
}

func TestBackfillIotSigningNameAuthorizationHeader(t *testing.T) {
	svc := iot.New(unit.Session)

	// Arbitrarily use ListAuthorizers since it doesn't require any input
	req, _ := svc.ListAuthorizersRequest(&iot.ListAuthorizersInput{})
	req.Sign()
	authorizationHeader := req.HTTPRequest.Header.Get("Authorization")
	if authorizationHeader == "" {
		t.Errorf("Expected `Authorization` header to be present")
	}

	r := regexp.MustCompile(
		`^AWS4-HMAC-SHA256 ` +
			`Credential=AKID/[[:digit:]]{8}/mock-region/iot/aws4_request, ` +
			`SignedHeaders=host;x-amz-date;x-amz-security-token, ` +
			`Signature=[[:xdigit:]]{64}$`)
	if !r.MatchString(authorizationHeader) {
		t.Errorf("Expect regex %v to match, but found: %v", r.String(), authorizationHeader)
	}
}
