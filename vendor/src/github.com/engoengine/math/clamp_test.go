package math

import (
	"testing"
)

func TestClamp(t *testing.T) {
	t.Parallel()
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
		{
			f:        NaN(),
			low:      1,
			high:     2,
			expected: NaN(),
		},
	}

	for i, test := range tests {
		if r := Clamp(test.f, test.low, test.high); r != test.expected && !(IsNaN(test.expected) && IsNaN(r)) {
			t.Errorf("[%d] Clamp(%f,%f,%f) = %f, want %f", i, test.f, test.low, test.high, r, test.expected)
		}
	}
}
