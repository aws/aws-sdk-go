### SDK Features

### SDK Enhancements

### SDK Bugs
* `service/dynamodb/dynamodbattribute`: Fix string alias unmarshal.
  * Fixes #3983 by correcting the unmarshaler's decoding of AttributeValue number (N) parameter into type that is a string alias.
