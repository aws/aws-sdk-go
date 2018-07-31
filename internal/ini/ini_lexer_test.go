package ini

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	cases := []struct {
		r              io.Reader
		expectedTokens []Token
		expectedError  bool
	}{
		{
			r: bytes.NewBuffer([]byte(`x = 123`)),
			expectedTokens: []Token{
				literalToken{
					Value: Value{
						Type: StringType,
						raw:  []rune("x"),
					},
					raw: []rune("x"),
				},
				wsToken{
					raw: []rune(" "),
				},
				opToken{
					opType: opTypeEqual,
				},
				wsToken{
					raw: []rune(" "),
				},
				literalToken{
					Value: Value{
						Type:    IntegerType,
						integer: 123,
						raw:     []rune("123"),
					},
					raw: []rune("123"),
				},
			},
		},
		{
			r: bytes.NewBuffer([]byte(`[ foo ]`)),
			expectedTokens: []Token{
				sepToken{
					sepType: sepTypeOpenBrace,
				},
				wsToken{
					raw: []rune(" "),
				},
				literalToken{
					Value: Value{
						Type: StringType,
						raw:  []rune("foo"),
					},
					raw: []rune("foo"),
				},
				wsToken{
					raw: []rune(" "),
				},
				sepToken{
					sepType: sepTypeCloseBrace,
				},
			},
		},
	}

	for _, c := range cases {
		lex := iniLexer{}
		tokens, err := lex.Tokenize(c.r)

		if e, a := c.expectedError, err != nil; e != a {
			t.Errorf("expected %t, but received %t", e, a)
		}

		if e, a := c.expectedTokens, tokens; !reflect.DeepEqual(e, a) {
			t.Errorf("expected %v, but received %v", e, a)
		}
	}
}
