package indicators

import (
	"errors"

	"github.com/rangertaha/gota/types"
)

// MovingAverage calculates the moving average of a price series
type MovingAverage struct {
	Type   types.MovingAverageType
	Period int
	Prices []float64
	Result []float64
}

// NewMovingAverage creates a new MovingAverage instance
func NewMovingAverage(maType types.MovingAverageType, period int) *MovingAverage {
	return &MovingAverage{
		Type:   maType,
		Period: period,
	}
}

// Calculate computes the moving average based on the specified type
func (ma *MovingAverage) Calculate(prices []float64) (*types.IndicatorResult, error) {
	if len(prices) < ma.Period {
		return nil, errors.New("insufficient data points for calculation")
	}

	switch ma.Type {
	case types.SMA:
		return ma.calculateSMA(prices)
	case types.EMA:
		return ma.calculateEMA(prices)
	case types.WMA:
		return ma.calculateWMA(prices)
	default:
		return nil, errors.New("unsupported moving average type")
	}
}

// calculateSMA computes the Simple Moving Average
func (ma *MovingAverage) calculateSMA(prices []float64) (*types.IndicatorResult, error) {
	result := make([]float64, len(prices)-ma.Period+1)

	for i := 0; i <= len(prices)-ma.Period; i++ {
		sum := 0.0
		for j := 0; j < ma.Period; j++ {
			sum += prices[i+j]
		}
		result[i] = sum / float64(ma.Period)
	}

	return &types.IndicatorResult{
		Values: result,
		Error:  nil,
	}, nil
}

// calculateEMA computes the Exponential Moving Average
func (ma *MovingAverage) calculateEMA(prices []float64) (*types.IndicatorResult, error) {
	result := make([]float64, len(prices))
	multiplier := 2.0 / float64(ma.Period+1)

	// Initialize EMA with SMA
	sum := 0.0
	for i := 0; i < ma.Period; i++ {
		sum += prices[i]
	}
	result[ma.Period-1] = sum / float64(ma.Period)

	// Calculate EMA
	for i := ma.Period; i < len(prices); i++ {
		result[i] = (prices[i]-result[i-1])*multiplier + result[i-1]
	}

	return &types.IndicatorResult{
		Values: result[ma.Period-1:],
		Error:  nil,
	}, nil
}

// calculateWMA computes the Weighted Moving Average
func (ma *MovingAverage) calculateWMA(prices []float64) (*types.IndicatorResult, error) {
	result := make([]float64, len(prices)-ma.Period+1)

	for i := 0; i <= len(prices)-ma.Period; i++ {
		sum := 0.0
		weightSum := 0.0
		for j := 0; j < ma.Period; j++ {
			weight := float64(j + 1)
			sum += prices[i+j] * weight
			weightSum += weight
		}
		result[i] = sum / weightSum
	}

	return &types.IndicatorResult{
		Values: result,
		Error:  nil,
	}, nil
}

// IsValid checks if the moving average calculation is valid
func (ma *MovingAverage) IsValid() bool {
	return ma.Period > 0 && len(ma.Prices) >= ma.Period
}

// GetResult returns the calculated moving average values
func (ma *MovingAverage) GetResult() []float64 {
	return ma.Result
}
