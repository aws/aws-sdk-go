### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/client`: Update client logging to only log request headers when logging is enabled
  * Updates the SDK's base client to only log request headers when logging is enabled. This fixes an issue where event stream API operations would always log request headers. Regardless if logging was enabled or not.
