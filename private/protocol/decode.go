package protocol

import (
	"io"
	"time"
)

// FieldUnmarshaler used by protocol unmarshaling to unmarshal a type's nested fields.
type FieldUnmarshaler interface {
	UnmarshalFields(FieldDecoder) error
}

// A FieldValue is a value that will be unmarshalered from the decoder to
// a concrete value.
type FieldValue interface{}

// ListDecoder provides the interface for unmarshaling list elements from the
// underlying decoder.
type ListDecoder interface {
	ListGet(fn func(v FieldValue))
	ListGetList(fn func(n int, ld ListDecoder))
	ListGetMap(fn func(ks []string, md MapDecoder))
	ListGetFields(m FieldUnmarshaler)
}

// MapDecoder provides the interface for unmarshaling map elements from the
// underlying decoder. The map key the value is retrieved from is k.
type MapDecoder interface {
	MapGet(k string, fn func(v FieldValue))
	MapGetList(k string, fn func(n int, ld ListDecoder))
	MapGetMap(k string, fn func(ks []string, fd FieldDecoder))
	MapGetFields(k string, fn func() FieldUnmarshaler)
}

// FieldDecoder provides the interface for unmarshaling values from a type. The
// value is retrieved from the location referenced to by the Target. The field
// name that the value is retrieved from is k.
type FieldDecoder interface {
	Get(t Target, k string, fn func(v FieldValue), meta Metadata)
	GetList(t Target, k string, fn func(n int, ld ListDecoder), meta Metadata)
	GetMap(t Target, k string, fn func(ks []string, md MapDecoder), meta Metadata)
	GetFields(t Target, k string, fn func() FieldUnmarshaler, meta Metadata)
}

// DecodeBool converts a FieldValue into a bool pointer, updating the value
// pointed to by the input.
func DecodeBool(vp **bool) func(FieldValue) {
	return func(v FieldValue) {
		*vp = new(bool)
		**vp = v.(bool)
	}
}

// DecodeString converts a FieldValue into a string pointer, updating the value
// pointed to by the input.
func DecodeString(vp **string) func(FieldValue) {
	return func(v FieldValue) {
		*vp = new(string)
		**vp = v.(string)
	}
}

// DecodeInt64 converts a FieldValue into an int64 pointer, updating the value
// pointed to by the input.
func DecodeInt64(vp **int64) func(FieldValue) {
	return func(v FieldValue) {
		*vp = new(int64)
		**vp = v.(int64)
	}
}

// DecodeFloat64 converts a FieldValue into a float64 pointer, updating the value
// pointed to by the input.
func DecodeFloat64(vp **float64) func(FieldValue) {
	return func(v FieldValue) {
		*vp = new(float64)
		**vp = v.(float64)
	}
}

// DecodeTime converts a FieldValue into a time pointer, updating the value
// pointed to by the input.
func DecodeTime(vp **time.Time) func(FieldValue) {
	return func(v FieldValue) {
		*vp = new(time.Time)
		**vp = v.(time.Time)
	}
}

// DecodeBytes converts a FieldValue into a bytes slice, updating the value
// pointed to by the input.
func DecodeBytes(vp *[]byte) func(FieldValue) {
	return func(v FieldValue) {
		*vp = v.([]byte)
	}
}

// DecodeReadCloser converts a FieldValue into an io.ReadCloser, updating the value
// pointed to by the input.
func DecodeReadCloser(vp *io.ReadCloser) func(FieldValue) {
	return func(v FieldValue) {
		*vp = v.(io.ReadCloser)
	}
}

// DecodeReadSeeker converts a FieldValue into an io.ReadCloser, updating the value
// pointed to by the input.
func DecodeReadSeeker(vp *io.ReadSeeker) func(FieldValue) {
	return func(v FieldValue) {
		*vp = v.(io.ReadSeeker)
	}
}
