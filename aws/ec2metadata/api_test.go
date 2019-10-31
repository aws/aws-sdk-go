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

type testServer struct {
	t             *testing.T
	serverType    serverType
	returnedToken []string
	httpMethod    string
	resp          string
}

type serverType int

const (
	insecure serverType = 0
	secure   serverType = 1
)

func initServer(s testServer) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/latest/api/token", s.tokenHandler)
	mux.HandleFunc("/latest/", s.metadataHandler)
	return httptest.NewServer(mux)
}

func (s testServer) tokenHandler(w http.ResponseWriter, r *http.Request) {
	switch s.serverType {
	case secure:
		if r.Method == "PUT" && r.Header.Get(ec2metadata.TTLHeader) != "" {
			w.Header().Set(ec2metadata.TTLHeader, r.Header.Get(ec2metadata.TTLHeader))
			if len(s.returnedToken)!=0{
				w.Write([]byte(s.returnedToken[0]))
			}
			return
		}
		http.Error(w, "bad request", http.StatusBadRequest)
	default:
		http.Error(w, "not found", http.StatusNotFound)
	}
}

func (s testServer) metadataHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != s.httpMethod {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	switch s.serverType {
	case secure:
		if r.Header.Get(ec2metadata.TokenHeader) == "token" {
			w.Write([]byte(s.resp))
			return
		}
		http.Error(w, "Unauthorized for an expired, invalid, or missing token header", http.StatusUnauthorized)
	default: // insecure
		w.Write([]byte(s.resp))
	}
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
		config           testServer
		expectedResponse string
		expectedError    string
	}{
		"Insecure server success case": {
			testServer{
				t:             t,
				returnedToken: nil,
				httpMethod:    "GET",
				resp:          "profile_name",
			},
			"profile_name",
			"",
		},
		"Secure server success case": {
			testServer{
				serverType:    secure,
				returnedToken: []string{"token"},
				httpMethod:    "GET",
				resp:          "profile_name", // real response includes suffix ,
			},
			"profile_name",
			"",
		},
		"Insecure server failure case": {
			testServer{
				returnedToken: nil,
				httpMethod:    "PUT",
				resp:          "",
			},
			"",
			"bad request",
		},
		"Secure server failure case": {
			testServer{
				serverType:    secure,
				returnedToken: []string{"invalid token"},
				httpMethod:    "GET",
				resp:          "", // real response includes suffix ,
			},
			"",
			"Unauthorized",
		},
	}

	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			server := initServer(in.config)
			defer server.Close()
			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			resp, err := c.GetMetadata("some/path")

			if err == nil && in.expectedError != "" {
				t.Errorf("expect %v, got %v", in.expectedError, err)
			}

			if e, a := in.expectedError, err; a != nil {
				if !strings.Contains(a.Error(), e) {
					t.Errorf("expect %v, got %v", in.expectedError, err)
				}
			}

			if e, a := in.expectedResponse, resp; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

