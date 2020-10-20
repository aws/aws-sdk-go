### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/credentials`: Fixed a race condition checking if credentials are expired. ([#3448](https://github.com/aws/aws-sdk-go/issues/3448))
  * Fixes [#3524](https://github.com/aws/aws-sdk-go/issues/3524)
* `internal/ini`: Fixes ini file parsing for cases when Right Hand Value is missed in the last statement of the ini file ([#3596](https://github.com/aws/aws-sdk-go/pull/3596)) 
  * related to [#2800](https://github.com/aws/aws-sdk-go/issues/2800)
