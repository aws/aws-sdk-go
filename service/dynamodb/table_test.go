package dynamodb

import (
	"fmt"
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type TableSuite struct {
	TableName  string
	PrimaryKey *PrimaryKey
	table      *Table
}

var table_suite = &TableSuite{
	TableName: "TestTable",
	PrimaryKey: &PrimaryKey{
		Type:        Hash,
		HashKeyName: "Id",
		HashKeyType: StringType,
	},
}

var table_suite_with_range_key = &TableSuite{
	TableName: "TestTable",
	PrimaryKey: &PrimaryKey{
		Type:         HashAndRange,
		HashKeyName:  "Id",
		HashKeyType:  StringType,
		RangeKeyName: "Range",
		RangeKeyType: NumberType,
	},
}

var _ = check.Suite(table_suite)
var _ = check.Suite(table_suite_with_range_key)

func signWithHeader(req *aws.Request) {
	headers := map[string]string{
		"X-Amz-Algorithm":     "",
		"X-Amz-Credential":    "",
		"X-Amz-SignedHeaders": "",
		"X-Amz-Signature":     "",
	}
	query := req.HTTPRequest.URL.Query()
	for header, _ := range headers {
		v := query.Get(header)
		if v == "" {
			req.Error = fmt.Errorf("'%s' was not found in the query string", header)
			return
		}
		headers[header] = v
	}
	authorization := fmt.Sprintf("%s Credential=%s, SignedHeaders=%s, Signature=%s",
		headers["X-Amz-Algorithm"], headers["X-Amz-Credential"],
		headers["X-Amz-SignedHeaders"], headers["X-Amz-Signature"])
	req.HTTPRequest.Header.Set("Authorization", authorization)
}

func (s *TableSuite) SetUpSuite(c *check.C) {
	creds := aws.Creds("DUMMY_KEY", "DUMMY_SECRET", "")
	config := &DynamoDBConfig{
		Config: &aws.Config{
			Credentials: creds,
			Endpoint:    "http://localhost:8000",
		},
	}
	db := New(config)
	db.Handlers.Sign.PushBack(signWithHeader) // required for DynamoDB Local

	s.table = db.NewTable(s.TableName, s.PrimaryKey)
	s.TearDownSuite(c) // cleanup from any previous runs

	err := s.table.CreateTable(&ProvisionedThroughput{
		ReadCapacityUnits:  aws.Long(1),
		WriteCapacityUnits: aws.Long(1),
	})
	if err != nil {
		c.Fatal(err)
	}
}

func (s *TableSuite) TearDownSuite(c *check.C) {
	if s.table == nil {
		return
	}
	s.table.DeleteTable()
}

func (s *TableSuite) TestPutGetDeleteItem(c *check.C) {
	key := &Key{HashKey: "Hash"}
	if s.PrimaryKey.Type == HashAndRange {
		key.RangeKey = "1"
	}

	type myInnterStruct struct {
		List []interface{}
	}
	type myStruct struct {
		Attr1  string
		Attr2  int64
		Nested myInnterStruct
	}
	in := myStruct{
		Attr1: "Attr1Val",
		Attr2: 1000000,
		Nested: myInnterStruct{
			List: []interface{}{true, false, nil, "some string", 3.14},
		},
	}

	// Put
	if err := s.table.PutItem(key, in); err != nil {
		c.Fatal(err)
	}

	// Get
	var out myStruct
	if err := s.table.GetItem(key, &out); err != nil {
		c.Fatal(err)
	}
	c.Assert(out.Attr1, check.DeepEquals, in.Attr1)
	c.Assert(out.Attr2, check.DeepEquals, in.Attr2)
	for i := range in.Nested.List {
		c.Assert(out.Nested.List[i], check.DeepEquals, in.Nested.List[i])
	}
	c.Assert(out.Nested.List, check.DeepEquals, in.Nested.List)

	// Delete
	if err := s.table.DeleteItem(key); err != nil {
		c.Fatal(err)
	}
	if err := s.table.GetItem(key, &out); err == nil {
		c.Fatalf("Expected error, item should have been deleted, %#v", out)
	}
}
