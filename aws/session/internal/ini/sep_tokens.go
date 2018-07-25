package ini

import (
	"fmt"
)

func isSep(b []rune) bool {
	if len(b) == 0 {
		return false
	}

	switch b[0] {
	case '[', ']':
		return true
	default:
		return false
	}
}

const (
	sepTypeNone = iota
	sepTypeOpenBrace
	sepTypeCloseBrace
)

// sepToken is a separator token which represents the concept of
// scoping in ini files.
type sepToken struct {
	emptyToken

	sepType int
	value   string
}

func newSepToken(b []rune) (sepToken, int, error) {
	tok := sepToken{}

	switch b[0] {
	case '[':
		tok.sepType = sepTypeOpenBrace
		tok.value = string(b[0])
	case ']':
		tok.sepType = sepTypeCloseBrace
		tok.value = string(b[0])
	default:
		return tok, 0, NewParseError(fmt.Sprintf("unexpected sep type, %v", b[0]))
	}
	return tok, 1, nil
}

func (token sepToken) Raw() string {
	return token.value
}

func (token sepToken) Type() TokenType {
	return TokenSep
}
