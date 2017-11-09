package protocol

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
)

type EscapeMode int

const (
	NoEscape EscapeMode = iota
	Base64Escape
	QuotedEscape
)

// EncodeJSONValue marshals the value into a JSON string, and optionally base64
// encodes the string before returning it.
func EncodeJSONValue(v aws.JSONValue, escape EscapeMode) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	switch escape {
	case Base64Escape:
		return base64.StdEncoding.EncodeToString(b), nil
	case QuotedEscape:
		return strconv.Quote(string(b)), nil
	}

	return string(b), nil
}

// DecodeJSONValue will attempt to decode the string input as a JSONValue.
// Optionally decoding base64 the value first before JSON unmarshaling.
func DecodeJSONValue(v string, escape EscapeMode) (aws.JSONValue, error) {
	var b []byte
	var err error

	switch escape {
	case Base64Escape:
		b, err = base64.StdEncoding.DecodeString(v)
	case QuotedEscape:
		var u string
		u, err = strconv.Unquote(v)
		fmt.Println("JSONVAlue decode", u, v, err)
		b = []byte(u)
	default:
		b = []byte(v)
	}

	if err != nil {
		return nil, err
	}

	m := aws.JSONValue{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
