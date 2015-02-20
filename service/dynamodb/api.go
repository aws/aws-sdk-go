package dynamodb

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"time"
)

// BatchGetItemRequest generates a request for the BatchGetItem operation.
func (c *DynamoDB) BatchGetItemRequest(input *BatchGetItemInput) (req *aws.Request, output *BatchGetItemOutput) {
	if opBatchGetItem == nil {
		opBatchGetItem = &aws.Operation{
			Name:       "BatchGetItem",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opBatchGetItem, input, output)
	output = &BatchGetItemOutput{}
	req.Data = output
	return
}

func (c *DynamoDB) BatchGetItem(input *BatchGetItemInput) (output *BatchGetItemOutput, err error) {
	req, out := c.BatchGetItemRequest(input)
	output = out
	err = req.Send()
	return
}

var opBatchGetItem *aws.Operation

// BatchWriteItemRequest generates a request for the BatchWriteItem operation.
func (c *DynamoDB) BatchWriteItemRequest(input *BatchWriteItemInput) (req *aws.Request, output *BatchWriteItemOutput) {
	if opBatchWriteItem == nil {
		opBatchWriteItem = &aws.Operation{
			Name:       "BatchWriteItem",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opBatchWriteItem, input, output)
	output = &BatchWriteItemOutput{}
	req.Data = output
	return
}

func (c *DynamoDB) BatchWriteItem(input *BatchWriteItemInput) (output *BatchWriteItemOutput, err error) {
	req, out := c.BatchWriteItemRequest(input)
	output = out
	err = req.Send()
	return
}

var opBatchWriteItem *aws.Operation

// CreateTableRequest generates a request for the CreateTable operation.
func (c *DynamoDB) CreateTableRequest(input *CreateTableInput) (req *aws.Request, output *CreateTableOutput) {
	if opCreateTable == nil {
		opCreateTable = &aws.Operation{
			Name:       "CreateTable",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCreateTable, input, output)
	output = &CreateTableOutput{}
	req.Data = output
	return
}

func (c *DynamoDB) CreateTable(input *CreateTableInput) (output *CreateTableOutput, err error) {
	req, out := c.CreateTableRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateTable *aws.Operation

// DeleteItemRequest generates a request for the DeleteItem operation.
func (c *DynamoDB) DeleteItemRequest(input *DeleteItemInput) (req *aws.Request, output *DeleteItemOutput) {
	if opDeleteItem == nil {
		opDeleteItem = &aws.Operation{
			Name:       "DeleteItem",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteItem, input, output)
	output = &DeleteItemOutput{}
	req.Data = output
	return
}

func (c *DynamoDB) DeleteItem(input *DeleteItemInput) (output *DeleteItemOutput, err error) {
	req, out := c.DeleteItemRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteItem *aws.Operation

// DeleteTableRequest generates a request for the DeleteTable operation.
func (c *DynamoDB) DeleteTableRequest(input *DeleteTableInput) (req *aws.Request, output *DeleteTableOutput) {
	if opDeleteTable == nil {
		opDeleteTable = &aws.Operation{
			Name:       "DeleteTable",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteTable, input, output)
	output = &DeleteTableOutput{}
	req.Data = output
	return
}

func (c *DynamoDB) DeleteTable(input *DeleteTableInput) (output *DeleteTableOutput, err error) {
	req, out := c.DeleteTableRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteTable *aws.Operation

// DescribeTableRequest generates a request for the DescribeTable operation.
func (c *DynamoDB) DescribeTableRequest(input *DescribeTableInput) (req *aws.Request, output *DescribeTableOutput) {
	if opDescribeTable == nil {
		opDescribeTable = &aws.Operation{
			Name:       "DescribeTable",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeTable, input, output)
	output = &DescribeTableOutput{}
	req.Data = output
	return
}

func (c *DynamoDB) DescribeTable(input *DescribeTableInput) (output *DescribeTableOutput, err error) {
	req, out := c.DescribeTableRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeTable *aws.Operation

// GetItemRequest generates a request for the GetItem operation.
func (c *DynamoDB) GetItemRequest(input *GetItemInput) (req *aws.Request, output *GetItemOutput) {
	if opGetItem == nil {
		opGetItem = &aws.Operation{
			Name:       "GetItem",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetItem, input, output)
	output = &GetItemOutput{}
	req.Data = output
	return
}

func (c *DynamoDB) GetItem(input *GetItemInput) (output *GetItemOutput, err error) {
	req, out := c.GetItemRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetItem *aws.Operation

// ListTablesRequest generates a request for the ListTables operation.
func (c *DynamoDB) ListTablesRequest(input *ListTablesInput) (req *aws.Request, output *ListTablesOutput) {
	if opListTables == nil {
		opListTables = &aws.Operation{
			Name:       "ListTables",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListTables, input, output)
	output = &ListTablesOutput{}
	req.Data = output
	return
}

func (c *DynamoDB) ListTables(input *ListTablesInput) (output *ListTablesOutput, err error) {
	req, out := c.ListTablesRequest(input)
	output = out
	err = req.Send()
	return
}

var opListTables *aws.Operation

// PutItemRequest generates a request for the PutItem operation.
func (c *DynamoDB) PutItemRequest(input *PutItemInput) (req *aws.Request, output *PutItemOutput) {
	if opPutItem == nil {
		opPutItem = &aws.Operation{
			Name:       "PutItem",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opPutItem, input, output)
	output = &PutItemOutput{}
	req.Data = output
	return
}

func (c *DynamoDB) PutItem(input *PutItemInput) (output *PutItemOutput, err error) {
	req, out := c.PutItemRequest(input)
	output = out
	err = req.Send()
	return
}

var opPutItem *aws.Operation

// QueryRequest generates a request for the Query operation.
func (c *DynamoDB) QueryRequest(input *QueryInput) (req *aws.Request, output *QueryOutput) {
	if opQuery == nil {
		opQuery = &aws.Operation{
			Name:       "Query",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opQuery, input, output)
	output = &QueryOutput{}
	req.Data = output
	return
}

func (c *DynamoDB) Query(input *QueryInput) (output *QueryOutput, err error) {
	req, out := c.QueryRequest(input)
	output = out
	err = req.Send()
	return
}

var opQuery *aws.Operation

// ScanRequest generates a request for the Scan operation.
func (c *DynamoDB) ScanRequest(input *ScanInput) (req *aws.Request, output *ScanOutput) {
	if opScan == nil {
		opScan = &aws.Operation{
			Name:       "Scan",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opScan, input, output)
	output = &ScanOutput{}
	req.Data = output
	return
}

func (c *DynamoDB) Scan(input *ScanInput) (output *ScanOutput, err error) {
	req, out := c.ScanRequest(input)
	output = out
	err = req.Send()
	return
}

var opScan *aws.Operation

// UpdateItemRequest generates a request for the UpdateItem operation.
func (c *DynamoDB) UpdateItemRequest(input *UpdateItemInput) (req *aws.Request, output *UpdateItemOutput) {
	if opUpdateItem == nil {
		opUpdateItem = &aws.Operation{
			Name:       "UpdateItem",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opUpdateItem, input, output)
	output = &UpdateItemOutput{}
	req.Data = output
	return
}

func (c *DynamoDB) UpdateItem(input *UpdateItemInput) (output *UpdateItemOutput, err error) {
	req, out := c.UpdateItemRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateItem *aws.Operation

// UpdateTableRequest generates a request for the UpdateTable operation.
func (c *DynamoDB) UpdateTableRequest(input *UpdateTableInput) (req *aws.Request, output *UpdateTableOutput) {
	if opUpdateTable == nil {
		opUpdateTable = &aws.Operation{
			Name:       "UpdateTable",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opUpdateTable, input, output)
	output = &UpdateTableOutput{}
	req.Data = output
	return
}

func (c *DynamoDB) UpdateTable(input *UpdateTableInput) (output *UpdateTableOutput, err error) {
	req, out := c.UpdateTableRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateTable *aws.Operation

type AttributeDefinition struct {
	AttributeName *string `type:"string" json:",omitempty"`
	AttributeType *string `type:"string" json:",omitempty"`

	metadataAttributeDefinition
}

type metadataAttributeDefinition struct {
	SDKShapeTraits bool `type:"structure" required:"AttributeName,AttributeType" json:",omitempty"`
}

type AttributeValue struct {
	B    *[]byte                     `type:"blob" json:",omitempty"`
	BOOL *bool                       `type:"boolean" json:",omitempty"`
	BS   *[]*[]byte                  `type:"list" json:",omitempty"`
	L    *[]*AttributeValue          `type:"list" json:",omitempty"`
	M    *map[string]*AttributeValue `type:"map" json:",omitempty"`
	N    *string                     `type:"string" json:",omitempty"`
	NS   *[]*string                  `type:"list" json:",omitempty"`
	NULL *bool                       `type:"boolean" json:",omitempty"`
	S    *string                     `type:"string" json:",omitempty"`
	SS   *[]*string                  `type:"list" json:",omitempty"`

	metadataAttributeValue
}

type metadataAttributeValue struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type AttributeValueUpdate struct {
	Action *string         `type:"string" json:",omitempty"`
	Value  *AttributeValue `type:"structure" json:",omitempty"`

	metadataAttributeValueUpdate
}

type metadataAttributeValueUpdate struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type BatchGetItemInput struct {
	RequestItems           *map[string]*KeysAndAttributes `type:"map" json:",omitempty"`
	ReturnConsumedCapacity *string                        `type:"string" json:",omitempty"`

	metadataBatchGetItemInput
}

type metadataBatchGetItemInput struct {
	SDKShapeTraits bool `type:"structure" required:"RequestItems" json:",omitempty"`
}

type BatchGetItemOutput struct {
	ConsumedCapacity *[]*ConsumedCapacity                       `type:"list" json:",omitempty"`
	Responses        *map[string]*[]*map[string]*AttributeValue `type:"map" json:",omitempty"`
	UnprocessedKeys  *map[string]*KeysAndAttributes             `type:"map" json:",omitempty"`

	metadataBatchGetItemOutput
}

type metadataBatchGetItemOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type BatchWriteItemInput struct {
	RequestItems                *map[string]*[]*WriteRequest `type:"map" json:",omitempty"`
	ReturnConsumedCapacity      *string                      `type:"string" json:",omitempty"`
	ReturnItemCollectionMetrics *string                      `type:"string" json:",omitempty"`

	metadataBatchWriteItemInput
}

type metadataBatchWriteItemInput struct {
	SDKShapeTraits bool `type:"structure" required:"RequestItems" json:",omitempty"`
}

type BatchWriteItemOutput struct {
	ConsumedCapacity      *[]*ConsumedCapacity                  `type:"list" json:",omitempty"`
	ItemCollectionMetrics *map[string]*[]*ItemCollectionMetrics `type:"map" json:",omitempty"`
	UnprocessedItems      *map[string]*[]*WriteRequest          `type:"map" json:",omitempty"`

	metadataBatchWriteItemOutput
}

type metadataBatchWriteItemOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type Capacity struct {
	CapacityUnits *float64 `type:"double" json:",omitempty"`

	metadataCapacity
}

type metadataCapacity struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type Condition struct {
	AttributeValueList *[]*AttributeValue `type:"list" json:",omitempty"`
	ComparisonOperator *string            `type:"string" json:",omitempty"`

	metadataCondition
}

type metadataCondition struct {
	SDKShapeTraits bool `type:"structure" required:"ComparisonOperator" json:",omitempty"`
}

type ConditionalCheckFailedException struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataConditionalCheckFailedException
}

type metadataConditionalCheckFailedException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ConsumedCapacity struct {
	CapacityUnits          *float64              `type:"double" json:",omitempty"`
	GlobalSecondaryIndexes *map[string]*Capacity `type:"map" json:",omitempty"`
	LocalSecondaryIndexes  *map[string]*Capacity `type:"map" json:",omitempty"`
	Table                  *Capacity             `type:"structure" json:",omitempty"`
	TableName              *string               `type:"string" json:",omitempty"`

	metadataConsumedCapacity
}

type metadataConsumedCapacity struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type CreateGlobalSecondaryIndexAction struct {
	IndexName             *string                `type:"string" json:",omitempty"`
	KeySchema             *[]*KeySchemaElement   `type:"list" json:",omitempty"`
	Projection            *Projection            `type:"structure" json:",omitempty"`
	ProvisionedThroughput *ProvisionedThroughput `type:"structure" json:",omitempty"`

	metadataCreateGlobalSecondaryIndexAction
}

type metadataCreateGlobalSecondaryIndexAction struct {
	SDKShapeTraits bool `type:"structure" required:"IndexName,KeySchema,Projection,ProvisionedThroughput" json:",omitempty"`
}

type CreateTableInput struct {
	AttributeDefinitions   *[]*AttributeDefinition  `type:"list" json:",omitempty"`
	GlobalSecondaryIndexes *[]*GlobalSecondaryIndex `type:"list" json:",omitempty"`
	KeySchema              *[]*KeySchemaElement     `type:"list" json:",omitempty"`
	LocalSecondaryIndexes  *[]*LocalSecondaryIndex  `type:"list" json:",omitempty"`
	ProvisionedThroughput  *ProvisionedThroughput   `type:"structure" json:",omitempty"`
	TableName              *string                  `type:"string" json:",omitempty"`

	metadataCreateTableInput
}

type metadataCreateTableInput struct {
	SDKShapeTraits bool `type:"structure" required:"AttributeDefinitions,TableName,KeySchema,ProvisionedThroughput" json:",omitempty"`
}

type CreateTableOutput struct {
	TableDescription *TableDescription `type:"structure" json:",omitempty"`

	metadataCreateTableOutput
}

type metadataCreateTableOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DeleteGlobalSecondaryIndexAction struct {
	IndexName *string `type:"string" json:",omitempty"`

	metadataDeleteGlobalSecondaryIndexAction
}

type metadataDeleteGlobalSecondaryIndexAction struct {
	SDKShapeTraits bool `type:"structure" required:"IndexName" json:",omitempty"`
}

type DeleteItemInput struct {
	ConditionExpression         *string                             `type:"string" json:",omitempty"`
	ConditionalOperator         *string                             `type:"string" json:",omitempty"`
	Expected                    *map[string]*ExpectedAttributeValue `type:"map" json:",omitempty"`
	ExpressionAttributeNames    *map[string]*string                 `type:"map" json:",omitempty"`
	ExpressionAttributeValues   *map[string]*AttributeValue         `type:"map" json:",omitempty"`
	Key                         *map[string]*AttributeValue         `type:"map" json:",omitempty"`
	ReturnConsumedCapacity      *string                             `type:"string" json:",omitempty"`
	ReturnItemCollectionMetrics *string                             `type:"string" json:",omitempty"`
	ReturnValues                *string                             `type:"string" json:",omitempty"`
	TableName                   *string                             `type:"string" json:",omitempty"`

	metadataDeleteItemInput
}

type metadataDeleteItemInput struct {
	SDKShapeTraits bool `type:"structure" required:"TableName,Key" json:",omitempty"`
}

type DeleteItemOutput struct {
	Attributes            *map[string]*AttributeValue `type:"map" json:",omitempty"`
	ConsumedCapacity      *ConsumedCapacity           `type:"structure" json:",omitempty"`
	ItemCollectionMetrics *ItemCollectionMetrics      `type:"structure" json:",omitempty"`

	metadataDeleteItemOutput
}

type metadataDeleteItemOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DeleteRequest struct {
	Key *map[string]*AttributeValue `type:"map" json:",omitempty"`

	metadataDeleteRequest
}

type metadataDeleteRequest struct {
	SDKShapeTraits bool `type:"structure" required:"Key" json:",omitempty"`
}

type DeleteTableInput struct {
	TableName *string `type:"string" json:",omitempty"`

	metadataDeleteTableInput
}

type metadataDeleteTableInput struct {
	SDKShapeTraits bool `type:"structure" required:"TableName" json:",omitempty"`
}

type DeleteTableOutput struct {
	TableDescription *TableDescription `type:"structure" json:",omitempty"`

	metadataDeleteTableOutput
}

type metadataDeleteTableOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeTableInput struct {
	TableName *string `type:"string" json:",omitempty"`

	metadataDescribeTableInput
}

type metadataDescribeTableInput struct {
	SDKShapeTraits bool `type:"structure" required:"TableName" json:",omitempty"`
}

type DescribeTableOutput struct {
	Table *TableDescription `type:"structure" json:",omitempty"`

	metadataDescribeTableOutput
}

type metadataDescribeTableOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ExpectedAttributeValue struct {
	AttributeValueList *[]*AttributeValue `type:"list" json:",omitempty"`
	ComparisonOperator *string            `type:"string" json:",omitempty"`
	Exists             *bool              `type:"boolean" json:",omitempty"`
	Value              *AttributeValue    `type:"structure" json:",omitempty"`

	metadataExpectedAttributeValue
}

type metadataExpectedAttributeValue struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetItemInput struct {
	AttributesToGet          *[]*string                  `type:"list" json:",omitempty"`
	ConsistentRead           *bool                       `type:"boolean" json:",omitempty"`
	ExpressionAttributeNames *map[string]*string         `type:"map" json:",omitempty"`
	Key                      *map[string]*AttributeValue `type:"map" json:",omitempty"`
	ProjectionExpression     *string                     `type:"string" json:",omitempty"`
	ReturnConsumedCapacity   *string                     `type:"string" json:",omitempty"`
	TableName                *string                     `type:"string" json:",omitempty"`

	metadataGetItemInput
}

type metadataGetItemInput struct {
	SDKShapeTraits bool `type:"structure" required:"TableName,Key" json:",omitempty"`
}

type GetItemOutput struct {
	ConsumedCapacity *ConsumedCapacity           `type:"structure" json:",omitempty"`
	Item             *map[string]*AttributeValue `type:"map" json:",omitempty"`

	metadataGetItemOutput
}

type metadataGetItemOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GlobalSecondaryIndex struct {
	IndexName             *string                `type:"string" json:",omitempty"`
	KeySchema             *[]*KeySchemaElement   `type:"list" json:",omitempty"`
	Projection            *Projection            `type:"structure" json:",omitempty"`
	ProvisionedThroughput *ProvisionedThroughput `type:"structure" json:",omitempty"`

	metadataGlobalSecondaryIndex
}

type metadataGlobalSecondaryIndex struct {
	SDKShapeTraits bool `type:"structure" required:"IndexName,KeySchema,Projection,ProvisionedThroughput" json:",omitempty"`
}

type GlobalSecondaryIndexDescription struct {
	Backfilling           *bool                             `type:"boolean" json:",omitempty"`
	IndexName             *string                           `type:"string" json:",omitempty"`
	IndexSizeBytes        *int64                            `type:"long" json:",omitempty"`
	IndexStatus           *string                           `type:"string" json:",omitempty"`
	ItemCount             *int64                            `type:"long" json:",omitempty"`
	KeySchema             *[]*KeySchemaElement              `type:"list" json:",omitempty"`
	Projection            *Projection                       `type:"structure" json:",omitempty"`
	ProvisionedThroughput *ProvisionedThroughputDescription `type:"structure" json:",omitempty"`

	metadataGlobalSecondaryIndexDescription
}

type metadataGlobalSecondaryIndexDescription struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GlobalSecondaryIndexUpdate struct {
	Create *CreateGlobalSecondaryIndexAction `type:"structure" json:",omitempty"`
	Delete *DeleteGlobalSecondaryIndexAction `type:"structure" json:",omitempty"`
	Update *UpdateGlobalSecondaryIndexAction `type:"structure" json:",omitempty"`

	metadataGlobalSecondaryIndexUpdate
}

type metadataGlobalSecondaryIndexUpdate struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InternalServerError struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataInternalServerError
}

type metadataInternalServerError struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ItemCollectionMetrics struct {
	ItemCollectionKey   *map[string]*AttributeValue `type:"map" json:",omitempty"`
	SizeEstimateRangeGB *[]*float64                 `type:"list" json:",omitempty"`

	metadataItemCollectionMetrics
}

type metadataItemCollectionMetrics struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ItemCollectionSizeLimitExceededException struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataItemCollectionSizeLimitExceededException
}

type metadataItemCollectionSizeLimitExceededException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type KeySchemaElement struct {
	AttributeName *string `type:"string" json:",omitempty"`
	KeyType       *string `type:"string" json:",omitempty"`

	metadataKeySchemaElement
}

type metadataKeySchemaElement struct {
	SDKShapeTraits bool `type:"structure" required:"AttributeName,KeyType" json:",omitempty"`
}

type KeysAndAttributes struct {
	AttributesToGet          *[]*string                     `type:"list" json:",omitempty"`
	ConsistentRead           *bool                          `type:"boolean" json:",omitempty"`
	ExpressionAttributeNames *map[string]*string            `type:"map" json:",omitempty"`
	Keys                     *[]*map[string]*AttributeValue `type:"list" json:",omitempty"`
	ProjectionExpression     *string                        `type:"string" json:",omitempty"`

	metadataKeysAndAttributes
}

type metadataKeysAndAttributes struct {
	SDKShapeTraits bool `type:"structure" required:"Keys" json:",omitempty"`
}

type LimitExceededException struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataLimitExceededException
}

type metadataLimitExceededException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListTablesInput struct {
	ExclusiveStartTableName *string `type:"string" json:",omitempty"`
	Limit                   *int    `type:"integer" json:",omitempty"`

	metadataListTablesInput
}

type metadataListTablesInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListTablesOutput struct {
	LastEvaluatedTableName *string    `type:"string" json:",omitempty"`
	TableNames             *[]*string `type:"list" json:",omitempty"`

	metadataListTablesOutput
}

type metadataListTablesOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type LocalSecondaryIndex struct {
	IndexName  *string              `type:"string" json:",omitempty"`
	KeySchema  *[]*KeySchemaElement `type:"list" json:",omitempty"`
	Projection *Projection          `type:"structure" json:",omitempty"`

	metadataLocalSecondaryIndex
}

type metadataLocalSecondaryIndex struct {
	SDKShapeTraits bool `type:"structure" required:"IndexName,KeySchema,Projection" json:",omitempty"`
}

type LocalSecondaryIndexDescription struct {
	IndexName      *string              `type:"string" json:",omitempty"`
	IndexSizeBytes *int64               `type:"long" json:",omitempty"`
	ItemCount      *int64               `type:"long" json:",omitempty"`
	KeySchema      *[]*KeySchemaElement `type:"list" json:",omitempty"`
	Projection     *Projection          `type:"structure" json:",omitempty"`

	metadataLocalSecondaryIndexDescription
}

type metadataLocalSecondaryIndexDescription struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type Projection struct {
	NonKeyAttributes *[]*string `type:"list" json:",omitempty"`
	ProjectionType   *string    `type:"string" json:",omitempty"`

	metadataProjection
}

type metadataProjection struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ProvisionedThroughput struct {
	ReadCapacityUnits  *int64 `type:"long" json:",omitempty"`
	WriteCapacityUnits *int64 `type:"long" json:",omitempty"`

	metadataProvisionedThroughput
}

type metadataProvisionedThroughput struct {
	SDKShapeTraits bool `type:"structure" required:"ReadCapacityUnits,WriteCapacityUnits" json:",omitempty"`
}

type ProvisionedThroughputDescription struct {
	LastDecreaseDateTime   *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	LastIncreaseDateTime   *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	NumberOfDecreasesToday *int64     `type:"long" json:",omitempty"`
	ReadCapacityUnits      *int64     `type:"long" json:",omitempty"`
	WriteCapacityUnits     *int64     `type:"long" json:",omitempty"`

	metadataProvisionedThroughputDescription
}

type metadataProvisionedThroughputDescription struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ProvisionedThroughputExceededException struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataProvisionedThroughputExceededException
}

type metadataProvisionedThroughputExceededException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type PutItemInput struct {
	ConditionExpression         *string                             `type:"string" json:",omitempty"`
	ConditionalOperator         *string                             `type:"string" json:",omitempty"`
	Expected                    *map[string]*ExpectedAttributeValue `type:"map" json:",omitempty"`
	ExpressionAttributeNames    *map[string]*string                 `type:"map" json:",omitempty"`
	ExpressionAttributeValues   *map[string]*AttributeValue         `type:"map" json:",omitempty"`
	Item                        *map[string]*AttributeValue         `type:"map" json:",omitempty"`
	ReturnConsumedCapacity      *string                             `type:"string" json:",omitempty"`
	ReturnItemCollectionMetrics *string                             `type:"string" json:",omitempty"`
	ReturnValues                *string                             `type:"string" json:",omitempty"`
	TableName                   *string                             `type:"string" json:",omitempty"`

	metadataPutItemInput
}

type metadataPutItemInput struct {
	SDKShapeTraits bool `type:"structure" required:"TableName,Item" json:",omitempty"`
}

type PutItemOutput struct {
	Attributes            *map[string]*AttributeValue `type:"map" json:",omitempty"`
	ConsumedCapacity      *ConsumedCapacity           `type:"structure" json:",omitempty"`
	ItemCollectionMetrics *ItemCollectionMetrics      `type:"structure" json:",omitempty"`

	metadataPutItemOutput
}

type metadataPutItemOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type PutRequest struct {
	Item *map[string]*AttributeValue `type:"map" json:",omitempty"`

	metadataPutRequest
}

type metadataPutRequest struct {
	SDKShapeTraits bool `type:"structure" required:"Item" json:",omitempty"`
}

type QueryInput struct {
	AttributesToGet           *[]*string                  `type:"list" json:",omitempty"`
	ConditionalOperator       *string                     `type:"string" json:",omitempty"`
	ConsistentRead            *bool                       `type:"boolean" json:",omitempty"`
	ExclusiveStartKey         *map[string]*AttributeValue `type:"map" json:",omitempty"`
	ExpressionAttributeNames  *map[string]*string         `type:"map" json:",omitempty"`
	ExpressionAttributeValues *map[string]*AttributeValue `type:"map" json:",omitempty"`
	FilterExpression          *string                     `type:"string" json:",omitempty"`
	IndexName                 *string                     `type:"string" json:",omitempty"`
	KeyConditions             *map[string]*Condition      `type:"map" json:",omitempty"`
	Limit                     *int                        `type:"integer" json:",omitempty"`
	ProjectionExpression      *string                     `type:"string" json:",omitempty"`
	QueryFilter               *map[string]*Condition      `type:"map" json:",omitempty"`
	ReturnConsumedCapacity    *string                     `type:"string" json:",omitempty"`
	ScanIndexForward          *bool                       `type:"boolean" json:",omitempty"`
	Select                    *string                     `type:"string" json:",omitempty"`
	TableName                 *string                     `type:"string" json:",omitempty"`

	metadataQueryInput
}

type metadataQueryInput struct {
	SDKShapeTraits bool `type:"structure" required:"TableName,KeyConditions" json:",omitempty"`
}

type QueryOutput struct {
	ConsumedCapacity *ConsumedCapacity              `type:"structure" json:",omitempty"`
	Count            *int                           `type:"integer" json:",omitempty"`
	Items            *[]*map[string]*AttributeValue `type:"list" json:",omitempty"`
	LastEvaluatedKey *map[string]*AttributeValue    `type:"map" json:",omitempty"`
	ScannedCount     *int                           `type:"integer" json:",omitempty"`

	metadataQueryOutput
}

type metadataQueryOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ResourceInUseException struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataResourceInUseException
}

type metadataResourceInUseException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ResourceNotFoundException struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataResourceNotFoundException
}

type metadataResourceNotFoundException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ScanInput struct {
	AttributesToGet           *[]*string                  `type:"list" json:",omitempty"`
	ConditionalOperator       *string                     `type:"string" json:",omitempty"`
	ExclusiveStartKey         *map[string]*AttributeValue `type:"map" json:",omitempty"`
	ExpressionAttributeNames  *map[string]*string         `type:"map" json:",omitempty"`
	ExpressionAttributeValues *map[string]*AttributeValue `type:"map" json:",omitempty"`
	FilterExpression          *string                     `type:"string" json:",omitempty"`
	IndexName                 *string                     `type:"string" json:",omitempty"`
	Limit                     *int                        `type:"integer" json:",omitempty"`
	ProjectionExpression      *string                     `type:"string" json:",omitempty"`
	ReturnConsumedCapacity    *string                     `type:"string" json:",omitempty"`
	ScanFilter                *map[string]*Condition      `type:"map" json:",omitempty"`
	Segment                   *int                        `type:"integer" json:",omitempty"`
	Select                    *string                     `type:"string" json:",omitempty"`
	TableName                 *string                     `type:"string" json:",omitempty"`
	TotalSegments             *int                        `type:"integer" json:",omitempty"`

	metadataScanInput
}

type metadataScanInput struct {
	SDKShapeTraits bool `type:"structure" required:"TableName" json:",omitempty"`
}

type ScanOutput struct {
	ConsumedCapacity *ConsumedCapacity              `type:"structure" json:",omitempty"`
	Count            *int                           `type:"integer" json:",omitempty"`
	Items            *[]*map[string]*AttributeValue `type:"list" json:",omitempty"`
	LastEvaluatedKey *map[string]*AttributeValue    `type:"map" json:",omitempty"`
	ScannedCount     *int                           `type:"integer" json:",omitempty"`

	metadataScanOutput
}

type metadataScanOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type TableDescription struct {
	AttributeDefinitions   *[]*AttributeDefinition             `type:"list" json:",omitempty"`
	CreationDateTime       *time.Time                          `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	GlobalSecondaryIndexes *[]*GlobalSecondaryIndexDescription `type:"list" json:",omitempty"`
	ItemCount              *int64                              `type:"long" json:",omitempty"`
	KeySchema              *[]*KeySchemaElement                `type:"list" json:",omitempty"`
	LocalSecondaryIndexes  *[]*LocalSecondaryIndexDescription  `type:"list" json:",omitempty"`
	ProvisionedThroughput  *ProvisionedThroughputDescription   `type:"structure" json:",omitempty"`
	TableName              *string                             `type:"string" json:",omitempty"`
	TableSizeBytes         *int64                              `type:"long" json:",omitempty"`
	TableStatus            *string                             `type:"string" json:",omitempty"`

	metadataTableDescription
}

type metadataTableDescription struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type UpdateGlobalSecondaryIndexAction struct {
	IndexName             *string                `type:"string" json:",omitempty"`
	ProvisionedThroughput *ProvisionedThroughput `type:"structure" json:",omitempty"`

	metadataUpdateGlobalSecondaryIndexAction
}

type metadataUpdateGlobalSecondaryIndexAction struct {
	SDKShapeTraits bool `type:"structure" required:"IndexName,ProvisionedThroughput" json:",omitempty"`
}

type UpdateItemInput struct {
	AttributeUpdates            *map[string]*AttributeValueUpdate   `type:"map" json:",omitempty"`
	ConditionExpression         *string                             `type:"string" json:",omitempty"`
	ConditionalOperator         *string                             `type:"string" json:",omitempty"`
	Expected                    *map[string]*ExpectedAttributeValue `type:"map" json:",omitempty"`
	ExpressionAttributeNames    *map[string]*string                 `type:"map" json:",omitempty"`
	ExpressionAttributeValues   *map[string]*AttributeValue         `type:"map" json:",omitempty"`
	Key                         *map[string]*AttributeValue         `type:"map" json:",omitempty"`
	ReturnConsumedCapacity      *string                             `type:"string" json:",omitempty"`
	ReturnItemCollectionMetrics *string                             `type:"string" json:",omitempty"`
	ReturnValues                *string                             `type:"string" json:",omitempty"`
	TableName                   *string                             `type:"string" json:",omitempty"`
	UpdateExpression            *string                             `type:"string" json:",omitempty"`

	metadataUpdateItemInput
}

type metadataUpdateItemInput struct {
	SDKShapeTraits bool `type:"structure" required:"TableName,Key" json:",omitempty"`
}

type UpdateItemOutput struct {
	Attributes            *map[string]*AttributeValue `type:"map" json:",omitempty"`
	ConsumedCapacity      *ConsumedCapacity           `type:"structure" json:",omitempty"`
	ItemCollectionMetrics *ItemCollectionMetrics      `type:"structure" json:",omitempty"`

	metadataUpdateItemOutput
}

type metadataUpdateItemOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type UpdateTableInput struct {
	AttributeDefinitions        *[]*AttributeDefinition        `type:"list" json:",omitempty"`
	GlobalSecondaryIndexUpdates *[]*GlobalSecondaryIndexUpdate `type:"list" json:",omitempty"`
	ProvisionedThroughput       *ProvisionedThroughput         `type:"structure" json:",omitempty"`
	TableName                   *string                        `type:"string" json:",omitempty"`

	metadataUpdateTableInput
}

type metadataUpdateTableInput struct {
	SDKShapeTraits bool `type:"structure" required:"TableName" json:",omitempty"`
}

type UpdateTableOutput struct {
	TableDescription *TableDescription `type:"structure" json:",omitempty"`

	metadataUpdateTableOutput
}

type metadataUpdateTableOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type WriteRequest struct {
	DeleteRequest *DeleteRequest `type:"structure" json:",omitempty"`
	PutRequest    *PutRequest    `type:"structure" json:",omitempty"`

	metadataWriteRequest
}

type metadataWriteRequest struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}