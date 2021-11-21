package custom_actions

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/organizations"
)

// MyPolicyDetail holds the data gathered from aws, this data is later processed and filtered
type MyPolicyDetail struct {
	Arn            *string
	PolicyDocument *string
	PolicyName     *string
}

type EffectivePermissions struct {
	parameters map[string]interface{}
}

func (perm *EffectivePermissions) GetParameters() map[string]interface{} {
	return perm.parameters
}

func NewEffectivePermissions(parameters map[string]interface{}) *EffectivePermissions {
	return &EffectivePermissions{parameters: parameters}
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

// getInfoMap gets the raw account authorization details from IAM.
func (perm *EffectivePermissions) getInfoMap() (*iam.GetAccountAuthorizationDetailsOutput, error) {
	awsSession, err := perm.GetSession()
	if err != nil {
		return nil, err
	}
	client := GetIAMClient(awsSession)
	if client == nil {
		return nil, errors.New("cant initialize client")
	}

	maxItems := int64(1000)

	details := iam.GetAccountAuthorizationDetailsOutput{}
	input := iam.GetAccountAuthorizationDetailsInput{
		MaxItems: &maxItems,
	}

	err = client.GetAccountAuthorizationDetailsPages(&input, func(out *iam.GetAccountAuthorizationDetailsOutput, lastPage bool) (shouldContinue bool) {
		details.UserDetailList = append(details.UserDetailList, out.UserDetailList...)

		details.GroupDetailList = append(details.GroupDetailList, out.GroupDetailList...)

		details.RoleDetailList = append(details.RoleDetailList, out.RoleDetailList...)

		details.Policies = append(details.Policies, out.Policies...)

		return out.Marker != nil
	})

	if err != nil {
		return nil, err
	}

	return &details, nil
}

// getUserAsMap finds the username in the infoMap.
func (perm *EffectivePermissions) getUserAsMap(infoMap iam.GetAccountAuthorizationDetailsOutput, userName string) *iam.UserDetail {
	for _, userDetail := range infoMap.UserDetailList {
		if *userDetail.UserName == userName {
			return userDetail
		}
	}

	return nil
}

// getUserInlinePolicies get the users PolicyList
func (perm *EffectivePermissions) getUserInlinePolicies(user iam.UserDetail) []*MyPolicyDetail {
	var InlinePolicies []*MyPolicyDetail

	for _, policy := range user.UserPolicyList {
		InlinePolicies = append(InlinePolicies, perm.convertIAMPolicy(policy))
	}

	return InlinePolicies
}

// getPolicyFromGroupArn finds the policy with the given arn.
func (perm *EffectivePermissions) getPolicyFromGroupArn(policies []*iam.ManagedPolicyDetail, arn *string) *iam.ManagedPolicyDetail {
	for _, policy := range policies {
		if *policy.Arn == *arn {
			return policy
		}
	}

	return nil
}

// getPolicyFromUserArn finds the policy with the given arn.
func (perm *EffectivePermissions) getPolicyFromUserArn(policies []*iam.ManagedPolicyDetail, arn *string) *iam.ManagedPolicyDetail {
	for _, p := range policies {
		if *p.Arn == *arn {
			return p
		}
	}

	return nil
}

func (perm *EffectivePermissions) getUserAttachedPolicies(policies []*iam.ManagedPolicyDetail, obj iam.UserDetail) []*MyPolicyDetail {
	var AttachedPolicies []*MyPolicyDetail

	for _, AttachedPolicy := range obj.AttachedManagedPolicies {
		policy := perm.getPolicyFromUserArn(policies, AttachedPolicy.PolicyArn)
		document := perm.getCorrectDocumentVersion(policy)

		if document != nil {
			NewPolicy := &MyPolicyDetail{
				Arn:            nil,
				PolicyDocument: document,
				PolicyName:     policy.PolicyName,
			}

			AttachedPolicies = append(AttachedPolicies, NewPolicy)
		}

	}

	return AttachedPolicies
}

func (perm *EffectivePermissions) getGroupAttachedPolicies(policies []*iam.ManagedPolicyDetail, obj *iam.GroupDetail) []*iam.ManagedPolicyDetail {
	var GroupAttachedPolicies []*iam.ManagedPolicyDetail

	for _, attachedPolicy := range obj.AttachedManagedPolicies {
		GroupAttachedPolicies = append(GroupAttachedPolicies, perm.getPolicyFromGroupArn(policies, attachedPolicy.PolicyArn))
	}

	return GroupAttachedPolicies
}

func (perm *EffectivePermissions) getGroupPolicies(infoMap iam.GetAccountAuthorizationDetailsOutput, groupName *string) []*iam.PolicyDetail {
	groupsData := infoMap.GroupDetailList
	policies := infoMap.Policies

	for _, group := range groupsData {
		if *group.GroupName == *groupName {
			groupInlinePolicies := group.GroupPolicyList
			for _, groupPolicy := range perm.getGroupAttachedPolicies(policies, group) {
				if s := perm.getCorrectDocumentVersion(groupPolicy); s != nil {
					newGroupPolicy := iam.PolicyDetail{
						PolicyDocument: s,
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

func (perm *EffectivePermissions) getCorrectDocumentVersion(p *iam.ManagedPolicyDetail) *string {
	for _, policyVersion := range p.PolicyVersionList {
		if *policyVersion.VersionId == *p.DefaultVersionId {
			return policyVersion.Document
		}
	}
	return nil
}

func (perm *EffectivePermissions) getUserGroupsPolicies(infoMap iam.GetAccountAuthorizationDetailsOutput, user iam.UserDetail) []*MyPolicyDetail {
	groups := user.GroupList
	var groupPolicies []*iam.PolicyDetail

	for _, group := range groups {
		groupPolicies = append(groupPolicies, perm.getGroupPolicies(infoMap, group)...)
	}

	var output []*MyPolicyDetail
	for _, policy := range groupPolicies {
		output = append(output, perm.convertIAMPolicy(policy))
	}

	return output
}

func (perm *EffectivePermissions) getSCPSStatements(scpIds []string) ([]*organizations.Policy, error) {
	awsSession, err := perm.GetSession()
	if err != nil {
		return nil, err
	}

	client := GetOrganizationClient(awsSession)
	if client == nil {
		return nil, errors.New("cant initialize client")
	}

	var describePolicies []organizations.DescribePolicyOutput

	for _, id := range scpIds {
		out, err := client.DescribePolicy(&organizations.DescribePolicyInput{PolicyId: &id})
		if err != nil {
			return nil, err
		}
		describePolicies = append(describePolicies, *out)
	}

	// convert DescribePolicyOutput to Policy
	policies := []*organizations.Policy{}
	for _, describePolicy := range describePolicies {
		policies = append(policies, describePolicy.Policy)
	}

	return policies, nil
}

func (perm *EffectivePermissions) convertIAMPolicy(p *iam.PolicyDetail) *MyPolicyDetail {
	return &MyPolicyDetail{
		Arn:            nil,
		PolicyDocument: p.PolicyDocument,
		PolicyName:     p.PolicyName,
	}
}

func (perm *EffectivePermissions) convertOrgPolicy(p *organizations.Policy) *MyPolicyDetail {
	return &MyPolicyDetail{
		Arn:            p.PolicySummary.Arn,
		PolicyDocument: p.Content,
		PolicyName:     p.PolicySummary.Name,
	}
}

func (perm *EffectivePermissions) getSCPSForTarget(targetId string) ([]*MyPolicyDetail, error) {
	awsSession, err := perm.GetSession()
	if err != nil {
		return nil, err
	}

	client := GetOrganizationClient(awsSession)
	if client == nil {
		return nil, errors.New("cant initialize client")
	}

	scpFilter := filterSCP

	input := &organizations.ListPoliciesForTargetInput{TargetId: &targetId, Filter: &scpFilter}

	policies, err := client.ListPoliciesForTarget(input)
	if err != nil {
		return nil, err
	}
	policyIds := []string{}
	for _, v := range policies.Policies {
		policyIds = append(policyIds, *v.Id)
	}

	policyDetails := []*MyPolicyDetail{}
	SCPStatements, err := perm.getSCPSStatements(policyIds)
	if err != nil {
		return nil, err
	}

	for _, SCPs := range SCPStatements {
		if SCPs == nil {
			continue
		}
		policyDetails = append(policyDetails, perm.convertOrgPolicy(SCPs))
	}

	return policyDetails, nil
}

// getUserAccountId extracts the account ID from the user arn.
func (perm *EffectivePermissions) getUserAccountId(user iam.UserDetail) string {
	return strings.Split(*user.Arn, ":")[4:5][0]
}

// getOrgId gets the organization ID.
func (perm *EffectivePermissions) getOrgId() (*string, error) {
	organization, err := perm.getOrganizationDescription()
	if err != nil {
		return nil, err
	}
	return organization.Organization.MasterAccountId, nil
}

func (perm *EffectivePermissions) getOrganizationDescription() (*organizations.DescribeOrganizationOutput, error) {
	awsSession, err := perm.GetSession()
	if err != nil {
		return nil, err
	}

	client := GetOrganizationClient(awsSession)
	if client == nil {
		return nil, errors.New("cant initialize client")
	}

	return client.DescribeOrganization(&organizations.DescribeOrganizationInput{})
}

func (perm *EffectivePermissions) getAllPoliciesForUser(infoMap iam.GetAccountAuthorizationDetailsOutput, user iam.UserDetail) (map[string]interface{}, error) {
	tags, err := perm.getUserTags(user.UserName)
	if err != nil {
		return nil, err
	}

	inline, err := filterStatements(perm.getUserInlinePolicies(user), tags)
	if err != nil {
		return nil, err
	}

	attached, err := filterStatements(perm.getUserAttachedPolicies(infoMap.Policies, user), tags)
	if err != nil {
		return nil, err
	}

	group, err := filterStatements(perm.getUserGroupsPolicies(infoMap, user), tags)
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

func (perm *EffectivePermissions) getAllPoliciesForOrg(user iam.UserDetail) (map[string]interface{}, error) {
	targetId, err := perm.getOrgId()
	if err != nil {
		return nil, err
	}

	tags, err := perm.getUserTags(user.UserName)
	if err != nil {
		return nil, err
	}

	accountSCPS, err := perm.getSCPSForTarget(perm.getUserAccountId(user))
	if err != nil {
		return nil, err
	}

	filteredAccountSCPS, err := filterStatements(accountSCPS, tags)
	if err != nil {
		return nil, err
	}

	organizationSCPS, err := perm.getSCPSForTarget(*targetId)
	if err != nil {
		return nil, err
	}

	filteredOrganizationSCPS, err := filterStatements(organizationSCPS, tags)
	if err != nil {
		return nil, err
	}

	allPoliciesNonFiltered := map[string]interface{}{
		"account_SCPS":      filteredAccountSCPS,
		"organization_SCPS": filteredOrganizationSCPS,
	}

	return allPoliciesNonFiltered, nil
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

	outputTags := map[string]string{}
	for _, tag := range output.Tags {
		outputTags["aws:PrincipalTag/"+*tag.Key] = *tag.Value
	}

	return outputTags, nil
}

func (perm *EffectivePermissions) GetUserEffectivePermissions() (map[string]interface{}, error) {
	infoMap, user, err := perm.getBaseValues()
	if err != nil {
		return nil, err
	}

	return perm.getAllPoliciesForUser(*infoMap, *user)
}

func (perm *EffectivePermissions) getBaseValues() (*iam.GetAccountAuthorizationDetailsOutput, *iam.UserDetail, error) {
	infoMap, err := perm.getInfoMap()
	if err != nil {
		return nil, nil, err
	}

	val, ok := perm.GetParameters()[userNameKey]
	if !ok {
		return nil, nil, fmt.Errorf("missing %s param", userNameKey)
	}

	user := perm.getUserAsMap(*infoMap, val.(string))
	if user == nil {
		return nil, nil, fmt.Errorf("user not found")
	}

	return infoMap, user, nil
}

func (perm *EffectivePermissions) GetOrgEffectivePermissions() (map[string]interface{}, error) {
	_, user, err := perm.getBaseValues()
	if err != nil {
		return nil, err
	}

	return perm.getAllPoliciesForOrg(*user)
}
