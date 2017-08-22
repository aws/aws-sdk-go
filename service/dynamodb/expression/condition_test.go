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
			name:               "nested path with path",
			input:              Path("foo").Equal(Path("bar")),
			expectedExpression: "#0 = #1",
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("bar"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{},
		},
		{
			name:  "nested path with value",
			input: Path("foo.yay.cool.rad").Equal(Value(5)),
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
			name:  "nested path with path size",
			input: Path("foo.yay.cool.rad").Equal(Path("baz").Size()),
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
			name:  "value with path",
			input: Value(5).Equal(Path("bar")),
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
			name: "nested value with path",
			input: Value(map[string]int{
				"five": 5,
			}).Equal(Path("bar")),
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
			name: "nested value with path size",
			input: Value(map[string]int{
				"five": 5,
			}).Equal(Path("baz").Size()),
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
			name:  "path size with path",
			input: Path("foo[1]").Size().Equal(Path("bar")),
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("bar"),
			},
			expectedValues:     map[string]*dynamodb.AttributeValue{},
			expectedExpression: "size (#0[1]) = #1",
		},
		{
			name:  "path size with value",
			input: Path("foo[1]").Size().Equal(Value(5)),
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
			name:  "path size with path size",
			input: Path("foo[1]").Size().Equal(Path("baz").Size()),
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("baz"),
			},
			expectedValues:     map[string]*dynamodb.AttributeValue{},
			expectedExpression: "size (#0[1]) = size (#1)",
		},
		{
			name:  "path size comparison with duplicate names",
			input: Path("foo.bar.baz").Size().Equal(Path("bar.qux.foo").Size()),
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
			name:  "path size comparison with duplicate names",
			input: Path("foo.bar.baz").Size().Equal(Path("bar.qux.foo").Size()),

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
			name:  "path NotEqual",
			input: Path("foo").NotEqual(Value(5)),

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
			name:  "path NotEqual",
			input: Path("foo").Size().NotEqual(Value(5)),

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
			name:  "path Less",
			input: Path("foo").Less(Value(5)),

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
			name:  "path Less",
			input: Path("foo").Size().Less(Value(5)),

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
			name:  "path LessEqual",
			input: Path("foo").LessEqual(Value(5)),

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
			name:  "path LessEqual",
			input: Path("foo").Size().LessEqual(Value(5)),

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
			name:  "path Greater",
			input: Path("foo").Greater(Value(5)),

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
			name:  "path Greater",
			input: Path("foo").Size().Greater(Value(5)),

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
			name:  "path GreaterEqual",
			input: Path("foo").GreaterEqual(Value(5)),

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
			name:  "path GreaterEqual",
			input: Path("foo").Size().GreaterEqual(Value(5)),

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
		// 	input: Path("").Size().Equal(Value(5)),
		// 	err:   invalidOperand,
		// },
		// {
		// 	name:  "invalid operand error NotEqual",
		// 	input: Path("").Size().NotEqual(Value(5)),
		// 	err:   invalidOperand,
		// },
		// {
		// 	name:  "invalid operand error Less",
		// 	input: Path("").Size().Less(Value(5)),
		// 	err:   invalidOperand,
		// },
		// {
		// 	name:  "invalid operand error LessEqual",
		// 	input: Path("").Size().LessEqual(Value(5)),
		// 	err:   invalidOperand,
		// },
		// {
		// 	name:  "invalid operand error Greater",
		// 	input: Path("").Size().Greater(Value(5)),
		// 	err:   invalidOperand,
		// },
		// {
		// 	name:  "invalid operand error GreaterEqual",
		// 	input: Path("").Size().GreaterEqual(Value(5)),
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

// func TestBuildCondition(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		input    ConditionBuilder
// 		expected ExprNode
// 		err      condErrorMode
// 	}{
// 		{
// 			name:  "no match error",
// 			input: ConditionBuilder{},
// 			err:   unsetCondition,
// 		},
// 	}
//
// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			expr, err := c.input.buildCondition()
//
// 			if c.err != noConditionError {
// 				if err == nil {
// 					t.Errorf("expect error %q, got no error", c.err)
// 				} else {
// 					if e, a := string(c.err), err.Error(); !strings.Contains(a, e) {
// 						t.Errorf("expect %q error message to be in %q", e, a)
// 					}
// 				}
// 			} else {
// 				if err != nil {
// 					t.Errorf("expect no error, got unexpected Error %q", err)
// 				}
//
// 				if e, a := c.expected, expr; !reflect.DeepEqual(a, e) {
// 					t.Errorf("expect %v, got %v", e, a)
// 				}
// 			}
// 		})
// 	}
// }
//
// func TestBoolCondition(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		input    ConditionBuilder
// 		expected expectedExpression
// 		err      condErrorMode
// 	}{
// 		{
// 			name:  "basic method and",
// 			input: Path("foo").Equal(Value(5)).And(Path("bar").Equal(Value("baz"))),
//
// 				expectedNames: map[string]*string{
// 					"#1": aws.String("bar"),
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": &dynamodb.AttributeValue{
// 						N: aws.String("5"),
// 					},
// 					":1": &dynamodb.AttributeValue{
// 						S: aws.String("baz"),
// 					},
// 				},
// 				expectedExpression: "(#0 = :0) AND (#1 = :1)",
// 			},
// 		},
// 		{
// 			name:  "basic method or",
// 			input: Path("foo").Equal(Value(5)).Or(Path("bar").Equal(Value("baz"))),
//
// 				expectedNames: map[string]*string{
// 					"#1": aws.String("bar"),
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": &dynamodb.AttributeValue{
// 						N: aws.String("5"),
// 					},
// 					":1": &dynamodb.AttributeValue{
// 						S: aws.String("baz"),
// 					},
// 				},
// 				expectedExpression: "(#0 = :0) OR (#1 = :1)",
// 			},
// 		},
// 		{
// 			name:  "variadic function and",
// 			input: And(Path("foo").Equal(Value(5)), Path("bar").Equal(Value("baz")), Path("qux").Equal(Value(true))),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 					"#1": aws.String("bar"),
// 					"#2": aws.String("qux"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": &dynamodb.AttributeValue{
// 						N: aws.String("5"),
// 					},
// 					":1": &dynamodb.AttributeValue{
// 						S: aws.String("baz"),
// 					},
// 					":2": &dynamodb.AttributeValue{
// 						BOOL: aws.Bool(true),
// 					},
// 				},
// 				expectedExpression: "(#0 = :0) AND (#1 = :1) AND (#2 = :2)",
// 			},
// 		},
// 		{
// 			name:  "variadic function or",
// 			input: Or(Path("foo").Equal(Value(5)), Path("bar").Equal(Value("baz")), Path("qux").Equal(Value(true))),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 					"#1": aws.String("bar"),
// 					"#2": aws.String("qux"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": &dynamodb.AttributeValue{
// 						N: aws.String("5"),
// 					},
// 					":1": &dynamodb.AttributeValue{
// 						S: aws.String("baz"),
// 					},
// 					":2": &dynamodb.AttributeValue{
// 						BOOL: aws.Bool(true),
// 					},
// 				},
// 				expectedExpression: "(#0 = :0) OR (#1 = :1) OR (#2 = :2)",
// 			},
// 		},
// 		{
// 			name:  "duplicate paths and",
// 			input: And(Path("foo").Equal(Path("foo")), Path("bar").Equal(Path("foo")), Path("qux").Equal(Path("foo"))),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 					"#1": aws.String("bar"),
// 					"#2": aws.String("qux"),
// 				},
// 				expectedExpression: "(#0 = #0) AND (#1 = #0) AND (#2 = #0)",
// 			},
// 		},
// 		{
// 			name:  "invalid operand error And",
// 			input: Path("").Size().GreaterEqual(Value(5)).And(Path("[5]").Between(Value(3), Value(9))),
// 			err:   invalidOperand,
// 		},
// 		{
// 			name:  "invalid operand error Or",
// 			input: Path("").Size().GreaterEqual(Value(5)).Or(Path("[5]").Between(Value(3), Value(9))),
// 			err:   invalidOperand,
// 		},
// 	}
//
// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			expr, err := c.input.BuildExpression()
// 			if c.err != noConditionError {
// 				if err == nil {
// 					t.Errorf("expect error %q, got no error", c.err)
// 				} else {
// 					if e, a := string(c.err), err.Error(); !strings.Contains(a, e) {
// 						t.Errorf("expect %q error message to be in %q", e, a)
// 					}
// 				}
// 			} else {
// 				if err != nil {
// 					t.Errorf("expect no error, got unexpected Error %q", err)
// 				}
//
// 				if e, a := c.expected, expr; !reflect.DeepEqual(a, e) {
// 					t.Errorf("expect %v, got %v", e, a)
// 				}
// 			}
// 		})
// 	}
// }
//
// func TestNotCondition(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		input    ConditionBuilder
// 		expected expectedExpression
// 		err      condErrorMode
// 	}{
// 		{
// 			name:  "basic method not",
// 			input: Path("foo").Equal(Value(5)).Not(),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": &dynamodb.AttributeValue{
// 						N: aws.String("5"),
// 					},
// 				},
// 				expectedExpression: "NOT (#0 = :0)",
// 			},
// 		},
// 		{
// 			name:  "basic function not",
// 			input: Not(Path("foo").Equal(Value(5))),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": &dynamodb.AttributeValue{
// 						N: aws.String("5"),
// 					},
// 				},
// 				expectedExpression: "NOT (#0 = :0)",
// 			},
// 		},
// 		{
// 			name:  "invalid operand error not",
// 			input: Path("").Size().GreaterEqual(Value(5)).Or(Path("[5]").Between(Value(3), Value(9))).Not(),
// 			err:   invalidOperand,
// 		},
// 	}
//
// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			expr, err := c.input.BuildExpression()
// 			if c.err != noConditionError {
// 				if err == nil {
// 					t.Errorf("expect error %q, got no error", c.err)
// 				} else {
// 					if e, a := string(c.err), err.Error(); !strings.Contains(a, e) {
// 						t.Errorf("expect %q error message to be in %q", e, a)
// 					}
// 				}
// 			} else {
// 				if err != nil {
// 					t.Errorf("expect no error, got unexpected Error %q", err)
// 				}
//
// 				if e, a := c.expected, expr; !reflect.DeepEqual(a, e) {
// 					t.Errorf("expect %v, got %v", e, a)
// 				}
// 			}
// 		})
// 	}
// }
//
// func TestBetweenCondition(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		input    ConditionBuilder
// 		expected expectedExpression
// 		err      condErrorMode
// 	}{
// 		{
// 			name:  "basic method between for path",
// 			input: Path("foo").Between(Value(5), Value(7)),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": &dynamodb.AttributeValue{
// 						N: aws.String("5"),
// 					},
// 					":1": &dynamodb.AttributeValue{
// 						N: aws.String("7"),
// 					},
// 				},
// 				expectedExpression: "#0 BETWEEN :0 AND :1",
// 			},
// 		},
// 		{
// 			name:  "basic method between for value",
// 			input: Value(6).Between(Value(5), Value(7)),
//
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": &dynamodb.AttributeValue{
// 						N: aws.String("6"),
// 					},
// 					":1": &dynamodb.AttributeValue{
// 						N: aws.String("5"),
// 					},
// 					":2": &dynamodb.AttributeValue{
// 						N: aws.String("7"),
// 					},
// 				},
// 				expectedExpression: ":0 BETWEEN :1 AND :2",
// 			},
// 		},
// 		{
// 			name:  "basic method between for size",
// 			input: Path("foo").Size().Between(Value(5), Value(7)),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": &dynamodb.AttributeValue{
// 						N: aws.String("5"),
// 					},
// 					":1": &dynamodb.AttributeValue{
// 						N: aws.String("7"),
// 					},
// 				},
// 				expectedExpression: "size (#0) BETWEEN :0 AND :1",
// 			},
// 		},
// 		{
// 			name:  "invalid operand error between",
// 			input: Path("[5]").Between(Value(3), Path("foo..bar")),
// 			err:   invalidOperand,
// 		},
// 	}
//
// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			expr, err := c.input.BuildExpression()
// 			if c.err != noConditionError {
// 				if err == nil {
// 					t.Errorf("expect error %q, got no error", c.err)
// 				} else {
// 					if e, a := string(c.err), err.Error(); !strings.Contains(a, e) {
// 						t.Errorf("expect %q error message to be in %q", e, a)
// 					}
// 				}
// 			} else {
// 				if err != nil {
// 					t.Errorf("expect no error, got unexpected Error %q", err)
// 				}
//
// 				if e, a := c.expected, expr; !reflect.DeepEqual(a, e) {
// 					t.Errorf("expect %v, got %v", e, a)
// 				}
// 			}
// 		})
// 	}
// }
//
// func TestInCondition(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		input    ConditionBuilder
// 		expected expectedExpression
// 		err      condErrorMode
// 	}{
// 		{
// 			name:  "basic method in for path",
// 			input: Path("foo").In(Value(5), Value(7)),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": &dynamodb.AttributeValue{
// 						N: aws.String("5"),
// 					},
// 					":1": &dynamodb.AttributeValue{
// 						N: aws.String("7"),
// 					},
// 				},
// 				expectedExpression: "#0 IN (:0, :1)",
// 			},
// 		},
// 		{
// 			name:  "basic method in for value",
// 			input: Value(6).In(Value(5), Value(7)),
//
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": &dynamodb.AttributeValue{
// 						N: aws.String("6"),
// 					},
// 					":1": &dynamodb.AttributeValue{
// 						N: aws.String("5"),
// 					},
// 					":2": &dynamodb.AttributeValue{
// 						N: aws.String("7"),
// 					},
// 				},
// 				expectedExpression: ":0 IN (:1, :2)",
// 			},
// 		},
// 		{
// 			name:  "basic method in for size",
// 			input: Path("foo").Size().In(Value(5), Value(7)),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": &dynamodb.AttributeValue{
// 						N: aws.String("5"),
// 					},
// 					":1": &dynamodb.AttributeValue{
// 						N: aws.String("7"),
// 					},
// 				},
// 				expectedExpression: "size (#0) IN (:0, :1)",
// 			},
// 		},
// 		{
// 			name:  "invalid operand error in",
// 			input: Path("[5]").In(Value(3), Path("foo..bar")),
// 			err:   invalidOperand,
// 		},
// 	}
//
// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			expr, err := c.input.BuildExpression()
// 			if c.err != noConditionError {
// 				if err == nil {
// 					t.Errorf("expect error %q, got no error", c.err)
// 				} else {
// 					if e, a := string(c.err), err.Error(); !strings.Contains(a, e) {
// 						t.Errorf("expect %q error message to be in %q", e, a)
// 					}
// 				}
// 			} else {
// 				if err != nil {
// 					t.Errorf("expect no error, got unexpected Error %q", err)
// 				}
//
// 				if e, a := c.expected, expr; !reflect.DeepEqual(a, e) {
// 					t.Errorf("expect %v, got %v", e, a)
// 				}
// 			}
// 		})
// 	}
// }
//
// func TestAttrExistsCondition(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		input    ConditionBuilder
// 		expected expectedExpression
// 		err      condErrorMode
// 	}{
// 		{
// 			name:  "basic attr exists",
// 			input: Path("foo").AttributeExists(),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedExpression: "attribute_exists (#0)",
// 			},
// 		},
// 		{
// 			name:  "basic attr not exists",
// 			input: Path("foo").AttributeNotExists(),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedExpression: "attribute_not_exists (#0)",
// 			},
// 		},
// 		{
// 			name:  "invalid operand error attr exists",
// 			input: AttributeExists(Path("")),
// 			err:   invalidOperand,
// 		},
// 		{
// 			name:  "invalid operand error attr not exists",
// 			input: AttributeNotExists(Path("foo..bar")),
// 			err:   invalidOperand,
// 		},
// 	}
//
// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			expr, err := c.input.BuildExpression()
// 			if c.err != noConditionError {
// 				if err == nil {
// 					t.Errorf("expect error %q, got no error", c.err)
// 				} else {
// 					if e, a := string(c.err), err.Error(); !strings.Contains(a, e) {
// 						t.Errorf("expect %q error message to be in %q", e, a)
// 					}
// 				}
// 			} else {
// 				if err != nil {
// 					t.Errorf("expect no error, got unexpected Error %q", err)
// 				}
//
// 				if e, a := c.expected, expr; !reflect.DeepEqual(a, e) {
// 					t.Errorf("expect %v, got %v", e, a)
// 				}
// 			}
// 		})
// 	}
// }
//
// func TestAttrTypeCondition(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		input    ConditionBuilder
// 		expected expectedExpression
// 		err      condErrorMode
// 	}{
// 		{
// 			name:  "attr type String",
// 			input: Path("foo").AttributeType(String),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": {
// 						S: aws.String("S"),
// 					},
// 				},
// 				expectedExpression: "attribute_type (#0, :0)",
// 			},
// 		},
// 		{
// 			name:  "attr type StringSet",
// 			input: Path("foo").AttributeType(StringSet),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": {
// 						S: aws.String("SS"),
// 					},
// 				},
// 				expectedExpression: "attribute_type (#0, :0)",
// 			},
// 		},
// 		{
// 			name:  "attr type Number",
// 			input: Path("foo").AttributeType(Number),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": {
// 						S: aws.String("N"),
// 					},
// 				},
// 				expectedExpression: "attribute_type (#0, :0)",
// 			},
// 		},
// 		{
// 			name:  "attr type Number Set",
// 			input: Path("foo").AttributeType(NumberSet),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": {
// 						S: aws.String("NS"),
// 					},
// 				},
// 				expectedExpression: "attribute_type (#0, :0)",
// 			},
// 		},
// 		{
// 			name:  "attr type Binary",
// 			input: Path("foo").AttributeType(Binary),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": {
// 						S: aws.String("B"),
// 					},
// 				},
// 				expectedExpression: "attribute_type (#0, :0)",
// 			},
// 		},
// 		{
// 			name:  "attr type Binary Set",
// 			input: Path("foo").AttributeType(BinarySet),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": {
// 						S: aws.String("BS"),
// 					},
// 				},
// 				expectedExpression: "attribute_type (#0, :0)",
// 			},
// 		},
// 		{
// 			name:  "attr type Boolean",
// 			input: Path("foo").AttributeType(Boolean),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": {
// 						S: aws.String("BOOL"),
// 					},
// 				},
// 				expectedExpression: "attribute_type (#0, :0)",
// 			},
// 		},
// 		{
// 			name:  "attr type Null",
// 			input: Path("foo").AttributeType(Null),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": {
// 						S: aws.String("NULL"),
// 					},
// 				},
// 				expectedExpression: "attribute_type (#0, :0)",
// 			},
// 		},
// 		{
// 			name:  "attr type List",
// 			input: Path("foo").AttributeType(List),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": {
// 						S: aws.String("L"),
// 					},
// 				},
// 				expectedExpression: "attribute_type (#0, :0)",
// 			},
// 		},
// 		{
// 			name:  "attr type Map",
// 			input: Path("foo").AttributeType(Map),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": {
// 						S: aws.String("M"),
// 					},
// 				},
// 				expectedExpression: "attribute_type (#0, :0)",
// 			},
// 		},
// 		{
// 			name:  "attr type invalid operand",
// 			input: Path("").AttributeType(Map),
// 			err:   invalidOperand,
// 		},
// 	}
//
// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			expr, err := c.input.BuildExpression()
// 			if c.err != noConditionError {
// 				if err == nil {
// 					t.Errorf("expect error %q, got no error", c.err)
// 				} else {
// 					if e, a := string(c.err), err.Error(); !strings.Contains(a, e) {
// 						t.Errorf("expect %q error message to be in %q", e, a)
// 					}
// 				}
// 			} else {
// 				if err != nil {
// 					t.Errorf("expect no error, got unexpected Error %q", err)
// 				}
//
// 				if e, a := c.expected, expr; !reflect.DeepEqual(a, e) {
// 					t.Errorf("expect %v, got %v", e, a)
// 				}
// 			}
// 		})
// 	}
// }
//
// func TestBeginsWithCondition(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		input    ConditionBuilder
// 		expected expectedExpression
// 		err      condErrorMode
// 	}{
// 		{
// 			name:  "basic begins with",
// 			input: Path("foo").BeginsWith("bar"),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": {
// 						S: aws.String("bar"),
// 					},
// 				},
// 				expectedExpression: "begins_with (#0, :0)",
// 			},
// 		},
// 		{
// 			name:  "begins with invalid operand",
// 			input: Path("").BeginsWith("bar"),
// 			err:   invalidOperand,
// 		},
// 	}
//
// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			expr, err := c.input.BuildExpression()
// 			if c.err != noConditionError {
// 				if err == nil {
// 					t.Errorf("expect error %q, got no error", c.err)
// 				} else {
// 					if e, a := string(c.err), err.Error(); !strings.Contains(a, e) {
// 						t.Errorf("expect %q error message to be in %q", e, a)
// 					}
// 				}
// 			} else {
// 				if err != nil {
// 					t.Errorf("expect no error, got unexpected Error %q", err)
// 				}
//
// 				if e, a := c.expected, expr; !reflect.DeepEqual(a, e) {
// 					t.Errorf("expect %v, got %v", e, a)
// 				}
// 			}
// 		})
// 	}
// }
//
// func TestContainsCondition(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		input    ConditionBuilder
// 		expected expectedExpression
// 		err      condErrorMode
// 	}{
// 		{
// 			name:  "basic contains",
// 			input: Path("foo").Contains("bar"),
//
// 				expectedNames: map[string]*string{
// 					"#0": aws.String("foo"),
// 				},
// 				expectedValues: map[string]*dynamodb.AttributeValue{
// 					":0": {
// 						S: aws.String("bar"),
// 					},
// 				},
// 				expectedExpression: "contains (#0, :0)",
// 			},
// 		},
// 		{
// 			name:  "contains invalid operand",
// 			input: Path("").Contains("bar"),
// 			err:   invalidOperand,
// 		},
// 	}
//
// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			expr, err := c.input.BuildExpression()
// 			if c.err != noConditionError {
// 				if err == nil {
// 					t.Errorf("expect error %q, got no error", c.err)
// 				} else {
// 					if e, a := string(c.err), err.Error(); !strings.Contains(a, e) {
// 						t.Errorf("expect %q error message to be in %q", e, a)
// 					}
// 				}
// 			} else {
// 				if err != nil {
// 					t.Errorf("expect no error, got unexpected Error %q", err)
// 				}
//
// 				if e, a := c.expected, expr; !reflect.DeepEqual(a, e) {
// 					t.Errorf("expect %v, got %v", e, a)
// 				}
// 			}
// 		})
// 	}
// }
//
// func TestCompoundBuildCondition(t *testing.T) {
// 	cases := []struct {
// 		name      string
// 		inputCond ConditionBuilder
// 		expected  string
// 	}{
// 		{
// 			name: "and",
// 			inputCond: ConditionBuilder{
// 				conditionList: []ConditionBuilder{
// 					ConditionBuilder{},
// 					ConditionBuilder{},
// 					ConditionBuilder{},
// 					ConditionBuilder{},
// 				},
// 				mode: andCond,
// 			},
// 			expected: "($c) AND ($c) AND ($c) AND ($c)",
// 		},
// 		{
// 			name: "or",
// 			inputCond: ConditionBuilder{
// 				conditionList: []ConditionBuilder{
// 					ConditionBuilder{},
// 					ConditionBuilder{},
// 					ConditionBuilder{},
// 					ConditionBuilder{},
// 					ConditionBuilder{},
// 					ConditionBuilder{},
// 					ConditionBuilder{},
// 				},
// 				mode: orCond,
// 			},
// 			expected: "($c) OR ($c) OR ($c) OR ($c) OR ($c) OR ($c) OR ($c)",
// 		},
// 	}
//
// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			en, err := compoundBuildCondition(c.inputCond, ExprNode{})
// 			if err != nil {
// 				t.Errorf("expect no error, got unexpected Error %q", err)
// 			}
//
// 			if e, a := c.expected, en.fmtExpr; !reflect.DeepEqual(a, e) {
// 				t.Errorf("expect %v, got %v", e, a)
// 			}
// 		})
// 	}
// }
//
// func TestInBuildCondition(t *testing.T) {
// 	cases := []struct {
// 		name      string
// 		inputCond ConditionBuilder
// 		expected  string
// 	}{
// 		{
// 			name: "in",
// 			inputCond: ConditionBuilder{
// 				operandList: []OperandBuilder{
// 					PathBuilder{},
// 					PathBuilder{},
// 					PathBuilder{},
// 					PathBuilder{},
// 					PathBuilder{},
// 					PathBuilder{},
// 					PathBuilder{},
// 				},
// 				mode: andCond,
// 			},
// 			expected: "$c IN ($c, $c, $c, $c, $c, $c)",
// 		},
// 	}
//
// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			en, err := inBuildCondition(c.inputCond, ExprNode{})
// 			if err != nil {
// 				t.Errorf("expect no error, got unexpected Error %q", err)
// 			}
//
// 			if e, a := c.expected, en.fmtExpr; !reflect.DeepEqual(a, e) {
// 				t.Errorf("expect %v, got %v", e, a)
// 			}
// 		})
// 	}
// }

// If there is time implement mapEquals
