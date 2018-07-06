package ini

// emptyToken is used to satisfy the Token interface
type emptyToken struct{}

func (token emptyToken) Type() TokenType {
	return TokenNone
}

func (token emptyToken) StringValue() string {
	return ""
}

func (token emptyToken) Raw() string {
	return ""
}

func (token emptyToken) IntValue() int64 {
	return 0
}

func (token emptyToken) FloatValue() float64 {
	return 0.0
}

func (token emptyToken) BoolValue() bool {
	return false
}
