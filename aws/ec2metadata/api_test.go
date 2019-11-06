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
	"strings"
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
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

const (
	ttlHeader   = "x-aws-ec2-metadata-token-ttl-seconds"
	tokenHeader = "x-aws-ec2-metadata-token"
)

type testType int

const (
	SecureTestType testType = iota
	InsecureTestType
	BadRequestTestType
	ServerErrorForTokenTestType
	retryTimeOutForTokenTestType
	retryTimeOutWith401TestType
)

type testServer struct {
	t *testing.T

	tokens      []string
	activeToken string
	data        string
}

type operationListProvider struct {
	operationsPerformed []string
}

func getTokenRequiredParams(t *testing.T, fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if e, a := "PUT", r.Method; e != a {
			t.Errorf("expect %v, http method got %v", e, a)
			http.Error(w, "wrong method", 400)
			return
		}
		if len(r.Header.Get(ttlHeader)) == 0 {
			t.Errorf("expect ttl header to be present in the request headers, got none")
			http.Error(w, "wrong method", 400)
			return
		}

		fn(w, r)
	}
}

func newTestServer(t *testing.T, testType testType, testServer *testServer) *httptest.Server {

	mux := http.NewServeMux()

	switch testType {
	case SecureTestType:
		mux.HandleFunc("/latest/api/token", getTokenRequiredParams(t, testServer.secureGetTokenHandler))
		mux.HandleFunc("/latest/", testServer.secureGetLatestHandler)
	case InsecureTestType:
		mux.HandleFunc("/latest/api/token", testServer.insecureGetTokenHandler)
		mux.HandleFunc("/latest/", testServer.insecureGetLatestHandler)
	case BadRequestTestType:
		mux.HandleFunc("/latest/api/token", getTokenRequiredParams(t, testServer.badRequestGetTokenHandler))
		mux.HandleFunc("/latest/", testServer.badRequestGetLatestHandler)
	case ServerErrorForTokenTestType:
		mux.HandleFunc("/latest/api/token", getTokenRequiredParams(t, testServer.serverErrorGetTokenHandler))
		mux.HandleFunc("/latest/", testServer.insecureGetLatestHandler)
	case retryTimeOutForTokenTestType:
		mux.HandleFunc("/latest/api/token", getTokenRequiredParams(t, testServer.retryableErrorGetTokenHandler))
		mux.HandleFunc("/latest/", testServer.insecureGetLatestHandler)
	case retryTimeOutWith401TestType:
		mux.HandleFunc("/latest/api/token", getTokenRequiredParams(t, testServer.retryableErrorGetTokenHandler))
		mux.HandleFunc("/latest/", testServer.unauthorizedGetLatestHandler)

	}

	return httptest.NewServer(mux)
}

func (s *testServer) secureGetTokenHandler(w http.ResponseWriter, r *http.Request) {

	token := s.tokens[0]

	// set the active token
	s.activeToken = token

	// rotate the token
	if len(s.tokens) > 1 {
		s.tokens = s.tokens[1:]
	}

	// set the header and response body
	w.Header().Set(ttlHeader, r.Header.Get(ttlHeader))
	w.Write([]byte(s.activeToken))
}

func (s *testServer) secureGetLatestHandler(w http.ResponseWriter, r *http.Request) {
	if len(s.activeToken) == 0 {
		s.t.Errorf("expect token to have been requested, was not")
		http.Error(w, "", 401)
		return
	}
	if e, a := s.activeToken, r.Header.Get(tokenHeader); e != a {
		s.t.Errorf("expect %v token, got %v", e, a)
		http.Error(w, "", 401)
		return
	}

	w.Header().Set("ttlheader", r.Header.Get(ttlHeader))
	w.Write([]byte(s.data))
}

func (s *testServer) insecureGetTokenHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", 404)
}

func (s *testServer) insecureGetLatestHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.Header.Get(tokenHeader)) != 0 {
		s.t.Errorf("Request token found, expected none")
		http.Error(w, "", 400)
		return
	}

	w.Write([]byte(s.data))
}

func (s *testServer) badRequestGetTokenHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", 400)
}

func (s *testServer) badRequestGetLatestHandler(w http.ResponseWriter, r *http.Request) {
	s.t.Errorf("Expected no call to this handler, incorrect behavior found")
}

