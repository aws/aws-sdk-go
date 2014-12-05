// Package dynamodb provides a client for Amazon DynamoDB.
package dynamodb

import (
	"fmt"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
)

// DynamoDB is a client for Amazon DynamoDB.
type DynamoDB struct {
	client aws.Client
}

// New returns a new DynamoDB client.
func New(key, secret, region string, client *http.Client) *DynamoDB {
	if client == nil {
		client = http.DefaultClient
	}

	return &DynamoDB{
		client: &aws.JSONClient{
			Client:       client,
			Region:       region,
			Endpoint:     fmt.Sprintf("https://dynamodb.%s.amazonaws.com", region),
			Prefix:       "dynamodb",
			Key:          key,
			Secret:       secret,
			JSONVersion:  "1.0",
			TargetPrefix: "DynamoDB_20120810",
		},
	}
}

// BatchGetItem the BatchGetItem operation returns the attributes of one or
// more items from one or more tables. You identify requested items by
// primary key. A single operation can retrieve up to 16 MB of data, which
// can contain as many as 100 items. BatchGetItem will return a partial
// result if the response size limit is exceeded, the table's provisioned
// throughput is exceeded, or an internal processing failure occurs. If a
// partial result is returned, the operation returns a value for
// UnprocessedKeys . You can use this value to retry the operation starting
// with the next item to get. For example, if you ask to retrieve 100
// items, but each individual item is 300 KB in size, the system returns 52
// items (so as not to exceed the 16 MB limit). It also returns an
// appropriate UnprocessedKeys value so you can get the next page of
// results. If desired, your application can include its own logic to
// assemble the pages of results into one data set. If none of the items
// can be processed due to insufficient provisioned throughput on all of
// the tables in the request, then BatchGetItem will return a
// ProvisionedThroughputExceededException . If at least one of the items is
// successfully processed, then BatchGetItem completes successfully, while
// returning the keys of the unread items in UnprocessedKeys If DynamoDB
// returns any unprocessed items, you should retry the batch operation on
// those items. However, we strongly recommend that you use an exponential
// backoff algorithm . If you retry the batch operation immediately, the
// underlying read or write requests can still fail due to throttling on
// the individual tables. If you delay the batch operation using
// exponential backoff, the individual requests in the batch are much more
// likely to succeed. For more information, go to Batch Operations and
// Error Handling in the Amazon DynamoDB Developer Guide By default,
// BatchGetItem performs eventually consistent reads on every table in the
// request. If you want strongly consistent reads instead, you can set
// ConsistentRead to true for any or all tables. In order to minimize
// response latency, BatchGetItem retrieves items in parallel. When
// designing your application, keep in mind that DynamoDB does not return
// attributes in any particular order. To help parse the response by item,
// include the primary key values for the items in your request in the
// AttributesToGet parameter. If a requested item does not exist, it is not
// returned in the result. Requests for nonexistent items consume the
// minimum read capacity units according to the type of read. For more
// information, see Capacity Units Calculations in the Amazon DynamoDB
// Developer Guide
func (c *DynamoDB) BatchGetItem(req BatchGetItemInput) (resp *BatchGetItemOutput, err error) {
	resp = &BatchGetItemOutput{}
	err = c.client.Do("BatchGetItem", "POST", "/", req, resp)
	return
}

