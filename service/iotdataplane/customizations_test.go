package iotdataplane_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/EMCECS/aws-sdk-go/aws"
	"github.com/EMCECS/aws-sdk-go/awstesting/unit"
	"github.com/EMCECS/aws-sdk-go/service/iotdataplane"
)

func TestRequireEndpointIfRegionProvided(t *testing.T) {
	svc := iotdataplane.New(unit.Session, &aws.Config{
		Region:                 aws.String("mock-region"),
		DisableParamValidation: aws.Bool(true),
	})
	req, _ := svc.GetThingShadowRequest(nil)
	err := req.Build()

	assert.Equal(t, "", svc.Endpoint)
	assert.Error(t, err)
	assert.Equal(t, aws.ErrMissingEndpoint, err)
}

func TestRequireEndpointIfNoRegionProvided(t *testing.T) {
	svc := iotdataplane.New(unit.Session, &aws.Config{
		DisableParamValidation: aws.Bool(true),
	})
	fmt.Println(svc.ClientInfo.SigningRegion)

	req, _ := svc.GetThingShadowRequest(nil)
	err := req.Build()

	assert.Equal(t, "", svc.Endpoint)
	assert.Error(t, err)
	assert.Equal(t, aws.ErrMissingEndpoint, err)
}

func TestRequireEndpointUsed(t *testing.T) {
	svc := iotdataplane.New(unit.Session, &aws.Config{
		Region:                 aws.String("mock-region"),
		DisableParamValidation: aws.Bool(true),
		Endpoint:               aws.String("https://endpoint"),
	})
	req, _ := svc.GetThingShadowRequest(nil)
	err := req.Build()

	assert.Equal(t, "https://endpoint", svc.Endpoint)
	assert.NoError(t, err)
}
