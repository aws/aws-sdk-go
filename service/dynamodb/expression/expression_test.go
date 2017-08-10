// +build go1.8

package expression

import (
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type exprErrorMode string

const (
	noExpressionError exprErrorMode = ""
	// invalidEscChar error will occer if the escape char '$' is either followed
	// by an unsupported character or if the escape char is the last character
	invalidEscChar = "invalid escape"
	// outOfRange error will occur if there are more escaped chars than there are
	// actual values to be aliased.
	outOfRange = "out of range"
	// nilAliasList error will occur if the aliasList passed in has not been
	// initialized
	nilAliasList = "aliasList is nil"
)

func TestBuildExpression(t *testing.T) {
	cases := []struct {
		name     string
		input    ExprNode
		expected Expression
		err      exprErrorMode
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
			err: invalidEscChar,
		},
		{
			name: "names out of range",
			input: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "$p.$p",
			},
			err: outOfRange,
		},
		{
			name: "values out of range",
			input: ExprNode{
				fmtExpr: "$v",
			},
			err: outOfRange,
		},
		{
			name: "children out of range",
			input: ExprNode{
				fmtExpr: "$c",
			},
			err: outOfRange,
		},
		{
			name: "invalid escape char",
			input: ExprNode{
				fmtExpr: "$!",
			},
			err: invalidEscChar,
		},
		{
			name:     "empty ExprNode",
			input:    ExprNode{},
			expected: Expression{},
		},
		{
			name:     "nil aliasList",
			input:    ExprNode{},
			expected: Expression{},
			err:      nilAliasList,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var expr Expression
			var err error
			if c.err == nilAliasList {
				expr, err = c.input.buildExprNodes(nil)
			} else {
				expr, err = c.input.buildExprNodes(&aliasList{})
			}

			if c.err != noExpressionError {
				if err == nil {
					t.Errorf("expect error %q, got no error", c.err)
				} else {
					if e, a := string(c.err), err.Error(); !strings.Contains(a, e) {
						t.Errorf("expect %q error message to be in %q", e, a)
					}
				}
			} else {
				if err != nil {
					t.Errorf("expect no error, got unexpected Error %q", err)
				}

				if e, a := c.expected, expr; !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestAliasValue(t *testing.T) {
	cases := []struct {
		name     string
		input    *aliasList
		expected string
		err      exprErrorMode
	}{
		{
			name:  "nil alias list",
			input: nil,
			err:   nilAliasList,
		},
		{
			name:     "first item",
			input:    &aliasList{},
			expected: ":0",
		},
		{
			name: "fifth item",
			input: &aliasList{
				valuesList: []dynamodb.AttributeValue{
					{},
					{},
					{},
					{},
				},
			},
			expected: ":4",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			str, err := c.input.aliasValue(dynamodb.AttributeValue{})

			if c.err != noExpressionError {
				if err == nil {
					t.Errorf("expect error %q, got no error", c.err)
				} else {
					if e, a := string(c.err), err.Error(); !strings.Contains(a, e) {
						t.Errorf("expect %q error message to be in %q", e, a)
					}
				}
			} else {
				if err != nil {
					t.Errorf("expect no error, got unexpected Error %q", err)
				}

				if e, a := c.expected, str; e != a {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestAliasPath(t *testing.T) {
	cases := []struct {
		name      string
		inputList *aliasList
		inputName string
		expected  string
		err       exprErrorMode
	}{
		{
			name:      "nil alias list",
			inputList: nil,
			err:       nilAliasList,
		},
		{
			name:      "new unique item",
			inputList: &aliasList{},
			inputName: "foo",
			expected:  "#0",
		},
		{
			name: "duplicate item",
			inputList: &aliasList{
				namesList: []string{
					"foo",
					"bar",
				},
			},
			inputName: "foo",
			expected:  "#0",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			str, err := c.inputList.aliasPath(c.inputName)

			if c.err != noExpressionError {
				if err == nil {
					t.Errorf("expect error %q, got no error", c.err)
				} else {
					if e, a := string(c.err), err.Error(); !strings.Contains(a, e) {
						t.Errorf("expect %q error message to be in %q", e, a)
					}
				}
			} else {
				if err != nil {
					t.Errorf("expect no error, got unexpected Error %q", err)
				}

				if e, a := c.expected, str; e != a {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestMergeMaps(t *testing.T) {
	cases := []struct {
		name     string
		input    []Expression
		expected Expression
		err      exprErrorMode
	}{
		{
			name: "default use",
			input: []Expression{
				{
					Names: map[string]*string{
						"#0": aws.String("foo"),
						"#1": aws.String("bar"),
						"#2": aws.String("baz"),
					},
					Values: map[string]*dynamodb.AttributeValue{
						":0": {
							S: aws.String("FOO"),
						},
					},
				},
				{
					Names: map[string]*string{
						"#3": aws.String("qux"),
						"#4": aws.String("quux"),
						"#5": aws.String("yar"),
					},
					Values: map[string]*dynamodb.AttributeValue{
						":1": {
							N: aws.String("5"),
						},
					},
				},
			},
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
					"#1": aws.String("bar"),
					"#2": aws.String("baz"),
					"#3": aws.String("qux"),
					"#4": aws.String("quux"),
					"#5": aws.String("yar"),
				},
				Values: map[string]*dynamodb.AttributeValue{
					":0": {
						S: aws.String("FOO"),
					},
					":1": {
						N: aws.String("5"),
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr := MergeMaps(c.input...)

			if e, a := c.expected, expr; !reflect.DeepEqual(e, a) {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}
