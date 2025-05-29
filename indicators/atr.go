package indicators

import (
	"errors"
	"math"

	"github.com/rangertaha/gota/types"
)

// ATR represents the Average True Range indicator
type ATR struct {
	Period int
	High   []float64
	Low    []float64
	Close  []float64
	Result []float64
}

// NewATR creates a new ATR instance with the specified period
// Standard period is 14
func NewATR(period int) *ATR {
	return &ATR{
		Period: period,
	}
}

// Calculate computes the ATR for the given price series
func (atr *ATR) Calculate(high, low, close []float64) (*types.IndicatorResult, error) {
	if len(high) != len(low) || len(high) != len(close) {
		return nil, errors.New("input arrays must have the same length")
	}
	if len(high) < atr.Period+1 {
		return nil, errors.New("insufficient data points for ATR calculation")
	}

	// Calculate True Range
	trueRange := make([]float64, len(high)-1)
	for i := 1; i < len(high); i++ {
		// True Range is the greatest of:
		// 1. Current High - Current Low
		// 2. |Current High - Previous Close|
		// 3. |Current Low - Previous Close|
		tr1 := high[i] - low[i]
		tr2 := math.Abs(high[i] - close[i-1])
		tr3 := math.Abs(low[i] - close[i-1])
		trueRange[i-1] = math.Max(tr1, math.Max(tr2, tr3))
	}

	// Calculate initial ATR (Simple Moving Average of True Range)
	atr.Result = make([]float64, len(trueRange)-atr.Period+1)
	sum := 0.0
	for i := 0; i < atr.Period; i++ {
		sum += trueRange[i]
	}
	atr.Result[0] = sum / float64(atr.Period)

	// Calculate remaining ATR values using Wilder's smoothing method
	for i := 1; i < len(atr.Result); i++ {
		atr.Result[i] = (atr.Result[i-1]*float64(atr.Period-1) + trueRange[i+atr.Period-1]) / float64(atr.Period)
	}

	return &types.IndicatorResult{
		Values: atr.Result,
		Error:  nil,
	}, nil
}

// IsValid checks if the ATR calculation is valid
func (atr *ATR) IsValid() bool {
	return atr.Period > 0 && len(atr.High) >= atr.Period+1
}

// GetResult returns the calculated ATR values
func (atr *ATR) GetResult() []float64 {
	return atr.Result
}

// IsVolatilityHigh checks if the current volatility is high compared to the average
func (atr *ATR) IsVolatilityHigh(index int, multiplier float64) bool {
	if atr.Result == nil || index >= len(atr.Result) {
		return false
	}

	// Calculate average ATR
	sum := 0.0
	for i := 0; i < len(atr.Result); i++ {
		sum += atr.Result[i]
	}
	avgATR := sum / float64(len(atr.Result))

	return atr.Result[index] > avgATR*multiplier
}

// IsVolatilityLow checks if the current volatility is low compared to the average
func (atr *ATR) IsVolatilityLow(index int, multiplier float64) bool {
	if atr.Result == nil || index >= len(atr.Result) {
		return false
	}

	// Calculate average ATR
	sum := 0.0
	for i := 0; i < len(atr.Result); i++ {
		sum += atr.Result[i]
	}
	avgATR := sum / float64(len(atr.Result))

	return atr.Result[index] < avgATR/multiplier
}

// GetVolatilityRatio returns the ratio of current ATR to average ATR
func (atr *ATR) GetVolatilityRatio(index int) float64 {
	if atr.Result == nil || index >= len(atr.Result) {
		return 0.0
	}

	// Calculate average ATR
	sum := 0.0
	for i := 0; i < len(atr.Result); i++ {
		sum += atr.Result[i]
	}
	avgATR := sum / float64(len(atr.Result))

	if avgATR == 0 {
		return 0.0
	}

	return atr.Result[index] / avgATR
}
