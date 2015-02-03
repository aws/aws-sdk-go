package query

import (
	"encoding/xml"
	"io"

	"github.com/awslabs/aws-sdk-go/aws"
)

func Unmarshal(r *aws.Request) {
	defer r.HTTPResponse.Body.Close()
	if r.DataFilled() {
		err := xml.NewDecoder(r.HTTPResponse.Body).Decode(r.Data)
		if err != nil && err != io.EOF {
			r.Error = err
			return
		}
	}
}
