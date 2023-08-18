### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/credentials/ssocreds`: Modify sso token provider logic to handle possible nil val returned by CreateToken.
  * Fixes [4947](https://github.com/aws/aws-sdk-go/issues/4947)