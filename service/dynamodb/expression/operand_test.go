package expression

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestBuildOperand(t *testing.T) {
	cases := []struct {
		input          OperandBuilder
		expected       ExprNode
		emptyPathError bool
		// alError        bool
	}{
		{
			input: NewPath("foo"),
			expected: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "%p",
			},
		},
		{
			input: NewPath("foo.foo"),
			expected: ExprNode{
				names:   []string{"foo", "foo"},
				fmtExpr: "%p.%p",
			},
		},
		{
			input: NewValue(5),
			expected: ExprNode{
				values: []dynamodb.AttributeValue{
					dynamodb.AttributeValue{
						N: aws.String("5"),
					},
				},
				fmtExpr: "%v",
			},
		},
		{
			input: NewPath("foo.bar"),
			expected: ExprNode{
				names:   []string{"foo", "bar"},
				fmtExpr: "%p.%p",
			},
		},
		{
			input: NewPath("foo.bar[0].baz"),
			expected: ExprNode{
				names:   []string{"foo", "bar", "baz"},
				fmtExpr: "%p.%p[0].%p",
			},
		},
		{
			input: NewPath("foo").Size(),
			expected: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "size (%p)",
			},
		},
		{
			input:          NewPath(""),
			expected:       ExprNode{},
			emptyPathError: true,
		},
	}

	for testNumber, c := range cases {
		en, err := c.input.BuildOperand()

		if c.emptyPathError {
			if err == nil {
				t.Errorf("TestBuildOperand Test Number %#v: Expected Error", testNumber)
			} else {
				continue
			}
		}

		if err != nil {
			t.Errorf("TestBuildOperand Test Number %#v: Unexpected Error %#v", testNumber, err)
		}

		if reflect.DeepEqual(c.expected, en) == false {
			t.Errorf("TestBuildOperand Test Number %#v: Got %#v, expected %#v\n", testNumber, en, c.expected)
		}
	}
}

func TestBuildExpression(t *testing.T) {
	cases := []struct {
		input    ExprNode
		expected Expression
	}{
		{
			input: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "%p",
			},
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
				},
				Expression: "#0",
			},
		},
		{
			input: ExprNode{
				values: []dynamodb.AttributeValue{
					dynamodb.AttributeValue{
						N: aws.String("5"),
					},
				},
				fmtExpr: "%v",
			},
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
			input: ExprNode{
				names:   []string{"foo", "bar"},
				fmtExpr: "%p.%p",
			},
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
					"#1": aws.String("bar"),
				},
				Expression: "#0.#1",
			},
		},
		{
			input: ExprNode{
				names:   []string{"foo", "bar", "baz"},
				fmtExpr: "%p.%p[0].%p",
			},
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
			input: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "size (%p)",
			},
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
				},
				Expression: "size (#0)",
			},
		},
		{
			input: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "size (%p)",
			},
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
				},
				Expression: "size (#0)",
			},
		},
		{
			input: ExprNode{
				names:   []string{"foo", "foo"},
				fmtExpr: "%p.%p",
			},
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
				},
				Expression: "#0.#0",
			},
		},
		{
			input: ExprNode{
				children: []ExprNode{
					ExprNode{
						names:   []string{"foo"},
						fmtExpr: "%p",
					},
					ExprNode{
						values: []dynamodb.AttributeValue{
							dynamodb.AttributeValue{
								N: aws.String("5"),
							},
						},
						fmtExpr: "%v",
					},
				},
				fmtExpr: "%c = %c",
			},
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
				},
				Values: map[string]*dynamodb.AttributeValue{
					":0": &dynamodb.AttributeValue{
						N: aws.String("5"),
					},
				},
				Expression: "#0 = :0",
			},
		},
		{
			input:    ExprNode{},
			expected: Expression{},
		},
	}

	for testNumber, c := range cases {
		expr, err := c.input.buildExpression(&aliasList{})
		if err != nil {
			t.Errorf("TestBuildExpression Test Number %#v: Unexpected Error %#v", testNumber, err)
		}

		if reflect.DeepEqual(expr, c.expected) != true {
			t.Errorf("TestBuildExpression Test Number %#v: Expected %#v, got %#v", testNumber, c.expected, expr)
		}
	}
}
