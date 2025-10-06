package math

// Add computes vector arithmetic addition of multiple float64 slices.
// Returns nil if no inputs are provided.
// All input slices must have the same length.
func Add(arrs ...[]float64) []float64 {
	if len(arrs) == 0 {
		return nil
	}
	result := make([]float64, len(arrs[0]))
	for _, arr := range arrs {
		for i, v := range arr {
			result[i] += v
		}
	}
	return result
}

// Div computes vector arithmetic division of multiple float64 slices.
// Returns nil if less than 2 inputs are provided.
// All input slices must have the same length.
// Division by zero is handled by preserving the original value.
func Div(arrs ...[]float64) []float64 {
	if len(arrs) < 2 {
		return nil
	}
	result := make([]float64, len(arrs[0]))
	copy(result, arrs[0])
	for i := 1; i < len(arrs); i++ {
		for j, v := range arrs[i] {
			if v != 0 {
				result[j] /= v
			}
		}
	}
	return result
}

// Max returns the highest value over a specified period in the input array.
// Returns 0 if the input is empty or timeperiod is invalid.
// The timeperiod must be positive and not exceed the input length.
func Max(arr []float64, timeperiod int) float64 {
	if len(arr) == 0 || timeperiod <= 0 {
		return 0
	}
	max := arr[0]
	for i := 1; i < timeperiod && i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// MaxIndex returns the index of the highest value over a specified period.
// Returns -1 if the input is empty or timeperiod is invalid.
// The timeperiod must be positive and not exceed the input length.
func MaxIndex(arr []float64, timeperiod int) int {
	if len(arr) == 0 || timeperiod <= 0 {
		return -1
	}
	maxIdx := 0
	for i := 1; i < timeperiod && i < len(arr); i++ {
		if arr[i] > arr[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}

// Min returns an array of the lowest values over a specified period for each position.
// Returns nil if the input is empty or timeperiod is invalid.
// The timeperiod must be positive and not exceed the input length.
func Min(arr []float64, timeperiod int) []float64 {
	if len(arr) == 0 || timeperiod <= 0 {
		return nil
	}
	result := make([]float64, len(arr))
	for i := 0; i < len(arr); i++ {
		end := i + timeperiod
		if end > len(arr) {
			end = len(arr)
		}
		min := arr[i]
		for j := i + 1; j < end; j++ {
			if arr[j] < min {
				min = arr[j]
			}
		}
		result[i] = min
	}
	return result
}

// MinIndex returns the index of the lowest value over a specified period.
// Returns -1 if the input is empty or timeperiod is invalid.
// The timeperiod must be positive and not exceed the input length.
func MinIndex(arr []float64, timeperiod int) int {
	if len(arr) == 0 || timeperiod <= 0 {
		return -1
	}
	minIdx := 0
	for i := 1; i < timeperiod && i < len(arr); i++ {
		if arr[i] < arr[minIdx] {
			minIdx = i
		}
	}
	return minIdx
}

// MinMax returns both the lowest and highest values over a specified period.
// Returns (0, 0) if the input is empty or timeperiod is invalid.
// The timeperiod must be positive and not exceed the input length.
func MinMax(arr []float64, timeperiod int) (float64, float64) {
	if len(arr) == 0 || timeperiod <= 0 {
		return 0, 0
	}
	min, max := arr[0], arr[0]
	for i := 1; i < timeperiod && i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	return min, max
}

// MinMaxIndex returns the indexes of both the lowest and highest values over a specified period.
// Returns (-1, -1) if the input is empty or timeperiod is invalid.
// The timeperiod must be positive and not exceed the input length.
func MinMaxIndex(arr []float64, timeperiod int) (int, int) {
	if len(arr) == 0 || timeperiod <= 0 {
		return -1, -1
	}
	minIdx, maxIdx := 0, 0
	for i := 1; i < timeperiod && i < len(arr); i++ {
		if arr[i] < arr[minIdx] {
			minIdx = i
		}
		if arr[i] > arr[maxIdx] {
			maxIdx = i
		}
	}
	return minIdx, maxIdx
}

// Mult computes vector arithmetic multiplication of two float64 slices.
// Returns nil if the input slices have different lengths.
func Mult(arr0 []float64, arr1 []float64) []float64 {
	if len(arr0) != len(arr1) {
		return nil
	}
	result := make([]float64, len(arr0))
	for i := range arr0 {
		result[i] = arr0[i] * arr1[i]
	}
	return result
}

// Sum computes vector arithmetic addition of two float64 slices.
// Returns nil if the input slices have different lengths.
func Sum(x []float64, y []float64) []float64 {
	if len(x) != len(y) {
		return nil
	}
	result := make([]float64, len(x))
	for i := range x {
		result[i] = x[i] + y[i]
	}
	return result
}

// CumulativeSum returns the running total (cumulative sum) of the input array.
// Returns an empty slice if the input is empty.
func CumulativeSum(arr []float64) []float64 {
	result := make([]float64, len(arr))
	var sum float64
	for i, v := range arr {
		sum += v
		result[i] = sum
	}
	return result
}

// Drawdown returns the drawdown at each point in the equity curve.
// Drawdown is calculated as (peak - current) / peak.
// Returns an empty slice if the input is empty.
func Drawdown(equity []float64) []float64 {
	result := make([]float64, len(equity))
	var peak float64
	for i, v := range equity {
		if v > peak {
			peak = v
		}
		if peak == 0 {
			result[i] = 0
		} else {
			result[i] = (peak - v) / peak
		}
	}
	return result
}

// MovingAverage returns the simple moving average of the input array with the given window size.
// Returns nil if the timeperiod is invalid (<= 0 or > input length).
// The result length is len(input) - timeperiod + 1.
func MovingAverage(arr []float64, timeperiod int) []float64 {
	if timeperiod <= 0 || timeperiod > len(arr) {
		return nil
	}
	result := make([]float64, len(arr)-timeperiod+1)
	var sum float64
	for i := 0; i < timeperiod; i++ {
		sum += arr[i]
	}
	result[0] = sum / float64(timeperiod)
	for i := timeperiod; i < len(arr); i++ {
		sum += arr[i] - arr[i-timeperiod]
		result[i-timeperiod+1] = sum / float64(timeperiod)
	}
	return result
}

// Histogram returns a map of value counts for the input array.
// The map keys are the unique values in the input array.
// The map values are the frequency of each value.
func Histogram(arr []float64) map[float64]int {
	hist := make(map[float64]int)
	for _, v := range arr {
		hist[v]++
	}
	return hist
}

// Mean returns the arithmetic mean (average) of the input array.
// Returns 0 if the input is empty.
func Mean(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}
	var sum float64
	for _, v := range arr {
		sum += v
	}
	return sum / float64(len(arr))
}

// Average returns the arithmetic mean (average) of the input array.
// Returns 0 if the input is empty.
// This is an alias for Mean() for clarity and consistency with common terminology.
func Average(arr []float64) float64 {
	return Mean(arr)
}

// Median returns the median value of the input array.
// Returns 0 if the input is empty.
// For even-length arrays, returns the average of the two middle values.
func Median(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}

	// Create a copy to avoid modifying the input
	sorted := make([]float64, len(arr))
	copy(sorted, arr)

	// Simple bubble sort (could be optimized for larger arrays)
	for i := 0; i < len(sorted)-1; i++ {
		for j := 0; j < len(sorted)-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2
	}
	return sorted[mid]
}

// Mode returns the most frequently occurring value in the input array.
// Returns 0 if the input is empty.
// If multiple values have the same frequency, returns the smallest value.
func Mode(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}

	freq := make(map[float64]int)
	for _, v := range arr {
		freq[v]++
	}

	var mode float64
	maxFreq := 0
	for v, f := range freq {
		if f > maxFreq || (f == maxFreq && v < mode) {
			mode = v
			maxFreq = f
		}
	}
	return mode
}

// Range returns the difference between the maximum and minimum values in the input array.
// Returns 0 if the input is empty.
func Range(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}

	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}
