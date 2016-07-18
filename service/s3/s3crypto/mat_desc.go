package s3crypto

import (
	"encoding/json"
)

// MaterialDescription is used for describing the materials
// that are being encrypted.
type MaterialDescription interface {
	Set(string, string)
	Get(string) (string, bool)
	EncodeDescription() (string, error)
	DecodeDescription([]byte) error
	// TODO: Should we do this differently?
	GetData() map[string]*string
}

// JSONMatDesc will use json marshalers to encode and decode
// the material description
type JSONMatDesc struct {
	data map[string]*string
}

// NewJSONMatDesc returns a MaterialDescription while making the
// necessary instantiations.
func NewJSONMatDesc() MaterialDescription {
	md := &JSONMatDesc{}
	md.data = make(map[string]*string)
	return md
}

// Set associates a given key to a value.
//
// Example:
// matdesc := s3crypto.NewJSONMatDesc()
// matdesc.Set("foo", "bar")
func (md *JSONMatDesc) Set(key, value string) {
	md.data[key] = &value
}

// Get returns the given value associated with key. We also return
// a bool to signify whether or not that key exists
//
// Example:
// matdesc := s3crypto.NewJSONMatDesc()
// matdesc.Set("foo", "bar")
// str := matdesc.Get("foo")
func (md *JSONMatDesc) Get(key string) (string, bool) {
	v, ok := md.data[key]
	if ok {
		return *v, ok
	}
	return "", ok
}

// EncodeDescription returns the proper json encoded string
// Maybe need a better name here
//
// Example:
// matdesc := s3crypto.NewJSONMatDesc()
// matdesc.Set("foo", "bar")
// data, err := matdesc.EncodeDescription()
func (md *JSONMatDesc) EncodeDescription() (string, error) {
	v, err := json.Marshal(&md.data)
	return string(v), err
}

// DecodeDescription return the proper json decoded string
//
// Example:
// matdesc := s3crypto.NewJSONMatDesc()
// data := []byte("{\"foo\": \"bar\"}")
// err := matdesc.DecodeDescription(data)
func (md *JSONMatDesc) DecodeDescription(b []byte) error {
	return json.Unmarshal(b, &md.data)
}

// GetData is used to assign to
func (md *JSONMatDesc) GetData() map[string]*string {
	return md.data
}
