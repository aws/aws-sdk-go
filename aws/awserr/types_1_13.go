//go:build go1.13
// +build go1.13

package awserr

import (
	"errors"
)

// Is reports whether this error matches target.
func (b baseError) Is(target error) bool {
	for _, e := range b.errs {
		if errors.Is(e, target) {
			return true
		}
	}
	return false
}

// As finds the first error in errs that matches target, and if so,
// sets target to that error value and returns true. Otherwise, it returns false.
func (b baseError) As(target interface{}) bool {
	for _, e := range b.errs {
		if errors.As(e, target) {
			return true
		}
	}
	return false
}
