### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/session`: Modify resolving sso credential logic to fix stack overflow bug while configuring shared config/profile via env var.
  * Fixes [4912](https://github.com/aws/aws-sdk-go/issues/4912)