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
		mode     ConditionMode
		expected Expression
	}{
		{
			lhs:  NewPath("foo.yay.cool.rad"),
			rhs:  NewPath("bar"),
			mode: EqualCond,
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
					"#1": aws.String("yay"),
					"#2": aws.String("cool"),
					"#3": aws.String("rad"),
					"#4": aws.String("bar"),
				},
				Expression: "#0.#1.#2.#3 = #4",
			},
		},
		{
			lhs:  NewPath("foo.yay.cool.rad"),
			rhs:  NewValue(5),
			mode: EqualCond,
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
					"#1": aws.String("yay"),
					"#2": aws.String("cool"),
					"#3": aws.String("rad"),
				},
				Values: map[string]*dynamodb.AttributeValue{
					":0": &dynamodb.AttributeValue{
						N: aws.String("5"),
					},
				},
				Expression: "#0.#1.#2.#3 = :0",
			},
		},
		{
			lhs:  NewPath("foo.yay.cool.rad"),
			rhs:  NewPath("baz").Size(),
			mode: EqualCond,
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
					"#1": aws.String("yay"),
					"#2": aws.String("cool"),
					"#3": aws.String("rad"),
					"#4": aws.String("baz"),
				},
				Expression: "#0.#1.#2.#3 = size (#4)",
			},
		},
		{
			lhs:  NewValue(5),
			rhs:  NewPath("bar"),
			mode: EqualCond,
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("bar"),
				},
				Values: map[string]*dynamodb.AttributeValue{
					":0": &dynamodb.AttributeValue{
						N: aws.String("5"),
					},
				},
				Expression: ":0 = #0",
			},
		},
		{
			lhs: NewValue(map[string]int{
				"five": 5,
			}),
			rhs:  NewPath("bar"),
			mode: EqualCond,
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("bar"),
				},
				Values: map[string]*dynamodb.AttributeValue{
					":0": &dynamodb.AttributeValue{
						M: map[string]*dynamodb.AttributeValue{
							"five": &dynamodb.AttributeValue{
								N: aws.String("5"),
							},
						},
					},
				},
				Expression: ":0 = #0",
			},
		},
		{
			lhs: NewValue(map[string]int{
				"five": 5,
			}),
			rhs:  NewValue(5),
			mode: EqualCond,
			expected: Expression{
				Values: map[string]*dynamodb.AttributeValue{
					":0": &dynamodb.AttributeValue{
						M: map[string]*dynamodb.AttributeValue{
							"five": &dynamodb.AttributeValue{
								N: aws.String("5"),
							},
						},
					},
					":1": &dynamodb.AttributeValue{
						N: aws.String("5"),
					},
				},
				Expression: ":0 = :1",
			},
		},
		{
			lhs: NewValue(map[string]int{
				"five": 5,
			}),
			rhs:  NewPath("baz").Size(),
			mode: EqualCond,
			expected: Expression{
				Values: map[string]*dynamodb.AttributeValue{
					":0": &dynamodb.AttributeValue{
						M: map[string]*dynamodb.AttributeValue{
							"five": &dynamodb.AttributeValue{
								N: aws.String("5"),
							},
						},
					},
				},
				Names: map[string]*string{
					"#0": aws.String("baz"),
				},
				Expression: ":0 = size (#0)",
			},
		},
		{
			lhs:  NewPath("foo[1]").Size(),
			rhs:  NewPath("bar"),
			mode: EqualCond,
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
					"#1": aws.String("bar"),
				},
				Expression: "size (#0[1]) = #1",
			},
		},
		{
			lhs:  NewPath("foo[1]").Size(),
			rhs:  NewValue(5),
			mode: EqualCond,
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
				},
				Values: map[string]*dynamodb.AttributeValue{
					":0": &dynamodb.AttributeValue{
						N: aws.String("5"),
					},
				},
				Expression: "size (#0[1]) = :0",
			},
		},
		{
			lhs:  NewPath("foo[1]").Size(),
			rhs:  NewPath("baz").Size(),
			mode: EqualCond,
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
					"#1": aws.String("baz"),
				},
				Expression: "size (#0[1]) = size (#1)",
			},
		},
		{
			lhs:  NewPath("foo.bar.baz").Size(),
			rhs:  NewPath("bar.qux.foo").Size(),
			mode: EqualCond,
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
					"#1": aws.String("bar"),
					"#2": aws.String("baz"),
					"#3": aws.String("qux"),
				},
				Expression: "size (#0.#1.#2) = size (#1.#3.#0)",
			},
		},
	}
	for testNumber, c := range cases {
		exprNode, err := c.lhs.Equal(c.rhs).buildCondition()
		if err != nil {
			t.Errorf("TestEquals Test Number %#v: Unexpected Error %#v", testNumber, err)
		}
		expr, err := exprNode.buildExprNodes(&aliasList{})
		if err != nil {
			t.Errorf("TestEquals Test Number %#v: Unexpected Error %#v", testNumber, err)
		}

		if reflect.DeepEqual(expr, c.expected) != true {
			t.Errorf("TestEquals Test Number %#v: Expected %#v, got %#v", testNumber, c.expected, expr)
		}

	}
}

func TestBuildCondition(t *testing.T) {
	cases := []struct {
		input                 ConditionBuilder
		expected              ExprNode
		buildListOperandError bool
		noMatchError          bool
		operandNumberError    bool
		conditionNumberError  bool
	}{
		{
			input:        ConditionBuilder{},
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
		expr, err := c.input.buildCondition()

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
