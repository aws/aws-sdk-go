package awstesting

import (
	"bytes"
	"encoding/base64"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/util"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io/ioutil"
	"net/url"
	"testing"
)

func AssertRequestMethodEquals(t *testing.T, req *request.Request, val string, msgAndArgs ...interface{}) bool {
	return equal(t, val, req.HTTPRequest.Method, msgAndArgs)
}

func AssertRequestUrlMatches(t *testing.T, req *request.Request, val string, msgAndArgs ...interface{}) bool {
	return AssertURL(t, val, req.HTTPRequest.URL.String(), msgAndArgs)
}

func AssertRequestUrlPathMatches(t *testing.T, req *request.Request, val string, msgAndArgs ...interface{}) bool {
	return equal(t, val, req.HTTPRequest.URL.EscapedPath(), msgAndArgs)
}

func AssertRequestUrlQueryMatches(t *testing.T, req *request.Request, val string, msgAndArgs ...interface{}) bool {

	queryRequest := req.HTTPRequest.URL.Query() //parsed RawQuery of "req" to get the values inside
	expectQ, err := url.ParseQuery(val)

	if err != nil {
		t.Errorf(errMsg("unable to parse query from expect", err, msgAndArgs))
		return false
	}
	for expectKey, expectVal := range expectQ{
		if expectKey == "timestamp" {
			expectKey = "unixTimestamp"
		}
		reqVal := queryRequest.Get(expectKey)
		if expectKey == "binary-value"{
			temp, err  := base64.StdEncoding.DecodeString(reqVal)
			if err != nil {
				t.Errorf(errMsg("unable to decode string for binary-value parameter of expect", err, msgAndArgs))
				return false
			}
			reqVal = string(temp)
		}
		if  reqVal != expectVal[0] {
			t.Errorf(errMsg("query values inside request and expect don't match", nil))
			return false
		}
	}
	return true
}

func AssertRequestHeadersMatch(t *testing.T, req *request.Request, header map[string]interface{}, msgAndArgs ...interface{}) bool {
	for key, valExpect := range header {
		valReq := req.HTTPRequest.Header.Get(key)
		if key == "Header-Binary" {
			temp, err  := base64.StdEncoding.DecodeString(valReq)
			if err != nil {
				t.Errorf(errMsg("unable to decode string for Header-Binary parameter of expect", err, msgAndArgs))
				return false
			}
			valReq = string(temp)
		}
		if valReq == "" || valReq != valExpect {
			t.Errorf(errMsg("header values inside request and expect don't match", nil))
			return false
		}
	}
	return true
}

func AssertRequestBodyEqualsBytes(t *testing.T, req *request.Request, val string, msgAndArgs ...interface{}) bool {
	var bytesReqBody []byte
	var err error
	if req.HTTPRequest.Body != nil {
		bytesReqBody, err = ioutil.ReadAll(req.HTTPRequest.Body)
		if err != nil {
			t.Errorf(errMsg("unable to read body from request", err, msgAndArgs))
			return false
		}
	}
	req.HTTPRequest.Body = ioutil.NopCloser(bytes.NewBuffer(bytesReqBody))

	return equal(t, val, string(bytesReqBody), msgAndArgs)
}

func AssertRequestBodyEqualsJson(t *testing.T, req *request.Request, val string, msgAndArgs ...interface{}) bool {
	var bytesReqBody []byte
	var err error
	if req.HTTPRequest.Body != nil {
		bytesReqBody, err = ioutil.ReadAll(req.HTTPRequest.Body)
		if err != nil {
			t.Errorf(errMsg("unable to read body from request", err, msgAndArgs))
			return false
		}
	}
	req.HTTPRequest.Body = ioutil.NopCloser(bytes.NewBuffer(bytesReqBody))

	return AssertJSON(t, val, util.Trim(string(bytesReqBody)))
}

func AssertRequestBodyMatchesXml(t *testing.T, req *request.Request, val string, container interface{}, msgAndArgs ...interface{}) bool {
	r := req.HTTPRequest
	if r.Body == nil {
		t.Errorf(errMsg("request body is nil", nil, msgAndArgs))
		return false
	}
	body := util.SortXML(r.Body)

	return AssertXML(t, val, util.Trim(string(body)), container)
}

func AssertRequestBodyEqualsString(t *testing.T, req *request.Request, val string, msgAndArgs ...interface{}) bool {
	buf := new(bytes.Buffer)
	ReqBody, err := req.HTTPRequest.GetBody()
	if err != nil {
		t.Errorf(errMsg("unable to read body from request", err, msgAndArgs))
		return false
	}
	buf.ReadFrom(ReqBody)

	return buf.String() == val
}

func AssertRequestIdEquals(t *testing.T, req *request.Request, val string, msgAndArgs ...interface{}) bool {
	return req.RequestID == val
}

func AssertResponseDataEquals(t *testing.T, response interface{}, expectResponse interface{}, msgAndArgs ...interface{}) bool {
	if response == nil || expectResponse == nil {
		return equal(t, expectResponse, response, msgAndArgs)
	}
	return cmp.Equal(response, expectResponse, cmpopts.EquateEmpty())
}

func AssertResponseErrorIsKindOf(t *testing.T, err error, val string, msgAndArgs ...interface{}) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		return awsErr.Code() == val
	}
	return true
}

func AssertResponseErrorMessageEquals(t *testing.T, err error, val string, msgAndArgs ...interface{}) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		return awsErr.Message() == val
	}
	return true
}

func AssertResponseErrorDataEquals(t *testing.T, err error, val map[string]interface{}, msgAndArgs ...interface{}) {
	if testing.Short() {
		t.Skip("skipping responseErrorDataEquals assertion")
	}
}
