package ec2query

import (
	"encoding/xml"
	"io"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/xml/xmlutil"
)

func Unmarshal(r *aws.Request) {
	defer r.HTTPResponse.Body.Close()
	if r.DataFilled() {
		err := xmlutil.UnmarshalXML(r.Data, xml.NewDecoder(r.HTTPResponse.Body))
		if err != nil && err != io.EOF {
			r.Error = err
			return
		}
	}
}
