package v4

import (
	"net/http"
	"strings"
	"testing"
	"time"
)

func buildSigner(serviceName string, region string, signTime time.Time, body string) signer {
	endpoint := "https://" + serviceName + "." + region + ".amazonaws.com"
	reader := strings.NewReader(body)
	req, _ := http.NewRequest("POST", endpoint, reader)
	req.URL.Opaque = "//example.org/bucket/key-._~,!@#$%^&*()"
	req.Header.Add("X-Amz-Target", "prefix.Operation")
	req.Header.Add("Content-Type", "application/x-amz-json-1.0")
	req.Header.Add("Content-Length", string(len(body)))
	req.Header.Add("X-Amz-Meta-Other-Header", "some-value=!@#$%^&* ()")

	return signer{
		Request:         req,
		Time:            signTime,
		ExpireTime:      300 * time.Second,
		Query:           req.URL.Query(),
		Body:            reader,
		ServiceName:     serviceName,
		Region:          region,
		AccessKeyID:     "AKID",
		SecretAccessKey: "SECRET",
		SessionToken:    "SESSION",
	}
}

func removeWS(text string) string {
	text = strings.Replace(text, " ", "", -1)
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\t", "", -1)
	return text
}

func assertEqual(t *testing.T, expected, given string) {
	if removeWS(expected) != removeWS(given) {
		t.Errorf("\nExpected: %s\nGiven:    %s", expected, given)
	}
}

func TestSignRequest(t *testing.T) {
	signer := buildSigner("dynamodb", "us-east-1", time.Unix(0, 0), "{}")
	signer.sign()

	expectedDate := "19700101T000000Z"
	expectedHeaders := "host;x-amz-meta-other-header;x-amz-target"
	expectedSig := "41c18d68f9191079dfeead4e3f034328f89d86c79f8e9d51dd48bb70eaf623fc"
	expectedCred := "AKID/19700101/us-east-1/dynamodb/aws4_request"

	q := signer.Request.URL.Query()
	assertEqual(t, expectedSig, q.Get("X-Amz-Signature"))
	assertEqual(t, expectedCred, q.Get("X-Amz-Credential"))
	assertEqual(t, expectedHeaders, q.Get("X-Amz-SignedHeaders"))
	assertEqual(t, expectedDate, q.Get("X-Amz-Date"))
}

func BenchmarkSignRequest(b *testing.B) {
	signer := buildSigner("dynamodb", "us-east-1", time.Now(), "{}")
	for i := 0; i < b.N; i++ {
		signer.sign()
	}
}
