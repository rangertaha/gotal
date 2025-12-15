package indicators

import (
	"fmt"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/pkg/opt"
	"github.com/rangertaha/gotal/internal/plugins/indicators"
	_ "github.com/rangertaha/gotal/internal/plugins/indicators/all"
)

var (
	err error

	// Indicator options
	With         = opt.With
	OnField      = opt.WithField
	OnFields     = opt.WithFields
	WithField    = opt.WithField
	WithFields   = opt.WithFields
	WithOutput   = opt.WithOutput
	WithPeriod   = opt.WithPeriod
	WithDuration = opt.WithDuration

	// for MACD,
	WithSlowPeriod   = opt.WithSlowPeriod
	WithFastPeriod   = opt.WithFastPeriod
	WithSignalPeriod = opt.WithSignalPeriod

	WithMAType = opt.WithMAType

	// Group of indicators
	// TREND  []internal.SeriesFunc

	// Mock indicator
	MOCK internal.IndicatorFunc

	// Moving Average indicators
	OHLC, OHLCV, SMA, EMA, WMA, DEMA, TEMA, TRIMA, KAMA, MAMA, T3 internal.IndicatorFunc

	// MACD
	MACD internal.IndicatorFunc
)

func init() {

	// Mock indicator
	MOCK, err = indicators.Get("mock")

	// OHLC, OHLCV
	OHLC, err = indicators.Series("ohlc")
	OHLCV, err = ih.Series("ohlcv")

	// TREND, err = ind.Group(ind.TREND)

	// Simple Moving Average
	SMA, err = i.Series("sma")

	// Exponential Moving Average
	EMA, err = i.Series("ema")

	// Weighted Moving Average
	WMA, err = i.Series("wma")

	// // Double Exponential Moving Average
	// DEMA, err = ind.Series("dema")
	// DEMAs, err = ind.Stream("dema")

	// Triple Exponential Moving Average
	// TEMA, err = ind.Get("tema")
	// TRIMA, err = ind.Get("trima")
	// KAMA, err = ind.Get("kama")
	// MAMA, err = ind.Get("mama")
	// T3, err = ind.Get("t3")

	// Moving Average Convergence Divergence
	MACD, err = i.Series("macd")

	if err != nil {
		fmt.Println("Error initializing indicators:", err)
		panic(err)
	}
}
