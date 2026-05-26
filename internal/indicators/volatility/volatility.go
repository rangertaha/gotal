// Package volatility implements TA-Lib's Volatility indicators.
package volatility

import (
	"math"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/indicators"
	"github.com/rangertaha/gotal/internal/indicators/util"
)

func trueRange(highs, lows, closes []float64) []float64 {
	n := len(highs)
	out := make([]float64, n)
	if n == 0 {
		return out
	}
	out[0] = highs[0] - lows[0]
	for j := 1; j < n; j++ {
		hl := highs[j] - lows[j]
		hpc := math.Abs(highs[j] - closes[j-1])
		lpc := math.Abs(lows[j] - closes[j-1])
		out[j] = math.Max(hl, math.Max(hpc, lpc))
	}
	return out
}

// ---------- TRANGE ----------

type trange struct{ Name string }

func newTRANGE(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &trange{Name: c.GetStr("name", "trange")}, nil
}
func (i *trange) Reset() error                          { return nil }
func (i *trange) Ready() bool                           { return true }
func (i *trange) Process(t internal.Tick) internal.Tick { return t }
func (i *trange) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.AttachField(ts, i.Name, trueRange(
		util.FieldOf(ts, "high"),
		util.FieldOf(ts, "low"),
		util.FieldOf(ts, "close"),
	))
}

// ---------- ATR ----------

type atr struct{ Name string; Period int }

func newATR(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &atr{Name: c.GetStr("name", "atr"), Period: c.GetInt("period", 14)}, nil
}
func (i *atr) Reset() error                          { return nil }
func (i *atr) Ready() bool                           { return true }
func (i *atr) Process(t internal.Tick) internal.Tick { return t }
func (i *atr) Compute(ts internal.TimeSeries) internal.TimeSeries {
	tr := trueRange(util.FieldOf(ts, "high"), util.FieldOf(ts, "low"), util.FieldOf(ts, "close"))
	n := len(tr)
	out := make([]float64, n)
	if i.Period <= 0 || n < i.Period {
		return util.AttachField(ts, i.Name, out)
	}
	var sum float64
	for j := 0; j < i.Period; j++ {
		sum += tr[j]
	}
	out[i.Period-1] = sum / float64(i.Period)
	for j := i.Period; j < n; j++ {
		out[j] = (out[j-1]*float64(i.Period-1) + tr[j]) / float64(i.Period)
	}
	return util.AttachField(ts, i.Name, out)
}

// ---------- NATR ----------

type natr struct{ Name string; Period int }

func newNATR(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &natr{Name: c.GetStr("name", "natr"), Period: c.GetInt("period", 14)}, nil
}
func (i *natr) Reset() error                          { return nil }
func (i *natr) Ready() bool                           { return true }
func (i *natr) Process(t internal.Tick) internal.Tick { return t }
func (i *natr) Compute(ts internal.TimeSeries) internal.TimeSeries {
	closes := util.FieldOf(ts, "close")
	tr := trueRange(util.FieldOf(ts, "high"), util.FieldOf(ts, "low"), closes)
	n := len(tr)
	out := make([]float64, n)
	if i.Period <= 0 || n < i.Period {
		return util.AttachField(ts, i.Name, out)
	}
	var sum float64
	for j := 0; j < i.Period; j++ {
		sum += tr[j]
	}
	atr := sum / float64(i.Period)
	if closes[i.Period-1] != 0 {
		out[i.Period-1] = atr / closes[i.Period-1] * 100
	}
	for j := i.Period; j < n; j++ {
		atr = (atr*float64(i.Period-1) + tr[j]) / float64(i.Period)
		if closes[j] != 0 {
			out[j] = atr / closes[j] * 100
		}
	}
	return util.AttachField(ts, i.Name, out)
}

// All runs every implemented Volatility indicator against ts.
func All(ts internal.TimeSeries, opts ...internal.ConfigOption) internal.TimeSeries {
	return util.RunGroup(ts, indicators.VOLATILITY, opts...)
}

func init() {
	util.Must(indicators.Add("TRANGE", newTRANGE, indicators.VOLATILITY))
	util.Must(indicators.Add("ATR", newATR, indicators.VOLATILITY))
	util.Must(indicators.Add("NATR", newNATR, indicators.VOLATILITY))
}
