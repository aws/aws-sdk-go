package ini

import (
	"reflect"
	"testing"
)

func TestIsNumberValue(t *testing.T) {
	cases := []struct {
		b        []byte
		expected bool
	}{
		{
			[]byte("123"),
			true,
		},
		{
			[]byte("123.456"),
			true,
		},
		{
			[]byte("a"),
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
		b             []byte
		expectedRead  int
		expectedToken literalToken
		expectedError bool
	}{
		{
			b:            []byte("123"),
			expectedRead: 3,
			expectedToken: literalToken{
				Value: UnionValue{
					Type:    IntegerType,
					integer: 123,
				},
			},
		},
		{
			b:            []byte("123.456"),
			expectedRead: 7,
			expectedToken: literalToken{
				Value: UnionValue{
					Type:    DecimalType,
					decimal: 123.456,
				},
			},
		},
		{
			b:            []byte("123 456"),
			expectedRead: 3,
			expectedToken: literalToken{
				Value: UnionValue{
					Type:    IntegerType,
					integer: 123,
				},
			},
		},
		{
			b:            []byte("123 abc"),
			expectedRead: 3,
			expectedToken: literalToken{
				Value: UnionValue{
					Type:    IntegerType,
					integer: 123,
				},
			},
		},
		{
			b:            []byte(`"Hello" 123`),
			expectedRead: 7,
			expectedToken: literalToken{
				Value: UnionValue{
					Type: StringType,
					str:  `Hello`,
				},
			},
		},
		{
			b:            []byte(`"Hello World"`),
			expectedRead: 13,
			expectedToken: literalToken{
				Value: UnionValue{
					Type: StringType,
					str:  "Hello World",
				},
			},
		},
		{
			b:            []byte("true"),
			expectedRead: 4,
			expectedToken: literalToken{
				Value: UnionValue{
					Type:    BoolType,
					boolean: true,
				},
			},
		},
		{
			b:            []byte("false"),
			expectedRead: 5,
			expectedToken: literalToken{
				Value: UnionValue{
					Type:    BoolType,
					boolean: false,
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