// BatchWriteItem the BatchWriteItem operation puts or deletes multiple
// items in one or more tables. A single call to BatchWriteItem can write
// up to 16 MB of data, which can comprise as many as 25 put or delete
// requests. Individual items to be written can be as large as 400 The
// individual PutItem and DeleteItem operations specified in BatchWriteItem
// are atomic; however BatchWriteItem as a whole is not. If any requested
// operations fail because the table's provisioned throughput is exceeded
// or an internal processing failure occurs, the failed operations are
// returned in the UnprocessedItems response parameter. You can investigate
// and optionally resend the requests. Typically, you would call
// BatchWriteItem in a loop. Each iteration would check for unprocessed
// items and submit a new BatchWriteItem request with those unprocessed
// items until all items have been processed. Note that if none of the
// items can be processed due to insufficient provisioned throughput on all
// of the tables in the request, then BatchWriteItem will return a
// ProvisionedThroughputExceededException If DynamoDB returns any
// unprocessed items, you should retry the batch operation on those items.
// However, we strongly recommend that you use an exponential backoff
// algorithm . If you retry the batch operation immediately, the underlying
// read or write requests can still fail due to throttling on the
// individual tables. If you delay the batch operation using exponential
// backoff, the individual requests in the batch are much more likely to
// succeed. For more information, go to Batch Operations and Error Handling
// in the Amazon DynamoDB Developer Guide With BatchWriteItem , you can
// efficiently write or delete large amounts of data, such as from Amazon
// Elastic MapReduce or copy data from another database into DynamoDB. In
// order to improve performance with these large-scale operations,
// BatchWriteItem does not behave in the same way as individual PutItem and
// DeleteItem calls would For example, you cannot specify conditions on
// individual put and delete requests, and BatchWriteItem does not return
// deleted items in the response. If you use a programming language that
// supports concurrency, such as Java, you can use threads to write items
// in parallel. Your application must include the necessary logic to manage
// the threads. With languages that don't support threading, such as you
// must update or delete the specified items one at a time. In both
// situations, BatchWriteItem provides an alternative where the API
// performs the specified put and delete operations in parallel, giving you
// the power of the thread pool approach without having to introduce
// complexity into your application. Parallel processing reduces latency,
// but each specified put and delete request consumes the same number of
// write capacity units whether it is processed in parallel or not. Delete
// operations on nonexistent items consume one write capacity unit. If one
// or more of the following is true, DynamoDB rejects the entire batch
// write operation: One or more tables specified in the BatchWriteItem
// request does not exist. Primary key attributes specified on an item in
// the request do not match those in the corresponding table's primary key
// schema. You try to perform multiple operations on the same item in the
// same BatchWriteItem request. For example, you cannot put and delete the
// same item in the same BatchWriteItem request. There are more than 25
// requests in the batch.
func (c *DynamoDB) BatchWriteItem(req BatchWriteItemInput) (resp *BatchWriteItemOutput, err error) {
	resp = &BatchWriteItemOutput{}
	err = c.client.Do("BatchWriteItem", "POST", "/", req, resp)
	return
}

// CreateTable the CreateTable operation adds a new table to your account.
// In an AWS account, table names must be unique within each region. That
// is, you can have two tables with same name if you create the tables in
// different regions. CreateTable is an asynchronous operation. Upon
// receiving a CreateTable request, DynamoDB immediately returns a response
// with a TableStatus of . After the table is created, DynamoDB sets the
// TableStatus to . You can perform read and write operations only on an
// table. If you want to create multiple tables with secondary indexes on
// them, you must create them sequentially. Only one table with secondary
// indexes can be in the state at any given time. You can use the
// DescribeTable API to check the table status.
func (c *DynamoDB) CreateTable(req CreateTableInput) (resp *CreateTableOutput, err error) {
	resp = &CreateTableOutput{}
	err = c.client.Do("CreateTable", "POST", "/", req, resp)
	return
}

// DeleteItem deletes a single item in a table by primary key. You can
// perform a conditional delete operation that deletes the item if it
// exists, or if it has an expected attribute value. In addition to
// deleting an item, you can also return the item's attribute values in the
// same operation, using the ReturnValues parameter. Unless you specify
// conditions, the DeleteItem is an idempotent operation; running it
// multiple times on the same item or attribute does not result in an error
// response. Conditional deletes are useful for deleting items only if
// specific conditions are met. If those conditions are met, DynamoDB
// performs the delete. Otherwise, the item is not deleted.
func (c *DynamoDB) DeleteItem(req DeleteItemInput) (resp *DeleteItemOutput, err error) {
	resp = &DeleteItemOutput{}
	err = c.client.Do("DeleteItem", "POST", "/", req, resp)
	return
}

