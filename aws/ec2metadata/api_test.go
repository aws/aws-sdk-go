package ec2metadata_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/awstesting/unit"
)

// Not used, but this is the signed plaintext contained in the PKCS7 data below
const instanceIdentityDocument = `{
  "devpayProductCodes" : null,
  "availabilityZone" : "us-east-1d",
  "privateIp" : "10.158.112.84",
  "version" : "2010-08-31",
  "region" : "us-east-1",
  "instanceId" : "i-1234567890abcdef0",
  "billingProducts" : null,
  "instanceType" : "t1.micro",
  "accountId" : "123456789012",
  "pendingTime" : "2015-11-19T16:32:11Z",
  "imageId" : "ami-5fb8c835",
  "kernelId" : "aki-919dcaf8",
  "ramdiskId" : null,
  "architecture" : "x86_64"
}`

const PKCS7 = `MIIDdQYJKoZIhvcNAQcCoIIDZjCCA2ICAQExCzAJBgUrDgMCGgUAMIIBzgYJKoZI
hvcNAQcBoIIBvwSCAbt7CiAgImRldnBheVByb2R1Y3RDb2RlcyIgOiBudWxsLAog
ICJhdmFpbGFiaWxpdHlab25lIiA6ICJ1cy1lYXN0LTFkIiwKICAicHJpdmF0ZUlw
IiA6ICIxMC4xNTguMTEyLjg0IiwKICAidmVyc2lvbiIgOiAiMjAxMC0wOC0zMSIs
CiAgInJlZ2lvbiIgOiAidXMtZWFzdC0xIiwKICAiaW5zdGFuY2VJZCIgOiAiaS0x
MjM0NTY3ODkwYWJjZGVmMCIsCiAgImJpbGxpbmdQcm9kdWN0cyIgOiBudWxsLAog
ICJpbnN0YW5jZVR5cGUiIDogInQxLm1pY3JvIiwKICAiYWNjb3VudElkIiA6ICIx
MjM0NTY3ODkwMTIiLAogICJwZW5kaW5nVGltZSIgOiAiMjAxNS0xMS0xOVQxNjoz
MjoxMVoiLAogICJpbWFnZUlkIiA6ICJhbWktNWZiOGM4MzUiLAogICJrZXJuZWxJ
ZCIgOiAiYWtpLTkxOWRjYWY4IiwKICAicmFtZGlza0lkIiA6IG51bGwsCiAgImFy
Y2hpdGVjdHVyZSIgOiAieDg2XzY0Igp9CjGCAXwwggF4AgEBMFIwRTELMAkGA1UE
BhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoMGEludGVybmV0IFdp
ZGdpdHMgUHR5IEx0ZAIJAILBjm7QgIp6MAkGBSsOAwIaBQCggdgwGAYJKoZIhvcN
AQkDMQsGCSqGSIb3DQEHATAcBgkqhkiG9w0BCQUxDxcNMTcxMDE4MjMyMzA4WjAj
BgkqhkiG9w0BCQQxFgQUjo+fcC+mFKZW46pOQppLyhF4vq4weQYJKoZIhvcNAQkP
MWwwajALBglghkgBZQMEASowCwYJYIZIAWUDBAEWMAsGCWCGSAFlAwQBAjAKBggq
hkiG9w0DBzAOBggqhkiG9w0DAgICAIAwDQYIKoZIhvcNAwICAUAwBwYFKw4DAgcw
DQYIKoZIhvcNAwICASgwCQYHKoZIzjgEAwQuMCwCFEswo7mYpSI+Jm3SEXRBu3Lu
U4jQAhQ04k3XPWh6qC6MXFgPpaCqCogJGA==`

