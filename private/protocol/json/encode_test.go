package json

import (
	"io"
	"io/ioutil"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/awstesting"
	"github.com/aws/aws-sdk-go/private/protocol"
)

func TestEncodeNestedShape(t *testing.T) {
	r, err := encode(baseShape{
		Payload: &payloadShape{
			Nested: &nestedShape{
				Value: aws.String("expected value"),
			},
		},
	})
	if err != nil {
		t.Fatalf("expect no marshal error, %v", err)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("expect no read error, %v", err)
	}

	expect := `{"nested":{"value":"expected value"}}`

	if e, a := expect, string(b); e != a {
		t.Errorf("expect bodies to match, did not.\n,\tExpect:\n%s\n\tActual:\n%s\n", e, a)
	}
}
func TestEncodeMapString(t *testing.T) {
	r, err := encode(baseShape{
		Payload: &payloadShape{
			MapStr: map[string]*string{
				"abc": aws.String("123"),
			},
		},
	})
	if err != nil {
		t.Fatalf("expect no marshal error, %v", err)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("expect no read error, %v", err)
	}

	expect := `{"mapstr":{"abc":"123"}}`

	if e, a := expect, string(b); e != a {
		t.Errorf("expect bodies to match, did not.\n,\tExpect:\n%s\n\tActual:\n%s\n", e, a)
	}
}
func TestEncodeMapShape(t *testing.T) {
	r, err := encode(baseShape{
		Payload: &payloadShape{
			MapShape: map[string]*nestedShape{
				"abc": {Value: aws.String("1")},
				"123": {IntVal: aws.Int64(123)},
			},
		},
	})
	if err != nil {
		t.Fatalf("expect no marshal error, %v", err)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("expect no read error, %v", err)
	}

	expect := `{"mapShape":{"abc":{"value":"1"},"123":{"intval":123}}}`

	awstesting.AssertJSON(t, expect, string(b), "expect bodies to match")
}
func TestEncodeListString(t *testing.T) {
	r, err := encode(baseShape{
		Payload: &payloadShape{
			ListStr: []*string{
				aws.String("abc"),
				aws.String("123"),
			},
		},
	})
	if err != nil {
		t.Fatalf("expect no marshal error, %v", err)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("expect no read error, %v", err)
	}

	expect := `{"liststr":["abc","123"]}`

	if e, a := expect, string(b); e != a {
		t.Errorf("expect bodies to match, did not.\n,\tExpect:\n%s\n\tActual:\n%s\n", e, a)
	}
}
func TestEncodeListFlatten(t *testing.T) {
	// TODO no JSON flatten
}
func TestEncodeListFlattened(t *testing.T) {
	// TODO No json flatten
}
func TestEncodeListNamed(t *testing.T) {
	// TODO no json named
}
func TestEncodeListShape(t *testing.T) {
	r, err := encode(baseShape{
		Payload: &payloadShape{
			ListShape: []*nestedShape{
				{Value: aws.String("abc")},
				{Value: aws.String("123")},
				{IntVal: aws.Int64(123)},
			},
		},
	})
	if err != nil {
		t.Fatalf("expect no marshal error, %v", err)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("expect no read error, %v", err)
	}

	expect := `{"listShape":[{"value":"abc"},{"value":"123"},{"intval":123}]}`

	if e, a := expect, string(b); e != a {
		t.Errorf("expect bodies to match, did not.\n,\tExpect:\n%s\n\tActual:\n%s\n", e, a)
	}
}

type baseShape struct {
	Payload *payloadShape
}

func (s *baseShape) MarshalFields(e protocol.FieldEncoder) error {
	if s.Payload != nil {
		e.SetFields(protocol.PayloadTarget, "payload", s.Payload, protocol.Metadata{})
	}
	return nil
}

type payloadShape struct {
	Value            *string
	IntVal           *int64
	TimeVal          *time.Time
	Nested           *nestedShape
	MapStr           map[string]*string
	MapFlatten       map[string]*string
	MapNamed         map[string]*string
	MapShape         map[string]*nestedShape
	MapFlattenShape  map[string]*nestedShape
	MapNamedShape    map[string]*nestedShape
	ListStr          []*string
	ListFlatten      []*string
	ListNamed        []*string
	ListShape        []*nestedShape
	ListFlattenShape []*nestedShape
	ListNamedShape   []*nestedShape
}

