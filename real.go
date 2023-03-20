package m74

import "math"

func ModR[T Real](n, d T) T {
	return T(float64(n) - math.Floor(float64(n)/float64(d))*float64(d))
}
