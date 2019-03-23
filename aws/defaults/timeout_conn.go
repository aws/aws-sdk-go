package defaults

import (
	"context"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

type Timeouts struct {
	Connect       time.Duration
	ConnKeepAlive time.Duration
	Read          time.Duration
	Write         time.Duration

	TLSHandshake   time.Duration
	ExpectContinue time.Duration
	IdleConn       time.Duration
	ResponseHeader time.Duration
}

func DefaultTimeouts() Timeouts {
	return Timeouts{
		Connect:       30 * time.Second,
		ConnKeepAlive: 30 * time.Second,
		Read:          30 * time.Second,
		Write:         30 * time.Second,

		TLSHandshake:   10 * time.Second,
		ExpectContinue: 1 * time.Second,

		IdleConn: 90 * time.Second,
	}
}

func NewClient(timeouts Timeouts) *http.Client {
	dialer := &net.Dialer{
		// Connect timeout.
		Timeout: timeouts.Connect,
		// protocol keep alive, e.g. TCP keep-alive
		KeepAlive: timeouts.ConnKeepAlive,
	}
	dialContextFn := dialer.DialContext

	// Don't use the timeout wrapper if not needed
	if timeouts.Read > 0 || timeouts.Write > 0 {
		wrapper := &dialTimeoutConnWrapper{
			Dialer:       dialer,
			ReadTimeout:  timeouts.Read,
			WriteTimeout: timeouts.Write,
		}
		dialContextFn = wrapper.DialContext
	}

	tr := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialContextFn,
		MaxIdleConns:          100,
		IdleConnTimeout:       timeouts.IdleConn,
		TLSHandshakeTimeout:   timeouts.TLSHandshake,
		ExpectContinueTimeout: timeouts.ExpectContinue,
		ResponseHeaderTimeout: timeouts.ResponseHeader,
	}

	return &http.Client{
		Transport: tr,
	}
}

type dialTimeoutConnWrapper struct {
	*net.Dialer

	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (d *dialTimeoutConnWrapper) Dial(network, address string) (net.Conn, error) {
	return d.DialContext(context.Background(), network, address)
}

func (d *dialTimeoutConnWrapper) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	conn, err := d.Dialer.DialContext(ctx, network, address)
	if err != nil {
		return nil, err
	}

	return &timeoutConn{
		Conn:         conn,
		ReadTimeout:  d.ReadTimeout,
		WriteTimeout: d.WriteTimeout,
	}, nil
}

type timeoutConn struct {
	net.Conn

	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	readDeadlineSet  int32
	writeDeadlineSet int32
}

func (c *timeoutConn) Read(b []byte) (int, error) {
	if c.ReadTimeout > 0 && !atomic.CompareAndSwapInt32(&c.readDeadlineSet, 1, 0) {
		c.Conn.SetReadDeadline(time.Now().Add(c.ReadTimeout))
	}

	return c.Conn.Read(b)
}

func (c *timeoutConn) Write(b []byte) (int, error) {
	if c.WriteTimeout > 0 && !atomic.CompareAndSwapInt32(&c.writeDeadlineSet, 1, 0) {
		c.Conn.SetWriteDeadline(time.Now().Add(c.WriteTimeout))
	}

	return c.Conn.Write(b)
}

func (c *timeoutConn) SetDeadline(v time.Time) error {
	atomic.StoreInt32(&c.writeDeadlineSet, 1)
	atomic.StoreInt32(&c.readDeadlineSet, 1)
	return c.Conn.SetDeadline(v)
}

func (c *timeoutConn) SetReadDeadline(v time.Time) error {
	atomic.StoreInt32(&c.readDeadlineSet, 1)
	return c.Conn.SetReadDeadline(v)
}

func (c *timeoutConn) SetWriteDeadline(v time.Time) error {
	atomic.StoreInt32(&c.writeDeadlineSet, 1)
	return c.Conn.SetWriteDeadline(v)
}
