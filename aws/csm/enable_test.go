package csm

import (
	"net"
	"testing"
)

func startUDPServer(done chan struct{}, fn func([]byte)) (string, error) {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		return "", err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return "", err
	}

	buf := make([]byte, 1024)
	i := 0
	go func() {
		defer conn.Close()
		for {
			i++
			select {
			case <-done:
				return
			default:
			}

			n, _, err := conn.ReadFromUDP(buf)
			fn(buf[:n])

			if err != nil {
				panic(err)
			}
		}
	}()

	return conn.LocalAddr().String(), nil
}

func TestInvalidPort(t *testing.T) {
	r, err := Start("clientID", ":0")
	if sender != nil {
		t.Errorf("expected sender to be a nil value")
	}

	if r != nil {
		t.Errorf("expected r to be a nil value")
	}

	if err == nil {
		t.Errorf("expected error, but received none")
	}

	sender = nil
}
func TestStartPause(t *testing.T) {
	done := make(chan struct{})
	url, err := startUDPServer(done, func(b []byte) {})
	if err != nil {
		t.Errorf("expected no error when starting UDP server, but received %v", err)
	}
	defer close(done)

	_, err = Start("clientID", url)
	if sender == nil {
		t.Errorf("expected sender to be a nil value")
	}

	if err != nil {
		t.Errorf("expected no error, but received %v", err)
	}

	r := Get()
	r.Pause()

	if !r.metricsCh.IsPaused() {
		t.Errorf("expected monitoring to be paused, but was not")
	}

	r.Continue()
	if r.metricsCh.IsPaused() {
		t.Errorf("expected monitoring to be resumed, but was not")
	}

	sender = nil
}
