### SDK Features

### SDK Enhancements
* `aws/arn`: aws/arn: Package for parsing and producing ARNs ([#1463](https://github.com/aws/aws-sdk-go/pull/1463))
  * Adds the `arn` package for AWS ARN parsing and building. Use this package to build AWS ARNs for services such as outlined in the [documentation](http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).

### SDK Bugs
* `aws/signer/v4`: Correct V4 presign signature to include content sha25 in URL ([#1469](https://github.com/aws/aws-sdk-go/pull/1469))
  * Updates the V4 signer so that when a Presign is generated the `X-Amz-Content-Sha256` header is added to the query string instead of being required to be in the header. This allows you to generate presigned URLs for GET requests, e.g S3.GetObject that do not require additional headers to be set by the downstream users of the presigned URL.
  * Related To: [#1467](https://github.com/aws/aws-sdk-go/issues/1467)

