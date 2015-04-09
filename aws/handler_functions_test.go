package aws

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEndpointHandler(t *testing.T) {
	os.Clearenv()
	svc := NewService(&Config{Region: "us-west-2"})
	svc.Handlers.Clear()
	svc.Handlers.Validate.PushBack(ValidateEndpointHandler)

	req := NewRequest(svc, &Operation{Name: "Operation"}, nil, nil)
	err := req.Build()

	assert.NoError(t, err)
}

func TestValidateEndpointHandlerErrorRegion(t *testing.T) {
	os.Clearenv()
	svc := NewService(nil)
	svc.Handlers.Clear()
	svc.Handlers.Validate.PushBack(ValidateEndpointHandler)

	req := NewRequest(svc, &Operation{Name: "Operation"}, nil, nil)
	err := req.Build()

	assert.Error(t, err)
	assert.Equal(t, ErrMissingRegion, err)
}

func TestValidateEndpointHandlerErrorEndpoint(t *testing.T) {
	os.Clearenv()
	svc := NewService(&Config{Endpoint: ""})
	svc.Handlers.Clear()
	svc.Handlers.Validate.PushBack(ValidateEndpointHandler)

	req := NewRequest(svc, &Operation{Name: "Operation"}, nil, nil)
	err := req.Build()

	assert.Error(t, err)
	assert.Equal(t, ErrMissingEndpoint, err)
}
