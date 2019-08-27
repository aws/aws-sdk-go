package awstesting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/util"
	"github.com/google/go-cmp/cmp"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/url"
	"reflect"
	"strings"
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
		log.Println("couldn't read the body from actual response")
	}
	ybyte, err2 := ioutil.ReadAll(y.(io.Reader))
	if err2 != nil {
		log.Println("couldn't read the body from expect response")
	}
	return (x != nil && y != nil) && (bytes.Compare(xbyte, ybyte) == 0)
}

// DiffReporter is a simple custom reporter that only records differences
// detected during comparison.
type DiffReporter struct {
	path  cmp.Path
	diffs []string
}

func (r *DiffReporter) PushStep(ps cmp.PathStep) {
	r.path = append(r.path, ps)
}

func (r *DiffReporter) Report(rs cmp.Result) {
	if !rs.Equal() {
		vx, vy := r.path.Last().Values()
		r.diffs = append(r.diffs, fmt.Sprintf("comparision failed at %#v:\n\t expect: %+v\n\t actual: %+v\n", r.path, vx, vy))
	}
}

func (r *DiffReporter) PopStep() {
	r.path = r.path[:len(r.path)-1]
}

func (r *DiffReporter) String() string {
	return strings.Join(r.diffs, "\n")
}

// StringEqual asserts that two strings are equal else returns false by wrapping an error message
func StringEqual(t *testing.T, expectVal, actualVal string) bool {
	if expectVal != actualVal {
		t.Errorf("%s\n",fmt.Sprintf("String comparision failed,\n\texpect: %s\n\tactual: %s\n", expectVal, actualVal))
		return false
	}
	return true
}

// 	Asserts if method field in request and expect value are equal
func AssertRequestMethodEquals(t *testing.T, expectVal string, actualVal string) bool {
	return StringEqual(t, expectVal, actualVal)
}

// Asserts if request URL in request and expect are equal. True
// if all URL fields, (path, query, hostname, protocol name etc)
// in request and expect are equal
func AssertRequestUrlMatches(t *testing.T, expectVal string, actualVal string) bool {
	return AssertURL(t, expectVal, actualVal)
}

// Asserts if the path field in request and expect are equal
func AssertRequestUrlPathMatches(t *testing.T, expectVal string, actualVal string) bool {
	return StringEqual(t, expectVal, actualVal)
}

// 	Asserts if query values in request and expect are equal. Values
// 	of query string in request and expect are equal even if they have
//	different orders
func AssertRequestUrlQueryMatches(t *testing.T, expectVal string, req *request.Request, msgAndArgs ...interface{}) bool {
	queryRequest := req.HTTPRequest.URL.Query() //parsed RawQuery of "req" to get the values inside
	expectQ, err := url.ParseQuery(expectVal)

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

// Asserts if headers in request and expect are equal. True if each
// header key in the expected map is present in the request with
// equal values. request may have additional headers outside the
// expected ones.
func AssertRequestHeadersMatch(t *testing.T, expectHeader map[string]interface{}, req *request.Request, msgAndArgs ...interface{}) bool {
	for key, valExpect := range expectHeader {
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

			var r DiffReporter
			for key1, val1 := range expectJsonValue {
				if !cmp.Equal(responseJsonValue[key1], val1, FloatIntEquate(), cmp.Reporter(&r)) {
					t.Errorf(errMsg("aws.JSON value from expect and response don't match", nil))
					return false
				}
			}
			continue
		}
		if valReq == "" || valReq != valExpect {
			t.Errorf("header values don't match,\nexpect: %s\nactual: %s", valExpect, valReq )
			return false
		}
	}
	return true
}

// Asserts if base64 encoded string inside request body is equal to expected value.
func AssertRequestBodyEqualsBytes(t *testing.T, expectVal string, req *request.Request) bool {
	var bytesReqBody []byte
	var err error
	if req.HTTPRequest.Body != nil {
		bytesReqBody, err = ioutil.ReadAll(req.HTTPRequest.Body)
		if err != nil {
			t.Errorf(errMsg("unable to read body from request", err))
			return false
		}
	}
	req.HTTPRequest.Body = ioutil.NopCloser(bytes.NewBuffer(bytesReqBody))

	return StringEqual(t, expectVal, string(bytesReqBody))
}

