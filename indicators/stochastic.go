package indicators

import (
	"errors"
	"math"

	"github.com/rangertaha/gotal/types"
)

// Stochastic represents the Stochastic Oscillator indicator
type Stochastic struct {
	KPeriod int // %K period
	DPeriod int // %D period (SMA of %K)
	Slowing int // Slowing period
	High    []float64
	Low     []float64
	Close   []float64
	Result  *StochasticResult
}

// StochasticResult contains all the Stochastic calculation results
type StochasticResult struct {
	K []float64 // %K line (Fast Stochastic)
	D []float64 // %D line (Signal line)
}

// NewStochastic creates a new Stochastic instance with the specified periods
// Standard parameters are: KPeriod=14, DPeriod=3, Slowing=3
func NewStochastic(kPeriod, dPeriod, slowing int) *Stochastic {
	return &Stochastic{
		KPeriod: kPeriod,
		DPeriod: dPeriod,
		Slowing: slowing,
	}
}

// Calculate computes the Stochastic Oscillator for the given price series
func (stoch *Stochastic) Calculate(high, low, close []float64) (*types.IndicatorResult, error) {
	if len(high) != len(low) || len(high) != len(close) {
		return nil, errors.New("input arrays must have the same length")
	}
	if len(high) < stoch.KPeriod+stoch.Slowing+stoch.DPeriod {
		return nil, errors.New("insufficient data points for Stochastic calculation")
	}

	// Calculate %K
	k := make([]float64, len(close)-stoch.KPeriod-stoch.Slowing+1)
	for i := 0; i < len(k); i++ {
		// Find highest high and lowest low in the period
		highestHigh := high[i]
		lowestLow := low[i]
		for j := 1; j < stoch.KPeriod; j++ {
			highestHigh = math.Max(highestHigh, high[i+j])
			lowestLow = math.Min(lowestLow, low[i+j])
		}

		// Calculate %K
		if highestHigh == lowestLow {
			k[i] = 50.0 // Default to middle when range is zero
		} else {
			k[i] = 100.0 * (close[i+stoch.KPeriod-1] - lowestLow) / (highestHigh - lowestLow)
		}
	}

	// Apply slowing period to %K
	if stoch.Slowing > 1 {
		smoothedK := make([]float64, len(k)-stoch.Slowing+1)
		for i := 0; i < len(smoothedK); i++ {
			sum := 0.0
			for j := 0; j < stoch.Slowing; j++ {
				sum += k[i+j]
			}
			smoothedK[i] = sum / float64(stoch.Slowing)
		}
		k = smoothedK
	}

	// Calculate %D (SMA of %K)
	d := make([]float64, len(k)-stoch.DPeriod+1)
	for i := 0; i < len(d); i++ {
		sum := 0.0
		for j := 0; j < stoch.DPeriod; j++ {
			sum += k[i+j]
		}
		d[i] = sum / float64(stoch.DPeriod)
	}

	// Store results
	stoch.Result = &StochasticResult{
		K: k[stoch.DPeriod-1:],
		D: d,
	}

	return &types.IndicatorResult{
		Values: k[stoch.DPeriod-1:],
		Error:  nil,
	}, nil
}

// IsValid checks if the Stochastic calculation is valid
func (stoch *Stochastic) IsValid() bool {
	return stoch.KPeriod > 0 && stoch.DPeriod > 0 && stoch.Slowing > 0
}

// GetResult returns the calculated Stochastic values
func (stoch *Stochastic) GetResult() *StochasticResult {
	return stoch.Result
}

// IsOverbought checks if the Stochastic indicates an overbought condition
func (stoch *Stochastic) IsOverbought(index int) bool {
	if stoch.Result == nil || index >= len(stoch.Result.K) {
		return false
	}
	return stoch.Result.K[index] > 80.0
}

// IsOversold checks if the Stochastic indicates an oversold condition
func (stoch *Stochastic) IsOversold(index int) bool {
	if stoch.Result == nil || index >= len(stoch.Result.K) {
		return false
	}
	return stoch.Result.K[index] < 20.0
}

// HasBullishDivergence checks for bullish divergence
func (stoch *Stochastic) HasBullishDivergence(prices []float64, index int) bool {
	if stoch.Result == nil || index < 2 || index >= len(stoch.Result.K) {
		return false
	}
	return prices[index] < prices[index-1] && stoch.Result.K[index] > stoch.Result.K[index-1]
}

// HasBearishDivergence checks for bearish divergence
func (stoch *Stochastic) HasBearishDivergence(prices []float64, index int) bool {
	if stoch.Result == nil || index < 2 || index >= len(stoch.Result.K) {
		return false
	}
	return prices[index] > prices[index-1] && stoch.Result.K[index] < stoch.Result.K[index-1]
}

// HasBullishCross checks for bullish crossover (%K crosses above %D)
func (stoch *Stochastic) HasBullishCross(index int) bool {
	if stoch.Result == nil || index < 1 || index >= len(stoch.Result.K) {
		return false
	}
	return stoch.Result.K[index] > stoch.Result.D[index] && stoch.Result.K[index-1] <= stoch.Result.D[index-1]
}

// HasBearishCross checks for bearish crossover (%K crosses below %D)
func (stoch *Stochastic) HasBearishCross(index int) bool {
	if stoch.Result == nil || index < 1 || index >= len(stoch.Result.K) {
		return false
	}
	return stoch.Result.K[index] < stoch.Result.D[index] && stoch.Result.K[index-1] >= stoch.Result.D[index-1]
}
