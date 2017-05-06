package imath

import (
	"unsafe"
)

// Integer limit values.
const (
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
)

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Cbrt returns the cube root of x.
func Cbrt(x int) int {
	panic("not implemented")
}

// Copysign returns a value with the magnitude of x and the sign of y.
func Copysign(x, y int) int {
	if x > 0 {
		if y > 0 {
			return x
		}
		return -x
	}
	if y > 0 {
		return -x
	}
	return x
}

// Dim returns the maximum of x-y or 0.
func Dim(x, y int) int {
	return Max(x-y, 0)
}

// Exp2 returns 2**x, the base-2 exponential of x.
func Exp2(x int) int {
	return 2 << uint(x)
}

// Intbits return the binary representation of i.
func Intbits(i int) uint {
	return *(*uint)(unsafe.Pointer(&i))
}

// Intfrombits returns the int represented from b.
func Intfrombits(b uint) int {
	return *(*int)(unsafe.Pointer(&b))
}

// Hypot returns Sqrt(p*p + q*q), taking care to avoid unnecessary overflow and
// underflow.
func Hypot(p, q int) int {
	if p < 0 {
		p = -p
	}
	if q < 0 {
		q = -q
	}
	if p < q {
		p, q = q, p
	}
	if p == 0 {
		return 0
	}
	q = q / p
	return p * Sqrt(1+q*q)
}

// Log returns the natural logarithm of x.
func Log(x int) int {
	panic("not implemented")
}

// log10 returns the decimal logarithm of x.
func log10(n int) int {
	// #define S(k, m) if (n >= UINT64_C(m)) { i += k; n /= UINT64_C(m); }
	// int i = -(n == 0);
	// S(16,10000000000000000); S(8,100000000); S(4,10000); S(2,100); S(1,10);
	// return i;
	/*var i int
	if n == 0 {
		i = -1
	} else {
		i = -0
	}

	if n >= 10000000000000000 {
		i += 16
		n /= 10000000000000000
	}

	if n >= 100000000 {
		i += 8
		n /= 100000000
	}

	if n >= 10000 {
		i += 4
		n /= 10000
	}

	if n >= 100 {
		i += 2
		n /= 100
	}

	if n >= 10 {
		i++
		n /= 10
	}

	return i*/
	panic("not implemented")
}

// log2 returns the binary logarithm of x.
func log2(n int) int {
	//S(k) if (n >= (UINT64_C(1) << k)) { i += k; n >>= k; }
	//int i = -(n == 0); S(32); S(16); S(8); S(4); S(2); S(1); return i;
	/*var i int
	if n == 0 {
		i = -1
	} else {
		i = -0
	}

	if n >= (1 << 32) {
		i += 32
		n >>= 32
	}

	if n >= (1 << 16) {
		i += 16
		n >>= 16
	}

	if n >= (1 << 8) {
		i += 8
		n >>= 8
	}

	if n >= (1 << 4) {
		i += 4
		n >>= 4
	}

	if n >= (1 << 2) {
		i += 2
		n >>= 2
	}

	if n >= (1 << 1) {
		i++
		n >>= 1
	}

	return i*/
	panic("not implemented")
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Mod returns the x%y.
func Mod(x, y int) int {
	return x % y
}

// Nextafter returns the next representable int value after x towards y.
func Nextafter(x, y int) (r int) {
	if x > y {
		return x + 1
	}
	return x - 1
}

// Pow returns x**y, the base-x exponential of y.
func Pow(x, y int) int {
	if y == 0 {
		return 1
	}
	if y == 1 {
		return x
	}

	tmp := Pow(x, y/2)
	if y%2 == 0 {
		return tmp * tmp
	}
	return x * tmp * tmp
}

// Pow10 returns 10**e, the base-10 exponential of e.
func Pow10(e int) int {
	if e == 0 {
		return 1
	}
	if e == 1 {
		return 10
	}

	tmp := Pow10(e / 2)
	if e%2 == 0 {
		return tmp * tmp
	}
	return 10 * tmp * tmp
}

// Signbit returns true if x is negative or negative zero.
func Signbit(x int) bool {
	if x < 0 {
		return false
	}
	return true
}

// Sqrt returns the square root of x.
func Sqrt(x int) int {
	op := x
	res := 0
	one := 1 << 30

	for one > op {
		one >>= 2
	}

	for one != 0 {
		if op >= res+one {
			op -= res + one
			res = (res >> 1) + one
		} else {
			res >>= 1
		}
		one >>= 2
	}
	return res

}

// func Int32bits(f float32) uint32
// func Int32frombits(b uint32) float32
// func Int64bits(f float32) uint32
// func Int64frombits(b uint32) float32
// func Nextafter32(x, y float32) (r float32)
// func Nextafter64(x, y float32) (r float32)
