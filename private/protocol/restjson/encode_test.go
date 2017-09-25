package restjson

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/private/protocol"
)

func TestEncodeNestedShape(t *testing.T) {
	_, reader, err := encode("PUT", "/path", shape{
		NestedShape: &nestedShape{
			Value: aws.String("some value"),
		},
	})
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatalf("expect no read error, %v", err)
	}

	expect := `{"nestedShape":{"value":"some value"}}`
	if e, a := expect, string(b); e != a {
		t.Errorf("expect bodies to match, did not.\n,\tExpect:\n%s\n\tActual:\n%s\n", e, a)
	}
}

func TestEncodePayloadShape(t *testing.T) {
	_, reader, err := encode("PUT", "/path", shape{
		PayloadShape: &nestedShape{
			Value: aws.String("some value"),
		},
	})
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatalf("expect no read error, %v", err)
	}

	expect := `{"value":"some value"}`
	if e, a := expect, string(b); e != a {
		t.Errorf("expect bodies to match, did not.\n,\tExpect:\n%s\n\tActual:\n%s\n", e, a)
	}
}

func TestEncodePayloadStream(t *testing.T) {
	_, reader, err := encode("PUT", "/path", shape{
		PayloadStream: strings.NewReader("some value"),
	})
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatalf("expect no read error, %v", err)
	}

	expect := "some value"
	if e, a := expect, string(b); e != a {
		t.Errorf("expect bodies to match, did not.\n,\tExpect:\n%s\n\tActual:\n%s\n", e, a)
	}
}

type shape struct {
	NestedShape   *nestedShape
	PayloadShape  *nestedShape
	PayloadStream io.ReadSeeker
}

func (s *shape) MarshalFields(e protocol.FieldEncoder) error {
	if s.NestedShape != nil {
		e.SetFields(protocol.BodyTarget, "nestedShape", s.NestedShape, protocol.Metadata{})
	}
	if s.PayloadShape != nil {
		e.SetFields(protocol.PayloadTarget, "payloadShape", s.PayloadShape, protocol.Metadata{})
	}
	if s.PayloadStream != nil {
		e.SetStream(protocol.PayloadTarget, "payloadReader", protocol.ReadSeekerStream{V: s.PayloadStream}, protocol.Metadata{})
	}
	return nil
}

type nestedShape struct {
	Value *string
}

func (s *nestedShape) MarshalFields(e protocol.FieldEncoder) error {
	if s.Value != nil {
		e.SetValue(protocol.BodyTarget, "value", protocol.StringValue(*s.Value), protocol.Metadata{})
	}
	return nil
}

func encode(method, path string, s shape) (*http.Request, io.ReadSeeker, error) {
	origReq, _ := http.NewRequest(method, "https://service.amazonaws.com"+path, nil)

	e := NewEncoder(origReq)
	s.MarshalFields(e)
	return e.Encode()
}