func (s *testServer) serverErrorGetTokenHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", 403)
}

func (s *testServer) retryableErrorGetTokenHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "service unavailable", http.StatusServiceUnavailable)
}

func (s *testServer) unauthorizedGetLatestHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", 401)
}

func (opListProvider *operationListProvider) addToOperationPerformedList(r *request.Request) {
	opListProvider.operationsPerformed = append(opListProvider.operationsPerformed, r.Operation.Name)
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
		NewServer                   func(t *testing.T) *httptest.Server
		expectedData                string
		expectedError               string
		expectedOperationsPerformed []string
	}{
		"Insecure server success case": {
			NewServer: func(t *testing.T) *httptest.Server {
				testType := InsecureTestType
				Ts := &testServer{
					t:           t,
					tokens:      nil,
					activeToken: "",
					data:        "IMDSProfileForGoSDK",
				}
				return newTestServer(t, testType, Ts)
			},
			expectedData:                "IMDSProfileForGoSDK",
			expectedError:               "",
			expectedOperationsPerformed: []string{"GetToken", "GetMetadata", "GetToken", "GetMetadata", "GetToken", "GetMetadata", "GetToken", "GetMetadata"},
		},
		"Secure server success case": {
			NewServer: func(t *testing.T) *httptest.Server {
				testType := SecureTestType
				Ts := &testServer{
					t:           t,
					tokens:      []string{"firstToken", "secondToken", "thirdToken"},
					activeToken: "",
					data:        "IMDSProfileForGoSDK",
				}
				return newTestServer(t, testType, Ts)
			},
			expectedData:                "IMDSProfileForGoSDK",
			expectedError:               "",
			expectedOperationsPerformed: []string{"GetToken", "GetMetadata", "GetMetadata", "GetToken", "GetMetadata", "GetToken", "GetMetadata"},
		},
		"Bad request case": {
			NewServer: func(t *testing.T) *httptest.Server {
				testType := BadRequestTestType
				Ts := &testServer{
					t:           t,
					tokens:      []string{"firstToken", "secondToken", "thirdToken"},
					activeToken: "",
					data:        "IMDSProfileForGoSDK",
				}
				return newTestServer(t, testType, Ts)
			},
			expectedData:                "",
			expectedError:               "400",
			expectedOperationsPerformed: []string{"GetToken", "GetMetadata", "GetToken", "GetMetadata", "GetToken", "GetMetadata", "GetToken", "GetMetadata"},
		},
		"ServerErrorForTokenTestType": {
			NewServer: func(t *testing.T) *httptest.Server {
				testType := ServerErrorForTokenTestType
				Ts := &testServer{
					t:           t,
					tokens:      []string{},
					activeToken: "",
					data:        "IMDSProfileForGoSDK",
				}
				return newTestServer(t, testType, Ts)
			},
			expectedData:                "IMDSProfileForGoSDK",
			expectedError:               "",
			expectedOperationsPerformed: []string{"GetToken", "GetMetadata", "GetToken", "GetMetadata", "GetToken", "GetMetadata", "GetToken", "GetMetadata"},
		},
	}

	for name, x := range cases {
		t.Run(name, func(t *testing.T) {

			server := x.NewServer(t)
			defer server.Close()

			op := &operationListProvider{}

			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			c.Handlers.Complete.PushBack(op.addToOperationPerformedList)

			resp, err := c.GetMetadata("some/path")

			// token should stay alive, since default duration is 26000 seconds
			resp, err = c.GetMetadata("some/path")

			// Set Token TTL to 0, this will reset the token, thus there should be a change in token
			c.SetTokenTTL(0)
			resp, err = c.GetMetadata("some/path")

			// since earlier token ttl was set to 0. The ttl has expired and this should fetch new token
			resp, err = c.GetMetadata("some/path")

			if e, a := x.expectedData, resp; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}

			if err != nil && len(x.expectedError) == 0 {
				t.Errorf("expected no error, got %v", err)
			}

			if err == nil && len(x.expectedError) != 0 {
				t.Errorf("expected %v, got none", x.expectedError)
			}

			if err != nil && !strings.Contains(err.Error(), x.expectedError) {
				t.Errorf("expect %v, got %v", x.expectedError, err)
			}

			if len(x.expectedOperationsPerformed) != len(op.operationsPerformed) {
				t.Errorf("expected operation list to be %v, got %v", x.expectedOperationsPerformed, op.operationsPerformed)
				return
			}

			for i := 0; i < len(op.operationsPerformed); i++ {
				if op.operationsPerformed[i] != x.expectedOperationsPerformed[i] {
					t.Errorf("expected operation list to be %v, got %v", x.expectedOperationsPerformed, op.operationsPerformed)
				}
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
		NewServer                   func(t *testing.T) *httptest.Server
		expectedData                string
		expectedError               string
		expectedOperationsPerformed []string
	}{
		"Insecure server success case": {
			NewServer: func(t *testing.T) *httptest.Server {
				testType := InsecureTestType
				Ts := &testServer{
					t:           t,
					tokens:      nil,
					activeToken: "",
					data:        "us-west-2a",
				}
				return newTestServer(t, testType, Ts)
			},
			expectedData:                "us-west-2",
			expectedError:               "",
			expectedOperationsPerformed: []string{"GetToken", "GetMetadata"},
		},
		"Secure server success case": {
			NewServer: func(t *testing.T) *httptest.Server {
				testType := SecureTestType
				Ts := &testServer{
					t:           t,
					tokens:      []string{"firstToken", "secondToken", "thirdToken"},
					activeToken: "",
					data:        "us-west-2a",
				}
				return newTestServer(t, testType, Ts)
			},
			expectedData:                "us-west-2",
			expectedError:               "",
			expectedOperationsPerformed: []string{"GetToken", "GetMetadata"},
		},
		"Bad request case": {
			NewServer: func(t *testing.T) *httptest.Server {
				testType := BadRequestTestType
				Ts := &testServer{
					t:           t,
					tokens:      []string{"firstToken", "secondToken", "thirdToken"},
					activeToken: "",
					data:        "us-west-2a",
				}
				return newTestServer(t, testType, Ts)
			},
			expectedData:                "",
			expectedError:               "400",
			expectedOperationsPerformed: []string{"GetToken", "GetMetadata"},
		},
		"ServerErrorForTokenTestType": {
			NewServer: func(t *testing.T) *httptest.Server {
				testType := ServerErrorForTokenTestType
				Ts := &testServer{
					t:           t,
					tokens:      []string{},
					activeToken: "",
					data:        "us-west-2a",
				}
				return newTestServer(t, testType, Ts)
			},
			expectedData:                "us-west-2",
			expectedError:               "",
			expectedOperationsPerformed: []string{"GetToken", "GetMetadata"},
		},
	}

	for name, x := range cases {
		t.Run(name, func(t *testing.T) {

			server := x.NewServer(t)
			defer server.Close()

			op := &operationListProvider{}

			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			c.Handlers.Complete.PushBack(op.addToOperationPerformedList)

			resp, err := c.Region()

			if e, a := x.expectedData, resp; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}

			if err != nil && len(x.expectedError) == 0 {
				t.Errorf("expected no error, got %v", err)
			}

			if err == nil && len(x.expectedError) != 0 {
				t.Errorf("expected %v, got none", x.expectedError)
			}

			if err != nil && !strings.Contains(err.Error(), x.expectedError) {
				t.Errorf("expect %v, got %v", x.expectedError, err)
			}

			if len(x.expectedOperationsPerformed) != len(op.operationsPerformed) {
				t.Errorf("expected operation list to be %v, got %v", x.expectedOperationsPerformed, op.operationsPerformed)
				return
			}

			for i := 0; i < len(op.operationsPerformed); i++ {
				if op.operationsPerformed[i] != x.expectedOperationsPerformed[i] {
					t.Errorf("expected operation list to be %v, got %v", x.expectedOperationsPerformed, op.operationsPerformed)
				}
			}

		})
	}
}

