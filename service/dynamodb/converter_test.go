package dynamodb

import (
	"math"
	"reflect"
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
)

type mySimpleStruct struct {
	String  string
	Int     int
	Uint    uint
	Float32 float32
	Float64 float64
	Bool    bool
	Null    *interface{}
}

type myComplexStruct struct {
	Simple []mySimpleStruct
}

type converterTestInput struct {
	input     interface{}
	expected  map[string]*AttributeValue
	inputType string // "enum" of types
}

var trueValue = true
var falseValue = false

var converterTestInputs = []converterTestInput{
	// Scalar tests
	converterTestInput{
		input:    map[string]interface{}{"string": "some string"},
		expected: map[string]*AttributeValue{"string": &AttributeValue{S: aws.String("some string")}},
	},
	converterTestInput{
		input:    map[string]interface{}{"bool": true},
		expected: map[string]*AttributeValue{"bool": &AttributeValue{BOOL: &trueValue}},
	},
	converterTestInput{
		input:    map[string]interface{}{"bool": false},
		expected: map[string]*AttributeValue{"bool": &AttributeValue{BOOL: &falseValue}},
	},
	converterTestInput{
		input:    map[string]interface{}{"null": nil},
		expected: map[string]*AttributeValue{"null": &AttributeValue{NULL: &trueValue}},
	},
	converterTestInput{
		input:    map[string]interface{}{"float": 3.14},
		expected: map[string]*AttributeValue{"float": &AttributeValue{N: aws.String("3.14")}},
	},
	converterTestInput{
		input:    map[string]interface{}{"float": math.MaxFloat32},
		expected: map[string]*AttributeValue{"float": &AttributeValue{N: aws.String("340282346638528860000000000000000000000")}},
	},
	converterTestInput{
		input:    map[string]interface{}{"float": math.MaxFloat64},
		expected: map[string]*AttributeValue{"float": &AttributeValue{N: aws.String("179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")}},
	},
	converterTestInput{
		input:    map[string]interface{}{"int": int(12)},
		expected: map[string]*AttributeValue{"int": &AttributeValue{N: aws.String("12")}},
	},
	// List
	converterTestInput{
		input: map[string]interface{}{"list": []interface{}{"a string", 12, 3.14, true, nil, false}},
		expected: map[string]*AttributeValue{
			"list": &AttributeValue{
				L: []*AttributeValue{
					&AttributeValue{S: aws.String("a string")},
					&AttributeValue{N: aws.String("12")},
					&AttributeValue{N: aws.String("3.14")},
					&AttributeValue{BOOL: &trueValue},
					&AttributeValue{NULL: &trueValue},
					&AttributeValue{BOOL: &falseValue},
				},
			},
		},
	},
	// Map
	converterTestInput{
		input: map[string]interface{}{"map": map[string]interface{}{"nestedint": 12}},
		expected: map[string]*AttributeValue{
			"map": &AttributeValue{
				M: &map[string]*AttributeValue{
					"nestedint": &AttributeValue{
						N: aws.String("12"),
					},
				},
			},
		},
	},
	// Structs
	converterTestInput{
		input: mySimpleStruct{},
		expected: map[string]*AttributeValue{
			"Bool":    &AttributeValue{BOOL: &falseValue},
			"Float32": &AttributeValue{N: aws.String("0")},
			"Float64": &AttributeValue{N: aws.String("0")},
			"Int":     &AttributeValue{N: aws.String("0")},
			"Null":    &AttributeValue{NULL: &trueValue},
			"String":  &AttributeValue{S: aws.String("")},
			"Uint":    &AttributeValue{N: aws.String("0")},
		},
		inputType: "mySimpleStruct",
	},
	converterTestInput{
		input: myComplexStruct{},
		expected: map[string]*AttributeValue{
			"Simple": &AttributeValue{NULL: &trueValue},
		},
		inputType: "myComplexStruct",
	},
	converterTestInput{
		input: myComplexStruct{Simple: []mySimpleStruct{mySimpleStruct{Int: -2}, mySimpleStruct{Uint: 5}}},
		expected: map[string]*AttributeValue{
			"Simple": &AttributeValue{
				L: []*AttributeValue{
					&AttributeValue{
						M: &map[string]*AttributeValue{
							"Bool":    &AttributeValue{BOOL: &falseValue},
							"Float32": &AttributeValue{N: aws.String("0")},
							"Float64": &AttributeValue{N: aws.String("0")},
							"Int":     &AttributeValue{N: aws.String("-2")},
							"Null":    &AttributeValue{NULL: &trueValue},
							"String":  &AttributeValue{S: aws.String("")},
							"Uint":    &AttributeValue{N: aws.String("0")},
						},
					},
					&AttributeValue{
						M: &map[string]*AttributeValue{
							"Bool":    &AttributeValue{BOOL: &falseValue},
							"Float32": &AttributeValue{N: aws.String("0")},
							"Float64": &AttributeValue{N: aws.String("0")},
							"Int":     &AttributeValue{N: aws.String("0")},
							"Null":    &AttributeValue{NULL: &trueValue},
							"String":  &AttributeValue{S: aws.String("")},
							"Uint":    &AttributeValue{N: aws.String("5")},
						},
					},
				},
			},
		},
		inputType: "myComplexStruct",
	},
}

func TestConvertTo(t *testing.T) {
	for _, test := range converterTestInputs {
		testConvertTo(t, test.input, test.expected)
	}
}

func testConvertTo(t *testing.T, in interface{}, expected map[string]*AttributeValue) {
	actual, err := ConvertTo(in)
	if err != nil {
		t.Fatal(err)
	}
	compareObjects(t, expected, actual)
}

func TestConvertFrom(t *testing.T) {
	// Using the same inputs from TestConvertTo, test the reverse mapping.
	for _, test := range converterTestInputs {
		switch test.inputType {
		case "mySimpleStruct":
			testConvertFromSimpleStruct(t, test.expected, test.input)
		case "myComplexStruct":
			testConvertFromComplexStruct(t, test.expected, test.input)
		default:
			testConvertFrom(t, test.expected, test.input)
		}
	}
}

func testConvertFrom(t *testing.T, in map[string]*AttributeValue, expected interface{}) {
	var actual map[string]interface{}
	if err := ConvertFrom(in, &actual); err != nil {
		t.Fatal(err)
	}
	compareObjects(t, expected, actual)
}

func testConvertFromSimpleStruct(t *testing.T, in map[string]*AttributeValue, expected interface{}) {
	var actual mySimpleStruct
	if err := ConvertFrom(in, &actual); err != nil {
		t.Fatal(err)
	}
	compareObjects(t, expected, actual)
}

func testConvertFromComplexStruct(t *testing.T, in map[string]*AttributeValue, expected interface{}) {
	var actual myComplexStruct
	if err := ConvertFrom(in, &actual); err != nil {
		t.Fatal(err)
	}
	compareObjects(t, expected, actual)
}

func compareObjects(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %s, got %s", awsutil.StringValue(expected), awsutil.StringValue(actual))
	}
}
