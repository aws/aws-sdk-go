### SDK Features

### SDK Enhancements
* Fixup SDK source formating, test error checking, and simplify type conervsions
  * [#2703](https://github.com/aws/aws-sdk-go/pull/2703), [#2704](https://github.com/aws/aws-sdk-go/pull/2704), [#2705](https://github.com/aws/aws-sdk-go/pull/2705), [#2706](https://github.com/aws/aws-sdk-go/pull/2706), [#2707](https://github.com/aws/aws-sdk-go/pull/2707), [#2708](https://github.com/aws/aws-sdk-go/pull/2708)

### SDK Bugs
* `aws/request`: Fix SDK error checking when seeking readers ([#2696](https://github.com/aws/aws-sdk-go/pull/2696))
  * Fixes the SDK handling of seeking a reader to ensure errors are not lost, and are bubbled up.
  * In several places the SDK ignored Seek errors when attempting to determine a reader's length, or rewinding the reader for retry attempts.
  * Related to [#2525](https://github.com/aws/aws-sdk-go/issues/2525)
