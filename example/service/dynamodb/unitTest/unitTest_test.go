// Unit tests for package unitTest.
package unitTest

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// A fakeDynamoDB instance. During testing, instatiate ItemGetter, then simply
// assign an instance of fakeDynamoDB to it.
type fakeDynamoDB struct {
	dynamodbiface.DynamoDBAPI
	payload map[string]string // Store expected return values
	err     error
}

// Mock GetItem such that the output returned carries values identical to input.
func (self *fakeDynamoDB) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
    output := new(dynamodb.GetItemOutput)
	output.Item = make(map[string]*dynamodb.AttributeValue)
	for key, value := range self.payload {
		output.Item[key] = &dynamodb.AttributeValue{
			S: aws.String(value),
		}
	}
	return output, self.err
}

func TestItemGetterGet(t *testing.T) {
    expected_key := "expected key"
    expected_value := "expected value"
    getter := new(ItemGetter)
	getter.DynamoDB = &fakeDynamoDB{
        payload: map[string]string{"id": expected_key, "value": expected_value},
	}
    if actual_value := getter.Get(expected_key); actual_value != expected_value {
		t.Errorf("Expected %q but got %q", expected_value, actual_value)
	}
}

// When DynamoDB.GetItem returns a non-nil error, expect an empty string.
func TestItemGetterGetFail(t *testing.T) {
    expected_key := "expected key"
    expected_value := "expected value"
    getter := new(ItemGetter)
	getter.DynamoDB = &fakeDynamoDB{
        payload: map[string]string{"id": expected_key, "value": expected_value},
		err:     errors.New("any error"),
	}
    if actual_value := getter.Get(expected_key); len(actual_value) > 0 {
		t.Errorf("Expected %q but got %q", expected_value, actual_value)
	}
}
