// +build integration

//Package autoscaling provides gucumber integration tests support.
package autoscaling

import (
	"github.com/EMCECS/aws-sdk-go/awstesting/integration/smoke"
	"github.com/EMCECS/aws-sdk-go/service/autoscaling"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@autoscaling", func() {
		gucumber.World["client"] = autoscaling.New(smoke.Session)
	})
}
