package m74

// Factorial performs an integer factorial of `v`
func Factorial[T Natural](v T) T {
	var prod T = 1

	for i := T(2); i < v; i++ {
		prod *= i
	}

	return prod
}

// NCR performs `n` choose `r`, or the number of ways you could select `r` items from a set of `n`
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

// GCD (Greatest Common Divisor) returns the largest integer divisor of `a` and `b`
func GCD[T Integer](a, b T) T {
	var (
		p, q = a, b
	)

	for q != 0 {
		p, q = q, p%q
	}

	return Max(p, -p)
}

// Mod is the modulo function
func Mod[T Integer](n, d T) T {
	return (n%d + d) % d
}
