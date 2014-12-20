package aws_test

import (
	"encoding/xml"
	"testing"

	"github.com/stripe/aws-go/aws"
)

type XMLRequest struct {
	XMLName xml.Name `xml:"http://whatever Request"`

	Integer aws.IntegerValue `xml:",omitempty"`
}

func (r *XMLRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return aws.MarshalXML(r, e, start)
}

func TestMarshalingXML(t *testing.T) {
	r := &XMLRequest{
		Integer: aws.Integer(0),
	}

	out, err := xml.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	if v, want := string(out), `<Request xmlns="http://whatever"><Integer>0</Integer></Request>`; v != want {
		t.Errorf("XML was \n%s\n but expected \n%s", v, want)
	}
}
