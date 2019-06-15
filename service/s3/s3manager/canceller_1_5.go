// +build !go1.9

package s3manager

import (
	"github.com/aws/aws-sdk-go/aws"
	"sync"
)

type cancellerFunc func()

type cancellerCtx struct {
	aws.Context
	mu     sync.Mutex
	closed bool
	ch     chan struct{}
}

func (c *cancellerCtx) Done() <-chan struct{} {
	return c.ch
}

// canceller mimics the interface of context.WithCancel for older versions
// of Go. It is a simplified implementation with no support for cancelling
// child contexts or attaching to the parent context. Consequently most of
// the benefits of cancellation are lost, but for the Copier's purposes it's
// a stable interface and correct enough.
func canceller(ctx aws.Context) (aws.Context, cancellerFunc) {
	cctx := cancellerCtx{Context: ctx, ch: make(chan struct{})}
	return &cctx, func() {
		cctx.mu.Lock()
		defer cctx.mu.Unlock()
		if !cctx.closed {
			close(cctx.ch)
			cctx.closed = true
		}
	}
}
