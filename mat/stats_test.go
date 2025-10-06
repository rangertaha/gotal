package math

import (
	"math"
	"testing"
)

func TestBeta(t *testing.T) {
	tests := []struct {
		name     string
		x        []float64
		y        []float64
		expected float64
	}{
		{
			name:     "empty input",
			x:        []float64{},
			y:        []float64{},
			expected: 0,
		},
		{
			name:     "mismatched lengths",
			x:        []float64{1, 2, 3},
			y:        []float64{1, 2},
			expected: 0,
		},
		{
			name:     "perfect correlation",
			x:        []float64{1, 2, 3, 4, 5},
			y:        []float64{2, 4, 6, 8, 10},
			expected: 2,
		},
		{
			name:     "negative correlation",
			x:        []float64{1, 2, 3, 4, 5},
			y:        []float64{10, 8, 6, 4, 2},
			expected: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Beta(tt.x, tt.y)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("Beta() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLinearRegression(t *testing.T) {
	tests := []struct {
		name              string
		x                 []float64
		y                 []float64
		expectedSlope     float64
		expectedIntercept float64
	}{
		{
			name:              "empty input",
			x:                 []float64{},
			y:                 []float64{},
			expectedSlope:     0,
			expectedIntercept: 0,
		},
		{
			name:              "mismatched lengths",
			x:                 []float64{1, 2, 3},
			y:                 []float64{1, 2},
			expectedSlope:     0,
			expectedIntercept: 0,
		},
		{
			name:              "perfect linear relationship",
			x:                 []float64{1, 2, 3, 4, 5},
			y:                 []float64{2, 4, 6, 8, 10},
			expectedSlope:     2,
			expectedIntercept: 0,
		},
		{
			name:              "offset linear relationship",
			x:                 []float64{1, 2, 3, 4, 5},
			y:                 []float64{3, 5, 7, 9, 11},
			expectedSlope:     2,
			expectedIntercept: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slope, intercept := LinearRegression(tt.x, tt.y)
			if math.Abs(slope-tt.expectedSlope) > 1e-10 || math.Abs(intercept-tt.expectedIntercept) > 1e-10 {
				t.Errorf("LinearRegression() = (%v, %v), want (%v, %v)",
					slope, intercept, tt.expectedSlope, tt.expectedIntercept)
			}
		})
	}
}

func TestLinearRegressionAngle(t *testing.T) {
	tests := []struct {
		name     string
		x        []float64
		y        []float64
		expected float64
	}{
		{
			name:     "empty input",
			x:        []float64{},
			y:        []float64{},
			expected: 0,
		},
		{
			name:     "45 degree angle",
			x:        []float64{1, 2, 3, 4, 5},
			y:        []float64{1, 2, 3, 4, 5},
			expected: 45,
		},
		{
			name:     "negative angle",
			x:        []float64{1, 2, 3, 4, 5},
			y:        []float64{5, 4, 3, 2, 1},
			expected: -45,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LinearRegressionAngle(tt.x, tt.y)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("LinearRegressionAngle() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLinearRegressionIntercept(t *testing.T) {
	tests := []struct {
		name     string
		x        []float64
		y        []float64
		expected float64
	}{
		{
			name:     "empty input",
			x:        []float64{},
			y:        []float64{},
			expected: 0,
		},
		{
			name:     "zero intercept",
			x:        []float64{1, 2, 3, 4, 5},
			y:        []float64{2, 4, 6, 8, 10},
			expected: 0,
		},
		{
			name:     "positive intercept",
			x:        []float64{1, 2, 3, 4, 5},
			y:        []float64{3, 5, 7, 9, 11},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LinearRegressionIntercept(tt.x, tt.y)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("LinearRegressionIntercept() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLinearRegressionSlope(t *testing.T) {
	tests := []struct {
		name     string
		x        []float64
		y        []float64
		expected float64
	}{
		{
			name:     "empty input",
			x:        []float64{},
			y:        []float64{},
			expected: 0,
		},
		{
			name:     "positive slope",
			x:        []float64{1, 2, 3, 4, 5},
			y:        []float64{2, 4, 6, 8, 10},
			expected: 2,
		},
		{
			name:     "negative slope",
			x:        []float64{1, 2, 3, 4, 5},
			y:        []float64{10, 8, 6, 4, 2},
			expected: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LinearRegressionSlope(tt.x, tt.y)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("LinearRegressionSlope() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestTimeSeriesForecast(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		periods  int
		expected []float64
	}{
		{
			name:     "empty input",
			input:    []float64{},
			periods:  3,
			expected: nil,
		},
		{
			name:     "zero periods",
			input:    []float64{1, 2, 3, 4, 5},
			periods:  0,
			expected: nil,
		},
		{
			name:     "linear trend",
			input:    []float64{1, 2, 3, 4, 5},
			periods:  3,
			expected: []float64{6, 7, 8},
		},
		{
			name:     "negative trend",
			input:    []float64{5, 4, 3, 2, 1},
			periods:  3,
			expected: []float64{0, -1, -2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TimeSeriesForecast(tt.input, tt.periods)
			if !compareSlices(result, tt.expected) {
				t.Errorf("TimeSeriesForecast() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestStandardDeviation(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{
			name:     "empty input",
			input:    []float64{},
			expected: 0,
		},
		{
			name:     "single value",
			input:    []float64{1},
			expected: 0,
		},
		{
			name:     "normal case",
			input:    []float64{2, 4, 4, 4, 5, 5, 7, 9},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StandardDeviation(tt.input)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("StdDev() = %v, want %v", result, tt.expected)
			}
		})
	}
}


func TestCorrelation(t *testing.T) {
	tests := []struct {
		name     string
		x        []float64
		y        []float64
		expected float64
	}{
		{
			name:     "empty input",
			x:        []float64{},
			y:        []float64{},
			expected: 0,
		},
		{
			name:     "perfect correlation",
			x:        []float64{1, 2, 3, 4, 5},
			y:        []float64{2, 4, 6, 8, 10},
			expected: 1,
		},
		{
			name:     "perfect negative correlation",
			x:        []float64{1, 2, 3, 4, 5},
			y:        []float64{10, 8, 6, 4, 2},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Correlation(tt.x, tt.y)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("Correlation() = %v, want %v", result, tt.expected)
			}
		})
	}
}

