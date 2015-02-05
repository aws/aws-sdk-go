package v4

import (
	"bytes"
	"crypto/sha256"
	"io"
	"io/ioutil"
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

type readSeekCloser struct {
	*bytes.Reader
}

func (r readSeekCloser) Close() error {
	return nil
}

func TestMakeShaReaderReadSeeker(t *testing.T) {
	var seeker io.ReadSeeker
	testData := []byte(`I am a read seeker`)
	seeker = &readSeekCloser{bytes.NewReader(testData)}
	mockRequest, _ := http.NewRequest("GET", "http://example.com", seeker)

	checksum := makeSha256Reader(mockRequest)
	expectedCheckum := sha256.Sum256(testData)

	if !bytes.Equal(checksum, expectedCheckum[:]) {
		t.Errorf("Checksum did not match expected; was %v instead of %v", checksum, sha256.Sum256(testData))
	}

	// Verify it was rewound
	data, err := ioutil.ReadAll(mockRequest.Body)
	if err != nil {
		t.Errorf("Unable to read seeker: %v", err)
	}
	if !bytes.Equal(data, testData) {
		t.Errorf("Re-reading reader got different result: %v != %v", string(data), string(testData))
	}
}

func TestMakeShaReaderNotSeeker(t *testing.T) {
	var reader io.Reader
	testData := []byte(`I am not a read seeker`)
	reader = bytes.NewBuffer(testData)
	if _, ok := reader.(io.ReadSeeker); ok {
		t.Fatal("We're trying to test something that's not a readseeker here. Fix the test or downgrade Go :)")
	}
	mockRequest, _ := http.NewRequest("GET", "http://example.com", reader)

	checksum := makeSha256Reader(mockRequest)
	expectedCheckum := sha256.Sum256(testData)

	if !bytes.Equal(checksum, expectedCheckum[:]) {
		t.Errorf("Checksum did not match expected; was %v instead of %v", checksum, sha256.Sum256(testData))
	}

	// Verify it was rewound
	data, err := ioutil.ReadAll(mockRequest.Body)
	if err != nil {
		t.Errorf("Unable to read seeker: %v", err)
	}
	if !bytes.Equal(data, testData) {
		t.Errorf("Re-reading reader got different result: %v != %v", string(data), string(testData))
	}
}
