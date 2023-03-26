package m74

import (
	"math"
)

// ConvVector converts a vector to a different type
func ConvVector[To, From Real](v Vec2[From]) Vec2[To] {
	return Vec2[To]{
		X: To(v.X),
		Y: To(v.Y),
	}
}

// VecLerp interpolates between two vectors
func VecLerp[T Real](x float64, v, w Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: Lerp(x, v.X, w.X),
		Y: Lerp(x, v.Y, w.Y),
	}
}

// Vec2 is 2 dimensional vector
type Vec2[T Real] struct {
	X, Y T
}

// Mag returns the magnitude of the receiver
func (v Vec2[T]) Mag() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

// Add adds `w` to the receiver
func (v Vec2[T]) Add(w Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v.X + w.X,
		Y: v.Y + w.Y,
	}
}

// Sub subtracts `w` from the receiver
func (v Vec2[T]) Sub(w Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v.X - w.X,
		Y: v.Y - w.Y,
	}
}

// Dot performs the dot product of the receiver and `w`
func (v Vec2[T]) Dot(w Vec2[T]) T {
	return v.X*w.X + v.Y*w.Y
}

// MulElem performs an element-wise multiplication of the receiver and `w`
func (v Vec2[T]) MulElem(w Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v.X * w.X,
		Y: v.Y * w.Y,
	}
}

// DivElem performs an element-wise division of the receiver and `w`
func (v Vec2[T]) DivElem(w Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v.X / w.X,
		Y: v.Y / w.Y,
	}
}

// Mul multiplies each element of a vector by `s`
func (v Vec2[T]) Mul(s T) Vec2[T] {
	return Vec2[T]{
		X: v.X * s,
		Y: v.Y * s,
	}
}

// Div divides each element of a vector by `s`
func (v Vec2[T]) Div(s T) Vec2[T] {
	return Vec2[T]{
		X: v.X / s,
		Y: v.Y / s,
	}
}
