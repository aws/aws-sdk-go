package ini

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

func isOp(b []byte) bool {
	if len(b) == 0 {
		return false
	}

	switch b[0] {
	case '=':
		return true
	default:
		return false
	}
}

const (
	opTypeNone = iota
	opTypeEqual
)

type opToken struct {
	emptyToken

	opType int
	value  string
}

func newOpToken(b []byte) (opToken, int, error) {
	tok := opToken{}

	switch b[0] {
	case '=':
		tok.opType = opTypeEqual
		tok.value = string(b[0])
	default:
		return tok, 0, awserr.New(ErrCodeParseError, fmt.Sprintf("unexpected op type, %v", b[0]), nil)
	}
	return tok, 1, nil
}

func (token opToken) StringValue() string {
	return token.value
}

func (token opToken) Raw() string {
	return token.value
}

func (token opToken) Type() tokenType {
	return tokenOp
}

func (token opToken) String() string {
	return token.value
}
