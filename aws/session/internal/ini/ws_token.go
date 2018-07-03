package ini

type wsToken struct {
	emptyToken
	raw string
}

// isWhitespace will return whether or not the character is
// a whitespace character
func isWhitespace(c byte) bool {
	return c == '\t' || c == ' '
}

func newWSToken(b []byte) (wsToken, int, error) {
	i := 0
	value := ""
	for ; i < len(b); i++ {
		if !isWhitespace(b[i]) {
			break
		}
		value += string(b[i])
	}

	return wsToken{
		raw: value,
	}, i, nil
}

func (tok wsToken) Raw() string {
	return tok.raw
}

func (tok wsToken) Type() tokenType {
	return tokenWS
}
