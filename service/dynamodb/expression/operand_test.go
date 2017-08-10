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
	// emptyPath error will occur if an empty string is passed into PathBuilder or
	// a nested path has an empty intermediary attribute name (i.e. foo.bar..baz)
	emptyPath = "path is an empty string"
	// invalidPathIndex error will occur if there is an invalid index between the
	// square brackets or there is no attribute that a square bracket iterates
	// over
	invalidPathIndex = "invalid path index"
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
			name:  "basic path",
			input: Path("foo"),
			expected: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "$p",
			},
		},
		{
			name:  "duplicate path name",
			input: Path("foo.foo"),
			expected: ExprNode{
				names:   []string{"foo", "foo"},
				fmtExpr: "$p.$p",
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
			name:  "nested path",
			input: Path("foo.bar"),
			expected: ExprNode{
				names:   []string{"foo", "bar"},
				fmtExpr: "$p.$p",
			},
		},
		{
			name:  "nested path with index",
			input: Path("foo.bar[0].baz"),
			expected: ExprNode{
				names:   []string{"foo", "bar", "baz"},
				fmtExpr: "$p.$p[0].$p",
			},
		},
		{
			name:  "basic size",
			input: Path("foo").Size(),
			expected: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "size ($p)",
			},
		},
		{
			name:     "empty path error",
			input:    Path(""),
			expected: ExprNode{},
			err:      emptyPath,
		},
		{
			name:     "invalid path",
			input:    Path("foo..bar"),
			expected: ExprNode{},
			err:      emptyPath,
		},
		{
			name:     "invalid index",
			input:    Path("[foo]"),
			expected: ExprNode{},
			err:      invalidPathIndex,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			en, err := c.input.BuildOperand()

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
