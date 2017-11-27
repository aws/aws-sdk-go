package rest

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/aws/aws-sdk-go/private/protocol"
)

// An Encoder provides encoding of REST URI path, query, and header components
// of an HTTP request. Can also encode a stream as the payload.
//
// Does not support SetFields.
type Encoder struct {
	req *http.Request

	path protocol.PathReplace

	query  url.Values
	header http.Header

	payload io.ReadSeeker

	err error
}

// NewEncoder creates a new encoder from the passed in request. All query and
// header values will be added on top of the request's existing values. Overwriting
// duplicate values.
func NewEncoder(req *http.Request) *Encoder {
	e := &Encoder{
		req: req,

		path:   protocol.NewPathReplace(req.URL.Path),
		query:  req.URL.Query(),
		header: req.Header,
	}

	return e
}

// Encode will return the request and body if one was set. If the body
// payload was not set the io.ReadSeeker will be nil.
//
// returns any error if one occured while encoding the API's parameters.
func (e *Encoder) Encode() (*http.Request, io.ReadSeeker, error) {
	if e.err != nil {
		return nil, nil, e.err
	}

	e.req.URL.Path, e.req.URL.RawPath = e.path.Encode()
	e.req.URL.RawQuery = e.query.Encode()
	e.req.Header = e.header

	return e.req, e.payload, nil
}

// SetValue will set a value to the header, path, query.
//
// If the request's method is GET all BodyTarget values will be written to
// the query string.
func (e *Encoder) SetValue(t protocol.Target, k string, v protocol.ValueMarshaler, meta protocol.Metadata) {
	if e.err != nil {
		return
	}

	var str string
	str, e.err = v.MarshalValue()
	if e.err != nil {
		return
	}

	switch t {
	case protocol.HeaderTarget:
		e.header.Set(k, str)
	case protocol.PathTarget:
		e.path.ReplaceElement(k, str)
	case protocol.QueryTarget:
		e.query.Set(k, str)
	case protocol.BodyTarget:
		if e.req.Method != "GET" {
			e.err = fmt.Errorf("body target not supported for rest non-GET methods %s, %s", t, k)
			return
		}
		e.query.Set(k, str)
	default:
		e.err = fmt.Errorf("unknown SetValue rest encode target, %s, %s", t, k)
	}
}

// SetStream will set the stream to the payload of the request.
func (e *Encoder) SetStream(t protocol.Target, k string, v protocol.StreamMarshaler, meta protocol.Metadata) {
	if e.err != nil {
		return
	}

	switch t {
	case protocol.PayloadTarget:
		e.payload, e.err = v.MarshalStream()
	default:
		e.err = fmt.Errorf("unknown SetStream rest encode target, %s, %s", t, k)
	}
}

// SetList will set the nested list values to the header or query.
func (e *Encoder) SetList(t protocol.Target, k string, fn func(le protocol.ListEncoder), meta protocol.Metadata) {
	if e.err != nil {
		return
	}

	switch t {
	case protocol.QueryTarget:
		nested := protocol.QueryListEncoder{Key: k, Query: e.query}
		fn(&nested)
		e.err = nested.Err
	case protocol.HeaderTarget:
		nested := protocol.HeaderListEncoder{Key: k, Header: e.header}
		fn(&nested)
		e.err = nested.Err
	default:
		e.err = fmt.Errorf("unknown SetList rest encode target, %s, %s", t, k)
	}
}

// SetMap will set the nested map values to the header or query.
func (e *Encoder) SetMap(t protocol.Target, k string, fn func(fe protocol.MapEncoder), meta protocol.Metadata) {
	if e.err != nil {
		return
	}

	switch t {
	case protocol.QueryTarget:
		nested := protocol.QueryMapEncoder{Query: e.query}
		fn(&nested)
		e.err = nested.Err
	case protocol.HeadersTarget:
		nested := protocol.HeaderMapEncoder{Prefix: k, Header: e.header}
		fn(&nested)
		e.err = nested.Err
	default:
		e.err = fmt.Errorf("unknown SetMap rest encode target, %s, %s", t, k)
	}
}

// SetFields is not supported for REST encoder.
func (e *Encoder) SetFields(t protocol.Target, k string, m protocol.FieldMarshaler, meta protocol.Metadata) {
	e.err = fmt.Errorf("rest encoder SetFields not supported, %s, %s", t, k)
}
