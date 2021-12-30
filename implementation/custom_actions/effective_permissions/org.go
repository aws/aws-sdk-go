package effective_permissions

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go/service/organizations"
)

func (perm *EffectivePermissions) getAllPoliciesForOrg(accountID string) (map[string]interface{}, error) {
	targetId, err := perm.getOrgId()
	if err != nil {
		return nil, err
	}

	accountSCPS, err := perm.getSCPSForTarget(accountID)
	if err != nil {
		return nil, err
	}

	formattedAccountSCPS, err := formatStatements(accountSCPS)
	if err != nil {
		return nil, err
	}
	organizationSCPS, err := perm.getSCPSForTarget(*targetId)
	if err != nil {
		return nil, err
	}

	formattedOrganizationSCPS, err := formatStatements(organizationSCPS)
	if err != nil {
		return nil, err
	}
	allPoliciesNonFiltered := map[string]interface{}{
		"account_SCPS":      formattedAccountSCPS,
		"organization_SCPS": formattedOrganizationSCPS,
	}

	return allPoliciesNonFiltered, nil
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

func (perm *EffectivePermissions) GetOrgEffectivePermissions() (map[string]interface{}, error) {
	accountID, ok := perm.GetParameters()[accountIDKey]

	if !ok {
		return nil, fmt.Errorf("missing %s param", accountIDKey)
	}

	return perm.getAllPoliciesForOrg(accountID.(string))
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
	var policies []*organizations.Policy
	for _, describePolicy := range describePolicies {
		policies = append(policies, describePolicy.Policy)
	}

	return policies, nil
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
	var policyIds []string
	for _, v := range policies.Policies {
		policyIds = append(policyIds, *v.Id)
	}

	var policyDetails []*MyPolicyDetail
	SCPStatements, err := perm.getSCPSStatements(policyIds)
	if err != nil {
		return nil, err
	}

	for _, SCPs := range SCPStatements {
		if SCPs != nil {
			policyDetails = append(policyDetails, convertOrgPolicy(SCPs))
		}
	}

	return policyDetails, nil
}
