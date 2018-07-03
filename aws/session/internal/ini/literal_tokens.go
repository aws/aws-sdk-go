package ini

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

var literalValues = []string{
	"true",
	"false",
}

func isBoolValue(b []byte) bool {
	for _, lv := range literalValues {
		if len(b) < len(lv) {
			continue
		}

		v := string(b[:len(lv)])
		if v == lv {
			return true
		}
	}
	return false
}

func isNumberValue(b []byte) bool {
	foundDecimal := false
	foundBinary := false
	foundOctal := false
	foundHex := false

	for i := 0; i < len(b); i++ {
		switch b[i] {
		case '.':
			if foundDecimal ||
				foundBinary ||
				foundOctal ||
				foundHex {
				return false
			}
			foundDecimal = true
			continue
		case 'b', 'o', 'x':
			if i == 0 {
				return false
			}
			if foundDecimal ||
				foundBinary ||
				foundOctal ||
				foundHex {
				return false
			}
			foundBinary = foundBinary || b[i] == 'b'
			foundOctal = foundOctal || b[i] == 'o'
			foundHex = foundHex || b[i] == 'x'
			continue
		}

		if isWhitespace(b[i]) ||
			isNewline(b[i:]) {
			return true
		}

		switch {
		case foundBinary:
			if b[i] != '0' && b[i] != '1' {
				return false
			}
		case foundOctal:
			switch b[i] {
			case '0', '1', '2', '3', '4', '5', '6', '7':
			default:
				return false
			}
		case foundHex:
			if !isHexByte(b[i]) {
				return false
			}
		case foundDecimal:
			if !isDigit(b[i]) {
				return false
			}
		default:
			if !isDigit(b[i]) {
				return false
			}
		}

	}

	return true
}

// isDigit will return whether or not something is an integer
func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isValid(b byte) bool {
	return utf8.ValidRune(rune(b)) && b != '=' && b != '[' && b != ']' && b != ' ' && b != '\n'
}

const (
	NoneType = iota
	DecimalType
	IntegerType
	StringType
	BoolType
)

type literalToken struct {
	Value UnionValue
	raw   string
}

type UnionValue struct {
	Type int

	integer int
	decimal float64
	boolean bool
	str     string
}

func (v UnionValue) String() string {
	switch v.Type {
	case DecimalType:
		return fmt.Sprintf("decimal: %f", v.decimal)
	case IntegerType:
		return fmt.Sprintf("integer: %d", v.integer)
	case StringType:
		return fmt.Sprintf("string: %s", v.str)
	case BoolType:
		return fmt.Sprintf("bool: %t", v.boolean)
	default:
		return "union not set"
	}
}

func newLitToken(b []byte) (literalToken, int, error) {
	value := ""
	n := 0
	var err error

	token := literalToken{}

	if isNumberValue(b) {
		value, n, err = getNumericalValue(b)
		if err != nil {
			return token, 0, err
		}

		token.raw = value

		// TODO: octal, hex, and binary values
		// TODO: scientific notation
		if strings.Contains(value, ".") {
			token.Value.Type = DecimalType
			token.Value.decimal, err = strconv.ParseFloat(value, 64)
		} else {
			token.Value.Type = IntegerType
			token.Value.integer, err = strconv.Atoi(value)
		}
	} else if isBoolValue(b) {
		value, n, err = getBoolValue(b)

		token.raw = value
		token.Value.Type = BoolType
		token.Value.boolean = value == "true"
	} else if b[0] == '"' {
		value, n, err = getStringValue(b)

		token.raw = value
		token.Value.Type = StringType
		token.Value.str = value
	} else {
		value, n, err = getValue(b)

		token.raw = value
		token.Value.Type = StringType
		token.Value.str = value
	}

	return token, n, err
}

func (token literalToken) IntValue() int {
	return token.Value.integer
}

func (token literalToken) FloatValue() float64 {
	return token.Value.decimal
}

func (token literalToken) BoolValue() bool {
	return token.Value.boolean
}

func (token literalToken) StringValue() string {
	switch token.Value.Type {
	case StringType:
		return token.Value.str
	default:
		return token.Raw()
	}

	return ""
}

func (token literalToken) Raw() string {
	return token.raw
}

func (token literalToken) Type() tokenType {
	return tokenLit
}

func (token literalToken) String() string {
	switch token.Value.Type {
	case DecimalType:
		return fmt.Sprintf("%f", token.FloatValue())
	case IntegerType:
		return fmt.Sprintf("%d", token.IntValue())
	case StringType:
		return fmt.Sprintf("%s", token.StringValue())
	case BoolType:
		return fmt.Sprintf("%t", token.BoolValue())
	}

	return "invalid token"
}
