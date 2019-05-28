### SDK Features

### SDK Enhancements
* `service/mediastoredata`: Add support for nonseekable io.Reader ([#2622](https://github.com/aws/aws-sdk-go/pull/2622))
  * Updates the SDK's documentation to clarify how you can use the SDK's `aws.ReadSeekCloser` utility function to wrap an io.Reader to be used with an API operation that allows streaming unsigned payload in the operation's request.
  * Adds example using ReadSeekCloser with AWS Elemental MediaStore Data's PutObject API operation.
* Update CI validation testing for Go module files ([#2626](https://github.com/aws/aws-sdk-go/pull/2626))
  * Suppress changes to the Go module definition files during CI code generation validation testing.

### SDK Bugs
* `service/pinpointemail`: Fix client unable to make API requests ([#2625](https://github.com/aws/aws-sdk-go/pull/2625))
  * Fixes the API client's code generation to ignore the `targetPrefix` modeled value. This value is not valid for the REST-JSON protocol.
  * Updates the SDK's code generation to ignore the `targetPrefix` for all protocols other than RPCJSON.
