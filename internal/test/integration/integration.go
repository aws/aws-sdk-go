// Package integration performs initialization and validation for integration
// tests.
package integration

import "github.com/awslabs/aws-sdk-go/aws"

// Imported is a marker to ensure that this package's init() function gets
// executed.
//
// To use this package, import it and add:
//
// 	 var _ = integration.Imported
const Imported = true

func init() {
	if aws.DefaultConfig.Region == "" {
		panic("AWS_REGION must be configured to run integration tests")
	}
}
