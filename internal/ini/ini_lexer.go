package ini

import (
	"bytes"
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
	Raw() []rune
}

// Tokenize will return a list of tokens during lexical analysis of the
// io.Reader.
// TODO: Change to use runes instead of bytes
func (l *iniLexer) Tokenize(r io.Reader) ([]Token, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, awserr.New(ErrCodeUnableToReadFile, "unable to read file", err)
	}

	return l.tokenize(b)
}

func (l *iniLexer) tokenize(b []byte) ([]Token, error) {
	runes := bytes.Runes(b)
	var tok Token
	var err error
	n := 0
	tokens := make([]Token, 0, len(runes))

	for len(runes) > 0 {
		switch {
		case isWhitespace(runes[0]):
			tok, n, err = newWSToken(runes)
		case isComma(runes[0]):
			tok, n = newCommaToken(), 1
		case isComment(runes):
			tok, n, err = newCommentToken(runes)
		case isNewline(runes):
			tok, n, err = newNewlineToken(runes)
		case isSep(runes):
			tok, n, err = newSepToken(runes)
		case isOp(runes):
			tok, n, err = newOpToken(runes)
		default:
			tok, n, err = newLitToken(runes)
		}

		if err != nil {
			return nil, err
		}

		tokens = append(tokens, tok)

		runes = runes[n:]
	}

	return tokens, nil
}
