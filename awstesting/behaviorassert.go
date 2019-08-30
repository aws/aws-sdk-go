// +build go1.9

package awstesting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/util"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// DiffReporter is a simple custom reporter that only records differences
// detected during comparison.
// Copied from go-cmp library "example_reporter_test.go"
type DiffReporter struct {
	path  cmp.Path
	diffs []string
}

// PushStep appends path to path step , used by DiffReporter
func (r *DiffReporter) PushStep(ps cmp.PathStep) {
	r.path = append(r.path, ps)
}

// Report constructs the error report and stores in r
func (r *DiffReporter) Report(rs cmp.Result) {
	if !rs.Equal() {
		vx, vy := r.path.Last().Values()
		r.diffs = append(r.diffs, fmt.Sprintf("comparision failed for key: %#v\n\t expect: %+v\n\t actual: %+v\n", r.path, vx, vy))
	}
}

// PopStep truncates path, used by DiffReporter
func (r *DiffReporter) PopStep() {
	r.path = r.path[:len(r.path)-1]
}

// String() returns the differences in string format
func (r *DiffReporter) String() string {
	return strings.Join(r.diffs, "\n")
}

// FloatIntEquate returns an option which compares floats with ints
// and return true if they have same value, Ex: (1.00 , 1)  are equal
func FloatIntEquate() cmp.Option {
	return cmp.FilterValues(EqualFloatVal, cmp.Comparer(equateAlways))
}

func equateAlways(_, _ interface{}) bool { return true }

// EqualFloatVal checks if the underlying value of x and y are same
// even if they are different types provided they are one of the int
// or float types
func EqualFloatVal(x, y interface{}) bool {
	xFloat, err1 := ToFloat(x)
	if err1 != nil {
		return false
	}
	yFloat, err2 := ToFloat(y)
	if err2 != nil {
		return false
	}
	return xFloat == yFloat
}

// ToFloat converts interface to float64, accepts only int and float types
func ToFloat(val interface{}) (float64, error) {
	floatType := reflect.TypeOf(float64(0))
	b := reflect.Indirect(reflect.ValueOf(val))
	if !b.Type().ConvertibleTo(floatType) {
		return 0, fmt.Errorf("cannot convert %v to float64", b.Type())
	}
	ans := b.Convert(floatType)
	return ans.Float(), nil
}

// EquateIoReader return an option to compare actual and expect bodies
func EquateIoReader() cmp.Option {
	return cmp.FilterValues(ioReaderCompare, cmp.Comparer(equateAlways))
}

// ioReaderCompare compares actual and expect bodies even if they have
// different types but have the same byte stream value. Body is built
// from shape_value_builder such that, actual would be of type
// "ioutil.nopCloser" and expect would be "aws.ReaderSeekerCloser"
func ioReaderCompare(expect, actual interface{}) bool {
	vActual, vExpect := reflect.ValueOf(actual), reflect.ValueOf(expect)
	if vActual.Type().String() != "ioutil.nopCloser" || vExpect.Type().String() != "aws.ReaderSeekerCloser" {
		return false
	}
	actualByte, err1 := ioutil.ReadAll(vActual.Interface().(io.Reader))
	if err1 != nil {
		log.Println("couldn't read the body from actual response")
	}
	expectByte, err2 := ioutil.ReadAll(vExpect.Interface().(io.Reader))
	if err2 != nil {
		log.Println("couldn't read the body from expect response")
	}
	return (actual != nil && expect != nil) && (bytes.Compare(actualByte, expectByte) == 0)
}

// StringEqual asserts that two strings are equal else returns false by wrapping an error message
func StringEqual(t *testing.T, expectVal, actualVal string) error {
	t.Helper()
	if expectVal != actualVal {
		t.Fatalf("String comparision failed,\n\texpect: %s\n\tactual: %s\n", expectVal, actualVal)
	}
	return nil
}

