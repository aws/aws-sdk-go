### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/session`: Removed typed literal parsing for config, everything is now treated as a string until a numeric value is needed.
  * This resolves an issue where the contents of a profile would silently be dropped with certain values.
