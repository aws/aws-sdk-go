// +build integration

package dynamodb_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
)

var dynamoDBClient *dynamodb.DynamoDB
var idAndNameTable string

func TestMain(m *testing.M) {
	var exitCode int
	createTables()
	defer os.Exit(exitCode)
	defer cleanup()

	exitCode = m.Run()
}

func TestGetAndPutItem(t *testing.T) {

	item := make(map[string]dynamodb.AttributeValue)
	item["Id"] = dynamodb.AttributeValue{
		N: aws.String("1"),
	}
	item["Name"] = dynamodb.AttributeValue{
		S: aws.String("Joe Smith"),
	}
	item["Age"] = dynamodb.AttributeValue{
		N: aws.String("40"),
	}
	item["ComplexShape"] = dynamodb.AttributeValue{
		M: map[string]dynamodb.AttributeValue{
			"List1": dynamodb.AttributeValue{
				M: map[string]dynamodb.AttributeValue{
					"Scores1": dynamodb.AttributeValue{
						L: []dynamodb.AttributeValue{
							dynamodb.AttributeValue{N: aws.String("15")},
							dynamodb.AttributeValue{N: aws.String("25")},
						},
					},
					"Scores2": dynamodb.AttributeValue{
						L: []dynamodb.AttributeValue{
							dynamodb.AttributeValue{N: aws.String("39")},
							dynamodb.AttributeValue{N: aws.String("49")},
						},
					},
				},
			},
		},
	}

	putItemInput := dynamodb.PutItemInput{
		TableName: &idAndNameTable,
		Item:      item,
		ReturnConsumedCapacity: aws.String("TOTAL"),
	}

	putItemOutput, err := dynamoDBClient.PutItem(&putItemInput)
	if err != nil {
		t.Fatal(err)
	}

	if *(putItemOutput.ConsumedCapacity.CapacityUnits) == 0.0 {
		t.Fatal("Failed to marshall consumed capacity")
	}

	key := make(map[string]dynamodb.AttributeValue)
	key["Id"] = dynamodb.AttributeValue{
		N: aws.String("1"),
	}
	key["Name"] = dynamodb.AttributeValue{
		S: aws.String("Joe Smith"),
	}
	getItemInput := dynamodb.GetItemInput{
		TableName: &idAndNameTable,
		Key:       key,
	}

	getItemOutput, err := dynamoDBClient.GetItem(&getItemInput)
	if err != nil {
		t.Fatal(err)
	}

	if len(getItemOutput.Item) != len(putItemInput.Item) {
		t.Fatal("Different number of attributes between get and put")
	}

	if *(getItemOutput.Item["Id"].N) != "1" {
		t.Fatal("Incorrect value for Id attribute: ", getItemOutput.Item["Id"].N)
	}

	if *(getItemOutput.Item["Name"].S) != "Joe Smith" {
		t.Fatal("Incorrect value for Name attribute: ", getItemOutput.Item["Name"].N)
	}

	if *(getItemOutput.Item["Age"].N) != "40" {
		t.Fatal("Incorrect value for Age attribute: ", getItemOutput.Item["Age"].N)
	}

	if *(getItemOutput.Item["ComplexShape"].M["List1"].M["Scores1"].L[0].N) != "15" {
		t.Fatal("Failed to find value in complex shape: ",
			*(getItemOutput.Item["ComplexShape"].M["List1"].M["Scores1"].L[0].N))
	}
}

func cleanup() {

	if idAndNameTable != "" {
		deleteTable(&idAndNameTable)
	}

}

func deleteTable(tableName *string) {

	input := dynamodb.DeleteTableInput{
		TableName: tableName,
	}

	dynamoDBClient.DeleteTable(&input)
}

func createTables() {

	idAndNameTable = "AWS_GO_ID_AND_NAME_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	dynamoDBClient = dynamodb.New(nil)

	attributes := make([]dynamodb.AttributeDefinition, 2, 2)
	attributes[0] = dynamodb.AttributeDefinition{
		AttributeName: aws.String("Id"),
		AttributeType: aws.String("N"),
	}
	attributes[1] = dynamodb.AttributeDefinition{
		AttributeName: aws.String("Name"),
		AttributeType: aws.String("S"),
	}

	provisionThroughput := dynamodb.ProvisionedThroughput{
		ReadCapacityUnits:  aws.Long(1),
		WriteCapacityUnits: aws.Long(1),
	}

	keySchema := make([]dynamodb.KeySchemaElement, 2, 2)
	keySchema[0] = dynamodb.KeySchemaElement{
		AttributeName: aws.String("Id"),
		KeyType:       aws.String("HASH"),
	}
	keySchema[1] = dynamodb.KeySchemaElement{
		AttributeName: aws.String("Name"),
		KeyType:       aws.String("RANGE"),
	}

	input := dynamodb.CreateTableInput{
		TableName:             &idAndNameTable,
		AttributeDefinitions:  attributes,
		ProvisionedThroughput: &provisionThroughput,
		KeySchema:             keySchema,
	}

	data, err := dynamoDBClient.CreateTable(&input)
	if err != nil {
		fmt.Errorf("Failed to create DynamoDB tables for integration tests: ", err)
		panic(err)
	}

	waitTillTableActive(data.TableDescription.TableName)
}

func waitTillTableActive(tableName *string) {

	describeInput := dynamodb.DescribeTableInput{
		TableName: tableName,
	}

	var describeOutput *dynamodb.DescribeTableOutput

	for describeOutput == nil || *(describeOutput.Table.TableStatus) != "ACTIVE" {
		time.Sleep(time.Second * 5)
		describeOutput, _ = dynamoDBClient.DescribeTable(&describeInput)
	}

	fmt.Println("Table ", *tableName, " is active.")
}
