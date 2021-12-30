package effective_permissions

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/service/iam"
)

func (perm *EffectivePermissions) getUserBaseValues() (*iam.UserDetail, error) {
	val, ok := perm.GetParameters()[userNameKey]
	if !ok {
		return nil, fmt.Errorf("missing %s param", userNameKey)
	}

	user := perm.getUserDetails(val.(string))
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

// getUserDetails finds the username in the infoMap.
func (perm *EffectivePermissions) getUserDetails(userName string) *iam.UserDetail {
	for _, userDetail := range perm.accountDetails.UserDetailList {
		if *userDetail.UserName == userName {
			return userDetail
		}
	}

	return nil
}

func (perm *EffectivePermissions) getUserTags(username *string) (map[string]string, error) {
	awsSession, err := perm.GetSession()
	if err != nil {
		return nil, err
	}

	client := GetIAMClient(awsSession)
	if client == nil {
		return nil, errors.New("cant initialize client")
	}
	maxItems := int64(1000)

	output := iam.ListUserTagsOutput{}
	input := iam.ListUserTagsInput{
		UserName: username,
		MaxItems: &maxItems,
	}

	err = client.ListUserTagsPages(&input, func(out *iam.ListUserTagsOutput, lastPage bool) (shouldContinue bool) {
		output.Tags = append(output.Tags, out.Tags...)
		return out.Marker != nil
	})
	if err != nil {
		return nil, err
	}

	return processTags(output.Tags), nil
}

func (perm *EffectivePermissions) getGroupAttachedPolicies(policies []*iam.ManagedPolicyDetail, obj *iam.GroupDetail) []*iam.ManagedPolicyDetail {
	var GroupAttachedPolicies []*iam.ManagedPolicyDetail

	for _, attachedPolicy := range obj.AttachedManagedPolicies {
		GroupAttachedPolicies = append(GroupAttachedPolicies, getPolicyFromArn(policies, attachedPolicy.PolicyArn))
	}

	return GroupAttachedPolicies
}

func (perm *EffectivePermissions) getGroupPolicies(groupName *string) []*iam.PolicyDetail {
	groupsData := perm.accountDetails.GroupDetailList
	policies := perm.accountDetails.Policies

	for _, group := range groupsData {
		if *group.GroupName == *groupName {
			groupInlinePolicies := group.GroupPolicyList
			for _, groupPolicy := range perm.getGroupAttachedPolicies(policies, group) {
				if document := getDefaultDocument(groupPolicy); document != nil {
					newGroupPolicy := iam.PolicyDetail{
						PolicyDocument: document,
						PolicyName:     groupPolicy.PolicyName,
					}
					groupInlinePolicies = append(groupInlinePolicies, &newGroupPolicy)
				}
			}
			return groupInlinePolicies
		}
	}
	return nil
}

func (perm *EffectivePermissions) getUserGroupsPolicies(user iam.UserDetail) []*MyPolicyDetail {
	groups := user.GroupList
	var groupPolicies []*iam.PolicyDetail

	for _, group := range groups {
		groupPolicies = append(groupPolicies, perm.getGroupPolicies(group)...)
	}

	return ConvertPolicyDetails(groupPolicies)
}

func (perm *EffectivePermissions) getAllPoliciesForUser(user iam.UserDetail) (map[string]interface{}, error) {
	tags, err := perm.getUserTags(user.UserName)
	if err != nil {
		return nil, err
	}

	inline, err := filterStatements(getUserInlinePolicies(user), tags)
	if err != nil {
		return nil, err
	}

	attached, err := filterStatements(getAttachedPolicies(perm.accountDetails.Policies, user.AttachedManagedPolicies), tags)
	if err != nil {
		return nil, err
	}

	group, err := filterStatements(perm.getUserGroupsPolicies(user), tags)
	if err != nil {
		return nil, err
	}

	allPolicies := map[string]interface{}{
		"inline":   inline,
		"attached": attached,
		"group":    group,
	}
	return allPolicies, nil
}

func (perm *EffectivePermissions) GetUserEffectivePermissions() (map[string]interface{}, error) {
	user, err := perm.getUserBaseValues()
	if err != nil {
		return nil, err
	}

	return perm.getAllPoliciesForUser(*user)
}
