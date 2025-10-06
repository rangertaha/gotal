package indicators

import (
	"fmt"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/pkg/opt"
	ind "github.com/rangertaha/gotal/internal/plugins/indicators"
	_ "github.com/rangertaha/gotal/internal/plugins/indicators/all"
)

var (
	err error

	// Indicator options
	With           = opt.With
	OnField        = opt.WithField
	OnFields       = opt.WithFields
	WithField      = opt.WithField
	WithFields     = opt.WithFields
	WithOutput     = opt.WithOutput
	WithPeriod     = opt.WithPeriod

	// for MACD, 
	WithSlowPeriod = opt.WithSlowPeriod
	WithFastPeriod = opt.WithFastPeriod
	WithSignalPeriod = opt.WithSignalPeriod
	
	WithMAType     = opt.WithMAType

	// Group of indicators
	TREND  []internal.SeriesFunc
	TRENDs []internal.StreamFunc

	// Moving Average indicators
	OHLCV, SMA, EMA, WMA, DEMA, TEMA, TRIMA, KAMA, MAMA, T3           internal.SeriesFunc
	OHLCVs, SMAs, EMAs, WMAs, DEMAs, TEMAs, TRIMAs, KAMAs, MAMAs, T3s internal.StreamFunc

	// MACD
	MACD  internal.SeriesFunc
	MACDs internal.StreamFunc
)

func init() {
	// OHLCV
	OHLCV, err = ind.Series("ohlcv")
	// OHLCVs, err = ind.Stream("ohlcv")

	// TREND, err = ind.Group(ind.TREND)

	// Simple Moving Average
	SMA, err = ind.Series("sma")
	SMAs, err = ind.Stream("sma")

	// Exponential Moving Average
	EMA, err = ind.Series("ema")
	EMAs, err = ind.Stream("ema")

	// Weighted Moving Average
	WMA, err = ind.Series("wma")
	WMAs, err = ind.Stream("wma")

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
	MACD, err = ind.Series("macd")
	MACDs, err = ind.Stream("macd")

	if err != nil {
		fmt.Println("Error initializing indicators:", err)
		panic(err)
	}
}