// Certificate created for test purpose. The corresponding private key
// is checked in testdata/
const testCert = `-----BEGIN CERTIFICATE-----
MIIDGjCCAtigAwIBAgIJAILBjm7QgIp6MAsGCWCGSAFlAwQDAjBFMQswCQYDVQQG
EwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50ZXJuZXQgV2lk
Z2l0cyBQdHkgTHRkMB4XDTE3MTAxODIyMDAwM1oXDTE3MTExNzIyMDAwM1owRTEL
MAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoMGEludGVy
bmV0IFdpZGdpdHMgUHR5IEx0ZDCCAbcwggEsBgcqhkjOOAQBMIIBHwKBgQCVNC8S
zv/6FpuxWsC9fzVmXQPyjxUc98xFcm75EneF6D+qsefSJeubcnWVRmLmNl1fU+iH
/BsNJbM0nXGNvMmT2pJCREsB8lNK7JkqfSpzYbwYoVtnTVppWTMZf4GO9i0fWDE0
vwffVviGWgTaBIKpcKaP+BvPt4d6OkBAiErcXQIVAI7aoB4F5deSrt8N0+pTKV/l
FhHdAoGBAJHkM1yXL5G/AhwnQeNH5Sp7e8wEnUpvA6wo6vdHJ1WlRlPeClKr4zgW
XYr3+1wNmIFW8nCnR5jXPrzBLMo7MZzlOk+1ogbgNG7nCCUMcFXYXRfYSYKoUrrA
/356mDD9ruiHb4HwwKq7ah22+6050cTXwvnigfvGDCt1NJE/GSLeA4GEAAKBgBq1
t3iLtyZVs8kxriLBapt+qPaYGCZ2MBruZRyDQNwaAeOvExMDZnja/CbxPyLTXWpT
E3LSFVAnvieAXE0ZcNg+RG5tYBr99g4ZlWm/XJsnzlglEtUxzL6ZndjKSNfWYIm/
8dKFVqq4BHvG5J+B4UGoIIQiSvsgLx0uEq5ezLaoo1AwTjAdBgNVHQ4EFgQUhaBr
c5u63ltwBhuB8IpH9aY5RUowHwYDVR0jBBgwFoAUhaBrc5u63ltwBhuB8IpH9aY5
RUowDAYDVR0TBAUwAwEB/zALBglghkgBZQMEAwIDLwAwLAIUK7XuiiURX/w5WagP
Y7uo+rp7S/oCFA15khEHbJ6oNutjJYwC60WmENT9
-----END CERTIFICATE-----`

const validIamInfo = `{
  "Code" : "Success",
  "LastUpdated" : "2016-03-17T12:27:32Z",
  "InstanceProfileArn" : "arn:aws:iam::123456789012:instance-profile/my-instance-profile",
  "InstanceProfileId" : "AIPAABCDEFGHIJKLMN123"
}`

const unsuccessfulIamInfo = `{
  "Code" : "Failed",
  "LastUpdated" : "2016-03-17T12:27:32Z",
  "InstanceProfileArn" : "arn:aws:iam::123456789012:instance-profile/my-instance-profile",
  "InstanceProfileId" : "AIPAABCDEFGHIJKLMN123"
}`

func initTestServer(path string, resp string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != path {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		w.Write([]byte(resp))
	}))
}

