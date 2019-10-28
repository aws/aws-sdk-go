// +build go1.7

package ec2metadata_test

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/awstesting/unit"
)

const instanceIdentityDocument = `{
  "devpayProductCodes" : null,
  "marketplaceProductCodes" : [ "1abc2defghijklm3nopqrs4tu" ], 
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

type testInput struct {
	name    string
	path    string
	headers map[string]string
	method  string
	resp    string
}

func initInsecureServer(input testInput) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.RequestURI {
		case "/latest/api/token":
			http.Error(w, "not found", http.StatusNotFound)
			return

		case input.path:
			break
		default:
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		if r.Method != input.method {
			http.Error(w, "Incorrect http method called", http.StatusBadRequest)
			return
		}

		w.Write([]byte(input.resp))
	}))
}

func initSecureServer(input testInput) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.RequestURI {
		case "/latest/api/token":
			if r.Method == "PUT" && r.Header.Get("x-aws-ec2-metadata-token-ttl-seconds") != "" {
				w.Header().Set("X-aws-ec2-metadata-token-ttl-seconds", "21600")
				w.Write([]byte("Aeioubcdefghijklmnopqrstuvw"))
				return
			}
			http.Error(w, "bad request", http.StatusBadRequest)
			return

		case input.path:
			if r.Method == "GET" && r.Header.Get("x-aws-ec2-metadata-token") == "Aeioubcdefghijklmnopqrstuvw" {
				break
			}
			http.Error(w, "Unauthorized for an expired, invalid, or missing token header", http.StatusUnauthorized)
			return

		default:
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		if r.Method != input.method {
			http.Error(w, "Incorrect http method called", http.StatusBadRequest)
			return
		}

		w.Write([]byte(input.resp))
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
	cases := map[string]struct {
		input            testInput
		server           func(input testInput) *httptest.Server
		expectedResponse string
	}{
		"Insecure server": {
			testInput{
				path:    "/latest/meta-data/some/path",
				headers: nil,
				method:  "GET",
				resp:    "success", // real response includes suffix ,
			},
			initInsecureServer,
			"success",
		},
		"Secure server": {
			testInput{
				path: "/latest/meta-data/some/path",
				headers: map[string]string{
					"x-aws-ec2-metadata-token": "Aeioubcdefghijklmnopqrstuv",
				},
				method: "GET",
				resp:   "success", // real response includes suffix ,
			},
			initSecureServer,
			"success",
		},
	}

	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			server := in.server(in.input)
			defer server.Close()
			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			resp, err := c.GetMetadata("some/path")

			if err != nil {
				t.Errorf("expect no error, got %v", err)
			}
			if e, a := in.expectedResponse, resp; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

func TestGetUserData(t *testing.T) {
	cases := map[string]struct {
		input            testInput
		server           func(input testInput) *httptest.Server
		expectedResponse string
	}{
		"Insecure server": {
			testInput{
				path:    "/latest/user-data",
				headers: nil,
				method:  "GET",
				resp:    "success", // real response includes suffix ,
			},
			initInsecureServer,
			"success",
		},
		"Secure server": {
			testInput{
				path: "/latest/user-data",
				headers: map[string]string{
					"x-aws-ec2-metadata-token": "Aeioubcdefghijklmnopqrstuv",
				},
				method: "GET",
				resp:   "success", // real response includes suffix ,
			},
			initSecureServer,
			"success",
		},
	}

	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			server := in.server(in.input)
			defer server.Close()
			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			resp, err := c.GetUserData()

			if err != nil {
				t.Errorf("expect no error, got %v", err)
			}
			if e, a := in.expectedResponse, resp; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
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
	cases := map[string]struct {
		input            testInput
		server           func(input testInput) *httptest.Server
		expectedResponse string
	}{
		"Insecure server": {
			testInput{
				path:    "/latest/meta-data/placement/availability-zone",
				headers: nil,
				method:  "GET",
				resp:    "us-west-2a", // real response includes suffix ,
			},
			initInsecureServer,
			"us-west-2",
		},
		"Secure server": {
			testInput{
				path: "/latest/meta-data/placement/availability-zone",
				headers: map[string]string{
					"x-aws-ec2-metadata-token": "Aeioubcdefghijklmnopqrstuv",
				},
				method: "GET",
				resp:   "us-west-2a", // real response includes suffix ,
			},
			initSecureServer,
			"us-west-2",
		},
	}

	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			server := in.server(in.input)
			defer server.Close()
			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			region, err := c.Region()

			if err != nil {
				t.Errorf("expect no error, got %v", err)
			}
			if e, a := in.expectedResponse, region; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}


func TestGetRegion_invalidResponse(t *testing.T) {
	cases := map[string]struct {
		input            testInput
		server           func(input testInput) *httptest.Server
		expectedResponse string
	}{
		"Insecure server": {
			testInput{
				path:    "/latest/meta-data/placement/availability-zone",
				headers: nil,
				method:  "GET",
				resp:    "", // real response includes suffix ,
			},
			initInsecureServer,
			"",
		},
		"Secure server": {
			testInput{
				path: "/latest/meta-data/placement/availability-zone",
				headers: map[string]string{
					"x-aws-ec2-metadata-token": "Aeioubcdefghijklmnopqrstuv",
				},
				method: "GET",
				resp:   "", // real response includes suffix ,
			},
			initSecureServer,
			"",
		},
	}

	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			server := in.server(in.input)
			defer server.Close()
			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			region, err := c.Region()

			if err == nil {
				t.Errorf("expect no error, got %v", err)
			}
			if e, a := in.expectedResponse, region; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

func TestMetadataAvailable(t *testing.T) {
	cases := map[string]struct {
		input            testInput
		server           func(input testInput) *httptest.Server
		expectedResponse string
	}{
		"Insecure server": {
			testInput{
				path:    "/latest/meta-data/instance-id",
				headers: nil,
				method:  "GET",
				resp:    "instance-id", // real response includes suffix ,
			},
			initInsecureServer,
			"instance-id",
		},
		"Secure server": {
			testInput{
				path: "/latest/meta-data/instance-id",
				headers: map[string]string{
					"x-aws-ec2-metadata-token": "Aeioubcdefghijklmnopqrstuv",
				},
				method: "GET",
				resp:   "instance-id", // real response includes suffix ,
			},
			initSecureServer,
			"instance-id",
		},
	}

	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			server := in.server(in.input)
			defer server.Close()
			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})

			if !c.Available() {
				t.Errorf("expect available")
			}
		})
	}
}


func TestMetadataIAMInfo_success(t *testing.T) {
	cases := map[string]struct {
		input            testInput
		server           func(input testInput) *httptest.Server
		expectedResponse string
	}{
		"Insecure server": {
			testInput{
				path:    "/latest/meta-data/iam/info",
				headers: nil,
				method:  "GET",
				resp:    validIamInfo, // real response includes suffix ,
			},
			initInsecureServer,
			validIamInfo,
		},
		"Secure server": {
			testInput{
				path: "/latest/meta-data/iam/info",
				headers: map[string]string{
					"x-aws-ec2-metadata-token": "Aeioubcdefghijklmnopqrstuv",
				},
				method: "GET",
				resp:   validIamInfo, // real response includes suffix ,
			},
			initSecureServer,
			validIamInfo,
		},
	}

	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			server := in.server(in.input)
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
		})
	}
}


func TestMetadataIAMInfo_failure(t *testing.T) {
	cases := map[string]struct {
		input            testInput
		server           func(input testInput) *httptest.Server
		expectedResponse string
	}{
		"Insecure server": {
			testInput{
				path:    "/latest/meta-data/iam/info",
				headers: nil,
				method:  "GET",
				resp:    unsuccessfulIamInfo, // real response includes suffix ,
			},
			initInsecureServer,
			unsuccessfulIamInfo,
		},
		"Secure server": {
			testInput{
				path: "/latest/meta-data/iam/info",
				headers: map[string]string{
					"x-aws-ec2-metadata-token": "Aeioubcdefghijklmnopqrstuv",
				},
				method: "GET",
				resp:   unsuccessfulIamInfo, // real response includes suffix ,
			},
			initSecureServer,
			unsuccessfulIamInfo,
		},
	}

	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			server := in.server(in.input)
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
		})
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
	cases := map[string]struct {
		input            testInput
		server           func(input testInput) *httptest.Server
		expectedResponse string
	}{
		"Insecure server": {
			testInput{
				path:    "/latest/dynamic/instance-identity/document",
				headers: nil,
				method:  "GET",
				resp:    instanceIdentityDocument, // real response includes suffix ,
			},
			initInsecureServer,
			instanceIdentityDocument,
		},
		"Secure server": {
			testInput{
				path: "/latest/dynamic/instance-identity/document",
				headers: map[string]string{
					"x-aws-ec2-metadata-token": "Aeioubcdefghijklmnopqrstuv",
				},
				method: "GET",
				resp:   instanceIdentityDocument, // real response includes suffix ,
			},
			initSecureServer,
			instanceIdentityDocument,
		},
	}

	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			server := in.server(in.input)
			defer server.Close()
			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			doc, err := c.GetInstanceIdentityDocument()
			if err != nil {
				t.Errorf("expect no error, got %v", err)
			}
			if e, a := doc.AccountID, "123456789012"; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
			if e, a := doc.AvailabilityZone, "us-east-1d"; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
			if e, a := doc.Region, "us-east-1"; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

