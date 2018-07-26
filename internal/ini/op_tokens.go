package ini

import (
	"fmt"
)

func isOp(b []rune) bool {
	if len(b) == 0 {
		return false
	}

	switch b[0] {
	case '=':
		return true
	case ':':
		return true
	default:
		return false
	}
}

const (
	opTypeNone = iota
	opTypeEqual
)

// opToken is an operation token that signifies an expression.
type opToken struct {
	emptyToken

	opType int
	value  string
}

func newOpToken(b []rune) (opToken, int, error) {
	tok := opToken{}

	switch b[0] {
	case '=', ':':
		tok.opType = opTypeEqual
		tok.value = string(b[0])
	default:
		return tok, 0, NewParseError(fmt.Sprintf("unexpected op type, %v", b[0]))
	}
	return tok, 1, nil
}

func (token opToken) Raw() string {
	return token.value
}

func (token opToken) Type() TokenType {
	return TokenOp
}
