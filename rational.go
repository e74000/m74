package m74

// Rational represents fraction (n/d)
type Rational struct {
	n int64
	d int64
}

// NewRational returns a rational
func NewRational[T Integer](n, d T) Rational {
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

// Positive returns true if the rational is positive or zero
func (r Rational) Positive() bool {
	return (r.n >= 0 && r.d >= 0) || (r.n < 0 && r.d < 0)
}

// Negative returns true if the rational is negative or zero
func (r Rational) Negative() bool {
	return (r.n < 0 && r.d >= 0) || (r.n >= 0 && r.d < 0)
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

	if s.d == 1 {
		return int(s.n), true
	} else {
		return 0, false
	}
}

func (r Rational) PosPow(n int) Rational {
	if n < 0 {
		panic("PosPow (Positive Power) only works for positive inputs.")
	}

	out := NewRational(1, 1)

	for i := 0; i < n; i++ {
		out = out.Mul(r)
	}

	return out.Simplify()
}

// Abs gets the absolute value of a rational, by multiplying it by -1 if negative
func (r Rational) Abs() Rational {
	return Rational{
		n: Abs(r.n),
		d: Abs(r.d),
	}
}

// Simplify removes any common divisors from the numerator/denominator
func (r Rational) Simplify() Rational {
	gcd := GCD(r.n, r.d)

	if gcd == 1 {
		return r
	}

	return Rational{
		n: r.n / gcd,
		d: r.d / gcd,
	}
}

// Mul multiplies the receiver by `s`
func (r Rational) Mul(s Rational) Rational {
	return Rational{
		n: r.n * s.n,
		d: r.d * s.d,
	}.Simplify()
}

// Div divides the receiver by `s`
func (r Rational) Div(s Rational) Rational {
	return Rational{
		n: r.n * s.d,
		d: r.d * s.n,
	}.Simplify()
}

// Add adds `s` to the receiver
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

// Sub subtracts `s` from the receiver
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
