// +build go1.8

package expression

import (
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// condErrorMode will help with error cases and checking error types
type condErrorMode int

const (
	noConditionError condErrorMode = iota
	// noMatchingMode error will occur when the ConditionBuilder's Mode is not
	// supported
	noMatchingMode
)

func (cem condErrorMode) String() string {
	switch cem {
	case noConditionError:
		return "no Error"
	case noMatchingMode:
		return "no matching"
	default:
		return ""
	}
}

//Compare
func TestCompare(t *testing.T) {
	cases := []struct {
		name     string
		input    ConditionBuilder
		expected Expression
	}{
		{
			name:  "nested path with path",
			input: Path("foo.yay.cool.rad").Equal(Path("bar")),
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
			name:  "nested path with value",
			input: Path("foo.yay.cool.rad").Equal(Value(5)),
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
			name:  "nested path with path size",
			input: Path("foo.yay.cool.rad").Equal(Path("baz").Size()),
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
			name:  "value with path",
			input: Value(5).Equal(Path("bar")),
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
			input: Value(map[string]int{
				"five": 5,
			}).Equal(Path("bar")),
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
			input: Value(map[string]int{
				"five": 5,
			}).Equal(Value(5)),
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
			input: Value(map[string]int{
				"five": 5,
			}).Equal(Path("baz").Size()),
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
			name:  "path size with path",
			input: Path("foo[1]").Size().Equal(Path("bar")),
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
					"#1": aws.String("bar"),
				},
				Expression: "size (#0[1]) = #1",
			},
		},
		{
			name:  "path size with value",
			input: Path("foo[1]").Size().Equal(Value(5)),
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
			name:  "path size with path size",
			input: Path("foo[1]").Size().Equal(Path("baz").Size()),
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
					"#1": aws.String("baz"),
				},
				Expression: "size (#0[1]) = size (#1)",
			},
		},
		{
			name:  "path size comparison with duplicate names",
			input: Path("foo.bar.baz").Size().Equal(Path("bar.qux.foo").Size()),
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
		{
			name:  "path size comparison with duplicate names",
			input: Path("foo.bar.baz").Size().Equal(Path("bar.qux.foo").Size()),
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
		t.Run(c.name, func(t *testing.T) {
			expr, err := c.input.BuildExpression()
			if err != nil {
				t.Errorf("expect no error, got error %v", err)
			}

			if e, a := c.expected, expr; !reflect.DeepEqual(e, a) {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

func TestBuildCondition(t *testing.T) {
	cases := []struct {
		name     string
		input    ConditionBuilder
		expected ExprNode
		err      condErrorMode
	}{
		{
			name:  "no match error",
			input: ConditionBuilder{},
			err:   noMatchingMode,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := c.input.buildCondition()

			if c.err != noConditionError {
				if err == nil {
					t.Errorf("expect error %q, got no error", c.err)
				} else {
					if e, a := c.err.String(), err.Error(); !strings.Contains(a, e) {
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

func TestBoolCondition(t *testing.T) {
	cases := []struct {
		name     string
		input    ConditionBuilder
		expected Expression
		err      condErrorMode
	}{
		{
			name:  "basic method and",
			input: Path("foo").Equal(Value(5)).And(Path("bar").Equal(Value("baz"))),
			expected: Expression{
				Names: map[string]*string{
					"#1": aws.String("bar"),
					"#0": aws.String("foo"),
				},
				Values: map[string]*dynamodb.AttributeValue{
					":0": &dynamodb.AttributeValue{
						N: aws.String("5"),
					},
					":1": &dynamodb.AttributeValue{
						S: aws.String("baz"),
					},
				},
				Expression: "(#0 = :0) AND (#1 = :1)",
			},
		},
		{
			name:  "variadic function and",
			input: And(Path("foo").Equal(Value(5)), Path("bar").Equal(Value("baz")), Path("qux").Equal(Value(true))),
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
			input: And(Path("foo").Equal(Path("foo")), Path("bar").Equal(Path("foo")), Path("qux").Equal(Path("foo"))),
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
					"#1": aws.String("bar"),
					"#2": aws.String("qux"),
				},
				Expression: "(#0 = #0) AND (#1 = #0) AND (#2 = #0)",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := c.input.BuildExpression()
			if c.err != noConditionError {
				if err == nil {
					t.Errorf("expect error %q, got no error", c.err)
				} else {
					if e, a := c.err.String(), err.Error(); !strings.Contains(a, e) {
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

func TestBetweenCondition(t *testing.T) {
	cases := []struct {
		name     string
		input    ConditionBuilder
		expected Expression
		err      condErrorMode
	}{
		{
			name:  "basic method between for path",
			input: Path("foo").Between(Value(5), Value(7)),
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
			input: Value(6).Between(Value(5), Value(7)),
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
			input: Path("foo").Size().Between(Value(5), Value(7)),
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
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := c.input.BuildExpression()
			if c.err != noConditionError {
				if err == nil {
					t.Errorf("expect error %q, got no error", c.err)
				} else {
					if e, a := c.err.String(), err.Error(); !strings.Contains(a, e) {
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

// If there is time implement mapEquals