// DeleteTable the DeleteTable operation deletes a table and all of its
// items. After a DeleteTable request, the specified table is in the state
// until DynamoDB completes the deletion. If the table is in the state, you
// can delete it. If a table is in or states, then DynamoDB returns a
// ResourceInUseException . If the specified table does not exist, DynamoDB
// returns a ResourceNotFoundException . If table is already in the state,
// no error is returned. When you delete a table, any indexes on that table
// are also deleted. Use the DescribeTable API to check the status of the
// table.
func (c *DynamoDB) DeleteTable(req DeleteTableInput) (resp *DeleteTableOutput, err error) {
	resp = &DeleteTableOutput{}
	err = c.client.Do("DeleteTable", "POST", "/", req, resp)
	return
}

// DescribeTable returns information about the table, including the current
// status of the table, when it was created, the primary key schema, and
// any indexes on the table.
func (c *DynamoDB) DescribeTable(req DescribeTableInput) (resp *DescribeTableOutput, err error) {
	resp = &DescribeTableOutput{}
	err = c.client.Do("DescribeTable", "POST", "/", req, resp)
	return
}

// GetItem the GetItem operation returns a set of attributes for the item
// with the given primary key. If there is no matching item, GetItem does
// not return any data. GetItem provides an eventually consistent read by
// default. If your application requires a strongly consistent read, set
// ConsistentRead to true . Although a strongly consistent read might take
// more time than an eventually consistent read, it always returns the last
// updated value.
func (c *DynamoDB) GetItem(req GetItemInput) (resp *GetItemOutput, err error) {
	resp = &GetItemOutput{}
	err = c.client.Do("GetItem", "POST", "/", req, resp)
	return
}

// ListTables returns an array of table names associated with the current
// account and endpoint. The output from ListTables is paginated, with each
// page returning a maximum of 100 table names.
func (c *DynamoDB) ListTables(req ListTablesInput) (resp *ListTablesOutput, err error) {
	resp = &ListTablesOutput{}
	err = c.client.Do("ListTables", "POST", "/", req, resp)
	return
}

// PutItem creates a new item, or replaces an old item with a new item. If
// an item that has the same primary key as the new item already exists in
// the specified table, the new item completely replaces the existing item.
// You can perform a conditional put operation (add a new item if one with
// the specified primary key doesn't exist), or replace an existing item if
// it has certain attribute values. In addition to putting an item, you can
// also return the item's attribute values in the same operation, using the
// ReturnValues parameter. When you add an item, the primary key
// attribute(s) are the only required attributes. Attribute values cannot
// be null. String and Binary type attributes must have lengths greater
// than zero. Set type attributes cannot be empty. Requests with empty
// values will be rejected with a ValidationException exception. You can
// request that PutItem return either a copy of the original item (before
// the update) or a copy of the updated item (after the update). For more
// information, see the ReturnValues description below. For more
// information about using this see Working with Items in the Amazon
// DynamoDB Developer Guide
func (c *DynamoDB) PutItem(req PutItemInput) (resp *PutItemOutput, err error) {
	resp = &PutItemOutput{}
	err = c.client.Do("PutItem", "POST", "/", req, resp)
	return
}

// Query a Query operation directly accesses items from a table using the
// table primary key, or from an index using the index key. You must
// provide a specific hash key value. You can narrow the scope of the query
// by using comparison operators on the range key value, or on the index
// key. You can use the ScanIndexForward parameter to get results in
// forward or reverse order, by range key or by index key. Queries that do
// not return results consume the minimum number of read capacity units for
// that type of read operation. If the total number of items meeting the
// query criteria exceeds the result set size limit of 1 MB, the query
// stops and results are returned to the user with LastEvaluatedKey to
// continue the query in a subsequent operation. Unlike a Scan operation, a
// Query operation never returns both an empty result set and a
// LastEvaluatedKey . The LastEvaluatedKey is only provided if the results
// exceed 1 MB, or if you have used Limit . You can query a table, a local
// secondary index, or a global secondary index. For a query on a table or
// on a local secondary index, you can set ConsistentRead to true and
// obtain a strongly consistent result. Global secondary indexes support
// eventually consistent reads only, so do not specify ConsistentRead when
// querying a global secondary index.
func (c *DynamoDB) Query(req QueryInput) (resp *QueryOutput, err error) {
	resp = &QueryOutput{}
	err = c.client.Do("Query", "POST", "/", req, resp)
	return
}

