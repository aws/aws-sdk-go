### SDK Features

### SDK Enhancements
* `aws/credentials`: Add support for context when getting credentials.
  * Adds `GetWithContext` to `Credentials` that allows canceling getting the credentials if the context is canceled, or times out. This fixes an issue where API operations would ignore their provide context when waiting for credentials to refresh.
  * Related to [#3127](https://github.com/aws/aws-sdk-go/pull/3127).

### SDK Bugs
* `service/s3/s3manager`: Fix resource leak on UploadPart failures ([#3144](https://github.com/aws/aws-sdk-go/pull/3144))
