// +build go1.9

package awstesting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/util"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// CompareValue type is a function which returns a string used in
// requestHeadersMatch assertion
type CompareValue func(actual string) string

// JSONValueCompareWith returns a compare function to compare Header JSON
// Strings which is used in requestHeadersMatch assertion when the
// "Header-Json-Value"
// The compare function returns a string which lists the differences between
// two decoded JSON strings. Length of the returned string would be 0 if they
// are equal
func JSONValueCompareWith(t *testing.T, expect string) CompareValue {
	t.Helper()
	return func(actual string) string {
		expectJSONValue, err1 := protocol.DecodeJSONValue(expect, protocol.Base64Escape)
		if err1 != nil {
			t.Fatalf("unable to parse expected JSON, %v", err1)
		}
		responseJSONValue, err2 := protocol.DecodeJSONValue(actual, protocol.Base64Escape)
		if err2 != nil {
			t.Fatalf("unable to parse response JSON, %v", err2)
		}
		return cmp.Diff(expectJSONValue, responseJSONValue)
	}
}

// DefaultCompareWith returns a compare function to compare header strings
// The compare function returns a string which lists the differences between
// two strings. Length of the returned string would be 0 if they are equal
func DefaultCompareWith(t *testing.T, expect string) CompareValue {
	return func(actual string) string {
		return cmp.Diff(expect, actual)
	}
}

// NumberEquals returns an option which compares floats with ints
// and return true if they have same value, Ex: (1.00 , 1)  are equal
func NumberEquals() cmp.Option {
	return cmp.FilterValues(EqualNumberValue, cmp.Comparer(equateAlways))
}

func equateAlways(_, _ interface{}) bool { return true }

// EqualNumberValue checks if the underlying value of x and y are same
// even if they are different types provided they are one of the int
// or float types
func EqualNumberValue(x, y interface{}) bool {
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
		panic("couldn't read the body from actual response")
	}
	expectByte, err2 := ioutil.ReadAll(vExpect.Interface().(io.Reader))
	if err2 != nil {
		panic("couldn't read the body from expect response")
	}
	return (actual != nil && expect != nil) && (bytes.Compare(actualByte, expectByte) == 0)
}

// CopyRequestBody returns the request body as byte slice without erasing it
func CopyRequestBody(t *testing.T, req *request.Request) []byte {
	t.Helper()
	var bytesReqBody []byte
	var err error
	if req.HTTPRequest.Body != nil {
		bytesReqBody, err = ioutil.ReadAll(req.HTTPRequest.Body)
		if err != nil {
			t.Fatalf("unable to read body from request, %v", err)
		}
	}
	req.SetBufferBody(bytesReqBody)
	return bytesReqBody
}

// AssertRequestMethodEquals asserts if method field in request and expect value are equal
func AssertRequestMethodEquals(t *testing.T, expectVal string, actualVal string) bool {
	if expectVal != actualVal {
		t.Error(cmp.Diff(expectVal, actualVal))
	}
	return true
}

// AssertRequestURLMatches asserts if request URL in request and expect are equal. True
// if all URL fields, (path, query, hostname, protocol name etc) in request and expect
// are equal
func AssertRequestURLMatches(t *testing.T, expectVal string, actualVal string) bool {
	return AssertURL(t, expectVal, actualVal)
}

// AssertRequestURLPathMatches asserts if the path field in request and expect are equal
func AssertRequestURLPathMatches(t *testing.T, expectVal string, actualVal string) bool {
	if expectVal != actualVal {
		t.Error(cmp.Diff(expectVal, actualVal))
	}
	return true
}

// AssertRequestURLQueryMatches asserts if query values in request and expect are equal.
// Values of query string in request and expect are equal even if they have different
// orders
func AssertRequestURLQueryMatches(t *testing.T, expectVal string, reqURL *url.URL) bool {
	queryRequest := reqURL.Query()
	expectQ, err := url.ParseQuery(expectVal)
	if err != nil {
		t.Fatalf("unable to parse query from expect, %v", err)
	}
	for expectKey, expectVal := range expectQ {
		reqVal := queryRequest[expectKey]
		sort.Strings(expectVal)
		sort.Strings(reqVal)
		if !cmp.Equal(expectVal, reqVal) {
			t.Errorf("query values inside request and expect don't match\n%s\n", cmp.Diff(expectQ, queryRequest))
		}
	}
	return true
}

// AssertRequestHeadersMatch asserts if headers in request and expect are equal. True if each
// header key in the expected map is present in the request with equal values. request may
// have additional headers outside the expected ones.
func AssertRequestHeadersMatch(t *testing.T, expectHeader map[string]CompareValue, actualHeader http.Header) bool {
	for key, Func := range expectHeader {
		if v := Func(actualHeader.Get(key)); len(v) != 0 {
			t.Errorf("Headers in expect and actual don't match,\n%v", v)
		}
	}
	return true
}

// AssertRequestBodyEqualsBytes asserts if base64 encoded string inside request body is equal
// to expected value.
func AssertRequestBodyEqualsBytes(t *testing.T, expectVal string, req *request.Request) bool {
	actualVal := string(CopyRequestBody(t, req))
	if expectVal != actualVal {
		t.Error(cmp.Diff(expectVal, actualVal))
	}
	return true
}

