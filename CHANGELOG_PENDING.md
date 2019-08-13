### SDK Features
* SDK code generation will no longer remove stutter from operations and type names in service API packages. New API operations and types will not have the service's name removed from them. The SDK was previously squashing API types if removing stutter from a type resulted in a name of a type that already existed. The existing type would be deleted. Only the renamed type remained. This has been fixed, and previously deleted types are now available.
  * `AWS Glue`'s `GlueTable` with `Table`.  The API's previously deleted `Table` is available as `TableData`.
  * `AWS IoT Events`'s `IotEventsAction` with `Action`. The previously deleted `Action` is available as `ActionData`.

### SDK Enhancements
* `service/kinesis`: Add support for retrying service specific API errors ([#2751](https://github.com/aws/aws-sdk-go/pull/2751)
  * Adds support for retrying the Kinesis API error, LimitExceededException.
  * Fixes [#1376](https://github.com/aws/aws-sdk-go/issues/1376)

### SDK Bugs
* `private/model/api`: Fix broken shape stutter rename during generation ([#2747](https://github.com/aws/aws-sdk-go/pull/2747))
  * Fixes the SDK's code generation incorrectly renaming types and operations. The code generation would incorrectly rename an API type by removing the service's name from the type's name. This was done without checking for if a type with the new name already existed. Causing the SDK to replace the existing type with the renamed one.
  * Fixes [#2741](https://github.com/aws/aws-sdk-go/issues/2741)
* `private/model/api`: Fix API doc being generated with wrong value ([#2748](https://github.com/aws/aws-sdk-go/pull/2748))
  * Fixes the SDK's generated API documentation for structure member being generated with the wrong documentation value when the member was included multiple times in the model doc-2.json file, but under different types.
