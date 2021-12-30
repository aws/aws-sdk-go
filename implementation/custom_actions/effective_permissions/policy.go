package effective_permissions

import (
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/organizations"
)

// MyPolicyDetail holds the data gathered from aws, this data is later processed and filtered
type MyPolicyDetail struct {
	Arn            *string
	PolicyDocument *string
	PolicyName     *string
}

// OutputPolicyDetail is returned by the filtering functions.
type OutputPolicyDetail struct {
	PolicyDocument []Statement
	PolicyName     *string
}

// ConvertPolicyDetails turns a slice of iam.PolicyDetail to slice of MyPolicyDetail
func ConvertPolicyDetails(policyList []*iam.PolicyDetail) []*MyPolicyDetail {
	var output []*MyPolicyDetail

	for _, policy := range policyList {
		output = append(output, convertIAMPolicy(policy))
	}

	return output
}

func NewPolicy(PolicyName *string, PolicyDocument *string, Arn *string) *MyPolicyDetail {
	return &MyPolicyDetail{
		PolicyName:     PolicyName,
		PolicyDocument: PolicyDocument,
		Arn:            Arn,
	}
}

// convertIAMPolicy turns iam.PolicyDetail to MyPolicyDetail
func convertIAMPolicy(p *iam.PolicyDetail) *MyPolicyDetail {
	return NewPolicy(p.PolicyName, p.PolicyDocument, nil)
}

// convertOrgPolicy turns organizations.Policy to MyPolicyDetail
func convertOrgPolicy(p *organizations.Policy) *MyPolicyDetail {
	return NewPolicy(p.PolicySummary.Name, p.Content, p.PolicySummary.Arn)
}

// getRoleInlinePolicies get the user's PolicyList
func getUserInlinePolicies(user iam.UserDetail) []*MyPolicyDetail {
	return ConvertPolicyDetails(user.UserPolicyList)
}

// getRoleInlinePolicies get the role's PolicyList
func getRoleInlinePolicies(roleDetail *iam.RoleDetail) []*MyPolicyDetail {
	return ConvertPolicyDetails(roleDetail.RolePolicyList)
}

// getPolicyFromArn finds the policy with the given arn.
func getPolicyFromArn(policies []*iam.ManagedPolicyDetail, arn *string) *iam.ManagedPolicyDetail {
	for _, p := range policies {
		if *p.Arn == *arn {
			return p
		}
	}

	return nil
}

func getAttachedPolicies(policies []*iam.ManagedPolicyDetail, attachedPolicies []*iam.AttachedPolicy) []*MyPolicyDetail {
	var AttachedPolicies []*MyPolicyDetail

	for _, AttachedPolicy := range attachedPolicies {
		policy := getPolicyFromArn(policies, AttachedPolicy.PolicyArn)
		document := getDefaultDocument(policy)

		if document != nil {
			AttachedPolicies = append(AttachedPolicies, NewPolicy(policy.PolicyName, document, AttachedPolicy.PolicyArn))
		}

	}

	return AttachedPolicies
}
