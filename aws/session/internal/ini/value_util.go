package ini

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

// getStringValue will return a quoted string and the amount
// of bytes read
//
// an error will be returned if the string is not properly formatted
func getStringValue(b []byte) (string, int, error) {
	if b[0] != '"' {
		return "", 0, awserr.New(ErrCodeParseError, "strings must start with '\"'", nil)
	}

	value := ""
	endQuote := false
	i := 1
	for ; i < len(b) && !endQuote; i++ {
		if escaped := isEscaped(value, b[i]); b[i] == '"' && !escaped {
			endQuote = true
			break
		} else if escaped {
			value = value[:len(value)-1]
			c, err := getEscapedByte(b[i])
			if err != nil {
				return "", 0, err
			}

			value += string(c)
			continue
		}
		value += string(b[i])
	}

	if !endQuote {
		return "", 0, awserr.New(ErrCodeParseError, "missing '\"' in string value", nil)
	}

	return value, i + 1, nil
}

// getBoolValue will return a boolean and the amount
// of bytes read
//
// an error will be returned if the boolean is not of a correct
// value
func getBoolValue(b []byte) (string, int, error) {
	if len(b) < 4 {
		return "", 0, awserr.New(ErrCodeParseError, "invalid boolean value", nil)
	}

	value := ""
	n := 0
	for _, lv := range literalValues {
		v := string(b[:len(lv)])
		if v == lv {
			value = v
			n = len(v)
		}
	}

	if n == 0 {
		return "", 0, awserr.New(ErrCodeParseError, "invalid boolean value", nil)
	}

	return value, n, nil
}

// getNumericalValue will return a numerical string and the amount
// of bytes read
//
// an error will be returned if the number is not of a correct
// value
func getNumericalValue(b []byte) (string, int, int, error) {
	if !isDigit(b[0]) {
		return "", 0, 0, awserr.New(ErrCodeParseError, "invalid digit value", nil)
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
					return "", 0, 0, awserr.New(ErrCodeParseError, "parse error '-'", nil)
				}

				neg, n := getNegativeNumber(b[i:])
				value += neg
				i += (n - 1)
				helper.Determine(b[i])
				continue
			case '.':
				if helper.Exists() {
					return "", 0, 0, awserr.New(ErrCodeParseError, "invalid decimal format", nil)
				}

				helper.Determine(b[i])
			case 'e', 'E':
				if helper.Exists() {
					return "", 0, 0, awserr.New(ErrCodeParseError, fmt.Sprintf("multiple number formats found, %s", string(b[i])), nil)
				}

				negativeIndex = 0
				helper.Determine(b[i])
			case 'b':
				if helper.hex {
					break
				}
				fallthrough
			case 'o', 'x':
				if i == 0 && value != "0" {
					return "", 0, 0, awserr.New(ErrCodeParseError, "incorrect base format", nil)
				}

				if helper.Exists() {
					return "", 0, 0, awserr.New(ErrCodeParseError, "multiple base format", nil)
				}

				if b[i-1] != '0' {
					return "", 0, 0, awserr.New(ErrCodeParseError, "incorrect base format", nil)
				}

				helper.Determine(b[i])
			default:
				if i > 0 && isWhitespace(b[i]) {
					break loop
				}

				if isNewline(b[i:]) {
					break loop
				}

				if !(helper.hex && isHexByte(b[i])) {
					if i+2 < len(b) && !isNewline(b[i:i+2]) {
						return "", 0, 0, awserr.New(ErrCodeParseError, "invalid numerical character", nil)
					} else if !isNewline([]byte{b[i]}) {
						return "", 0, 0, awserr.New(ErrCodeParseError, "invalid numerical character", nil)
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
func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func hasExponent(v string) bool {
	return strings.Contains(v, "e") || strings.Contains(v, "E")
}

func isBinaryByte(b byte) bool {
	switch b {
	case '0', '1':
		return true
	default:
		return false
	}
}

func isOctalByte(b byte) bool {
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7':
		return true
	default:
		return false
	}
}

func isHexByte(b byte) bool {
	if isDigit(b) {
		return true
	}
	return (b >= 'A' && b <= 'F') ||
		(b >= 'a' && b <= 'f')
}

func getValue(b []byte) (string, int, error) {
	value := ""
	i := 0
	if !isValid(b[0]) {
		return "", 0, awserr.New(ErrCodeParseError, "invalid id name", nil)
	}

	for ; i < len(b); i++ {
		if isWhitespace(b[i]) {
			break
		}

		if !(isDigit(b[i]) || isValid(b[i])) {
			break
		}
		value += string(b[i])
	}

	return value, i, nil
}

// getNegativeNumber will return a negative number from a
// byte slice. This will iterate through all characters until
// a non-digit has been found.
func getNegativeNumber(b []byte) (string, int) {
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
func isEscaped(value string, b byte) bool {
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

func getEscapedByte(b byte) (byte, error) {
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
		return b, awserr.New(ErrCodeParseError, fmt.Sprintf("invalid escaped character %c", b), nil)
	}
}
