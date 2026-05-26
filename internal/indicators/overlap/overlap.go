// Package overlap implements TA-Lib's Overlap Studies indicators.
package overlap

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/indicators"
	"github.com/rangertaha/gotal/internal/indicators/util"
)

type base struct{ Name, Source string }
type windowed struct {
	base
	Period int
	Alpha  float64
}

func cfg(opts []internal.ConfigOption, name string, defaultPeriod int) windowed {
	c := util.Cfg(opts...)
	return windowed{
		base:   base{Name: c.GetStr("name", name), Source: c.GetStr("source", "close")},
		Period: c.GetInt("period", defaultPeriod),
		Alpha:  float64(c.GetFloat("alpha", 0)),
	}
}

// ---------- WMA ----------

type wma struct{ windowed }

func newWMA(opts ...internal.ConfigOption) (internal.Indicator, error) {
	return &wma{cfg(opts, "wma", 14)}, nil
}
func (i *wma) Reset() error                          { return nil }
func (i *wma) Ready() bool                           { return true }
func (i *wma) Process(t internal.Tick) internal.Tick { return t }
func (i *wma) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		n := len(src)
		out := make([]float64, n)
		if i.Period <= 0 || n < i.Period {
			return out
		}
		denom := float64(i.Period*(i.Period+1)) / 2
		for j := i.Period - 1; j < n; j++ {
			var s float64
			for k := 0; k < i.Period; k++ {
				s += src[j-i.Period+1+k] * float64(k+1)
			}
			out[j] = s / denom
		}
		return out
	})
}

// ---------- DEMA ----------

type dema struct{ windowed }

func newDEMA(opts ...internal.ConfigOption) (internal.Indicator, error) {
	return &dema{cfg(opts, "dema", 14)}, nil
}
func (i *dema) Reset() error                          { return nil }
func (i *dema) Ready() bool                           { return true }
func (i *dema) Process(t internal.Tick) internal.Tick { return t }
func (i *dema) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		e1 := util.EMA(src, i.Period, i.Alpha)
		e2 := util.EMA(e1, i.Period, i.Alpha)
		out := make([]float64, len(src))
		for j := range out {
			out[j] = 2*e1[j] - e2[j]
		}
		return out
	})
}

// ---------- TEMA ----------

type tema struct{ windowed }

func newTEMA(opts ...internal.ConfigOption) (internal.Indicator, error) {
	return &tema{cfg(opts, "tema", 14)}, nil
}
func (i *tema) Reset() error                          { return nil }
func (i *tema) Ready() bool                           { return true }
func (i *tema) Process(t internal.Tick) internal.Tick { return t }
func (i *tema) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		e1 := util.EMA(src, i.Period, i.Alpha)
		e2 := util.EMA(e1, i.Period, i.Alpha)
		e3 := util.EMA(e2, i.Period, i.Alpha)
		out := make([]float64, len(src))
		for j := range out {
			out[j] = 3*e1[j] - 3*e2[j] + e3[j]
		}
		return out
	})
}

// ---------- TRIMA ----------

type trima struct{ windowed }

func newTRIMA(opts ...internal.ConfigOption) (internal.Indicator, error) {
	return &trima{cfg(opts, "trima", 14)}, nil
}
func (i *trima) Reset() error                          { return nil }
func (i *trima) Ready() bool                           { return true }
func (i *trima) Process(t internal.Tick) internal.Tick { return t }
func (i *trima) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		half := i.Period / 2
		if i.Period%2 == 0 {
			return util.SMA(util.SMA(src, half), half+1)
		}
		return util.SMA(util.SMA(src, half+1), half+1)
	})
}

// ---------- MIDPOINT ----------

type midpoint struct{ windowed }

func newMIDPOINT(opts ...internal.ConfigOption) (internal.Indicator, error) {
	return &midpoint{cfg(opts, "midpoint", 14)}, nil
}
func (i *midpoint) Reset() error                          { return nil }
func (i *midpoint) Ready() bool                           { return true }
func (i *midpoint) Process(t internal.Tick) internal.Tick { return t }
func (i *midpoint) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		hi := util.RollingMax(src, i.Period)
		lo := util.RollingMin(src, i.Period)
		out := make([]float64, len(src))
		for j := range out {
			if hi[j] != 0 || lo[j] != 0 {
				out[j] = (hi[j] + lo[j]) / 2
			}
		}
		return out
	})
}

// ---------- MIDPRICE ----------

type midprice struct {
	Name   string
	Period int
}

func newMIDPRICE(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &midprice{
		Name:   c.GetStr("name", "midprice"),
		Period: c.GetInt("period", 14),
	}, nil
}
func (i *midprice) Reset() error                          { return nil }
func (i *midprice) Ready() bool                           { return true }
func (i *midprice) Process(t internal.Tick) internal.Tick { return t }
func (i *midprice) Compute(ts internal.TimeSeries) internal.TimeSeries {
	highs := util.RollingMax(util.FieldOf(ts, "high"), i.Period)
	lows := util.RollingMin(util.FieldOf(ts, "low"), i.Period)
	out := make([]float64, len(highs))
	for j := range out {
		if highs[j] != 0 || lows[j] != 0 {
			out[j] = (highs[j] + lows[j]) / 2
		}
	}
	return util.AttachField(ts, i.Name, out)
}

