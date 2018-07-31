package ini

import (
	"reflect"
	"testing"
)

func TestIsNumberValue(t *testing.T) {
	cases := []struct {
		b        []rune
		expected bool
	}{
		{
			[]rune("123"),
			true,
		},
		{
			[]rune("-123"),
			true,
		},
		{
			[]rune("123.456"),
			true,
		},
		{
			[]rune("1e234"),
			true,
		},
		{
			[]rune("1E234"),
			true,
		},
		{
			[]rune("1ea4"),
			false,
		},
		{
			[]rune("1-23"),
			false,
		},
		{
			[]rune("-1-23"),
			false,
		},
		{
			[]rune("123-"),
			false,
		},
		{
			[]rune("a"),
			false,
		},
	}

	for i, c := range cases {
		if e, a := c.expected, isNumberValue(c.b); e != a {
			t.Errorf("%d: expected %t, but received %t", i+1, e, a)
		}
	}
}

// TODO: test errors
func TestNewLiteralToken(t *testing.T) {
	cases := []struct {
		b             []rune
		expectedRead  int
		expectedToken literalToken
		expectedError bool
	}{
		{
			b:            []rune("123"),
			expectedRead: 3,
			expectedToken: literalToken{
				Value: Value{
					Type:    IntegerType,
					integer: 123,
					raw:     []rune("123"),
				},
			},
		},
		{
			b:            []rune("123.456"),
			expectedRead: 7,
			expectedToken: literalToken{
				Value: Value{
					Type:    DecimalType,
					decimal: 123.456,
					raw:     []rune("123.456"),
				},
			},
		},
		{
			b:            []rune("123 456"),
			expectedRead: 3,
			expectedToken: literalToken{
				Value: Value{
					Type:    IntegerType,
					integer: 123,
					raw:     []rune("123"),
				},
			},
		},
		{
			b:            []rune("123 abc"),
			expectedRead: 3,
			expectedToken: literalToken{
				Value: Value{
					Type:    IntegerType,
					integer: 123,
					raw:     []rune("123"),
				},
			},
		},
		{
			b:            []rune(`"Hello" 123`),
			expectedRead: 7,
			expectedToken: literalToken{
				Value: Value{
					Type: QuotedStringType,
					raw:  []rune("Hello"),
				},
			},
		},
		{
			b:            []rune(`"Hello World"`),
			expectedRead: 13,
			expectedToken: literalToken{
				Value: Value{
					Type: QuotedStringType,
					raw:  []rune("Hello World"),
				},
			},
		},
		{
			b:            []rune("true"),
			expectedRead: 4,
			expectedToken: literalToken{
				Value: Value{
					Type:    BoolType,
					boolean: true,
					raw:     []rune("true"),
				},
			},
		},
		{
			b:            []rune("false"),
			expectedRead: 5,
			expectedToken: literalToken{
				Value: Value{
					Type:    BoolType,
					boolean: false,
					raw:     []rune("false"),
				},
			},
		},
	}

	for i, c := range cases {
		tok, n, err := newLitToken(c.b)

		if e, a := c.expectedToken.Value, tok.Value; !reflect.DeepEqual(e, a) {
			t.Errorf("%d: expected %v, but received %v", i+1, e, a)
		}

		if e, a := c.expectedRead, n; e != a {
			t.Errorf("%d: expected %v, but received %v", i+1, e, a)
		}

		if e, a := c.expectedError, err != nil; e != a {
			t.Errorf("%d: expected %v, but received %v", i+1, e, a)
		}
	}
}
