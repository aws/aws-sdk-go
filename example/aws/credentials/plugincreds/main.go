// +build example,go18

package main

import (
	"fmt"
	"plugin"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/plugincreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Example application which loads a Go Plugin file, and uses the credential
// provider defined within the plugin to get credentials for making a S3
// request.
//
// Build:
//   go build -tags example -o usePlugin usePlugin.go
//
// Usage:
//   ./usePlugin <compiled plugin file>
func main() {
	// Open plugin, and load it into the process.
	p, err := plugin.Open("somefile.so")
	if err != nil {
		return nil, err
	}

	// Create a new Credentials value which will source the provider's Retrieve
	// and IsExpired functions from the plugin.
	creds, err := plugincreds.NewCredentials(p)
	if err != nil {
		return nil, err
	}

	// Example to configure a Session with the newly created credentials that
	// will be sourced using the plugin's functionality.
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: creds,
	}))

	svc := s3.New(sess)

	result, err := svc.HeadObject(&s3.HeadObject{
		Bucket: aws.String("myBucket"),
		Key:    aws.String("myKey"),
	})

	fmt.Println(result, err)
}
