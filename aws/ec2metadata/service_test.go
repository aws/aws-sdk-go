package ec2metadata_test

import (
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/stretchr/testify/assert"
)

func TestClientOverrideDefaultHTTPClientDialTimeout(t *testing.T) {
	svc := ec2metadata.New(session.New())

	assert.NotEqual(t, http.DefaultClient, svc.Config.HTTPClient)

	tr, ok := svc.Config.HTTPClient.Transport.(*http.Transport)
	assert.True(t, ok)
	assert.NotNil(t, tr)

	assert.NotNil(t, tr.Dial)
}

func TestClientNotOverrideDefaultHTTPClientDialTimeout(t *testing.T) {
	origClient := *http.DefaultClient
	http.DefaultClient.Transport = &http.Transport{}
	defer func() {
		http.DefaultClient = &origClient
	}()

	svc := ec2metadata.New(session.New())

	assert.Equal(t, http.DefaultClient, svc.Config.HTTPClient)

	tr, ok := svc.Config.HTTPClient.Transport.(*http.Transport)
	assert.True(t, ok)
	assert.NotNil(t, tr)

	assert.Nil(t, tr.Dial)
}
