package ini

// CommaToken represents a comma character token
type commaToken struct {
	emptyToken
}

func newCommaToken() commaToken {
	return commaToken{}
}

// Type will return the TokenType
func (tok commaToken) Type() TokenType {
	return TokenComma
}

func isComma(b rune) bool {
	return b == ','
}
