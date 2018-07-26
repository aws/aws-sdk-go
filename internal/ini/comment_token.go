package ini

// commentToken represents a token in an ini configuration.
// Comments may start with a ';', '#', or "//".
type commentToken struct {
	emptyToken
	comment string
}

// isComment will return whether or not the next byte(s) is a
// comment.
func isComment(b []rune) bool {
	if len(b) == 0 {
		return false
	}

	switch b[0] {
	case ';':
		return true
	case '#':
		return true
	case '/':
		if len(b) > 1 {
			return b[1] == '/'
		}
	}

	return false
}

// newCommentToken will create a comment token and
// return how many bytes were read.
func newCommentToken(b []rune) (commentToken, int, error) {
	i := 0
	value := ""
	for ; i < len(b); i++ {
		if b[i] == '\n' {
			break
		}

		if len(b) > 2 && b[i] == '\r' && b[i+1] == '\n' {
			break
		}
		value += string(b[i])
	}

	return commentToken{
		comment: value,
	}, i, nil
}

// Raw will return the raw value from the token
func (token commentToken) Raw() string {
	return token.comment
}

func (token commentToken) StringValue() string {
	return token.comment
}

// Type will return the TokenType
func (token commentToken) Type() TokenType {
	return TokenComment
}
