package s3_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
)

var s3LocationTests = []struct {
	body string
	loc  string
}{
	{`<?xml version="1.0" encoding="UTF-8"?><LocationConstraint xmlns="http://s3.amazonaws.com/doc/2006-03-01/"/>`, ``},
	{`<?xml version="1.0" encoding="UTF-8"?><LocationConstraint xmlns="http://s3.amazonaws.com/doc/2006-03-01/">EU</LocationConstraint>`, `EU`},
}

func TestGetBucketLocation(t *testing.T) {
	for _, test := range s3LocationTests {
		s := s3.New(baseConfig)
		s.Handlers.Send.Clear()
		s.Handlers.Send.PushBack(func(r *aws.Request) {
			reader := ioutil.NopCloser(bytes.NewReader([]byte(test.body)))
			r.HTTPResponse = &http.Response{StatusCode: 200, Body: reader}
		})

		resp, err := s.GetBucketLocation(&s3.GetBucketLocationInput{Bucket: aws.String("bucket")})
		assert.NoError(t, err)
		if test.loc == "" {
			assert.Nil(t, resp.LocationConstraint)
		} else {
			assert.Equal(t, test.loc, *resp.LocationConstraint)
		}
	}
}
