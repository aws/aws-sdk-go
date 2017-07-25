package expression

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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
		{
			input: NewValue(5),
			expected: Expression{
				Values: map[string]*dynamodb.AttributeValue{
					":" + encode(fmt.Sprint(dynamodb.AttributeValue{
						N: aws.String("5"),
					})): &dynamodb.AttributeValue{
						N: aws.String("5"),
					},
				},
				Expression: ":" + encode(fmt.Sprint(dynamodb.AttributeValue{
					N: aws.String("5"),
				})),
			},
		},
		{
			input: NewPath("foo.bar[2].baz"),
			expected: Expression{
				Names: map[string]*string{
					"#" + encode("foo"): aws.String("foo"),
					"#" + encode("bar"): aws.String("bar"),
					"#" + encode("baz"): aws.String("baz"),
				},
				Expression: "#" + encode("foo") + ".#" + encode("bar") + "[2]" + ".#" + encode("baz"),
			},
		},
		{
			input: NewValue(map[string]int{
				"even": 2,
				"odd":  1,
			}),
			expected: Expression{
				Values: map[string]*dynamodb.AttributeValue{
					":" + encode(fmt.Sprint(dynamodb.AttributeValue{
						M: map[string]*dynamodb.AttributeValue{
							"even": &dynamodb.AttributeValue{
								N: aws.String("2"),
							},
							"odd": &dynamodb.AttributeValue{
								N: aws.String("1"),
							},
						},
					})): &dynamodb.AttributeValue{
						M: map[string]*dynamodb.AttributeValue{
							"even": &dynamodb.AttributeValue{
								N: aws.String("2"),
							},
							"odd": &dynamodb.AttributeValue{
								N: aws.String("1"),
							},
						},
					},
				},
				Expression: ":" + encode(fmt.Sprint(dynamodb.AttributeValue{
					M: map[string]*dynamodb.AttributeValue{
						"even": &dynamodb.AttributeValue{
							N: aws.String("2"),
						},
						"odd": &dynamodb.AttributeValue{
							N: aws.String("1"),
						},
					},
				})),
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
