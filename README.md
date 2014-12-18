# aws-go

[![GoDoc](https://godoc.org/github.com/stripe/aws-go?status.svg)](http://godoc.org/github.com/stripe/aws-go)
[![Build Status](https://travis-ci.org/stripe/aws-go.svg?branch=master)](https://travis-ci.org/stripe/aws-go)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/stripe/aws-go/blob/master/LICENSE)

aws-go is a set of clients for all Amazon Web Services APIs,
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

    $ go get github.com/stripe/aws-go/gen/ec2

## Using

```go
import "github.com/stripe/aws-go/aws"
import "github.com/stripe/aws-go/gen/ec2"

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
 * Elasticache
 * Elastic Beanstalk
 * Elastic Transcoder
 * ELB
 * EMR
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
