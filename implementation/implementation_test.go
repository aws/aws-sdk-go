package implementation

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

type STSMockClient struct {
	stsiface.STSAPI
}

func (s *STSMockClient) AssumeRole(input *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error) {
	if *input.ExternalId == "eligezer" {
		return &sts.AssumeRoleOutput{}, fmt.Errorf("invalid external ID")
	}
	return &sts.AssumeRoleOutput{
		Credentials: &sts.Credentials{
			AccessKeyId: aws.String("ASDASKDSDASDAS"),
			SecretAccessKey: aws.String("SADASDQWDDASAS"),
			SessionToken: aws.String("ASJDHASKJDAS"),
			},
	}, nil
}

func (s *STSMockClient) AssumeRoleWithWebIdentity(input *sts.AssumeRoleWithWebIdentityInput) (*sts.AssumeRoleWithWebIdentityOutput, error) {
	if *input.RoleArn == "arn:aws:iam::12345678910:role/eligezer" {
		return &sts.AssumeRoleWithWebIdentityOutput{}, fmt.Errorf("invalid role ARN")
	}
	return &sts.AssumeRoleWithWebIdentityOutput{
		Credentials: &sts.Credentials{
			AccessKeyId: aws.String("ASDASKDSDASDAS"),
			SecretAccessKey: aws.String("SADASDQWDDASAS"),
			SessionToken: aws.String("ASJDHASKJDAS"),
		},
	}, nil
}


func TestAssumeRole(t *testing.T) {
	type args struct {
		role string
		externalID string
	}
	tests := []struct {
		name           string
		args           args
		wantErr        string
	}{
		{
			name:           "successful trused identity assume role",
			args:           args{role: "arn:aws:iam::12345678910:role/good-role", externalID: "waaaa"},
			wantErr:        "",
		},
		{
			name:           "unsuccessful trused identity assume role",
			args:           args{role: "arn:aws:iam::630441286508:role/bad-role", externalID: "eligezer"},
			wantErr:       "invalid external ID",
		},
		{
			name:           "successful web identity assume role",
			args:           args{role: "arn:aws:iam::630441286508:role/good-role", externalID: ""},
			wantErr:        "",
		},
		{
			name:           "unsuccessful web identity assume role",
			args:           args{role: "arn:aws:iam::12345678910:role/eligezer", externalID: ""},
			wantErr:        "invalid role ARN",
		},
	}
	file := "/tmp/lewl123"
	err := os.WriteFile(file, []byte(""), 0644)
	defer os.Remove(file)
	if err != nil {
		t.Error("unable to create temporary file")
	}
	os.Setenv("AWS_WEB_IDENTITY_TOKEN_FILE", file)

	client := &STSMockClient{}
	for _, tt := range tests {
		t.Run("test assumeRole(): "+tt.name, func(t *testing.T) {
			_, _, _, err := assumeRole(client, tt.args.role, tt.args.externalID)
			if tt.wantErr != "" {
				require.NotNil(t, err, tt.name)
				assert.Contains(t, err.Error(), tt.wantErr, tt.name)
			} else {
				require.Nil(t, err, tt.name)
			}
		})
	}
}

func TestDetectConnectionType(t *testing.T) {
	type args struct {
		awsCreds map[string]string
	}
	tests := []struct {
		name           string
		args           args
		credsType      string
	}{
		{
			name:           "user based credentials",
			args:           args{awsCreds: map[string]string{awsAccessKeyId: "ASDHASDSDHADHASD", awsSecretAccessKey: "ASDSADASDASAS", roleArn: "", externalID: ""}},
			credsType:      "userBased",
		},
		{
			name:           "role based credentials",
			args:           args{awsCreds: map[string]string{awsAccessKeyId: "", awsSecretAccessKey: "", roleArn: "arn:aws:iam::12345678910:role/test", externalID: "eligezer"}},
			credsType:      "roleBased",
		},
		{
			name:           "bad credentials",
			args:           args{awsCreds: map[string]string{awsAccessKeyId: "", awsSecretAccessKey: "", roleArn: "", externalID: ""}},
			credsType:      "",
		},
	}
	for _, tt := range tests {
		t.Run("test ActionExist(): "+tt.name, func(t *testing.T) {
			result, _, _ := detectConnectionType(tt.args.awsCreds)
			assert.Equal(t, tt.credsType, result)
		})
	}
}