// AssertRequestBodyEqualsJson verifies that the json value in request body
// string matches the expectVal map
func AssertRequestBodyEqualsJson(t *testing.T, expectVal map[string]interface{}, req *request.Request, msgAndArgs ...interface{}) bool {
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

	actualVal := map[string]interface{}{}
	if err := json.Unmarshal(bytesReqBody, &actualVal); err != nil {
		t.Errorf(errMsg("unable to parse expected JSON", err, msgAndArgs...))
		return false
	}
	var r DiffReporter
	if !cmp.Equal(actualVal, expectVal, FloatIntEquate(), cmp.Reporter(&r)){
		fmt.Print(r.String())
		return false
	}
 	return true
}

// AssertRequestBodyMatchesXml verifies that the expect xml string matches the actual. True if
// XML string inside request and expect are equal. For XML string, order of
// elements and attributes are significant while whitespaces are not.
func AssertRequestBodyMatchesXml(t *testing.T, expectVal string, req *request.Request, container interface{}, msgAndArgs ...interface{}) bool {
	r := req.HTTPRequest
	if r.Body == nil {
		t.Errorf(errMsg("request body is nil", nil, msgAndArgs))
		return false
	}
	body := util.SortXML(r.Body)

	return AssertXML(t, expectVal, util.Trim(string(body)), container)
}

func AssertRequestBodyEqualsString(t *testing.T, expectVal string, req *request.Request, msgAndArgs ...interface{}) bool {
	var bytesReqBody []byte
	var err error
	if req.HTTPRequest.Body != nil {
		bytesReqBody, err = ioutil.ReadAll(req.HTTPRequest.Body)
		if err != nil {
			t.Errorf(errMsg("unable to read body from request", err))
			return false
		}
	}
	req.HTTPRequest.Body = ioutil.NopCloser(bytes.NewBuffer(bytesReqBody))

	return StringEqual(t, expectVal, string(bytesReqBody))
}

// Asserts if requestID field in request and expect value are equal
func AssertRequestIdEquals(t *testing.T, expectVal string, actualVal string) bool {
	return StringEqual(t, expectVal, actualVal)
}

// Asserts if data in response and error are equal. True if all fields
// inside the structure parsed from expect value are equal to the
// corresponding response fields
func AssertResponseDataEquals(t *testing.T, expectResponse interface{}, actualResponse interface{}, msgAndArgs ...interface{}) bool {
	if actualResponse == nil || expectResponse == nil {
		return equal(t, expectResponse, actualResponse, msgAndArgs)
	}
	var r DiffReporter
	if !cmp.Equal(actualResponse, expectResponse, EquateIoReader(),cmp.Reporter(&r)){
		fmt.Print(r.String())
		return false
	}
	return true
}

// Asserts if code in response error and expect value are equal
func AssertResponseErrorIsKindOf(t *testing.T, expectVal string, err error) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		return StringEqual(t, expectVal, awsErr.Code())
	}
	return true
}

// Asserts if error message in response error and expect value are equal
func AssertResponseErrorMessageEquals(t *testing.T, expectVal string, err error) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		return StringEqual(t, expectVal, awsErr.Message())
	}
	return true
}

// Asserts if requestID field inside response error and expect
//	value are equal
func AssertResponseErrorRequestIdEquals(t *testing.T, expectVal string, err error) bool{
	if reqErr, ok := err.(awserr.RequestFailure); ok{
		return StringEqual(t, expectVal, reqErr.RequestID())
	}
	return true
}

// Asserts if all fields inside the structure parsed from expect value
// are equal to the corresponding response error data fields. This is not
// implemented in because Go SDK V1 doesn't expose the error data
func AssertResponseErrorDataEquals(t *testing.T, expectVal map[string]interface{}, err error, msgAndArgs ...interface{}){
	if testing.Short() {
		t.Skip("\n\nskipping responseErrorDataEquals assertion")
	}
}
