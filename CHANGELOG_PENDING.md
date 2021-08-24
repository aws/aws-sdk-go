### SDK Features

### SDK Enhancements

### SDK Bugs
* `private/protocol`: Add support for parsing RFC 3339 timestamp without trailing Z
  * Adds support for parsing RFC 3339 timestamp but without the `Z` character, nor UTC offset.
  * Related to [aws/aws-sdk-go-v2#1387](https://github.com/aws/aws-sdk-go-v2/issues/1387)
