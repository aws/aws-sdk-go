//go:build go1.7
// +build go1.7

package ini

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

const nested = `[default]
foo = ;comment
  bar = baz
  qux = wux
i=j`

func TestTokenize(t *testing.T) {
	cases := []struct {
		r              io.Reader
		expectedTokens []Token
		expectedError  bool
	}{
		{
			r: bytes.NewBuffer([]byte(nested)),
			expectedTokens: []Token{
				newToken(TokenSep, []rune("["), NoneType),
				newToken(TokenLit, []rune("default"), StringType),
				newToken(TokenSep, []rune("]"), NoneType),
				newToken(TokenNL, []rune("\n"), NoneType),
				newToken(TokenLit, []rune("foo"), StringType),
				newToken(TokenWS, []rune(" "), NoneType),
				newToken(TokenOp, []rune("="), NoneType),
				newToken(TokenWS, []rune(" "), NoneType),
				newToken(TokenComment, []rune(";comment"), NoneType),
				newToken(TokenNL, []rune("\n"), NoneType),
				newToken(TokenLit, []rune("  bar = baz\n  qux = wux"), StringType),
				newToken(TokenNL, []rune("\n"), NoneType),
				newToken(TokenLit, []rune("i"), StringType),
				newToken(TokenOp, []rune("="), NoneType),
				newToken(TokenLit, []rune("j"), StringType),
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
