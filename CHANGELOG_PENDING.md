### SDK Features

### SDK Enhancements
* `private/protocol/json/jsonutil`: Use json.Decoder to decrease memory allocation ([#2115](https://github.com/aws/aws-sdk-go/pull/2115))
  * Updates the SDK's JSON protocol marshaler to use `json.Decoder` instead of `ioutil.ReadAll`. This reduces the memory unmarshaling JSON payloads by about 50%.
  * Fix [#2114](https://github.com/aws/aws-sdk-go/pull/2114)

### SDK Bugs
