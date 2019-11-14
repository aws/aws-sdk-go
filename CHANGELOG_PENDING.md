### SDK Features

### SDK Enhancements
* `aws/endpoints`: Add support for regional S3 us-east-1 endpoint
  * Adds support for S3 configuring an SDK Amazon S3 client for the regional us-east-1 endpoint instead of the default global S3 endpoint.
  * Adds a new configuration option, `S3UsEast1RegionalEndpoint` which when set to `RegionalS3UsEast1Endpoint`, and region is `us-east-1` the S3 client will resolve the `us-east-1` regional endpoint, `s3.us-east-1.amazonaws.com` instead of the global S3 endpoint, `s3.amazonaws.com`. The SDK defaults to the current global S3 endpoint resolution for backwards compatibility.
  * Opt-in to the `us-east-1` regional endpoint via the SDK's Config, environment variable, `AWS_S3_US_EAST_1_REGIONAL_ENDPOINT=regional`, or shared config option, `s3_us_east_1_regional_endpoint=regional`.
  * Note the SDK does not support the shared configuration file by default.  You must opt-in to that behavior via Session Option `SharedConfigState`, or `AWS_SDK_LOAD_CONFIG=true` environment variable.

### SDK Bugs
