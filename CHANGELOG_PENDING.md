### SDK Features

### SDK Enhancements
* `codegen`: Add XXX_Values functions for getting slice of API enums by type.
  * Fixes [#3441](https://github.com/aws/aws-sdk-go/issues/3441) by adding a new XXX_Values function for each API enum type that returns a slice of enum values, e.g `DomainStatus_Values`.

### SDK Bugs
* `codegen`: Export event stream constructor for easier mocking ([#3473](https://github.com/aws/aws-sdk-go/pull/3473))
  * Fixes [#3412](https://github.com/aws/aws-sdk-go/issues/3412) by exporting the operation's EventStream type's constructor function so it can be used to fully initialize fully when mocking out behavior for API operations with event streams.
* `service/ec2`: Fix max retries with client customizations ([#3465](https://github.com/aws/aws-sdk-go/pull/3465))
  * Fixes [#3374](https://github.com/aws/aws-sdk-go/issues/3374) by correcting the EC2 API client's customization for ModifyNetworkInterfaceAttribute and AssignPrivateIpAddresses operations to use the aws.Config.MaxRetries value if set. Previously the API client's customizations would ignore MaxRetries specified in the SDK's aws.Config.MaxRetries field.
