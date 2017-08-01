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
		name     string
		lhs      testCompare
		rhs      OperandBuilder
		mode     ConditionMode
		expected Expression
	}{
		{
			name: "nested path with path",
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
			name: "nested path with value",
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
			name: "nested path with path size",
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
			name: "value with path",
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
			name: "nested value with path",
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
			name: "nested value with value",
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
			name: "nested value with path size",
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
			name: "path size with path",
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
			name: "path size with value",
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
			name: "path size with path size",
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
			name: "path size comparison with duplicate names",
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
	for _, c := range cases {
		exprNode, err := c.lhs.Equal(c.rhs).buildCondition()
		if err != nil {
			t.Errorf("TestEquals %#v: Unexpected Error %#v", c.name, err)
		}
		expr, err := exprNode.buildExprNodes(&aliasList{})
		if err != nil {
			t.Errorf("TestEquals %#v: Unexpected Error %#v", c.name, err)
		}

		if reflect.DeepEqual(expr, c.expected) != true {
			t.Errorf("TestEquals %#v: Expected %#v, got %#v", c.name, c.expected, expr)
		}

	}
}

func TestBuildCondition(t *testing.T) {
	cases := []struct {
		name                  string
		input                 ConditionBuilder
		expected              ExprNode
		buildListOperandError bool
		noMatchError          bool
		operandNumberError    bool
		conditionNumberError  bool
	}{
		{
			name:         "no match error",
			input:        ConditionBuilder{},
			noMatchError: true,
		},
		{
			name: "operand number error",
			input: ConditionBuilder{
				Mode: EqualCond,
			},
			operandNumberError: true,
		},
		{
			name: "condition number error",
			input: ConditionBuilder{
				Mode: EqualCond,
				conditionList: []ConditionBuilder{
					ConditionBuilder{},
				},
			},
			conditionNumberError: true,
		},
	}

	for _, c := range cases {
		expr, err := c.input.buildCondition()

		if c.buildListOperandError {
			if err == nil {
				t.Errorf("TestBuildCondition %#v: Expected list operand error but got no error", c.name)
			} else {
				continue
			}
		}

		if c.noMatchError {
			if err == nil {
				t.Errorf("TestBuildCondition %#v: Expected no matching mode error but got no error", c.name)
			} else {
				continue
			}
		}

		if c.operandNumberError {
			if err == nil {
				t.Errorf("TestBuildCondition %#v: Expected operand number error but got no error", c.name)
			} else {
				continue
			}
		}

		if c.conditionNumberError {
			if err == nil {
				t.Errorf("TestBuildCondition %#v: Expected condition number error but got no error", c.name)
			} else {
				continue
			}
		}

		if err != nil {
			t.Errorf("TestBuildCondition %#v: Unexpected Error %#v", c.name, err)
		}

		if reflect.DeepEqual(expr, c.expected) != true {
			t.Errorf("TestBuildCondition %#v: Expected %#v, got %#v", c.name, c.expected, expr)
		}
	}
}

func TestBoolCondition(t *testing.T) {
	cases := []struct {
		name     string
		input    ConditionBuilder
		expected Expression
		err      bool
	}{
		{
			name:  "basic method and",
			input: NewPath("foo").Equal(NewValue(5)).And(NewPath("bar").Equal(NewValue("baz"))),
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("bar"),
					"#1": aws.String("foo"),
				},
				Values: map[string]*dynamodb.AttributeValue{
					":1": &dynamodb.AttributeValue{
						N: aws.String("5"),
					},
					":0": &dynamodb.AttributeValue{
						S: aws.String("baz"),
					},
				},
				Expression: "(#0 = :0) AND (#1 = :1)",
			},
		},
		{
			name:  "variadic function and",
			input: And(NewPath("foo").Equal(NewValue(5)), NewPath("bar").Equal(NewValue("baz")), NewPath("qux").Equal(NewValue(true))),
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
					"#1": aws.String("bar"),
					"#2": aws.String("qux"),
				},
				Values: map[string]*dynamodb.AttributeValue{
					":0": &dynamodb.AttributeValue{
						N: aws.String("5"),
					},
					":1": &dynamodb.AttributeValue{
						S: aws.String("baz"),
					},
					":2": &dynamodb.AttributeValue{
						BOOL: aws.Bool(true),
					},
				},
				Expression: "(#0 = :0) AND (#1 = :1) AND (#2 = :2)",
			},
		},
		{
			name:  "duplicate paths and",
			input: And(NewPath("foo").Equal(NewPath("foo")), NewPath("bar").Equal(NewPath("foo")), NewPath("qux").Equal(NewPath("foo"))),
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
					"#1": aws.String("bar"),
					"#2": aws.String("qux"),
				},
				Expression: "(#0 = #0) AND (#1 = #0) AND (#2 = #0)",
			},
		},
		{
			name: "empty condition",
			input: ConditionBuilder{
				Mode: AndCond,
			},
			err: true,
		},
		{
			name: "non-empty operand list",
			input: ConditionBuilder{
				conditionList: []ConditionBuilder{
					NewValue("foo").Equal(NewPath("bar")),
				},
				operandList: []OperandBuilder{
					NewPath("foo"),
				},
				Mode: AndCond,
			},
			err: true,
		},
	}

	for _, c := range cases {
		expr, err := c.input.BuildExpression()
		if c.err {
			if err == nil {
				t.Errorf("TestBuildCondition %#v: Unexpected Error", c.name)
			} else {
				continue
			}
		}
		if err != nil {
			t.Errorf("TestBuildCondition %#v: Unexpected Error %#v", c.name, err)
		}

		if reflect.DeepEqual(expr, c.expected) != true {
			t.Errorf("TestBuildCondition %#v: Expected %#v, got %#v", c.name, c.expected, expr)
		}
	}
}

