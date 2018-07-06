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

// isNumberValue will return whether not the leading characters in
// a byte slice is a number. A number is delimited by whitespace or
// the newline token.
//
// A number is defined to be in a binary, octal, decimal (int | float), hex format,
// or in scientific notation.
func isNumberValue(b []byte) bool {
	negativeIndex := 0
	helper := numberHelper{}

	for i := 0; i < len(b); i++ {
		negativeIndex++

		switch b[i] {
		case '-':
			if helper.IsNegative() || negativeIndex != 1 {
				return false
			}
			helper.Determine(b[i])
			continue
		case 'e', 'E':
			if helper.Exists() {
				return false
			}

			negativeIndex = 0
			helper.Determine(b[i])
			continue
		case 'b':
			if helper.hex {
				break
			}
			fallthrough
		case 'o', 'x':
			if i == 0 {
				return false
			}

			fallthrough
		case '.':
			if helper.Exists() {
				return false
			}
			helper.Determine(b[i])
			continue
		}

		if isWhitespace(b[i]) ||
			isNewline(b[i:]) {
			return true
		}

		if !helper.CorrectByte(b[i]) {
			return false
		}
	}

	return true
}

func isValid(b byte) bool {
	return utf8.ValidRune(rune(b)) && b != '=' && b != '[' && b != ']' && b != ' ' && b != '\n'
}

// UnionValueType is an enum that will signify what type
// the UnionValue is
type UnionValueType int

func (v UnionValueType) String() string {
	switch v {
	case NoneType:
		return "NONE"
	case DecimalType:
		return "FLOAT"
	case IntegerType:
		return "INT"
	case StringType:
		return "STRING"
	case BoolType:
		return "BOOL"
	}

	return ""
}

// UnionValueType enums
const (
	NoneType = UnionValueType(iota)
	DecimalType
	IntegerType
	StringType
	QuotedStringType
	BoolType
)

type literalToken struct {
	Value UnionValue
	raw   string
}

// UnionValue is a union container
type UnionValue struct {
	Type UnionValueType

	integer int64
	decimal float64
	boolean bool
	str     string
}

// Append will append values and change the type to a string
// type.
func (v *UnionValue) Append(tok Token) {
	if v.Type != QuotedStringType {
		v.Type = StringType
	}
	v.str += tok.StringValue()
}

func (v UnionValue) String() string {
	switch v.Type {
	case DecimalType:
		return fmt.Sprintf("decimal: %f", v.decimal)
	case IntegerType:
		return fmt.Sprintf("integer: %d", v.integer)
	case StringType:
		return fmt.Sprintf("string: %s", v.str)
	case QuotedStringType:
		return fmt.Sprintf("quoted string: %s", v.str)
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
		var base int
		value, base, n, err = getNumericalValue(b)
		if err != nil {
			return token, 0, err
		}

		token.raw = value
		if strings.Contains(value, ".") || hasExponent(value) {
			token.Value.Type = DecimalType
			token.Value.decimal, err = strconv.ParseFloat(value, 64)
		} else {
			if base != 10 {
				// strip off 0b, 0o, or 0x so strconv.ParseInt can
				// parse the value.
				value = value[2:]
			}
			token.Value.Type = IntegerType
			token.Value.integer, err = strconv.ParseInt(value, base, 64)
		}
	} else if isBoolValue(b) {
		value, n, err = getBoolValue(b)

		token.raw = value
		token.Value.Type = BoolType
		token.Value.boolean = value == "true"
	} else if b[0] == '"' {
		value, n, err = getStringValue(b)

		token.raw = value
		token.Value.Type = QuotedStringType
		token.Value.str = value
	} else {
		value, n, err = getValue(b)

		token.raw = value
		token.Value.Type = StringType
		token.Value.str = value
	}

	return token, n, err
}

func (token literalToken) IntValue() int64 {
	return token.Value.integer
}

func (token literalToken) FloatValue() float64 {
	return token.Value.decimal
}

func (token literalToken) BoolValue() bool {
	return token.Value.boolean
}

func isTrimmable(r rune) bool {
	switch r {
	case '\n', ' ':
		return true
	}
	return false
}

func (token literalToken) StringValue() string {
	switch token.Value.Type {
	case StringType:
		return strings.TrimFunc(token.Value.str, isTrimmable)
	case QuotedStringType:
		// preserve all characters in the quotes
		return token.Value.str
	default:
		return strings.TrimFunc(token.Raw(), isTrimmable)
	}
}

func (token literalToken) Raw() string {
	return token.raw
}

func (token literalToken) Type() TokenType {
	return TokenLit
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
