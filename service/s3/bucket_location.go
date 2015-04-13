package s3

import (
	"io/ioutil"
	"regexp"

	"github.com/awslabs/aws-sdk-go/aws"
)

var reBucketLocation = regexp.MustCompile(`>([^<>]+)<\/Location`)

func buildGetBucketLocation(r *aws.Request) {
	if r.DataFilled() {
		out := r.Data.(*GetBucketLocationOutput)
		b, err := ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			r.Error = err
			return
		}

		match := reBucketLocation.FindSubmatch(b)
		if len(match) > 1 {
			loc := string(match[1])
			out.LocationConstraint = &loc
		}
	}
}
