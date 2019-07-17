// +build go1.7

package stscreds

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/sts"
)

type mockSTS struct {
	*sts.STS
	AssumeRoleWithWebIdentityFn func(input *sts.AssumeRoleWithWebIdentityInput) (*sts.AssumeRoleWithWebIdentityOutput, error)
}

func (m *mockSTS) AssumeRoleWithWebIdentity(input *sts.AssumeRoleWithWebIdentityInput) (*sts.AssumeRoleWithWebIdentityOutput, error) {
	if m.AssumeRoleWithWebIdentityFn != nil {
		return m.AssumeRoleWithWebIdentityFn(input)
	}

	return nil, nil
}

func TestWebIdentityProviderRetrieve(t *testing.T) {
	now = func() time.Time {
		return time.Time{}
	}

	cases := []struct {
		name              string
		mockSTS           *mockSTS
		roleARN           string
		tokenFilepath     string
		sessionName       string
		expectedError     error
		expectedCredValue credentials.Value
	}{
		{
			name:          "session name case",
			roleARN:       "arn",
			tokenFilepath: "testdata/token.jwt",
			sessionName:   "foo",
			mockSTS: &mockSTS{
				AssumeRoleWithWebIdentityFn: func(input *sts.AssumeRoleWithWebIdentityInput) (*sts.AssumeRoleWithWebIdentityOutput, error) {
					if e, a := "foo", *input.RoleSessionName; !reflect.DeepEqual(e, a) {
						t.Errorf("expected %v, but received %v", e, a)
					}

					return &sts.AssumeRoleWithWebIdentityOutput{
						Credentials: &sts.Credentials{
							Expiration:      aws.Time(time.Now()),
							AccessKeyId:     aws.String("access-key-id"),
							SecretAccessKey: aws.String("secret-access-key"),
							SessionToken:    aws.String("session-token"),
						},
					}, nil
				},
			},
			expectedCredValue: credentials.Value{
				AccessKeyID:     "access-key-id",
				SecretAccessKey: "secret-access-key",
				SessionToken:    "session-token",
				ProviderName:    WebIdentityProviderName,
			},
		},
		{
			name:          "valid case",
			roleARN:       "arn",
			tokenFilepath: "testdata/token.jwt",
			mockSTS: &mockSTS{
				AssumeRoleWithWebIdentityFn: func(input *sts.AssumeRoleWithWebIdentityInput) (*sts.AssumeRoleWithWebIdentityOutput, error) {
					if e, a := fmt.Sprintf("%d", now().UnixNano()), *input.RoleSessionName; !reflect.DeepEqual(e, a) {
						t.Errorf("expected %v, but received %v", e, a)
					}

					return &sts.AssumeRoleWithWebIdentityOutput{
						Credentials: &sts.Credentials{
							Expiration:      aws.Time(time.Now()),
							AccessKeyId:     aws.String("access-key-id"),
							SecretAccessKey: aws.String("secret-access-key"),
							SessionToken:    aws.String("session-token"),
						},
					}, nil
				},
			},
			expectedCredValue: credentials.Value{
				AccessKeyID:     "access-key-id",
				SecretAccessKey: "secret-access-key",
				SessionToken:    "session-token",
				ProviderName:    WebIdentityProviderName,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p := NewWebIdentityRoleProvider(c.mockSTS, c.roleARN, c.sessionName, c.tokenFilepath)
			credValue, err := p.Retrieve()
			if e, a := c.expectedError, err; !reflect.DeepEqual(e, a) {
				t.Errorf("expected %v, but received %v", e, a)
			}

			if e, a := c.expectedCredValue, credValue; !reflect.DeepEqual(e, a) {
				t.Errorf("expected %v, but received %v", e, a)
			}
		})
	}
}
