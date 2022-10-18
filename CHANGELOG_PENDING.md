### SDK Features

### SDK Enhancements
* `aws/session`: Modified config resolution strategy when `$HOME` or `%USERPROFILE%` environment variables are not set.
  * When the environment variables are not set, the SDK will attempt to determine the home directory using `user.Current()`. 

### SDK Bugs
