package restxml

import (
	"bytes"
	"encoding/xml"
	"io"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/protocol/query"
	"github.com/awslabs/aws-sdk-go/aws/protocol/rest"
	"github.com/awslabs/aws-sdk-go/internal/protocol/xml/xmlutil"
)

func Build(r *aws.Request) {
	rest.Build(r)

	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		var buf bytes.Buffer
		err := xmlutil.BuildXML(r.Params, xml.NewEncoder(&buf))
		if err != nil {
			r.Error = err
			return
		}
		r.SetBufferBody(buf.Bytes())
	}
}

func Unmarshal(r *aws.Request) {
	if t := rest.PayloadType(r.Data); t == "structure" || t == "" {
		defer r.HTTPResponse.Body.Close()
		err := xmlutil.UnmarshalXML(r.Data, xml.NewDecoder(r.HTTPResponse.Body))
		if err != nil && err != io.EOF {
			r.Error = err
			return
		}
	}
}

func UnmarshalMeta(r *aws.Request) {
	rest.Unmarshal(r)
}

func UnmarshalError(r *aws.Request) {
	query.UnmarshalError(r)
}
