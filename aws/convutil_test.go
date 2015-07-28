package aws_test

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
)

var testCasesStringSlice = [][]string{
	[]string{"a", "b", "c", "d", "e"},
	[]string{"a", "b", "", "", "e"},
}

func TestStringSlice(t *testing.T) {
	for idx, in := range testCasesStringSlice {
		if in == nil {
			continue
		}
		out := aws.StringSlice(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			assert.Equal(t, in[i], *(out[i]), "Unexpected value at idx %d", idx)
		}

		out2 := aws.StringValueSlice(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		assert.Equal(t, in, out2, "Unexpected value at idx %d", idx)
	}
}

var testCasesStringValueSlice = [][]*string{
	[]*string{aws.String("a"), aws.String("b"), nil, aws.String("c")},
}

func TestStringValueSlice(t *testing.T) {
	for idx, in := range testCasesStringValueSlice {
		if in == nil {
			continue
		}
		out := aws.StringValueSlice(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			if in[i] == nil {
				assert.Empty(t, out[i], "Unexpected value at idx %d", idx)
			} else {
				assert.Equal(t, *(in[i]), out[i], "Unexpected value at idx %d", idx)
			}
		}

		out2 := aws.StringSlice(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		for i := range out2 {
			if in[i] == nil {
				assert.Empty(t, *(out2[i]), "Unexpected value at idx %d", idx)
			} else {
				assert.Equal(t, in[i], out2[i], "Unexpected value at idx %d", idx)
			}
		}
	}
}

var testCasesStringMap = []map[string]string{
	map[string]string{"a": "1", "b": "2", "c": "3"},
}

func TestStringMap(t *testing.T) {
	for idx, in := range testCasesStringMap {
		if in == nil {
			continue
		}
		out := aws.StringMap(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			assert.Equal(t, in[i], *(out[i]), "Unexpected value at idx %d", idx)
		}

		out2 := aws.StringValueMap(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		assert.Equal(t, in, out2, "Unexpected value at idx %d", idx)
	}
}

var testCasesBoolSlice = [][]bool{
	[]bool{true, true, false, false},
}

func TestBoolSlice(t *testing.T) {
	for idx, in := range testCasesBoolSlice {
		if in == nil {
			continue
		}
		out := aws.BoolSlice(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			assert.Equal(t, in[i], *(out[i]), "Unexpected value at idx %d", idx)
		}

		out2 := aws.BoolValueSlice(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		assert.Equal(t, in, out2, "Unexpected value at idx %d", idx)
	}
}

var testCasesBoolValueSlice = [][]*bool{}

func TestBoolValueSlice(t *testing.T) {
	for idx, in := range testCasesBoolValueSlice {
		if in == nil {
			continue
		}
		out := aws.BoolValueSlice(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			if in[i] == nil {
				assert.Empty(t, out[i], "Unexpected value at idx %d", idx)
			} else {
				assert.Equal(t, *(in[i]), out[i], "Unexpected value at idx %d", idx)
			}
		}

		out2 := aws.BoolSlice(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		for i := range out2 {
			if in[i] == nil {
				assert.Empty(t, *(out2[i]), "Unexpected value at idx %d", idx)
			} else {
				assert.Equal(t, in[i], out2[i], "Unexpected value at idx %d", idx)
			}
		}
	}
}

var testCasesBoolMap = []map[string]bool{
	map[string]bool{"a": true, "b": false, "c": true},
}

func TestBoolMap(t *testing.T) {
	for idx, in := range testCasesBoolMap {
		if in == nil {
			continue
		}
		out := aws.BoolMap(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			assert.Equal(t, in[i], *(out[i]), "Unexpected value at idx %d", idx)
		}

		out2 := aws.BoolValueMap(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		assert.Equal(t, in, out2, "Unexpected value at idx %d", idx)
	}
}

var testCasesIntSlice = [][]int{
	[]int{1, 2, 3, 4},
}

func TestIntSlice(t *testing.T) {
	for idx, in := range testCasesIntSlice {
		if in == nil {
			continue
		}
		out := aws.IntSlice(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			assert.Equal(t, in[i], *(out[i]), "Unexpected value at idx %d", idx)
		}

		out2 := aws.IntValueSlice(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		assert.Equal(t, in, out2, "Unexpected value at idx %d", idx)
	}
}

var testCasesIntValueSlice = [][]*int{}

func TestIntValueSlice(t *testing.T) {
	for idx, in := range testCasesIntValueSlice {
		if in == nil {
			continue
		}
		out := aws.IntValueSlice(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			if in[i] == nil {
				assert.Empty(t, out[i], "Unexpected value at idx %d", idx)
			} else {
				assert.Equal(t, *(in[i]), out[i], "Unexpected value at idx %d", idx)
			}
		}

		out2 := aws.IntSlice(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		for i := range out2 {
			if in[i] == nil {
				assert.Empty(t, *(out2[i]), "Unexpected value at idx %d", idx)
			} else {
				assert.Equal(t, in[i], out2[i], "Unexpected value at idx %d", idx)
			}
		}
	}
}

var testCasesIntMap = []map[string]int{
	map[string]int{"a": 3, "b": 2, "c": 1},
}

func TestIntMap(t *testing.T) {
	for idx, in := range testCasesIntMap {
		if in == nil {
			continue
		}
		out := aws.IntMap(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			assert.Equal(t, in[i], *(out[i]), "Unexpected value at idx %d", idx)
		}

		out2 := aws.IntValueMap(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		assert.Equal(t, in, out2, "Unexpected value at idx %d", idx)
	}
}

var testCasesInt64Slice = [][]int64{
	[]int64{1, 2, 3, 4},
}

func TestInt64Slice(t *testing.T) {
	for idx, in := range testCasesInt64Slice {
		if in == nil {
			continue
		}
		out := aws.Int64Slice(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			assert.Equal(t, in[i], *(out[i]), "Unexpected value at idx %d", idx)
		}

		out2 := aws.Int64ValueSlice(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		assert.Equal(t, in, out2, "Unexpected value at idx %d", idx)
	}
}

var testCasesInt64ValueSlice = [][]*int64{}

func TestInt64ValueSlice(t *testing.T) {
	for idx, in := range testCasesInt64ValueSlice {
		if in == nil {
			continue
		}
		out := aws.Int64ValueSlice(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			if in[i] == nil {
				assert.Empty(t, out[i], "Unexpected value at idx %d", idx)
			} else {
				assert.Equal(t, *(in[i]), out[i], "Unexpected value at idx %d", idx)
			}
		}

		out2 := aws.Int64Slice(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		for i := range out2 {
			if in[i] == nil {
				assert.Empty(t, *(out2[i]), "Unexpected value at idx %d", idx)
			} else {
				assert.Equal(t, in[i], out2[i], "Unexpected value at idx %d", idx)
			}
		}
	}
}

var testCasesInt64Map = []map[string]int64{
	map[string]int64{"a": 3, "b": 2, "c": 1},
}

func TestInt64Map(t *testing.T) {
	for idx, in := range testCasesInt64Map {
		if in == nil {
			continue
		}
		out := aws.Int64Map(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			assert.Equal(t, in[i], *(out[i]), "Unexpected value at idx %d", idx)
		}

		out2 := aws.Int64ValueMap(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		assert.Equal(t, in, out2, "Unexpected value at idx %d", idx)
	}
}

var testCasesFloat64Slice = [][]float64{
	[]float64{1, 2, 3, 4},
}

func TestFloat64Slice(t *testing.T) {
	for idx, in := range testCasesFloat64Slice {
		if in == nil {
			continue
		}
		out := aws.Float64Slice(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			assert.Equal(t, in[i], *(out[i]), "Unexpected value at idx %d", idx)
		}

		out2 := aws.Float64ValueSlice(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		assert.Equal(t, in, out2, "Unexpected value at idx %d", idx)
	}
}

var testCasesFloat64ValueSlice = [][]*float64{}

func TestFloat64ValueSlice(t *testing.T) {
	for idx, in := range testCasesFloat64ValueSlice {
		if in == nil {
			continue
		}
		out := aws.Float64ValueSlice(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			if in[i] == nil {
				assert.Empty(t, out[i], "Unexpected value at idx %d", idx)
			} else {
				assert.Equal(t, *(in[i]), out[i], "Unexpected value at idx %d", idx)
			}
		}

		out2 := aws.Float64Slice(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		for i := range out2 {
			if in[i] == nil {
				assert.Empty(t, *(out2[i]), "Unexpected value at idx %d", idx)
			} else {
				assert.Equal(t, in[i], out2[i], "Unexpected value at idx %d", idx)
			}
		}
	}
}

var testCasesFloat64Map = []map[string]float64{
	map[string]float64{"a": 3, "b": 2, "c": 1},
}

func TestFloat64Map(t *testing.T) {
	for idx, in := range testCasesFloat64Map {
		if in == nil {
			continue
		}
		out := aws.Float64Map(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			assert.Equal(t, in[i], *(out[i]), "Unexpected value at idx %d", idx)
		}

		out2 := aws.Float64ValueMap(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		assert.Equal(t, in, out2, "Unexpected value at idx %d", idx)
	}
}

var testCasesTimeSlice = [][]time.Time{
	[]time.Time{time.Now(), time.Now().AddDate(100, 0, 0)},
}

func TestTimeSlice(t *testing.T) {
	for idx, in := range testCasesTimeSlice {
		if in == nil {
			continue
		}
		out := aws.TimeSlice(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			assert.Equal(t, in[i], *(out[i]), "Unexpected value at idx %d", idx)
		}

		out2 := aws.TimeValueSlice(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		assert.Equal(t, in, out2, "Unexpected value at idx %d", idx)
	}
}

var testCasesTimeValueSlice = [][]*time.Time{}

func TestTimeValueSlice(t *testing.T) {
	for idx, in := range testCasesTimeValueSlice {
		if in == nil {
			continue
		}
		out := aws.TimeValueSlice(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			if in[i] == nil {
				assert.Empty(t, out[i], "Unexpected value at idx %d", idx)
			} else {
				assert.Equal(t, *(in[i]), out[i], "Unexpected value at idx %d", idx)
			}
		}

		out2 := aws.TimeSlice(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		for i := range out2 {
			if in[i] == nil {
				assert.Empty(t, *(out2[i]), "Unexpected value at idx %d", idx)
			} else {
				assert.Equal(t, in[i], out2[i], "Unexpected value at idx %d", idx)
			}
		}
	}
}

var testCasesTimeMap = []map[string]time.Time{
	map[string]time.Time{"a": time.Now().AddDate(-100, 0, 0), "b": time.Now()},
}

func TestTimeMap(t *testing.T) {
	for idx, in := range testCasesTimeMap {
		if in == nil {
			continue
		}
		out := aws.TimeMap(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			assert.Equal(t, in[i], *(out[i]), "Unexpected value at idx %d", idx)
		}

		out2 := aws.TimeValueMap(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		assert.Equal(t, in, out2, "Unexpected value at idx %d", idx)
	}
}
