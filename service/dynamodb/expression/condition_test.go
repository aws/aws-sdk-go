// +build go1.7

package expression

import (
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// condErrorMode will help with error cases and checking error types
type condErrorMode string

const (
	noConditionError condErrorMode = ""
	// unsetCondition error will occur when BuildExpression is called on an empty
	// ConditionBuilder
	unsetCondition = "UnsetCondition"
	// invalidOperand error will occur when an invalid OperandBuilder is used as
	// an argument
	invalidOperand = "BuildOperand error"
)

//Compare
func TestCompare(t *testing.T) {
	cases := []struct {
		name               string
		input              ConditionBuilder
		expectedExpression string
		expectedNames      map[string]*string
		expectedValues     map[string]*dynamodb.AttributeValue
		err                condErrorMode
	}{
		{
			name:               "nested name with name",
			input:              Name("foo").Equal(Name("bar")),
			expectedExpression: "#0 = #1",
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("bar"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{},
		},
		{
			name:  "nested name with value",
			input: Name("foo.yay.cool.rad").Equal(Value(5)),
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("yay"),
				"#2": aws.String("cool"),
				"#3": aws.String("rad"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: "#0.#1.#2.#3 = :0",
		},
		{
			name:  "nested name with name size",
			input: Name("foo.yay.cool.rad").Equal(Name("baz").Size()),
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("yay"),
				"#2": aws.String("cool"),
				"#3": aws.String("rad"),
				"#4": aws.String("baz"),
			},
			expectedValues:     map[string]*dynamodb.AttributeValue{},
			expectedExpression: "#0.#1.#2.#3 = size (#4)",
		},
		{
			name:  "value with name",
			input: Value(5).Equal(Name("bar")),
			expectedNames: map[string]*string{
				"#0": aws.String("bar"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: ":0 = #0",
		},
		{
			name: "nested value with name",
			input: Value(map[string]int{
				"five": 5,
			}).Equal(Name("bar")),
			expectedNames: map[string]*string{
				"#0": aws.String("bar"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					M: map[string]*dynamodb.AttributeValue{
						"five": &dynamodb.AttributeValue{
							N: aws.String("5"),
						},
					},
				},
			},
			expectedExpression: ":0 = #0",
		},
		{
			name: "nested value with value",
			input: Value(map[string]int{
				"five": 5,
			}).Equal(Value(5)),
			expectedNames: map[string]*string{},
			expectedValues: map[string]*dynamodb.AttributeValue{
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
			expectedExpression: ":0 = :1",
		},
		{
			name: "nested value with name size",
			input: Value(map[string]int{
				"five": 5,
			}).Equal(Name("baz").Size()),
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					M: map[string]*dynamodb.AttributeValue{
						"five": &dynamodb.AttributeValue{
							N: aws.String("5"),
						},
					},
				},
			},
			expectedNames: map[string]*string{
				"#0": aws.String("baz"),
			},
			expectedExpression: ":0 = size (#0)",
		},
		{
			name:  "name size with name",
			input: Name("foo[1]").Size().Equal(Name("bar")),
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("bar"),
			},
			expectedValues:     map[string]*dynamodb.AttributeValue{},
			expectedExpression: "size (#0[1]) = #1",
		},
		{
			name:  "name size with value",
			input: Name("foo[1]").Size().Equal(Value(5)),
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: "size (#0[1]) = :0",
		},
		{
			name:  "name size with name size",
			input: Name("foo[1]").Size().Equal(Name("baz").Size()),
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("baz"),
			},
			expectedValues:     map[string]*dynamodb.AttributeValue{},
			expectedExpression: "size (#0[1]) = size (#1)",
		},
		{
			name:  "name size comparison with duplicate names",
			input: Name("foo.bar.baz").Size().Equal(Name("bar.qux.foo").Size()),
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("bar"),
				"#2": aws.String("baz"),
				"#3": aws.String("qux"),
			},
			expectedValues:     map[string]*dynamodb.AttributeValue{},
			expectedExpression: "size (#0.#1.#2) = size (#1.#3.#0)",
		},
		{
			name:  "name size comparison with duplicate names",
			input: Name("foo.bar.baz").Size().Equal(Name("bar.qux.foo").Size()),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("bar"),
				"#2": aws.String("baz"),
				"#3": aws.String("qux"),
			},
			expectedValues:     map[string]*dynamodb.AttributeValue{},
			expectedExpression: "size (#0.#1.#2) = size (#1.#3.#0)",
		},
		{
			name:  "name NotEqual",
			input: Name("foo").NotEqual(Value(5)),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: "#0 <> :0",
		},
		{
			name:          "value NotEqual",
			input:         Value(8).NotEqual(Value(5)),
			expectedNames: map[string]*string{},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("8"),
				},
				":1": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: ":0 <> :1",
		},
		{
			name:  "name NotEqual",
			input: Name("foo").Size().NotEqual(Value(5)),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: "size (#0) <> :0",
		},
		{
			name:  "name Less",
			input: Name("foo").Less(Value(5)),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: "#0 < :0",
		},
		{
			name:          "value Less",
			input:         Value(8).Less(Value(5)),
			expectedNames: map[string]*string{},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("8"),
				},
				":1": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: ":0 < :1",
		},
		{
			name:  "name Less",
			input: Name("foo").Size().Less(Value(5)),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: "size (#0) < :0",
		},
		{
			name:  "name LessEqual",
			input: Name("foo").LessEqual(Value(5)),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: "#0 <= :0",
		},
		{
			name:  "value LessEqual",
			input: Value(8).LessEqual(Value(5)),

			expectedNames: map[string]*string{},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("8"),
				},
				":1": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: ":0 <= :1",
		},
		{
			name:  "name LessEqual",
			input: Name("foo").Size().LessEqual(Value(5)),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: "size (#0) <= :0",
		},
		{
			name:  "name Greater",
			input: Name("foo").Greater(Value(5)),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: "#0 > :0",
		},
		{
			name:  "value Greater",
			input: Value(8).Greater(Value(5)),

			expectedNames: map[string]*string{},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("8"),
				},
				":1": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: ":0 > :1",
		},
		{
			name:  "name Greater",
			input: Name("foo").Size().Greater(Value(5)),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: "size (#0) > :0",
		},
		{
			name:  "name GreaterEqual",
			input: Name("foo").GreaterEqual(Value(5)),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: "#0 >= :0",
		},
		{
			name:  "value GreaterEqual",
			input: Value(10).GreaterEqual(Value(5)),

			expectedNames: map[string]*string{},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("10"),
				},
				":1": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: ":0 >= :1",
		},
		{
			name:  "name GreaterEqual",
			input: Name("foo").Size().GreaterEqual(Value(5)),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: "size (#0) >= :0",
		},
		// {
		// 	name:  "invalid operand error Equal",
		// 	input: Name("").Size().Equal(Value(5)),
		// 	err:   invalidOperand,
		// },
		// {
		// 	name:  "invalid operand error NotEqual",
		// 	input: Name("").Size().NotEqual(Value(5)),
		// 	err:   invalidOperand,
		// },
		// {
		// 	name:  "invalid operand error Less",
		// 	input: Name("").Size().Less(Value(5)),
		// 	err:   invalidOperand,
		// },
		// {
		// 	name:  "invalid operand error LessEqual",
		// 	input: Name("").Size().LessEqual(Value(5)),
		// 	err:   invalidOperand,
		// },
		// {
		// 	name:  "invalid operand error Greater",
		// 	input: Name("").Size().Greater(Value(5)),
		// 	err:   invalidOperand,
		// },
		// {
		// 	name:  "invalid operand error GreaterEqual",
		// 	input: Name("").Size().GreaterEqual(Value(5)),
		// 	err:   invalidOperand,
		// },
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := c.input.BuildExpression()
			if c.err != noConditionError {
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

				if e, a := aws.String(c.expectedExpression), expr.ConditionExpression(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedNames, expr.Names(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedValues, expr.Values(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
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
			err:   unsetCondition,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := c.input.BuildTree()

			if c.err != noConditionError {
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

func TestBoolCondition(t *testing.T) {
	cases := []struct {
		name               string
		input              ConditionBuilder
		expectedNames      map[string]*string
		expectedValues     map[string]*dynamodb.AttributeValue
		expectedExpression string
		err                condErrorMode
	}{
		{
			name:  "basic method and",
			input: Name("foo").Equal(Value(5)).And(Name("bar").Equal(Value("baz"))),

			expectedNames: map[string]*string{
				"#1": aws.String("bar"),
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
				":1": &dynamodb.AttributeValue{
					S: aws.String("baz"),
				},
			},
			expectedExpression: "(#0 = :0) AND (#1 = :1)",
		},
		{
			name:  "basic method or",
			input: Name("foo").Equal(Value(5)).Or(Name("bar").Equal(Value("baz"))),

			expectedNames: map[string]*string{
				"#1": aws.String("bar"),
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
				":1": &dynamodb.AttributeValue{
					S: aws.String("baz"),
				},
			},
			expectedExpression: "(#0 = :0) OR (#1 = :1)",
		},
		{
			name:  "variadic function and",
			input: And(Name("foo").Equal(Value(5)), Name("bar").Equal(Value("baz")), Name("qux").Equal(Value(true))),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("bar"),
				"#2": aws.String("qux"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
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
			expectedExpression: "(#0 = :0) AND (#1 = :1) AND (#2 = :2)",
		},
		{
			name:  "variadic function or",
			input: Or(Name("foo").Equal(Value(5)), Name("bar").Equal(Value("baz")), Name("qux").Equal(Value(true))),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("bar"),
				"#2": aws.String("qux"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
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
			expectedExpression: "(#0 = :0) OR (#1 = :1) OR (#2 = :2)",
		},
		{
			name:           "duplicate names and",
			input:          And(Name("foo").Equal(Name("foo")), Name("bar").Equal(Name("foo")), Name("qux").Equal(Name("foo"))),
			expectedValues: map[string]*dynamodb.AttributeValue{},
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("bar"),
				"#2": aws.String("qux"),
			},
			expectedExpression: "(#0 = #0) AND (#1 = #0) AND (#2 = #0)",
		},
		// {
		// 	name:  "invalid operand error And",
		// 	input: Name("").Size().GreaterEqual(Value(5)).And(Name("[5]").Between(Value(3), Value(9))),
		// 	err:   invalidOperand,
		// },
		// {
		// 	name:  "invalid operand error Or",
		// 	input: Name("").Size().GreaterEqual(Value(5)).Or(Name("[5]").Between(Value(3), Value(9))),
		// 	err:   invalidOperand,
		// },
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := c.input.BuildExpression()
			if c.err != noConditionError {
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

				if e, a := aws.String(c.expectedExpression), expr.ConditionExpression(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedNames, expr.Names(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedValues, expr.Values(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestNotCondition(t *testing.T) {
	cases := []struct {
		name               string
		input              ConditionBuilder
		expectedNames      map[string]*string
		expectedValues     map[string]*dynamodb.AttributeValue
		expectedExpression string
		err                condErrorMode
	}{
		{
			name:  "basic method not",
			input: Name("foo").Equal(Value(5)).Not(),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: "NOT (#0 = :0)",
		},
		{
			name:  "basic function not",
			input: Not(Name("foo").Equal(Value(5))),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: "NOT (#0 = :0)",
		},
		// {
		// 	name:  "invalid operand error not",
		// 	input: Name("").Size().GreaterEqual(Value(5)).Or(Name("[5]").Between(Value(3), Value(9))).Not(),
		// 	err:   invalidOperand,
		// },
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := c.input.BuildExpression()
			if c.err != noConditionError {
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

				if e, a := aws.String(c.expectedExpression), expr.ConditionExpression(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedNames, expr.Names(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedValues, expr.Values(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestBetweenCondition(t *testing.T) {
	cases := []struct {
		name               string
		input              ConditionBuilder
		expectedNames      map[string]*string
		expectedValues     map[string]*dynamodb.AttributeValue
		expectedExpression string
		err                condErrorMode
	}{
		{
			name:  "basic method between for name",
			input: Name("foo").Between(Value(5), Value(7)),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
				":1": &dynamodb.AttributeValue{
					N: aws.String("7"),
				},
			},
			expectedExpression: "#0 BETWEEN :0 AND :1",
		},
		{
			name:          "basic method between for value",
			input:         Value(6).Between(Value(5), Value(7)),
			expectedNames: map[string]*string{},
			expectedValues: map[string]*dynamodb.AttributeValue{
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
			expectedExpression: ":0 BETWEEN :1 AND :2",
		},
		{
			name:  "basic method between for size",
			input: Name("foo").Size().Between(Value(5), Value(7)),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
				":1": &dynamodb.AttributeValue{
					N: aws.String("7"),
				},
			},
			expectedExpression: "size (#0) BETWEEN :0 AND :1",
		},
		// {
		// 	name:  "invalid operand error between",
		// 	input: Name("[5]").Between(Value(3), Name("foo..bar")),
		// 	err:   invalidOperand,
		// },
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := c.input.BuildExpression()
			if c.err != noConditionError {
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

				if e, a := aws.String(c.expectedExpression), expr.ConditionExpression(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedNames, expr.Names(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedValues, expr.Values(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestInCondition(t *testing.T) {
	cases := []struct {
		name               string
		input              ConditionBuilder
		expectedNames      map[string]*string
		expectedValues     map[string]*dynamodb.AttributeValue
		expectedExpression string
		err                condErrorMode
	}{
		{
			name:  "basic method in for name",
			input: Name("foo").In(Value(5), Value(7)),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
				":1": &dynamodb.AttributeValue{
					N: aws.String("7"),
				},
			},
			expectedExpression: "#0 IN (:0, :1)",
		},
		{
			name:          "basic method in for value",
			input:         Value(6).In(Value(5), Value(7)),
			expectedNames: map[string]*string{},
			expectedValues: map[string]*dynamodb.AttributeValue{
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
			expectedExpression: ":0 IN (:1, :2)",
		},
		{
			name:  "basic method in for size",
			input: Name("foo").Size().In(Value(5), Value(7)),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
				":1": &dynamodb.AttributeValue{
					N: aws.String("7"),
				},
			},
			expectedExpression: "size (#0) IN (:0, :1)",
		},
		// {
		// 	name:  "invalid operand error in",
		// 	input: Name("[5]").In(Value(3), Name("foo..bar")),
		// 	err:   invalidOperand,
		// },
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := c.input.BuildExpression()
			if c.err != noConditionError {
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

				if e, a := aws.String(c.expectedExpression), expr.ConditionExpression(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedNames, expr.Names(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedValues, expr.Values(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestAttrExistsCondition(t *testing.T) {
	cases := []struct {
		name               string
		input              ConditionBuilder
		expectedNames      map[string]*string
		expectedValues     map[string]*dynamodb.AttributeValue
		expectedExpression string
		err                condErrorMode
	}{
		{
			name:           "basic attr exists",
			input:          Name("foo").AttributeExists(),
			expectedValues: map[string]*dynamodb.AttributeValue{},
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedExpression: "attribute_exists (#0)",
		},
		{
			name:           "basic attr not exists",
			input:          Name("foo").AttributeNotExists(),
			expectedValues: map[string]*dynamodb.AttributeValue{},
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedExpression: "attribute_not_exists (#0)",
		},
		// {
		// 	name:  "invalid operand error attr exists",
		// 	input: AttributeExists(Name("")),
		// 	err:   invalidOperand,
		// },
		// {
		// 	name:  "invalid operand error attr not exists",
		// 	input: AttributeNotExists(Name("foo..bar")),
		// 	err:   invalidOperand,
		// },
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := c.input.BuildExpression()
			if c.err != noConditionError {
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

				if e, a := aws.String(c.expectedExpression), expr.ConditionExpression(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedNames, expr.Names(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedValues, expr.Values(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestAttrTypeCondition(t *testing.T) {
	cases := []struct {
		name               string
		input              ConditionBuilder
		expectedNames      map[string]*string
		expectedValues     map[string]*dynamodb.AttributeValue
		expectedExpression string
		err                condErrorMode
	}{
		{
			name:  "attr type String",
			input: Name("foo").AttributeType(String),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": {
					S: aws.String("S"),
				},
			},
			expectedExpression: "attribute_type (#0, :0)",
		},
		{
			name:  "attr type StringSet",
			input: Name("foo").AttributeType(StringSet),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": {
					S: aws.String("SS"),
				},
			},
			expectedExpression: "attribute_type (#0, :0)",
		},
		{
			name:  "attr type Number",
			input: Name("foo").AttributeType(Number),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": {
					S: aws.String("N"),
				},
			},
			expectedExpression: "attribute_type (#0, :0)",
		},
		{
			name:  "attr type Number Set",
			input: Name("foo").AttributeType(NumberSet),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": {
					S: aws.String("NS"),
				},
			},
			expectedExpression: "attribute_type (#0, :0)",
		},
		{
			name:  "attr type Binary",
			input: Name("foo").AttributeType(Binary),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": {
					S: aws.String("B"),
				},
			},
			expectedExpression: "attribute_type (#0, :0)",
		},
		{
			name:  "attr type Binary Set",
			input: Name("foo").AttributeType(BinarySet),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": {
					S: aws.String("BS"),
				},
			},
			expectedExpression: "attribute_type (#0, :0)",
		},
		{
			name:  "attr type Boolean",
			input: Name("foo").AttributeType(Boolean),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": {
					S: aws.String("BOOL"),
				},
			},
			expectedExpression: "attribute_type (#0, :0)",
		},
		{
			name:  "attr type Null",
			input: Name("foo").AttributeType(Null),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": {
					S: aws.String("NULL"),
				},
			},
			expectedExpression: "attribute_type (#0, :0)",
		},
		{
			name:  "attr type List",
			input: Name("foo").AttributeType(List),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": {
					S: aws.String("L"),
				},
			},
			expectedExpression: "attribute_type (#0, :0)",
		},
		{
			name:  "attr type Map",
			input: Name("foo").AttributeType(Map),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": {
					S: aws.String("M"),
				},
			},
			expectedExpression: "attribute_type (#0, :0)",
		},
		// {
		// 	name:  "attr type invalid operand",
		// 	input: Name("").AttributeType(Map),
		// 	err:   invalidOperand,
		// },
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := c.input.BuildExpression()
			if c.err != noConditionError {
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

				if e, a := aws.String(c.expectedExpression), expr.ConditionExpression(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedNames, expr.Names(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedValues, expr.Values(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestBeginsWithCondition(t *testing.T) {
	cases := []struct {
		name               string
		input              ConditionBuilder
		expectedNames      map[string]*string
		expectedValues     map[string]*dynamodb.AttributeValue
		expectedExpression string
		err                condErrorMode
	}{
		{
			name:  "basic begins with",
			input: Name("foo").BeginsWith("bar"),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": {
					S: aws.String("bar"),
				},
			},
			expectedExpression: "begins_with (#0, :0)",
		},
		// {
		// 	name:  "begins with invalid operand",
		// 	input: Name("").BeginsWith("bar"),
		// 	err:   invalidOperand,
		// },
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := c.input.BuildExpression()
			if c.err != noConditionError {
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

				if e, a := aws.String(c.expectedExpression), expr.ConditionExpression(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedNames, expr.Names(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedValues, expr.Values(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestContainsCondition(t *testing.T) {
	cases := []struct {
		name               string
		input              ConditionBuilder
		expectedNames      map[string]*string
		expectedValues     map[string]*dynamodb.AttributeValue
		expectedExpression string
		err                condErrorMode
	}{
		{
			name:  "basic contains",
			input: Name("foo").Contains("bar"),

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": {
					S: aws.String("bar"),
				},
			},
			expectedExpression: "contains (#0, :0)",
		},
		// {
		// 	name:  "contains invalid operand",
		// 	input: Name("").Contains("bar"),
		// 	err:   invalidOperand,
		// },
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := c.input.BuildExpression()
			if c.err != noConditionError {
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

				if e, a := aws.String(c.expectedExpression), expr.ConditionExpression(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedNames, expr.Names(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
				if e, a := c.expectedValues, expr.Values(); !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestCompoundBuildCondition(t *testing.T) {
	cases := []struct {
		name      string
		inputCond ConditionBuilder
		expected  string
	}{
		{
			name: "and",
			inputCond: ConditionBuilder{
				conditionList: []ConditionBuilder{
					ConditionBuilder{},
					ConditionBuilder{},
					ConditionBuilder{},
					ConditionBuilder{},
				},
				mode: andCond,
			},
			expected: "($c) AND ($c) AND ($c) AND ($c)",
		},
		{
			name: "or",
			inputCond: ConditionBuilder{
				conditionList: []ConditionBuilder{
					ConditionBuilder{},
					ConditionBuilder{},
					ConditionBuilder{},
					ConditionBuilder{},
					ConditionBuilder{},
					ConditionBuilder{},
					ConditionBuilder{},
				},
				mode: orCond,
			},
			expected: "($c) OR ($c) OR ($c) OR ($c) OR ($c) OR ($c) OR ($c)",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			en, err := compoundBuildCondition(c.inputCond, ExprNode{})
			if err != nil {
				t.Errorf("expect no error, got unexpected Error %q", err)
			}

			if e, a := c.expected, en.fmtExpr; !reflect.DeepEqual(a, e) {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

func TestInBuildCondition(t *testing.T) {
	cases := []struct {
		name      string
		inputCond ConditionBuilder
		expected  string
	}{
		{
			name: "in",
			inputCond: ConditionBuilder{
				operandList: []OperandBuilder{
					NameBuilder{},
					NameBuilder{},
					NameBuilder{},
					NameBuilder{},
					NameBuilder{},
					NameBuilder{},
					NameBuilder{},
				},
				mode: andCond,
			},
			expected: "$c IN ($c, $c, $c, $c, $c, $c)",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			en, err := inBuildCondition(c.inputCond, ExprNode{})
			if err != nil {
				t.Errorf("expect no error, got unexpected Error %q", err)
			}

			if e, a := c.expected, en.fmtExpr; !reflect.DeepEqual(a, e) {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

// If there is time implement mapEquals
