package protocol

import (
	"fmt"
	"net/url"
)

// QueryMapEncoder builds a query string.
type QueryMapEncoder struct {
	Prefix string
	Query  url.Values
	Err    error
}

// MapSetValue adds a single value to the query.
func (e *QueryMapEncoder) MapSetValue(k string, v ValueMarshaler) {
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

	e.Query.Add(k, str)
}

// MapSetList executes the passed in callback with a list encoder based on
// the context of this QueryMapEncoder.
func (e *QueryMapEncoder) MapSetList(k string, fn func(le ListEncoder)) {
	if e.Err != nil {
		return
	}

	if len(e.Prefix) > 0 {
		k = e.Prefix + k
	}

	nested := QueryListEncoder{Key: k, Query: e.Query}
	fn(&nested)
	e.Err = nested.Err
}

// MapSetMap sets the query string element with nested maps appending the
// passed in k to the prefix if one was set.
func (e *QueryMapEncoder) MapSetMap(k string, fn func(me MapEncoder)) {
	if e.Err != nil {
		return
	}

	if len(e.Prefix) > 0 {
		k = e.Prefix + k
	}

	nested := QueryMapEncoder{Prefix: k, Query: e.Query}
	e.Err = nested.Err
}

// MapSetFields Is not implemented, query map of map is undefined.
func (e *QueryMapEncoder) MapSetFields(k string, m FieldMarshaler) {
	e.Err = fmt.Errorf("query map encoder MapSetFields not supported, %s", e.Prefix)
}

// QueryListEncoder will encode list values nested into a query key.
type QueryListEncoder struct {
	Key   string
	Query url.Values
	Err   error
}

// ListAddValue encodes an individual list value into the querystring.
func (e *QueryListEncoder) ListAddValue(v ValueMarshaler) {
	if e.Err != nil {
		return
	}

	str, err := v.MarshalValue()
	if err != nil {
		e.Err = err
		return
	}

	e.Query.Add(e.Key, str)
}

// ListAddList Is not implemented, query list of list is undefined.
func (e *QueryListEncoder) ListAddList(fn func(le ListEncoder)) {
	e.Err = fmt.Errorf("query list encoder ListAddList not supported, %s", e.Key)
}

// ListAddMap Is not implemented, query list of map is undefined.
func (e *QueryListEncoder) ListAddMap(fn func(fe MapEncoder)) {
	e.Err = fmt.Errorf("query list encoder ListAddMap not supported, %s", e.Key)
}

// ListAddFields Is not implemented, query list of FieldMarshaler is undefined.
func (e *QueryListEncoder) ListAddFields(m FieldMarshaler) {
	e.Err = fmt.Errorf("query list encoder ListAddFields not supported, %s", e.Key)
}
