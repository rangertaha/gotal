package math

import "math"

// Beta calculates the beta coefficient between two arrays.
// Beta measures the volatility of an asset compared to a benchmark.
// Returns 0 if the arrays have different lengths or are empty.
func Beta(x, y []float64) float64 {
	if len(x) != len(y) || len(x) == 0 {
		return 0
	}
	covariance := 0.0
	variance := 0.0
	meanX := Mean(x)
	meanY := Mean(y)

	for i := 0; i < len(x); i++ {
		dx := x[i] - meanX
		dy := y[i] - meanY
		covariance += dx * dy
		variance += dx * dx
	}

	if variance == 0 {
		return 0
	}
	return covariance / variance
}

// LinearRegression performs a simple linear regression between two arrays.
// Returns the slope and intercept of the regression line.
// Returns (0, 0) if the arrays have different lengths or are empty.
func LinearRegression(x, y []float64) (slope, intercept float64) {
	if len(x) != len(y) || len(x) == 0 {
		return 0, 0
	}

	n := float64(len(x))
	sumX := 0.0
	sumY := 0.0
	sumXY := 0.0
	sumX2 := 0.0

	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}

	slope = (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	intercept = (sumY - slope*sumX) / n

	return slope, intercept
}

// LinearRegressionAngle returns the angle of the regression line in degrees.
// Returns 0 if the arrays have different lengths or are empty.
func LinearRegressionAngle(x, y []float64) float64 {
	slope, _ := LinearRegression(x, y)
	return math.Atan(slope) * 180 / math.Pi
}

// LinearRegressionIntercept returns the y-intercept of the regression line.
// Returns 0 if the arrays have different lengths or are empty.
func LinearRegressionIntercept(x, y []float64) float64 {
	_, intercept := LinearRegression(x, y)
	return intercept
}

// LinearRegressionSlope returns the slope of the regression line.
// Returns 0 if the arrays have different lengths or are empty.
func LinearRegressionSlope(x, y []float64) float64 {
	slope, _ := LinearRegression(x, y)
	return slope
}

// TimeSeriesForecast performs a linear regression forecast for the next n periods.
// Returns nil if the arrays have different lengths or are empty.
// The forecast is based on the linear regression of y against time (0,1,2,...).
func TimeSeriesForecast(y []float64, periods int) []float64 {
	if len(y) == 0 || periods <= 0 {
		return nil
	}

	// Create time series x (0,1,2,...)
	x := make([]float64, len(y))
	for i := range x {
		x[i] = float64(i)
	}

	slope, intercept := LinearRegression(x, y)

	// Generate forecast
	forecast := make([]float64, periods)
	for i := 0; i < periods; i++ {
		forecast[i] = slope*float64(len(y)+i) + intercept
	}

	return forecast
}
// Variance returns the population variance of the input array.
// Returns 0 if the input is empty.
// The variance is the average of the squared differences from the mean.
func Variance(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}

	mean := Mean(arr)
	var sumSqDiff float64
	for _, v := range arr {
		diff := v - mean
		sumSqDiff += diff * diff
	}
	return sumSqDiff / float64(len(arr))
}

// StdDev returns the standard deviation of the input array.
// Returns 0 if the input is empty.
// The standard deviation is the square root of the variance.
func StandardDeviation(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}
	var sum, mean, sqDiff float64
	for _, v := range arr {
		sum += v
	}
	mean = sum / float64(len(arr))
	for _, v := range arr {
		sqDiff += (v - mean) * (v - mean)
	}
	return math.Sqrt(sqDiff / float64(len(arr)))
}

// Correlation computes the Pearson correlation coefficient between two arrays.
// Returns 0 if the arrays have different lengths or are empty.
// The result ranges from -1 to 1, where:
//
//	-1 indicates perfect negative correlation
//	0 indicates no correlation
//	1 indicates perfect positive correlation
func Correlation(x, y []float64) float64 {
	if len(x) != len(y) || len(x) == 0 {
		return 0
	}
	var sumX, sumY, sumXY, sumX2, sumY2 float64
	n := float64(len(x))
	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
		sumY2 += y[i] * y[i]
	}
	numerator := n*sumXY - sumX*sumY
	denominator := math.Sqrt((n*sumX2 - sumX*sumX) * (n*sumY2 - sumY*sumY))
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}
