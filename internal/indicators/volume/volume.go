// Package volume implements TA-Lib's Volume indicators.
package volume

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/indicators"
	"github.com/rangertaha/gotal/internal/indicators/util"
)

// ---------- OBV ----------

type obv struct{ Name string }

func newOBV(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &obv{Name: c.GetStr("name", "obv")}, nil
}
func (i *obv) Reset() error                          { return nil }
func (i *obv) Ready() bool                           { return true }
func (i *obv) Process(t internal.Tick) internal.Tick { return t }
func (i *obv) Compute(ts internal.TimeSeries) internal.TimeSeries {
	closes := util.FieldOf(ts, "close")
	vols := util.FieldOf(ts, "volume")
	n := len(closes)
	out := make([]float64, n)
	if n == 0 {
		return util.AttachField(ts, i.Name, out)
	}
	out[0] = vols[0]
	for j := 1; j < n; j++ {
		switch {
		case closes[j] > closes[j-1]:
			out[j] = out[j-1] + vols[j]
		case closes[j] < closes[j-1]:
			out[j] = out[j-1] - vols[j]
		default:
			out[j] = out[j-1]
		}
	}
	return util.AttachField(ts, i.Name, out)
}

// ---------- AD (Chaikin A/D Line) ----------

type ad struct{ Name string }

func newAD(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &ad{Name: c.GetStr("name", "ad")}, nil
}
func (i *ad) Reset() error                          { return nil }
func (i *ad) Ready() bool                           { return true }
func (i *ad) Process(t internal.Tick) internal.Tick { return t }
func (i *ad) Compute(ts internal.TimeSeries) internal.TimeSeries {
	highs := util.FieldOf(ts, "high")
	lows := util.FieldOf(ts, "low")
	closes := util.FieldOf(ts, "close")
	vols := util.FieldOf(ts, "volume")
	n := len(closes)
	out := make([]float64, n)
	var cum float64
	for j := 0; j < n; j++ {
		rng := highs[j] - lows[j]
		var mfm float64
		if rng != 0 {
			mfm = ((closes[j] - lows[j]) - (highs[j] - closes[j])) / rng
		}
		cum += mfm * vols[j]
		out[j] = cum
	}
	return util.AttachField(ts, i.Name, out)
}

// ---------- ADOSC (Chaikin A/D Oscillator) ----------
// ADOSC = EMA(AD, fast) - EMA(AD, slow)

type adosc struct {
	Name       string
	Fast, Slow int
}

func newADOSC(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &adosc{
		Name: c.GetStr("name", "adosc"),
		Fast: c.GetInt("fast", 3), Slow: c.GetInt("slow", 10),
	}, nil
}
func (i *adosc) Reset() error                          { return nil }
func (i *adosc) Ready() bool                           { return true }
func (i *adosc) Process(t internal.Tick) internal.Tick { return t }
func (i *adosc) Compute(ts internal.TimeSeries) internal.TimeSeries {
	highs := util.FieldOf(ts, "high")
	lows := util.FieldOf(ts, "low")
	closes := util.FieldOf(ts, "close")
	vols := util.FieldOf(ts, "volume")
	n := len(closes)
	adLine := make([]float64, n)
	var cum float64
	for j := 0; j < n; j++ {
		rng := highs[j] - lows[j]
		var mfm float64
		if rng != 0 {
			mfm = ((closes[j] - lows[j]) - (highs[j] - closes[j])) / rng
		}
		cum += mfm * vols[j]
		adLine[j] = cum
	}
	fast := util.EMA(adLine, i.Fast, 0)
	slow := util.EMA(adLine, i.Slow, 0)
	out := make([]float64, n)
	for j := range out {
		out[j] = fast[j] - slow[j]
	}
	return util.AttachField(ts, i.Name, out)
}

// ---------- MFI (Money Flow Index) ----------

type mfi struct {
	Name   string
	Period int
}

