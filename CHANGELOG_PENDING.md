### SDK Features
* `aws/request`: Fixes bug in WithSetRequestHeaders where the header key was added to the header map directly
  * Addresses an issue where the header keys being added were being added directly to the header map, and did not have the canonical header casing applied. This introduced bugs where instead of overwriting existing header key, it added another map entry.

### SDK Enhancements

### SDK Bugs
