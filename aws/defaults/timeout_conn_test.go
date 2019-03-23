package defaults

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestTimeoutConn(t *testing.T) {

	cases := map[string]struct {
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
		ReadWait     time.Duration
		WriteWait    time.Duration
		ReadErr      bool
		WriteErr     bool
	}{
		"Successful": {
			ReadTimeout:  time.Millisecond,
			WriteTimeout: time.Millisecond,
		},
		"Read timeout": {
			ReadTimeout:  time.Millisecond,
			WriteTimeout: time.Millisecond,
			ReadWait:     10 * time.Second,
			ReadErr:      true,
		},
		"Write timeout": {
			ReadTimeout:  time.Millisecond,
			WriteTimeout: time.Millisecond,
			WriteWait:    10 * time.Second,
			WriteErr:     true,
		},
		"Success no timeout": {
			ReadWait:  10 * time.Second,
			WriteWait: 10 * time.Second,
		},
	}

	buf := make([]byte, 10)
	for nc, c := range cases {
		t.Run(nc, func(t *testing.T) {
			mc := &mockConn{
				readWait:  c.ReadWait,
				writeWait: c.WriteWait,
			}
			conn := &timeoutConn{
				Conn:         mc,
				ReadTimeout:  c.ReadTimeout,
				WriteTimeout: c.WriteTimeout,
			}

			_, err := conn.Read(buf)
			if e, a := c.ReadErr, (err != nil); e != a {
				t.Errorf("expect read error %t, got %t", e, a)
			}

			_, err = conn.Write(buf)
			if e, a := c.WriteErr, (err != nil); e != a {
				t.Errorf("expect write error %t, got %t", e, a)
			}
		})
	}
}

type mockConn struct {
	net.Conn

	readWait  time.Duration
	writeWait time.Duration

	readDeadline  time.Time
	writeDeadline time.Time
}

func (c *mockConn) Read(b []byte) (int, error) {
	if c.readDeadline.IsZero() {
		return len(b), nil
	}

	if time.Now().Add(c.readWait).After(c.readDeadline) {
		return 0, fmt.Errorf("read deadline exceeded")
	}

	return len(b), nil
}

func (c *mockConn) Write(b []byte) (int, error) {
	if c.writeDeadline.IsZero() {
		return len(b), nil
	}

	if time.Now().Add(c.writeWait).After(c.writeDeadline) {
		return 0, fmt.Errorf("write deadline exceeded")
	}

	return len(b), nil
}

func (c *mockConn) SetDeadline(v time.Time) error {
	c.readDeadline = v
	c.writeDeadline = v
	return nil
}

func (c *mockConn) SetReadDeadline(v time.Time) error {
	c.readDeadline = v
	return nil
}

func (c *mockConn) SetWriteDeadline(v time.Time) error {
	c.writeDeadline = v
	return nil
}
