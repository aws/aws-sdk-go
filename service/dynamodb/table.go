package dynamodb

import (
	"fmt"

	"github.com/awslabs/aws-sdk-go/aws"
)

type Table struct {
	db         *DynamoDB
	Name       string
	PrimaryKey *PrimaryKey
}

type PrimaryKeyType string

const (
	Hash         PrimaryKeyType = "Hash"
	HashAndRange                = "HashAndRange"
)

const (
	StringType string = "S"
	NumberType        = "N"
)

type PrimaryKey struct {
	Type         PrimaryKeyType
	HashKeyName  string
	HashKeyType  string
	RangeKeyName string
	RangeKeyType string
}

type Key struct {
	HashKey  string
	RangeKey string
}

func (db *DynamoDB) NewTable(name string, pk *PrimaryKey) *Table {
	return &Table{
		db:         db,
		Name:       name,
		PrimaryKey: pk,
	}
}

func (t *Table) CreateTable(throughput *ProvisionedThroughput) error {
	attributes := make([]AttributeDefinition, 1, 2)
	attributes[0] = AttributeDefinition{
		AttributeName: &t.PrimaryKey.HashKeyName,
		AttributeType: &t.PrimaryKey.HashKeyType,
	}
	if t.PrimaryKey.hasRange() {
		attributes = append(attributes, AttributeDefinition{
			AttributeName: &t.PrimaryKey.RangeKeyName,
			AttributeType: &t.PrimaryKey.RangeKeyType,
		})
	}

	keySchema := make([]KeySchemaElement, 1, 2)
	keySchema[0] = KeySchemaElement{
		AttributeName: &t.PrimaryKey.HashKeyName,
		KeyType:       aws.String("HASH"),
	}
	if t.PrimaryKey.hasRange() {
		keySchema = append(keySchema, KeySchemaElement{
			AttributeName: &t.PrimaryKey.RangeKeyName,
			KeyType:       aws.String("RANGE"),
		})
	}

	input := &CreateTableInput{
		TableName:             &t.Name,
		AttributeDefinitions:  attributes,
		ProvisionedThroughput: throughput,
		KeySchema:             keySchema,
	}

	_, err := t.db.CreateTable(input)
	return err
}

func (t *Table) DeleteTable() error {
	input := &DeleteTableInput{
		TableName: &t.Name,
	}
	_, err := t.db.DeleteTable(input)
	return err
}

func (t *Table) PutItem(key *Key, data interface{}) error {
	item, err := Marshal(data)
	if err != nil {
		return err
	}
	item, err = t.addPrimaryKey(item, key)
	if err != nil {
		return err
	}
	input := &PutItemInput{
		Item:      item,
		TableName: &t.Name,
	}
	_, err = t.db.PutItem(input)
	return err
}

func (t *Table) GetItem(key *Key, v interface{}) error {
	item, err := t.addPrimaryKey(nil, key)
	if err != nil {
		return err
	}
	input := &GetItemInput{
		Key:       item,
		TableName: &t.Name,
	}
	output, err := t.db.GetItem(input)
	if err != nil {
		return err
	}
	if output.Item == nil {
		return fmt.Errorf("Item not found")
	}
	err = Unmarshal(output.Item, v)
	return err
}

func (t *Table) DeleteItem(key *Key) error {
	item, err := t.addPrimaryKey(nil, key)
	if err != nil {
		return err
	}
	input := &DeleteItemInput{
		Key:       item,
		TableName: &t.Name,
	}
	_, err = t.db.DeleteItem(input)
	return err
}

func (t *Table) addPrimaryKey(item map[string]AttributeValue, key *Key) (map[string]AttributeValue, error) {
	if item == nil {
		item = make(map[string]AttributeValue)
	}
	a, err := t.PrimaryKey.attributeValue(t.PrimaryKey.HashKeyType, key.HashKey)
	if err != nil {
		return nil, err
	}
	item[t.PrimaryKey.HashKeyName] = a
	if t.PrimaryKey.hasRange() {
		a, err := t.PrimaryKey.attributeValue(t.PrimaryKey.RangeKeyType, key.RangeKey)
		if err != nil {
			return nil, err
		}
		item[t.PrimaryKey.RangeKeyName] = a
	}
	return item, nil
}

func (pk *PrimaryKey) attributeValue(t, v string) (a AttributeValue, err error) {
	switch t {
	case StringType:
		a.S = aws.String(v)
	case NumberType:
		a.N = aws.String(v)
	default:
		err = fmt.Errorf("'%s' is not a valid primary key type", t)
	}
	return
}

func (pk *PrimaryKey) hasRange() bool {
	return pk.Type == HashAndRange
}
