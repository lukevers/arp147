// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build arm64

#include "textflag.h"

TEXT ·Abs(SB),NOSPLIT,$0
	JMP ·abs(SB)

// func Sqrt(x float32) float32	
TEXT ·Sqrt(SB),NOSPLIT,$0
	JMP ·sqrt(SB)
