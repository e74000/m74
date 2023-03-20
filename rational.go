package m74

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

	if gcd == 1 {
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
