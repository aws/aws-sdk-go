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
}

func (e *OnceError) Err() error {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.err
}

func (e *OnceError) SetOnce(err error) {
	if err == nil {
		return
	}

	e.mu.Lock()
	defer e.mu.Unlock()
	if e.err == nil {
		e.err = err
	}
}
