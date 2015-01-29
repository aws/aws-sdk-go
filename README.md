# AWS SDK for Go

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/awslabs/aws-sdk-go)
[![Build Status](https://img.shields.io/travis/awslabs/aws-sdk-go.svg)](https://travis-ci.org/awslabs/aws-sdk-go)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/awslabs/aws-sdk-go/blob/master/LICENSE)

aws-sdk-go is a set of clients for all Amazon Web Services APIs,
automatically generated from the JSON schemas shipped with
[botocore](http://github.com/boto/botocore).

It supports all known AWS services, and maps exactly to the documented
APIs, with some allowances for Go-specific idioms (e.g. `ID` vs. `Id`).

## Caution

It is currently **highly untested**, so please be patient and report any
bugs or problems you experience. The APIs may change radically without
much warning, so please vendor your dependencies w/ Godep or similar.

Please do not confuse this for a stable, feature-complete library.

## Installing

Let's say you want to use EC2:

    $ go get github.com/awslabs/aws-sdk-go/gen/ec2

## Using

```go
import "github.com/awslabs/aws-sdk-go/aws"
import "github.com/awslabs/aws-sdk-go/gen/ec2"

creds := aws.Creds(accessKey, secretKey, "")
cli := ec2.New(creds, "us-west-2", nil)
resp, err := cli.DescribeInstances(nil)
if err != nil {
    panic(err)
}
fmt.Println(resp.Reservations)
```

## Supported Services

 * AutoScaling
 * CloudFormation
 * CloudFront
 * CloudHSM
 * CloudSearch
 * CloudSearchdomain
 * CloudTrail
 * CloudWatch Metrics
 * CloudWatch Logs
 * CodeDeploy
 * Cognito Identity
 * Cognito Sync
 * Config
 * Data Pipeline
 * Direct Connect
 * DynamoDB
 * EC2
 * EC2 Container Service
 * Elasticache
 * Elastic Beanstalk
 * Elastic Transcoder
 * ELB
 * EMR
 * Glacier
 * IAM
 * Import/Export
 * Kinesis
 * Key Management Service
 * Lambda
 * OpsWorks
 * RDS
 * RedShift
 * Route53
 * Route53 Domains
 * S3
 * SimpleDB
 * Simple Email Service
 * SNS
 * SQS
 * Storage Gateway
 * STS
 * Support
 * SWF
