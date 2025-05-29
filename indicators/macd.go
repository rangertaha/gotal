package indicators

import (
	"errors"

	"github.com/rangertaha/gota/types"
)

// MACD represents the Moving Average Convergence Divergence indicator
type MACD struct {
	FastPeriod   int
	SlowPeriod   int
	SignalPeriod int
	Prices       []float64
	Result       *MACDResult
}

// MACDResult contains all the MACD calculation results
type MACDResult struct {
	MACDLine   []float64 // MACD Line (Fast EMA - Slow EMA)
	SignalLine []float64 // Signal Line (EMA of MACD Line)
	Histogram  []float64 // Histogram (MACD Line - Signal Line)
	FastEMA    []float64 // Fast EMA values
	SlowEMA    []float64 // Slow EMA values
}

// NewMACD creates a new MACD instance with the specified periods
// Standard periods are: FastPeriod=12, SlowPeriod=26, SignalPeriod=9
func NewMACD(fastPeriod, slowPeriod, signalPeriod int) *MACD {
	return &MACD{
		FastPeriod:   fastPeriod,
		SlowPeriod:   slowPeriod,
		SignalPeriod: signalPeriod,
	}
}

// Calculate computes the MACD for the given price series
func (macd *MACD) Calculate(prices []float64) (*types.IndicatorResult, error) {
	if len(prices) < macd.SlowPeriod+macd.SignalPeriod {
		return nil, errors.New("insufficient data points for MACD calculation")
	}

	// Calculate Fast EMA
	fastEMA := calculateEMA(prices, macd.FastPeriod)
	if len(fastEMA) == 0 {
		return nil, errors.New("error calculating fast EMA")
	}

	// Calculate Slow EMA
	slowEMA := calculateEMA(prices, macd.SlowPeriod)
	if len(slowEMA) == 0 {
		return nil, errors.New("error calculating slow EMA")
	}

	// Calculate MACD Line (Fast EMA - Slow EMA)
	macdLine := make([]float64, len(slowEMA))
	for i := range slowEMA {
		macdLine[i] = fastEMA[i+macd.SlowPeriod-macd.FastPeriod] - slowEMA[i]
	}

	// Calculate Signal Line (EMA of MACD Line)
	signalLine := calculateEMA(macdLine, macd.SignalPeriod)
	if len(signalLine) == 0 {
		return nil, errors.New("error calculating signal line")
	}

	// Calculate Histogram (MACD Line - Signal Line)
	histogram := make([]float64, len(signalLine))
	for i := range signalLine {
		histogram[i] = macdLine[i+macd.SignalPeriod-1] - signalLine[i]
	}

	// Store results
	macd.Result = &MACDResult{
		MACDLine:   macdLine[macd.SignalPeriod-1:],
		SignalLine: signalLine,
		Histogram:  histogram,
		FastEMA:    fastEMA,
		SlowEMA:    slowEMA,
	}

	return &types.IndicatorResult{
		Values: macdLine[macd.SignalPeriod-1:],
		Error:  nil,
	}, nil
}

// calculateEMA computes the Exponential Moving Average
func calculateEMA(prices []float64, period int) []float64 {
	if len(prices) < period {
		return nil
	}

	multiplier := 2.0 / float64(period+1)
	ema := make([]float64, len(prices))

	// Initialize EMA with SMA
	sum := 0.0
	for i := 0; i < period; i++ {
		sum += prices[i]
	}
	ema[period-1] = sum / float64(period)

	// Calculate EMA
	for i := period; i < len(prices); i++ {
		ema[i] = (prices[i]-ema[i-1])*multiplier + ema[i-1]
	}

	return ema[period-1:]
}

// IsValid checks if the MACD calculation is valid
func (macd *MACD) IsValid() bool {
	return macd.FastPeriod > 0 && macd.SlowPeriod > macd.FastPeriod && macd.SignalPeriod > 0
}

// GetResult returns the calculated MACD values
func (macd *MACD) GetResult() *MACDResult {
	return macd.Result
}

// IsBullish checks if the MACD indicates a bullish signal
func (macd *MACD) IsBullish(index int) bool {
	if macd.Result == nil || index >= len(macd.Result.Histogram) {
		return false
	}
	return macd.Result.Histogram[index] > 0 && macd.Result.Histogram[index] > macd.Result.Histogram[index-1]
}

// IsBearish checks if the MACD indicates a bearish signal
func (macd *MACD) IsBearish(index int) bool {
	if macd.Result == nil || index >= len(macd.Result.Histogram) {
		return false
	}
	return macd.Result.Histogram[index] < 0 && macd.Result.Histogram[index] < macd.Result.Histogram[index-1]
}

// HasBullishDivergence checks for bullish divergence
func (macd *MACD) HasBullishDivergence(prices []float64, index int) bool {
	if macd.Result == nil || index < 2 || index >= len(macd.Result.MACDLine) {
		return false
	}
	return prices[index] < prices[index-1] && macd.Result.MACDLine[index] > macd.Result.MACDLine[index-1]
}

// HasBearishDivergence checks for bearish divergence
func (macd *MACD) HasBearishDivergence(prices []float64, index int) bool {
	if macd.Result == nil || index < 2 || index >= len(macd.Result.MACDLine) {
		return false
	}
	return prices[index] > prices[index-1] && macd.Result.MACDLine[index] < macd.Result.MACDLine[index-1]
}
