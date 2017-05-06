// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import (
	m "math"
)

// Nextafter returns the next representable float32 value after x towards y.
// Special cases:
//	Nextafter32(x, x)   = x
//      Nextafter32(NaN, y) = NaN
//      Nextafter32(x, NaN) = NaN
//
// Since this is a float32 math package the 32 bit version has no number and the
// 64 bit version has the number in the method name.
func Nextafter(x, y float32) (r float32) {
	switch {
	case IsNaN(x) || IsNaN(y): // special case
		r = NaN()
	case x == y:
		r = x
	case x == 0:
		r = Copysign(Float32frombits(1), y)
	case (y > x) == (x > 0):
		r = Float32frombits(Float32bits(x) + 1)
	default:
		r = Float32frombits(Float32bits(x) - 1)
	}
	return
}

// Nextafter64 returns the next representable float64 value after x towards y.
// Special cases:
//      Nextafter64(x, x)   = x
//      Nextafter64(NaN, y) = NaN
//      Nextafter64(x, NaN) = NaN
//
// Since this is a float32 math package the 32 bit version has no number and the
// 64 bit version has the number in the method name.
func Nextafter64(x, y float64) (r float64) {
	return m.Nextafter(x, y)
}
