package expression

import (
	"reflect"
	"testing"
)

type testEquals interface {
	Equal(right OperandBuilder) CompareBuilder
}

//Compare
//Equal
func TestEquals(t *testing.T) {
	cases := []struct {
		equal                  testEquals
		rhs                    OperandBuilder
		expectedCompareBuilder CompareBuilder
	}{
		{
			equal: NewPath("foo.yay.cool.rad"),
			rhs:   NewPath("bar"),
			expectedCompareBuilder: CompareBuilder{
				Left:  NewPath("foo.yay.cool.rad"),
				Right: NewPath("bar"),
				Type:  "=",
			},
		},
		{
			equal: NewPath("foo.yay.cool.rad"),
			rhs:   NewValue(5),
			expectedCompareBuilder: CompareBuilder{
				Left:  NewPath("foo.yay.cool.rad"),
				Right: NewValue(5),
				Type:  "=",
			},
		},
		{
			equal: NewPath("foo.yay.cool.rad"),
			rhs:   NewPath("baz").Size(),
			expectedCompareBuilder: CompareBuilder{
				Left:  NewPath("foo.yay.cool.rad"),
				Right: NewPath("baz").Size(),
				Type:  "=",
			},
		},
		{
			equal: NewValue(map[string]int{
				"five": 5,
			}),
			rhs: NewPath("bar"),
			expectedCompareBuilder: CompareBuilder{
				Left: NewValue(map[string]int{
					"five": 5,
				}),
				Right: NewPath("bar"),
				Type:  "=",
			},
		},
		{
			equal: NewValue(map[string]int{
				"five": 5,
			}),
			rhs: NewValue(5),
			expectedCompareBuilder: CompareBuilder{
				Left: NewValue(map[string]int{
					"five": 5,
				}),
				Right: NewValue(5),
				Type:  "=",
			},
		},
		{
			equal: NewValue(map[string]int{
				"five": 5,
			}),
			rhs: NewPath("baz").Size(),
			expectedCompareBuilder: CompareBuilder{
				Left: NewValue(map[string]int{
					"five": 5,
				}),
				Right: NewPath("baz").Size(),
				Type:  "=",
			},
		},
		{
			equal: NewPath("foo[1]").Size(),
			rhs:   NewPath("bar"),
			expectedCompareBuilder: CompareBuilder{
				Left:  NewPath("foo[1]").Size(),
				Right: NewPath("bar"),
				Type:  "=",
			},
		},
		{
			equal: NewPath("foo[1]").Size(),
			rhs:   NewValue(5),
			expectedCompareBuilder: CompareBuilder{
				Left:  NewPath("foo[1]").Size(),
				Right: NewValue(5),
				Type:  "=",
			},
		},
		{
			equal: NewPath("foo[1]").Size(),
			rhs:   NewPath("baz").Size(),
			expectedCompareBuilder: CompareBuilder{
				Left:  NewPath("foo[1]").Size(),
				Right: NewPath("baz").Size(),
				Type:  "=",
			},
		},
	}
	for _, c := range cases {
		input := c.equal.Equal(c.rhs)
		expr, err := input.BuildCondition()
		if err != nil {
			t.Error(err)
		}
		expected, err := c.expectedCompareBuilder.BuildCondition()
		if err != nil {
			t.Error(err)
		}

		if reflect.DeepEqual(expr, expected) != true {
			t.Errorf("Condition Equal with input %#v returned %#v, expected %#v", input, expr, expected)
		}
	}
}

// If there is time implement mapEquals
