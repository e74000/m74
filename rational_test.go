package m74

import (
	"math"
	"testing"
)

func FuzzRationalApprox(f *testing.F) {
	f.Add(0.0)
	f.Add(1.0)
	f.Add(math.Inf(+1))
	f.Add(math.Inf(-1))
	f.Add(math.NaN())

	f.Fuzz(func(t *testing.T, f float64) {
		r := RationalApprox(f, -1)

		if math.Abs(f-r.AsFloat64()) > 0.000001 {
			t.Fail()
		}
	})
}
