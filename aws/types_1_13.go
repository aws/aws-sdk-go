//go:build go1.13
// +build go1.13

package aws

import stderrors "errors"

// Is reports whether this error matches target.
func (es errors) Is(target error) bool {
	for _, e := range es {
		if stderrors.Is(e, target) {
			return true
		}
	}
	return false
}

// As finds the first error in errs that matches target, and if so,
// sets target to that error value and returns true. Otherwise, it returns false.
func (es errors) As(target interface{}) bool {
	for _, e := range es {
		if stderrors.As(e, target) {
			return true
		}
	}
	return false
}
