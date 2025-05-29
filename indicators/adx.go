package indicators

import (
	"errors"
	"math"

	"github.com/rangertaha/gota/types"
)

// ADX represents the Average Directional Index indicator
type ADX struct {
	Period int // ADX period
	High   []float64
	Low    []float64
	Close  []float64
	Result *ADXResult
}

// ADXResult contains all the ADX calculation results
type ADXResult struct {
	ADX     []float64 // Average Directional Index
	PlusDI  []float64 // Plus Directional Indicator
	MinusDI []float64 // Minus Directional Indicator
}

// NewADX creates a new ADX instance with the specified period
// Standard period is 14
func NewADX(period int) *ADX {
	return &ADX{
		Period: period,
	}
}

// Calculate computes the ADX for the given price series
func (adx *ADX) Calculate(high, low, close []float64) (*types.IndicatorResult, error) {
	if len(high) != len(low) || len(high) != len(close) {
		return nil, errors.New("input arrays must have the same length")
	}
	if len(high) < adx.Period+1 {
		return nil, errors.New("insufficient data points for ADX calculation")
	}

	// Calculate True Range (TR)
	tr := make([]float64, len(close))
	tr[0] = high[0] - low[0]
	for i := 1; i < len(close); i++ {
		tr[i] = math.Max(high[i]-low[i], math.Max(math.Abs(high[i]-close[i-1]), math.Abs(low[i]-close[i-1])))
	}

	// Calculate Directional Movement
	plusDM := make([]float64, len(close))
	minusDM := make([]float64, len(close))
	for i := 1; i < len(close); i++ {
		upMove := high[i] - high[i-1]
		downMove := low[i-1] - low[i]

		if upMove > downMove && upMove > 0 {
			plusDM[i] = upMove
		} else {
			plusDM[i] = 0
		}

		if downMove > upMove && downMove > 0 {
			minusDM[i] = downMove
		} else {
			minusDM[i] = 0
		}
	}

	// Calculate Smoothed TR and DM
	smoothedTR := make([]float64, len(close))
	smoothedPlusDM := make([]float64, len(close))
	smoothedMinusDM := make([]float64, len(close))

	// First value is simple sum
	smoothedTR[adx.Period] = 0
	smoothedPlusDM[adx.Period] = 0
	smoothedMinusDM[adx.Period] = 0
	for i := 1; i <= adx.Period; i++ {
		smoothedTR[adx.Period] += tr[i]
		smoothedPlusDM[adx.Period] += plusDM[i]
		smoothedMinusDM[adx.Period] += minusDM[i]
	}

	// Subsequent values use Wilder's smoothing
	for i := adx.Period + 1; i < len(close); i++ {
		smoothedTR[i] = smoothedTR[i-1] - (smoothedTR[i-1] / float64(adx.Period)) + tr[i]
		smoothedPlusDM[i] = smoothedPlusDM[i-1] - (smoothedPlusDM[i-1] / float64(adx.Period)) + plusDM[i]
		smoothedMinusDM[i] = smoothedMinusDM[i-1] - (smoothedMinusDM[i-1] / float64(adx.Period)) + minusDM[i]
	}

	// Calculate Plus and Minus Directional Indicators
	plusDI := make([]float64, len(close))
	minusDI := make([]float64, len(close))
	for i := adx.Period; i < len(close); i++ {
		plusDI[i] = 100 * (smoothedPlusDM[i] / smoothedTR[i])
		minusDI[i] = 100 * (smoothedMinusDM[i] / smoothedTR[i])
	}

	// Calculate Directional Index (DX)
	dx := make([]float64, len(close))
	for i := adx.Period; i < len(close); i++ {
		dx[i] = 100 * math.Abs(plusDI[i]-minusDI[i]) / (plusDI[i] + minusDI[i])
	}

	// Calculate ADX
	adxValues := make([]float64, len(close))
	adxValues[adx.Period*2-1] = 0
	for i := adx.Period; i < adx.Period*2; i++ {
		adxValues[adx.Period*2-1] += dx[i]
	}
	adxValues[adx.Period*2-1] /= float64(adx.Period)

	for i := adx.Period * 2; i < len(close); i++ {
		adxValues[i] = ((adxValues[i-1] * float64(adx.Period-1)) + dx[i]) / float64(adx.Period)
	}

	// Store results
	adx.Result = &ADXResult{
		ADX:     adxValues[adx.Period*2-1:],
		PlusDI:  plusDI[adx.Period*2-1:],
		MinusDI: minusDI[adx.Period*2-1:],
	}

	return &types.IndicatorResult{
		Values: adxValues[adx.Period*2-1:],
		Error:  nil,
	}, nil
}

// IsValid checks if the ADX calculation is valid
func (adx *ADX) IsValid() bool {
	return adx.Period > 0
}

// GetResult returns the calculated ADX values
func (adx *ADX) GetResult() *ADXResult {
	return adx.Result
}

// IsStrongTrend checks if the ADX indicates a strong trend
func (adx *ADX) IsStrongTrend(index int) bool {
	if adx.Result == nil || index >= len(adx.Result.ADX) {
		return false
	}
	return adx.Result.ADX[index] > 25.0
}

// IsVeryStrongTrend checks if the ADX indicates a very strong trend
func (adx *ADX) IsVeryStrongTrend(index int) bool {
	if adx.Result == nil || index >= len(adx.Result.ADX) {
		return false
	}
	return adx.Result.ADX[index] > 50.0
}

// IsBullishTrend checks if the trend is bullish (PlusDI > MinusDI)
func (adx *ADX) IsBullishTrend(index int) bool {
	if adx.Result == nil || index >= len(adx.Result.PlusDI) {
		return false
	}
	return adx.Result.PlusDI[index] > adx.Result.MinusDI[index]
}

// IsBearishTrend checks if the trend is bearish (MinusDI > PlusDI)
func (adx *ADX) IsBearishTrend(index int) bool {
	if adx.Result == nil || index >= len(adx.Result.PlusDI) {
		return false
	}
	return adx.Result.MinusDI[index] > adx.Result.PlusDI[index]
}

// HasTrendReversal checks for trend reversals
func (adx *ADX) HasTrendReversal(index int) bool {
	if adx.Result == nil || index < 1 || index >= len(adx.Result.PlusDI) {
		return false
	}
	return (adx.Result.PlusDI[index] > adx.Result.MinusDI[index] && adx.Result.PlusDI[index-1] <= adx.Result.MinusDI[index-1]) ||
		(adx.Result.MinusDI[index] > adx.Result.PlusDI[index] && adx.Result.MinusDI[index-1] <= adx.Result.PlusDI[index-1])
}
