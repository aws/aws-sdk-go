package protocol

import (
	"io"
	"time"
)

// A FieldMarshaler interface is used to marshal struct fields when encoding.
type FieldMarshaler interface {
	MarshalFields(FieldEncoder) error
}

// ValueMarshaler provides a generic type for all encoding field values to be
// passed into a encoder's methods with.
type ValueMarshaler interface {
	MarshalValue() (string, error)
	MarshalValueBuf([]byte) ([]byte, error)
}

// A StreamMarshaler interface is used to marshal a stream when encoding.
type StreamMarshaler interface {
	MarshalStream() (io.ReadSeeker, error)
}

// A ListEncoder provides the interface for encoders that will encode List elements.
type ListEncoder interface {
	ListAddValue(v ValueMarshaler)
	ListAddList(fn func(ListEncoder))
	ListAddMap(fn func(MapEncoder))
	ListAddFields(m FieldMarshaler)
}

// A MapEncoder provides the interface for encoders that will encode map elements.
type MapEncoder interface {
	MapSetValue(k string, v ValueMarshaler)
	MapSetList(k string, fn func(ListEncoder))
	MapSetMap(k string, fn func(MapEncoder))
	MapSetFields(k string, m FieldMarshaler)
}

// A FieldEncoder provides the interface for encoding struct field members.
type FieldEncoder interface {
	SetValue(t Target, k string, v ValueMarshaler, meta Metadata)
	SetStream(t Target, k string, v StreamMarshaler, meta Metadata)
	SetList(t Target, k string, fn func(ListEncoder), meta Metadata)
	SetMap(t Target, k string, fn func(MapEncoder), meta Metadata)
	SetFields(t Target, k string, m FieldMarshaler, meta Metadata)
}

// EncodeStringList returns a function that will add the slice's values to
// a list Encoder.
func EncodeStringList(vs []*string) func(ListEncoder) {
	return func(le ListEncoder) {
		for _, v := range vs {
			le.ListAddValue(StringValue(*v))
		}
	}
}

// EncodeStringMap returns a function that will add the map's values to
// a map Encoder.
func EncodeStringMap(vs map[string]*string) func(MapEncoder) {
	return func(me MapEncoder) {
		for k, v := range vs {
			me.MapSetValue(k, StringValue(*v))
		}
	}
}

// EncodeInt64List returns a function that will add the slice's values to
// a list Encoder.
func EncodeInt64List(vs []*int64) func(ListEncoder) {
	return func(le ListEncoder) {
		for _, v := range vs {
			le.ListAddValue(Int64Value(*v))
		}
	}
}

// EncodeInt64Map returns a function that will add the map's values to
// a map Encoder.
func EncodeInt64Map(vs map[string]*int64) func(MapEncoder) {
	return func(me MapEncoder) {
		for k, v := range vs {
			me.MapSetValue(k, Int64Value(*v))
		}
	}
}

// EncodeFloat64List returns a function that will add the slice's values to
// a list Encoder.
func EncodeFloat64List(vs []*float64) func(ListEncoder) {
	return func(le ListEncoder) {
		for _, v := range vs {
			le.ListAddValue(Float64Value(*v))
		}
	}
}

// EncodeFloat64Map returns a function that will add the map's values to
// a map Encoder.
func EncodeFloat64Map(vs map[string]*float64) func(MapEncoder) {
	return func(me MapEncoder) {
		for k, v := range vs {
			me.MapSetValue(k, Float64Value(*v))
		}
	}
}

// EncodeBoolList returns a function that will add the slice's values to
// a list Encoder.
func EncodeBoolList(vs []*bool) func(ListEncoder) {
	return func(le ListEncoder) {
		for _, v := range vs {
			le.ListAddValue(BoolValue(*v))
		}
	}
}

// EncodeBoolMap returns a function that will add the map's values to
// a map Encoder.
func EncodeBoolMap(vs map[string]*bool) func(MapEncoder) {
	return func(me MapEncoder) {
		for k, v := range vs {
			me.MapSetValue(k, BoolValue(*v))
		}
	}
}

// EncodeTimeList returns a function that will add the slice's values to
// a list Encoder.
func EncodeTimeList(vs []*time.Time) func(ListEncoder) {
	return func(le ListEncoder) {
		for _, v := range vs {
			le.ListAddValue(TimeValue{V: *v})
		}
	}
}

// EncodeTimeMap returns a function that will add the map's values to
// a map Encoder.
func EncodeTimeMap(vs map[string]*time.Time) func(MapEncoder) {
	return func(me MapEncoder) {
		for k, v := range vs {
			me.MapSetValue(k, TimeValue{V: *v})
		}
	}
}

// A FieldBuffer provides buffering of fields so the number of
// allocations are reduced by providng a persistent buffer that is
// used between fields.
type FieldBuffer struct {
	buf []byte
}

// GetValue will retrieve the ValueMarshaler's value by appending the
// value to the buffer. Will return the buffer that was populated.
//
// This buffer is only valid until the next time GetValue is called.
func (b *FieldBuffer) GetValue(m ValueMarshaler) ([]byte, error) {
	v, err := m.MarshalValueBuf(b.buf)
	b.buf = v
	b.buf = b.buf[0:0]
	return v, err
}
