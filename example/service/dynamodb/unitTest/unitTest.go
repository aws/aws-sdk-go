// Example of how to unit test your code that uses DynamoDB without needing to
// pass a connector to every function that interacts with the database.
package unitTest

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// After instantiating ItemGetter, assign the DynamoDB connector like:
//	svc := dynamodb.DynamoDB(sess)
//	getter.DynamoDB = dynamodbiface.DynamoDBAPI(svc)
type ItemGetter struct {
	DynamoDB dynamodbiface.DynamoDBAPI
}

// Retrieve values from DynamoDB with a table containing entries like:
// {"id": "my primary key", "value": "valuable value"}
func (self *ItemGetter) Get(id string) (value string) {
	var input = &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String("my_table"),
		AttributesToGet: []*string{
			aws.String("value"),
		},
	}
	if output, err := self.DynamoDB.GetItem(input); err == nil {
		if _, ok := output.Item["value"]; ok {
			dynamodbattribute.Unmarshal(output.Item["value"], &value)
		}
	}
	return
}
