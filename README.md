# aws-go

aws-go is a set of clients for all Amazon Web Services APIs,
automatically generated from the JSON schemas shipped with
[botocore](http://github.com/boto/botocore).

It supports all known AWS services, and maps exactly to the documented
APIs, with some allowances for Go-specific idioms (e.g. `ID` vs. `Id`).

It is currently **highly untested**, so please be patient and report any
bugs or problems you experience.

## Installing

    $ go get github.com/stripe/aws-go/aws/gen/...

## Using

```go
import "github.com/stripe/aws-go/aws/gen/ec2"

cli := ec2.New(accessKey, secretKey, "us-west-2", nil)
resp, err := cli.DescribeInstances(ec2.DescribeInstancesRequest{})
if err != nil {
    panic(err)
}
fmt.Println(resp.Reservations)
```
