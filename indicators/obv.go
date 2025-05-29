package indicators

import (
	"errors"

	"github.com/rangertaha/gotal/types"
)

// OBV represents the On-Balance Volume indicator
type OBV struct {
	Close  []float64
	Volume []float64
	Result *OBVResult
}

// OBVResult contains all the OBV calculation results
type OBVResult struct {
	OBV       []float64 // On-Balance Volume
	OBVChange []float64 // OBV Change
	OBVSlope  []float64 // OBV Slope
}

// NewOBV creates a new OBV instance
func NewOBV() *OBV {
	return &OBV{}
}

// Calculate computes the OBV for the given price and volume series
func (obv *OBV) Calculate(close, volume []float64) (*types.IndicatorResult, error) {
	if len(close) != len(volume) {
		return nil, errors.New("input arrays must have the same length")
	}
	if len(close) < 2 {
		return nil, errors.New("insufficient data points for OBV calculation")
	}

	// Calculate OBV
	obvValues := make([]float64, len(close))
	obvValues[0] = volume[0]

	for i := 1; i < len(close); i++ {
		if close[i] > close[i-1] {
			obvValues[i] = obvValues[i-1] + volume[i]
		} else if close[i] < close[i-1] {
			obvValues[i] = obvValues[i-1] - volume[i]
		} else {
			obvValues[i] = obvValues[i-1]
		}
	}

	// Calculate OBV Change
	obvChange := make([]float64, len(close))
	obvChange[0] = 0
	for i := 1; i < len(close); i++ {
		obvChange[i] = obvValues[i] - obvValues[i-1]
	}

	// Calculate OBV Slope (5-period)
	obvSlope := make([]float64, len(close))
	period := 5
	for i := period; i < len(close); i++ {
		obvSlope[i] = (obvValues[i] - obvValues[i-period]) / float64(period)
	}

	// Store results
	obv.Result = &OBVResult{
		OBV:       obvValues,
		OBVChange: obvChange,
		OBVSlope:  obvSlope,
	}

	return &types.IndicatorResult{
		Values: obvValues,
		Error:  nil,
	}, nil
}

// IsValid checks if the OBV calculation is valid
func (obv *OBV) IsValid() bool {
	return len(obv.Close) > 0 && len(obv.Volume) > 0
}

// GetResult returns the calculated OBV values
func (obv *OBV) GetResult() *OBVResult {
	return obv.Result
}

// IsBullish checks if the OBV indicates a bullish trend
func (obv *OBV) IsBullish(index int) bool {
	if obv.Result == nil || index < 1 || index >= len(obv.Result.OBV) {
		return false
	}
	return obv.Result.OBV[index] > obv.Result.OBV[index-1]
}

// IsBearish checks if the OBV indicates a bearish trend
func (obv *OBV) IsBearish(index int) bool {
	if obv.Result == nil || index < 1 || index >= len(obv.Result.OBV) {
		return false
	}
	return obv.Result.OBV[index] < obv.Result.OBV[index-1]
}

// HasBullishDivergence checks for bullish divergence
func (obv *OBV) HasBullishDivergence(prices []float64, index int) bool {
	if obv.Result == nil || index < 2 || index >= len(obv.Result.OBV) {
		return false
	}
	return prices[index] < prices[index-1] && obv.Result.OBV[index] > obv.Result.OBV[index-1]
}

// HasBearishDivergence checks for bearish divergence
func (obv *OBV) HasBearishDivergence(prices []float64, index int) bool {
	if obv.Result == nil || index < 2 || index >= len(obv.Result.OBV) {
		return false
	}
	return prices[index] > prices[index-1] && obv.Result.OBV[index] < obv.Result.OBV[index-1]
}

// IsAccumulation checks if the OBV indicates accumulation
func (obv *OBV) IsAccumulation(index int) bool {
	if obv.Result == nil || index < 5 || index >= len(obv.Result.OBVSlope) {
		return false
	}
	return obv.Result.OBVSlope[index] > 0
}

// IsDistribution checks if the OBV indicates distribution
func (obv *OBV) IsDistribution(index int) bool {
	if obv.Result == nil || index < 5 || index >= len(obv.Result.OBVSlope) {
		return false
	}
	return obv.Result.OBVSlope[index] < 0
}
