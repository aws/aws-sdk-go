### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/signer/v4`: Fix Signing Unordered Multi Value Query Parameters ([#1491](https://github.com/aws/aws-sdk-go/pull/1491))
  * Removes sorting of query string values when calculating v4 signing as this is not part of the spec. The spec only requires the keys, not values, to be sorted which is achieved by Query.Encode().
