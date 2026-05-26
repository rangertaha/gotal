// Copyright (C) 2026 Rangertaha <rangertaha@gmail.com>
// SPDX-License-Identifier: GPL-3.0-or-later
package gotal

import (
	"github.com/rangertaha/gotal/internal/indicators"
	"github.com/rangertaha/gotal/internal/indicators/util"
)

// Overlap runs every implemented Overlap Studies indicator (SMA, EMA, WMA,
// DEMA, TEMA, TRIMA, MIDPOINT, MIDPRICE, BBANDS, HMA, ...) against ts and
// returns the (mutated) series with each indicator's output attached.
// Best-effort: stubs and indicators that can't run with the given opts are
// skipped silently.
func Overlap(ts TimeSeries, opts ...ConfigOption) TimeSeries {
	return util.RunGroup(ts, indicators.OVERLAP, opts...)
}

// Momentum runs every implemented Momentum indicator (RSI, MACD, MOM, ROC,
// STOCH, AROON, TRIX, CCI, WILLR, CMO, APO, PPO, ...) against ts.
func Momentum(ts TimeSeries, opts ...ConfigOption) TimeSeries {
	return util.RunGroup(ts, indicators.MOMENTUM, opts...)
}

// Volume runs every implemented Volume indicator (OBV, AD, ADOSC, VWMA, VWAP, ...).
func Volume(ts TimeSeries, opts ...ConfigOption) TimeSeries {
	return util.RunGroup(ts, indicators.VOLUME, opts...)
}

// Volatility runs every implemented Volatility indicator (TRANGE, ATR, NATR).
func Volatility(ts TimeSeries, opts ...ConfigOption) TimeSeries {
	return util.RunGroup(ts, indicators.VOLATILITY, opts...)
}

// Cycle runs every implemented Cycle / Hilbert Transform indicator. All of
// these are currently stubs, so the function is a no-op until they're filled in.
func Cycle(ts TimeSeries, opts ...ConfigOption) TimeSeries {
	return util.RunGroup(ts, indicators.CYCLE, opts...)
}

// Pattern runs every implemented Candlestick Pattern recognizer. Currently
// all CDL* patterns are stubs, so this is a no-op until they're filled in.
func Pattern(ts TimeSeries, opts ...ConfigOption) TimeSeries {
	return util.RunGroup(ts, indicators.PATTERN, opts...)
}

// Price runs every implemented Price Transform (AVGPRICE, MEDPRICE, TYPPRICE,
// WCLPRICE, HLC3, OHLC4, HEIKINASHI, ...).
func Price(ts TimeSeries, opts ...ConfigOption) TimeSeries {
	return util.RunGroup(ts, indicators.PRICE, opts...)
}

// Statistic runs every implemented Statistic Function (STDDEV, VARIANCE,
// LINEARREG family, TSF, BETA, CORREL).
func Statistic(ts TimeSeries, opts ...ConfigOption) TimeSeries {
	return util.RunGroup(ts, indicators.STATISTIC, opts...)
}

// MathOps runs every implemented Math Operator (MAX, MIN, SUMWINDOW, ADD,
// SUB, MULT, DIV, MAXINDEX, MININDEX).
func MathOps(ts TimeSeries, opts ...ConfigOption) TimeSeries {
	return util.RunGroup(ts, indicators.MATHOP, opts...)
}

// MathTransforms runs every implemented Math Transform (ACOS, ASIN, ATAN,
// COS, SIN, TAN, COSH, SINH, TANH, CEIL, FLOOR, EXP, LN, LOG10, SQRT).
func MathTransforms(ts TimeSeries, opts ...ConfigOption) TimeSeries {
	return util.RunGroup(ts, indicators.MATHTRANSFORM, opts...)
}

// Signal runs every implemented buy/sell/hold signal generator (RSISIGNAL,
// MFISIGNAL, CCISIGNAL, WILLRSIGNAL, MOMSIGNAL, ROCSIGNAL, TRIXSIGNAL,
// PPOSIGNAL, CMOSIGNAL, MACDSIGNAL_CROSS, STOCHSIGNAL, ...). Each output
// field carries +100 (buy), -100 (sell), or 0 (hold) per tick — the same
// convention TA-Lib uses for candlestick pattern outputs.
func Signal(ts TimeSeries, opts ...ConfigOption) TimeSeries {
	return util.RunGroup(ts, indicators.SIGNAL, opts...)
}

// RunAll runs every implemented indicator across every group against ts. Use
// this for exploratory passes — the resulting series will have 80+ attached
// fields. Stubs are skipped silently.
func RunAll(ts TimeSeries, opts ...ConfigOption) TimeSeries {
	return util.RunIDs(ts, indicators.All(), opts...)
}

// Group returns every indicator id registered under the given group.
func Group(g string) []string {
	return indicators.Group(indicators.GroupType(g))
}

// Groups returns every known group constant.
func Groups() []string {
	return []string{
		string(indicators.OVERLAP),
		string(indicators.MOMENTUM),
		string(indicators.VOLATILITY),
		string(indicators.VOLUME),
		string(indicators.CYCLE),
		string(indicators.PATTERN),
		string(indicators.PRICE),
		string(indicators.STATISTIC),
		string(indicators.MATHOP),
		string(indicators.MATHTRANSFORM),
		string(indicators.SIGNAL),
		string(indicators.TREND),
		string(indicators.OTHER),
	}
}
