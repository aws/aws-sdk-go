### SDK Features

### SDK Enhancements

### SDK Bugs
* `private/model/api`: Update client ServiceName to be based on name of service for new services. ([#1997](https://github.com/aws/aws-sdk-go/pull/1997))
    * Fixes the SDK's `ServiceName` AWS service client package value to be unique based on the service name for new AWS services. Does not change exiting client packages.
