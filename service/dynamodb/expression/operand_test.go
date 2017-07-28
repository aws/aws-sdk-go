package expression

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestBuildList(t *testing.T) {
	cases := []struct {
		input               OperandBuilder
		expected            aliasList
		incompletePathError bool
		emptyPathError      bool
	}{
		{
			input: NewPath("foo"),
			expected: aliasList{
				NamesList: []string{
					"foo",
				},
				ValuesList: nil,
			},
		},
		{
			input: NewValue(5),
			expected: aliasList{
				NamesList:  nil,
				ValuesList: nil,
			},
		},
		{
			input: NewPath("foo.bar[7].baz"),
			expected: aliasList{
				NamesList: []string{
					"foo",
					"bar",
					"baz",
				},
				ValuesList: nil,
			},
		},
		{
			input:          NewPath(""),
			expected:       aliasList{},
			emptyPathError: true,
		},
		{
			input:               NewPath("foo..bar"),
			expected:            aliasList{},
			incompletePathError: true,
		},
		{
			input: NewPath("foo").Size(),
			expected: aliasList{
				NamesList: []string{
					"foo",
				},
				ValuesList: nil,
			},
		},
	}

	for testNumber, c := range cases {
		oe, err := c.input.BuildOperand()
		if err != nil {
			t.Errorf("TestBuildList Test Number %#v: Unexpected Error %#v", testNumber, err)
		}

		al, err := oe.buildList()

		if c.emptyPathError {
			if err == nil {
				t.Errorf("TestBuildList Test Number %#v: Expected empty path error but got no error", testNumber)
			} else {
				continue
			}
		}
		if c.incompletePathError {
			if err == nil {
				t.Errorf("TestBuildList Test Number %#v: Expected incomplete path error but got no error", testNumber)
			} else {
				continue
			}
		}

		if err != nil {
			t.Errorf("TestBuildList Test Number %#v: Unexpected Error %#v", testNumber, err)
		}

		if reflect.DeepEqual(al, c.expected) != true {
			t.Errorf("TestBuildList Test Number %#v: Expected %#v, got %#v", testNumber, c.expected, al)
		}
	}
}

func TestBuildExpression(t *testing.T) {
	cases := []struct {
		input          OperandBuilder
		expected       Expression
		emptyPathError bool
		alError        bool
	}{
		{
			input: NewPath("foo"),
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
				},
				Expression: "#0",
			},
		},
		{
			input: NewValue(5),
			expected: Expression{
				Values: map[string]*dynamodb.AttributeValue{
					":0": &dynamodb.AttributeValue{
						N: aws.String("5"),
					},
				},
				Expression: ":0",
			},
		},
		{
			input: NewPath("foo.bar"),
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
					"#1": aws.String("bar"),
				},
				Expression: "#0.#1",
			},
		},
		{
			input: NewPath("foo.bar[0].baz"),
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
					"#1": aws.String("bar"),
					"#2": aws.String("baz"),
				},
				Expression: "#0.#1[0].#2",
			},
		},
		{
			input: NewPath("foo").Size(),
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
				},
				Expression: "size (#0)",
			},
		},
		{
			input:          NewPath(""),
			expected:       Expression{},
			emptyPathError: true,
		},
		{
			input:    NewPath("foo"),
			expected: Expression{},
			alError:  true,
		},
		{
			input:    NewPath("foo").Size(),
			expected: Expression{},
			alError:  true,
		},
	}

	for testNumber, c := range cases {
		oe, err := c.input.BuildOperand()
		if err != nil {
			t.Error(err)
		}

		al, err := oe.buildList()

		if c.emptyPathError {
			if err == nil {
				t.Errorf("TestBuildExpression Test Number %#v: Expected counter error but got no error", testNumber)
			} else {
				continue
			}
		}

		if err != nil {
			t.Error(err)
		}

		if c.alError {
			al.NamesList = al.NamesList[1:]
		}

		operand, err := oe.buildExpression(al)

		if c.alError {
			if err == nil {
				t.Errorf("TestBuildExpression Test Number %#v: Expected List error but got no error", testNumber)
			} else {
				continue
			}
		}

		if err != nil {
			t.Errorf("TestBuildExpression Test Number %#v: Unexpected Error %#v", testNumber, err)
		}

		if operand.Expression != c.expected.Expression {
			t.Errorf("TestBuildExpression Test Number %#v: BuildOperand returned an unexpected Expression string %#v, expected %#v\n", testNumber, operand.Expression, c.expected.Expression)
		}

		if reflect.DeepEqual(c.expected.Names, operand.Names) != true {
			t.Errorf("TestBuildExpression Test Number %#v: BuildOperand returned an unexpected Name Map %#v, expected %#v\n", testNumber, operand.Names, c.expected.Names)
		}

		if reflect.DeepEqual(c.expected.Values, operand.Values) != true {
			t.Errorf("TestBuildExpression Test Number %#v: BuildOperand returned an unexpected Name Map %#v, expected %#v\n", testNumber, operand.Values, c.expected.Values)
		}
	}
}
