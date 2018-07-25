package ini

import (
	"fmt"
	"strings"
)

// getStringValue will return a quoted string and the amount
// of bytes read
//
// an error will be returned if the string is not properly formatted
func getStringValue(b []rune) (string, int, error) {
	if b[0] != '"' {
		return "", 0, NewParseError("strings must start with '\"'")
	}

	value := []rune{}
	endQuote := false
	i := 1

	for ; i < len(b) && !endQuote; i++ {
		if escaped := isEscaped(string(value), b[i]); b[i] == '"' && !escaped {
			endQuote = true
			break
		} else if escaped {
			value = value[:len(value)-1]
			c, err := getEscapedByte(b[i])
			if err != nil {
				return "", 0, err
			}

			value = append(value, c)
			continue
		}
		value = append(value, b[i])
	}

	if !endQuote {
		return "", 0, NewParseError("missing '\"' in string value")
	}

	return string(value), i + 1, nil
}

// getBoolValue will return a boolean and the amount
// of bytes read
//
// an error will be returned if the boolean is not of a correct
// value
func getBoolValue(b []rune) (string, int, error) {
	if len(b) < 4 {
		return "", 0, NewParseError("invalid boolean value")
	}

	value := ""
	n := 0
	for _, lv := range literalValues {
		if len(lv) > len(b) {
			continue
		}
		v := string(b[:len(lv)])
		if v == lv {
			value = v
			n = len(v)
		}
	}

	if n == 0 {
		return "", 0, NewParseError("invalid boolean value")
	}

	return value, n, nil
}

// getNumericalValue will return a numerical string, the amount
// of bytes read, and the base of the number
//
// an error will be returned if the number is not of a correct
// value
func getNumericalValue(b []rune) (string, int, int, error) {
	if !isDigit(b[0]) {
		return "", 0, 0, NewParseError("invalid digit value")
	}

	value := ""
	i := 0
	helper := numberHelper{}

loop:
	for negativeIndex := 0; i < len(b); i++ {
		negativeIndex++

		if !isDigit(b[i]) {
			switch b[i] {
			case '-':
				if helper.IsNegative() || negativeIndex != 1 {
					return "", 0, 0, NewParseError("parse error '-'")
				}

				neg, n := getNegativeNumber(b[i:])
				value += neg
				i += (n - 1)
				helper.Determine(b[i])
				continue
			case '.':
				if err := helper.Determine(b[i]); err != nil {
					return "", 0, 0, err
				}
			case 'e', 'E':
				if err := helper.Determine(b[i]); err != nil {
					return "", 0, 0, err
				}

				negativeIndex = 0
			case 'b':
				if helper.numberFormat == hex {
					break
				}
				fallthrough
			case 'o', 'x':
				if i == 0 && value != "0" {
					return "", 0, 0, NewParseError("incorrect base format, expected leading '0'")
				}

				if i != 1 {
					return "", 0, 0, NewParseError(fmt.Sprintf("incorrect base format found %s at %d index", string(b[i]), i))
				}

				if err := helper.Determine(b[i]); err != nil {
					return "", 0, 0, err
				}
			default:
				if i > 0 && isWhitespace(b[i]) {
					break loop
				}

				if isNewline(b[i:]) {
					break loop
				}

				if !(helper.numberFormat == hex && isHexByte(b[i])) {
					if i+2 < len(b) && !isNewline(b[i:i+2]) {
						return "", 0, 0, NewParseError("invalid numerical character")
					} else if !isNewline([]rune{b[i]}) {
						return "", 0, 0, NewParseError("invalid numerical character")
					}

					break loop
				}
			}
		}
		value += string(b[i])
	}

	return value, helper.Base(), i, nil
}

// isDigit will return whether or not something is an integer
func isDigit(b rune) bool {
	return b >= '0' && b <= '9'
}

func hasExponent(v string) bool {
	return strings.Contains(v, "e") || strings.Contains(v, "E")
}

func isBinaryByte(b rune) bool {
	switch b {
	case '0', '1':
		return true
	default:
		return false
	}
}

func isOctalByte(b rune) bool {
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7':
		return true
	default:
		return false
	}
}

func isHexByte(b rune) bool {
	if isDigit(b) {
		return true
	}
	return (b >= 'A' && b <= 'F') ||
		(b >= 'a' && b <= 'f')
}

func getValue(b []rune) (string, int, error) {
	value := ""
	i := 0

	for i < len(b) {
		if isWhitespace(b[i]) {
			break
		}

		if isOp(b[i:]) {
			break
		}

		valid, n, err := isValid(b[i:])
		if err != nil {
			return "", 0, err
		}

		if !valid {
			break
		}

		value += string(b[i : i+n])
		i += n
	}

	return value, i, nil
}

// getNegativeNumber will return a negative number from a
// byte slice. This will iterate through all characters until
// a non-digit has been found.
func getNegativeNumber(b []rune) (string, int) {
	if b[0] != '-' {
		return "", 0
	}
	value := string(b[0])
	for i := 1; i < len(b); i++ {
		if !isDigit(b[i]) {
			return value, len(value)
		}

		value += string(b[i])
	}

	return value, len(value)
}

// isEscaped will return whether or not the character is an escaped
// character.
func isEscaped(value string, b rune) bool {
	if len(value) == 0 {
		return false
	}

	switch b {
	case '\'': // single quote
	case '"': // quote
	case 'n': // newline
	case 't': // tab
	case '\\': // backslash
	default:
		return false
	}

	return value[len(value)-1] == '\\'
}

func getEscapedByte(b rune) (rune, error) {
	switch b {
	case '\'': // single quote
		return '\'', nil
	case '"': // quote
		return '"', nil
	case 'n': // newline
		return '\n', nil
	case 't': // table
		return '\t', nil
	case '\\': // backslash
		return '\\', nil
	default:
		return b, NewParseError(fmt.Sprintf("invalid escaped character %c", b))
	}
}