// AssertRequestBodyEqualsJSON verifies that the json value in request body
// string matches the expectVal map
func AssertRequestBodyEqualsJSON(t *testing.T, expectVal map[string]interface{}, req *request.Request) bool {
	bytesReqBody := CopyRequestBody(t, req)
	actualVal := map[string]interface{}{}
	if err := json.Unmarshal(bytesReqBody, &actualVal); err != nil {
		t.Fatalf("unable to parse expected JSON, %v", err)
	}
	for key, val := range expectVal {
		if key == "JsonValue" && !AssertJSON(t, reflect.ValueOf(val).String(), reflect.ValueOf(actualVal[key]).String()) {
			t.Errorf("AssertJSON failed when key is %s,\nexpect: %s\n\t\tactual: %s\t", key, val, actualVal[key])
		} else if key != "JsonValue" {
			if !cmp.Equal(val, actualVal[key], NumberEquals()) {
				t.Error(cmp.Diff(val, actualVal[key], NumberEquals()))
			}
		}
	}
	return true
}

// AssertRequestBodyMatchesXML verifies that the expect xml string matches the actual. True if
// XML string inside request and expect are equal. For XML string, order of
// elements and attributes are significant while whitespaces are not.
func AssertRequestBodyMatchesXML(t *testing.T, expectVal string, req *request.Request, container interface{}) bool {
	t.Skip("skipping RequestBodyMatchesXML assertion, this is fixed in PR# 2804 in V1 SDK Go")
	r := req.HTTPRequest
	if r.Body == nil {
		t.Fatalf("request body is nil")
	}
	body := util.SortXML(r.Body)

	return AssertXML(t, expectVal, util.Trim(string(body)), container)
}

// AssertRequestBodyEqualsString asserts if the request body and expect value are exactly
// equal
func AssertRequestBodyEqualsString(t *testing.T, expectVal string, req *request.Request) bool {
	actualVal := string(CopyRequestBody(t, req))
	if expectVal != actualVal {
		t.Error(cmp.Diff(expectVal, actualVal))
	}
	return true
}

// AssertRequestIDEquals asserts if requestID field in request and expect value are equal
func AssertRequestIDEquals(t *testing.T, expectVal string, actualVal string) bool {
	if expectVal != actualVal {
		t.Error(cmp.Diff(expectVal, actualVal))
	}
	return true
}

// AssertResponseValueEquals asserts if data in expectResponse and actualResponse are equal.
// expectResponse and actualResponse are structures and the function returns true if all fields
// inside the structure parsed from expect value are equal to the
// corresponding actualResponse fields
func AssertResponseValueEquals(t *testing.T, expectResponse interface{}, actualResponse interface{}) bool {
	// EquateIoReader considers expectResponse's body and actualResponse's body field as
	// equal even when they have different wrapping types (refer to the function definition)
	// if their underlying byte stream is same

	// EquateApprox considers two values to be equal if their difference is less than 1E-7. This
	// is used to compare floats or doubles
	if !cmp.Equal(expectResponse, actualResponse, EquateIoReader(), cmpopts.EquateApprox(1E-6, 1E-7)) {
		t.Error(cmp.Diff(expectResponse, actualResponse, EquateIoReader(), cmpopts.EquateApprox(1E-6, 1E-7)))
	}
	return true
}

// AssertResponseErrorIsKindOf asserts if code in response error and expect value are equal
func AssertResponseErrorIsKindOf(t *testing.T, expectVal string, err error) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		actualVal := awsErr.Code()
		if expectVal != actualVal {
			t.Error(cmp.Diff(expectVal, actualVal))
		}
		return true
	}
	t.Error("err is not of type awserr.Error")
	return false
}

// AssertResponseErrorMessageEquals asserts if error message in response error and expect
// value are equal
func AssertResponseErrorMessageEquals(t *testing.T, expectVal string, err error) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		actualVal := awsErr.Message()
		if expectVal != actualVal {
			t.Error(cmp.Diff(expectVal, actualVal))
		}
		return true
	}
	t.Error("err is not of type awserr.Error")
	return false
}

// AssertResponseErrorRequestIDEquals asserts if requestID field inside response error and
// expect value are equal
func AssertResponseErrorRequestIDEquals(t *testing.T, expectVal string, err error) bool {
	if reqErr, ok := err.(awserr.RequestFailure); ok {
		actualVal := reqErr.RequestID()
		if expectVal != actualVal {
			t.Error(cmp.Diff(expectVal, actualVal))
		}
		return true
	}
	t.Error("err is not of type awserr.Error")
	return false
}

// AssertResponseErrorDataEquals asserts if all fields inside the structure parsed from expect
// value are equal to the corresponding response error data fields. This is not implemented in
// because Go SDK V1 doesn't expose the error data
func AssertResponseErrorDataEquals(t *testing.T, expectVal interface{}, err error) {
	t.Skip("skipping responseErrorDataEquals assertion because the SDK doesn't expose the error data yet")
}
