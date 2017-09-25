### SDK Features
* Add dep Go dependency management metadata files (#1544)
  * Adds the Go `dep` dependency management metadata files to the SDK.
  * Fixes [#1451](https://github.com/aws/aws-sdk-go/issues/1451)
  * Fixes [#634](https://github.com/aws/aws-sdk-go/issues/634)
* `service/dynamodb/expression`: Add expression building utility for DynamoDB ([#1527](https://github.com/aws/aws-sdk-go/pull/1527))
  * Adds a new package, expression, to the SDK providing builder utilities to create DynamoDB expressions safely taking advantage of type safety.
* `API Marshaler`: Add generated marshalers for RESTXML protocol ([#1409](https://github.com/aws/aws-sdk-go/pull/1409))
  * Updates the RESTXML protocol marshaler to use generated code instead of reflection for REST XML based services.
* `API Marshaler`: Add generated marshalers for RESTJSON protocol ([#1547](https://github.com/aws/aws-sdk-go/pull/1547))
  * Updates the RESTJSON protocol marshaler to use generated code instead of reflection for REST JSON based services.

### SDK Enhancements
* `private/protocol`: Update format of REST JSON and XMl benchmarks ([#1546](https://github.com/aws/aws-sdk-go/pull/1546))
  * Updates the format of the REST JSON and XML benchmarks to be readable. RESTJSON benchmarks were updated to more accurately bench building of the protocol.

### SDK Bugs
