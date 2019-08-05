package awstesting

import (
	"bytes"
	"encoding/base64"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/util"
	"github.com/google/go-cmp/cmp"

	"io/ioutil"
	"net/url"
	"testing"
)

func AssertRequestMethodEquals(t *testing.T, req *request.Request, val string) bool {
	return req.HTTPRequest.Method == val
}

func AssertRequestUrlMatches(t *testing.T, req *request.Request, val string) bool {
	return AssertURL(t, val, req.HTTPRequest.URL.String())
}

func AssertRequestUrlPathMatches(t *testing.T, req *request.Request, val string) bool {
	return req.HTTPRequest.URL.RequestURI() == val
}

func AssertRequestUrlQueryMatches(t *testing.T, req *request.Request, val string) bool {
	structExpect, err := url.Parse(val) // parsed val into a structure
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	queryRequest := req.HTTPRequest.URL.Query() //parsed RawQuery of "req" to get the values inside
	queryExpect := structExpect.Query()         //parsed RawQuery of "val" to get the values inside

	return queryRequest.Encode() == queryExpect.Encode()
}

func AssertRequestHeadersMatch(t *testing.T, req *request.Request, header map[string]interface{}) bool {
	for key, valExpect := range header {
		valReq := req.HTTPRequest.Header.Get(key)
		if valReq == "" || valReq[0] != valExpect {
			return false
		}
	}
	return true
}

func AssertRequestBodyEqualsBytes(t *testing.T, req *request.Request, val string) bool {
	var bytesReqBody []byte
	bytesExpect, err := base64.StdEncoding.DecodeString(val)

	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	if req.HTTPRequest.Body != nil {
		bytesReqBody, err = ioutil.ReadAll(req.HTTPRequest.Body)
		if err != nil {
			t.Errorf("expect no error, got %v", err)
		}
	}

	req.HTTPRequest.Body = ioutil.NopCloser(bytes.NewBuffer(bytesReqBody))

	return bytes.Compare(bytesReqBody, bytesExpect) == 0
}

func AssertRequestBodyEqualsJson(t *testing.T, req *request.Request, val string) bool {
	var bytesReqBody []byte
	var err error
	if req.HTTPRequest.Body != nil {
		bytesReqBody, err = ioutil.ReadAll(req.HTTPRequest.Body)
		if err != nil {
			t.Errorf("expect no error, got %v", err)
		}
	}

	req.HTTPRequest.Body = ioutil.NopCloser(bytes.NewBuffer(bytesReqBody))

	return AssertJSON(t, val, util.Trim(string(bytesReqBody)))
}

func AssertRequestBodyMatchesXml(t *testing.T, req *request.Request, val string, container interface{}) bool {
	r := req.HTTPRequest

	if r.Body == nil {
		t.Errorf("expect body not to be nil")
	}
	body := util.SortXML(r.Body)

	return AssertXML(t, val, util.Trim(string(body)), container)
}

func AssertRequestBodyEqualsString(t *testing.T, req *request.Request, val string) bool {
	var bytesReqBody []byte
	var err error
	if req.HTTPRequest.Body != nil {
		bytesReqBody, err = ioutil.ReadAll(req.HTTPRequest.Body)
		if err != nil {
			t.Errorf("expect no error, got %v", err)
		}
	}

	req.HTTPRequest.Body = ioutil.NopCloser(bytes.NewBuffer(bytesReqBody))
	stringReqBody := string(bytesReqBody)

	return stringReqBody == val
}

func AssertRequestIdEquals(t *testing.T, req *request.Request, val string) bool {
	return req.RequestID == val
}

func AssertResponseDataEquals(t *testing.T, response interface{}, expectResponse interface{}) bool {
	if response == nil || expectResponse == nil {
		return response == expectResponse
	}
	return cmp.Equal(expectResponse, response)
}

func AssertResponseErrorIsKindOf(t *testing.T, err error, val string) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		return awsErr.Code() == val
	}
	return true
}

func AssertResponseErrorMessageEquals(t *testing.T, err error, val string) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		return awsErr.Message() == val
	}
	return true
}

func AssertResponseErrorDataEquals(t *testing.T, err error, val map[string]interface{}) {
	if testing.Short() {
		t.Skip("skipping responseErrorDataEquals assertion")
	}
}
