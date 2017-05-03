package apigateway

import (
	"github.com/EMCECS/aws-sdk-go/aws/client"
	"github.com/EMCECS/aws-sdk-go/aws/request"
)

func init() {
	initClient = func(c *client.Client) {
		c.Handlers.Build.PushBack(func(r *request.Request) {
			r.HTTPRequest.Header.Add("Accept", "application/json")
		})
	}
}