// ReadBody returns the request body as byte slice without erasing it
func ReadBody(t *testing.T, req *request.Request) []byte {
	t.Helper()
	var bytesReqBody []byte
	var err error
	if req.HTTPRequest.Body != nil {
		bytesReqBody, err = ioutil.ReadAll(req.HTTPRequest.Body)
		if err != nil {
			t.Fatalf("unable to read body from request, %v", err)
		}
	}
	req.HTTPRequest.Body = ioutil.NopCloser(bytes.NewBuffer(bytesReqBody))
	return bytesReqBody
}

// AssertRequestMethodEquals asserts if method field in request and expect value are equal
func AssertRequestMethodEquals(t *testing.T, expectVal string, actualVal string) bool {
	return StringEqual(t, expectVal, actualVal) == nil
}

// AssertRequestURLMatches asserts if request URL in request and expect are equal. True
// if all URL fields, (path, query, hostname, protocol name etc) in request and expect
// are equal
func AssertRequestURLMatches(t *testing.T, expectVal string, actualVal string) bool {
	return AssertURL(t, expectVal, actualVal)
}

// AssertRequestURLPathMatches asserts if the path field in request and expect are equal
func AssertRequestURLPathMatches(t *testing.T, expectVal string, actualVal string) bool {
	return StringEqual(t, expectVal, actualVal) == nil
}

// AssertRequestURLQueryMatches asserts if query values in request and expect are equal.
// Values of query string in request and expect are equal even if they have different
// orders
func AssertRequestURLQueryMatches(t *testing.T, expectVal string, req *request.Request, msgAndArgs ...interface{}) bool {
	queryRequest := req.HTTPRequest.URL.Query() //parsed RawQuery of "req" to get the values inside
	expectQ, err := url.ParseQuery(expectVal)

	if err != nil {
		t.Fatalf("unable to parse query from expect, %v", err)
	}
	var r DiffReporter
	for expectKey, expectVal := range expectQ {
		reqVal := queryRequest.Get(expectKey)
		if !cmp.Equal(expectVal[0], reqVal, cmp.Reporter(&r)) {
			t.Fatalf("query values inside request and expect don't match\n%s\n", r.String())
		}
	}
	return true
}

// AssertRequestHeadersMatch asserts if headers in request and expect are equal. True if each
// header key in the expected map is present in the request with equal values. request may
// have additional headers outside the expected ones.
func AssertRequestHeadersMatch(t *testing.T, expectHeader map[string]interface{}, req *request.Request, msgAndArgs ...interface{}) bool {
	for key, valExpect := range expectHeader {
		valReq := req.HTTPRequest.Header.Get(key)
		if key == "Header-Json-Value" {
			expectJSONValue, err1 := protocol.DecodeJSONValue(valExpect.(string), protocol.Base64Escape)
			if err1 != nil {
				t.Fatalf("unable to parse expected JSON, %v", err1)
			}
			responseJSONValue, err2 := protocol.DecodeJSONValue(valReq, protocol.Base64Escape)
			if err2 != nil {
				t.Fatalf("unable to parse response JSON, %v", err2)
			}

			var r DiffReporter
			for key1, val1 := range expectJSONValue {
				if !cmp.Equal(val1, responseJSONValue[key1], FloatIntEquate(), cmp.Reporter(&r)) {
					t.Fatalf("aws.JSON value from expect and response don't match\n%s\n", r.String())
				}
			}
			continue
		}
		if valReq == "" || valReq != valExpect {
			t.Fatalf("header values don't match for key: %q,\nexpect: %s\nactual: %s", key, valExpect, valReq)
		}
	}
	return true
}

// AssertRequestBodyEqualsBytes asserts if base64 encoded string inside request body is equal
// to expected value.
func AssertRequestBodyEqualsBytes(t *testing.T, expectVal string, req *request.Request) bool {
	bytesReqBody := ReadBody(t, req)
	return StringEqual(t, expectVal, string(bytesReqBody)) == nil
}

