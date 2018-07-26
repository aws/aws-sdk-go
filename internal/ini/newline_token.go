package ini

// newlineToken acts as a delimeter in ini is will be used
// primarily to handle nesting expressions.
type newlineToken struct {
	emptyToken
	raw string
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
	value := string(b[0])
	if value[0] == '\r' && isNewline(b[1:]) {
		value += string(b[1])
	}

	if !isNewline([]rune(value)) {
		return newlineToken{}, 0, NewParseError("invalid new line token")
	}

	return newlineToken{
		raw: value,
	}, len(value), nil
}

func (tok newlineToken) Raw() string {
	return tok.raw
}

func (tok newlineToken) Type() TokenType {
	return TokenNL
}
