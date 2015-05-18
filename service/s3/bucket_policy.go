package s3

import (
	"io/ioutil"

	"github.com/awslabs/aws-sdk-go/aws"
)

// unmarshalGetBucketPolicy extracts the policy string from the response body and
// sets it to the Output object's Policy.
//
// GetBucketPolicy comes back as a raw JSON string, not XML wrapped. So a custom
// unmarshal is needed to read the string from the body and set it to the Policy field.
func unmarshalGetBucketPolicy(r *aws.Request) {
	if !r.DataFilled() {
		return
	}

	b, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		r.Error = err
		return
	}
	r.HTTPResponse.Body.Close()

	out := r.Data.(*GetBucketPolicyOutput)

	policy := string(b)
	out.Policy = &policy
}
