package ini

type newlineToken struct {
	emptyToken
	raw string
}

func isNewline(b []byte) bool {
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

func newNewlineToken(b []byte) (newlineToken, int, error) {
	value := string(b[0])
	if isNewline(b[1:]) {
		value += string(b[1])
	}

	return newlineToken{
		raw: value,
	}, len(value), nil
}

func (tok newlineToken) Raw() string {
	return tok.raw
}

func (tok newlineToken) Type() tokenType {
	return tokenNL
}
