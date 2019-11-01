// +build go1.7

package ec2metadata_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awserr"

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

type testServer struct {
	t                          *testing.T
	tokenProvider              *tokenAccessProvider
	dataProvider               *dataAccessProvider
	expectedTokens             []int
	expectedTokenProviderError string
}

type tokenAccessProvider struct {
	token      int
	HTTPMethod string
}

type dataAccessProvider struct {
	data       string
	HTTPMethod string
}

func initSecureServer(s testServer) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/latest/api/token", s.secureTokenHandler)
	mux.HandleFunc("/latest/", s.secureMetadataHandler)
	return httptest.NewServer(mux)
}

func initInsecureServer(s testServer) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/latest/api/token", s.insecureTokenHandler)
	mux.HandleFunc("/latest/", s.insecureMetadataHandler)
	return httptest.NewServer(mux)
}

func (s *testServer) insecureTokenHandler(w http.ResponseWriter, r *http.Request) {
	s.t.Logf("Insecure token handler called \n")
	http.Error(w, "not found", http.StatusNotFound)
	if !strings.Contains("not found", s.expectedTokenProviderError) {
		s.t.Errorf("Expected %v, got not found error \n", s.expectedTokenProviderError)
	}
}

func (s *testServer) insecureMetadataHandler(w http.ResponseWriter, r *http.Request) {
	s.t.Logf("Insecure metadata handler called \n")

	if r.Method != s.dataProvider.HTTPMethod {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	w.Write([]byte(s.dataProvider.data))
}

func (s *testServer) secureTokenHandler(w http.ResponseWriter, r *http.Request) {
	s.t.Logf("Secure token handler called \n")
	if r.Method != s.tokenProvider.HTTPMethod {
		http.Error(w, "bad request", http.StatusBadRequest)
		if !strings.Contains("bad request", s.expectedTokenProviderError) {
			s.t.Errorf("Expected %v, got bad request error \n", s.expectedTokenProviderError)
		}
		return
	}

	if r.Header.Get(ec2metadata.TTLHeader) == "" {
		http.Error(w, "TTL not set for token retrieval", http.StatusBadRequest)
		if !strings.Contains("TTL not set for token retrieval", s.expectedTokenProviderError) {
			s.t.Errorf("Expected %v, got TTL not set for token retrieval error \n", s.expectedTokenProviderError)
		}
		return
	}

	// increment token for each call to secure server for token
	s.tokenProvider.token++

	w.Header().Set(ec2metadata.TTLHeader, r.Header.Get(ec2metadata.TTLHeader))
	w.Write([]byte(strconv.Itoa(s.tokenProvider.token)))
}

func (s *testServer) secureMetadataHandler(w http.ResponseWriter, r *http.Request) {
	s.t.Logf("secure metadata handler called \n")

	if r.Method != s.dataProvider.HTTPMethod {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if token := r.Header.Get(ec2metadata.TokenHeader); token != "" {
		if i, ok := strconv.Atoi(token); ok == nil && i != s.expectedTokens[i-1] {
			http.Error(w, "Unauthorized for an expired, invalid, or missing token header", http.StatusUnauthorized)
			return
		}
	}

	w.Header().Set(ec2metadata.TTLHeader, r.Header.Get(ec2metadata.TTLHeader))
	w.Write([]byte(s.dataProvider.data))
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
		s                         testServer
		expectedDataResponse      string
		expectedDataProviderError string
		serverType                string
	}{
		"Insecure server success case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       "profile_name",
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{},
				expectedTokenProviderError: "not found",
			},
			expectedDataResponse:      "profile_name",
			expectedDataProviderError: "",
			serverType:                "insecure",
		},
		"Insecure server failure case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       "profile_name",
					HTTPMethod: "PUT",
				},
				expectedTokens:             []int{},
				expectedTokenProviderError: "not found",
			},
			expectedDataResponse:      "",
			expectedDataProviderError: "bad request",
			serverType:                "insecure",
		},
		"Secure server success case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       "profile_name",
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{1, 2, 3},
				expectedTokenProviderError: "",
			},
			expectedDataResponse:      "profile_name",
			expectedDataProviderError: "",
			serverType:                "secure",
		},
		"Secure server failure case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       "profile_name",
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{0, 1, 2},
				expectedTokenProviderError: "",
			},
			expectedDataResponse:      "",
			expectedDataProviderError: "Unauthorized",
			serverType:                "secure",
		},
	}

	for name, x := range cases {
		t.Run(name, func(t *testing.T) {
			var server *httptest.Server
			// change this
			if x.serverType == "insecure" {
				server = initInsecureServer(x.s)
			} else {
				server = initSecureServer(x.s)
			}

			defer server.Close()

			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			resp, err := c.GetMetadata("some/path")

			// Set Token TTL to 0, this will reset the token, thus there should be a change in token
			c.SetTokenTTL(0)
			resp, err = c.GetMetadata("some/path")

			// since earlier token was expires this should fetch new token
			resp, err = c.GetMetadata("some/path")

			if err != nil && x.expectedDataProviderError == "" {
				t.Errorf("expected no error, got %v", err)
			}

			if e, a := x.expectedDataProviderError, err; a != nil {
				if !strings.Contains(a.Error(), e) {
					t.Errorf("expect %v, got %v", x.expectedDataProviderError, err)
				}
			} else if x.expectedDataProviderError != "" {
				t.Errorf("expected %v, got none", x.expectedDataProviderError)
			}

			if e, a := x.expectedDataResponse, resp; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

func TestGetUserData(t *testing.T) {
	cases := map[string]struct {
		s                         testServer
		expectedDataResponse      string
		expectedDataProviderError string
		serverType                string
	}{
		"Insecure server success case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       "user_data",
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{},
				expectedTokenProviderError: "not found",
			},
			expectedDataResponse:      "user_data",
			expectedDataProviderError: "",
			serverType:                "insecure",
		},
		"Insecure server failure case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       "user_data",
					HTTPMethod: "PUT",
				},
				expectedTokens:             []int{},
				expectedTokenProviderError: "not found",
			},
			expectedDataResponse:      "",
			expectedDataProviderError: "bad request",
			serverType:                "insecure",
		},
		"Secure server success case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       "user_data",
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{1},
				expectedTokenProviderError: "",
			},
			expectedDataResponse:      "user_data",
			expectedDataProviderError: "",
			serverType:                "secure",
		},
		"Secure server failure case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       "user_data",
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{0},
				expectedTokenProviderError: "",
			},
			expectedDataResponse:      "",
			expectedDataProviderError: "Unauthorized",
			serverType:                "secure",
		},
	}
	for name, x := range cases {
		t.Run(name, func(t *testing.T) {
			var server *httptest.Server
			// change this
			if x.serverType == "insecure" {
				server = initInsecureServer(x.s)
			} else {
				server = initSecureServer(x.s)
			}

			defer server.Close()

			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			resp, err := c.GetUserData()

			if err != nil && x.expectedDataProviderError == "" {
				t.Errorf("expected no error, got %v", err)
			}

			if e, a := x.expectedDataProviderError, err; a != nil {
				if !strings.Contains(a.Error(), e) {
					t.Errorf("expect %v, got %v", x.expectedDataProviderError, err)
				}
			} else if x.expectedDataProviderError != "" {
				t.Errorf("expected %v, got none", x.expectedDataProviderError)
			}

			if e, a := x.expectedDataResponse, resp; e != a {
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

	if requestFailedError, ok := err.(awserr.RequestFailure); ok {
		if e, a := http.StatusNotFound, requestFailedError.StatusCode(); e != a {
			t.Errorf("expect %v, got %v", e, a)
		}
	}
}

func TestGetRegion(t *testing.T) {
	cases := map[string]struct {
		s                         testServer
		expectedDataResponse      string
		expectedDataProviderError string
		serverType                string
	}{
		"Insecure server success case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       "us-west-2a",
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{},
				expectedTokenProviderError: "not found",
			},
			expectedDataResponse:      "us-west-2",
			expectedDataProviderError: "",
			serverType:                "insecure",
		},
		"Insecure server failure case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       "us-west-2a",
					HTTPMethod: "PUT",
				},
				expectedTokens:             []int{},
				expectedTokenProviderError: "not found",
			},
			expectedDataResponse:      "",
			expectedDataProviderError: "bad request",
			serverType:                "insecure",
		},
		"Secure server success case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       "us-west-2a",
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{1},
				expectedTokenProviderError: "",
			},
			expectedDataResponse:      "us-west-2",
			expectedDataProviderError: "",
			serverType:                "secure",
		},
		"Secure server failure case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       "us-west-2a",
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{0},
				expectedTokenProviderError: "",
			},
			expectedDataResponse:      "",
			expectedDataProviderError: "Unauthorized",
			serverType:                "secure",
		},
	}
	for name, x := range cases {
		t.Run(name, func(t *testing.T) {
			var server *httptest.Server
			// change this
			if x.serverType == "insecure" {
				server = initInsecureServer(x.s)
			} else {
				server = initSecureServer(x.s)
			}

			defer server.Close()

			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			resp, err := c.Region()

			if err != nil && x.expectedDataProviderError == "" {
				t.Errorf("expected no error, got %v", err)
			}

			if e, a := x.expectedDataProviderError, err; a != nil {
				if !strings.Contains(a.Error(), e) {
					t.Errorf("expect %v, got %v", x.expectedDataProviderError, err)
				}
			} else if x.expectedDataProviderError != "" {
				t.Errorf("expected %v, got none", x.expectedDataProviderError)
			}

			if e, a := x.expectedDataResponse, resp; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

func TestMetadataAvailable(t *testing.T) {
	cases := map[string]struct {
		s                    testServer
		expectedDataResponse bool
		serverType           string
	}{
		"Insecure server success case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       "instance-id",
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{},
				expectedTokenProviderError: "not found",
			},
			expectedDataResponse: true,
			serverType:           "insecure",
		},
		"Insecure server failure case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       "instance-id",
					HTTPMethod: "PUT",
				},
				expectedTokens:             []int{},
				expectedTokenProviderError: "not found",
			},
			expectedDataResponse: false,
			serverType:           "insecure",
		},
		"Secure server success case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       "instance-id",
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{1},
				expectedTokenProviderError: "",
			},
			expectedDataResponse: true,
			serverType:           "secure",
		},
		"Secure server failure case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       "instance-id",
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{0},
				expectedTokenProviderError: "",
			},
			expectedDataResponse: false,
			serverType:           "secure",
		},
	}
	for name, x := range cases {
		t.Run(name, func(t *testing.T) {
			var server *httptest.Server
			// change this
			if x.serverType == "insecure" {
				server = initInsecureServer(x.s)
			} else {
				server = initSecureServer(x.s)
			}

			defer server.Close()

			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			if e, a := x.expectedDataResponse, c.Available(); e != a {
				t.Errorf("expect availability to be %v, got %v \n", e, a)
			}
		})
	}
}