// Scan the Scan operation returns one or more items and item attributes by
// accessing every item in the table. To have DynamoDB return fewer items,
// you can provide a ScanFilter operation. If the total number of scanned
// items exceeds the maximum data set size limit of 1 MB, the scan stops
// and results are returned to the user as a LastEvaluatedKey value to
// continue the scan in a subsequent operation. The results also include
// the number of items exceeding the limit. A scan can result in no table
// data meeting the filter criteria. The result set is eventually
// consistent. By default, Scan operations proceed sequentially; however,
// for faster performance on large tables, applications can request a
// parallel Scan operation by specifying the Segment and TotalSegments
// parameters. For more information, see Parallel Scan in the Amazon
// DynamoDB Developer Guide
func (c *DynamoDB) Scan(req ScanInput) (resp *ScanOutput, err error) {
	resp = &ScanOutput{}
	err = c.client.Do("Scan", "POST", "/", req, resp)
	return
}

// UpdateItem edits an existing item's attributes, or adds a new item to
// the table if it does not already exist. You can put, delete, or add
// attribute values. You can also perform a conditional update (insert a
// new attribute name-value pair if it doesn't exist, or replace an
// existing name-value pair if it has certain expected attribute values).
// You can also return the item's attribute values in the same UpdateItem
// operation using the ReturnValues parameter.
func (c *DynamoDB) UpdateItem(req UpdateItemInput) (resp *UpdateItemOutput, err error) {
	resp = &UpdateItemOutput{}
	err = c.client.Do("UpdateItem", "POST", "/", req, resp)
	return
}

// UpdateTable updates the provisioned throughput for the given table.
// Setting the throughput for a table helps you manage performance and is
// part of the provisioned throughput feature of DynamoDB. The provisioned
// throughput values can be upgraded or downgraded based on the maximums
// and minimums listed in the Limits section in the Amazon DynamoDB
// Developer Guide The table must be in the state for this operation to
// succeed. UpdateTable is an asynchronous operation; while executing the
// operation, the table is in the state. While the table is in the state,
// the table still has the provisioned throughput from before the call. The
// new provisioned throughput setting is in effect only when the table
// returns to the state after the UpdateTable operation. You cannot add,
// modify or delete indexes using UpdateTable . Indexes can only be defined
// at table creation time.
func (c *DynamoDB) UpdateTable(req UpdateTableInput) (resp *UpdateTableOutput, err error) {
	resp = &UpdateTableOutput{}
	err = c.client.Do("UpdateTable", "POST", "/", req, resp)
	return
}

type AttributeDefinition struct {
	AttributeName string `json:"AttributeName"`
	AttributeType string `json:"AttributeType"`
}

type AttributeValue struct {
	B    []byte                    `json:"B,omitempty"`
	BOOL bool                      `json:"BOOL,omitempty"`
	BS   [][]byte                  `json:"BS,omitempty"`
	L    []AttributeValue          `json:"L,omitempty"`
	M    map[string]AttributeValue `json:"M,omitempty"`
	N    string                    `json:"N,omitempty"`
	NS   []string                  `json:"NS,omitempty"`
	NULL bool                      `json:"NULL,omitempty"`
	S    string                    `json:"S,omitempty"`
	SS   []string                  `json:"SS,omitempty"`
}

type AttributeValueUpdate struct {
	Action string         `json:"Action,omitempty"`
	Value  AttributeValue `json:"Value,omitempty"`
}

type BatchGetItemInput struct {
	RequestItems           map[string]KeysAndAttributes `json:"RequestItems"`
	ReturnConsumedCapacity string                       `json:"ReturnConsumedCapacity,omitempty"`
}

type BatchGetItemOutput struct {
	ConsumedCapacity []ConsumedCapacity                     `json:"ConsumedCapacity,omitempty"`
	Responses        map[string][]map[string]AttributeValue `json:"Responses,omitempty"`
	UnprocessedKeys  map[string]KeysAndAttributes           `json:"UnprocessedKeys,omitempty"`
}

type BatchWriteItemInput struct {
	RequestItems                map[string][]WriteRequest `json:"RequestItems"`
	ReturnConsumedCapacity      string                    `json:"ReturnConsumedCapacity,omitempty"`
	ReturnItemCollectionMetrics string                    `json:"ReturnItemCollectionMetrics,omitempty"`
}

