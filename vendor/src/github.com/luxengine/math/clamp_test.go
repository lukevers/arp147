package math

import (
	"testing"
)

func TestClamp(t *testing.T) {
	tests := []struct {
		f, low, high, expected float32
	}{
		{
			f:        0.5,
			low:      0,
			high:     1,
			expected: 0.5,
		},
		{
			f:        -0.5,
			low:      0,
			high:     1,
			expected: 0,
		},
		{
			f:        1.5,
			low:      0,
			high:     1,
			expected: 1,
		},

		{
			f:        1.5,
			low:      1,
			high:     2,
			expected: 1.5,
		},
		{
			f:        0.5,
			low:      1,
			high:     2,
			expected: 1,
		},
		{
			f:        2.5,
			low:      1,
			high:     2,
			expected: 2,
		},
	}

	for i, test := range tests {
		if r := Clamp(test.f, test.low, test.high); r != test.expected {
			t.Errorf("[%d] Clamp(%f,%f,%f) = %f, want %f", i, test.f, test.low, test.high, r, test.expected)
		}
	}
}
