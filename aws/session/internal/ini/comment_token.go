package ini

type CommentToken struct {
	emptyToken
	comment string
}

func isComment(b []byte) bool {
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

func newCommentToken(b []byte) (CommentToken, int, error) {
	i := 0
	value := ""
	for ; i < len(b); i++ {
		if b[i] == '\n' {
			break
		}
		if b[i] == '\r' && b[i+1] == '\n' {
			break
		}
		value += string(b[i])
	}

	return CommentToken{
		comment: value,
	}, i, nil
}

func (token CommentToken) Raw() string {
	return token.comment
}

func (token CommentToken) StringValue() string {
	return token.comment
}

func (token CommentToken) Type() tokenType {
	return tokenComment
}