// AssertRequestBodyEqualsJSON verifies that the json value in request body
// string matches the expectVal map
func AssertRequestBodyEqualsJSON(t *testing.T, expectVal map[string]interface{}, req *request.Request, msgAndArgs ...interface{}) bool {
	bytesReqBody := ReadBody(t, req)
	actualVal := map[string]interface{}{}
	if err := json.Unmarshal(bytesReqBody, &actualVal); err != nil {
		t.Fatalf("unable to parse expected JSON, %v", err)
	}
	var r DiffReporter
	for key, val := range expectVal {
		if key == "JsonValue" && !AssertJSON(t, reflect.ValueOf(val).String(), reflect.ValueOf(actualVal[key]).String()) {
			t.Fatalf("AssertJSON failed when key is %s,\nexpect: %s\n\t\tactual: %s\t", key, val, actualVal[key])
		} else if key != "JsonValue" {
			if !cmp.Equal(val, actualVal[key], FloatIntEquate(), cmp.Reporter(&r)) {
				t.Fatalf(r.String())
			}
		}
	}
	return true
}

// AssertRequestBodyMatchesXML verifies that the expect xml string matches the actual. True if
// XML string inside request and expect are equal. For XML string, order of
// elements and attributes are significant while whitespaces are not.
func AssertRequestBodyMatchesXML(t *testing.T, expectVal string, req *request.Request, container interface{}, msgAndArgs ...interface{}) bool {
	r := req.HTTPRequest
	if r.Body == nil {
		t.Fatalf("request body is nil")
	}
	body := util.SortXML(r.Body)

	return AssertXML(t, expectVal, util.Trim(string(body)), container)
}

// AssertRequestBodyEqualsString asserts if the request body and expect value are exactly
// equal
func AssertRequestBodyEqualsString(t *testing.T, expectVal string, req *request.Request, msgAndArgs ...interface{}) bool {
	bytesReqBody := ReadBody(t, req)
	return StringEqual(t, expectVal, string(bytesReqBody)) == nil
}

// AssertRequestIDEquals asserts if requestID field in request and expect value are equal
func AssertRequestIDEquals(t *testing.T, expectVal string, actualVal string) bool {
	return StringEqual(t, expectVal, actualVal) == nil
}

// AssertResponseDataEquals asserts if data in response and error are equal. True if all
// fields inside the structure parsed from expect value are equal to the corresponding
// response fields
func AssertResponseDataEquals(t *testing.T, expectResponse interface{}, actualResponse interface{}, msgAndArgs ...interface{}) bool {
	var r DiffReporter
	// EquateApprox considers two values to be equal if their difference is less than 1E-7
	if !cmp.Equal(expectResponse, actualResponse, EquateIoReader(), cmpopts.EquateApprox(1E-6, 1E-7), cmp.Reporter(&r)) {
		t.Fatalf(r.String())
	}
	return true
}

// AssertResponseErrorIsKindOf asserts if code in response error and expect value are equal
func AssertResponseErrorIsKindOf(t *testing.T, expectVal string, err error) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		return StringEqual(t, expectVal, awsErr.Code()) == nil
	}
	return true
}

// AssertResponseErrorMessageEquals asserts if error message in response error and expect
// value are equal
func AssertResponseErrorMessageEquals(t *testing.T, expectVal string, err error) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		return StringEqual(t, expectVal, awsErr.Message()) == nil
	}
	return true
}

// AssertResponseErrorRequestIDEquals asserts if requestID field inside response error and
// expect value are equal
func AssertResponseErrorRequestIDEquals(t *testing.T, expectVal string, err error) bool {
	if reqErr, ok := err.(awserr.RequestFailure); ok {
		return StringEqual(t, expectVal, reqErr.RequestID()) == nil
	}
	return true
}

// AssertResponseErrorDataEquals asserts if all fields inside the structure parsed from expect
// value are equal to the corresponding response error data fields. This is not implemented in
// because Go SDK V1 doesn't expose the error data
func AssertResponseErrorDataEquals(t *testing.T, expectVal map[string]interface{}, err error, msgAndArgs ...interface{}) {
	if testing.Short() {
		t.Skip("\nskipping responseErrorDataEquals assertion")
	}
}
