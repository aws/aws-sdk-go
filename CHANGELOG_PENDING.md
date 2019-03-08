### SDK Features
*  `service/kinesis`: Enable support for SubscribeToStream API operation ([#2402](https://github.com/aws/aws-sdk-go/pull/2402))
  * Adds support for Kinesis's SubscribeToStream API operation. The API operation response type, `SubscribeToStreamOutput` member, EventStream has a method `Events` which returns a channel to read Kinesis record events from.

### SDK Enhancements

### SDK Bugs
