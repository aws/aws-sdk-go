package expression

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type testCompare interface {
	Equal(right OperandBuilder) ConditionBuilder
}

//Compare
//Equal
func TestCompare(t *testing.T) {
	cases := []struct {
		lhs      testCompare
		rhs      OperandBuilder
		expected ConditionBuilder
	}{
		{
			lhs: NewPath("foo.yay.cool.rad"),
			rhs: NewPath("bar"),
			expected: ConditionBuilder{
				operandList: []OperandExpression{
					OperandExpression{
						Mode: PathOpe,
						path: "foo.yay.cool.rad",
					},
					OperandExpression{
						Mode: PathOpe,
						path: "bar",
					},
				},
			},
		},
		{
			lhs: NewPath("foo.yay.cool.rad"),
			rhs: NewValue(5),
			expected: ConditionBuilder{
				operandList: []OperandExpression{
					OperandExpression{
						Mode: PathOpe,
						path: "foo.yay.cool.rad",
					},
					OperandExpression{
						Mode: ValueOpe,
						value: dynamodb.AttributeValue{
							N: aws.String("5"),
						},
					},
				},
			},
		},
		{
			lhs: NewPath("foo.yay.cool.rad"),
			rhs: NewPath("baz").Size(),
			expected: ConditionBuilder{
				operandList: []OperandExpression{
					OperandExpression{
						Mode: PathOpe,
						path: "foo.yay.cool.rad",
					},
					OperandExpression{
						Mode: SizeOpe,
						path: "baz",
					},
				},
			},
		},
		{
			lhs: NewValue(5),
			rhs: NewPath("bar"),
			expected: ConditionBuilder{
				operandList: []OperandExpression{
					OperandExpression{
						Mode: ValueOpe,
						value: dynamodb.AttributeValue{
							N: aws.String("5"),
						},
					},
					OperandExpression{
						Mode: PathOpe,
						path: "bar",
					},
				},
			},
		},
		{
			lhs: NewValue(map[string]int{
				"five": 5,
			}),
			rhs: NewPath("bar"),
			expected: ConditionBuilder{
				operandList: []OperandExpression{
					OperandExpression{
						Mode: ValueOpe,
						value: dynamodb.AttributeValue{
							M: map[string]*dynamodb.AttributeValue{
								"five": &dynamodb.AttributeValue{
									N: aws.String("5"),
								},
							},
						},
					},
					OperandExpression{
						Mode: PathOpe,
						path: "bar",
					},
				},
			},
		},
		{
			lhs: NewValue(map[string]int{
				"five": 5,
			}),
			rhs: NewValue(5),
			expected: ConditionBuilder{
				operandList: []OperandExpression{
					OperandExpression{
						Mode: ValueOpe,
						value: dynamodb.AttributeValue{
							M: map[string]*dynamodb.AttributeValue{
								"five": &dynamodb.AttributeValue{
									N: aws.String("5"),
								},
							},
						},
					},
					OperandExpression{
						Mode: ValueOpe,
						value: dynamodb.AttributeValue{
							N: aws.String("5"),
						},
					},
				},
			},
		},
		{
			lhs: NewValue(map[string]int{
				"five": 5,
			}),
			rhs: NewPath("baz").Size(),
			expected: ConditionBuilder{
				operandList: []OperandExpression{
					OperandExpression{
						Mode: ValueOpe,
						value: dynamodb.AttributeValue{
							M: map[string]*dynamodb.AttributeValue{
								"five": &dynamodb.AttributeValue{
									N: aws.String("5"),
								},
							},
						},
					},
					OperandExpression{
						Mode: SizeOpe,
						path: "baz",
					},
				},
			},
		},
		{
			lhs: NewPath("foo[1]").Size(),
			rhs: NewPath("bar"),
			expected: ConditionBuilder{
				operandList: []OperandExpression{
					OperandExpression{
						Mode: SizeOpe,
						path: "foo[1]",
					},
					OperandExpression{
						Mode: PathOpe,
						path: "bar",
					},
				},
			},
		},
		{
			lhs: NewPath("foo[1]").Size(),
			rhs: NewValue(5),
			expected: ConditionBuilder{
				operandList: []OperandExpression{
					OperandExpression{
						Mode: SizeOpe,
						path: "foo[1]",
					},
					OperandExpression{
						Mode: ValueOpe,
						value: dynamodb.AttributeValue{
							N: aws.String("5"),
						},
					},
				},
			},
		},
		{
			lhs: NewPath("foo[1]").Size(),
			rhs: NewPath("baz").Size(),
			expected: ConditionBuilder{
				operandList: []OperandExpression{
					OperandExpression{
						Mode: SizeOpe,
						path: "foo[1]",
					},
					OperandExpression{
						Mode: SizeOpe,
						path: "baz",
					},
				},
			},
		},
		{
			lhs: NewPath("foo.bar.baz").Size(),
			rhs: NewPath("bar.qux.foo").Size(),
			expected: ConditionBuilder{
				operandList: []OperandExpression{
					OperandExpression{
						Mode: SizeOpe,
						path: "foo.bar.baz",
					},
					OperandExpression{
						Mode: SizeOpe,
						path: "bar.qux.foo",
					},
				},
			},
		},
	}
	for testNumber, c := range cases {

		c.expected.Mode = EqualCond
		expr, err := c.lhs.Equal(c.rhs).BuildCondition()
		if err != nil {
			t.Errorf("TestEquals Test Number %#v: Unexpected Error %#v", testNumber, err)
		}
		expected, err := c.expected.BuildCondition()
		if err != nil {
			t.Errorf("TestEquals Test Number %#v: Unexpected Error %#v", testNumber, err)
		}

		if reflect.DeepEqual(expr, expected) != true {
			t.Errorf("TestEquals Test Number %#v: Expected %#v, got %#v", testNumber, expected, expr)
		}

	}
}