type BatchWriteItemOutput struct {
	ConsumedCapacity      []ConsumedCapacity                 `json:"ConsumedCapacity,omitempty"`
	ItemCollectionMetrics map[string][]ItemCollectionMetrics `json:"ItemCollectionMetrics,omitempty"`
	UnprocessedItems      map[string][]WriteRequest          `json:"UnprocessedItems,omitempty"`
}

type Capacity struct {
	CapacityUnits float64 `json:"CapacityUnits,omitempty"`
}

type Condition struct {
	AttributeValueList []AttributeValue `json:"AttributeValueList,omitempty"`
	ComparisonOperator string           `json:"ComparisonOperator"`
}

type ConsumedCapacity struct {
	CapacityUnits          float64             `json:"CapacityUnits,omitempty"`
	GlobalSecondaryIndexes map[string]Capacity `json:"GlobalSecondaryIndexes,omitempty"`
	LocalSecondaryIndexes  map[string]Capacity `json:"LocalSecondaryIndexes,omitempty"`
	Table                  Capacity            `json:"Table,omitempty"`
	TableName              string              `json:"TableName,omitempty"`
}

type CreateTableInput struct {
	AttributeDefinitions   []AttributeDefinition  `json:"AttributeDefinitions"`
	GlobalSecondaryIndexes []GlobalSecondaryIndex `json:"GlobalSecondaryIndexes,omitempty"`
	KeySchema              []KeySchemaElement     `json:"KeySchema"`
	LocalSecondaryIndexes  []LocalSecondaryIndex  `json:"LocalSecondaryIndexes,omitempty"`
	ProvisionedThroughput  ProvisionedThroughput  `json:"ProvisionedThroughput"`
	TableName              string                 `json:"TableName"`
}

type CreateTableOutput struct {
	TableDescription TableDescription `json:"TableDescription,omitempty"`
}

type DeleteItemInput struct {
	ConditionExpression         string                            `json:"ConditionExpression,omitempty"`
	ConditionalOperator         string                            `json:"ConditionalOperator,omitempty"`
	Expected                    map[string]ExpectedAttributeValue `json:"Expected,omitempty"`
	ExpressionAttributeNames    map[string]string                 `json:"ExpressionAttributeNames,omitempty"`
	ExpressionAttributeValues   map[string]AttributeValue         `json:"ExpressionAttributeValues,omitempty"`
	Key                         map[string]AttributeValue         `json:"Key"`
	ReturnConsumedCapacity      string                            `json:"ReturnConsumedCapacity,omitempty"`
	ReturnItemCollectionMetrics string                            `json:"ReturnItemCollectionMetrics,omitempty"`
	ReturnValues                string                            `json:"ReturnValues,omitempty"`
	TableName                   string                            `json:"TableName"`
}

type DeleteItemOutput struct {
	Attributes            map[string]AttributeValue `json:"Attributes,omitempty"`
	ConsumedCapacity      ConsumedCapacity          `json:"ConsumedCapacity,omitempty"`
	ItemCollectionMetrics ItemCollectionMetrics     `json:"ItemCollectionMetrics,omitempty"`
}

type DeleteRequest struct {
	Key map[string]AttributeValue `json:"Key"`
}

type DeleteTableInput struct {
	TableName string `json:"TableName"`
}

type DeleteTableOutput struct {
	TableDescription TableDescription `json:"TableDescription,omitempty"`
}

type DescribeTableInput struct {
	TableName string `json:"TableName"`
}

type DescribeTableOutput struct {
	Table TableDescription `json:"Table,omitempty"`
}

type ExpectedAttributeValue struct {
	AttributeValueList []AttributeValue `json:"AttributeValueList,omitempty"`
	ComparisonOperator string           `json:"ComparisonOperator,omitempty"`
	Exists             bool             `json:"Exists,omitempty"`
	Value              AttributeValue   `json:"Value,omitempty"`
}

