package v2

import (
	"bytes"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/stretchr/testify/assert"
)

func buildSigner(serviceName string, region string, signTime time.Time, query url.Values) signer {
	endpoint := "https://" + serviceName + "." + region + ".amazonaws.com"
	body := []byte(query.Encode())
	reader := bytes.NewReader(body)
	req, _ := http.NewRequest("POST", endpoint, reader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", string(len(body)))

	signer := signer{
		Request: req,
		Time:    signTime,
		Credentials: credentials.NewStaticCredentials(
			"AKID",
			"SECRET",
			"SESSION"),
	}

	if os.Getenv("DEBUG") != "" {
		signer.Debug = aws.LogDebug
		signer.Logger = aws.NewDefaultLogger()
	}

	return signer
}

func TestSimpleSignRequest(t *testing.T) {
	query := make(url.Values)
	query.Add("Action", "CreateDomain")
	query.Add("DomainName", "TestDomain-1437033376")
	query.Add("Version", "2009-04-15")

	timestamp := time.Date(2015, 7, 16, 7, 56, 16, 0, time.UTC)
	signer := buildSigner("sdb", "ap-southeast-2", timestamp, query)

	err := signer.Sign()
	assert.Nil(t, err)
	assert.Equal(t, "Ch6qv3rzXB1SLqY2vFhsgA1WQ9rnQIE2WJCigOvAJwI=", signer.signature)
	assert.Equal(t, 9, len(signer.Query))
	assert.Equal(t, "AKID", signer.Query.Get("AWSAccessKeyId"))
	assert.Equal(t, "2015-07-16T07:56:16Z", signer.Query.Get("Timestamp"))
	assert.Equal(t, "HmacSHA256", signer.Query.Get("SignatureMethod"))
	assert.Equal(t, "2", signer.Query.Get("SignatureVersion"))
	assert.Equal(t, "Ch6qv3rzXB1SLqY2vFhsgA1WQ9rnQIE2WJCigOvAJwI=", signer.Query.Get("Signature"))
	assert.Equal(t, "CreateDomain", signer.Query.Get("Action"))
	assert.Equal(t, "TestDomain-1437033376", signer.Query.Get("DomainName"))
	assert.Equal(t, "2009-04-15", signer.Query.Get("Version"))
}

func TestMoreComplexSignRequest(t *testing.T) {
	query := make(url.Values)
	query.Add("Action", "PutAttributes")
	query.Add("DomainName", "TestDomain-1437041569")
	query.Add("Version", "2009-04-15")
	query.Add("Attribute.2.Name", "Attr2")
	query.Add("Attribute.2.Value", "Value2")
	query.Add("Attribute.2.Replace", "true")
	query.Add("Attribute.1.Name", "Attr1-%\\+ %")
	query.Add("Attribute.1.Value", " \tValue1 +!@#$%^&*(){}[]\"';:?/.>,<\x12\x00")
	query.Add("Attribute.1.Replace", "true")
	query.Add("ItemName", "Item 1")

	timestamp := time.Date(2015, 7, 16, 10, 12, 51, 0, time.UTC)
	signer := buildSigner("sdb", "ap-southeast-2", timestamp, query)

	err := signer.Sign()
	assert.Nil(t, err)
	assert.Equal(t, "WNdE62UJKLKoA6XncVY/9RDbrKmcVMdQPQOTAs8SgwQ=", signer.signature)
}

func TestGet(t *testing.T) {
	assert := assert.New(t)
	svc := service.New(&aws.Config{
		Credentials: credentials.NewStaticCredentials("AKID", "SECRET", "TOKEN"),
		Region:      aws.String("ap-southeast-2"),
	})
	r := svc.NewRequest(
		&request.Operation{
			Name:       "OpName",
			HTTPMethod: "GET",
			HTTPPath:   "/",
		},
		nil,
		nil,
	)

	r.Build()
	assert.Equal("GET", r.HTTPRequest.Method)
	assert.Equal("", r.HTTPRequest.URL.Query().Get("Signature"))

	Sign(r)
	assert.NoError(r.Error)
	t.Logf("Signature: %s", r.HTTPRequest.URL.Query().Get("Signature"))
	assert.NotEqual("", r.HTTPRequest.URL.Query().Get("Signature"))
}

func TestAnonymousCredentials(t *testing.T) {
	svc := service.New(&aws.Config{
		Credentials: credentials.AnonymousCredentials,
		Region:      aws.String("ap-southeast-2"),
	})
	r := svc.NewRequest(
		&request.Operation{
			Name:       "PutAttributes",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		},
		nil,
		nil,
	)
	r.Build()

	Sign(r)

	req := r.HTTPRequest
	req.ParseForm()

	assert.Empty(t, req.PostForm.Get("Signature"))
}
