### SDK Features

### SDK Enhancements
* `service/s3`: Amazon S3 now supports AWS PrivateLink, providing direct access to S3 via a private endpoint within your virtual private network.

### SDK Bugs
* `aws/session`: Fixed a bug that prevented credentials from being sourced from the environment if the loaded shared config profile contained partial SSO configuration. ([#3769](https://github.com/aws/aws-sdk-go/pull/3769))
  * Fixes ([#3768](https://github.com/aws/aws-sdk-go/issues/3768))
