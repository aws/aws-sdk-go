package ini

import (
	"unicode"
)

type wsToken struct {
	emptyToken
	raw string
}

// isWhitespace will return whether or not the character is
// a whitespace character.
//
// Whitespace is defined as a space or tab.
func isWhitespace(c rune) bool {
	return unicode.IsSpace(c) && c != '\n' && c != '\r'
}

func newWSToken(b []rune) (wsToken, int, error) {
	i := 0
	value := ""
	for ; i < len(b); i++ {
		if !isWhitespace(b[i]) {
			break
		}
		value += string(b[i])
	}

	return wsToken{
		raw: value,
	}, i, nil
}

func (tok wsToken) Raw() string {
	return tok.raw
}

func (tok wsToken) StringValue() string {
	return tok.raw
}

func (tok wsToken) Type() TokenType {
	return TokenWS
}
