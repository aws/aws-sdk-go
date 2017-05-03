// +build integration

//Package codepipeline provides gucumber integration tests support.
package codepipeline

import (
	"github.com/EMCECS/aws-sdk-go/awstesting/integration/smoke"
	"github.com/EMCECS/aws-sdk-go/service/codepipeline"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@codepipeline", func() {
		gucumber.World["client"] = codepipeline.New(smoke.Session)
	})
}
