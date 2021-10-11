### SDK Features
* Update SDK's serialization of REST-JSON API input and Content-Type
  * Updates the SDK's serialization of REST-JSON based API input parameters into HTTP request message payload, and Content-Type are set correctly. API operations with input structure members that are modeled to be serialized to the request payload will always have at least an empty JSON object serialized. Even if all members targeting the payload are nil. Also fixes REST-JSON serialization so that Content-Type is not sent if the input parameter has no members target the request payload.

### SDK Enhancements

### SDK Bugs
