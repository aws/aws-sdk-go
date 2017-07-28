package expression

import (
	"reflect"
	"testing"
)

type testCompare interface {
	Equal(right OperandBuilder) ConditionBuilder
}

//Compare
//Equal
func TestCompare(t *testing.T) {
	cases := []struct {
		lhs      OperandBuilder
		rhs      OperandBuilder
		expected ConditionBuilder
	}{
		{
			lhs: NewPath("foo.yay.cool.rad"),
			rhs: NewPath("bar"),
			expected: ConditionBuilder{
				operandList: []OperandExpression{
					OperandExpression{
						Mode: PathOpe,
						path: "foo.yay.cool.rad",
					},
					OperandExpression{
						Mode: PathOpe,
						path: "bar",
					},
				},
			},
		},
		// {
		// 	equal: NewPath("foo.yay.cool.rad"),
		// 	rhs:   NewValue(5),
		// 	expectedCondition: ConditionBuilder{
		// 		OperandList: []OperandBuilder{
		// 			NewPath("foo.yay.cool.rad"),
		// 			NewValue(5),
		// 		},
		// 		Mode: EqualCond,
		// 	},
		// },
		// {
		// 	equal: NewPath("foo.yay.cool.rad"),
		// 	rhs:   NewPath("baz").Size(),
		// 	expectedCondition: ConditionBuilder{
		// 		OperandList: []OperandBuilder{
		// 			NewPath("foo.yay.cool.rad"),
		// 			NewPath("baz").Size(),
		// 		},
		// 		Mode: EqualCond,
		// 	},
		// },
		// {
		// 	equal: NewValue(5),
		// 	rhs:   NewPath("bar"),
		// 	expectedCondition: ConditionBuilder{
		// 		OperandList: []OperandBuilder{
		// 			NewValue(5),
		// 			NewPath("bar"),
		// 		},
		// 		Mode: EqualCond,
		// 	},
		// },
		// {
		// 	equal: NewValue(map[string]int{
		// 		"five": 5,
		// 	}),
		// 	rhs: NewPath("bar"),
		// 	expectedCondition: ConditionBuilder{
		// 		OperandList: []OperandBuilder{
		// 			NewValue(map[string]int{
		// 				"five": 5,
		// 			}),
		// 			NewPath("bar"),
		// 		},
		// 		Mode: EqualCond,
		// 	},
		// },
		// {
		// 	equal: NewValue(map[string]int{
		// 		"five": 5,
		// 	}),
		// 	rhs: NewValue(5),
		// 	expectedCondition: ConditionBuilder{
		// 		OperandList: []OperandBuilder{
		// 			NewValue(map[string]int{
		// 				"five": 5,
		// 			}),
		// 			NewValue(5),
		// 		},
		// 		Mode: EqualCond,
		// 	},
		// },
		// {
		// 	equal: NewValue(map[string]int{
		// 		"five": 5,
		// 	}),
		// 	rhs: NewPath("baz").Size(),
		// 	expectedCondition: ConditionBuilder{
		// 		OperandList: []OperandBuilder{
		// 			NewValue(map[string]int{
		// 				"five": 5,
		// 			}),
		// 			NewPath("baz").Size(),
		// 		},
		// 		Mode: EqualCond,
		// 	},
		// },
		// {
		// 	equal: NewPath("foo[1]").Size(),
		// 	rhs:   NewPath("bar"),
		// 	expectedCondition: ConditionBuilder{
		// 		OperandList: []OperandBuilder{
		// 			NewPath("foo[1]").Size(),
		// 			NewPath("bar"),
		// 		},
		// 		Mode: EqualCond,
		// 	},
		// },
		// {
		// 	equal: NewPath("foo[1]").Size(),
		// 	rhs:   NewValue(5),
		// 	expectedCondition: ConditionBuilder{
		// 		OperandList: []OperandBuilder{
		// 			NewPath("foo[1]").Size(),
		// 			NewValue(5),
		// 		},
		// 		Mode: EqualCond,
		// 	},
		// },
		// {
		// 	equal: NewPath("foo[1]").Size(),
		// 	rhs:   NewPath("baz").Size(),
		// 	expectedCondition: ConditionBuilder{
		// 		OperandList: []OperandBuilder{
		// 			NewPath("foo[1]").Size(),
		// 			NewPath("baz").Size(),
		// 		},
		// 		Mode: EqualCond,
		// 	},
		// },
		// {
		// 	equal: NewPath("foo.bar.baz").Size(),
		// 	rhs:   NewPath("bar.qux.foo").Size(),
		// 	expectedCondition: ConditionBuilder{
		// 		OperandList: []OperandBuilder{
		// 			NewPath("foo.bar.baz").Size(),
		// 			NewPath("bar.qux.foo").Size(),
		// 		},
		// 		Mode: EqualCond,
		// 	},
		// },
	}
	for testNumber, c := range cases {

		c.expected.Mode = EqualCond
		expr, err := c.lhs.Equal(c.rhs).BuildCondition()
		if err != nil {
			t.Errorf("TestEquals Test Number %#v: Unexpected Error %#v", testNumber, err)
		}
		expected, err := c.expected.BuildCondition()
		if err != nil {
			t.Errorf("TestEquals Test Number %#v: Unexpected Error %#v", testNumber, err)
		}

		if reflect.DeepEqual(expr, expected) != true {
			t.Errorf("TestEquals Test Number %#v: Expected %#v, got %#v", testNumber, expected, expr)
		}

		expr, err = Equal(c.lhs, c.rhs).BuildCondition()
		if err != nil {
			t.Errorf("TestEquals Test Number %#v: Unexpected Error %#v", testNumber, err)
		}
		expected, err = c.expected.BuildCondition()
		if err != nil {
			t.Errorf("TestEquals Test Number %#v: Unexpected Error %#v", testNumber, err)
		}

		if reflect.DeepEqual(expr, expected) != true {
			t.Errorf("TestEquals Test Number %#v: Expected %#v, got %#v", testNumber, expected, expr)
		}
	}
}

