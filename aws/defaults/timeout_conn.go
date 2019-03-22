package defaults

import (
	"context"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

type Timeouts struct {
	Connect              time.Duration
	NetProtocolKeepalive time.Duration
	Read                 time.Duration
	Write                time.Duration

	TLSHandshake   time.Duration
	ExpectContinue time.Duration
	IdleConn       time.Duration
	ResponseHeader time.Duration
}

func DefaultTimeouts() Timeouts {
	return Timeouts{
		Connect:              30 * time.Second,
		NetProtocolKeepalive: 30 * time.Second,
		Read:                 30 * time.Second,
		Write:                30 * time.Second,

		TLSHandshake:   10 * time.Second,
		ExpectContinue: 1 * time.Second,

		IdleConn: 90 * time.Second,
	}
}

func NewClient(timeouts Timeouts) *http.Client {
	dialer := dialTimeoutConnWrapper{
		Dialer: &net.Dialer{
			// Connect timeout.
			Timeout: timeouts.Connect,
			// protocol keep alive, e.g. TCP keep-alive
			KeepAlive: timeouts.NetProtocolKeepalive,
		},
		ReadTimeout:  timeouts.Read,
		WriteTimeout: timeouts.Write,
	}

	tr := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
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
	if !atomic.CompareAndSwapInt32(&c.readDeadlineSet, 1, 0) {
		c.Conn.SetReadDeadline(time.Now().Add(c.ReadTimeout))
	}

	n, err := c.Conn.Read(b)

	atomic.StoreInt32(&c.readDeadlineSet, 0)
	return n, err
}

func (c *timeoutConn) Write(b []byte) (int, error) {
	if !atomic.CompareAndSwapInt32(&c.writeDeadlineSet, 1, 0) {
		c.Conn.SetWriteDeadline(time.Now().Add(c.WriteTimeout))
	}

	atomic.StoreInt32(&c.writeDeadlineSet, 0)
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
