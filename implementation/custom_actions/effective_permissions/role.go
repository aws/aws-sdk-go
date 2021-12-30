package effective_permissions

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go/service/iam"
)

func (perm *EffectivePermissions) getRoleBaseValues() (*iam.RoleDetail, error) {
	// get the ARN from the parameters
	arn, ok := perm.GetParameters()[roleNameKey]
	if !ok {
		return nil, fmt.Errorf("missing %s param", roleNameKey)
	}

	role := perm.getRoleDetails(arn.(string))
	if role == nil {
		return nil, fmt.Errorf("role not found")
	}

	return role, nil
}

// getUserDetails finds the role arn in the RoleDetailList.
func (perm *EffectivePermissions) getRoleDetails(roleArn string) *iam.RoleDetail {
	for _, roleDetail := range perm.accountDetails.RoleDetailList {
		if *roleDetail.Arn == roleArn {
			return roleDetail
		}
	}

	return nil
}

func (perm *EffectivePermissions) getRoleTags(roleName *string) (map[string]string, error) {
	awsSession, err := perm.GetSession()
	if err != nil {
		return nil, err
	}

	client := GetIAMClient(awsSession)
	if client == nil {
		return nil, errors.New("cant initialize client")
	}
	maxItems := int64(1000)

	input := iam.ListRoleTagsInput{
		RoleName: roleName,
		MaxItems: &maxItems,
	}

	output, err := client.ListRoleTags(&input)
	if err != nil {
		return nil, err
	}

	return processTags(output.Tags), nil
}

func (perm *EffectivePermissions) getAllPoliciesForRole(role *iam.RoleDetail) (map[string]interface{}, error) {
	tags, err := perm.getRoleTags(role.RoleName)
	if err != nil {
		return nil, err
	}

	inline, err := filterStatements(getRoleInlinePolicies(role), tags)
	if err != nil {
		return nil, err
	}

	attached, err := filterStatements(getAttachedPolicies(perm.accountDetails.Policies, role.AttachedManagedPolicies), tags)
	if err != nil {
		return nil, err
	}

	allPolicies := map[string]interface{}{
		"inline":   inline,
		"attached": attached,
	}
	return allPolicies, nil
}

func (perm *EffectivePermissions) GetRoleEffectivePermissions() (map[string]interface{}, error) {
	role, err := perm.getRoleBaseValues()
	if err != nil {
		return nil, err
	}

	return perm.getAllPoliciesForRole(role)
}
