package awstesting

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/util"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/url"
	"reflect"
	"testing"
)


func FloatIntEquate() cmp.Option {
	return cmp.Options{
		cmp.FilterValues(areNaNsF64s, cmp.Comparer(equateAlways)),
		cmp.FilterValues(areNaNsF32s, cmp.Comparer(equateAlways)),
		cmp.FilterValues(areNaNsI32s, cmp.Comparer(equateAlways)),
		cmp.FilterValues(areNaNsI64s, cmp.Comparer(equateAlways)),
	}
}

func equateAlways(_, _ interface{}) bool { return true }

func areNaNsF64s(x, y float64) bool {
	return math.IsNaN(x) && math.IsNaN(y)
}
func areNaNsF32s(x, y float32) bool {
	return areNaNsF64s(float64(x), float64(y))
}

func areNaNsI32s(x, y int32) bool {
	return areNaNsF64s(float64(x), float64(y))
}
func areNaNsI64s(x, y int64) bool {
	return areNaNsF64s(float64(x), float64(y))
}

func EquateIoReader() cmp.Option {
	return cmp.FilterValues(ioReaderCompare, cmp.Comparer(equateAlways))
}

func ioReaderCompare(x, y interface{}) bool {
	vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
	if vx.Type().String() != "ioutil.nopCloser" || vy.Type().String() != "aws.ReaderSeekerCloser" {
		return false
	}
	xbyte, err1 := ioutil.ReadAll(x.(io.Reader))
	if err1 != nil {
		log.Println("couldn't read the body from response")
	}
	ybyte, err2 := ioutil.ReadAll(y.(io.Reader))
	if err2 != nil {
		log.Println("couldn't read the body from expect response")
	}
	return (x != nil && y != nil) && (bytes.Compare(xbyte, ybyte) == 0)
}

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
		reqVal := queryRequest.Get(expectKey)
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
		if key == "Header-Json-Value" {
			expectJsonValue, err1 := protocol.DecodeJSONValue(valExpect.(string), protocol.Base64Escape)
			if err1 != nil {
				t.Errorf(errMsg("unable to parse expected JSON", err1, msgAndArgs...))
			}
			responseJsonValue, err2 := protocol.DecodeJSONValue(valReq, protocol.Base64Escape)
			if err2 != nil {
				t.Errorf(errMsg("unable to parse response JSON", err2, msgAndArgs...))
			}
			for key1, val1 := range expectJsonValue{
				if !cmp.Equal(responseJsonValue[key1], val1, FloatIntEquate()) {
					t.Errorf(errMsg("aws.JSON value from expect and response don't match", nil))
					return false
				}
			}
			continue
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
	return cmp.Equal(response, expectResponse, cmpopts.EquateEmpty(), EquateIoReader())
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
