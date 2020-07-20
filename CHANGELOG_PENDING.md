### SDK Features

### SDK Enhancements

### SDK Bugs
* `service/s3/s3crypto`: Fix client's temporary file buffer error on retry ([#3344](https://github.com/aws/aws-sdk-go/pull/3344))
  * Fixes the Crypto client's temporary file buffer cleanup returning an error when the request is retried.
