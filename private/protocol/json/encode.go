package json

import (
	"bytes"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/private/protocol"
)

// An Encoder provides encoding of the AWS JSON protocol. This encoder will will
// write all content to JSON. Only supports body and payload targets.
type Encoder struct {
	encoder
	root bool
}

// NewEncoder creates a new encoder for encoding AWS JSON protocol. Only encodes
// fields into the JSON body, and error is returned if target is anything other
// than Body or Payload.
func NewEncoder() *Encoder {
	e := &Encoder{
		encoder: encoder{
			buf:      bytes.NewBuffer([]byte{'{'}),
			fieldBuf: &protocol.FieldBuffer{},
		},
		root: true,
	}

	return e
}

// Encode returns the encoded XMl reader. An error will be returned if one was
// encountered while building the JSON body.
func (e *Encoder) Encode() (io.ReadSeeker, error) {
	b, err := e.encode()
	if err != nil {
		return nil, err
	}

	if len(b) == 2 {
		// Account for first starting object in buffer
		return nil, nil
	}

	return bytes.NewReader(b), nil
}

// SetValue sets an individual value to the JSON body.
func (e *Encoder) SetValue(t protocol.Target, k string, v protocol.ValueMarshaler, meta protocol.Metadata) {
	e.writeSep()
	e.writeKey(k)
	e.writeValue(v)
}

// SetStream is not supported for JSON protocol marshaling.
func (e *Encoder) SetStream(t protocol.Target, k string, v protocol.StreamMarshaler, meta protocol.Metadata) {
	if e.err != nil {
		return
	}
	e.err = fmt.Errorf("json encoder SetStream not supported, %s, %s", t, k)
}

// SetList creates an JSON list and calls the passed in fn callback with a list encoder.
func (e *Encoder) SetList(t protocol.Target, k string, fn func(le protocol.ListEncoder), meta protocol.Metadata) {
	e.writeSep()
	e.writeKey(k)
	e.writeList(func(enc encoder) error {
		nested := listEncoder{encoder: enc}
		fn(&nested)
		return nested.err
	})
}

// SetMap creates an JSON map and calls the passed in fn callback with a map encoder.
func (e *Encoder) SetMap(t protocol.Target, k string, fn func(me protocol.MapEncoder), meta protocol.Metadata) {
	e.writeSep()
	e.writeKey(k)
	e.writeObject(func(enc encoder) error {
		nested := mapEncoder{encoder: enc}
		fn(&nested)
		return nested.err
	})
}

// SetFields sets the nested fields to the JSON body.
func (e *Encoder) SetFields(t protocol.Target, k string, m protocol.FieldMarshaler, meta protocol.Metadata) {
	if t == protocol.PayloadTarget {
		// Ignore payload key and only marshal body without wrapping in object first.
		nested := Encoder{
			encoder: encoder{
				buf:      e.encoder.buf,
				fieldBuf: e.encoder.fieldBuf,
			},
		}
		m.MarshalFields(&nested)
		e.err = nested.err
		return
	}

	e.writeSep()
	e.writeKey(k)
	e.writeObject(func(enc encoder) error {
		nested := Encoder{encoder: enc}
		m.MarshalFields(&nested)
		return nested.err
	})
}

// A listEncoder encodes elements within a list for the JSON encoder.
type listEncoder struct {
	encoder
}

// ListAddValue will add the value to the list.
func (e *listEncoder) ListAddValue(v protocol.ValueMarshaler) {
	e.writeSep()
	e.writeValue(v)
}

// ListAddList adds a list nested within another list.
func (e *listEncoder) ListAddList(fn func(le protocol.ListEncoder)) {
	e.writeSep()
	e.writeList(func(enc encoder) error {
		nested := listEncoder{encoder: enc}
		fn(&nested)
		return nested.err
	})
}

// ListAddMap adds a map nested within a list.
func (e *listEncoder) ListAddMap(fn func(me protocol.MapEncoder)) {
	e.writeSep()
	e.writeObject(func(enc encoder) error {
		nested := mapEncoder{encoder: enc}
		fn(&nested)
		return nested.err
	})
}

// ListAddFields will set the nested type's fields to the list.
func (e *listEncoder) ListAddFields(m protocol.FieldMarshaler) {
	e.writeSep()
	e.writeObject(func(enc encoder) error {
		nested := Encoder{encoder: enc}
		m.MarshalFields(&nested)
		return nested.err
	})
}

// A mapEncoder encodes key values pair map values for the JSON encoder.
type mapEncoder struct {
	encoder
}

// MapSetValue sets a map value.
func (e *mapEncoder) MapSetValue(k string, v protocol.ValueMarshaler) {
	e.writeSep()
	e.writeKey(k)
	e.writeValue(v)
}

// MapSetList encodes a list nested within the map.
func (e *mapEncoder) MapSetList(k string, fn func(le protocol.ListEncoder)) {
	e.writeSep()
	e.writeKey(k)
	e.writeList(func(enc encoder) error {
		nested := listEncoder{encoder: enc}
		fn(&nested)
		return nested.err
	})
}

// MapSetMap encodes a map nested within another map.
func (e *mapEncoder) MapSetMap(k string, fn func(me protocol.MapEncoder)) {
	e.writeSep()
	e.writeKey(k)
	e.writeObject(func(enc encoder) error {
		nested := mapEncoder{encoder: enc}
		fn(&nested)
		return nested.err
	})
}

// MapSetFields will set the nested type's fields under the map.
func (e *mapEncoder) MapSetFields(k string, m protocol.FieldMarshaler) {
	e.writeSep()
	e.writeKey(k)
	e.writeObject(func(enc encoder) error {
		nested := Encoder{encoder: enc}
		m.MarshalFields(&nested)
		return nested.err
	})
}

type encoder struct {
	buf      *bytes.Buffer
	fieldBuf *protocol.FieldBuffer
	started  bool
	err      error
}

func (e encoder) encode() ([]byte, error) {
	if e.err != nil {
		return nil, e.err
	}

	// Close the root object
	e.buf.WriteByte('}')

	return e.buf.Bytes(), nil
}

func (e *encoder) writeSep() {
	if e.started {
		e.buf.WriteByte(',')
	} else {
		e.started = true
	}

}
func (e *encoder) writeKey(k string) {
	e.buf.WriteByte('"')
	e.buf.WriteString(k) // TODO escape?
	e.buf.WriteByte('"')
	e.buf.WriteByte(':')
}

func (e *encoder) writeValue(v protocol.ValueMarshaler) {
	if e.err != nil {
		return
	}

	b, err := e.fieldBuf.GetValue(v)
	if err != nil {
		e.err = err
		return
	}

	var asStr bool
	switch v.(type) {
	case protocol.StringValue, protocol.BytesValue:
		asStr = true
	}

	if asStr {
		escapeStringBytes(e.buf, b)
	} else {
		e.buf.Write(b)
	}
}

func (e *encoder) writeList(fn func(encoder) error) {
	if e.err != nil {
		return
	}

	e.buf.WriteByte('[')
	e.err = fn(encoder{buf: e.buf, fieldBuf: e.fieldBuf})
	e.buf.WriteByte(']')
}

func (e *encoder) writeObject(fn func(encoder) error) {
	if e.err != nil {
		return
	}

	e.buf.WriteByte('{')
	e.err = fn(encoder{buf: e.buf, fieldBuf: e.fieldBuf})
	e.buf.WriteByte('}')
}
