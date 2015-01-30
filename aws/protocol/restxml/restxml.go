package restxml

import (
	"encoding/xml"
	"io"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/protocol/rest"
)

func Build(r *aws.Request) {
	rest.Build(r)

	if m := rest.PayloadMember(r.Params, r.Operation.InPayload); m != nil {
		b, err := xml.Marshal(m)
		if err != nil {
			r.Error = err
			return
		}
		r.SetBufferBody(b)
	}
}

func Unmarshal(r *aws.Request) {
	rest.Unmarshal(r)

	if m := rest.PayloadMember(r.Data, r.Operation.OutPayload); m != nil {
		defer r.HTTPResponse.Body.Close()
		err := xml.NewDecoder(r.HTTPResponse.Body).Decode(r.Data)
		if err != nil && err != io.EOF {
			r.Error = err
			return
		}
	}
}
