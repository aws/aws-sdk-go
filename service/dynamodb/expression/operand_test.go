package expression

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestBuildOperand(t *testing.T) {
	cases := []struct {
		name           string
		input          OperandBuilder
		expected       ExprNode
		emptyPathError bool
	}{
		{
			name:  "basic path",
			input: NewPath("foo"),
			expected: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "$p",
			},
		},
		{
			name:  "duplicate path name",
			input: NewPath("foo.foo"),
			expected: ExprNode{
				names:   []string{"foo", "foo"},
				fmtExpr: "$p.$p",
			},
		},
		{
			name:  "basic value",
			input: NewValue(5),
			expected: ExprNode{
				values: []dynamodb.AttributeValue{
					dynamodb.AttributeValue{
						N: aws.String("5"),
					},
				},
				fmtExpr: "$v",
			},
		},
		{
			name:  "nested path",
			input: NewPath("foo.bar"),
			expected: ExprNode{
				names:   []string{"foo", "bar"},
				fmtExpr: "$p.$p",
			},
		},
		{
			name:  "nested path with index",
			input: NewPath("foo.bar[0].baz"),
			expected: ExprNode{
				names:   []string{"foo", "bar", "baz"},
				fmtExpr: "$p.$p[0].$p",
			},
		},
		{
			name:  "basic size",
			input: NewPath("foo").Size(),
			expected: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "size ($p)",
			},
		},
		{
			name:           "empty path error",
			input:          NewPath(""),
			expected:       ExprNode{},
			emptyPathError: true,
		},
		{
			name:           "invalid path",
			input:          NewPath("foo..bar"),
			expected:       ExprNode{},
			emptyPathError: true,
		},
		{
			name:           "invalid index",
			input:          NewPath("[foo]"),
			expected:       ExprNode{},
			emptyPathError: true,
		},
	}

	for _, c := range cases {
		en, err := c.input.BuildOperand()

		if c.emptyPathError {
			if err == nil {
				t.Errorf("Test %#v: Expected Error", c.name)
			} else {
				continue
			}
		}

		if err != nil {
			t.Errorf("Test %#v: Unexpected Error %#v", c.name, err)
		}

		if reflect.DeepEqual(c.expected, en) == false {
			t.Errorf("Test %#v: Got %#v, expected %#v\n", c.name, en, c.expected)
		}
	}
}

func TestBuildExpression(t *testing.T) {
	cases := []struct {
		name              string
		input             ExprNode
		expected          Expression
		invalEscError     bool
		outOfRangeError   bool
		nilAliasListError bool
	}{
		{
			name: "basic path",
			input: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "$p",
			},
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
				},
				Expression: "#0",
			},
		},
		{
			name: "basic value",
			input: ExprNode{
				values: []dynamodb.AttributeValue{
					dynamodb.AttributeValue{
						N: aws.String("5"),
					},
				},
				fmtExpr: "$v",
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
			name: "nested path",
			input: ExprNode{
				names:   []string{"foo", "bar"},
				fmtExpr: "$p.$p",
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
			name: "nested path with index",
			input: ExprNode{
				names:   []string{"foo", "bar", "baz"},
				fmtExpr: "$p.$p[0].$p",
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
			name: "basic size",
			input: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "size ($p)",
			},
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
				},
				Expression: "size (#0)",
			},
		},
		{
			name: "duplicate path name",
			input: ExprNode{
				names:   []string{"foo", "foo"},
				fmtExpr: "$p.$p",
			},
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
				},
				Expression: "#0.#0",
			},
		},
		{
			name: "equal expression",
			input: ExprNode{
				children: []ExprNode{
					ExprNode{
						names:   []string{"foo"},
						fmtExpr: "$p",
					},
					ExprNode{
						values: []dynamodb.AttributeValue{
							dynamodb.AttributeValue{
								N: aws.String("5"),
							},
						},
						fmtExpr: "$v",
					},
				},
				fmtExpr: "$c = $c",
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
			name: "missing char after $",
			input: ExprNode{
				names:   []string{"foo", "foo"},
				fmtExpr: "$p.$",
			},
			invalEscError: true,
		},
		{
			name: "names out of range",
			input: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "$p.$p",
			},
			outOfRangeError: true,
		},
		{
			name: "values out of range",
			input: ExprNode{
				fmtExpr: "$v",
			},
			outOfRangeError: true,
		},
		{
			name: "children out of range",
			input: ExprNode{
				fmtExpr: "$c",
			},
			outOfRangeError: true,
		},
		{
			name: "invalid escape char",
			input: ExprNode{
				fmtExpr: "$!",
			},
			outOfRangeError: true,
		},
		{
			name:     "empty ExprNode",
			input:    ExprNode{},
			expected: Expression{},
		},
		{
			name:              "nil aliasList",
			input:             ExprNode{},
			expected:          Expression{},
			nilAliasListError: true,
		},
	}

	for _, c := range cases {
		if c.nilAliasListError {
			_, err := c.input.buildExprNodes(nil)
			if err == nil {
				t.Errorf("Test %#v: Expected Error", c.name)
			} else {
				continue
			}
		}

		expr, err := c.input.buildExprNodes(&aliasList{})
		if c.invalEscError || c.outOfRangeError {
			if err == nil {
				t.Errorf("Test %#v: Expected Error", c.name)
			} else {
				continue
			}
		}
		if err != nil {
			t.Errorf("Test %#v: Unexpected Error %#v", c.name, err)
		}

		if reflect.DeepEqual(expr, c.expected) != true {
			t.Errorf("Test %#v: Expected %#v, got %#v", c.name, c.expected, expr)
		}
	}
}
