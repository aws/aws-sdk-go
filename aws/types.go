package aws

import (
	"math"
	"strconv"
	"time"
)

// String converts a Go string into a StringValue.
func String(v string) *string {
	return &v
}

// Boolean converts a Go bool into a BooleanValue.
func Boolean(v bool) *bool {
	return &v
}

// Long converts a Go int64 into a LongValue.
func Long(v int64) *int64 {
	return &v
}

// Double converts a Go float64 into a DoubleValue.
func Double(v float64) *float64 {
	return &v
}

func NewTime(t time.Time) *Time {
	return &Time{t}
}

// A Time is a timestamp
type Time struct {
	time.Time
}

// MarshalJSON marshals the timestamp as a float.
func (t Time) MarshalJSON() (text []byte, err error) {
	n := float64(t.Time.UnixNano()) / 1e9
	s := strconv.FormatFloat(n, 'f', -1, 64)
	return []byte(s), nil
}

// UnmarshalJSON unmarshals the timestamp from a float.
func (t *Time) UnmarshalJSON(text []byte) error {
	f, err := strconv.ParseFloat(string(text), 64)
	if err != nil {
		return err
	}

	sec := math.Floor(f)
	nsec := (f - sec) * 1e9

	t.Time = time.Unix(int64(sec), int64(nsec)).UTC()
	return nil
}
