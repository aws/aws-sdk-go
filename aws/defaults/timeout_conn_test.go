package defaults

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestTimeoutConn(t *testing.T) {
	mc := &mockConn{}
	conn := &timeoutConn{
		Conn:         mc,
		ReadTimeout:  time.Millisecond,
		WriteTimeout: time.Millisecond,
	}
	buf := make([]byte, 10)

	cases := map[string]struct {
		ReadWait  time.Duration
		WriteWait time.Duration
		ReadErr   bool
		WriteErr  bool
	}{
		"Successful": {},
		"Read timeout": {
			ReadWait: 10 * time.Second,
			ReadErr:  true,
		},
		"Write timeout": {
			WriteWait: 10 * time.Second,
			WriteErr:  true,
		},
	}

	for nc, c := range cases {
		t.Run(nc, func(t *testing.T) {
			mc.readWait = c.ReadWait
			mc.writeWait = c.WriteWait

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
	if time.Now().Add(c.readWait).After(c.readDeadline) {
		return 0, fmt.Errorf("read deadline exceeded")
	}
	return len(b), nil
}

func (c *mockConn) Write(b []byte) (int, error) {
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
