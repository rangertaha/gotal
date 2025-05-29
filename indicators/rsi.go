package indicators

import (
	"errors"

	"github.com/rangertaha/gota/types"
)

// RSI represents the Relative Strength Index indicator
type RSI struct {
	Period int
	Prices []float64
	Result []float64
}

// NewRSI creates a new RSI instance with the specified period
func NewRSI(period int) *RSI {
	return &RSI{
		Period: period,
	}
}

// Calculate computes the RSI for the given price series
func (rsi *RSI) Calculate(prices []float64) (*types.IndicatorResult, error) {
	if len(prices) < rsi.Period+1 {
		return nil, errors.New("insufficient data points for RSI calculation")
	}

	// Calculate price changes
	changes := make([]float64, len(prices)-1)
	for i := 1; i < len(prices); i++ {
		changes[i-1] = prices[i] - prices[i-1]
	}

	// Separate gains and losses
	gains := make([]float64, len(changes))
	losses := make([]float64, len(changes))
	for i, change := range changes {
		if change >= 0 {
			gains[i] = change
			losses[i] = 0
		} else {
			gains[i] = 0
			losses[i] = -change
		}
	}

	// Calculate initial average gain and loss
	avgGain := 0.0
	avgLoss := 0.0
	for i := 0; i < rsi.Period; i++ {
		avgGain += gains[i]
		avgLoss += losses[i]
	}
	avgGain /= float64(rsi.Period)
	avgLoss /= float64(rsi.Period)

	// Calculate RSI values
	result := make([]float64, len(prices)-rsi.Period)
	result[0] = 100 - (100 / (1 + avgGain/avgLoss))

	// Calculate remaining RSI values using Wilder's smoothing method
	for i := 1; i < len(result); i++ {
		avgGain = (avgGain*float64(rsi.Period-1) + gains[i+rsi.Period-1]) / float64(rsi.Period)
		avgLoss = (avgLoss*float64(rsi.Period-1) + losses[i+rsi.Period-1]) / float64(rsi.Period)

		if avgLoss == 0 {
			result[i] = 100
		} else {
			rs := avgGain / avgLoss
			result[i] = 100 - (100 / (1 + rs))
		}
	}

	return &types.IndicatorResult{
		Values: result,
		Error:  nil,
	}, nil
}

// IsValid checks if the RSI calculation is valid
func (rsi *RSI) IsValid() bool {
	return rsi.Period > 0 && len(rsi.Prices) >= rsi.Period+1
}

// GetResult returns the calculated RSI values
func (rsi *RSI) GetResult() []float64 {
	return rsi.Result
}

// IsOverbought checks if the RSI value indicates an overbought condition
func (rsi *RSI) IsOverbought(value float64) bool {
	return value >= 70
}

// IsOversold checks if the RSI value indicates an oversold condition
func (rsi *RSI) IsOversold(value float64) bool {
	return value <= 30
}
