package sts

import "github.com/awslabs/aws-sdk-go/aws"

func init() {
	initRequest = func(r *aws.Request) {
		switch r.Operation {
		case opAssumeRoleWithSAML, opAssumeRoleWithWebIdentity:
			r.Handlers.Sign.Init() // these operations are unsigned
		}
	}
}
