### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/signer/v4`: Fix Signer not trimming header value spaces
  * Fixes the AWS Sigv4 signer to trim header value's whitespace when computing the canonical headers block of the string to sign.
