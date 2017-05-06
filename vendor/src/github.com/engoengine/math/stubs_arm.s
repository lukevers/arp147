// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build arm

#include "textflag.h"

// func Sqrt(x float32) float32	
TEXT ·Sqrt(SB),NOSPLIT,$0
	JMP ·sqrt(SB)

// func Cos(x float32) float32	
TEXT ·Cos(SB),NOSPLIT,$0
	JMP ·cos(SB)

// func Sin(x float32) float32	
TEXT ·Sin(SB),NOSPLIT,$0
	JMP ·sin(SB)
