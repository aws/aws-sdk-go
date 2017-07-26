package expression

import (
	"reflect"
	"testing"
)

type testEquals interface {
	Equal(right OperandBuilder) Condition
}

//Compare
//Equal
func TestEquals(t *testing.T) {
	cases := []struct {
		equal             testEquals
		rhs               OperandBuilder
		expectedCondition Condition
	}{
		{
			equal: NewPath("foo.yay.cool.rad"),
			rhs:   NewPath("bar"),
			expectedCondition: Condition{
				OperandList: []OperandBuilder{
					NewPath("foo.yay.cool.rad"),
					NewPath("bar"),
				},
				Mode: EqualCond,
			},
		},
		{
			equal: NewPath("foo.yay.cool.rad"),
			rhs:   NewValue(5),
			expectedCondition: Condition{
				OperandList: []OperandBuilder{
					NewPath("foo.yay.cool.rad"),
					NewValue(5),
				},
				Mode: EqualCond,
			},
		},
		{
			equal: NewPath("foo.yay.cool.rad"),
			rhs:   NewPath("baz").Size(),
			expectedCondition: Condition{
				OperandList: []OperandBuilder{
					NewPath("foo.yay.cool.rad"),
					NewPath("baz").Size(),
				},
				Mode: EqualCond,
			},
		},
		{
			equal: NewValue(5),
			rhs:   NewPath("bar"),
			expectedCondition: Condition{
				OperandList: []OperandBuilder{
					NewValue(5),
					NewPath("bar"),
				},
				Mode: EqualCond,
			},
		},
		{
			equal: NewValue(map[string]int{
				"five": 5,
			}),
			rhs: NewPath("bar"),
			expectedCondition: Condition{
				OperandList: []OperandBuilder{
					NewValue(map[string]int{
						"five": 5,
					}),
					NewPath("bar"),
				},
				Mode: EqualCond,
			},
		},
		{
			equal: NewValue(map[string]int{
				"five": 5,
			}),
			rhs: NewValue(5),
			expectedCondition: Condition{
				OperandList: []OperandBuilder{
					NewValue(map[string]int{
						"five": 5,
					}),
					NewValue(5),
				},
				Mode: EqualCond,
			},
		},
		{
			equal: NewValue(map[string]int{
				"five": 5,
			}),
			rhs: NewPath("baz").Size(),
			expectedCondition: Condition{
				OperandList: []OperandBuilder{
					NewValue(map[string]int{
						"five": 5,
					}),
					NewPath("baz").Size(),
				},
				Mode: EqualCond,
			},
		},
		{
			equal: NewPath("foo[1]").Size(),
			rhs:   NewPath("bar"),
			expectedCondition: Condition{
				OperandList: []OperandBuilder{
					NewPath("foo[1]").Size(),
					NewPath("bar"),
				},
				Mode: EqualCond,
			},
		},
		{
			equal: NewPath("foo[1]").Size(),
			rhs:   NewValue(5),
			expectedCondition: Condition{
				OperandList: []OperandBuilder{
					NewPath("foo[1]").Size(),
					NewValue(5),
				},
				Mode: EqualCond,
			},
		},
		{
			equal: NewPath("foo[1]").Size(),
			rhs:   NewPath("baz").Size(),
			expectedCondition: Condition{
				OperandList: []OperandBuilder{
					NewPath("foo[1]").Size(),
					NewPath("baz").Size(),
				},
				Mode: EqualCond,
			},
		},
		{
			equal: NewPath("foo.bar.baz").Size(),
			rhs:   NewPath("bar.qux.foo").Size(),
			expectedCondition: Condition{
				OperandList: []OperandBuilder{
					NewPath("foo.bar.baz").Size(),
					NewPath("bar.qux.foo").Size(),
				},
				Mode: EqualCond,
			},
		},
	}
	for testNumber, c := range cases {
		input := c.equal.Equal(c.rhs)
		expr, err := input.BuildCondition()
		if err != nil {
			t.Errorf("TestEquals Test Number %#v: Unexpected Error %#v", testNumber, err)
		}
		expected, err := c.expectedCondition.BuildCondition()
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
		input                 Condition
		expected              Expression
		buildListOperandError bool
		noMatchError          bool
		operandNumberError    bool
		conditionNumberError  bool
	}{
		{
			input: Condition{
				OperandList: []OperandBuilder{
					PathBuilder{
						path: "",
					},
				},
			},
			expected:              Expression{},
			buildListOperandError: true,
		},
		{
			input: Condition{
				Mode: 0,
			},
			noMatchError: true,
		},
		{
			input: Condition{
				Mode: EqualCond,
			},
			operandNumberError: true,
		},
		{
			input: Condition{
				Mode: EqualCond,
				ConditionList: []Condition{
					Condition{},
				},
			},
			conditionNumberError: true,
		},
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
