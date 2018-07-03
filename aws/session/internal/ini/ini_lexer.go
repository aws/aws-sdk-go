// package ini TODOs:
// [ ] support UTF-8
package ini

import (
	"io"
	"io/ioutil"
)

type tokenType int

func (t tokenType) String() string {
	switch t {
	case tokenNone:
		return "none"
	case tokenLit:
		return "literal"
	case tokenSep:
		return "sep"
	case tokenOp:
		return "op"
	case tokenWS:
		return "ws"
	case tokenNL:
		return "newline"
	case tokenComment:
		return "comment"
	default:
		return ""
	}
}

const (
	tokenNone = tokenType(iota)
	tokenLit
	tokenSep
	tokenOp
	tokenWS
	tokenNL
	tokenComment
)

type iniLexer struct{}

type iniToken interface {
	Type() tokenType

	Raw() string
	StringValue() string
	IntValue() int
	FloatValue() float64
	BoolValue() bool
}

func (l *iniLexer) Tokenize(r io.Reader) ([]iniToken, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var tok iniToken
	n := 0
	tokens := []iniToken{}
	i := 0

	for i < len(b) {
		subB := b[i:]

		switch {
		case isComment(subB):
			tok, n, err = newCommentToken(subB)
		case isWhitespace(subB[0]):
			tok, n, err = newWSToken(subB)
		case isNewline(subB):
			tok, n, err = newNewlineToken(subB)
		case isSep(subB):
			tok, n, err = newSepToken(subB)
		case isOp(subB):
			tok, n, err = newOpToken(subB)
		default:
			tok, n, err = newLitToken(subB)
		}

		if err != nil {
			return nil, err
		}

		tokens = append(tokens, tok)
		i += n
	}

	return tokens, nil
}
