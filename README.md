# aws-go

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

    $ go get github.com/stripe/aws-go/aws/...

## Using

```go
import "github.com/stripe/aws-go/aws"
import "github.com/stripe/aws-go/aws/gen/ec2"

creds := aws.Creds(accessKey, secretKey, "")
cli := ec2.New(creds, "us-west-2", nil)
resp, err := cli.DescribeInstances(ec2.DescribeInstancesRequest{})
if err != nil {
    panic(err)
}
fmt.Println(resp.Reservations)
```
