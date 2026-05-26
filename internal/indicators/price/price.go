// Package price implements TA-Lib's Price Transform indicators.
package price

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/indicators"
	"github.com/rangertaha/gotal/internal/indicators/util"
)

type priceFn struct {
	Name string
	fn   func(o, h, l, c float64) float64
}

func (i *priceFn) Reset() error                          { return nil }
func (i *priceFn) Ready() bool                           { return true }
func (i *priceFn) Process(t internal.Tick) internal.Tick { return t }
func (i *priceFn) Compute(ts internal.TimeSeries) internal.TimeSeries {
	opens := util.FieldOf(ts, "open")
	highs := util.FieldOf(ts, "high")
	lows := util.FieldOf(ts, "low")
	closes := util.FieldOf(ts, "close")
	n := len(closes)
	out := make([]float64, n)
	for j := 0; j < n; j++ {
		out[j] = i.fn(opens[j], highs[j], lows[j], closes[j])
	}
	return util.AttachField(ts, i.Name, out)
}

func ctor(name string, fn func(o, h, l, c float64) float64) func(opts ...internal.ConfigOption) (internal.Indicator, error) {
	return func(opts ...internal.ConfigOption) (internal.Indicator, error) {
		c := util.Cfg(opts...)
		return &priceFn{Name: c.GetStr("name", name), fn: fn}, nil
	}
}

// ---------- HEIKINASHI ----------
//
// Outputs <name>_open/_high/_low/_close. Recursive — open[i] uses prior HA.

type heikinashi struct{ Name string }

func newHEIKINASHI(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &heikinashi{Name: c.GetStr("name", "ha")}, nil
}
func (i *heikinashi) Reset() error                          { return nil }
func (i *heikinashi) Ready() bool                           { return true }
func (i *heikinashi) Process(t internal.Tick) internal.Tick { return t }
func (i *heikinashi) Compute(ts internal.TimeSeries) internal.TimeSeries {
	opens := util.FieldOf(ts, "open")
	highs := util.FieldOf(ts, "high")
	lows := util.FieldOf(ts, "low")
	closes := util.FieldOf(ts, "close")
	n := len(closes)
	ho := make([]float64, n)
	hh := make([]float64, n)
	hl := make([]float64, n)
	hc := make([]float64, n)
	if n == 0 {
		return ts
	}
	hc[0] = (opens[0] + highs[0] + lows[0] + closes[0]) / 4
	ho[0] = (opens[0] + closes[0]) / 2
	hh[0] = highs[0]
	hl[0] = lows[0]
	for j := 1; j < n; j++ {
		hc[j] = (opens[j] + highs[j] + lows[j] + closes[j]) / 4
		ho[j] = (ho[j-1] + hc[j-1]) / 2
		hh[j] = maxOf(highs[j], ho[j], hc[j])
		hl[j] = minOf(lows[j], ho[j], hc[j])
	}
	util.AttachField(ts, i.Name+"_open", ho)
	util.AttachField(ts, i.Name+"_high", hh)
	util.AttachField(ts, i.Name+"_low", hl)
	util.AttachField(ts, i.Name+"_close", hc)
	return ts
}

func maxOf(vals ...float64) float64 {
	m := vals[0]
	for _, v := range vals[1:] {
		if v > m {
			m = v
		}
	}
	return m
}
func minOf(vals ...float64) float64 {
	m := vals[0]
	for _, v := range vals[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

// All runs every implemented Price Transform indicator against ts.
func All(ts internal.TimeSeries, opts ...internal.ConfigOption) internal.TimeSeries {
	return util.RunGroup(ts, indicators.PRICE, opts...)
}

func init() {
	util.Must(indicators.Add("AVGPRICE", ctor("avgprice", func(o, h, l, c float64) float64 { return (o + h + l + c) / 4 }), indicators.PRICE))
	util.Must(indicators.Add("MEDPRICE", ctor("medprice", func(o, h, l, c float64) float64 { return (h + l) / 2 }), indicators.PRICE))
	util.Must(indicators.Add("TYPPRICE", ctor("typprice", func(o, h, l, c float64) float64 { return (h + l + c) / 3 }), indicators.PRICE))
	util.Must(indicators.Add("WCLPRICE", ctor("wclprice", func(o, h, l, c float64) float64 { return (h + l + 2*c) / 4 }), indicators.PRICE))
	util.Must(indicators.Add("HLC3", ctor("hlc3", func(o, h, l, c float64) float64 { return (h + l + c) / 3 }), indicators.PRICE))
	util.Must(indicators.Add("OHLC4", ctor("ohlc4", func(o, h, l, c float64) float64 { return (o + h + l + c) / 4 }), indicators.PRICE))
	util.Must(indicators.Add("HEIKINASHI", newHEIKINASHI, indicators.OTHER))
}
