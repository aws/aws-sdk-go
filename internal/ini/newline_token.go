package ini

// newlineToken acts as a delimeter in ini is will be used
// primarily to handle nesting expressions.
type newlineToken struct {
	emptyToken
	raw []rune
}

func isNewline(b []rune) bool {
	if len(b) == 0 {
		return false
	}

	if b[0] == '\n' {
		return true
	}

	if len(b) < 2 {
		return false
	}

	return b[0] == '\r' && b[1] == '\n'
}

func newNewlineToken(b []rune) (newlineToken, int, error) {
	i := 1
	if b[0] == '\r' && isNewline(b[1:]) {
		i++
	}

	if !isNewline([]rune(b[:i])) {
		return newlineToken{}, 0, NewParseError("invalid new line token")
	}

	return newlineToken{
		raw: b[:i],
	}, i, nil
}

func (tok newlineToken) Raw() []rune {
	return tok.raw
}

func (tok newlineToken) Type() TokenType {
	return TokenNL
}
