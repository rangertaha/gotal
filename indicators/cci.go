package indicators

import (
	"errors"
	"math"

	"github.com/rangertaha/gota/types"
)

// CCI represents the Commodity Channel Index indicator
type CCI struct {
	Period int // CCI period
	High   []float64
	Low    []float64
	Close  []float64
	Result *CCIResult
}

// CCIResult contains all the CCI calculation results
type CCIResult struct {
	CCI     []float64 // Commodity Channel Index
	TP      []float64 // Typical Price
	MeanTP  []float64 // Mean of Typical Price
	MeanDev []float64 // Mean Deviation
}

// NewCCI creates a new CCI instance with the specified period
// Standard period is 20
func NewCCI(period int) *CCI {
	return &CCI{
		Period: period,
	}
}

// Calculate computes the CCI for the given price series
func (cci *CCI) Calculate(high, low, close []float64) (*types.IndicatorResult, error) {
	if len(high) != len(low) || len(high) != len(close) {
		return nil, errors.New("input arrays must have the same length")
	}
	if len(high) < cci.Period {
		return nil, errors.New("insufficient data points for CCI calculation")
	}

	// Calculate Typical Price (TP)
	tp := make([]float64, len(close))
	for i := 0; i < len(close); i++ {
		tp[i] = (high[i] + low[i] + close[i]) / 3.0
	}

	// Calculate Mean of Typical Price
	meanTP := make([]float64, len(close)-cci.Period+1)
	for i := 0; i < len(meanTP); i++ {
		sum := 0.0
		for j := 0; j < cci.Period; j++ {
			sum += tp[i+j]
		}
		meanTP[i] = sum / float64(cci.Period)
	}

	// Calculate Mean Deviation
	meanDev := make([]float64, len(close)-cci.Period+1)
	for i := 0; i < len(meanDev); i++ {
		sum := 0.0
		for j := 0; j < cci.Period; j++ {
			sum += math.Abs(tp[i+j] - meanTP[i])
		}
		meanDev[i] = sum / float64(cci.Period)
	}

	// Calculate CCI
	cciValues := make([]float64, len(close)-cci.Period+1)
	for i := 0; i < len(cciValues); i++ {
		if meanDev[i] == 0 {
			cciValues[i] = 0 // Avoid division by zero
		} else {
			cciValues[i] = (tp[i+cci.Period-1] - meanTP[i]) / (0.015 * meanDev[i])
		}
	}

	// Store results
	cci.Result = &CCIResult{
		CCI:     cciValues,
		TP:      tp[cci.Period-1:],
		MeanTP:  meanTP,
		MeanDev: meanDev,
	}

	return &types.IndicatorResult{
		Values: cciValues,
		Error:  nil,
	}, nil
}

// IsValid checks if the CCI calculation is valid
func (cci *CCI) IsValid() bool {
	return cci.Period > 0
}

// GetResult returns the calculated CCI values
func (cci *CCI) GetResult() *CCIResult {
	return cci.Result
}

// IsOverbought checks if the CCI indicates an overbought condition
func (cci *CCI) IsOverbought(index int) bool {
	if cci.Result == nil || index >= len(cci.Result.CCI) {
		return false
	}
	return cci.Result.CCI[index] > 100.0
}

// IsOversold checks if the CCI indicates an oversold condition
func (cci *CCI) IsOversold(index int) bool {
	if cci.Result == nil || index >= len(cci.Result.CCI) {
		return false
	}
	return cci.Result.CCI[index] < -100.0
}

// IsExtremelyOverbought checks if the CCI indicates an extremely overbought condition
func (cci *CCI) IsExtremelyOverbought(index int) bool {
	if cci.Result == nil || index >= len(cci.Result.CCI) {
		return false
	}
	return cci.Result.CCI[index] > 200.0
}

// IsExtremelyOversold checks if the CCI indicates an extremely oversold condition
func (cci *CCI) IsExtremelyOversold(index int) bool {
	if cci.Result == nil || index >= len(cci.Result.CCI) {
		return false
	}
	return cci.Result.CCI[index] < -200.0
}

// HasBullishDivergence checks for bullish divergence
func (cci *CCI) HasBullishDivergence(prices []float64, index int) bool {
	if cci.Result == nil || index < 2 || index >= len(cci.Result.CCI) {
		return false
	}
	return prices[index] < prices[index-1] && cci.Result.CCI[index] > cci.Result.CCI[index-1]
}

// HasBearishDivergence checks for bearish divergence
func (cci *CCI) HasBearishDivergence(prices []float64, index int) bool {
	if cci.Result == nil || index < 2 || index >= len(cci.Result.CCI) {
		return false
	}
	return prices[index] > prices[index-1] && cci.Result.CCI[index] < cci.Result.CCI[index-1]
}

// HasZeroLineCross checks for zero line crossovers
func (cci *CCI) HasZeroLineCross(index int) bool {
	if cci.Result == nil || index < 1 || index >= len(cci.Result.CCI) {
		return false
	}
	return (cci.Result.CCI[index] > 0 && cci.Result.CCI[index-1] <= 0) ||
		(cci.Result.CCI[index] < 0 && cci.Result.CCI[index-1] >= 0)
}