func TestBetweenCondition(t *testing.T) {
	cases := []struct {
		name     string
		input    ConditionBuilder
		expected Expression
		err      bool
	}{
		{
			name:  "basic method between for path",
			input: NewPath("foo").Between(NewValue(5), NewValue(7)),
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
				},
				Values: map[string]*dynamodb.AttributeValue{
					":0": &dynamodb.AttributeValue{
						N: aws.String("5"),
					},
					":1": &dynamodb.AttributeValue{
						N: aws.String("7"),
					},
				},
				Expression: "#0 BETWEEN :0 AND :1",
			},
		},
		{
			name:  "basic method between for value",
			input: NewValue(6).Between(NewValue(5), NewValue(7)),
			expected: Expression{
				Values: map[string]*dynamodb.AttributeValue{
					":0": &dynamodb.AttributeValue{
						N: aws.String("6"),
					},
					":1": &dynamodb.AttributeValue{
						N: aws.String("5"),
					},
					":2": &dynamodb.AttributeValue{
						N: aws.String("7"),
					},
				},
				Expression: ":0 BETWEEN :1 AND :2",
			},
		},
		{
			name:  "basic method between for size",
			input: NewPath("foo").Size().Between(NewValue(5), NewValue(7)),
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
				},
				Values: map[string]*dynamodb.AttributeValue{
					":0": &dynamodb.AttributeValue{
						N: aws.String("5"),
					},
					":1": &dynamodb.AttributeValue{
						N: aws.String("7"),
					},
				},
				Expression: "size (#0) BETWEEN :0 AND :1",
			},
		},
		{
			name: "non-empty condition list",
			input: ConditionBuilder{
				conditionList: []ConditionBuilder{
					NewValue("foo").Equal(NewPath("bar")),
				},
				Mode: BetweenCond,
			},
			err: true,
		},
		{
			name: "invalid operand list",
			input: ConditionBuilder{
				operandList: []OperandBuilder{
					NewPath("foo"),
				},
				Mode: BetweenCond,
			},
			err: true,
		},
	}

	for _, c := range cases {
		expr, err := c.input.BuildExpression()
		if c.err {
			if err == nil {
				t.Errorf("TestBuildCondition %#v: Unexpected Error", c.name)
			} else {
				continue
			}
		}
		if err != nil {
			t.Errorf("TestBuildCondition %#v: Unexpected Error %#v", c.name, err)
		}

		if reflect.DeepEqual(expr, c.expected) != true {
			t.Errorf("TestBuildCondition %#v: Expected %#v, got %#v", c.name, c.expected, expr)
		}
	}
}

// If there is time implement mapEquals
