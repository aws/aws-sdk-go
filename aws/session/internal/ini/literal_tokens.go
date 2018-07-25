package ini

import (
	"fmt"
	"strconv"
	"strings"
)

var literalValues = []string{
	"true",
	"false",
}

func isBoolValue(b []rune) bool {
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
func isNumberValue(b []rune) bool {
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
			if err := helper.Determine(b[i]); err != nil {
				return false
			}
			negativeIndex = 0
			continue
		case 'b':
			if helper.numberFormat == hex {
				break
			}
			fallthrough
		case 'o', 'x':
			if i == 0 {
				return false
			}

			fallthrough
		case '.':
			if err := helper.Determine(b[i]); err != nil {
				return false
			}
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

func isValid(b []rune) (bool, int, error) {
	if len(b) == 0 {
		// TODO: should probably return an error
		return false, 0, nil
	}

	return isValidRune(b[0]), 1, nil
}

func isValidRune(r rune) bool {
	return r != '=' && r != '[' && r != ']' && r != ' ' && r != '\n'
}

// ValueType is an enum that will signify what type
// the Value is
type ValueType int

func (v ValueType) String() string {
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

// ValueType enums
const (
	NoneType = ValueType(iota)
	DecimalType
	IntegerType
	StringType
	QuotedStringType
	BoolType
)

type literalToken struct {
	Value Value
	raw   string
}

// Value is a union container
type Value struct {
	Type ValueType
	raw  string

	integer int64
	decimal float64
	boolean bool
	str     string
}

// Append will append values and change the type to a string
// type.
func (v *Value) Append(tok Token) {
	if v.Type != QuotedStringType {
		v.Type = StringType
	}
	if litToken, ok := tok.(literalToken); ok {
		v.str += litToken.Value.StringValue()
	} else {
		v.str += tok.Raw()
	}
}

func (v Value) String() string {
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

func newLitToken(b []rune) (literalToken, int, error) {
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
		token.Value.raw = value
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
		token.Value.raw = value
		token.Value.Type = BoolType
		token.Value.boolean = value == "true"
	} else if b[0] == '"' {
		value, n, err = getStringValue(b)

		token.raw = value
		token.Value.raw = value
		token.Value.Type = QuotedStringType
		token.Value.str = value
	} else {
		value, n, err = getValue(b)

		token.raw = value
		token.Value.raw = value
		token.Value.Type = StringType
		token.Value.str = value
	}

	return token, n, err
}

func (v Value) IntValue() int64 {
	return v.integer
}

func (v Value) FloatValue() float64 {
	return v.decimal
}

func (v Value) BoolValue() bool {
	return v.boolean
}

func isTrimmable(r rune) bool {
	switch r {
	case '\n', ' ':
		return true
	}
	return false
}

func (v Value) StringValue() string {
	switch v.Type {
	case StringType:
		return strings.TrimFunc(v.str, isTrimmable)
	case QuotedStringType:
		// preserve all characters in the quotes
		return v.str
	default:
		return strings.TrimFunc(v.raw, isTrimmable)
	}
}

func (token literalToken) Raw() string {
	return token.raw
}

func (token literalToken) Type() TokenType {
	return TokenLit
}

func (token literalToken) String() string {
	return token.Value.String()
}
