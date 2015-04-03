package dynamodb_test

import (
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
)

func TestCustomRetryRules(t *testing.T) {
	db := dynamodb.New(&aws.Config{MaxRetries: -1})
	assert.Equal(t, db.MaxRetries(), uint(10))
}