func TestMetadataIAMInfo_success(t *testing.T) {
	cases := map[string]struct {
		NewServer                   func(t *testing.T) *httptest.Server
		expectedData                string
		expectedError               string
		expectedOperationsPerformed []string
	}{
		"Insecure server success case": {
			NewServer: func(t *testing.T) *httptest.Server {
				testType := InsecureTestType
				Ts := &testServer{
					t:           t,
					tokens:      nil,
					activeToken: "",
					data:        validIamInfo,
				}
				return newTestServer(t, testType, Ts)
			},
			expectedData:                validIamInfo,
			expectedError:               "",
			expectedOperationsPerformed: []string{"GetToken", "GetMetadata"},
		},
		"Secure server success case": {
			NewServer: func(t *testing.T) *httptest.Server {
				testType := SecureTestType
				Ts := &testServer{
					t:           t,
					tokens:      []string{"firstToken", "secondToken", "thirdToken"},
					activeToken: "",
					data:        validIamInfo,
				}
				return newTestServer(t, testType, Ts)
			},
			expectedData:                validIamInfo,
			expectedError:               "",
			expectedOperationsPerformed: []string{"GetToken", "GetMetadata"},
		},
	}

	for name, x := range cases {
		t.Run(name, func(t *testing.T) {

			server := x.NewServer(t)
			defer server.Close()

			op := &operationListProvider{}

			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			c.Handlers.Complete.PushBack(op.addToOperationPerformedList)

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

			if len(x.expectedOperationsPerformed) != len(op.operationsPerformed) {
				t.Errorf("expected operation list to be %v, got %v", x.expectedOperationsPerformed, op.operationsPerformed)
				return
			}

			for i := 0; i < len(op.operationsPerformed); i++ {
				if op.operationsPerformed[i] != x.expectedOperationsPerformed[i] {
					t.Errorf("expected operation list to be %v, got %v", x.expectedOperationsPerformed, op.operationsPerformed)
				}
			}

		})
	}
}

