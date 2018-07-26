package ini

// emptyToken is used to satisfy the Token interface
type emptyToken struct{}

func (token emptyToken) Type() TokenType {
	return TokenNone
}

func (token emptyToken) Raw() string {
	return ""
}
