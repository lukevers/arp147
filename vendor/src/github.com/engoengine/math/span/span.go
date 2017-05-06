package span

import (
	"github.com/engoengine/math"
)

// Span represents an interval.
type Span struct {
	Min, Max float32
}

// Add 2 span togheter
//	[a, b] + [c, d] = [a+c, b+d]
func (s0 Span) Add(s1 Span) Span {
	return Span{s0.Min + s1.Min, s0.Max + s1.Max}
}

// Sub 2 span togheter
//	[a, b] - [c, d] = [a-c, b-d]
func (s0 Span) Sub(s1 Span) Span {
	return Span{s0.Min - s1.Min, s0.Max - s1.Max}
}

// Mul multiply this these 2 span togheter
//	[a, b] * [c, d] = [min(ac, ad, bc, bd), max(ac, ad, bc, bd)]
func (s0 Span) Mul(s1 Span) Span {
	return Span{
		math.Min(math.Min(s0.Min*s1.Min, s0.Max*s1.Max), math.Min(s0.Max*s1.Min, s0.Max*s1.Min)),
		math.Max(math.Max(s0.Min*s1.Min, s0.Max*s1.Max), math.Max(s0.Max*s1.Min, s0.Max*s1.Min)),
	}
}

// Div returns s0/s1
func (s0 Span) Div(s1 Span) Span {
	s2 := Span{1 / s1.Min, 1 / s1.Max}
	return s0.Mul(s2)
}

// Abs return the absolute of the given span.
func Abs(s Span) Span {
	return Span{math.Abs(s.Min), math.Abs(s.Max)}
}
