package aws

// A StringValue is a string which may or may not be present.
type StringValue *string

// String converts a Go string into a StringValue.
func String(v string) StringValue {
	return &v
}

// A BooleanValue is a boolean which may or may not be present.
type BooleanValue *bool

// Boolean converts a Go bool into a BooleanValue.
func Boolean(v bool) BooleanValue {
	return &v
}

// True is the BooleanValue equivalent of the Go literal true.
func True() BooleanValue {
	return &t
}

// False is the BooleanValue equivalent of the Go literal false.
func False() BooleanValue {
	return &f
}

var (
	t = true
	f = false
)

// An IntegerValue is an integer which may or may not be present.
type IntegerValue *int

// Integer converts a Go int into an IntegerValue.
func Integer(v int) IntegerValue {
	return &v
}

// A LongValue is a 64-bit integer which may or may not be present.
type LongValue *int64

// Long converts a Go int64 into a LongValue.
func Long(v int64) LongValue {
	return &v
}

// A FloatValue is a 32-bit floating point number which may or may not be
// present.
type FloatValue *float32

// Float converts a Go float32 into a FloatValue.
func Float(v float32) FloatValue {
	return &v
}

// A DoubleValue is a 64-bit floating point number which may or may not be
// present.
type DoubleValue *float64

// Double converts a Go float64 into a DoubleValue.
func Double(v float64) DoubleValue {
	return &v
}