type GetItemInput struct {
	AttributesToGet          []string                  `json:"AttributesToGet,omitempty"`
	ConsistentRead           bool                      `json:"ConsistentRead,omitempty"`
	ExpressionAttributeNames map[string]string         `json:"ExpressionAttributeNames,omitempty"`
	Key                      map[string]AttributeValue `json:"Key"`
	ProjectionExpression     string                    `json:"ProjectionExpression,omitempty"`
	ReturnConsumedCapacity   string                    `json:"ReturnConsumedCapacity,omitempty"`
	TableName                string                    `json:"TableName"`
}

type GetItemOutput struct {
	ConsumedCapacity ConsumedCapacity          `json:"ConsumedCapacity,omitempty"`
	Item             map[string]AttributeValue `json:"Item,omitempty"`
}

type GlobalSecondaryIndex struct {
	IndexName             string                `json:"IndexName"`
	KeySchema             []KeySchemaElement    `json:"KeySchema"`
	Projection            Projection            `json:"Projection"`
	ProvisionedThroughput ProvisionedThroughput `json:"ProvisionedThroughput"`
}

type GlobalSecondaryIndexDescription struct {
	IndexName             string                           `json:"IndexName,omitempty"`
	IndexSizeBytes        int                              `json:"IndexSizeBytes,omitempty"`
	IndexStatus           string                           `json:"IndexStatus,omitempty"`
	ItemCount             int                              `json:"ItemCount,omitempty"`
	KeySchema             []KeySchemaElement               `json:"KeySchema,omitempty"`
	Projection            Projection                       `json:"Projection,omitempty"`
	ProvisionedThroughput ProvisionedThroughputDescription `json:"ProvisionedThroughput,omitempty"`
}

type GlobalSecondaryIndexUpdate struct {
	Update UpdateGlobalSecondaryIndexAction `json:"Update,omitempty"`
}

type ItemCollectionMetrics struct {
	ItemCollectionKey   map[string]AttributeValue `json:"ItemCollectionKey,omitempty"`
	SizeEstimateRangeGB []float64                 `json:"SizeEstimateRangeGB,omitempty"`
}

type KeySchemaElement struct {
	AttributeName string `json:"AttributeName"`
	KeyType       string `json:"KeyType"`
}

type KeysAndAttributes struct {
	AttributesToGet          []string                    `json:"AttributesToGet,omitempty"`
	ConsistentRead           bool                        `json:"ConsistentRead,omitempty"`
	ExpressionAttributeNames map[string]string           `json:"ExpressionAttributeNames,omitempty"`
	Keys                     []map[string]AttributeValue `json:"Keys"`
	ProjectionExpression     string                      `json:"ProjectionExpression,omitempty"`
}

type ListTablesInput struct {
	ExclusiveStartTableName string `json:"ExclusiveStartTableName,omitempty"`
	Limit                   int    `json:"Limit,omitempty"`
}

type ListTablesOutput struct {
	LastEvaluatedTableName string   `json:"LastEvaluatedTableName,omitempty"`
	TableNames             []string `json:"TableNames,omitempty"`
}

type LocalSecondaryIndex struct {
	IndexName  string             `json:"IndexName"`
	KeySchema  []KeySchemaElement `json:"KeySchema"`
	Projection Projection         `json:"Projection"`
}

type LocalSecondaryIndexDescription struct {
	IndexName      string             `json:"IndexName,omitempty"`
	IndexSizeBytes int                `json:"IndexSizeBytes,omitempty"`
	ItemCount      int                `json:"ItemCount,omitempty"`
	KeySchema      []KeySchemaElement `json:"KeySchema,omitempty"`
	Projection     Projection         `json:"Projection,omitempty"`
}

type Projection struct {
	NonKeyAttributes []string `json:"NonKeyAttributes,omitempty"`
	ProjectionType   string   `json:"ProjectionType,omitempty"`
}

type ProvisionedThroughput struct {
	ReadCapacityUnits  int `json:"ReadCapacityUnits"`
	WriteCapacityUnits int `json:"WriteCapacityUnits"`
}

type ProvisionedThroughputDescription struct {
	LastDecreaseDateTime   time.Time `json:"LastDecreaseDateTime,omitempty"`
	LastIncreaseDateTime   time.Time `json:"LastIncreaseDateTime,omitempty"`
	NumberOfDecreasesToday int       `json:"NumberOfDecreasesToday,omitempty"`
	ReadCapacityUnits      int       `json:"ReadCapacityUnits,omitempty"`
	WriteCapacityUnits     int       `json:"WriteCapacityUnits,omitempty"`
}

