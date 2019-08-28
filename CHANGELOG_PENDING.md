### SDK Features

### SDK Enhancements

### SDK Bugs
*	`aws/client`: The `Retry-After` duration specified in the request is now added to the Retry delay for throttled exception
     * Fixes broken test for retry delays for throttled exceptions [#2796](https://github.com/aws/aws-sdk-go/pull/2796)
  	 * Fixes [#2795](https://github.com/aws/aws-sdk-go/issues/2795)
