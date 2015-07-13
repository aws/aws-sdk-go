package aws

import (
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

// Bool returns a pointer of the bool value passed in.
func Bool(v bool) *bool {
	return &v
}

// BoolValue returns the value of the bool pointer passed in or false if the
// pointer is nil.
func BoolValue(a *bool) bool {
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

// Int64 returns a pointer of the int64 value passed in.
func Int64(v int64) *int64 {
	return &v
}

// Int64Value returns the value of the int64 pointer passed in or zero if the
// pointer is nil.
func Int64Value(a *int64) int64 {
	if a != nil {
		return *a
	}
	return 0
}

// Float64 returns a pointer of the float64 value passed in.
func Float64(v float64) *float64 {
	return &v
}

// Float64Value returns the value of the float64 pointer passed in or zero if the
// pointer is nil.
func Float64Value(a *float64) float64 {
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
