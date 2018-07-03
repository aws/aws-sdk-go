package ini

// emptyToken is used to satisfy the iniToken interface
type emptyToken struct{}

func (token emptyToken) Type() tokenType {
	return tokenNone
}

func (token emptyToken) StringValue() string {
	return ""
}

func (token emptyToken) Raw() string {
	return ""
}

func (token emptyToken) IntValue() int {
	return 0
}

func (token emptyToken) FloatValue() float64 {
	return 0.0
}

func (token emptyToken) BoolValue() bool {
	return false
}
