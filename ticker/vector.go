package ticker

import (
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
)

type vector []float64

func (v vector) Values() []float64 {
	return []float64(v)
}

func (v vector) Vec() *mat.VecDense {
	return mat.NewVecDense(len(v), v)
}

// Mode calculates the most common value of the specified field across all ticks.
func (v vector) Mode(weights []float64) (val float64, count float64) {
	return stat.Mode(v, weights)
}

// Mean calculates the arithmetic mean of the specified field across all ticks.
func (v vector) Mean() (output float64) {
	return mat.Sum(v.Vec()) / float64(len(v))
}

// Median calculates the median value of the specified field across all ticks.
func (v vector) Median() (output float64) {
	return stat.Quantile(0.5, stat.Empirical, v, nil)
}

// Range returns the difference between the maximum and minimum values
func (v vector) Range() (output float64) {
	return mat.Max(v.Vec()) - mat.Min(v.Vec())
}

// Sum calculates the total sum of the specified field across all ticks.
func (v vector) Sum() (output float64) {
	return mat.Sum(v.Vec())
}

// Min returns the minimum value of the specified field across all ticks.
func (v vector) Min() (output float64) {
	return mat.Min(v.Vec())
}

// Max returns the maximum value of the specified field across all ticks.
func (v vector) Max() (output float64) {
	return mat.Max(v.Vec())
}

// First returns the first value of the specified field.
func (v vector) First() (output float64) {
	return v[0]
}

// Last returns the last value of the specified field.
func (v vector) Last() (output float64) {
	return v[len(v)-1]
}

// Std calculates the standard deviation of the specified field across all ticks.
func (v vector) StdDev(weights ...float64) (output float64) {
	return stat.StdDev(v, weights)
}

// Var calculates the variance of the specified field across all ticks.
func (v vector) Variance(weights ...float64) (output float64) {
	return stat.Variance(v, weights)
}

// Norm calculates the L1 norm (sum of absolute values) of the specified field across all ticks.
func (v vector) Norm(norm float64) (output float64) {
	return mat.Norm(v.Vec(), norm)
}

// Quantile calculates the quantile value of the specified field across all ticks.
func (v vector) Quantile(p float64, c stat.CumulantKind, weights []float64) float64 {
	return stat.Quantile(p, c, v, weights)
}