func TestMetadataIAMInfo_failure(t *testing.T) {
	cases := map[string]struct {
		NewServer                   func(t *testing.T) *httptest.Server
		expectedData                string
		expectedError               string
		expectedOperationsPerformed []string
	}{
		"Insecure server success case": {
			NewServer: func(t *testing.T) *httptest.Server {
				testType := InsecureTestType
				Ts := &testServer{
					t:           t,
					tokens:      nil,
					activeToken: "",
					data:        unsuccessfulIamInfo,
				}
				return newTestServer(t, testType, Ts)
			},
			expectedData:                unsuccessfulIamInfo,
			expectedError:               "",
			expectedOperationsPerformed: []string{"GetToken", "GetMetadata"},
		},
		"Secure server success case": {
			NewServer: func(t *testing.T) *httptest.Server {
				testType := SecureTestType
				Ts := &testServer{
					t:           t,
					tokens:      []string{"firstToken", "secondToken", "thirdToken"},
					activeToken: "",
					data:        unsuccessfulIamInfo,
				}
				return newTestServer(t, testType, Ts)
			},
			expectedData:                unsuccessfulIamInfo,
			expectedError:               "",
			expectedOperationsPerformed: []string{"GetToken", "GetMetadata"},
		},
	}

	for name, x := range cases {
		t.Run(name, func(t *testing.T) {

			server := x.NewServer(t)
			defer server.Close()

			op := &operationListProvider{}

			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			c.Handlers.Complete.PushBack(op.addToOperationPerformedList)

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

			if len(x.expectedOperationsPerformed) != len(op.operationsPerformed) {
				t.Errorf("expected operation list to be %v, got %v", x.expectedOperationsPerformed, op.operationsPerformed)
				return
			}

			for i := 0; i < len(op.operationsPerformed); i++ {
				if op.operationsPerformed[i] != x.expectedOperationsPerformed[i] {
					t.Errorf("expected operation list to be %v, got %v", x.expectedOperationsPerformed, op.operationsPerformed)
				}
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
		NewServer                   func(t *testing.T) *httptest.Server
		expectedData                string
		expectedError               string
		expectedOperationsPerformed []string
	}{
		"Insecure server success case": {
			NewServer: func(t *testing.T) *httptest.Server {
				testType := InsecureTestType
				Ts := &testServer{
					t:           t,
					tokens:      nil,
					activeToken: "",
					data:        instanceIdentityDocument,
				}
				return newTestServer(t, testType, Ts)
			},
			expectedData:                instanceIdentityDocument,
			expectedError:               "",
			expectedOperationsPerformed: []string{"GetToken", "GetDynamicData"},
		},
		"Secure server success case": {
			NewServer: func(t *testing.T) *httptest.Server {
				testType := SecureTestType
				Ts := &testServer{
					t:           t,
					tokens:      []string{"firstToken", "secondToken", "thirdToken"},
					activeToken: "",
					data:        instanceIdentityDocument,
				}
				return newTestServer(t, testType, Ts)
			},
			expectedData:                instanceIdentityDocument,
			expectedError:               "",
			expectedOperationsPerformed: []string{"GetToken", "GetDynamicData"},
		},
	}

	for name, x := range cases {
		t.Run(name, func(t *testing.T) {

			server := x.NewServer(t)
			defer server.Close()

			op := &operationListProvider{}

			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			c.Handlers.Complete.PushBack(op.addToOperationPerformedList)
			doc, err := c.GetInstanceIdentityDocument()

			if err != nil {
				t.Errorf("Expected no error, got %v ", err)
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

			if len(x.expectedOperationsPerformed) != len(op.operationsPerformed) {
				t.Errorf("expected operation list to be %v, got %v", x.expectedOperationsPerformed, op.operationsPerformed)
				return
			}

			for i := 0; i < len(op.operationsPerformed); i++ {
				if op.operationsPerformed[i] != x.expectedOperationsPerformed[i] {
					t.Errorf("expected operation list to be %v, got %v", x.expectedOperationsPerformed, op.operationsPerformed)
				}
			}

		})
	}
}

func TestEC2MetadataRetryFailure(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/latest/api/token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" && r.Header.Get(ttlHeader) != "" {
			w.Header().Set(ttlHeader, "200")
			http.Error(w, "service unavailable", http.StatusServiceUnavailable)
			return
		}
		http.Error(w, "bad request", http.StatusBadRequest)
	})

	// meta-data endpoint for this test, just returns the token
	mux.HandleFunc("/latest/meta-data/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("profile_name"))
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})

	c.Handlers.AfterRetry.PushBack(func(i *request.Request) {
		t.Logf("%v received, retrying operation %v", i.HTTPResponse.StatusCode, i.Operation.Name)
	})
	c.Handlers.Complete.PushBack(func(i *request.Request) {
		t.Logf("%v operation exited with status %v", i.Operation.Name, i.HTTPResponse.StatusCode)
	})

	resp, err := c.GetMetadata("some/path")
	resp, err = c.GetMetadata("some/path")

	if resp != "profile_name" {
		t.Errorf("Expected response to be profile_name, got %v", resp)
	}
	if err != nil {
		t.Errorf("Expected none, got error %v", err)
	}
}