func newMFI(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &mfi{Name: c.GetStr("name", "mfi"), Period: c.GetInt("period", 14)}, nil
}
func (i *mfi) Reset() error                          { return nil }
func (i *mfi) Ready() bool                           { return true }
func (i *mfi) Process(t internal.Tick) internal.Tick { return t }
func (i *mfi) Compute(ts internal.TimeSeries) internal.TimeSeries {
	highs := util.FieldOf(ts, "high")
	lows := util.FieldOf(ts, "low")
	closes := util.FieldOf(ts, "close")
	vols := util.FieldOf(ts, "volume")
	n := len(closes)
	out := make([]float64, n)
	if i.Period <= 0 || n <= i.Period {
		return util.AttachField(ts, i.Name, out)
	}
	tp := make([]float64, n)
	for j := 0; j < n; j++ {
		tp[j] = (highs[j] + lows[j] + closes[j]) / 3
	}
	posFlow := make([]float64, n)
	negFlow := make([]float64, n)
	for j := 1; j < n; j++ {
		rmf := tp[j] * vols[j]
		switch {
		case tp[j] > tp[j-1]:
			posFlow[j] = rmf
		case tp[j] < tp[j-1]:
			negFlow[j] = rmf
		}
	}
	for j := i.Period; j < n; j++ {
		var pos, neg float64
		for k := j - i.Period + 1; k <= j; k++ {
			pos += posFlow[k]
			neg += negFlow[k]
		}
		if neg == 0 {
			out[j] = 100
		} else {
			ratio := pos / neg
			out[j] = 100 - 100/(1+ratio)
		}
	}
	return util.AttachField(ts, i.Name, out)
}

// ---------- VWMA (Volume Weighted Moving Average) ----------

type vwma struct {
	Name, Source string
	Period       int
}

func newVWMA(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &vwma{
		Name: c.GetStr("name", "vwma"), Source: c.GetStr("source", "close"),
		Period: c.GetInt("period", 14),
	}, nil
}
func (i *vwma) Reset() error                          { return nil }
func (i *vwma) Ready() bool                           { return true }
func (i *vwma) Process(t internal.Tick) internal.Tick { return t }
func (i *vwma) Compute(ts internal.TimeSeries) internal.TimeSeries {
	src := util.FieldOf(ts, i.Source)
	vols := util.FieldOf(ts, "volume")
	n := len(src)
	out := make([]float64, n)
	if i.Period <= 0 || n < i.Period {
		return util.AttachField(ts, i.Name, out)
	}
	for j := i.Period - 1; j < n; j++ {
		var num, den float64
		for k := j - i.Period + 1; k <= j; k++ {
			num += src[k] * vols[k]
			den += vols[k]
		}
		if den != 0 {
			out[j] = num / den
		}
	}
	return util.AttachField(ts, i.Name, out)
}

// ---------- VWAP (cumulative session VWAP) ----------

type vwap struct{ Name string }

func newVWAP(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &vwap{Name: c.GetStr("name", "vwap")}, nil
}
func (i *vwap) Reset() error                          { return nil }
func (i *vwap) Ready() bool                           { return true }
func (i *vwap) Process(t internal.Tick) internal.Tick { return t }
func (i *vwap) Compute(ts internal.TimeSeries) internal.TimeSeries {
	highs := util.FieldOf(ts, "high")
	lows := util.FieldOf(ts, "low")
	closes := util.FieldOf(ts, "close")
	vols := util.FieldOf(ts, "volume")
	n := len(closes)
	out := make([]float64, n)
	var cumPV, cumV float64
	for j := 0; j < n; j++ {
		tp := (highs[j] + lows[j] + closes[j]) / 3
		cumPV += tp * vols[j]
		cumV += vols[j]
		if cumV != 0 {
			out[j] = cumPV / cumV
		}
	}
	return util.AttachField(ts, i.Name, out)
}

// All runs every implemented Volume indicator against ts.
func All(ts internal.TimeSeries, opts ...internal.ConfigOption) internal.TimeSeries {
	return util.RunGroup(ts, indicators.VOLUME, opts...)
}

func init() {
	util.Must(indicators.Add("OBV", newOBV, indicators.VOLUME))
	util.Must(indicators.Add("AD", newAD, indicators.VOLUME))
	util.Must(indicators.Add("ADOSC", newADOSC, indicators.VOLUME))
	util.Must(indicators.Add("MFI", newMFI, indicators.MOMENTUM))
	util.Must(indicators.Add("VWMA", newVWMA, indicators.VOLUME))
	util.Must(indicators.Add("VWAP", newVWAP, indicators.VOLUME))
}
