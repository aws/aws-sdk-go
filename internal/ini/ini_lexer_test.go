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
						str:  "x",
						raw:  "x",
					},
					raw: "x",
				},
				wsToken{
					raw: " ",
				},
				opToken{
					opType: opTypeEqual,
					value:  "=",
				},
				wsToken{
					raw: " ",
				},
				literalToken{
					Value: Value{
						Type:    IntegerType,
						integer: 123,
						raw:     "123",
					},
					raw: "123",
				},
			},
		},
		{
			r: bytes.NewBuffer([]byte(`[ foo ]`)),
			expectedTokens: []Token{
				sepToken{
					sepType: sepTypeOpenBrace,
					value:   "[",
				},
				wsToken{
					raw: " ",
				},
				literalToken{
					Value: Value{
						Type: StringType,
						str:  "foo",
						raw:  "foo",
					},
					raw: "foo",
				},
				wsToken{
					raw: " ",
				},
				sepToken{
					sepType: sepTypeCloseBrace,
					value:   "]",
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
