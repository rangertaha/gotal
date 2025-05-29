package indicators

import (
	"errors"

	"github.com/rangertaha/gota/types"
)

// ROC represents the Rate of Change indicator
type ROC struct {
	Period int // ROC period
	Prices []float64
	Result *ROCResult
}

// ROCResult contains all the ROC calculation results
type ROCResult struct {
	ROC      []float64 // Rate of Change
	Momentum []float64 // Momentum (price difference)
}

// NewROC creates a new ROC instance with the specified period
// Standard period is 14
func NewROC(period int) *ROC {
	return &ROC{
		Period: period,
	}
}

// Calculate computes the ROC for the given price series
func (roc *ROC) Calculate(prices []float64) (*types.IndicatorResult, error) {
	if len(prices) < roc.Period+1 {
		return nil, errors.New("insufficient data points for ROC calculation")
	}

	// Calculate ROC and Momentum
	rocValues := make([]float64, len(prices)-roc.Period)
	momentum := make([]float64, len(prices)-roc.Period)

	for i := 0; i < len(rocValues); i++ {
		// Calculate momentum (price difference)
		momentum[i] = prices[i+roc.Period] - prices[i]

		// Calculate ROC
		if prices[i] == 0 {
			rocValues[i] = 0 // Avoid division by zero
		} else {
			rocValues[i] = (momentum[i] / prices[i]) * 100
		}
	}

	// Store results
	roc.Result = &ROCResult{
		ROC:      rocValues,
		Momentum: momentum,
	}

	return &types.IndicatorResult{
		Values: rocValues,
		Error:  nil,
	}, nil
}

// IsValid checks if the ROC calculation is valid
func (roc *ROC) IsValid() bool {
	return roc.Period > 0
}

// GetResult returns the calculated ROC values
func (roc *ROC) GetResult() *ROCResult {
	return roc.Result
}

// IsOverbought checks if the ROC indicates an overbought condition
func (roc *ROC) IsOverbought(index int) bool {
	if roc.Result == nil || index >= len(roc.Result.ROC) {
		return false
	}
	return roc.Result.ROC[index] > 10.0
}

// IsOversold checks if the ROC indicates an oversold condition
func (roc *ROC) IsOversold(index int) bool {
	if roc.Result == nil || index >= len(roc.Result.ROC) {
		return false
	}
	return roc.Result.ROC[index] < -10.0
}

// HasBullishDivergence checks for bullish divergence
func (roc *ROC) HasBullishDivergence(prices []float64, index int) bool {
	if roc.Result == nil || index < 2 || index >= len(roc.Result.ROC) {
		return false
	}
	return prices[index] < prices[index-1] && roc.Result.ROC[index] > roc.Result.ROC[index-1]
}

// HasBearishDivergence checks for bearish divergence
func (roc *ROC) HasBearishDivergence(prices []float64, index int) bool {
	if roc.Result == nil || index < 2 || index >= len(roc.Result.ROC) {
		return false
	}
	return prices[index] > prices[index-1] && roc.Result.ROC[index] < roc.Result.ROC[index-1]
}

// HasZeroLineCross checks for zero line crossovers
func (roc *ROC) HasZeroLineCross(index int) bool {
	if roc.Result == nil || index < 1 || index >= len(roc.Result.ROC) {
		return false
	}
	return (roc.Result.ROC[index] > 0 && roc.Result.ROC[index-1] <= 0) ||
		(roc.Result.ROC[index] < 0 && roc.Result.ROC[index-1] >= 0)
}
