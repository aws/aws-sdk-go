package custom_actions

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/organizations"
)

func GetOrganizationClient(awsSession *session.Session) *organizations.Organizations {
	return organizations.New(awsSession) // create a new client from the session.
}

func GetIAMClient(awsSession *session.Session) *iam.IAM {
	return iam.New(awsSession) // create a new client from the session.
}