func TestGetUserData(t *testing.T) {
	cases := map[string]struct {
		config           testServer
		expectedResponse string
		expectedError    string
	}{
		"Insecure server successful case": {
			testServer{
				serverType:    insecure,
				returnedToken: nil,
				httpMethod:    "GET",
				resp:          "user-data",
			},
			"user-data",
			"",
		},

		"Secure server successful case": {
			testServer{
				serverType:    secure,
				returnedToken: []string{"token"},
				httpMethod:    "GET",
				resp:          "user-data",
			},
			"user-data",
			"",
		},
		"Insecure server failure case": {
			testServer{
				serverType:    insecure,
				returnedToken: nil,
				httpMethod:    "PUT",
				resp:          "user-data",
			},
			"",
			"bad request",
		},

		"Secure server failure case": {
			testServer{
				serverType:    secure,
				returnedToken: nil,
				httpMethod:    "GET",
				resp:          "user-data",
			},
			"",
			"Unauthorized",
		},
	}

	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			server := initServer(in.config)
			defer server.Close()
			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			resp, err := c.GetUserData()

			if err == nil && in.expectedError != "" {
				t.Errorf("expect nil, got %v", in.expectedError)
			}

			if e, a := in.expectedError, err; a != nil {
				if !strings.Contains(a.Error(), e) {
					t.Errorf("expect %v, got %v", in.expectedError, err)
				}
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

	if requestFailedError, ok := err.(awserr.RequestFailure); ok {
		if e, a := http.StatusNotFound, requestFailedError.StatusCode(); e != a {
			t.Errorf("expect %v, got %v", e, a)
		}
	}
}

func TestGetRegion(t *testing.T) {
	cases := map[string]struct {
		config           testServer
		expectedResponse string
		expectedError    string
	}{
		"Insecure server successful case": {
			testServer{
				serverType:    insecure,
				returnedToken: nil,
				httpMethod:    "GET",
				resp:          "us-west-2a",
			},
			"us-west-2",
			"",
		},

		"Secure server successful case": {
			testServer{
				serverType:    secure,
				returnedToken: []string{"token"},
				httpMethod:    "GET",
				resp:          "us-west-2a",
			},
			"us-west-2",
			"",
		},
		"Insecure server failure case": {
			testServer{
				serverType:    insecure,
				returnedToken: nil,
				httpMethod:    "PUT",
				resp:          "us-west-2a",
			},
			"",
			"bad request",
		},

		"Secure server failure case": {
			testServer{
				serverType:    secure,
				returnedToken: nil,
				httpMethod:    "GET",
				resp:          "us-west-2a",
			},
			"",
			"Unauthorized",
		},
	}

	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			server := initServer(in.config)
			defer server.Close()
			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			region, err := c.Region()

			if err == nil && in.expectedError != "" {
				t.Errorf("expect nil, got %v", in.expectedError)
			}

			if e, a := in.expectedError, err; a != nil {
				if !strings.Contains(a.Error(), e) {
					t.Errorf("expect %v, got %v", in.expectedError, err)
				}
			}

			if e, a := in.expectedResponse, region; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

func TestMetadataAvailable(t *testing.T) {
	cases := map[string]struct {
		config           testServer
		expectedResponse bool
		expectedError    string
	}{
		"Insecure server successful case": {
			testServer{
				serverType:    insecure,
				returnedToken: nil,
				httpMethod:    "GET",
				resp:          "instance-id",
			},
			true,
			"",
		},

		"Secure server successful case": {
			testServer{
				serverType:    secure,
				returnedToken: []string{"token"},
				httpMethod:    "GET",
				resp:          "instance-id",
			},
			true,
			"",
		},
		"Insecure server failure case": {
			testServer{
				serverType:    insecure,
				returnedToken: nil,
				httpMethod:    "PUT",
				resp:          "",
			},
			false,
			"bad request",
		},

		"Secure server failure case": {
			testServer{
				serverType:    secure,
				returnedToken: nil,
				httpMethod:    "GET",
				resp:          "",
			},
			false,
			"Unauthorized",
		},
	}

	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			server := initServer(in.config)
			defer server.Close()
			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			if e, a := in.expectedResponse, c.Available(); e != a {
				t.Errorf("expect availability as %v, got %v", e, a)
			}
		})
	}
}

func TestMetadataIAMInfo_success(t *testing.T) {
	cases := map[string]struct {
		config        testServer
		expectedError string
	}{
		"Insecure server successful case": {
			testServer{
				serverType:    insecure,
				returnedToken: nil,
				httpMethod:    "GET",
				resp:          validIamInfo,
			},
			"",
		},
		"Secure server successful case": {
			testServer{
				serverType:    secure,
				returnedToken: []string{"token"},

				httpMethod: "GET",
				resp:       validIamInfo,
			},
			"",
		},
	}

	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			server := initServer(in.config)
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
		config        testServer
		expectedError string
	}{
		"Insecure server failure case": {
			testServer{
				serverType:    insecure,
				returnedToken: nil,
				httpMethod:    "PUT",
				resp:          unsuccessfulIamInfo,
			},
			"bad request",
		},

		"Secure server failure case": {
			testServer{
				serverType:    secure,
				returnedToken: nil,
				httpMethod:    "GET",
				resp:          unsuccessfulIamInfo,
			},
			"Unauthorized",
		},
	}
	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			server := initServer(in.config)
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

//
func TestEC2RoleProviderInstanceIdentity(t *testing.T) {
	cases := map[string]struct {
		config           testServer
		expectedResponse string
		expectedError    string
	}{
		"Insecure server successful case": {
			testServer{
				serverType:    insecure,
				returnedToken: nil,
				httpMethod:    "GET",
				resp:          instanceIdentityDocument,
			},
			instanceIdentityDocument,
			"",
		},

		"Secure server successful case": {
			testServer{
				serverType:    secure,
				returnedToken: []string{"token"},
				httpMethod:    "GET",
				resp:          instanceIdentityDocument,
			},
			instanceIdentityDocument,
			"",
		},
	}

	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			server := initServer(in.config)
			defer server.Close()
			c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
			doc, err := c.GetInstanceIdentityDocument()

			if err == nil && in.expectedError != "" {
				t.Errorf("expect nil, got %v", in.expectedError)
			}

			if e, a := in.expectedError, err; a != nil {
				if !strings.Contains(a.Error(), e) {
					t.Errorf("expect %v, got %v", in.expectedError, err)
				}
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

func TestEC2MetadataRandomRetries(t *testing.T) {
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
