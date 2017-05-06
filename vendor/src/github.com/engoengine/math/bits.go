// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

//IEEE float: 1 sign 8 expo 23 mantissa
//IEEE double: 1 sign 11 expo 52 mantissa

// Constants used for various operations.
const (
	Totalsize    = 32
	MantissaSize = 23
	ExponentSize = 8
	Signsize     = 1
	//uvnan = 0x7FF8000000000001
	//0111 1111 1111 1000 0000 0000 0000 0000
	//0000 0000 0000 0000 0000 0000 0000 0001

	Uvnan = 0x7FC00001
	//0111 1111 1100 0000
	//0000 0000 0000 0001

	//uvinf = 0x7FF0000000000000
	//0111 1111 1111 0000 0000 0000 0000 0000
	//0000 0000 0000 0000 0000 0000 0000 0000

	Uvinf = 0x7F800000
	//0111 1111 1000 0000
	//0000 0000 0000 0000

	//uvneginf = 0xFFF0000000000000
	//1111 1111 1111 0000 0000 0000 0000 0000
	//0000 0000 0000 0000 0000 0000 0000 0000

	Uvneginf = 0xFF800000
	//1111 1111 1000 0000
	//0000 0000 0000 0000

	//mask = 0x7FF
	//0111 1111 1111

	//mask for exponent part
	Mask = 0xFF
	//1111 1111

	//shift = 64-11-1

	Shift = Totalsize - ExponentSize - Signsize

	//bias = 1023

	//bias = (2 ^ (exponentSize - 1)) - 1
	Bias = 127
)

// Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
func Inf(sign int) float32 {
	var v uint32
	if sign >= 0 {
		v = Uvinf
	} else {
		v = Uvneginf
	}
	return Float32frombits(v)
}

// NaN returns an IEEE 754 ``not-a-number'' value.
func NaN() float32 { return Float32frombits(Uvnan) }

// IsNaN reports whether f is an IEEE 754 ``not-a-number'' value.
func IsNaN(f float32) (is bool) {
	// IEEE 754 says that only NaNs satisfy f != f.
	// To avoid the floating-point hardware, could use:
	//	x := Float32bits(f);
	//	return uint32(x>>shift)&mask == mask && x != uvinf && x != uvneginf
	return f != f
}

// IsInf reports whether f is an infinity, according to sign.
// If sign > 0, IsInf reports whether f is positive infinity.
// If sign < 0, IsInf reports whether f is negative infinity.
// If sign == 0, IsInf reports whether f is either infinity.
func IsInf(f float32, sign int) bool {
	// Test for infinity by comparing against maximum float.
	// To avoid the floating-point hardware, could use:
	//	x := Float64bits(f);
	//	return sign >= 0 && x == uvinf || sign <= 0 && x == uvneginf;
	return sign >= 0 && f > MaxFloat32 || sign <= 0 && f < -MaxFloat32
}

// normalize returns a normal number y and exponent exp
// satisfying x == y × 2**exp. It assumes x is finite and non-zero.
func normalize(x float32) (y float32, exp int) {
	const SmallestNormal = 1.175494350822287507968736537222245677818665556e-38 // 2**-126
	if Abs(x) < SmallestNormal {
		return x * (1 << 23), -23
	}
	return x, 0
}
