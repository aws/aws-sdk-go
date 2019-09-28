### SDK Features

### SDK Enhancements

### SDK Bugs
* ` internal/ini`: Fix ini parser to handle empty values [#2860](https://github.com/aws/aws-sdk-go/pull/2860)
  * Fixes incorrect modifications to the previous token value of the skipper. Adds checks for cases where a skipped statement should be marked as complete and not be ignored.
  * Adds tests for nested and empty field value parsing, along with tests suggested in [#2801](https://github.com/aws/aws-sdk-go/pull/2801)
  * Fixes [#2800](https://github.com/aws/aws-sdk-go/issues/2800)
