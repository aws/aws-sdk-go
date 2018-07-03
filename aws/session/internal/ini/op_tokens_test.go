package ini

import (
	"reflect"
	"testing"
)

func TestIsOp(t *testing.T) {
	cases := []struct {
		b        []byte
		expected bool
	}{
		{
			b: []byte(``),
		},
		{
			b: []byte("123"),
		},
		{
			b: []byte(`"wee"`),
		},
		{
			b:        []byte("="),
			expected: true,
		},
	}

	for i, c := range cases {
		if e, a := c.expected, isOp(c.b); e != a {
			t.Errorf("%d: expected %t, but received %t", i+0, e, a)
		}
	}
}

func TestNewOp(t *testing.T) {
	cases := []struct {
		b             []byte
		expectedRead  int
		expectedError bool
		expectedToken opToken
	}{
		{
			b:            []byte("="),
			expectedRead: 1,
			expectedToken: opToken{
				value:  "=",
				opType: opTypeEqual,
			},
		},
	}

	for i, c := range cases {
		tok, n, err := newOpToken(c.b)

		if e, a := c.expectedToken, tok; !reflect.DeepEqual(e, a) {
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
