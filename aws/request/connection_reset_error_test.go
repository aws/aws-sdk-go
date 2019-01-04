// +build !appengine,!plan9

package request_test

import (
	"net"
	"syscall"
)

var (
	stubConnectionResetErrorAccept = &net.OpError{Op: "accept", Err: syscall.ECONNRESET}
	stubConnectionResetErrorRead   = &net.OpError{Op: "read", Err: syscall.ECONNRESET}
)
