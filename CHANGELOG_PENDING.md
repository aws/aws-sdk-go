### SDK Features

### SDK Enhancements

### SDK Bugs
* `service/s3`,`service/kinesis`: Fix streaming APIs' Err method closing stream ([#2882](https://github.com/aws/aws-sdk-go/pull/2882))
  * Fixes calling the Err method on SDK's Amazon Kinesis's SubscribeToShared and Amazon S3's SelectObjectContent response EventStream members closing the stream. This would cause unexpected read errors, or early termination of the streams. Only the Close method of the streaming members will close the streams.
  * Related to [#2769](https://github.com/aws/aws-sdk-go/issues/2769)