func TestMetadataIAMInfo_success(t *testing.T) {
	cases := map[string]struct {
		s                         testServer
		expectedDataResponse      string
		expectedDataProviderError string
		serverType                string
	}{
		"Insecure server success case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       validIamInfo,
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{},
				expectedTokenProviderError: "not found",
			},
			expectedDataResponse:      validIamInfo,
			expectedDataProviderError: "",
			serverType:                "insecure",
		},
		"Secure server success case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       validIamInfo,
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{1},
				expectedTokenProviderError: "",
			},
			expectedDataResponse:      validIamInfo,
			expectedDataProviderError: "",
			serverType:                "secure",
		},
	}
	for name, x := range cases {
		t.Run(name, func(t *testing.T) {
			var server *httptest.Server
			// change this
			if x.serverType == "insecure" {
				server = initInsecureServer(x.s)
			} else {
				server = initSecureServer(x.s)
			}

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
		s                         testServer
		expectedDataResponse      string
		expectedDataProviderError string
		serverType                string
	}{
		"Insecure server failure case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       unsuccessfulIamInfo,
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{},
				expectedTokenProviderError: "not found",
			},
			expectedDataResponse:      unsuccessfulIamInfo,
			expectedDataProviderError: "",
			serverType:                "insecure",
		},
		"Secure server failure case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       unsuccessfulIamInfo,
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{1},
				expectedTokenProviderError: "",
			},
			expectedDataResponse:      unsuccessfulIamInfo,
			expectedDataProviderError: "",
			serverType:                "secure",
		},
	}
	for name, x := range cases {
		t.Run(name, func(t *testing.T) {
			var server *httptest.Server
			// change this
			if x.serverType == "insecure" {
				server = initInsecureServer(x.s)
			} else {
				server = initSecureServer(x.s)
			}

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
		s                         testServer
		expectedDataResponse      string
		expectedDataProviderError string
		serverType                string
	}{
		"Insecure server success case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       instanceIdentityDocument,
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{},
				expectedTokenProviderError: "not found",
			},
			expectedDataResponse:      instanceIdentityDocument,
			expectedDataProviderError: "",
			serverType:                "insecure",
		},
		"Secure server success case": {
			s: testServer{
				t: t,
				tokenProvider: &tokenAccessProvider{
					token:      0,
					HTTPMethod: "PUT",
				},
				dataProvider: &dataAccessProvider{
					data:       instanceIdentityDocument,
					HTTPMethod: "GET",
				},
				expectedTokens:             []int{1},
				expectedTokenProviderError: "",
			},
			expectedDataResponse:      instanceIdentityDocument,
			expectedDataProviderError: "",
			serverType:                "secure",
		},
	}
	for name, x := range cases {
		t.Run(name, func(t *testing.T) {
			var server *httptest.Server
			// change this
			if x.serverType == "insecure" {
				server = initInsecureServer(x.s)
			} else {
				server = initSecureServer(x.s)
			}

			defer server.Close()

			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			doc, err := c.GetInstanceIdentityDocument()

			if err != nil {
				t.Errorf("Expected no error, got %v \n", err)
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

func TestTokenExpired(t *testing.T) {
	var token = 100
	mux := http.NewServeMux()
	mux.HandleFunc("/latest/api/token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" && r.Header.Get(ec2metadata.TTLHeader) != "" {
			w.Header().Set(ec2metadata.TTLHeader, "0")
			w.Write([]byte(strconv.Itoa(token)))
			token++
			return
		}
		http.Error(w, "bad request", http.StatusBadRequest)
	})

	// meta-data endpoint for this test, just returns the token
	mux.HandleFunc("/latest/meta-data/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.Header.Get(ec2metadata.TokenHeader)))
	})

	server := httptest.NewServer(mux)
	defer server.Close()
	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
	f, _ := c.GetMetadata("some/path")
	s, _ := c.GetMetadata("some/path")

	if f == s {
		t.Errorf("Expected different tokens to be returned, got %v instead", f)
	}
}

func TestTokenAlive(t *testing.T) {
	mux := http.NewServeMux()
	var token = 100
	mux.HandleFunc("/latest/api/token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" && r.Header.Get(ec2metadata.TTLHeader) != "" {
			w.Header().Set(ec2metadata.TTLHeader, "200")
			w.Write([]byte(strconv.Itoa(token)))
			token++
			return
		}
		http.Error(w, "bad request", http.StatusBadRequest)
	})

	// meta-data endpoint for this test, just returns the token
	mux.HandleFunc("/latest/meta-data/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.Header.Get(ec2metadata.TokenHeader)))
	})

	server := httptest.NewServer(mux)
	defer server.Close()
	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
	f, _ := c.GetMetadata("some/path")
	s, _ := c.GetMetadata("some/path")

	if f != s {
		t.Errorf("Expected same token to be returned, got %v instead of %v", f, s)
	}
}

