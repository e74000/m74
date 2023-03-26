package m74

import (
	"fmt"
	"math"
)

// Rational represents fraction (N/D)
type Rational struct {
	N int64
	D int64
}

// NewRational returns a rational
func NewRational[T Integer](n, d T) Rational {
	if Negative(d) {
		return Rational{
			N: -int64(n),
			D: int64(d),
		}
	}

	return Rational{
		N: int64(n),
		D: int64(d),
	}
}

// Positive returns true if the rational is positive or zero
func (r Rational) Positive() bool {
	return (r.N >= 0 && r.D >= 0) || (r.N < 0 && r.D < 0)
}

// Negative returns true if the rational is negative or zero
func (r Rational) Negative() bool {
	return (r.N < 0 && r.D >= 0) || (r.N >= 0 && r.D < 0)
}

// Sign returns +1 if the rational is positive or zero, and -1 otherwise
func (r Rational) Sign() int {
	if r.Positive() {
		return +1
	} else {
		return -1
	}
}

// IsInt returns true if the receiver is integer, as well as the value as an integer
func (r Rational) IsInt() (i int, ok bool) {
	s := r.Simplify()

	if s.D == 1 {
		return int(s.N), true
	} else {
		return 0, false
	}
}

// Pow returns the nth power of the receiver
func (r Rational) Pow(n int) Rational {
	if n > 0 {
		out := NewRational(1, 1)

		for i := 0; i < n; i++ {
			out = out.Mul(r)
		}

		return out.Simplify()
	} else if n < 0 {
		out := NewRational(1, 1)

		for i := 0; i < n; i++ {
			out = out.Div(r)
		}

		return out.Simplify()
	}

	return NewRational(1, 1)
}

// Abs gets the absolute value of a rational, by multiplying it by -1 if negative
func (r Rational) Abs() Rational {
	return Rational{
		N: Abs(r.N),
		D: Abs(r.D),
	}
}

// Simplify removes any common divisors from the numerator/denominator
func (r Rational) Simplify() Rational {
	gcd := GCD(r.N, r.D)

	if gcd == 1 || gcd == 0 {
		return r
	}

	return Rational{
		N: r.N / gcd,
		D: r.D / gcd,
	}
}

// Mul multiplies the receiver by `s`
func (r Rational) Mul(s Rational) Rational {
	return Rational{
		N: r.N * s.N,
		D: r.D * s.D,
	}.Simplify()
}

// Div divides the receiver by `s`
func (r Rational) Div(s Rational) Rational {
	return Rational{
		N: r.N * s.D,
		D: r.D * s.N,
	}.Simplify()
}

// Add adds `s` to the receiver
func (r Rational) Add(s Rational) Rational {
	if r.D == s.D {
		return Rational{
			N: r.N + s.N,
			D: r.D,
		}.Simplify()
	}

	return Rational{
		N: r.N*s.D + s.N*r.D,
		D: r.D * s.D,
	}.Simplify()
}

// Sub subtracts `s` from the receiver
func (r Rational) Sub(s Rational) Rational {
	if r.D == s.D {
		return Rational{
			N: r.N - s.N,
			D: r.D,
		}.Simplify()
	}

	return Rational{
		N: r.N*s.D - s.N*r.D,
		D: r.D * s.D,
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
