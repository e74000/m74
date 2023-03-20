package m74

func Factorial[T Natural](v T) T {
	var prod T = 1

	for i := T(2); i < v; i++ {
		prod *= i
	}

	return prod
}

func NCR[T Natural](n, r T) T {
	if r > n {
		panic("n must be greater than r")
	}

	var (
		num T = 1
		den T = 1
	)

	if r > n-r {
		r = n - r
	}

	for i := T(0); i < r; i++ {
		num *= n - i
		den *= r - i
	}

	return num / den
}

func GCD[T Integer](a, b T) T {
	var (
		p, q = a, b
	)

	for q != 0 {
		p, q = q, p%q
	}

	return Max(p, -p)
}

func Mod[T Integer](n, d T) T {
	return (n%d + d) % d
}