func TestEC2MetadataRetryFailure(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/latest/api/token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" && r.Header.Get(ec2metadata.TTLHeader) != "" {
			w.Header().Set(ec2metadata.TTLHeader, "200")
			http.Error(w, "service unavailable", http.StatusServiceUnavailable)
			t.Logf("%v received, retrying", http.StatusServiceUnavailable)
			return
		}
		http.Error(w, "bad request", http.StatusBadRequest)
	})

	// meta-data endpoint for this test, just returns the token
	mux.HandleFunc("/latest/meta-data/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.Header.Get(ec2metadata.TokenHeader)))
	})

	server := httptest.NewServer(mux)
	defer server.Close()
	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
	_, err := c.GetMetadata("some/path")

	if err != nil {
		t.Errorf("Expected none, got error %v", err)
	}
}

func TestEC2MetadataRetryOnce(t *testing.T) {
	var secureDataFlow bool
	var retry = true
	mux := http.NewServeMux()

	mux.HandleFunc("/latest/api/token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" && r.Header.Get(ec2metadata.TTLHeader) != "" {
			w.Header().Set(ec2metadata.TTLHeader, "200")
			for retry {
				retry = false
				http.Error(w, "service unavailable", http.StatusServiceUnavailable)
				t.Logf("%v received, retrying", http.StatusServiceUnavailable)
				return
			}
			w.Write([]byte("token"))
			secureDataFlow = true
			return
		}
		http.Error(w, "bad request", http.StatusBadRequest)
	})

	// meta-data endpoint for this test, just returns the token
	mux.HandleFunc("/latest/meta-data/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.Header.Get(ec2metadata.TokenHeader)))
	})

	server := httptest.NewServer(mux)
	defer server.Close()
	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
	_, err := c.GetMetadata("some/path")

	if !secureDataFlow {
		t.Errorf("Expected secure data flow to be %v, got %v", secureDataFlow, !secureDataFlow)
	}

	if err != nil {
		t.Errorf("Expected none, got error %v", err)
	}
}

func TestTokenErrorWith400(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/latest/api/token", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "bad request", http.StatusBadRequest)
	})

	// meta-data endpoint for this test, just returns the token
	mux.HandleFunc("/latest/meta-data/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("token"))
	})

	server := httptest.NewServer(mux)
	defer server.Close()
	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
	resp, err := c.GetMetadata("some/path")

	if err == nil {
		t.Error("Expected error, got none")
	}

	if resp == "token" {
		t.Errorf("Expected empty response, got %v", resp)
	}
}

func TestEC2Metadata_Concurrency(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/latest/api/token", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("token"))
	})

	// meta-data endpoint for this test, just returns the token
	mux.HandleFunc("/latest/meta-data/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("profile_name"))
	})

	server := httptest.NewServer(mux)
	defer server.Close()
	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			resp, err := c.GetMetadata("some/data")
			if err != nil  {
				t.Errorf("expect no error, got %v", err)
			}

			if e, a := "profile_name", resp; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
