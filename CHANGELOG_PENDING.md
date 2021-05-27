### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/request`: Fix handling of endpoints with trailing slashes
  * Fixes the SDK's handling of endpoint URLs that contain a trailing slash when the API operation's modeled path is suffixed. Also ensures any endpoint URL query string is squashed consistently. 
