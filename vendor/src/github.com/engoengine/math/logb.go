// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Logb returns the binary exponent of x.
//
// Special cases are:
//	Logb(±Inf) = +Inf
//	Logb(0) = -Inf
//	Logb(NaN) = NaN
func Logb(x float32) float32 {
	// special cases
	switch {
	case x == 0:
		return Inf(-1)
	case IsInf(x, 0):
		return Inf(1)
	case IsNaN(x):
		return x
	}
	return float32(ilogb(x))
}

// Ilogb returns the binary exponent of x as an integer.
//
// Special cases are:
//	Ilogb(±Inf) = MaxInt32
//	Ilogb(0) = MinInt32
//	Ilogb(NaN) = MaxInt32
func Ilogb(x float32) int {
	// special cases
	switch {
	case x == 0:
		return MinInt32
	case IsNaN(x):
		return MaxInt32
	case IsInf(x, 0):
		return MaxInt32
	}
	return ilogb(x)
}

// logb returns the binary exponent of x. It assumes x is finite and
// non-zero.
func ilogb(x float32) int {
	x, exp := normalize(x)
	return int((Float32bits(x)>>Shift)&Mask) - Bias + exp
}
