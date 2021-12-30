package effective_permissions

const (
	PrincipalTag = "aws:PrincipalTag"
	accountIDKey = "accountID"
	userNameKey  = "userName"
	roleNameKey  = "roleArn"
	filterSCP    = "SERVICE_CONTROL_POLICY"
	serviceTag   = "_Service"
	ExactMatch   = "StringEquals"
	NegatedMatch = "StringNotEquals"
)