// ---------- BBANDS ----------
//
// Outputs <name>_upper, <name>_middle, <name>_lower.

type bbands struct {
	Name, Source string
	Period       int
	NbDevUp      float64
	NbDevDn      float64
}

func newBBANDS(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &bbands{
		Name: c.GetStr("name", "bbands"), Source: c.GetStr("source", "close"),
		Period:  c.GetInt("period", 20),
		NbDevUp: float64(c.GetFloat("nbdevup", 2)),
		NbDevDn: float64(c.GetFloat("nbdevdn", 2)),
	}, nil
}
func (i *bbands) Reset() error                          { return nil }
func (i *bbands) Ready() bool                           { return true }
func (i *bbands) Process(t internal.Tick) internal.Tick { return t }
func (i *bbands) Compute(ts internal.TimeSeries) internal.TimeSeries {
	src := util.FieldOf(ts, i.Source)
	n := len(src)
	mid := util.SMA(src, i.Period)
	upper := make([]float64, n)
	lower := make([]float64, n)
	if i.Period <= 1 || n < i.Period {
		util.AttachField(ts, i.Name+"_upper", upper)
		util.AttachField(ts, i.Name+"_middle", mid)
		util.AttachField(ts, i.Name+"_lower", lower)
		return ts
	}
	for j := i.Period - 1; j < n; j++ {
		var sq float64
		for k := j - i.Period + 1; k <= j; k++ {
			d := src[k] - mid[j]
			sq += d * d
		}
		sd := mathSqrt(sq / float64(i.Period))
		upper[j] = mid[j] + i.NbDevUp*sd
		lower[j] = mid[j] - i.NbDevDn*sd
	}
	util.AttachField(ts, i.Name+"_upper", upper)
	util.AttachField(ts, i.Name+"_middle", mid)
	util.AttachField(ts, i.Name+"_lower", lower)
	return ts
}

func mathSqrt(x float64) float64 {
	if x < 0 {
		return 0
	}
	// Newton-Raphson; small input range, keep dependency-free.
	z := x
	for k := 0; k < 30; k++ {
		if z == 0 {
			return 0
		}
		z = (z + x/z) / 2
	}
	return z
}

// ---------- HMA (Hull Moving Average) ----------
//
// HMA(n) = WMA( 2*WMA(close, n/2) - WMA(close, n), sqrt(n) )

type hma struct {
	Name, Source string
	Period       int
}

func newHMA(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &hma{
		Name: c.GetStr("name", "hma"), Source: c.GetStr("source", "close"),
		Period: c.GetInt("period", 14),
	}, nil
}
func (i *hma) Reset() error                          { return nil }
func (i *hma) Ready() bool                           { return true }
func (i *hma) Process(t internal.Tick) internal.Tick { return t }
func (i *hma) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		halfN := i.Period / 2
		if halfN < 1 {
			halfN = 1
		}
		wmaHalf := computeWMA(src, halfN)
		wmaFull := computeWMA(src, i.Period)
		diff := make([]float64, len(src))
		for j := range diff {
			diff[j] = 2*wmaHalf[j] - wmaFull[j]
		}
		sqrtN := int(mathSqrt(float64(i.Period)))
		if sqrtN < 1 {
			sqrtN = 1
		}
		return computeWMA(diff, sqrtN)
	})
}

func computeWMA(src []float64, period int) []float64 {
	n := len(src)
	out := make([]float64, n)
	if period <= 0 || n < period {
		return out
	}
	denom := float64(period*(period+1)) / 2
	for j := period - 1; j < n; j++ {
		var s float64
		for k := 0; k < period; k++ {
			s += src[j-period+1+k] * float64(k+1)
		}
		out[j] = s / denom
	}
	return out
}

// All runs every implemented Overlap Studies indicator against ts and returns
// the (mutated) series with each indicator's output attached as a named field.
// Best-effort: stubs and indicators that can't construct with the given opts
// are skipped silently.
func All(ts internal.TimeSeries, opts ...internal.ConfigOption) internal.TimeSeries {
	return util.RunGroup(ts, indicators.OVERLAP, opts...)
}

func init() {
	util.Must(indicators.Add("WMA", newWMA, indicators.OVERLAP))
	util.Must(indicators.Add("DEMA", newDEMA, indicators.OVERLAP))
	util.Must(indicators.Add("TEMA", newTEMA, indicators.OVERLAP))
	util.Must(indicators.Add("TRIMA", newTRIMA, indicators.OVERLAP))
	util.Must(indicators.Add("MIDPOINT", newMIDPOINT, indicators.OVERLAP))
	util.Must(indicators.Add("MIDPRICE", newMIDPRICE, indicators.OVERLAP))
	util.Must(indicators.Add("BBANDS", newBBANDS, indicators.OVERLAP))
	util.Must(indicators.Add("HMA", newHMA, indicators.OVERLAP))
}
