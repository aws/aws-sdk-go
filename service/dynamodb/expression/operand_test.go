// +build go1.7

package expression

import (
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// opeErrorMode will help with error cases and checking error types
type opeErrorMode string

const (
	noOperandError opeErrorMode = ""
	// emptyName error will occur if an empty string is passed into NameBuilder or
	// a nested name has an empty intermediary attribute name (i.e. foo.bar..baz)
	emptyName = "name is an empty string"
	// invalidNameIndex error will occur if there is an invalid index between the
	// square brackets or there is no attribute that a square bracket iterates
	// over
	invalidNameIndex = "invalid name index"
	// unsetExpr error will occur if an unset Expression is passed into
	// mergeExpressionMaps
	unsetExpr = "expression is unset"
)

func TestBuildOperand(t *testing.T) {
	cases := []struct {
		name     string
		input    OperandBuilder
		expected ExprNode
		err      opeErrorMode
	}{
		{
			name:  "basic name",
			input: Name("foo"),
			expected: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "$n",
			},
		},
		{
			name:  "duplicate name name",
			input: Name("foo.foo"),
			expected: ExprNode{
				names:   []string{"foo", "foo"},
				fmtExpr: "$n.$n",
			},
		},
		{
			name:  "basic value",
			input: Value(5),
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
			name:  "nested name",
			input: Name("foo.bar"),
			expected: ExprNode{
				names:   []string{"foo", "bar"},
				fmtExpr: "$n.$n",
			},
		},
		{
			name:  "nested name with index",
			input: Name("foo.bar[0].baz"),
			expected: ExprNode{
				names:   []string{"foo", "bar", "baz"},
				fmtExpr: "$n.$n[0].$n",
			},
		},
		{
			name:  "basic size",
			input: Name("foo").Size(),
			expected: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "size ($n)",
			},
		},
		{
			name:     "empty name error",
			input:    Name(""),
			expected: ExprNode{},
			err:      emptyName,
		},
		{
			name:     "invalid name",
			input:    Name("foo..bar"),
			expected: ExprNode{},
			err:      emptyName,
		},
		{
			name:     "invalid index",
			input:    Name("[foo]"),
			expected: ExprNode{},
			err:      invalidNameIndex,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			en, err := c.input.Build()

			if c.err != noOperandError {
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

				if e, a := c.expected, en; !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}