func TestBuildCondition(t *testing.T) {
	cases := []struct {
		input                 ConditionBuilder
		expected              Expression
		buildListOperandError bool
		noMatchError          bool
		operandNumberError    bool
		conditionNumberError  bool
	}{
		{
			input: ConditionBuilder{
				operandList: []OperandExpression{
					OperandExpression{
						Mode: PathOpe,
						path: "",
					},
				},
			},
			expected:              Expression{},
			buildListOperandError: true,
		},
		// {
		// 	input: ConditionBuilder{
		// 		Mode: 0,
		// 	},
		// 	noMatchError: true,
		// },
		// {
		// 	input: ConditionBuilder{
		// 		Mode: EqualCond,
		// 	},
		// 	operandNumberError: true,
		// },
		// {
		// 	input: ConditionBuilder{
		// 		Mode: EqualCond,
		// 		ConditionList: []ConditionBuilder{
		// 			ConditionBuilder{},
		// 		},
		// 	},
		// 	conditionNumberError: true,
		// },
	}

	for testNumber, c := range cases {
		expr, err := c.input.BuildCondition()

		if c.buildListOperandError {
			if err == nil {
				t.Errorf("TestBuildCondition Test Number %#v: Expected list operand error but got no error", testNumber)
			} else {
				continue
			}
		}

		if c.noMatchError {
			if err == nil {
				t.Errorf("TestBuildCondition Test Number %#v: Expected no matching mode error but got no error", testNumber)
			} else {
				continue
			}
		}

		if c.operandNumberError {
			if err == nil {
				t.Errorf("TestBuildCondition Test Number %#v: Expected operand number error but got no error", testNumber)
			} else {
				continue
			}
		}

		if c.conditionNumberError {
			if err == nil {
				t.Errorf("TestBuildCondition Test Number %#v: Expected condition number error but got no error", testNumber)
			} else {
				continue
			}
		}

		if err != nil {
			t.Errorf("TestBuildCondition Test Number %#v: Unexpected Error %#v", testNumber, err)
		}

		if reflect.DeepEqual(expr, c.expected) != true {
			t.Errorf("TestBuildCondition Test Number %#v: Expected %#v, got %#v", testNumber, c.expected, expr)
		}
	}
}

// If there is time implement mapEquals
