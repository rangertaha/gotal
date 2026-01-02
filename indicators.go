package gotal

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/config"
	"github.com/rangertaha/gotal/internal/indicators"
)

var (
	// Batch indicators functions
	EMA, SMA, RSI, MACD, BB internal.IndicatorFunc

	// Stream indicators functions
	// TODO: Implement stream indicators functions
)

func init() {
	// EMA is the Exponential Moving Average indicator function.
	// It is a moving average indicator that places a greater weight and significance on the most recent data points.
	//
	// Usage example:
	//   ema, err := gotal.EMA(inputSeries, config.With("period", 10))
	//   if err != nil { /* handle error */ }
	//
	// The parameters typically include:
	//   - "period": the length of the EMA window (int)
	//   - "alpha": optional smoothing factor (float64)
	//
	// See also: Ema (convenience wrapper).
	EMA = indicators.Func("ema")

	// ... and other indicators ...
}

// Ema computes the Exponential Moving Average (EMA) of the given input series.
//
// Arguments:
//   - input: the input time series (Series)
//   - period: the window length of the EMA (int)
//   - alpha: the smoothing factor (float64)
//
// Returns:
//   - Series: the resulting EMA time series
//   - error: non-nil if computation fails
//
// Example usage:
//   ema, err := gotal.Ema(prices, 14, 0.2)
//   if err != nil { /* handle error */ }
//   ema.Print()
//
// This is a convenience wrapper around the generic EMA indicator, accepting period and alpha directly.
func Ema(input internal.Series, period int, alpha float64) (internal.Series, error) {
	return EMA(input, config.With("period", period), config.With("alpha", alpha))
}