func TestBuildCondition(t *testing.T) {
	cases := []struct {
		input                 ConditionBuilder
		expected              Expression
		buildListOperandError bool
		noMatchError          bool
		operandNumberError    bool
		conditionNumberError  bool
	}{
		{
			input: ConditionBuilder{
				operandList: []OperandExpression{
					OperandExpression{
						Mode: PathOpe,
						path: "",
					},
				},
			},
			expected:              Expression{},
			buildListOperandError: true,
		},
		{
			input: ConditionBuilder{
				Mode: 0,
			},
			noMatchError: true,
		},
		{
			input: ConditionBuilder{
				Mode: EqualCond,
			},
			operandNumberError: true,
		},
		{
			input: ConditionBuilder{
				Mode: EqualCond,
				conditionList: []ConditionBuilder{
					ConditionBuilder{},
				},
			},
			conditionNumberError: true,
		},
	}

	for testNumber, c := range cases {
		expr, err := c.input.BuildCondition()

		if c.buildListOperandError {
			if err == nil {
				t.Errorf("TestBuildCondition Test Number %#v: Expected list operand error but got no error", testNumber)
			} else {
				continue
			}
		}

		if c.noMatchError {
			if err == nil {
				t.Errorf("TestBuildCondition Test Number %#v: Expected no matching mode error but got no error", testNumber)
			} else {
				continue
			}
		}

		if c.operandNumberError {
			if err == nil {
				t.Errorf("TestBuildCondition Test Number %#v: Expected operand number error but got no error", testNumber)
			} else {
				continue
			}
		}

		if c.conditionNumberError {
			if err == nil {
				t.Errorf("TestBuildCondition Test Number %#v: Expected condition number error but got no error", testNumber)
			} else {
				continue
			}
		}

		if err != nil {
			t.Errorf("TestBuildCondition Test Number %#v: Unexpected Error %#v", testNumber, err)
		}

		if reflect.DeepEqual(expr, c.expected) != true {
			t.Errorf("TestBuildCondition Test Number %#v: Expected %#v, got %#v", testNumber, c.expected, expr)
		}
	}
}

// If there is time implement mapEquals
