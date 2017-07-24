package protocol

import (
	"fmt"
	"net/http"
)

// HeaderMapEncoder builds a map valu
type HeaderMapEncoder struct {
	Prefix string
	Header http.Header
	Err    error
}

// MapSetValue adds a single value to the header.
func (e *HeaderMapEncoder) MapSetValue(k string, v ValueMarshaler) {
	if e.Err != nil {
		return
	}

	str, err := v.MarshalValue()
	if err != nil {
		e.Err = err
		return
	}

	if len(e.Prefix) > 0 {
		k = e.Prefix + k
	}

	e.Header.Set(k, str)
}

// MapSetList executes the passed in callback with a list encoder based on
// the context of this HeaderMapEncoder.
func (e *HeaderMapEncoder) MapSetList(k string, fn func(le ListEncoder)) {
	if e.Err != nil {
		return
	}

	if len(e.Prefix) > 0 {
		k = e.Prefix + k
	}

	nested := HeaderListEncoder{Key: k, Header: e.Header}
	fn(&nested)
	e.Err = nested.Err
}

// MapSetMap sets the header element with nested maps appending the
// passed in k to the prefix if one was set.
func (e *HeaderMapEncoder) MapSetMap(k string, fn func(me MapEncoder)) {
	if e.Err != nil {
		return
	}

	if len(e.Prefix) > 0 {
		k = e.Prefix + k
	}

	nested := HeaderMapEncoder{Prefix: k, Header: e.Header}
	fn(&nested)
	e.Err = nested.Err
}

// MapSetFields Is not implemented, query map of FieldMarshaler is undefined.
func (e *HeaderMapEncoder) MapSetFields(k string, m FieldMarshaler) {
	e.Err = fmt.Errorf("header map encoder MapSetFields not supported, %s", k)
}

// HeaderListEncoder will encode list values nested into a header key.
type HeaderListEncoder struct {
	Key    string
	Header http.Header
	Err    error
}

// ListAddValue encodes an individual list value into the header.
func (e *HeaderListEncoder) ListAddValue(v ValueMarshaler) {
	if e.Err != nil {
		return
	}

	str, err := v.MarshalValue()
	if err != nil {
		e.Err = err
		return
	}

	e.Header.Add(e.Key, str)
}

// ListAddList Is not implemented, header list of list is undefined.
func (e *HeaderListEncoder) ListAddList(fn func(ListEncoder)) {
	e.Err = fmt.Errorf("header list encoder ListAddList not supported, %s", e.Key)
}

// ListAddMap Is not implemented, header list of map is undefined.
func (e *HeaderListEncoder) ListAddMap(fn func(MapEncoder)) {
	e.Err = fmt.Errorf("header list encoder ListAddMap not supported, %s", e.Key)
}

// ListAddFields Is not implemented, query list of FieldMarshaler is undefined.
func (e *HeaderListEncoder) ListAddFields(m FieldMarshaler) {
	e.Err = fmt.Errorf("header list encoder ListAddFields not supported, %s", e.Key)
}
