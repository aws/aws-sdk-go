package ini

type CommaToken struct {
	emptyToken
}

func newCommaToken() CommaToken {
	return CommaToken{}
}

func (tok CommaToken) Type() tokenType {
	return tokenComma
}

func isComma(b byte) bool {
	return b == ','
}
