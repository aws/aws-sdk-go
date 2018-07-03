package ini

import (
	"reflect"
	"testing"
)

func TestIsSep(t *testing.T) {
	cases := []struct {
		b        []byte
		expected bool
	}{
		{
			b: []byte(``),
		},
		{
			b: []byte(`"wee"`),
		},
		{
			b:        []byte("["),
			expected: true,
		},
		{
			b:        []byte("]"),
			expected: true,
		},
	}

	for i, c := range cases {
		if e, a := c.expected, isSep(c.b); e != a {
			t.Errorf("%d: expected %t, but received %t", i+0, e, a)
		}
	}
}

func TestNewSep(t *testing.T) {
	cases := []struct {
		b             []byte
		expectedRead  int
		expectedError bool
		expectedToken sepToken
	}{
		{
			b:            []byte("["),
			expectedRead: 1,
			expectedToken: sepToken{
				value:   "[",
				sepType: sepTypeOpenBrace,
			},
		},
		{
			b:            []byte("]"),
			expectedRead: 1,
			expectedToken: sepToken{
				value:   "]",
				sepType: sepTypeCloseBrace,
			},
		},
	}

	for i, c := range cases {
		tok, n, err := newSepToken(c.b)

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
