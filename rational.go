package m74

import (
	"fmt"
	"math"
)

type Rational struct {
	n int64
	d int64
}

func newRational[T Integer](n, d T) Rational {
	if Negative(d) {
		return Rational{
			n: -int64(n),
			d: int64(d),
		}
	}

	return Rational{
		n: int64(n),
		d: int64(d),
	}
}

func (r Rational) Positive() bool {
	return (r.n >= 0 && r.d >= 0) || (r.n < 0 && r.d < 0)
}

func (r Rational) Negative() bool {
	return (r.n < 0 && r.d >= 0) || (r.n >= 0 && r.d < 0)
}

func (r Rational) Sign() int {
	if r.Positive() {
		return 1
	} else {
		return -1
	}
}

func (r Rational) Abs() Rational {
	return Rational{
		n: Abs(r.n),
		d: Abs(r.d),
	}
}

func (r Rational) Simplify() Rational {
	gcd := GCD(r.n, r.d)

	if gcd == 1 || gcd == 0 {
		return r
	}

	return Rational{
		n: r.n / gcd,
		d: r.d / gcd,
	}
}

func (r Rational) Mul(s Rational) Rational {
	return Rational{
		n: r.n * s.n,
		d: r.d * s.d,
	}.Simplify()
}

func (r Rational) Div(s Rational) Rational {
	return Rational{
		n: r.n * s.d,
		d: r.d * s.n,
	}.Simplify()
}

func (r Rational) Add(s Rational) Rational {
	if r.d == s.d {
		return Rational{
			n: r.n + s.n,
			d: r.d,
		}.Simplify()
	}

	return Rational{
		n: r.n*s.d + s.n*r.d,
		d: r.d * s.d,
	}.Simplify()
}

func (r Rational) Sub(s Rational) Rational {
	if r.d == s.d {
		return Rational{
			n: r.n - s.n,
			d: r.d,
		}.Simplify()
	}

	return Rational{
		n: r.n*s.d - s.n*r.d,
		d: r.d * s.d,
	}.Simplify()
}

func (r Rational) AsFloat64() float64 {
	return float64(r.n) / float64(r.d)
}

func (r Rational) AsFloat32() float32 {
	return float32(r.n) / float32(r.d)
}

func (r Rational) String() string {
	return fmt.Sprintf("%d / %d", r.n, r.d)
}

func RationalApprox[T Real](v T, err float64) Rational {
	if math.IsInf(float64(v), +1) {
		return newRational(1, 0)
	} else if math.IsInf(float64(v), -1) {
		return newRational(-1, 0)
	} else if math.IsNaN(float64(v)) {
		return newRational(0, 0)
	}

	if err == -1 {
		err = 0.000001
	}

	n := math.Floor(float64(v))
	x := float64(v) - n

	if x < err {
		return newRational(int(n), 1)
	} else if x > 1-err {
		return newRational(int(n)+1, 1)
	}

	upper := newRational(0, 1)
	lower := newRational(1, 1)

	for {
		mid := Rational{
			n: upper.n + lower.n,
			d: upper.d + lower.d,
		}

		if float64(mid.d)*(x+err) < float64(mid.n) {
			lower = mid
		} else if float64(mid.n) < (x-err)*float64(mid.d) {
			upper = mid
		} else {
			return mid.Add(newRational(int(n), 1))
		}
	}
}
