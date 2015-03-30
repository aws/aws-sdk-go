package aws_test

import (
	"fmt"
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
)

// Use DynamoDB methods for simplicity
func TestPagination(t *testing.T) {
	db := dynamodb.New(nil)
	tokens, pages, numPages, gotToEnd := []string{}, []string{}, 0, false

	reqNum := 0
	resps := []*dynamodb.ListTablesOutput{
		&dynamodb.ListTablesOutput{TableNames: []*string{aws.String("Table1"), aws.String("Table2")}, LastEvaluatedTableName: aws.String("Table2")},
		&dynamodb.ListTablesOutput{TableNames: []*string{aws.String("Table3"), aws.String("Table4")}, LastEvaluatedTableName: aws.String("Table4")},
		&dynamodb.ListTablesOutput{TableNames: []*string{aws.String("Table5")}},
	}

	db.Handlers.Send.Init() // mock sending
	db.Handlers.Unmarshal.Init()
	db.Handlers.UnmarshalMeta.Init()
	db.Handlers.ValidateResponse.Init()
	db.Handlers.Build.PushBack(func(r *aws.Request) {
		in := r.Params.(*dynamodb.ListTablesInput)
		if in == nil {
			tokens = append(tokens, "")
		} else if in.ExclusiveStartTableName != nil {
			tokens = append(tokens, *in.ExclusiveStartTableName)
		}
	})
	db.Handlers.Unmarshal.PushBack(func(r *aws.Request) {
		r.Data = resps[reqNum]
		reqNum++
	})

	params := &dynamodb.ListTablesInput{Limit: aws.Long(2)}
	db.ListTablesPages(params, func(p *dynamodb.ListTablesOutput, err error) bool {
		if p != nil {
			numPages++
			for _, t := range p.TableNames {
				pages = append(pages, *t)
			}
		} else if p == nil && err == nil {
			gotToEnd = true
		}

		return true
	})

	assert.Equal(t, []string{"Table2", "Table4"}, tokens)
	assert.Equal(t, []string{"Table1", "Table2", "Table3", "Table4", "Table5"}, pages)
	assert.Equal(t, 3, numPages)
	assert.True(t, gotToEnd)
}

// Use DynamoDB methods for simplicity
func TestPaginationEarlyExit(t *testing.T) {
	db := dynamodb.New(nil)
	numPages, gotToEnd := 0, false

	reqNum := 0
	resps := []*dynamodb.ListTablesOutput{
		&dynamodb.ListTablesOutput{TableNames: []*string{aws.String("Table1"), aws.String("Table2")}, LastEvaluatedTableName: aws.String("Table2")},
		&dynamodb.ListTablesOutput{TableNames: []*string{aws.String("Table3"), aws.String("Table4")}, LastEvaluatedTableName: aws.String("Table4")},
		&dynamodb.ListTablesOutput{TableNames: []*string{aws.String("Table5")}},
	}

	db.Handlers.Send.Init() // mock sending
	db.Handlers.Unmarshal.Init()
	db.Handlers.UnmarshalMeta.Init()
	db.Handlers.ValidateResponse.Init()
	db.Handlers.Unmarshal.PushBack(func(r *aws.Request) {
		r.Data = resps[reqNum]
		reqNum++
	})

	params := &dynamodb.ListTablesInput{Limit: aws.Long(2)}
	db.ListTablesPages(params, func(p *dynamodb.ListTablesOutput, err error) bool {
		fmt.Println("page", numPages, p)
		if p == nil && err == nil {
			gotToEnd = true
		} else {
			numPages++
		}
		if numPages == 2 {
			fmt.Println("BREAKING")
			return false
		}
		return true
	})

	assert.Equal(t, 2, numPages)
	assert.False(t, gotToEnd)
}
