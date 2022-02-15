### SDK Features
* `codegen`: Updates the SDK's code generation to stop supporting new API modeled JSONValue parameters. The SDK's JSONValue type is only compatible with JSON documents with a top level JSON Object. JSON Lists, Strings, Scalars, are not compatible. This prevents JSON Value working with some APIs such as Amazon Lex Runtime Service's operations. 
  * Related to [#4264](https://github.com/aws/aws-sdk-go/pull/4264) and [#4258](https://github.com/aws/aws-sdk-go/issues/4258)

### SDK Enhancements

### SDK Bugs
* `service/lexruntimeservice`: Introduces a breaking change for following parameters from a JSONValue to string type, because the SDKs JSONValue is not compatible with JSON documents of lists.
  * PostContentInput.ActiveContexts
  * PutContentOutput.AlternativeIntents
  * PutContentOutput.ActiveContexts
  * PutSessionOutput.ActiveContexts
  * Fixes [#4258](https://github.com/aws/aws-sdk-go/issues/4258)
