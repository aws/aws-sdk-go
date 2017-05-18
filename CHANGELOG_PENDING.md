### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/request`: Fix logging from reporting wrong retry request errors #1281
  * Fixes the SDK's retry request logging to report the the actual error that occurred, not a stubbed Unknown error message.
  * Fixes the SDK's response logger to not output the response log multiple times per retry.
