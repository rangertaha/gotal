package vec

import "math"

type Vec2[T Number] struct {
	X, Y T
}

func (v Vec2[T]) Add(u Vec2[T]) Vec2[T] { return Vec2[T]{v.X + u.X, v.Y + u.Y} }
func (v Vec2[T]) Sub(u Vec2[T]) Vec2[T] { return Vec2[T]{v.X - u.X, v.Y - u.Y} }
func (v Vec2[T]) Scale(s T) Vec2[T]     { return Vec2[T]{v.X * s, v.Y * s} }

func (v Vec2[T]) Dot(u Vec2[T]) T { return v.X*u.X + v.Y*u.Y }

func (v Vec2[T]) Len() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func (v Vec2[T]) Normalize() Vec2[T] {
	l := v.Len()
	if l == 0 {
		panic("cannot normalize zero vector")
	}
	return Vec2[T]{T(float64(v.X) / l), T(float64(v.Y) / l)}
}
