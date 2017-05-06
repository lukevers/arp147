// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ppc64 ppc64le

#include "textflag.h"

// func Sqrt(x float32) float32	
TEXT ·Sqrt(SB),NOSPLIT,$0
	JMP ·sqrt(SB)
