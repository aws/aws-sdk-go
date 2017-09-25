package rest

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/private/protocol"
)

type shape struct {
	HeaderValue   *string
	PathValue     *string
	QueryValue    *string
	QueryList     []*string
	HeadersMap    map[string]*string
	QueryMap      map[string]*string
	QueryMapList  map[string][]*string
	BodyValue     *string
	PayloadString *string
	PayloadBytes  []byte
	PayloadReader io.ReadSeeker
}

func (s *shape) MarshalAWS(e protocol.FieldEncoder) {
	if s.HeaderValue != nil {
		e.SetValue(protocol.HeaderTarget, "header-value", protocol.StringValue(*s.HeaderValue), protocol.Metadata{})
	}
	if s.PathValue != nil {
		e.SetValue(protocol.PathTarget, "key", protocol.StringValue(*s.PathValue), protocol.Metadata{})
	}
	if s.QueryValue != nil {
		e.SetValue(protocol.QueryTarget, "queryKey", protocol.StringValue(*s.QueryValue), protocol.Metadata{})
	}
	if len(s.QueryList) > 0 {
		e.SetList(protocol.QueryTarget, "queryKey", func(le protocol.ListEncoder) {
			for _, v := range s.QueryList {
				le.ListAddValue(protocol.StringValue(*v))
			}
		}, protocol.Metadata{})
	}
	if len(s.HeadersMap) > 0 {
		e.SetMap(protocol.HeadersTarget, "prefix-", func(me protocol.MapEncoder) {
			for k, v := range s.HeadersMap {
				me.MapSetValue(k, protocol.StringValue(*v))
			}
		}, protocol.Metadata{})
	}
	if len(s.QueryMap) > 0 {
		e.SetMap(protocol.QueryTarget, "unused", func(me protocol.MapEncoder) {
			for k, v := range s.QueryMap {
				me.MapSetValue(k, protocol.StringValue(*v))
			}
		}, protocol.Metadata{})
	}
	if len(s.QueryMapList) > 0 {
		e.SetMap(protocol.QueryTarget, "unused", func(me protocol.MapEncoder) {
			for k, v := range s.QueryMapList {
				me.MapSetList(k, func(le protocol.ListEncoder) {
					for _, v := range v {
						le.ListAddValue(protocol.StringValue(*v))
					}
				})
			}
		}, protocol.Metadata{})
	}
	if s.BodyValue != nil {
		e.SetValue(protocol.BodyTarget, "bodyValue", protocol.StringValue(*s.BodyValue), protocol.Metadata{})
	}
	if s.PayloadString != nil {
		e.SetStream(protocol.PayloadTarget, "payloadString", protocol.StringStream(*s.PayloadString), protocol.Metadata{})
	}
	if s.PayloadBytes != nil {
		e.SetStream(protocol.PayloadTarget, "payloadBytes", protocol.BytesStream(s.PayloadBytes), protocol.Metadata{})
	}
	if s.PayloadReader != nil {
		e.SetStream(protocol.PayloadTarget, "payloadReader", protocol.ReadSeekerStream{V: s.PayloadReader}, protocol.Metadata{})
	}
}

func TestSetHeaderValue(t *testing.T) {
	req, body, err := encode("GET", "/path", shape{
		HeaderValue: aws.String("thevalue"),
	})

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if body != nil {
		t.Errorf("expect no body, got %v", body)
	}

	if e, a := "thevalue", req.Header.Get("header-value"); e != a {
		t.Errorf("expect %s header value, got %s", e, a)
	}
}

func TestSetPathValue(t *testing.T) {
	req, body, err := encode("GET", "/{key}", shape{
		PathValue: aws.String("value"),
	})

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if body != nil {
		t.Errorf("expect no body, got %v", body)
	}

	if e, a := "/value", req.URL.Path; e != a {
		t.Errorf("expect %s path, got %s", e, a)
	}
}

func TestSetQueryValue(t *testing.T) {
	req, body, err := encode("GET", "/path", shape{
		QueryValue: aws.String("queryValue"),
	})

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if body != nil {
		t.Errorf("expect no body, got %v", body)
	}

	query := req.URL.Query()
	if e, a := 1, len(query); e != a {
		t.Errorf("expect %d query values, got %d", e, a)
	}
	if e, a := "queryValue", query.Get("queryKey"); e != a {
		t.Errorf("expect %s for 'queryKey' querystring, got %s", e, a)
	}
}

func TestSetBody(t *testing.T) {
	req, body, err := encode("GET", "/path", shape{
		BodyValue: aws.String("a value"),
	})

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if body != nil {
		t.Errorf("expect no body, got %v", body)
	}

	query := req.URL.Query()
	if e, a := 1, len(query); e != a {
		t.Errorf("expect %d query values, got %d", e, a)
	}
	if e, a := "a value", query.Get("bodyValue"); e != a {
		t.Errorf("expect %s for 'bodyValue' querystring, got %s", e, a)
	}
}

func TestSetBody_Error(t *testing.T) {
	req, body, err := encode("POST", "/path", shape{
		BodyValue: aws.String("a value"),
	})

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "body target not supported", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %q to be in error %q, was not", e, a)
	}

	if body != nil {
		t.Errorf("expect no body, got %v", body)
	}
	if req != nil {
		t.Errorf("expect no req, got %v", req)
	}
}

