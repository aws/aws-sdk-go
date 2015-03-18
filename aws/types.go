package aws

import "time"

// String converts a Go string into a string pointer.
func String(v string) *string {
	return &v
}

// Boolean converts a Go bool into a boolean pointer.
func Boolean(v bool) *bool {
	return &v
}

// Long converts a Go int64 into a long pointer.
func Long(v int64) *int64 {
	return &v
}

// Double converts a Go float64 into a double pointer.
func Double(v float64) *float64 {
	return &v
}

// Time converts a Go Time into a Time pointer
func Time(t time.Time) *time.Time {
	return &t
}
