package math

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		inputs   [][]float64
		expected []float64
	}{
		{
			name:     "empty input",
			inputs:   [][]float64{},
			expected: nil,
		},
		{
			name:     "single array",
			inputs:   [][]float64{{1, 2, 3}},
			expected: []float64{1, 2, 3},
		},
		{
			name:     "multiple arrays",
			inputs:   [][]float64{{1, 2, 3}, {4, 5, 6}},
			expected: []float64{5, 7, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.inputs...)
			if !compareSlices(result, tt.expected) {
				t.Errorf("Add() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		name     string
		inputs   [][]float64
		expected []float64
	}{
		{
			name:     "insufficient inputs",
			inputs:   [][]float64{{1, 2, 3}},
			expected: nil,
		},
		{
			name:     "division by zero",
			inputs:   [][]float64{{1, 2, 3}, {0, 1, 2}},
			expected: []float64{1, 2, 1.5},
		},
		{
			name:     "normal division",
			inputs:   [][]float64{{6, 8, 10}, {2, 2, 2}},
			expected: []float64{3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Div(tt.inputs...)
			if !compareSlices(result, tt.expected) {
				t.Errorf("Div() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name       string
		input      []float64
		timeperiod int
		expected   float64
	}{
		{
			name:       "empty input",
			input:      []float64{},
			timeperiod: 3,
			expected:   0,
		},
		{
			name:       "normal case",
			input:      []float64{1, 5, 3, 8, 2},
			timeperiod: 3,
			expected:   5,
		},
		{
			name:       "period larger than input",
			input:      []float64{1, 2, 3},
			timeperiod: 5,
			expected:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Max(tt.input, tt.timeperiod)
			if result != tt.expected {
				t.Errorf("Max() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name       string
		input      []float64
		timeperiod int
		expected   []float64
	}{
		{
			name:       "empty input",
			input:      []float64{},
			timeperiod: 3,
			expected:   nil,
		},
		{
			name:       "normal case",
			input:      []float64{5, 2, 8, 1, 9},
			timeperiod: 3,
			expected:   []float64{2, 2, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Min(tt.input, tt.timeperiod)
			if !compareSlices(result, tt.expected) {
				t.Errorf("Min() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMinMax(t *testing.T) {
	tests := []struct {
		name       string
		input      []float64
		timeperiod int
		expected   struct {
			min float64
			max float64
		}
	}{
		{
			name:       "empty input",
			input:      []float64{},
			timeperiod: 3,
			expected: struct {
				min float64
				max float64
			}{0, 0},
		},
		{
			name:       "normal case",
			input:      []float64{5, 2, 8, 1, 9},
			timeperiod: 3,
			expected: struct {
				min float64
				max float64
			}{2, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			min, max := MinMax(tt.input, tt.timeperiod)
			if min != tt.expected.min || max != tt.expected.max {
				t.Errorf("MinMax() = (%v, %v), want (%v, %v)", min, max, tt.expected.min, tt.expected.max)
			}
		})
	}
}

func TestMult(t *testing.T) {
	tests := []struct {
		name     string
		input0   []float64
		input1   []float64
		expected []float64
	}{
		{
			name:     "mismatched lengths",
			input0:   []float64{1, 2, 3},
			input1:   []float64{1, 2},
			expected: nil,
		},
		{
			name:     "normal multiplication",
			input0:   []float64{1, 2, 3},
			input1:   []float64{2, 3, 4},
			expected: []float64{2, 6, 12},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Mult(tt.input0, tt.input1)
			if !compareSlices(result, tt.expected) {
				t.Errorf("Mult() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		input0   []float64
		input1   []float64
		expected []float64
	}{
		{
			name:     "mismatched lengths",
			input0:   []float64{1, 2, 3},
			input1:   []float64{1, 2},
			expected: nil,
		},
		{
			name:     "normal addition",
			input0:   []float64{1, 2, 3},
			input1:   []float64{2, 3, 4},
			expected: []float64{3, 5, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sum(tt.input0, tt.input1)
			if !compareSlices(result, tt.expected) {
				t.Errorf("Sum() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCumulativeSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected []float64
	}{
		{
			name:     "empty input",
			input:    []float64{},
			expected: []float64{},
		},
		{
			name:     "normal case",
			input:    []float64{1, 2, 3, 4},
			expected: []float64{1, 3, 6, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CumulativeSum(tt.input)
			if !compareSlices(result, tt.expected) {
				t.Errorf("CumulativeSum() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDrawdown(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected []float64
	}{
		{
			name:     "empty input",
			input:    []float64{},
			expected: []float64{},
		},
		{
			name:     "normal case",
			input:    []float64{100, 90, 95, 80},
			expected: []float64{0, 0.1, 0.05, 0.2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Drawdown(tt.input)
			if !compareSlices(result, tt.expected) {
				t.Errorf("Drawdown() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMovingAverage(t *testing.T) {
	tests := []struct {
		name       string
		input      []float64
		timeperiod int
		expected   []float64
	}{
		{
			name:       "invalid period",
			input:      []float64{1, 2, 3},
			timeperiod: 0,
			expected:   nil,
		},
		{
			name:       "normal case",
			input:      []float64{1, 2, 3, 4, 5},
			timeperiod: 3,
			expected:   []float64{2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MovingAverage(tt.input, tt.timeperiod)
			if !compareSlices(result, tt.expected) {
				t.Errorf("MovingAverage() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMean(t *testing.T) {
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
			input:    []float64{5},
			expected: 5,
		},
		{
			name:     "normal case",
			input:    []float64{1, 2, 3, 4, 5},
			expected: 3,
		},
		{
			name:     "negative values",
			input:    []float64{-1, -2, -3, -4, -5},
			expected: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Mean(tt.input)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("Mean() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAverage(t *testing.T) {
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
			input:    []float64{5},
			expected: 5,
		},
		{
			name:     "normal case",
			input:    []float64{1, 2, 3, 4, 5},
			expected: 3,
		},
		{
			name:     "negative values",
			input:    []float64{-1, -2, -3, -4, -5},
			expected: -3,
		},
		{
			name:     "decimal values",
			input:    []float64{1.5, 2.5, 3.5, 4.5},
			expected: 3,
		},
		{
			name:     "mixed positive and negative",
			input:    []float64{-2, -1, 0, 1, 2},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Average(tt.input)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("Average() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMedian(t *testing.T) {
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
			input:    []float64{5},
			expected: 5,
		},
		{
			name:     "odd length",
			input:    []float64{1, 3, 2, 5, 4},
			expected: 3,
		},
		{
			name:     "even length",
			input:    []float64{1, 2, 3, 4},
			expected: 2.5,
		},
		{
			name:     "unsorted input",
			input:    []float64{5, 2, 8, 1, 9},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Median(tt.input)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("Median() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMode(t *testing.T) {
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
			input:    []float64{5},
			expected: 5,
		},
		{
			name:     "single mode",
			input:    []float64{1, 2, 2, 3, 4},
			expected: 2,
		},
		{
			name:     "multiple modes",
			input:    []float64{1, 2, 2, 3, 3, 4},
			expected: 2,
		},
		{
			name:     "all unique values",
			input:    []float64{1, 2, 3, 4, 5},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Mode(tt.input)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("Mode() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestRange(t *testing.T) {
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
			input:    []float64{5},
			expected: 0,
		},
		{
			name:     "normal case",
			input:    []float64{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "negative values",
			input:    []float64{-5, -3, -1, 1, 3},
			expected: 8,
		},
		{
			name:     "mixed values",
			input:    []float64{-2, 0, 2, 4, 6},
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Range(tt.input)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("Range() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestVariance(t *testing.T) {
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
			input:    []float64{5},
			expected: 0,
		},
		{
			name:     "constant values",
			input:    []float64{2, 2, 2, 2, 2},
			expected: 0,
		},
		{
			name:     "normal case",
			input:    []float64{2, 4, 4, 4, 5, 5, 7, 9},
			expected: 4,
		},
		{
			name:     "negative values",
			input:    []float64{-2, -4, -4, -4, -5, -5, -7, -9},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Variance(tt.input)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("Variance() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Helper function to compare float64 slices with tolerance
func compareSlices(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > 1e-10 {
			return false
		}
	}
	return true
}