func TestSetQueryList(t *testing.T) {
	req, body, err := encode("GET", "/path", shape{
		QueryList: []*string{aws.String("a"), aws.String("b")},
	})

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if body != nil {
		t.Errorf("expect no body, got %v", body)
	}

	query := req.URL.Query()
	if e, a := 1, len(query); e != a {
		t.Errorf("expect %d query values, got %d", e, a)
	}
	vs, ok := query["queryKey"]
	if !ok {
		t.Fatalf("expect queryKey to exist, was not found")
	}
	for i, v := range []string{"a", "b"} {
		if e, a := v, vs[i]; e != a {
			t.Errorf("expect %s for 'queryKey[%d]' querystring, got %s", e, i, a)
		}
	}
}

func TestSetHeadersMap(t *testing.T) {
	req, body, err := encode("GET", "/path", shape{
		HeadersMap: map[string]*string{
			"abc":  aws.String("123"),
			"else": aws.String("other"),
		},
	})

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if body != nil {
		t.Errorf("expect no body, got %v", body)
	}

	if e, a := 2, len(req.Header); e != a {
		t.Errorf("expect %d header values, got %d", e, a)
	}
	for k, v := range map[string]string{"abc": "123", "else": "other"} {
		if e, a := v, req.Header.Get("prefix-"+k); e != a {
			t.Errorf("expect %s for 'prefix-%s' header, got %s", e, k, a)
		}
	}
}

func TestSetQueryMap(t *testing.T) {
	req, body, err := encode("GET", "/path", shape{
		QueryMap: map[string]*string{
			"a": aws.String("abc"),
			"b": aws.String("123"),
		},
	})

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if body != nil {
		t.Errorf("expect no body, got %v", body)
	}

	query := req.URL.Query()
	if e, a := 2, len(query); e != a {
		t.Errorf("expect %d query values, got %d", e, a)
	}
	for k, v := range map[string]string{"a": "abc", "b": "123"} {
		if e, a := v, query.Get(k); e != a {
			t.Errorf("expect %s for '%s' querystring, got %s", e, k, a)
		}
	}
}

func TestSetQueryMapList(t *testing.T) {
	req, body, err := encode("GET", "/path", shape{
		QueryMapList: map[string][]*string{
			"a": {aws.String("a1"), aws.String("a2")},
			"b": {aws.String("b1"), aws.String("b2")},
		},
	})

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if body != nil {
		t.Errorf("expect no body, got %v", body)
	}

	query := req.URL.Query()
	if e, a := 2, len(query); e != a {
		t.Errorf("expect %d query values, got %d", e, a)
	}
	for k, vs := range map[string][]string{"a": []string{"a1", "a2"}, "b": []string{"b1", "b2"}} {
		as, ok := query[k]
		if !ok {
			t.Fatalf("expect %s to exist, was not", k)
		}
		for i, v := range vs {
			if e, a := v, as[i]; e != a {
				t.Errorf("expect %s for '%s' querystring, got %s", e, k, a)
			}
		}
	}
}

func TestSetPayloadString(t *testing.T) {
	_, body, err := encode("POST", "/path", shape{
		PayloadString: aws.String("a value"),
	})

	if err != nil {
		t.Fatalf("expect no encode error, got %v", err)
	}
	if body == nil {
		t.Fatalf("expect body, got none")
	}

	b, err := ioutil.ReadAll(body)
	if err != nil {
		t.Fatalf("expect no body read error, got %v", err)
	}
	if e, a := "a value", string(b); e != a {
		t.Errorf("expect %s body, got %s", e, a)
	}
}
func TestSetPayloadBytes(t *testing.T) {
	_, body, err := encode("POST", "/path", shape{
		PayloadBytes: []byte("a value"),
	})

	if err != nil {
		t.Fatalf("expect no encode error, got %v", err)
	}
	if body == nil {
		t.Fatalf("expect body, got none")
	}

	b, err := ioutil.ReadAll(body)
	if err != nil {
		t.Fatalf("expect no body read error, got %v", err)
	}
	if e, a := "a value", string(b); e != a {
		t.Errorf("expect %s body, got %s", e, a)
	}
}
func TestSetPayloadReader(t *testing.T) {
	_, body, err := encode("POST", "/path", shape{
		PayloadReader: strings.NewReader("a value"),
	})

	if err != nil {
		t.Fatalf("expect no encode error, got %v", err)
	}
	if body == nil {
		t.Fatalf("expect body, got none")
	}

	b, err := ioutil.ReadAll(body)
	if err != nil {
		t.Fatalf("expect no body read error, got %v", err)
	}
	if e, a := "a value", string(b); e != a {
		t.Errorf("expect %s body, got %s", e, a)
	}
}

func encode(method, path string, s shape) (*http.Request, io.ReadSeeker, error) {
	origReq, _ := http.NewRequest(method, "https://service.amazonaws.com"+path, nil)

	e := NewEncoder(origReq)
	s.MarshalAWS(e)
	return e.Encode()
}
