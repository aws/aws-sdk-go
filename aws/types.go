package aws

type StringValue *string

func String(v string) StringValue {
	return &v
}

type BooleanValue *bool

func True() BooleanValue {
	return &t
}

func False() BooleanValue {
	return &f
}

type IntegerValue *int

func Integer(v int) IntegerValue {
	return &v
}

type LongValue *int64

func Long(v int64) LongValue {
	return &v
}

type FloatValue *float32

func Float(v float32) FloatValue {
	return &v
}

type DoubleValue *float64

func Double(v float64) DoubleValue {
	return &v
}

var (
	t = true
	f = false
)