type PutItemInput struct {
	ConditionExpression         string                            `json:"ConditionExpression,omitempty"`
	ConditionalOperator         string                            `json:"ConditionalOperator,omitempty"`
	Expected                    map[string]ExpectedAttributeValue `json:"Expected,omitempty"`
	ExpressionAttributeNames    map[string]string                 `json:"ExpressionAttributeNames,omitempty"`
	ExpressionAttributeValues   map[string]AttributeValue         `json:"ExpressionAttributeValues,omitempty"`
	Item                        map[string]AttributeValue         `json:"Item"`
	ReturnConsumedCapacity      string                            `json:"ReturnConsumedCapacity,omitempty"`
	ReturnItemCollectionMetrics string                            `json:"ReturnItemCollectionMetrics,omitempty"`
	ReturnValues                string                            `json:"ReturnValues,omitempty"`
	TableName                   string                            `json:"TableName"`
}

type PutItemOutput struct {
	Attributes            map[string]AttributeValue `json:"Attributes,omitempty"`
	ConsumedCapacity      ConsumedCapacity          `json:"ConsumedCapacity,omitempty"`
	ItemCollectionMetrics ItemCollectionMetrics     `json:"ItemCollectionMetrics,omitempty"`
}

type PutRequest struct {
	Item map[string]AttributeValue `json:"Item"`
}

type QueryInput struct {
	AttributesToGet           []string                  `json:"AttributesToGet,omitempty"`
	ConditionalOperator       string                    `json:"ConditionalOperator,omitempty"`
	ConsistentRead            bool                      `json:"ConsistentRead,omitempty"`
	ExclusiveStartKey         map[string]AttributeValue `json:"ExclusiveStartKey,omitempty"`
	ExpressionAttributeNames  map[string]string         `json:"ExpressionAttributeNames,omitempty"`
	ExpressionAttributeValues map[string]AttributeValue `json:"ExpressionAttributeValues,omitempty"`
	FilterExpression          string                    `json:"FilterExpression,omitempty"`
	IndexName                 string                    `json:"IndexName,omitempty"`
	KeyConditions             map[string]Condition      `json:"KeyConditions"`
	Limit                     int                       `json:"Limit,omitempty"`
	ProjectionExpression      string                    `json:"ProjectionExpression,omitempty"`
	QueryFilter               map[string]Condition      `json:"QueryFilter,omitempty"`
	ReturnConsumedCapacity    string                    `json:"ReturnConsumedCapacity,omitempty"`
	ScanIndexForward          bool                      `json:"ScanIndexForward,omitempty"`
	Select                    string                    `json:"Select,omitempty"`
	TableName                 string                    `json:"TableName"`
}

type QueryOutput struct {
	ConsumedCapacity ConsumedCapacity            `json:"ConsumedCapacity,omitempty"`
	Count            int                         `json:"Count,omitempty"`
	Items            []map[string]AttributeValue `json:"Items,omitempty"`
	LastEvaluatedKey map[string]AttributeValue   `json:"LastEvaluatedKey,omitempty"`
	ScannedCount     int                         `json:"ScannedCount,omitempty"`
}

type ScanInput struct {
	AttributesToGet           []string                  `json:"AttributesToGet,omitempty"`
	ConditionalOperator       string                    `json:"ConditionalOperator,omitempty"`
	ExclusiveStartKey         map[string]AttributeValue `json:"ExclusiveStartKey,omitempty"`
	ExpressionAttributeNames  map[string]string         `json:"ExpressionAttributeNames,omitempty"`
	ExpressionAttributeValues map[string]AttributeValue `json:"ExpressionAttributeValues,omitempty"`
	FilterExpression          string                    `json:"FilterExpression,omitempty"`
	Limit                     int                       `json:"Limit,omitempty"`
	ProjectionExpression      string                    `json:"ProjectionExpression,omitempty"`
	ReturnConsumedCapacity    string                    `json:"ReturnConsumedCapacity,omitempty"`
	ScanFilter                map[string]Condition      `json:"ScanFilter,omitempty"`
	Segment                   int                       `json:"Segment,omitempty"`
	Select                    string                    `json:"Select,omitempty"`
	TableName                 string                    `json:"TableName"`
	TotalSegments             int                       `json:"TotalSegments,omitempty"`
}

