package series

import (
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
)

// Statistical methods

// Vec creates a dense vector from the specified field values.
// Returns a VecDense containing the field values in chronological order.
func (s *Series) Vec(field string) *mat.VecDense {
	return mat.NewVecDense(s.Len(), s.Field(field))
}

// Vecs creates a map of field names to dense vectors from the specified fields.
// Returns a map where the keys are field names and the values are VecDense containing the field values in chronological order.
func (s *Series) VecMap(fields ...string) map[string]*mat.VecDense {
	vecs := make(map[string]*mat.VecDense)
	for _, field := range fields {
		if s.HasField(field) {
			vecs[field] = s.Vec(field)
		}
	}
	return vecs
}

// Mode calculates the most common value of the specified field across all ticks.
func (s *Series) Mode(field string, weights []float64) (val float64, count float64) {
	return stat.Mode(s.Field(field), weights)
}

// Mean calculates the arithmetic mean of the specified field across all ticks.
func (s *Series) Mean(field string) (output float64) {
	return mat.Sum(s.Vec(field)) / float64(s.Len())
}

// Median calculates the median value of the specified field across all ticks.
func (s *Series) Median(field string) (output float64) {
	return stat.Quantile(0.5, stat.Empirical, s.Field(field), nil)
}

// Range returns the difference between the maximum and minimum values
func (s *Series) Range(field string) (output float64) {
	return mat.Max(s.Vec(field)) - mat.Min(s.Vec(field))
}

// Sum calculates the total sum of the specified field across all ticks.
func (s *Series) Sum(field string) (output float64) {
	return mat.Sum(s.Vec(field))
}

// Min returns the minimum value of the specified field across all ticks.
func (s *Series) Min(field string) (output float64) {
	return mat.Min(s.Vec(field))
}

// Max returns the maximum value of the specified field across all ticks.
func (s *Series) Max(field string) (output float64) {
	return mat.Max(s.Vec(field))
}

// First returns the first value of the specified field.
func (s *Series) First(field string) (output float64) {
	return s.Vec(field).At(0, 0)
}

// Last returns the last value of the specified field.
func (s *Series) Last(field string) (output float64) {
	return s.Vec(field).At(s.Vec(field).Len()-1, 0)
}

// Std calculates the standard deviation of the specified field across all ticks.
func (s *Series) Std(field string) (output float64) {
	return stat.StdDev(s.Field(field), nil)
}

// Var calculates the variance of the specified field across all ticks.
func (s *Series) Var(field string) (output float64) {
	return stat.Variance(s.Field(field), nil)
}

// Norm calculates the L1 norm (sum of absolute values) of the specified field across all ticks.
func (s *Series) Norm(field string) (output float64) {
	return mat.Norm(s.Vec(field), 1)
}

// Quantile calculates the quantile value of the specified field across all ticks.
func (s *Series) Quantile(field string) func(p float64, c stat.CumulantKind, weights []float64) float64 {
	return func(p float64, c stat.CumulantKind, weights []float64) float64 {
		return stat.Quantile(p, c, s.Field(field), weights)
	}
}
