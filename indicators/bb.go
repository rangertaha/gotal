package indicators

import (
	"errors"
	"math"

	"github.com/rangertaha/gotal/types"
)

// BollingerBands represents the Bollinger Bands indicator
type BollingerBands struct {
	Period int
	StdDev float64
	Prices []float64
	Result *BollingerBandsResult
}

// BollingerBandsResult contains all the Bollinger Bands calculation results
type BollingerBandsResult struct {
	MiddleBand []float64 // Middle Band (SMA)
	UpperBand  []float64 // Upper Band (Middle + StdDev * Standard Deviation)
	LowerBand  []float64 // Lower Band (Middle - StdDev * Standard Deviation)
	Width      []float64 // Band Width ((Upper - Lower) / Middle)
	PercentB   []float64 // %B ((Price - Lower) / (Upper - Lower))
}

// NewBollingerBands creates a new Bollinger Bands instance
// Standard parameters are: Period=20, StdDev=2.0
func NewBollingerBands(period int, stdDev float64) *BollingerBands {
	return &BollingerBands{
		Period: period,
		StdDev: stdDev,
	}
}

// Calculate computes the Bollinger Bands for the given price series
func (bb *BollingerBands) Calculate(prices []float64) (*types.IndicatorResult, error) {
	if len(prices) < bb.Period {
		return nil, errors.New("insufficient data points for Bollinger Bands calculation")
	}

	// Calculate middle band (SMA)
	middleBand := calculateSMA(prices, bb.Period)
	if len(middleBand) == 0 {
		return nil, errors.New("error calculating middle band")
	}

	// Calculate standard deviation
	stdDev := calculateStandardDeviation(prices, bb.Period)
	if len(stdDev) == 0 {
		return nil, errors.New("error calculating standard deviation")
	}

	// Calculate upper and lower bands
	upperBand := make([]float64, len(middleBand))
	lowerBand := make([]float64, len(middleBand))
	width := make([]float64, len(middleBand))
	percentB := make([]float64, len(middleBand))

	for i := range middleBand {
		upperBand[i] = middleBand[i] + bb.StdDev*stdDev[i]
		lowerBand[i] = middleBand[i] - bb.StdDev*stdDev[i]
		width[i] = (upperBand[i] - lowerBand[i]) / middleBand[i]

		// Calculate %B
		price := prices[i+bb.Period-1]
		bandWidth := upperBand[i] - lowerBand[i]
		if bandWidth != 0 {
			percentB[i] = (price - lowerBand[i]) / bandWidth
		} else {
			percentB[i] = 0.5 // Default to middle when bands are equal
		}
	}

	// Store results
	bb.Result = &BollingerBandsResult{
		MiddleBand: middleBand,
		UpperBand:  upperBand,
		LowerBand:  lowerBand,
		Width:      width,
		PercentB:   percentB,
	}

	return &types.IndicatorResult{
		Values: middleBand,
		Error:  nil,
	}, nil
}

// calculateSMA computes the Simple Moving Average
func calculateSMA(prices []float64, period int) []float64 {
	if len(prices) < period {
		return nil
	}

	result := make([]float64, len(prices)-period+1)
	for i := 0; i <= len(prices)-period; i++ {
		sum := 0.0
		for j := 0; j < period; j++ {
			sum += prices[i+j]
		}
		result[i] = sum / float64(period)
	}
	return result
}

// calculateStandardDeviation computes the standard deviation
func calculateStandardDeviation(prices []float64, period int) []float64 {
	if len(prices) < period {
		return nil
	}

	result := make([]float64, len(prices)-period+1)
	for i := 0; i <= len(prices)-period; i++ {
		// Calculate mean
		sum := 0.0
		for j := 0; j < period; j++ {
			sum += prices[i+j]
		}
		mean := sum / float64(period)

		// Calculate variance
		variance := 0.0
		for j := 0; j < period; j++ {
			diff := prices[i+j] - mean
			variance += diff * diff
		}
		variance /= float64(period)

		// Calculate standard deviation
		result[i] = math.Sqrt(variance)
	}
	return result
}

// IsValid checks if the Bollinger Bands calculation is valid
func (bb *BollingerBands) IsValid() bool {
	return bb.Period > 0 && bb.StdDev > 0
}

// GetResult returns the calculated Bollinger Bands values
func (bb *BollingerBands) GetResult() *BollingerBandsResult {
	return bb.Result
}

// IsOverbought checks if the price is overbought based on %B
func (bb *BollingerBands) IsOverbought(index int) bool {
	if bb.Result == nil || index >= len(bb.Result.PercentB) {
		return false
	}
	return bb.Result.PercentB[index] > 1.0
}

// IsOversold checks if the price is oversold based on %B
func (bb *BollingerBands) IsOversold(index int) bool {
	if bb.Result == nil || index >= len(bb.Result.PercentB) {
		return false
	}
	return bb.Result.PercentB[index] < 0.0
}

// IsSqueeze checks if the bands are in a squeeze (narrowing)
func (bb *BollingerBands) IsSqueeze(index int) bool {
	if bb.Result == nil || index < 1 || index >= len(bb.Result.Width) {
		return false
	}
	return bb.Result.Width[index] < bb.Result.Width[index-1]
}

// IsExpansion checks if the bands are expanding
func (bb *BollingerBands) IsExpansion(index int) bool {
	if bb.Result == nil || index < 1 || index >= len(bb.Result.Width) {
		return false
	}
	return bb.Result.Width[index] > bb.Result.Width[index-1]
}
