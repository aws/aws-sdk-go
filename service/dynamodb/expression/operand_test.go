package expression

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

func TestBuildOperand(t *testing.T) {
	cases := []struct {
		input    OperandBuilder
		expected Expression
	}{
		{
			input: NewPath("foo"),
			expected: Expression{
				Names: map[string]*string{
					"#" + encode("foo"): aws.String("foo"),
				},
				Expression: "#" + encode("foo"),
			},
		},
	}

	for _, c := range cases {
		operand, err := c.input.BuildOperand()
		if err != nil {
			t.Error(err)
		}

		if reflect.DeepEqual(operand, c.expected) != true {
			t.Errorf("BuildOperand with input %#v returned %#v, expected %#v", c.input, operand, c.expected)
		}
	}
}
