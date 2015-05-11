package integration

import "github.com/awslabs/aws-sdk-go/aws"

const Imported = true

func init() {
	if aws.DefaultConfig.Region == "" {
		panic("AWS_REGION must be configured to run integration tests")
	}
}
