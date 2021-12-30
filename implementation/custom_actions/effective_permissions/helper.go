package effective_permissions

import (
	"strings"

	"github.com/aws/aws-sdk-go/service/iam"
)

func processTags(tags []*iam.Tag) map[string]string {
	outputTags := map[string]string{}
	for _, tag := range tags {
		outputTags[PrincipalTag+"/"+*tag.Key] = *tag.Value
	}
	return outputTags
}

// getUserAccountId extracts the account ID from the user arn.
func getUserAccountId(user *iam.UserDetail) string {
	return strings.Split(*user.Arn, ":")[4:5][0]
}

func getDefaultDocument(p *iam.ManagedPolicyDetail) *string {
	for _, policyVersion := range p.PolicyVersionList {
		if *policyVersion.VersionId == *p.DefaultVersionId {
			return policyVersion.Document
		}
	}
	return nil
}
