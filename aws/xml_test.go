package aws

import (
	"encoding/xml"
	"testing"
)

type XMLRequest struct {
	XMLName xml.Name `xml:"http://whatever Request"`

	Integer    IntegerValue `xml:",omitempty"`
	DangerZone string       `xml:"-"`
}

func (r *XMLRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return MarshalXML(r, e, start)
}

func TestMarshalingXML(t *testing.T) {
	r := &XMLRequest{
		Integer:    Integer(0),
		DangerZone: "a zone of danger",
	}

	out, err := xml.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	if v, want := string(out), `<Request xmlns="http://whatever"><Integer>0</Integer></Request>`; v != want {
		t.Errorf("XML was \n%s\n but expected \n%s", v, want)
	}
}
