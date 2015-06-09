# AWS SDK for Go

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/aws/aws-sdk-go)
[![Build Status](https://img.shields.io/travis/aws/aws-sdk-go.svg)](https://travis-ci.org/aws/aws-sdk-go)
[![Apache V2 License](http://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/aws/aws-sdk-go/blob/master/LICENSE.txt)

aws-sdk-go is the official AWS SDK for the Go programming language.

[**Check out the official Developer Preview announcement (New - June 3rd 2015)**](https://aws.amazon.com/blogs/aws/developer-preview-of-aws-sdk-for-go-is-now-available/)

## Caution

The SDK is currently in the process of being developed, and not everything
may be working fully yet. Please be patient and report any bugs or problems
you experience. The APIs may change radically without much warning, so please
vendor your dependencies with Godep or similar.

Please do not confuse this for a stable, feature-complete library.

Note that while most AWS protocols are currently supported, not all services
available in this package are implemented fully, as some require extra
customizations to work with the SDK. If you've encountered such a scenario,
please open a [GitHub issue](https://github.com/aws/aws-sdk-go/issues)
so we can track work for the service.

## Installing

Install your specific service package with the following `go get` command.
For example, EC2 support might be installed with:

    $ go get github.com/aws/aws-sdk-go/service/ec2

You can also install the entire SDK by installing the root package, including all of the SDK's dependancies:

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

## Using

To use a service in the SDK, create a service variable by calling the `New()`
function. Once you have a service, you can call API operations which each
return response data and a possible error.

To list a set of instance IDs from EC2, you could run:

```go
package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	// Create an EC2 service object in the "us-west-2" region
	// Note that you can also configure your region globally by
	// exporting the AWS_REGION environment variable
	svc := ec2.New(&aws.Config{Region: "us-west-2"})

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
			fmt.Println("    - Instance ID: ", *inst.InstanceID)
		}
	}
}
```

You can find more information and operations in our
[API documentation](http://godoc.org/github.com/aws/aws-sdk-go).

## License

This SDK is distributed under the
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0),
see LICENSE.txt and NOTICE.txt for more information.
