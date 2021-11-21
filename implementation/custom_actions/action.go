package custom_actions

import "github.com/aws/aws-sdk-go/aws/session"

type CustomAction interface {
	GetParameters() map[string]interface{}
	GetSession() *session.Session
}
