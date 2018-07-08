package ini

import (
	"io"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

const (
	// ErrCodeUnableToReadFile is used when a file is failed to be
	// opened or read from.
	ErrCodeUnableToReadFile = "FailedRead"
)

// TokenType represents the various different tokens types
type TokenType int

func (t TokenType) String() string {
	switch t {
	case TokenNone:
		return "none"
	case TokenLit:
		return "literal"
	case TokenSep:
		return "sep"
	case TokenOp:
		return "op"
	case TokenWS:
		return "ws"
	case TokenNL:
		return "newline"
	case TokenComment:
		return "comment"
	case TokenComma:
		return "comma"
	default:
		return ""
	}
}

// TokenType enums
const (
	TokenNone = TokenType(iota)
	TokenLit
	TokenSep
	TokenComma
	TokenOp
	TokenWS
	TokenNL
	TokenComment
)

type iniLexer struct{}

// Token is represents a token used in lexical analysis
type Token interface {
	Type() TokenType

	Raw() string
	StringValue() string
	IntValue() int64
	FloatValue() float64
	BoolValue() bool
}

// Tokenize will return a list of tokens during lexical analysis of the
// io.Reader.
func (l *iniLexer) Tokenize(r io.Reader) ([]Token, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, awserr.New(ErrCodeUnableToReadFile, "unable to read file", err)
	}

	var tok Token
	n := 0
	tokens := []Token{}
	i := 0

	for i < len(b) {
		subB := b[i:]

		switch {
		case isWhitespace(subB[0]):
			tok, n, err = newWSToken(subB)
		case isComma(subB[0]):
			tok, n = newCommaToken(), 1
		case isComment(subB):
			tok, n, err = newCommentToken(subB)
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
