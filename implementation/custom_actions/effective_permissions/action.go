package effective_permissions

import "github.com/aws/aws-sdk-go/aws/session"

type CustomAction interface {
	GetParameters() map[string]interface{}
	GetSession() *session.Session
}
