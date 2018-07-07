package ini

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

// numberHelper is used to dictate what format a number is in
// and what to do for negative values. Since -1e-4 is a valid
// number, we cannot just simply check for duplicate negatives.
type numberHelper struct {
	binary  bool
	octal   bool
	decimal bool
	hex     bool

	exponent bool
	negative bool
}

func (b numberHelper) Exists() bool {
	return b.decimal || b.binary || b.octal || b.hex || b.exponent
}

func (b numberHelper) IsNegative() bool {
	return b.negative
}

func (b *numberHelper) Determine(c byte) error {
	if b.Exists() {
		return awserr.New(ErrCodeParseError, fmt.Sprintf("multiple number formats: 0%v", string(c)), nil)
	}

	switch c {
	case 'b':
		b.binary = true
	case 'o':
		b.octal = true
	case 'x':
		b.hex = true
	case 'e', 'E':
		b.exponent = true
		b.negative = false
	case '-':
		b.negative = true
	case '.':
		b.decimal = true
	default:
		return awserr.New(ErrCodeParseError, fmt.Sprintf("invalid number character: %v", string(c)), nil)
	}

	return nil
}

func (b numberHelper) CorrectByte(c byte) bool {
	switch {
	case b.binary:
		if !isBinaryByte(c) {
			return false
		}
	case b.octal:
		if !isOctalByte(c) {
			return false
		}
	case b.hex:
		if !isHexByte(c) {
			return false
		}
	case b.decimal:
		if !isDigit(c) {
			return false
		}
	case b.exponent:
		if !isDigit(c) {
			return false
		}
	case b.negative:
		if !isDigit(c) {
			return false
		}
	default:
		if !isDigit(c) {
			return false
		}
	}

	return true
}

func (b numberHelper) Base() int {
	switch {
	case b.binary:
		return 2
	case b.octal:
		return 8
	case b.hex:
		return 16
	default:
		return 10
	}
}

func (b numberHelper) String() string {
	buf := bytes.Buffer{}
	i := 0

	if b.binary {
		i++
		buf.WriteString(string(i+'0') + ": binary format\n")
	}

	if b.octal {
		i++
		buf.WriteString(string(i+'0') + ": octal format\n")
	}

	if b.hex {
		i++
		buf.WriteString(string(i+'0') + ": hex format\n")
	}

	if b.exponent {
		i++
		buf.WriteString(string(i+'0') + ": exponent format\n")
	}

	if b.negative {
		i++
		buf.WriteString(string(i+'0') + ": negative format\n")
	}

	return buf.String()
}
