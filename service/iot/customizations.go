package iot

import (
	"github.com/aws/aws-sdk-go/aws/client"
)

func init() {
	initClient = func(c *client.Client) {
		c.ClientInfo.SigningName = ServiceName
	}
}
