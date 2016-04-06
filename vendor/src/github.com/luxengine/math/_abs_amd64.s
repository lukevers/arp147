// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

//func Abs(x float32) float32
TEXT ·Abs(SB),NOSPLIT,$0
	//FMOVF = F(to floating-point register)MOV(move)F(move a float, 32 bit) x(the argument)+(FP)(at the start of the frame pointer), F0(to the first register of the floating point registers)
	FMOVF   x+0(FP), F0  // F0=x
	//FABS = the F0 register is now the abs of what it was
	FABS                 // F0=|x|
	//FMOVF(like FMOVF)P(pointer ?) F0(source), ret+8(FP)(we're on amd64 so the destination is the 8 bytes behind the frame pointer, but its a float32 operation.)
	FMOVFP  F0, ret+8(FP)
	RET