type ScanOutput struct {
	ConsumedCapacity ConsumedCapacity            `json:"ConsumedCapacity,omitempty"`
	Count            int                         `json:"Count,omitempty"`
	Items            []map[string]AttributeValue `json:"Items,omitempty"`
	LastEvaluatedKey map[string]AttributeValue   `json:"LastEvaluatedKey,omitempty"`
	ScannedCount     int                         `json:"ScannedCount,omitempty"`
}

type TableDescription struct {
	AttributeDefinitions   []AttributeDefinition             `json:"AttributeDefinitions,omitempty"`
	CreationDateTime       time.Time                         `json:"CreationDateTime,omitempty"`
	GlobalSecondaryIndexes []GlobalSecondaryIndexDescription `json:"GlobalSecondaryIndexes,omitempty"`
	ItemCount              int                               `json:"ItemCount,omitempty"`
	KeySchema              []KeySchemaElement                `json:"KeySchema,omitempty"`
	LocalSecondaryIndexes  []LocalSecondaryIndexDescription  `json:"LocalSecondaryIndexes,omitempty"`
	ProvisionedThroughput  ProvisionedThroughputDescription  `json:"ProvisionedThroughput,omitempty"`
	TableName              string                            `json:"TableName,omitempty"`
	TableSizeBytes         int                               `json:"TableSizeBytes,omitempty"`
	TableStatus            string                            `json:"TableStatus,omitempty"`
}

type UpdateGlobalSecondaryIndexAction struct {
	IndexName             string                `json:"IndexName"`
	ProvisionedThroughput ProvisionedThroughput `json:"ProvisionedThroughput"`
}

type UpdateItemInput struct {
	AttributeUpdates            map[string]AttributeValueUpdate   `json:"AttributeUpdates,omitempty"`
	ConditionExpression         string                            `json:"ConditionExpression,omitempty"`
	ConditionalOperator         string                            `json:"ConditionalOperator,omitempty"`
	Expected                    map[string]ExpectedAttributeValue `json:"Expected,omitempty"`
	ExpressionAttributeNames    map[string]string                 `json:"ExpressionAttributeNames,omitempty"`
	ExpressionAttributeValues   map[string]AttributeValue         `json:"ExpressionAttributeValues,omitempty"`
	Key                         map[string]AttributeValue         `json:"Key"`
	ReturnConsumedCapacity      string                            `json:"ReturnConsumedCapacity,omitempty"`
	ReturnItemCollectionMetrics string                            `json:"ReturnItemCollectionMetrics,omitempty"`
	ReturnValues                string                            `json:"ReturnValues,omitempty"`
	TableName                   string                            `json:"TableName"`
	UpdateExpression            string                            `json:"UpdateExpression,omitempty"`
}

type UpdateItemOutput struct {
	Attributes            map[string]AttributeValue `json:"Attributes,omitempty"`
	ConsumedCapacity      ConsumedCapacity          `json:"ConsumedCapacity,omitempty"`
	ItemCollectionMetrics ItemCollectionMetrics     `json:"ItemCollectionMetrics,omitempty"`
}

type UpdateTableInput struct {
	GlobalSecondaryIndexUpdates []GlobalSecondaryIndexUpdate `json:"GlobalSecondaryIndexUpdates,omitempty"`
	ProvisionedThroughput       ProvisionedThroughput        `json:"ProvisionedThroughput,omitempty"`
	TableName                   string                       `json:"TableName"`
}

type UpdateTableOutput struct {
	TableDescription TableDescription `json:"TableDescription,omitempty"`
}

type WriteRequest struct {
	DeleteRequest DeleteRequest `json:"DeleteRequest,omitempty"`
	PutRequest    PutRequest    `json:"PutRequest,omitempty"`
}

var _ time.Time // to avoid errors if the time package isn't referenced
