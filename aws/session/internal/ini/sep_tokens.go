package ini

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

func isSep(b []byte) bool {
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

type sepToken struct {
	emptyToken

	sepType int
	value   string
}

func newSepToken(b []byte) (sepToken, int, error) {
	tok := sepToken{}

	switch b[0] {
	case '[':
		tok.sepType = sepTypeOpenBrace
		tok.value = string(b[0])
	case ']':
		tok.sepType = sepTypeCloseBrace
		tok.value = string(b[0])
	default:
		return tok, 0, awserr.New(ErrCodeParseError, fmt.Sprintf("unexpected sep type, %v", b[0]), nil)
	}
	return tok, 1, nil
}

func (token sepToken) StringValue() string {
	return token.value
}

func (token sepToken) Raw() string {
	return token.value
}

func (token sepToken) Type() tokenType {
	return tokenSep
}
