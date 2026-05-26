// Copyright (C) 2026 Rangertaha <rangertaha@gmail.com>
// SPDX-License-Identifier: GPL-3.0-or-later
package gotal

import (
	"fmt"

	"github.com/rangertaha/gotal/internal/config"
	"github.com/rangertaha/gotal/internal/indicators"
	_ "github.com/rangertaha/gotal/internal/indicators/all"
)

// Get returns a batch IndicatorFunc for the given registered id. If the id is
// unknown, it returns a function that errors at call time rather than panicking.
func Get(id string) IndicatorFunc {
	if indicators.Has(id) {
		return indicators.Func(id)
	}
	return func(ts TimeSeries, opts ...ConfigOption) (TimeSeries, error) {
		return nil, fmt.Errorf("gotal: indicator %q is not registered", id)
	}
}

// Has reports whether an indicator with the given id is registered.
func Has(id string) bool { return indicators.Has(id) }

// List returns every registered indicator id.
func List() []string { return indicators.All() }

// With creates a config option that sets a key/value pair on indicator config.
func With(key string, value any) ConfigOption {
	return config.With(key, value)
}

// Overlap Studies
var (
	SMA      = Get("SMA")
	EMA      = Get("EMA")
	DEMA     = Get("DEMA")
	TEMA     = Get("TEMA")
	T3       = Get("T3")
	KAMA     = Get("KAMA")
	MAMA     = Get("MAMA")
	TRIMA    = Get("TRIMA")
	WMA      = Get("WMA")
	MA       = Get("MA")
	MAVP     = Get("MAVP")
	MIDPOINT = Get("MIDPOINT")
	MIDPRICE = Get("MIDPRICE")
	BBANDS   = Get("BBANDS")
	SAR      = Get("SAR")
	SAREXT   = Get("SAREXT")
)

// Momentum Indicators
var (
	MACD     = Get("MACD")
	MACDEXT  = Get("MACDEXT")
	MACDFIX  = Get("MACDFIX")
	RSI      = Get("RSI")
	APO      = Get("APO")
	PPO      = Get("PPO")
	MOM      = Get("MOM")
	CMO      = Get("CMO")
	TRIX     = Get("TRIX")
	ROC      = Get("ROC")
	ROCP     = Get("ROCP")
	ROCR     = Get("ROCR")
	ROCR100  = Get("ROCR100")
	STOCH    = Get("STOCH")
	STOCHF   = Get("STOCHF")
	STOCHRSI = Get("STOCHRSI")
	ULTOSC   = Get("ULTOSC")
	WILLR    = Get("WILLR")
	CCI      = Get("CCI")
	BOP      = Get("BOP")
	ADX      = Get("ADX")
	ADXR     = Get("ADXR")
	DX       = Get("DX")
	PLUS_DI  = Get("PLUS_DI")
	MINUS_DI = Get("MINUS_DI")
	PLUS_DM  = Get("PLUS_DM")
	MINUS_DM = Get("MINUS_DM")
	AROON    = Get("AROON")
	AROONOSC = Get("AROONOSC")
	MFI      = Get("MFI")
)

// Volatility
var (
	ATR    = Get("ATR")
	NATR   = Get("NATR")
	TRANGE = Get("TRANGE")
)

// Volume
var (
	OBV   = Get("OBV")
	AD    = Get("AD")
	ADOSC = Get("ADOSC")
	VWMA  = Get("VWMA")
)

// Cycle (Hilbert Transform)
var (
	HT_DCPERIOD  = Get("HT_DCPERIOD")
	HT_DCPHASE   = Get("HT_DCPHASE")
	HT_PHASOR    = Get("HT_PHASOR")
	HT_SINE      = Get("HT_SINE")
	HT_TRENDLINE = Get("HT_TRENDLINE")
	HT_TRENDMODE = Get("HT_TRENDMODE")
)

// Price Transform
var (
	AVGPRICE = Get("AVGPRICE")
	MEDPRICE = Get("MEDPRICE")
	TYPPRICE = Get("TYPPRICE")
	WCLPRICE = Get("WCLPRICE")
)

// Statistic Functions
var (
	BETA                = Get("BETA")
	CORREL              = Get("CORREL")
	STDDEV              = Get("STDDEV")
	VARIANCE            = Get("VARIANCE")
	LINEARREG           = Get("LINEARREG")
	LINEARREG_SLOPE     = Get("LINEARREG_SLOPE")
	LINEARREG_ANGLE     = Get("LINEARREG_ANGLE")
	LINEARREG_INTERCEPT = Get("LINEARREG_INTERCEPT")
	TSF                 = Get("TSF")
)

