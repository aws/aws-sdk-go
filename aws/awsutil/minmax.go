package awsutil

// MinInt64 returns the minimum of two int64 values.
func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}

	return b
}

// MaxInt64 returns the maximum of two int64 values.
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}