func (s *payloadShape) MarshalFields(e protocol.FieldEncoder) error {
	if s.Value != nil {
		e.SetValue(protocol.BodyTarget, "value", protocol.StringValue(*s.Value), protocol.Metadata{})
	}
	if s.IntVal != nil {
		e.SetValue(protocol.BodyTarget, "intval", protocol.Int64Value(*s.IntVal), protocol.Metadata{})
	}
	if s.TimeVal != nil {
		e.SetValue(protocol.BodyTarget, "timeval", protocol.TimeValue{
			V: *s.TimeVal, Format: protocol.UnixTimeFormat,
		}, protocol.Metadata{})
	}
	if s.Nested != nil {
		e.SetFields(protocol.BodyTarget, "nested", s.Nested, protocol.Metadata{})
	}
	if len(s.MapStr) > 0 {
		e.SetMap(protocol.BodyTarget, "mapstr", func(me protocol.MapEncoder) {
			for k, v := range s.MapStr {
				me.MapSetValue(k, protocol.StringValue(*v))
			}
		}, protocol.Metadata{})
	}
	if len(s.MapFlatten) > 0 {
		e.SetMap(protocol.BodyTarget, "mapFlatten", func(me protocol.MapEncoder) {
			for k, v := range s.MapFlatten {
				me.MapSetValue(k, protocol.StringValue(*v))
			}
		}, protocol.Metadata{
			Flatten: true,
		})
	}
	if len(s.MapNamed) > 0 {
		e.SetMap(protocol.BodyTarget, "mapNamed", func(me protocol.MapEncoder) {
			for k, v := range s.MapNamed {
				me.MapSetValue(k, protocol.StringValue(*v))
			}
		}, protocol.Metadata{
			MapLocationNameKey: "namedKey", MapLocationNameValue: "namedValue",
		})
	}
	if len(s.MapShape) > 0 {
		e.SetMap(protocol.BodyTarget, "mapShape", encodeNestedShapeMap(s.MapShape), protocol.Metadata{})
	}
	if len(s.MapFlattenShape) > 0 {
		e.SetMap(protocol.BodyTarget, "mapFlattenShape", encodeNestedShapeMap(s.MapFlattenShape), protocol.Metadata{
			Flatten: true,
		})
	}
	if len(s.MapNamedShape) > 0 {
		e.SetMap(protocol.BodyTarget, "mapNamedShape", encodeNestedShapeMap(s.MapNamedShape), protocol.Metadata{
			MapLocationNameKey: "namedKey", MapLocationNameValue: "namedValue",
		})
	}
	if len(s.ListStr) > 0 {
		e.SetList(protocol.BodyTarget, "liststr", func(le protocol.ListEncoder) {
			for _, v := range s.ListStr {
				le.ListAddValue(protocol.StringValue(*v))
			}
		}, protocol.Metadata{})
	}
	if len(s.ListFlatten) > 0 {
		e.SetList(protocol.BodyTarget, "listFlatten", func(le protocol.ListEncoder) {
			for _, v := range s.ListFlatten {
				le.ListAddValue(protocol.StringValue(*v))
			}
		}, protocol.Metadata{
			Flatten: true,
		})
	}
	if len(s.ListNamed) > 0 {
		e.SetList(protocol.BodyTarget, "listNamed", func(le protocol.ListEncoder) {
			for _, v := range s.ListNamed {
				le.ListAddValue(protocol.StringValue(*v))
			}
		}, protocol.Metadata{
			ListLocationName: "namedMember",
		})
	}
	if len(s.ListShape) > 0 {
		e.SetList(protocol.BodyTarget, "listShape", encodeNestedShapeList(s.ListShape), protocol.Metadata{})
	}
	if len(s.ListFlattenShape) > 0 {
		e.SetList(protocol.BodyTarget, "listFlattenShape", encodeNestedShapeList(s.ListFlattenShape), protocol.Metadata{
			Flatten: true,
		})
	}
	if len(s.ListNamedShape) > 0 {
		e.SetList(protocol.BodyTarget, "listNamedShape", encodeNestedShapeList(s.ListNamedShape), protocol.Metadata{
			ListLocationName: "namedMember",
		})
	}
	return nil
}

type nestedShape struct {
	Value    *string
	IntVal   *int64
	Prefixed *string
}

func (s *nestedShape) MarshalFields(e protocol.FieldEncoder) error {
	if s.Value != nil {
		e.SetValue(protocol.BodyTarget, "value", protocol.StringValue(*s.Value), protocol.Metadata{})
	}
	if s.IntVal != nil {
		e.SetValue(protocol.BodyTarget, "intval", protocol.Int64Value(*s.IntVal), protocol.Metadata{})
	}
	if s.Prefixed != nil {
		e.SetValue(protocol.BodyTarget, "prefixed", protocol.StringValue(*s.Prefixed), protocol.Metadata{})
	}
	return nil
}
func encodeNestedShapeMap(vs map[string]*nestedShape) func(protocol.MapEncoder) {
	return func(me protocol.MapEncoder) {
		for k, v := range vs {
			me.MapSetFields(k, v)
		}
	}
}
func encodeNestedShapeList(vs []*nestedShape) func(protocol.ListEncoder) {
	return func(le protocol.ListEncoder) {
		for _, v := range vs {
			le.ListAddFields(v)
		}
	}
}

func encode(s baseShape) (io.ReadSeeker, error) {
	e := NewEncoder()
	s.MarshalFields(e)
	return e.Encode()
}
