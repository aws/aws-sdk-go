package v4_test

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
)

var standaloneSignCases = []struct {
	OrigURI                    string
	OrigQuery                  string
	Region, Service, SubDomain string
	ExpSig                     string
	EscapedURI                 string
}{
	{
		OrigURI:   `/logs-*/_search`,
		OrigQuery: `pretty=true`,
		Region:    "us-west-2", Service: "es", SubDomain: "hostname-clusterkey",
		EscapedURI: `/logs-%2A/_search`,
		ExpSig:     `AWS4-HMAC-SHA256 Credential=AKID/19700101/us-west-2/es/aws4_request, SignedHeaders=host;x-amz-date;x-amz-security-token, Signature=79d0760751907af16f64a537c1242416dacf51204a7dd5284492d15577973b91`,
	},
}

func TestPresignHandler(t *testing.T) {
	svc := s3.New(unit.Session)
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket:             aws.String("bucket"),
		Key:                aws.String("key"),
		ContentDisposition: aws.String("a+b c$d"),
		ACL:                aws.String("public-read"),
	})
	req.Time = time.Unix(0, 0)
	urlstr, err := req.Presign(5 * time.Minute)

	assert.NoError(t, err)

	expectedDate := "19700101T000000Z"
	expectedHeaders := "content-disposition;host;x-amz-acl"
	expectedSig := "b2754ba8ffeb74a40b94767017e24c4672107d6d5a894648d5d332ca61f5ffe4"
	expectedCred := "AKID/19700101/mock-region/s3/aws4_request"

	u, _ := url.Parse(urlstr)
	urlQ := u.Query()
	assert.Equal(t, expectedSig, urlQ.Get("X-Amz-Signature"))
	assert.Equal(t, expectedCred, urlQ.Get("X-Amz-Credential"))
	assert.Equal(t, expectedHeaders, urlQ.Get("X-Amz-SignedHeaders"))
	assert.Equal(t, expectedDate, urlQ.Get("X-Amz-Date"))
	assert.Equal(t, "300", urlQ.Get("X-Amz-Expires"))

	assert.NotContains(t, urlstr, "+") // + encoded as %20
}

func TestPresignRequest(t *testing.T) {
	svc := s3.New(unit.Session)
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket:             aws.String("bucket"),
		Key:                aws.String("key"),
		ContentDisposition: aws.String("a+b c$d"),
		ACL:                aws.String("public-read"),
	})
	req.Time = time.Unix(0, 0)
	urlstr, headers, err := req.PresignRequest(5 * time.Minute)

	assert.NoError(t, err)

	expectedDate := "19700101T000000Z"
	expectedHeaders := "content-disposition;host;x-amz-acl;x-amz-content-sha256"
	expectedSig := "0d200ba61501d752acd06f39ef4dbe7d83ffd5ea15978dc3476dfc00b8eb574e"
	expectedCred := "AKID/19700101/mock-region/s3/aws4_request"
	expectedHeaderMap := http.Header{
		"x-amz-acl":            []string{"public-read"},
		"content-disposition":  []string{"a+b c$d"},
		"x-amz-content-sha256": []string{"UNSIGNED-PAYLOAD"},
	}

	u, _ := url.Parse(urlstr)
	urlQ := u.Query()
	assert.Equal(t, expectedSig, urlQ.Get("X-Amz-Signature"))
	assert.Equal(t, expectedCred, urlQ.Get("X-Amz-Credential"))
	assert.Equal(t, expectedHeaders, urlQ.Get("X-Amz-SignedHeaders"))
	assert.Equal(t, expectedDate, urlQ.Get("X-Amz-Date"))
	assert.Equal(t, expectedHeaderMap, headers)
	assert.Equal(t, "300", urlQ.Get("X-Amz-Expires"))

	assert.NotContains(t, urlstr, "+") // + encoded as %20
}

func TestStandaloneSign_CustomURIEscape(t *testing.T) {
	var expectSig = `AWS4-HMAC-SHA256 Credential=AKID/19700101/us-east-1/es/aws4_request, SignedHeaders=host;x-amz-date;x-amz-security-token, Signature=0c2bebb26be55ecfb2c51068a753da9242910b551d01738f20232b116bff9c24`
	var customEscapeCalled bool

	creds := unit.Session.Config.Credentials
	signer := v4.NewSigner(creds, func(s *v4.Signer) {
		s.URIPathEscapeFn = func(in string) string {
			customEscapeCalled = true
			return `/log-%252A/_search`
		}
	})

	host := "https://subdomain.us-east-1.es.amazonaws.com"
	req, err := http.NewRequest("GET", host, nil)
	assert.NoError(t, err)

	req.URL.Path = `/log-*/_search`

	_, err = signer.Sign(req, nil, "es", "us-east-1", time.Unix(0, 0))
	assert.NoError(t, err)

	actual := req.Header.Get("Authorization")
	assert.Equal(t, expectSig, actual)
	assert.True(t, customEscapeCalled)
}
