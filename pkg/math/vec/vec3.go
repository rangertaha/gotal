package vec

import "math"


type Vec3[T Number] struct {
	X, Y, Z T
}

func (v Vec3[T]) Add(u Vec3[T]) Vec3[T] {
	return Vec3[T]{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

func (v Vec3[T]) Sub(u Vec3[T]) Vec3[T] {
	return Vec3[T]{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

func (v Vec3[T]) Scale(s T) Vec3[T] {
	return Vec3[T]{v.X * s, v.Y * s, v.Z * s}
}

func (v Vec3[T]) Dot(u Vec3[T]) T {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

func (v Vec3[T]) Cross(u Vec3[T]) Vec3[T] {
	return Vec3[T]{
		v.Y*u.Z - v.Z*u.Y,
		v.Z*u.X - v.X*u.Z,
		v.X*u.Y - v.Y*u.X,
	}
}

func (v Vec3[T]) Len() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
}

func (v Vec3[T]) Normalize() Vec3[T] {
	l := v.Len()
	if l == 0 {
		panic("cannot normalize zero vector")
	}
	return Vec3[T]{
		T(float64(v.X) / l),
		T(float64(v.Y) / l),
		T(float64(v.Z) / l),
	}
}
