package aws

import (
	"fmt"
	"io"
	"time"
)

// String returns a pointer of the string value passed in.
func String(v string) *string {
	return &v
}

// StringValue returns the value of the string pointer passed in or empty
// string if the pointer is nil.
func StringValue(a *string) string {
	if a != nil {
		return *a
	}
	return ""
}

// Boolean returns a pointer of the bool value passed in.
func Boolean(v bool) *bool {
	return &v
}

// BooleanValue returns the value of the bool pointer passed in or false if the
// pointer is nil.
func BooleanValue(a *bool) bool {
	if a != nil {
		return *a
	}
	return false
}

// Int returns a pointer of the int value passed in.
func Int(v int) *int {
	return &v
}

// IntValue returns the value of the int pointer passed in or zero if the
// pointer is nil.
func IntValue(a *int) int {
	if a != nil {
		return *a
	}
	return 0
}

// Long returns a pointer of the int64 value passed in.
func Long(v int64) *int64 {
	return &v
}

// LongValue returns the value of the int64 pointer passed in or zero if the
// pointer is nil.
func LongValue(a *int64) int64 {
	if a != nil {
		return *a
	}
	return 0
}

// Double returns a pointer of the float64 value passed in.
func Double(v float64) *float64 {
	return &v
}

// DoubleValue returns the value of the float64 pointer passed in or zero if the
// pointer is nil.
func DoubleValue(a *float64) float64 {
	if a != nil {
		return *a
	}
	return 0
}

// Time returns a pointer of the Time value passed in.
func Time(t time.Time) *time.Time {
	return &t
}

// TimeValue returns the value of the Time pointer passed in or zero
// time if the pointer is nil.
func TimeValue(a *time.Time) time.Time {
	if a != nil {
		return *a
	}
	return time.Time{}
}

// ReadSeekCloser wraps a io.Reader returning a ReaderSeakerCloser
func ReadSeekCloser(r io.Reader) ReaderSeekerCloser {
	return ReaderSeekerCloser{r}
}

// ReaderSeekerCloser represents a reader that can also delegate io.Seeker and
// io.Closer interfaces to the underlying object if they are available.
type ReaderSeekerCloser struct {
	r io.Reader
}

// Read reads from the reader up to size of p. The number of bytes read, and
// error if it occurred will be returned.
//
// If the reader is not an io.Reader zero bytes read, and nil error will be returned.
//
// Performs the same functionality as io.Reader Read
func (r ReaderSeekerCloser) Read(p []byte) (int, error) {
	switch t := r.r.(type) {
	case io.Reader:
		return t.Read(p)
	}
	return 0, nil
}

// Seek sets the offset for the next Read to offset, interpreted according to
// whence: 0 means relative to the origin of the file, 1 means relative to the
// current offset, and 2 means relative to the end. Seek returns the new offset
// and an error, if any.
//
// If the ReaderSeekerCloser is not an io.Seeker nothing will be done.
func (r ReaderSeekerCloser) Seek(offset int64, whence int) (int64, error) {
	switch t := r.r.(type) {
	case io.Seeker:
		return t.Seek(offset, whence)
	}
	return int64(0), nil
}

// Close closes the ReaderSeekerCloser.
//
// If the ReaderSeekerCloser is not an io.Closer nothing will be done.
func (r ReaderSeekerCloser) Close() error {
	switch t := r.r.(type) {
	case io.Closer:
		return t.Close()
	}
	return nil
}

// A SettableBool provides a boolean value which includes the state if
// the value was set or unset.  The set state is in addition to the value's
// value(true|false)
type SettableBool struct {
	value bool
	set   bool
}

// SetBool returns a SettableBool with a value set
func SetBool(value bool) SettableBool {
	return SettableBool{value: value, set: true}
}

// Get returns the value. Will always be false if the SettableBool was not set.
func (b *SettableBool) Get() bool {
	if !b.set {
		return false
	}
	return b.value
}

// Set sets the value and updates the state that the value has been set.
func (b *SettableBool) Set(value bool) {
	b.value = value
	b.set = true
}

// IsSet returns if the value has been set
func (b *SettableBool) IsSet() bool {
	return b.set
}

// Reset resets the state and value of the SettableBool to its initial default
// state of not set and zero value.
func (b *SettableBool) Reset() {
	b.value = false
	b.set = false
}

// String returns the string representation of the value if set. Zero if not set.
func (b *SettableBool) String() string {
	return fmt.Sprintf("%t", b.Get())
}

// GoString returns the string representation of the SettableBool value and state
func (b *SettableBool) GoString() string {
	return fmt.Sprintf("Bool{value:%t, set:%t}", b.value, b.set)
}
