package ini

import (
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
			value += string(getEscapedCharacter(b[i]))
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
	base := 10

	foundDecimal := false
	foundBinary := false
	foundOctal := false
	foundHex := false
	foundExponent := false
	// negativate variable is mostly to indicate whether or not
	// a '-' character is in a valid area
	foundNegative := false

loop:
	for negativeIndex := 0; i < len(b); i++ {
		negativeIndex++

		if !isDigit(b[i]) {
			switch b[i] {
			case '-':
				if foundNegative || negativeIndex != 1 {
					return "", 0, 0, awserr.New(ErrCodeParseError, "parse error '-'", nil)
				}

				neg, n := getNegativeNumber(b[i:])
				value += neg
				i += n
				continue
			case '.':
				switch {
				case foundDecimal:
					return "", 0, 0, awserr.New(ErrCodeParseError, "found multiple decimals", nil)
				case foundBinary, foundOctal, foundHex:
					return "", 0, 0, awserr.New(ErrCodeParseError, "float value not valid", nil)
				}

				foundDecimal = true
			case 'e', 'E':
				switch {
				case foundDecimal:
					return "", 0, 0, awserr.New(ErrCodeParseError, "exponent value not valid", nil)
				case foundBinary:
					return "", 0, 0, awserr.New(ErrCodeParseError, "exponent value not valid", nil)
				case foundOctal:
					return "", 0, 0, awserr.New(ErrCodeParseError, "exponent value not valid", nil)
				case foundHex:
					return "", 0, 0, awserr.New(ErrCodeParseError, "exponent value not valid", nil)
				case foundExponent:
					return "", 0, 0, awserr.New(ErrCodeParseError, "exponent value not valid", nil)
				}

				foundExponent = true
				foundNegative = false
				negativeIndex = 0
			case 'b', 'o', 'x':
				if i == 0 && value != "0" {
					return "", 0, 0, awserr.New(ErrCodeParseError, "incorrect base format", nil)
				}

				switch {
				case foundDecimal:
					return "", 0, 0, awserr.New(ErrCodeParseError, "found decimal in binary, octal, or hex format", nil)
				case foundBinary, foundOctal, foundHex:
					return "", 0, 0, awserr.New(ErrCodeParseError, "multiple base formats", nil)
				case foundExponent:
					return "", 0, 0, awserr.New(ErrCodeParseError, "found exponent in bainry, octal, or hex format", nil)
				}

				if b[i-1] != '0' {
					return "", 0, 0, awserr.New(ErrCodeParseError, "incorrect base format", nil)
				}

				if foundBinary = foundBinary || b[i] == 'b'; foundBinary {
					base = 2
				}
				if foundOctal = foundOctal || b[i] == 'o'; foundOctal {
					base = 8
				}
				if foundHex = foundHex || b[i] == 'x'; foundHex {
					base = 16
				}
			default:
				if i > 0 && isWhitespace(b[i]) {
					break loop
				}

				if !(foundHex && isHexByte(b[i])) {
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

	return value, base, i, nil
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

func isEscaped(value string, b byte) bool {
	if len(value) == 0 {
		return false
	}

	switch b {
	case '"': // quote
	case 'n': // newline
	case 't': // table
	default:
		return false
	}

	return value[len(value)-1] == '\\'
}

func getEscapedCharacter(b byte) byte {
	switch b {
	case '\'':
		return '\''
	case '"': // quote
		return '"'
	case 'n': // newline
		return '\n'
	case 't': // table
		return '\t'
	case '\\':
		return '\\'
	default:
		return '\\'
	}
}
