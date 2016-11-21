package polly

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/awstesting/unit"

	"github.com/stretchr/testify/assert"
)

func TestRestGETStrategy(t *testing.T) {
	svc := New(unit.Session, &aws.Config{Region: aws.String("us-west-2")})
	r, _ := svc.SynthesizeSpeechRequest(nil)
	err := restGETStrategy(r)
	assert.NoError(t, err)
	assert.Equal(t, "GET", r.HTTPRequest.Method)
	assert.NotEqual(t, nil, r.Operation.PresignStrategy)
}