func TestEndpoint(t *testing.T) {
	c := ec2metadata.New(unit.Session)
	op := &request.Operation{
		Name:       "GetMetadata",
		HTTPMethod: "GET",
		HTTPPath:   path.Join("/", "meta-data", "testpath"),
	}

	req := c.NewRequest(op, nil, nil)
	if e, a := "http://169.254.169.254/latest", req.ClientInfo.Endpoint; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "http://169.254.169.254/latest/meta-data/testpath", req.HTTPRequest.URL.String(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestGetMetadata(t *testing.T) {
	server := initTestServer(
		"/latest/meta-data/some/path",
		"success", // real response includes suffix
	)
	defer server.Close()
	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})

	resp, err := c.GetMetadata("some/path")

	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if e, a := "success", resp; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestGetUserData(t *testing.T) {
	server := initTestServer(
		"/latest/user-data",
		"success", // real response includes suffix
	)
	defer server.Close()
	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})

	resp, err := c.GetUserData()

	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if e, a := "success", resp; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestGetUserData_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reader := strings.NewReader(`<?xml version="1.0" encoding="iso-8859-1"?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
         "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
 <head>
  <title>404 - Not Found</title>
 </head>
 <body>
  <h1>404 - Not Found</h1>
 </body>
</html>`)
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", reader.Len()))
		w.WriteHeader(http.StatusNotFound)
		io.Copy(w, reader)
	}))

	defer server.Close()
	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})

	resp, err := c.GetUserData()
	if err == nil {
		t.Errorf("expect error")
	}
	if len(resp) != 0 {
		t.Errorf("expect empty, got %v", resp)
	}

	aerr := err.(awserr.Error)
	if e, a := "NotFoundError", aerr.Code(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestGetRegion(t *testing.T) {
	server := initTestServer(
		"/latest/meta-data/placement/availability-zone",
		"us-west-2a", // real response includes suffix
	)
	defer server.Close()
	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})

	region, err := c.Region()

	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if e, a := "us-west-2", region; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestMetadataAvailable(t *testing.T) {
	server := initTestServer(
		"/latest/meta-data/instance-id",
		"instance-id",
	)
	defer server.Close()
	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})

	if !c.Available() {
		t.Errorf("expect available")
	}
}

func TestMetadataIAMInfo_success(t *testing.T) {
	server := initTestServer(
		"/latest/meta-data/iam/info",
		validIamInfo,
	)
	defer server.Close()
	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})

	iamInfo, err := c.IAMInfo()
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if e, a := "Success", iamInfo.Code; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "arn:aws:iam::123456789012:instance-profile/my-instance-profile", iamInfo.InstanceProfileArn; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "AIPAABCDEFGHIJKLMN123", iamInfo.InstanceProfileID; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestMetadataIAMInfo_failure(t *testing.T) {
	server := initTestServer(
		"/latest/meta-data/iam/info",
		unsuccessfulIamInfo,
	)
	defer server.Close()
	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})

	iamInfo, err := c.IAMInfo()
	if err == nil {
		t.Errorf("expect error")
	}
	if e, a := "", iamInfo.Code; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "", iamInfo.InstanceProfileArn; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "", iamInfo.InstanceProfileID; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestMetadataNotAvailable(t *testing.T) {
	c := ec2metadata.New(unit.Session)
	c.Handlers.Send.Clear()
	c.Handlers.Send.PushBack(func(r *request.Request) {
		r.HTTPResponse = &http.Response{
			StatusCode: int(0),
			Status:     http.StatusText(int(0)),
			Body:       ioutil.NopCloser(bytes.NewReader([]byte{})),
		}
		r.Error = awserr.New("RequestError", "send request failed", nil)
		r.Retryable = aws.Bool(true) // network errors are retryable
	})

	if c.Available() {
		t.Errorf("expect not available")
	}
}

func TestMetadataErrorResponse(t *testing.T) {
	c := ec2metadata.New(unit.Session)
	c.Handlers.Send.Clear()
	c.Handlers.Send.PushBack(func(r *request.Request) {
		r.HTTPResponse = &http.Response{
			StatusCode: http.StatusBadRequest,
			Status:     http.StatusText(http.StatusBadRequest),
			Body:       ioutil.NopCloser(strings.NewReader("error message text")),
		}
		r.Retryable = aws.Bool(false) // network errors are retryable
	})

	data, err := c.GetMetadata("uri/path")
	if len(data) != 0 {
		t.Errorf("expect empty, got %v", data)
	}
	if e, a := "error message text", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v to be in %v", e, a)
	}
}

func TestEC2RoleProviderInstanceIdentity(t *testing.T) {
	server := initTestServer(
		"/latest/dynamic/instance-identity/pkcs7",
		PKCS7,
	)
	defer server.Close()
	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})

	doc, err := c.GetInstanceIdentityDocument(testCert)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if e, a := "123456789012", doc.AccountID; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "us-east-1d", doc.AvailabilityZone; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "us-east-1", doc.Region; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}
