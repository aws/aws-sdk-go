### SDK Features

### SDK Enhancements

### SDK Bugs
* Fix generated errors for some JSON APIs not including a message ([#3088](https://github.com/aws/aws-sdk-go/issues/3088))
  * Fixes the SDK's generated errors to all include the `Message` member regardless if it was modeled on the error shape. This fixes the bug identified in #3088 where some JSON errors were not modeled with the Message member.
