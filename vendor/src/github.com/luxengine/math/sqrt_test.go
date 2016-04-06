package math

import (
	"testing"
)

var sqrtv = []float32{
	2.2313699855653004178179799e+00,
	2.7818829105618685382239619e+00,
	5.2619393351308207940064676e-01,
	2.2384377203502809905444337e+00,
	3.104237975937876203857968e+00,
	1.7106657465582673083304144e+00,
	2.2867189460131345235538447e+00,
	1.6516476149988743582497364e+00,
	1.3510396309834566963559155e+00,
	2.9471892592823585310668477e+00,
}

func TestSqrtSoftware(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		a := Abs(vf[i])
		a = Abs(vf[i])
		if f := sqrt(a); sqrtv[i] != f {
			t.Errorf("sqrt(%g) = %g, want %g", a, f, sqrtv[i])
		}
	}
	var r float32
	for i := 0; i < len(vf); i++ {
		a := Abs(vf[i])
		a = Abs(vf[i])
		if sqrtC(a, &r); sqrtv[i] != r {
			t.Errorf("sqrt(%g) = %g, want %g", a, r, sqrtv[i])
		}
	}

	for i := 0; i < len(vfsqrtSC); i++ {
		if f := sqrt(vfsqrtSC[i]); !alike(sqrtSC[i], f) {
			t.Errorf("sqrt(%g) = %g, want %g", vfsqrtSC[i], f, sqrtSC[i])
		}
	}

	// subnormal float with LSB 0
	value := Float32frombits(uint32(2))
	expected := float32(5.293956e-23)
	if f := sqrt(value); expected != f {
		t.Errorf("sqrt(%g) = %g, want %g", value, f, expected)
	}
}
