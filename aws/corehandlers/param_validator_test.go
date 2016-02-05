package corehandlers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/corehandlers"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/stretchr/testify/require"
)

var testSvc = func() *client.Client {
	s := &client.Client{
		Config: aws.Config{},
		ClientInfo: metadata.ClientInfo{
			ServiceName: "mock-service",
			APIVersion:  "2015-01-01",
		},
	}
	return s
}()

type StructShape struct {
	RequiredList   []*ConditionalStructShape          `required:"true"`
	RequiredMap    map[string]*ConditionalStructShape `required:"true"`
	RequiredBool   *bool                              `required:"true"`
	OptionalStruct *ConditionalStructShape

	hiddenParameter *string
	_               struct{}
}

type ConditionalStructShape struct {
	Name *string `required:"true"`
	_    struct{}
}

func TestNoErrors(t *testing.T) {
	input := &StructShape{
		RequiredList: []*ConditionalStructShape{},
		RequiredMap: map[string]*ConditionalStructShape{
			"key1": {Name: aws.String("Name")},
			"key2": {Name: aws.String("Name")},
		},
		RequiredBool:   aws.Bool(true),
		OptionalStruct: &ConditionalStructShape{Name: aws.String("Name")},
	}

	req := testSvc.NewRequest(&request.Operation{}, input, nil)
	corehandlers.ValidateParametersHandler.Fn(req)
	require.NoError(t, req.Error)
}

func TestMissingRequiredParameters(t *testing.T) {
	input := &StructShape{}
	req := testSvc.NewRequest(&request.Operation{}, input, nil)
	corehandlers.ValidateParametersHandler.Fn(req)

	require.Error(t, req.Error)
	assert.Equal(t, "InvalidParameter", req.Error.(awserr.Error).Code())
	assert.Equal(t, "3 validation errors:\n- missing required parameter: RequiredList\n- missing required parameter: RequiredMap\n- missing required parameter: RequiredBool", req.Error.(awserr.Error).Message())
}

func TestNestedMissingRequiredParameters(t *testing.T) {
	input := &StructShape{
		RequiredList: []*ConditionalStructShape{{}},
		RequiredMap: map[string]*ConditionalStructShape{
			"key1": {Name: aws.String("Name")},
			"key2": {},
		},
		RequiredBool:   aws.Bool(true),
		OptionalStruct: &ConditionalStructShape{},
	}

	req := testSvc.NewRequest(&request.Operation{}, input, nil)
	corehandlers.ValidateParametersHandler.Fn(req)

	require.Error(t, req.Error)
	assert.Equal(t, "InvalidParameter", req.Error.(awserr.Error).Code())
	assert.Equal(t, "3 validation errors:\n- missing required parameter: RequiredList[0].Name\n- missing required parameter: RequiredMap[\"key2\"].Name\n- missing required parameter: OptionalStruct.Name", req.Error.(awserr.Error).Message())
}

type testInput struct {
	StringField string            `min:"5"`
	PtrStrField *string           `min:"2"`
	ListField   []string          `min:"3"`
	MapField    map[string]string `min:"4"`
}

var testsFieldMin = []struct {
	err awserr.Error
	in  testInput
}{
	{
		err: awserr.New("InvalidParameter", "1 validation errors:\n- field too short, minimum length 5: StringField", nil),
		in:  testInput{StringField: "abcd"},
	},
	{
		err: awserr.New("InvalidParameter", "2 validation errors:\n- field too short, minimum length 5: StringField\n- field too short, minimum length 3: ListField", nil),
		in:  testInput{StringField: "abcd", ListField: []string{"a", "b"}},
	},
	{
		err: awserr.New("InvalidParameter", "3 validation errors:\n- field too short, minimum length 5: StringField\n- field too short, minimum length 3: ListField\n- field too short, minimum length 4: MapField", nil),
		in:  testInput{StringField: "abcd", ListField: []string{"a", "b"}, MapField: map[string]string{"a": "a", "b": "b"}},
	},
	{
		err: awserr.New("InvalidParameter", "1 validation errors:\n- field too short, minimum length 2: PtrStrField", nil),
		in:  testInput{StringField: "abcde", PtrStrField: aws.String("v")},
	},
	{
		err: nil,
		in: testInput{StringField: "abcde", PtrStrField: aws.String("value"),
			ListField: []string{"a", "b", "c"}, MapField: map[string]string{"a": "a", "b": "b", "c": "c", "d": "d"}},
	},
}

func TestValidateFieldMinParameter(t *testing.T) {
	for i, c := range testsFieldMin {
		req := testSvc.NewRequest(&request.Operation{}, &c.in, nil)
		corehandlers.ValidateParametersHandler.Fn(req)

		require.Equal(t, c.err, req.Error, "%d case failed", i)
	}
}

func BenchmarkValidateAny(b *testing.B) {
	input := &kinesis.PutRecordsInput{
		StreamName: aws.String("stream"),
	}
	for i := 0; i < 100; i++ {
		record := &kinesis.PutRecordsRequestEntry{
			Data:         make([]byte, 10000),
			PartitionKey: aws.String("partition"),
		}
		input.Records = append(input.Records, record)
	}

	req, _ := kinesis.New(session.New()).PutRecordsRequest(input)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		corehandlers.ValidateParametersHandler.Fn(req)
		if err := req.Error; err != nil {
			b.Fatalf("validation failed: %v", err)
		}
	}
}
