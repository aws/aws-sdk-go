package sts

import "github.com/aws/aws-sdk-go/aws/service"

func init() {
	initRequest = func(r *service.Request) {
		switch r.Operation.Name {
		case opAssumeRoleWithSAML, opAssumeRoleWithWebIdentity:
			r.Handlers.Sign.Clear() // these operations are unsigned
		}
	}
}
