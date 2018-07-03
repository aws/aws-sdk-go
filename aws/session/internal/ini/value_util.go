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
		// TODO: handle escaped strings, ie) "\"hello\""
		if b[i] == '"' {
			endQuote = true
			break
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
func getNumericalValue(b []byte) (string, int, error) {
	if !isDigit(b[0]) {
		return "", 0, awserr.New(ErrCodeParseError, "invalid digit value", nil)
	}

	value := ""
	i := 0

	foundDecimal := false
	foundBinary := false
	foundOctal := false
	foundHex := false

loop:
	for ; i < len(b); i++ {
		// TODO: handle scientific notation
		if !isDigit(b[i]) {
			switch b[i] {
			case '.':
				switch {
				case foundDecimal:
					return "", 0, awserr.New(ErrCodeParseError, "found multiple decimals", nil)
				case foundBinary, foundOctal, foundHex:
					return "", 0, awserr.New(ErrCodeParseError, "float value not valid", nil)
				}

				foundDecimal = true
			case 'b', 'o', 'x':
				if i == 0 {
					return "", 0, awserr.New(ErrCodeParseError, "incorrect base format", nil)
				}

				switch {
				case foundDecimal:
					return "", 0, awserr.New(ErrCodeParseError, "found decimal in binary, octal, hex format", nil)
				case foundBinary, foundOctal, foundHex:
					return "", 0, awserr.New(ErrCodeParseError, "multiple base formats", nil)
				}

				if b[i-1] != '0' {
					return "", 0, awserr.New(ErrCodeParseError, "incorrect base format", nil)
				}

				foundBinary = foundBinary || b[i] == 'b'
				foundOctal = foundOctal || b[i] == 'o'
				foundHex = foundHex || b[i] == 'x'
			default:
				if i > 0 && isWhitespace(b[i]) {
					break loop
				}

				if !(foundHex && isHexByte(b[i])) {
					if i+2 < len(b) && !isNewline(b[i:i+2]) {
						return "", 0, awserr.New(ErrCodeParseError, "invalid numerical character", nil)
					} else if !isNewline([]byte{b[i]}) {
						return "", 0, awserr.New(ErrCodeParseError, "invalid numerical character", nil)
					}

					break loop
				}
			}
		}
		value += string(b[i])
	}

	return value, i, nil
}

func isHexByte(b byte) bool {
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
