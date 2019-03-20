### SDK Features
* `aws/credentials/stscreds`: Update StdinTokenProvider to prompt on stder ([#2481](https://github.com/aws/aws-sdk-go/pull/2481))
  * Updates the `stscreds` package default MFA token provider, `StdinTokenProvider`, to prompt on `stderr` instead of `stdout`. This is to make it possible to redirect/pipe output when using `StdinTokenProvider` and still seeing the prompt text.

### SDK Enhancements

### SDK Bugs
