// +build example,go18

package main

import (
	"fmt"
	"os"
	"plugin"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/plugincreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Example application which loads a Go Plugin file, and uses the credential
// provider defined within the plugin to get credentials for making a S3
// request.
//
// Build:
//   go build -tags example -o usePlugin main.go
//
// Usage:
//   ./usePlugin <compiled plugin>
func main() {
	if len(os.Args) < 2 {
		exitErrorPrintf("Usage: usePlugin <compiled plugin>")
	}

	// Open plugin, and load it into the process.
	p, err := plugin.Open(os.Args[1])
	if err != nil {
		exitErrorPrintf("failed to open plugin, %s, %v", os.Args[1], err)
	}

	// Create a new Credentials value which will source the provider's Retrieve
	// and IsExpired functions from the plugin.
	creds, err := plugincreds.NewCredentials(p)
	if err != nil {
		exitErrorPrintf("failed to load plugin provider, %v", err)
	}

	// Example to configure a Session with the newly created credentials that
	// will be sourced using the plugin's functionality.
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: creds,
	}))

	svc := s3.New(sess)
	result, err := svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String("myBucket"),
		Key:    aws.String("myKey"),
	})

	fmt.Println(result, err)
}

func exitErrorPrintf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
