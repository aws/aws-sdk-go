package ini

import (
	"testing"
)

func TestStringValue(t *testing.T) {
	cases := []struct {
		b             []byte
		expectedRead  int
		expectedError bool
		expectedValue string
	}{
		{
			b:             []byte(`"foo"`),
			expectedRead:  5,
			expectedValue: "foo",
		},
		{
			b:             []byte(`"123 !$_ 456 abc"`),
			expectedRead:  17,
			expectedValue: "123 !$_ 456 abc",
		},
		{
			b:             []byte("foo"),
			expectedError: true,
		},
		{
			b:             []byte(` "foo"`),
			expectedError: true,
		},
	}

	for i, c := range cases {
		a, n, err := getStringValue(c.b)

		if e := c.expectedValue; e != a {
			t.Errorf("%d: expected %v, but received %v", i+1, e, a)
		}

		if e, a := c.expectedRead, n; e != a {
			t.Errorf("%d: expected %v, but received %v", i+1, e, a)
		}

		if e, a := c.expectedError, err != nil; e != a {
			t.Errorf("%d: expected %v, but received %v", i+1, e, a)
		}
	}
}

func TestBoolValue(t *testing.T) {
	cases := []struct {
		b             []byte
		expectedRead  int
		expectedError bool
		expectedValue string
	}{
		{
			b:             []byte("true"),
			expectedRead:  4,
			expectedValue: "true",
		},
		{
			b:             []byte("false"),
			expectedRead:  5,
			expectedValue: "false",
		},
		{
			b:             []byte(`"false"`),
			expectedError: true,
		},
	}

	for _, c := range cases {
		a, n, err := getBoolValue(c.b)

		if e := c.expectedValue; e != a {
			t.Errorf("expected %v, but received %v", e, a)
		}

		if e, a := c.expectedRead, n; e != a {
			t.Errorf("expected %v, but received %v", e, a)
		}

		if e, a := c.expectedError, err != nil; e != a {
			t.Errorf("expected %v, but received %v", e, a)
		}
	}
}

func TestNumericalValue(t *testing.T) {
	cases := []struct {
		b             []byte
		expectedRead  int
		expectedError bool
		expectedValue string
	}{
		{
			b:             []byte("1.2"),
			expectedRead:  3,
			expectedValue: "1.2",
		},
		{
			b:             []byte("123"),
			expectedRead:  3,
			expectedValue: "123",
		},
		{
			b:             []byte("0x123A"),
			expectedRead:  6,
			expectedValue: "0x123A",
		},
		{
			b:             []byte("0b101"),
			expectedRead:  5,
			expectedValue: "0b101",
		},
		{
			b:             []byte("0o7"),
			expectedRead:  3,
			expectedValue: "0o7",
		},
		{
			b:             []byte(`"123"`),
			expectedError: true,
		},
		{
			b:             []byte("0xo123"),
			expectedError: true,
		},
		{
			b:             []byte("123A"),
			expectedError: true,
		},
	}

	for i, c := range cases {
		a, n, err := getNumericalValue(c.b)

		if e := c.expectedValue; e != a {
			t.Errorf("%d: expected %v, but received %v", i+1, e, a)
		}

		if e, a := c.expectedRead, n; e != a {
			t.Errorf("%d: expected %v, but received %v", i+1, e, a)
		}

		if e, a := c.expectedError, err != nil; e != a {
			t.Errorf("%d: expected %v, but received %v", i+1, e, a)
		}
	}
}
