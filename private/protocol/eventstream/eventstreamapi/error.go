package eventstreamapi

import (
	"fmt"
	"sync"
)

type messageError struct {
	code string
	msg  string
}

func (e messageError) Code() string {
	return e.code
}

func (e messageError) Message() string {
	return e.msg
}

func (e messageError) Error() string {
	return fmt.Sprintf("%s: %s", e.code, e.msg)
}

func (e messageError) OrigErr() error {
	return nil
}

type OnceError struct {
	mu  sync.RWMutex
	err error
	ch  chan struct{}
}

func NewOnceError() *OnceError {
	return &OnceError{
		ch: make(chan struct{}, 1),
	}
}

func (e *OnceError) Err() error {
	e.mu.RLock()
	err := e.err
	e.mu.RUnlock()

	return err
}

func (e *OnceError) SetError(err error) {
	if err == nil {
		return
	}

	e.mu.Lock()
	if e.err == nil {
		e.err = err
		close(e.ch)
	}
	e.mu.Unlock()
}

func (e *OnceError) ErrorSet() <-chan struct{} {
	return e.ch
}
