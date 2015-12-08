// This program creates a table in dynamodb, puts an item (INSERT),
// Update the item (adding new attributes (columns))

package main

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main() {

	// Create a new instance of the DynamoDB service client
	svc := dynamodb.New(session.New(&aws.Config{Region: aws.String("us-west-2")}))

	tableName := "AWS_SDK_TEST_TABLE"

	// Use CreateTable API Operation to create a table on DynamoDB
	log.Println("Creating table: ", tableName)
	if _, err := svc.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Filename"),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Filename"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	}); err != nil {
		log.Println("Failed to create Amazon DynamoDB table,", err)
		return
	}
	log.Println("Succesfully created table:", tableName, "\n")

	// Now we will list the tables and will see if the table
	// that we created above is printed or not
	log.Println("Getting tablenames from DynamoDB")
	result, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Tables:")
	for i, table := range result.TableNames {
		log.Println("\t", i+1, *table)
	}
	log.Println()

	// If we create a table and immediately Put data, we get
	// a HTTP 400 error. So sleep for 5 seconds and try
	time.Sleep(5 * time.Second)

	// Let us create a dummy struct which we can Put (INSERT)
	// into the dynamodb table
	type file struct {
		Filename string
		Ctime    string
	}

	item := file{
		"1",
		time.Now().String(),
	}

	// Create attributevalues
	log.Println("Creating an item that could be put into DynamoDB")
	var av, keyMap map[string]*dynamodb.AttributeValue
	av, err = dynamodbattribute.ConvertToMap(item)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Putting the Item into DynamoDB")
	_, err = svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})
	if err != nil {
		log.Printf("unable to record result, %v", err)
		return
	}
	log.Println("Succesfully put the item:", item, "\n")

	// Let us search for the item that we just put
	// and update it
	key := struct {
		Filename string
	}{"1"}

	keyMap, err = dynamodbattribute.ConvertToMap(key)
	if err != nil {
		log.Println(err)
		return
	}

	// We will update the item by adding two more properties (columns)
	aCol := "Description"
	bCol := "Mtime"
	m1 := make(map[string]*string)
	// #a and #b are names which we will refer in the UpdateExpression below
	m1["#a"] = &aCol
	m1["#b"] = &bCol

	m2 := make(map[string]*dynamodb.AttributeValue)
	aVal := "A newly added Description"
	bVal := time.Now().String()
	m2[":aVal"] = &dynamodb.AttributeValue{S: &aVal}
	m2[":bVal"] = &dynamodb.AttributeValue{S: &bVal}

	log.Println("Updating the item by adding Description and Mtime")
	_, err = svc.UpdateItem(&dynamodb.UpdateItemInput{
		Key:                       keyMap,
		TableName:                 aws.String(tableName),
		UpdateExpression:          aws.String("SET #a = :aVal, #b = :bVal"),
		ExpressionAttributeNames:  m1,
		ExpressionAttributeValues: m2,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Succesfully modified the item in the dynamodb. Check AWS console.")
}
