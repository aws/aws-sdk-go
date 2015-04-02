# AWS SDK for Go

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/awslabs/aws-sdk-go)
[![Build Status](https://img.shields.io/travis/awslabs/aws-sdk-go.svg)](https://travis-ci.org/awslabs/aws-sdk-go)
[![Apache V2 License](http://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/awslabs/aws-sdk-go/blob/master/LICENSE.txt)

aws-sdk-go is the official AWS SDK for the Go programming language.

## Caution

The SDK is currently in the process of being developed, and not everything
may be working fully yet. Please be patient and report any bugs or problems
you experience. The APIs may change radically without much warning, so please
vendor your dependencies with Godep or similar.

Please do not confuse this for a stable, feature-complete library.

Note that while most AWS protocols are currently supported, not all services
available in this package are implemented fully, as some require extra
customizations to work with the SDK. If you've encountered such a scenario,
please open a [GitHub issue](https://github.com/awslabs/aws-sdk-go/issues)
so we can track work for the service.

## Installing

Install your specific service package with the following `go get` command.
For example, EC2 support might be installed with:

    $ go get github.com/awslabs/aws-sdk-go/service/ec2

You can also install the entire SDK by installing the root package:

    $ go get github.com/awslabs/aws-sdk-go

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

## Using

To use a service in the SDK, create a service variable by calling the `New()`
function. Once you have a service, you can call API operations which each
return response data and a possible error.

To list a set of instance IDs from EC2, you could run:

```go
package main

import (
	"fmt"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/ec2"
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
[API documentation](http://godoc.org/github.com/awslabs/aws-sdk-go).

## TODO

The following list of feature enhancements are planned or are currently
in the works:

* [ ] `sqs.SendMessage()`, `sqs.SendMessageBatch()`, `sqs.ReceiveMessage()`:
  validate MD5 checksums of message bodies in responses (#162).
* [ ] `cloudsearchdomain`: Add support for CloudSearchDomain client (#163).
* [ ] DynamoDB CRC checksum validation (#164).
* [ ] `glacier`: Add support for Glacier (#165).
* [ ] `s3`: Compute MD5 checksums for operations that require this parameter (#154).
* [ ] `s3`: Improve error messages for empty response payloads (#166).
* [ ] `s3`: Validate that SSL is used when SSE keys are provided (#167).
* [ ] `s3`: Support path-style bucket access for non-DNS-compatible bucket names (#168).
* [ ] `s3.CreateBucket()`: Auto-populate `LocationConstraint` with `Config.Region` (#169).
* [ ] Error if credentials cannot be resolved when sending request (#170).
* [ ] Error if region/endpoint cannot be resolved when sending request (#171).

## License

This SDK is distributed under the
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0),
see LICENSE.txt and NOTICE.txt for more information.
