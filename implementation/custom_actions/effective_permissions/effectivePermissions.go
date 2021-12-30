package effective_permissions

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

type EffectivePermissions struct {
	parameters     map[string]interface{}
	accountDetails *iam.GetAccountAuthorizationDetailsOutput
}

func (perm *EffectivePermissions) GetParameters() map[string]interface{} {
	return perm.parameters
}

func NewEffectivePermissions(parameters map[string]interface{}) (*EffectivePermissions, error) {
	a := &EffectivePermissions{parameters: parameters}

	val, err := a.getIAMAccountAuthorizationDetails()
	if err != nil {
		return nil, err
	}

	a.accountDetails = val

	return a, nil
}

func (perm *EffectivePermissions) GetSession() (*session.Session, error) {
	// get the aws session from the parameters
	awsSession, ok := perm.GetParameters()[serviceTag].(session.Session)
	if !ok {
		return nil, errors.New("failed to get AWS service")
	}
	// refresh the credentials
	if _, err := awsSession.Config.Credentials.Get(); err != nil {
		return nil, err
	}
	return &awsSession, nil
}

// getIAMAccountAuthorizationDetails gets the raw account authorization details from IAM.
func (perm *EffectivePermissions) getIAMAccountAuthorizationDetails() (*iam.GetAccountAuthorizationDetailsOutput, error) {
	awsSession, err := perm.GetSession()
	if err != nil {
		return nil, err
	}
	client := GetIAMClient(awsSession)
	if client == nil {
		return nil, errors.New("failed to initialize client")
	}

	maxItems := int64(1000)

	details := iam.GetAccountAuthorizationDetailsOutput{}
	input := iam.GetAccountAuthorizationDetailsInput{
		MaxItems: &maxItems,
	}

	if err = client.GetAccountAuthorizationDetailsPages(&input, func(out *iam.GetAccountAuthorizationDetailsOutput, lastPage bool) (shouldContinue bool) {
		details.UserDetailList = append(details.UserDetailList, out.UserDetailList...)

		details.GroupDetailList = append(details.GroupDetailList, out.GroupDetailList...)

		details.RoleDetailList = append(details.RoleDetailList, out.RoleDetailList...)

		details.Policies = append(details.Policies, out.Policies...)

		return out.Marker != nil
	}); err != nil {
		return nil, err
	}

	return &details, nil
}
