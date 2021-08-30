//go:build go1.13
// +build go1.13

package aws

import (
	"context"
	stderrors "errors"
	"io"
	"os"
	"testing"
)

type errCloser struct {
	err error
}

func (c errCloser) Close() error { return c.err }

func TestMultiCloser_Close(t *testing.T) {
	cases := []struct {
		name    string
		closers []io.Closer
		expects []error
	}{
		{
			name:    "no error",
			closers: []io.Closer{errCloser{err: nil}},
		},
		{
			name: "multiple errors",
			closers: []io.Closer{
				errCloser{err: os.ErrNotExist},
				errCloser{err: nil},
				errCloser{err: context.Canceled},
			},
			expects: []error{os.ErrNotExist, context.Canceled},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := MultiCloser(c.closers).Close()
			for _, e := range c.expects {
				if !stderrors.Is(err, e) {
					t.Errorf("expect %v is a %v", err, e)
				}
			}
		})
	}
}
