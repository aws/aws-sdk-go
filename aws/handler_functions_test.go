package aws

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidateCredentialsHandler(t *testing.T) {
	os.Clearenv()
	svc := NewService(&Config{
		Credentials: DetectCreds("AKID", "SECRET", ""),
	})
	svc.Handlers.Clear()
	svc.Handlers.Validate.PushBack(ValidateCredentialsHandler)

	req := NewRequest(svc, &Operation{Name: "Operation"}, nil, nil)
	err := req.Build()

	assert.NoError(t, err)
}

func TestValidateCredentialsHandlerError(t *testing.T) {
	os.Clearenv()
	creds, _ := ProfileCreds("example.ini", "missing", 10*time.Minute)

	svc := NewService(&Config{Credentials: creds})
	svc.Handlers.Clear()
	svc.Handlers.Validate.PushBack(ValidateCredentialsHandler)

	req := NewRequest(svc, &Operation{Name: "Operation"}, nil, nil)
	err := req.Build()

	assert.Error(t, err)
	assert.Equal(t, ErrMissingCredentials, err)

	// Now try without any credentials object at all
	svc = NewService(nil)
	svc.Handlers.Clear()
	svc.Handlers.Validate.PushBack(ValidateCredentialsHandler)

	req = NewRequest(svc, &Operation{Name: "Operation"}, nil, nil)
	err = req.Build()

	assert.Error(t, err)
	assert.Equal(t, ErrMissingCredentials, err)
}

func TestValidateEndpointHandler(t *testing.T) {
	os.Clearenv()
	svc := NewService(&Config{Region: "us-west-2"})
	svc.Handlers.Clear()
	svc.Handlers.Validate.PushBack(ValidateEndpointHandler)

	req := NewRequest(svc, &Operation{Name: "Operation"}, nil, nil)
	err := req.Build()

	assert.NoError(t, err)
}

func TestValidateEndpointHandlerError(t *testing.T) {
	os.Clearenv()
	svc := NewService(nil)
	svc.Handlers.Clear()
	svc.Handlers.Validate.PushBack(ValidateEndpointHandler)

	req := NewRequest(svc, &Operation{Name: "Operation"}, nil, nil)
	err := req.Build()

	assert.Error(t, err)
	assert.Equal(t, ErrMissingRegion, err)
}
