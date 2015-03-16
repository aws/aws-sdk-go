// +build integration

package dynamodb_test

import (
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/util/utilassert"
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
)

var (
	_ = assert.Equal
	_ = utilassert.Match
)

func TestMakingABasicRequest(t *testing.T) {
	client := dynamodb.New(nil)
	resp, e := client.ListTables(&dynamodb.ListTablesInput{
		Limit: aws.Long(1),
	})
	err := aws.Error(e)
	_, _, _ = resp, e, err // avoid unused warnings

	assert.NoError(t, nil, err)

}

func TestErrorHandling(t *testing.T) {
	client := dynamodb.New(nil)
	resp, e := client.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String("fake-table"),
	})
	err := aws.Error(e)
	_, _, _ = resp, e, err // avoid unused warnings

	assert.NotEqual(t, nil, err)
	assert.Equal(t, "ResourceNotFoundException", err.Code)
	utilassert.Match(t, "Requested resource not found: Table: fake-table not found", err.Message)

}
