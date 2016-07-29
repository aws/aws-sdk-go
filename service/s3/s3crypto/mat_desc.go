package s3crypto

import (
	"encoding/json"
)

// MaterialDescription ...
type MaterialDescription map[string]*string

// EncodeDescription returns the proper json encoded string
// Maybe need a better name here
//
// Example:
//	matdesc := s3crypto.NewJSONMatDesc()
//	matdesc.Set("foo", "bar")
//	data, err := matdesc.EncodeDescription()
func (md *MaterialDescription) encodeDescription() ([]byte, error) {
	v, err := json.Marshal(&md)
	return v, err
}

// DecodeDescription return the proper json decoded string
//
// Example:
//	matdesc := s3crypto.NewJSONMatDesc()
//	data := []byte("{\"foo\": \"bar\"}")
//	err := matdesc.DecodeDescription(data)
func (md *MaterialDescription) decodeDescription(b []byte) error {
	return json.Unmarshal(b, &md)
}

// GetData is used to assign to
func (md MaterialDescription) GetData() map[string]*string {
	m := map[string]*string{}
	m = md
	return m
}
