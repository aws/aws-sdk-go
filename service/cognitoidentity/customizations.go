package cognitoidentity

import "github.com/aws/aws-sdk-go/aws/service"

func init() {
	initRequest = func(r *service.Request) {
		switch r.Operation.Name {
		case opGetOpenIdToken, opGetId, opGetCredentialsForIdentity:
			r.Handlers.Sign.Clear() // these operations are unsigned
		}
	}
}
