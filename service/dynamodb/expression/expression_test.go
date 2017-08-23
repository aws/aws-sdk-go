// +build go1.7

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
	nilAliasList = "AliasList is nil"
	// invalidFactoryBuilderOperand error will occur if an invalid operand is used
	// as input for BuildFactory()
	invalidFactoryBuildOperand = "BuildOperand error"
	// emptyFactoryBuilder error will occur if BuildFactory() is called on an
	// empty FactoryBuilder
	emptyFactoryBuilder = "EmptyFactoryBuilder"
)

func TestBuildFactory(t *testing.T) {
	cases := []struct {
		name     string
		input    FactoryBuilder
		expected Factory
		err      exprErrorMode
	}{
		{
			name:  "condition",
			input: Condition(Name("foo").Equal(Value(5))),
			expected: Factory{
				expressionMap: map[expressionType]TreeBuilder{
					condition: ConditionBuilder{
						operandList: []OperandBuilder{
							NameBuilder{
								name: "foo",
							},
							ValueBuilder{
								value: 5,
							},
						},
						mode: equalCond,
					},
				},
			},
		},
		{
			name:  "projection",
			input: Projection(NamesList(Name("foo"), Name("bar"), Name("baz"))),
			expected: Factory{
				expressionMap: map[expressionType]TreeBuilder{
					projection: ProjectionBuilder{
						names: []NameBuilder{
							{
								name: "foo",
							},
							{
								name: "bar",
							},
							{
								name: "baz",
							},
						},
					},
				},
			},
		},
		// {
		// 	name:  "keyCondition",
		// 	input: KeyCondition(Name("foo").Equal(Value(5))),
		// 	expected: Factory{
		// 		expressionMap: map[expressionType]TreeBuilder{
		// 			keyCondition: KeyConditionBuilder{
		// 				operandList: []OperandBuilder{
		// 					NameBuilder{
		// 						name: "foo",
		// 					},
		// 					ValueBuilder{
		// 						value: 5,
		// 					},
		// 				},
		// 				mode: equalCond,
		// 			},
		// 		},
		// 	},
		// },
		{
			name:  "filter",
			input: Filter(Name("foo").Equal(Value(5))),
			expected: Factory{
				expressionMap: map[expressionType]TreeBuilder{
					filter: ConditionBuilder{
						operandList: []OperandBuilder{
							NameBuilder{
								name: "foo",
							},
							ValueBuilder{
								value: 5,
							},
						},
						mode: equalCond,
					},
				},
			},
		},
		// {
		// 	name:  "update",
		// 	input: Update(Name("foo").Equal(Value(5))),
		// 	expected: Factory{
		// 		expressionMap: map[expressionType]TreeBuilder{
		// 			update: UpdateBuilder{
		// 				operandList: []OperandBuilder{
		// 					NameBuilder{
		// 						name: "foo",
		// 					},
		// 					ValueBuilder{
		// 						value: 5,
		// 					},
		// 				},
		// 				mode: equalCond,
		// 			},
		// 		},
		// 	},
		// },
		{
			name:  "compound",
			input: Condition(Name("foo").Equal(Value(5))).Filter(Name("bar").LessThan(Value(6))).Projection(NamesList(Name("foo"), Name("bar"), Name("baz"))),
			expected: Factory{
				expressionMap: map[expressionType]TreeBuilder{
					condition: ConditionBuilder{
						operandList: []OperandBuilder{
							NameBuilder{
								name: "foo",
							},
							ValueBuilder{
								value: 5,
							},
						},
						mode: equalCond,
					},
					filter: ConditionBuilder{
						operandList: []OperandBuilder{
							NameBuilder{
								name: "bar",
							},
							ValueBuilder{
								value: 6,
							},
						},
						mode: lessThanCond,
					},
					projection: ProjectionBuilder{
						names: []NameBuilder{
							{
								name: "foo",
							},
							{
								name: "bar",
							},
							{
								name: "baz",
							},
						},
					},
				},
			},
		},
		{
			name:  "invalid FactoryBuilder",
			input: Condition(Name("").Equal(Value(5))),
			err:   invalidFactoryBuildOperand,
		},
		{
			name:  "empty FactoryBuilder",
			input: FactoryBuilder{},
			err:   emptyFactoryBuilder,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := c.input.BuildFactory()
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

				if e, a := c.expected, actual; !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestCondition(t *testing.T) {
	cases := []struct {
		name     string
		input    Factory
		expected *string
	}{
		{
			name: "condition",
			input: Factory{
				expressionMap: map[expressionType]TreeBuilder{
					condition: Name("foo").Equal(Value(5)),
				},
			},
			expected: aws.String("#0 = :0"),
		},
		{
			name:     "nil",
			input:    Factory{},
			expected: nil,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.input.Condition()
			if e, a := c.expected, actual; !reflect.DeepEqual(a, e) {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	cases := []struct {
		name     string
		input    Factory
		expected *string
	}{
		{
			name: "filter",
			input: Factory{
				expressionMap: map[expressionType]TreeBuilder{
					filter: Name("foo").Equal(Value(5)),
				},
			},
			expected: aws.String("#0 = :0"),
		},
		{
			name:     "nil",
			input:    Factory{},
			expected: nil,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.input.Filter()
			if e, a := c.expected, actual; !reflect.DeepEqual(a, e) {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

func TestProjection(t *testing.T) {
	cases := []struct {
		name     string
		input    Factory
		expected *string
	}{
		{
			name: "projection",
			input: Factory{
				expressionMap: map[expressionType]TreeBuilder{
					projection: NamesList(Name("foo"), Name("bar"), Name("baz")),
				},
			},
			expected: aws.String("#0, #1, #2"),
		},
		{
			name:     "nil",
			input:    Factory{},
			expected: nil,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.input.Projection()
			if e, a := c.expected, actual; !reflect.DeepEqual(a, e) {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

func TestNames(t *testing.T) {
	cases := []struct {
		name     string
		input    Factory
		expected map[string]*string
	}{
		{
			name: "projection",
			input: Factory{
				expressionMap: map[expressionType]TreeBuilder{
					projection: NamesList(Name("foo"), Name("bar"), Name("baz")),
				},
			},
			expected: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("bar"),
				"#2": aws.String("baz"),
			},
		},
		{
			name: "aggregate",
			input: Factory{
				expressionMap: map[expressionType]TreeBuilder{
					condition: ConditionBuilder{
						operandList: []OperandBuilder{
							NameBuilder{
								name: "foo",
							},
							ValueBuilder{
								value: 5,
							},
						},
						mode: equalCond,
					},
					filter: ConditionBuilder{
						operandList: []OperandBuilder{
							NameBuilder{
								name: "bar",
							},
							ValueBuilder{
								value: 6,
							},
						},
						mode: lessThanCond,
					},
					projection: ProjectionBuilder{
						names: []NameBuilder{
							{
								name: "foo",
							},
							{
								name: "bar",
							},
							{
								name: "baz",
							},
						},
					},
				},
			},
			expected: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("bar"),
				"#2": aws.String("baz"),
			},
		},
		{
			name:     "empty",
			input:    Factory{},
			expected: map[string]*string{},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.input.Names()
			if e, a := c.expected, actual; !reflect.DeepEqual(a, e) {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

func TestValues(t *testing.T) {
	cases := []struct {
		name     string
		input    Factory
		expected map[string]*dynamodb.AttributeValue
	}{
		{
			name: "condition",
			input: Factory{
				expressionMap: map[expressionType]TreeBuilder{
					condition: Name("foo").Equal(Value(5)),
				},
			},
			expected: map[string]*dynamodb.AttributeValue{
				":0": {
					N: aws.String("5"),
				},
			},
		},
		{
			name: "aggregate",
			input: Factory{
				expressionMap: map[expressionType]TreeBuilder{
					condition: ConditionBuilder{
						operandList: []OperandBuilder{
							NameBuilder{
								name: "foo",
							},
							ValueBuilder{
								value: 5,
							},
						},
						mode: equalCond,
					},
					filter: ConditionBuilder{
						operandList: []OperandBuilder{
							NameBuilder{
								name: "bar",
							},
							ValueBuilder{
								value: 6,
							},
						},
						mode: lessThanCond,
					},
					projection: ProjectionBuilder{
						names: []NameBuilder{
							{
								name: "foo",
							},
							{
								name: "bar",
							},
							{
								name: "baz",
							},
						},
					},
				},
			},
			expected: map[string]*dynamodb.AttributeValue{
				":0": {
					N: aws.String("5"),
				},
				":1": {
					N: aws.String("6"),
				},
			},
		},
		{
			name:     "empty",
			input:    Factory{},
			expected: map[string]*dynamodb.AttributeValue{},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.input.Values()
			if e, a := c.expected, actual; !reflect.DeepEqual(a, e) {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

func TestBuildChildTrees(t *testing.T) {
	cases := []struct {
		name              string
		input             Factory
		expectedAliasList *AliasList
		expectedStringMap map[expressionType]string
	}{
		{
			name: "aggregate",
			input: Factory{
				expressionMap: map[expressionType]TreeBuilder{
					condition: ConditionBuilder{
						operandList: []OperandBuilder{
							NameBuilder{
								name: "foo",
							},
							ValueBuilder{
								value: 5,
							},
						},
						mode: equalCond,
					},
					filter: ConditionBuilder{
						operandList: []OperandBuilder{
							NameBuilder{
								name: "bar",
							},
							ValueBuilder{
								value: 6,
							},
						},
						mode: lessThanCond,
					},
					projection: ProjectionBuilder{
						names: []NameBuilder{
							{
								name: "foo",
							},
							{
								name: "bar",
							},
							{
								name: "baz",
							},
						},
					},
				},
			},
			expectedAliasList: &AliasList{
				namesList: []string{"foo", "bar", "baz"},
				valuesList: []dynamodb.AttributeValue{
					{
						N: aws.String("5"),
					},
					{
						N: aws.String("6"),
					},
				},
			},
			expectedStringMap: map[expressionType]string{
				condition:  "#0 = :0",
				filter:     "#1 < :1",
				projection: "#0, #1, #2",
			},
		},
		{
			name:              "empty",
			input:             Factory{},
			expectedAliasList: &AliasList{},
			expectedStringMap: map[expressionType]string{},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualAL, actualSM := c.input.buildChildTrees()
			if e, a := c.expectedAliasList, actualAL; !reflect.DeepEqual(a, e) {
				t.Errorf("expect %v, got %v", e, a)
			}
			if e, a := c.expectedStringMap, actualSM; !reflect.DeepEqual(a, e) {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

func TestBuildExpressionString(t *testing.T) {
	cases := []struct {
		name               string
		input              ExprNode
		expectedNames      map[string]*string
		expectedValues     map[string]*dynamodb.AttributeValue
		expectedExpression string
		err                exprErrorMode
	}{
		{
			name: "basic name",
			input: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "$n",
			},

			expectedValues: map[string]*dynamodb.AttributeValue{},
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedExpression: "#0",
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
			expectedNames: map[string]*string{},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: ":0",
		},
		{
			name: "nested path",
			input: ExprNode{
				names:   []string{"foo", "bar"},
				fmtExpr: "$n.$n",
			},

			expectedValues: map[string]*dynamodb.AttributeValue{},
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("bar"),
			},
			expectedExpression: "#0.#1",
		},
		{
			name: "nested path with index",
			input: ExprNode{
				names:   []string{"foo", "bar", "baz"},
				fmtExpr: "$n.$n[0].$n",
			},
			expectedValues: map[string]*dynamodb.AttributeValue{},
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
				"#1": aws.String("bar"),
				"#2": aws.String("baz"),
			},
			expectedExpression: "#0.#1[0].#2",
		},
		{
			name: "basic size",
			input: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "size ($n)",
			},
			expectedValues: map[string]*dynamodb.AttributeValue{},
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedExpression: "size (#0)",
		},
		{
			name: "duplicate path name",
			input: ExprNode{
				names:   []string{"foo", "foo"},
				fmtExpr: "$n.$n",
			},
			expectedValues: map[string]*dynamodb.AttributeValue{},
			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedExpression: "#0.#0",
		},
		{
			name: "equal expression",
			input: ExprNode{
				children: []ExprNode{
					ExprNode{
						names:   []string{"foo"},
						fmtExpr: "$n",
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

			expectedNames: map[string]*string{
				"#0": aws.String("foo"),
			},
			expectedValues: map[string]*dynamodb.AttributeValue{
				":0": &dynamodb.AttributeValue{
					N: aws.String("5"),
				},
			},
			expectedExpression: "#0 = :0",
		},
		{
			name: "missing char after $",
			input: ExprNode{
				names:   []string{"foo", "foo"},
				fmtExpr: "$n.$",
			},
			err: invalidEscChar,
		},
		{
			name: "names out of range",
			input: ExprNode{
				names:   []string{"foo"},
				fmtExpr: "$n.$n",
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
			name:               "empty ExprNode",
			input:              ExprNode{},
			expectedExpression: "",
		},
		{
			name:  "nil aliasList",
			input: ExprNode{},
			err:   nilAliasList,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var expr string
			var err error
			if c.err == nilAliasList {
				expr, err = c.input.BuildExpressionString(nil)
			} else {
				expr, err = c.input.BuildExpressionString(&AliasList{})
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

				if e, a := c.expectedExpression, expr; !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestAliasValue(t *testing.T) {
	cases := []struct {
		name     string
		input    *AliasList
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
			input:    &AliasList{},
			expected: ":0",
		},
		{
			name: "fifth item",
			input: &AliasList{
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
		inputList *AliasList
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
			inputList: &AliasList{},
			inputName: "foo",
			expected:  "#0",
		},
		{
			name: "duplicate item",
			inputList: &AliasList{
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
