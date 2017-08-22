// +build go1.7

package expression

import (
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

// projErrorMode will help with error cases and checking error types
type projErrorMode string

const (
	noProjError projErrorMode = ""
	// invalidProjectionOperand error will occur when an invalid OperandBuilder is used as
	// an argument
	invalidProjectionOperand = "BuildOperand error"
	// emptyNameList error will occur if the names member of ProjectionBuilder is
	// empty
	emptyNameList = "name list is empty"
)

func TestProjection(t *testing.T) {
	cases := []struct {
		name               string
		input              ProjectionBuilder
		expectedNames      map[string]*string
		expectedExpression string
		err                projErrorMode
	}{
		{
			name:  "basic projection",
			input: NamesList(Name("foo"), Name("bar")),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("bar"),
			},
			expectedExpression: "#0, #1",
		},
		{
			name:  "basic projection",
			input: Name("foo").NamesList(Name("bar")),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("bar"),
			},
			expectedExpression: "#0, #1",
		},
		{
			name:  "add name",
			input: Name("foo").NamesList(Name("bar")).AddNames(Name("baz"), Name("qux")),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("bar"),
				"#2": aws.String("baz"),
				"#3": aws.String("qux"),
			},
			expectedExpression: "#0, #1, #2, #3",
		},
		// {
		// 	name:  "invalid operand",
		// 	input: NamesList(Name("")),
		// 	err:   invalidProjectionOperand,
		// },
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := c.input.BuildExpression()
			if c.err != noProjError {
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

				if e, a := aws.String(c.expectedExpression), actual.ProjectionExpression(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedNames, actual.Names(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestBuildProjection(t *testing.T) {
	cases := []struct {
		name     string
		input    ProjectionBuilder
		expected string
		err      projErrorMode
	}{
		{
			name:     "build projection 3",
			input:    NamesList(Name("foo"), Name("bar"), Name("baz")),
			expected: "$c, $c, $c",
		},
		{
			name:     "build projection 5",
			input:    NamesList(Name("foo"), Name("bar"), Name("baz")).AddNames(Name("qux"), Name("quux")),
			expected: "$c, $c, $c, $c, $c",
		},
		{
			name:  "empty ProjectionBuilder",
			input: ProjectionBuilder{},
			err:   emptyNameList,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := c.input.BuildTree()
			if c.err != noProjError {
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

				if e, a := c.expected, actual.fmtExpr; !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestBuildChildNodes(t *testing.T) {
	cases := []struct {
		name     string
		input    ProjectionBuilder
		expected []ExprNode
		err      projErrorMode
	}{
		{
			name:  "build child nodes",
			input: NamesList(Name("foo"), Name("bar"), Name("baz")),
			expected: []ExprNode{
				{
					names:   []string{"foo"},
					fmtExpr: "$p",
				},
				{
					names:   []string{"bar"},
					fmtExpr: "$p",
				},
				{
					names:   []string{"baz"},
					fmtExpr: "$p",
				},
			},
		},
		{
			name:  "operand error",
			input: NamesList(Name("")),
			err:   invalidProjectionOperand,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := c.input.buildChildNodes()
			if c.err != noProjError {
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

				if e, a := c.expected, actual; !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}
