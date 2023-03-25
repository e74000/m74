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

func VecInvLerp[T Real](v, a, b Vec2[T]) float64 {
	if !Colinear(v, a, b) {
		return -1
	}

	if b.X-a.X != 0 {
		return float64(v.X-a.X) / float64(b.X-a.X)
	} else if b.Y-a.Y != 0 {
		return float64(v.Y-a.Y) / float64(b.Y-a.Y)
	}

	return 0
}

// VecMapRange can be treated as VecLerp(VecInvLerp(v, a0, b0), a1, b1)
func VecMapRange[T Real](v, a0, b0, a1, b1 Vec2[T]) Vec2[T] {
	return VecLerp(VecInvLerp(v, a0, b0), a1, b1)
}

func TriangleArea[T Real](a, b, c Vec2[T]) T {
	return (a.X*(b.Y-c.Y) + b.X*(c.Y-a.Y) + c.X*(a.Y-b.Y)) / 2
}

func Colinear[T Real](vs ...Vec2[T]) bool {
	if len(vs) <= 2 {
		return true
	}

	for i := 2; i < len(vs); i++ {
		if TriangleArea(vs[0], vs[1], vs[i]) != 0 {
			return false
		}
	}

	return true
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

type Rectangle[T Real] struct {
	Min, Max Vec2[T]
}

func Rect[T Real](x0, y0, x1, y1 T) Rectangle[T] {
	return Rectangle[T]{
		Min: Vec2[T]{x0, y0},
		Max: Vec2[T]{x1, y1},
	}
}

func (r Rectangle[T]) PointIn(v Vec2[T]) bool {
	offset := v.Sub(r.Min)

	if offset.X < r.Max.X && offset.Y < r.Max.Y {
		return true
	} else {
		return false
	}
}
