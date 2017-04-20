package xmlutil

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

type mockBody struct {
	DoneErr error
	Body    io.Reader
}

func (m *mockBody) Read(p []byte) (int, error) {
	n, err := m.Body.Read(p)
	if (n == 0 || err == io.EOF) && m.DoneErr != nil {
		return n, m.DoneErr
	}

	return n, err
}

func TestUnmarshalXML_UnexpectedEOF(t *testing.T) {
	const partialXMLBody = `<?xml version="1.0" encoding="UTF-8"?>
	<First>first value</First>
	<Second>Second val`

	out := struct {
		First  *string `locationName:"First" type:"string"`
		Second *string `locationName:"Second" type:"string"`
	}{}

	expect := out
	expect.First = aws.String("first")
	expect.Second = aws.String("second")

	expectErr := fmt.Errorf("expected read error")

	body := &mockBody{
		DoneErr: expectErr,
		Body:    strings.NewReader(partialXMLBody),
	}

	decoder := xml.NewDecoder(body)
	err := UnmarshalXML(&out, decoder, "")

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := expectErr, err; e != a {
		t.Errorf("expect %v error in %v, but was not", e, a)
	}
}
