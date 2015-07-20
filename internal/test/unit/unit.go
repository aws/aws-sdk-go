// Package unit performs initialization and validation for unit tests
package unit

import (
	"github.com/aws/aws-sdk-go/aws/awscfg"
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
	awscfg.DefaultConfig.Credentials =
		credentials.NewStaticCredentials("AKID", "SECRET", "SESSION")
	awscfg.DefaultConfig.WithRegion("mock-region")
}
