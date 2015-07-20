// Package integration performs initialization and validation for integration
// tests.
package integration

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws/awscfg"
	"github.com/aws/aws-sdk-go/aws/awsconv"
)

// Imported is a marker to ensure that this package's init() function gets
// executed.
//
// To use this package, import it and add:
//
// 	 var _ = integration.Imported
const Imported = true

func init() {
	if os.Getenv("DEBUG") != "" {
		awscfg.DefaultConfig.LogLevel = awscfg.LogLevel(awscfg.LogDebug)
	}
	if os.Getenv("DEBUG_SIGNING") != "" {
		awscfg.DefaultConfig.LogLevel = awscfg.LogLevel(awscfg.LogDebugWithSigning)
	}
	if os.Getenv("DEBUG_BODY") != "" {
		awscfg.DefaultConfig.LogLevel = awscfg.LogLevel(awscfg.LogDebugWithSigning | awscfg.LogDebugWithHTTPBody)
	}

	if awsconv.StringValue(awscfg.DefaultConfig.Region) == "" {
		panic("AWS_REGION must be configured to run integration tests")
	}
}

// UniqueID returns a unique UUID-like identifier for use in generating
// resources for integration tests.
func UniqueID() string {
	uuid := make([]byte, 16)
	io.ReadFull(rand.Reader, uuid)
	return fmt.Sprintf("%x", uuid)
}
