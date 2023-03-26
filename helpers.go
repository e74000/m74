package m74

// Positive returns true if a value is positive or zero
func Positive[T Real](v T) bool {
	return v >= 0
}

// Negative returns true if a value if negative
func Negative[T Real](v T) bool {
	return v < 0
}

// Max returns the maximum of a set of values
func Max[T Real](vals ...T) (max T) {
	max = vals[0]
	idx := -1

	for i, v := range vals {
		if v > max {
			max = v
			idx = i
		}
	}

	if idx == -1 {
		return 0
	}

	return max
}

// Min returns the minimum of a set of values
func Min[T Real](vals ...T) (min T) {
	min = vals[0]
	idx := -1

	for i, v := range vals {
		if v < min {
			min = v
			idx = i
		}
	}

	if idx == -1 {
		return 0
	}

	return min
}

// MaxIdx returns the index/value of the maximum from a set of values
func MaxIdx[T Real](vals ...T) (max T, idx int) {
	max = vals[0]
	idx = -1

	for i, v := range vals {
		if v > max {
			max = v
			idx = i
		}
	}

	if idx == -1 {
		return 0, -1
	}

	return max, idx
}

// MinIdx returns the index/value of the minimum from a set of values
func MinIdx[T Real](vals ...T) (min T, idx int) {
	min = vals[0]
	idx = -1

	for i, v := range vals {
		if v < min {
			min = v
			idx = i
		}
	}

	if idx == -1 {
		return 0, -1
	}

	return min, idx
}

// Mean returns the mean value from a set of values
func Mean[T Real](vals ...T) float64 {
	var sum T

	for _, v := range vals {
		sum += v
	}

	return float64(sum) / float64(len(vals))
}

// Conv Is used to evade weirdness when it comes to returning -1 with unsigned integers
func Conv[To, From Real](v From) To {
	return To(v)
}

// Sign returns the sign (+/-) of the input
func Sign[T Real](v T) T {
	if v >= 0 {
		return 1
	} else {
		return Conv[T, int](-1)
	}
}

// Abs returns the absolute value of input
func Abs[T Real](v T) T {
	return v * Sign(v)
}

// Lerp linearly interpolates between a and b.
func Lerp[T Real](x float64, a, b T) T {
	return T(x*float64(b-a) + float64(a))
}

// InvLerp Performs the inverse of Lerp
func InvLerp[T Real](v, a, b T) float64 {
	return float64(v-a) / float64(b-a)
}

// MapRange can be treated as Lerp(InvLerp(v, a0, b0), a1, b1)
func MapRange[T Real](v, a0, b0, a1, b1 T) T {
	return T((float64(v-a0)/float64(b0-a0))*float64(b1-a1) + float64(a1))
}
