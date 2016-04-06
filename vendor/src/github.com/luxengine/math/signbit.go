// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Signbit returns true if x is negative or negative zero.
func Signbit(x float32) bool {
	return Float32bits(x)&(1<<31) != 0
}
