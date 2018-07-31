package ini

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	runesTrue  = []rune("true")
	runesFalse = []rune("false")
)

var literalValues = [][]rune{
	runesTrue,
	runesFalse,
}

func isBoolValue(b []rune) bool {
	for _, lv := range literalValues {
		if isLitValue(lv, b) {
			return true
		}
	}
	return false
}

func isLitValue(want, have []rune) bool {
	if len(have) < len(want) {
		return false
	}

	for i := 0; i < len(want); i++ {
		if want[i] != have[i] {
			return false
		}
	}

	return true
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

		if i > 0 && (isNewline(b[i:]) || isWhitespace(b[i])) {
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
	raw   []rune
}

// Value is a union container
type Value struct {
	Type ValueType
	raw  []rune

	integer int64
	decimal float64
	boolean bool
}

// Append will append values and change the type to a string
// type.
func (v *Value) Append(tok Token) {
	if v.Type != QuotedStringType {
		v.Type = StringType
	}
	if litToken, ok := tok.(literalToken); ok {
		v.raw = append(v.raw, []rune(litToken.Value.StringValue())...)
	} else {
		v.raw = append(v.raw, tok.Raw()...)
	}
}

func (v Value) String() string {
	switch v.Type {
	case DecimalType:
		return fmt.Sprintf("decimal: %f", v.decimal)
	case IntegerType:
		return fmt.Sprintf("integer: %d", v.integer)
	case StringType:
		return fmt.Sprintf("string: %s", string(v.raw))
	case QuotedStringType:
		return fmt.Sprintf("quoted string: %s", string(v.raw))
	case BoolType:
		return fmt.Sprintf("bool: %t", v.boolean)
	default:
		return "union not set"
	}
}

func newLitToken(b []rune) (literalToken, int, error) {
	n := 0
	var err error

	token := literalToken{}

	if isNumberValue(b) {
		var base int
		base, n, err = getNumericalValue(b)
		if err != nil {
			return token, 0, err
		}

		token.raw = b[:n]
		token.Value.raw = token.raw
		value := token.raw
		if contains(value, '.') || hasExponent(value) {
			token.Value.Type = DecimalType
			// TODO: use buffer
			token.Value.decimal, err = strconv.ParseFloat(string(value), 64)
		} else {
			if base != 10 {
				// strip off 0b, 0o, or 0x so strconv.ParseInt can
				// parse the value.
				value = value[2:]
			}
			token.Value.Type = IntegerType
			token.Value.integer, err = strconv.ParseInt(string(value), base, 64)
		}
	} else if isBoolValue(b) {
		n, err = getBoolValue(b)

		token.raw = b[:n]
		token.Value.raw = token.raw
		token.Value.Type = BoolType
		token.Value.boolean = runeCompare(token.raw, runesTrue)
	} else if b[0] == '"' {
		n, err = getStringValue(b)
		if err != nil {
			return token, n, err
		}

		// remove quotes
		token.raw = b[1 : n-1]
		token.Value.raw = token.raw
		token.Value.Type = QuotedStringType
	} else {
		n, err = getValue(b)

		token.raw = b[:n]
		token.Value.raw = token.raw
		token.Value.Type = StringType
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
		return strings.TrimFunc(string(v.raw), isTrimmable)
	case QuotedStringType:
		// preserve all characters in the quotes
		return string(v.raw)
	default:
		return strings.TrimFunc(string(v.raw), isTrimmable)
	}
}

func (token literalToken) Raw() []rune {
	return []rune(token.raw)
}

func (token literalToken) Type() TokenType {
	return TokenLit
}

func (token literalToken) String() string {
	return token.Value.String()
}

func contains(runes []rune, c rune) bool {
	for i := 0; i < len(runes); i++ {
		if runes[i] == c {
			return true
		}
	}

	return false
}

func runeCompare(v1 []rune, v2 []rune) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := 0; i < len(v1); i++ {
		if v1[i] != v2[i] {
			return false
		}
	}

	return true
}