// Math Operators
var (
	ADD         = Get("ADD")
	SUB         = Get("SUB")
	MULT        = Get("MULT")
	DIV         = Get("DIV")
	MAX         = Get("MAX")
	MIN         = Get("MIN")
	MAXINDEX    = Get("MAXINDEX")
	MININDEX    = Get("MININDEX")
	MINMAX      = Get("MINMAX")
	MINMAXINDEX = Get("MINMAXINDEX")
	SUMWINDOW   = Get("SUMWINDOW")
)

// Community extensions
var (
	HMA        = Get("HMA")
	VWAP       = Get("VWAP")
	HEIKINASHI = Get("HEIKINASHI")
	HLC3       = Get("HLC3")
	OHLC4      = Get("OHLC4")
)

// Math Transforms
var (
	ACOS  = Get("ACOS")
	ASIN  = Get("ASIN")
	ATAN  = Get("ATAN")
	COS   = Get("COS")
	SIN   = Get("SIN")
	TAN   = Get("TAN")
	COSH  = Get("COSH")
	SINH  = Get("SINH")
	TANH  = Get("TANH")
	CEIL  = Get("CEIL")
	FLOOR = Get("FLOOR")
	EXP   = Get("EXP")
	LN    = Get("LN")
	LOG10 = Get("LOG10")
	SQRT  = Get("SQRT")
)

// Typed convenience wrappers for the most common indicators.

func Ema(input TimeSeries, period int, alpha float64) (TimeSeries, error) {
	return EMA(input, With("period", period), With("alpha", alpha))
}
func Sma(input TimeSeries, period int) (TimeSeries, error) {
	return SMA(input, With("period", period))
}
func Wma(input TimeSeries, period int) (TimeSeries, error) {
	return WMA(input, With("period", period))
}
func Dema(input TimeSeries, period int) (TimeSeries, error) {
	return DEMA(input, With("period", period))
}
func Tema(input TimeSeries, period int) (TimeSeries, error) {
	return TEMA(input, With("period", period))
}
func Trima(input TimeSeries, period int) (TimeSeries, error) {
	return TRIMA(input, With("period", period))
}
func Midpoint(input TimeSeries, period int) (TimeSeries, error) {
	return MIDPOINT(input, With("period", period))
}
func Midprice(input TimeSeries, period int) (TimeSeries, error) {
	return MIDPRICE(input, With("period", period))
}
func Rsi(input TimeSeries, period int) (TimeSeries, error) {
	return RSI(input, With("period", period))
}
func Mom(input TimeSeries, period int) (TimeSeries, error) {
	return MOM(input, With("period", period))
}
func Roc(input TimeSeries, period int) (TimeSeries, error) {
	return ROC(input, With("period", period))
}
func Willr(input TimeSeries, period int) (TimeSeries, error) {
	return WILLR(input, With("period", period))
}
func Cci(input TimeSeries, period int) (TimeSeries, error) {
	return CCI(input, With("period", period))
}
func Atr(input TimeSeries, period int) (TimeSeries, error) {
	return ATR(input, With("period", period))
}
func Natr(input TimeSeries, period int) (TimeSeries, error) {
	return NATR(input, With("period", period))
}
func Trange(input TimeSeries) (TimeSeries, error) { return TRANGE(input) }
func Obv(input TimeSeries) (TimeSeries, error)    { return OBV(input) }
func Ad(input TimeSeries) (TimeSeries, error)     { return AD(input) }
func Stddev(input TimeSeries, period int) (TimeSeries, error) {
	return STDDEV(input, With("period", period))
}
func Var_(input TimeSeries, period int) (TimeSeries, error) {
	return VARIANCE(input, With("period", period))
}
func Avgprice(input TimeSeries) (TimeSeries, error) { return AVGPRICE(input) }
func Medprice(input TimeSeries) (TimeSeries, error) { return MEDPRICE(input) }
func Typprice(input TimeSeries) (TimeSeries, error) { return TYPPRICE(input) }
func Wclprice(input TimeSeries) (TimeSeries, error) { return WCLPRICE(input) }
func Hma(input TimeSeries, period int) (TimeSeries, error) {
	return HMA(input, With("period", period))
}
func Vwap(input TimeSeries) (TimeSeries, error)       { return VWAP(input) }
func Heikinashi(input TimeSeries) (TimeSeries, error) { return HEIKINASHI(input) }
func Hlc3(input TimeSeries) (TimeSeries, error)       { return HLC3(input) }
func Ohlc4(input TimeSeries) (TimeSeries, error)      { return OHLC4(input) }
