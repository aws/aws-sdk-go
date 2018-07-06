package ini

import (
	"reflect"
	"testing"
)

func TestSkipper(t *testing.T) {
	idTok, _, _ := newLitToken([]byte("id"))
	nlTok := newlineToken{}

	cases := []struct {
		name               string
		Fn                 func(s *skipper)
		param              Token
		expected           bool
		expectedShouldSkip bool
		expectedPrevTok    Token
	}{
		{
			name: "empty case",
			Fn: func(s *skipper) {
			},
			param: emptyToken{},
		},
		{
			name: "skip case",
			Fn: func(s *skipper) {
				s.Skip()
			},
			param:              idTok,
			expectedShouldSkip: true,
			expected:           true,
		},
		{
			name: "continue case",
			Fn: func(s *skipper) {
				s.Continue()
			},
			param: emptyToken{},
		},
		{
			name: "skip then continue case",
			Fn: func(s *skipper) {
				s.Skip()
				s.Continue()
			},
			param: emptyToken{},
		},
		{
			name: "do not skip case",
			Fn: func(s *skipper) {
				s.Skip()
				s.prevTok = nlTok
			},
			param:              idTok,
			expectedShouldSkip: true,
			expectedPrevTok:    nlTok,
		},
	}

	for _, c := range cases {
		s := skipper{}
		c.Fn(&s)

		if e, a := c.expectedShouldSkip, s.shouldSkip; e != a {
			t.Errorf("%s: expected %t, but received %t", c.name, e, a)
		}

		if e, a := c.expectedPrevTok, s.prevTok; !reflect.DeepEqual(e, a) {
			t.Errorf("%s: expected %v, but received %v", c.name, e, a)
		}

		if e, a := c.expected, s.ShouldSkip(c.param); e != a {
			t.Errorf("%s: expected %t, but received %t", c.name, e, a)
		}
	}
}
