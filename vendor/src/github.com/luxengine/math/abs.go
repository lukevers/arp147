// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Abs returns the absolute value of x.
//
// Special cases are:
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
func Abs(x float32) float32 {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0 // return correctly abs(-0)
	}
	return x
}

//func Abs(x float32) float32 //{ return float32(m.Abs(float64(x))) }

// Note: the software version of abs is actually faster then the ASM version,
// due to some inlining that doesn't happen with ASM, so we're just gonna remove
// it for now.
