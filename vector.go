package m74

import (
	"math"
)

func ConvVector[To, From Real](v Vec2[From]) Vec2[To] {
	return Vec2[To]{
		X: To(v.X),
		Y: To(v.Y),
	}
}

func VecLerp[T Real](x float64, v, w Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: Lerp(x, v.X, w.X),
		Y: Lerp(x, v.Y, w.Y),
	}
}

type Vec2[T Real] struct {
	X, Y T
}

func (v Vec2[T]) Mag() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func (v Vec2[T]) Add(w Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v.X + w.X,
		Y: v.Y + w.Y,
	}
}

func (v Vec2[T]) Sub(w Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v.X - w.X,
		Y: v.Y - w.Y,
	}
}

func (v Vec2[T]) Dot(w Vec2[T]) T {
	return v.X*w.X + v.Y*w.Y
}

func (v Vec2[T]) MulElem(w Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v.X * w.X,
		Y: v.Y * w.Y,
	}
}

func (v Vec2[T]) DivElem(w Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v.X / w.X,
		Y: v.Y / w.Y,
	}
}

func (v Vec2[T]) Mul(s T) Vec2[T] {
	return Vec2[T]{
		X: v.X * s,
		Y: v.Y * s,
	}
}

func (v Vec2[T]) Div(s T) Vec2[T] {
	return Vec2[T]{
		X: v.X / s,
		Y: v.Y / s,
	}
}
