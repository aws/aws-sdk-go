### SDK Features
* `service/dynamodb/dynamodbattribute`: Support has been added for empty string and byte values.
  * `Encoder` has added two new configuration options for controlling whether empty string and byte values are sent as null or empty.
    * `NullEmptyString`: Whether string values that are empty will be sent as null (default: `true`).
    * `NullEmptyByteSlice`: Whether byte slice that are empty will be sent as null (default: `true`).
    * The default value for these options retrains the existing behavior of the SDK in prior releases.

### SDK Enhancements

### SDK Bugs
