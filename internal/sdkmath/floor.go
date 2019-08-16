// +build go1.10

package sdkmath

import "math"

func Round(x float64) float64 {
	return math.Round(x)
}

func Modf(f float64) (int float64, frac float64) {
	return math.Modf(f)
}