func TestEC2MetadataRetryOnce(t *testing.T) {
	var secureDataFlow bool
	var retry = true
	mux := http.NewServeMux()

	mux.HandleFunc("/latest/api/token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" && r.Header.Get(ttlHeader) != "" {
			w.Header().Set(ttlHeader, "200")
			for retry {
				retry = false
				http.Error(w, "service unavailable", http.StatusServiceUnavailable)
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
		w.Write([]byte(r.Header.Get(tokenHeader)))
	})

	var tokenRetryCount int

	server := httptest.NewServer(mux)
	defer server.Close()
	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})

	// Handler on client that logs if retried
	c.Handlers.AfterRetry.PushBack(func(i *request.Request) {
		t.Logf("%v received, retrying operation %v", i.HTTPResponse.StatusCode, i.Operation.Name)
		tokenRetryCount++
	})

	_, err := c.GetMetadata("some/path")

	if tokenRetryCount != 1 {
		t.Errorf("Expected number of retries for fetching token to be 1, got %v", tokenRetryCount)
	}

	if !secureDataFlow {
		t.Errorf("Expected secure data flow to be %v, got %v", secureDataFlow, !secureDataFlow)
	}

	if err != nil {
		t.Errorf("Expected none, got error %v", err)
	}
}

func TestEC2Metadata_Concurrency(t *testing.T) {
	ts := &testServer{
		t:           t,
		tokens:      []string{"firstToken"},
		activeToken: "",
		data:        "IMDSProfileForSDKGo",
	}

	server := newTestServer(t, SecureTestType, ts)
	defer server.Close()

	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				resp, err := c.GetMetadata("some/data")
				if err != nil {
					t.Errorf("expect no error, got %v", err)
				}

				if e, a := "IMDSProfileForSDKGo", resp; e != a {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestRequestOnMetadata(t *testing.T) {
	ts := &testServer{
		t:           t,
		tokens:      []string{"firstToken", "secondToken"},
		activeToken: "",
		data:        "profile_name",
	}
	server := newTestServer(t, SecureTestType, ts)
	defer server.Close()

	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
	req := c.NewRequest(&request.Operation{
		Name:            "Ec2Metadata request",
		HTTPMethod:      "GET",
		HTTPPath:        "/latest",
		Paginator:       nil,
		BeforePresignFn: nil,
	}, nil, nil)

	op := operationListProvider{}
	c.Handlers.Complete.PushBack(op.addToOperationPerformedList)
	err := req.Send()

	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	if len(op.operationsPerformed) < 1 {
		t.Errorf("Expected atleast one operation GetToken to be called on EC2Metadata client")
		return
	}

	if op.operationsPerformed[0] != "GetToken" {
		t.Errorf("Expected GetToken operation to be called")
	}

}

func TestExhaustiveRetryToFetchToken(t *testing.T) {
	ts := &testServer{
		t:           t,
		tokens:      []string{"firstToken", "secondToken"},
		activeToken: "",
		data:        "IMDSProfileForSDKGo",
	}

	server := newTestServer(t, retryTimeOutForTokenTestType, ts)
	defer server.Close()

	op := &operationListProvider{}

	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
	c.Handlers.Complete.PushBack(op.addToOperationPerformedList)

	c.Handlers.AfterRetry.PushBack(func(i *request.Request) {
		t.Logf("Retried %v operation for status code %v", i.Operation.Name, i.HTTPResponse.StatusCode)
	})

	c.Handlers.Complete.PushBack(func(i *request.Request) {
		t.Logf("Operation %v completed, with status %v", i.Operation.Name, i.HTTPResponse.StatusCode)
	})

	resp, err := c.GetMetadata("/some/path")
	resp, err = c.GetMetadata("/some/path")
	resp, err = c.GetMetadata("/some/path")
	resp, err = c.GetMetadata("/some/path")

	expectedOperationsPerformed := []string{"GetToken", "GetMetadata", "GetMetadata", "GetMetadata", "GetMetadata"}

	if e, a := "IMDSProfileForSDKGo", resp; e != a {
		t.Errorf("Expected %v, got %v", e, a)
	}

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(expectedOperationsPerformed) != len(op.operationsPerformed) {
		t.Errorf("expected operation list to be %v, got %v", expectedOperationsPerformed, op.operationsPerformed)
		return
	}

	for i := 0; i < len(op.operationsPerformed); i++ {
		if op.operationsPerformed[i] != expectedOperationsPerformed[i] {
			t.Errorf("expected operation list to be %v, got %v", expectedOperationsPerformed, op.operationsPerformed)
		}
	}

}

func TestExhaustiveRetryWith401(t *testing.T) {
	ts := &testServer{
		t:           t,
		tokens:      []string{"firstToken", "secondToken"},
		activeToken: "",
		data:        "IMDSProfileForSDKGo",
	}

	server := newTestServer(t, retryTimeOutWith401TestType, ts)
	defer server.Close()

	op := &operationListProvider{}

	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
	c.Handlers.Complete.PushBack(op.addToOperationPerformedList)

	resp, err := c.GetMetadata("/some/path")
	resp, err = c.GetMetadata("/some/path")
	resp, err = c.GetMetadata("/some/path")
	resp, err = c.GetMetadata("/some/path")

	expectedOperationsPerformed := []string{"GetToken", "GetMetadata", "GetToken", "GetMetadata", "GetToken", "GetMetadata", "GetToken", "GetMetadata"}

	if e, a := "", resp; e != a {
		t.Errorf("Expected %v, got %v", e, a)
	}

	if err == nil {
		t.Errorf("Expected %v error, got none", err)
	}

	if len(expectedOperationsPerformed) != len(op.operationsPerformed) {
		t.Errorf("expected operation list to be %v, got %v", expectedOperationsPerformed, op.operationsPerformed)
		return
	}

	for i := 0; i < len(op.operationsPerformed); i++ {
		if op.operationsPerformed[i] != expectedOperationsPerformed[i] {
			t.Errorf("expected operation list to be %v, got %v", expectedOperationsPerformed, op.operationsPerformed)
		}
	}

}
