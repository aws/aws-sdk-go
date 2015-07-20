// Package unit performs initialization and validation for unit tests
package unit

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

// Imported is a marker to ensure that this package's init() function gets
// executed.
//
// To use this package, import it and add:
//
// 	 var _ = unit.Imported
const Imported = true

func init() {
	// mock region and credentials
	aws.DefaultConfig.Credentials =
		credentials.NewStaticCredentials("AKID", "SECRET", "SESSION")
	aws.DefaultConfig.WithRegion("mock-region")
}
