### SDK Features

### SDK Enhancements

### SDK Bugs

* `service/s3`: Populate X-Amz-Content-Sha256 header when using s3 outpost arn
  * Using an outpost ARN results in a different signing name from the resolved endpoint. This signing name was not included in the signer logic to indicate the `X-Amz-Content-Sha256` header should be added to the request which is required by S3.
