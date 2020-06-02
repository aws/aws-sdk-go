// +build go1.13

package awserr

import (
	"errors"
	"fmt"
	"testing"
)

func TestBaseErrorSupportsUnwrap(t *testing.T) {
	wrapped := fmt.Errorf("snap!")
	cases := []struct {
		input   error
		unwraps bool
	}{
		{
			input:   New("NotGood", "not so good", nil),
			unwraps: false,
		},
		{
			input:   New("Bad", "a bit bad", wrapped),
			unwraps: true,
		},
		{
			input:   New("Worse", "somewhat bad", New("Bad", "a bit bad", wrapped)),
			unwraps: true,
		},
		{
			input:   New("Worst", "so bad", New("Worse", "somewhat bad", New("Bad", "a bit bad", wrapped))),
			unwraps: true,
		},
		{
			input:   New("VeryWorst", "so very bad", New("Worse", "somewhat bad", New("Bad", "a bit bad", nil))),
			unwraps: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.input.Error(), func(t *testing.T) {
			_, unwrappable := tc.input.(interface {
				Unwrap() error
			})
			if !unwrappable {
				t.Errorf("%q does not implement Unwrap method", tc.input)
			}
			if unwrapped := errors.Is(tc.input, wrapped); unwrapped != tc.unwraps {
				t.Errorf("expected errors.Is(...)==%v but got: %v", tc.unwraps, unwrapped)
			}
		})
	}
}

func TestBatchedErrorsDontUnwrap(t *testing.T) {
	wrapped := fmt.Errorf("snap!")
	for _, err := range []error{
		NewBatchError("Code1", "one", nil),
		NewBatchError("Code2", "two", []error{wrapped}),
		NewBatchError("Code3", "three", []error{wrapped, wrapped}),
	} {
		_, unwrappable := err.(interface {
			Unwrap() error
		})
		if unwrappable {
			t.Errorf("%q unexpectedly implements Unwrap method", err)
		}
		if errors.Is(err, wrapped) {
			t.Errorf("%q unexpectedly has the wrapped error in its chain", err)
		}
	}
}
