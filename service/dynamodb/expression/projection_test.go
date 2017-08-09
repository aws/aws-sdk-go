package expression

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

func TestProjection(t *testing.T) {
	cases := []struct {
		name     string
		input    ProjectionBuilder
		expected Expression
	}{
		{
			name:  "basic projection",
			input: Projection(Path("foo"), Path("bar")),
			expected: Expression{
				Names: map[string]*string{
					"#0": aws.String("foo"),
					"#1": aws.String("bar"),
				},
				Expression: "#0, #1",
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr, err := c.input.BuildExpression()
			if err != nil {
				t.Errorf("expect no error, got unexpected Error %q", err)
			}

			if e, a := c.expected, expr; !reflect.DeepEqual(a, e) {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}
