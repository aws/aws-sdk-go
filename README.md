# AWS SDK for Go

[![API Reference](http://img.shields.io/badge/api-reference-blue.svg)](http://docs.aws.amazon.com/sdk-for-go/api)
[![Join the chat at https://gitter.im/aws/aws-sdk-go](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/aws/aws-sdk-go?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Build Status](https://img.shields.io/travis/aws/aws-sdk-go.svg)](https://travis-ci.org/aws/aws-sdk-go)
[![Apache V2 License](http://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/aws/aws-sdk-go/blob/master/LICENSE.txt)

aws-sdk-go is the official AWS SDK for the Go programming language.

Checkout our [release notes](https://github.com/aws/aws-sdk-go/releases) for information about the latest bug fixes, updates, and features added to the SDK.

**Release [v0.10.0](http://aws.amazon.com/releasenotes/xxx) introduced a breaking change to the SDK.**

This change refactors the way the SDK provides configuration to service clients.  
To migrate to this change you'll need to add a session value to each
place a service cleint is created. Naively this can be just
`s3.New(session.New())` in place of `s3.New(nil)`. If aws.Config value
was passed in your code would become, `s3.New(session.New(), cfg)`. If
you used `default.DefaultConfig` you can create a new session, set the
config you used on the default config like `session.New(myCfg)` or
`sess.New(); sess.Config.LogLevel = aws.LogLevel(aws.Debug)`.

Examples:
```go
    // Create a session with the default config and request handlers.
    sess := session.New()

    // Create a session with a custom region
    sess := session.New(&aws.Config{Region: aws.String("us-east-1")})

    // Create a session, and add additional handlers for all service
    // clients created with the session to inherit.
    sess := session.New()
    sess.Handlers.Build.pushBack(func(r *request.Handlers) {
        // Log every request made and its payload
        logger.Println("Request: %s/%s, Payload: %s", r.ClientInfo.ServiceName, r.Operation, r.Params)
    })

    // Create a S3 client instance from a session
    sess := session.New()
    svc := s3.New(sess)

    // Create a S3 client with additional configuration
    svc := s3.New(sess, aws.NewConfig().WithRegion("us-west-2"))
```

## Caution

The SDK is currently in the process of being developed, and not everything
may be working fully yet. Please be patient and report any bugs or problems
you experience. The APIs may change radically without much warning, so please
vendor your dependencies with Godep or similar.

Please do not confuse this for a stable, feature-complete library.

If you've encountered any issues usign the SDK please open a [GitHub issue](https://github.com/aws/aws-sdk-go/issues).

## Installing

Install your specific service package with the following `go get` command.
For example, EC2 support might be installed with:

    $ go get github.com/aws/aws-sdk-go/service/ec2

You can also install the entire SDK by installing the root package, including all of the SDK's dependencies:

    $ go get -u github.com/aws/aws-sdk-go/...

## Configuring Credentials

Before using the SDK, ensure that you've configured credentials. The best
way to configure credentials on a development machine is to use the
`~/.aws/credentials` file, which might look like:

```
[default]
aws_access_key_id = AKID1234567890
aws_secret_access_key = MY-SECRET-KEY
```

You can learn more about the credentials file from this
[blog post](http://blogs.aws.amazon.com/security/post/Tx3D6U6WSFGOK2H/A-New-and-Standardized-Way-to-Manage-Credentials-in-the-AWS-SDKs).

Alternatively, you can set the following environment variables:

```
AWS_ACCESS_KEY_ID=AKID1234567890
AWS_SECRET_ACCESS_KEY=MY-SECRET-KEY
```

## Using the Go SDK

To use a service in the SDK, create a service variable by calling the `New()`
function. Once you have a service client, you can call API operations which each
return response data and a possible error.

To list a set of instance IDs from EC2, you could run:

```go
package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	// Create an EC2 service object in the "us-west-2" region
	// Note that you can also configure your region globally by
	// exporting the AWS_REGION environment variable
	svc := ec2.New(session.New(), aws.Config{Region: aws.String("us-west-2")})

	// Call the DescribeInstances Operation
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	// resp has all of the response data, pull out instance IDs:
	fmt.Println("> Number of reservation sets: ", len(resp.Reservations))
	for idx, res := range resp.Reservations {
		fmt.Println("  > Number of instances: ", len(res.Instances))
		for _, inst := range resp.Reservations[idx].Instances {
			fmt.Println("    - Instance ID: ", *inst.InstanceId)
		}
	}
}
```

You can find more information and operations in our
[API documentation](http://docs.aws.amazon.com/sdk-for-go/api/).

## License

This SDK is distributed under the
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0),
see LICENSE.txt and NOTICE.txt for more information.
