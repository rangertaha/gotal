package vec

import "math"

// Constraint for numeric types allowed
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}


type Vec[T Number] []T

// NewVec creates a vector of length n initialized with zeroes
func NewVec[T Number](n int) Vec[T] {
	return make(Vec[T], n)
}
 
// Dim
func (v Vec[T]) Dim() int { return len(v) }

func (v Vec[T]) Add(u Vec[T]) Vec[T] {
	if len(v) != len(u) {
		panic("dimension mismatch")
	}
	r := make(Vec[T], len(v))
	for i := range v {
		r[i] = v[i] + u[i]
	}
	return r
}

func (v Vec[T]) Sub(u Vec[T]) Vec[T] {
	if len(v) != len(u) {
		panic("dimension mismatch")
	}
	r := make(Vec[T], len(v))
	for i := range v {
		r[i] = v[i] - u[i]
	}
	return r
}

func (v Vec[T]) Scale(s T) Vec[T] {
	r := make(Vec[T], len(v))
	for i := range v {
		r[i] = v[i] * s
	}
	return r
}

func (v Vec[T]) Dot(u Vec[T]) T {
	if len(v) != len(u) {
		panic("dimension mismatch")
	}
	var sum T
	for i := range v {
		sum += v[i] * u[i]
	}
	return sum
}

func (v Vec[T]) Len() float64 {
	var sum float64
	for _, x := range v {
		sum += float64(x * x)
	}
	return math.Sqrt(sum)
}

func (v Vec[T]) Normalize() Vec[T] {
	l := v.Len()
	if l == 0 {
		panic("cannot normalize zero vector")
	}
	r := make(Vec[T], len(v))
	for i := range v {
		r[i] = T(float64(v[i]) / l)
	}
	return r
}

// cosine_sim(u,v)=∥u∥∥v∥u⋅v​
func (v VecN[T]) CosineSimilarity(u VecN[T]) float64 {
	if len(v) != len(u) {
		panic("dimension mismatch")
	}
	dot := v.Dot(u)
	return float64(dot) / (v.Len() * u.Len())
}

// Hypotenuse calculates the length of the hypotenuse of a right-angled triangle,
// given the lengths of the other two sides (a and b), using the formula:
//
//	c = sqrt(a^2 + b^2)
//
// It uses functions from the standard [math package](
