// Package statistic implements TA-Lib's Statistic Functions.
package statistic

import (
	"math"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/indicators"
	"github.com/rangertaha/gotal/internal/indicators/util"
)

func rollingMoment(src []float64, period int, sample bool) (variance, stddev []float64) {
	n := len(src)
	variance = make([]float64, n)
	stddev = make([]float64, n)
	if period <= 1 || n < period {
		return
	}
	denom := float64(period)
	if sample {
		denom = float64(period - 1)
	}
	for i := period - 1; i < n; i++ {
		var sum float64
		for j := i - period + 1; j <= i; j++ {
			sum += src[j]
		}
		mean := sum / float64(period)
		var sq float64
		for j := i - period + 1; j <= i; j++ {
			d := src[j] - mean
			sq += d * d
		}
		variance[i] = sq / denom
		stddev[i] = math.Sqrt(variance[i])
	}
	return
}

// ---------- STDDEV ----------

type stddev struct{ Name, Source string; Period int; NbDev float64 }

func newSTDDEV(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &stddev{
		Name: c.GetStr("name", "stddev"), Source: c.GetStr("source", "close"),
		Period: c.GetInt("period", 5),
		NbDev:  float64(c.GetFloat("nbdev", 1)),
	}, nil
}
func (i *stddev) Reset() error                          { return nil }
func (i *stddev) Ready() bool                           { return true }
func (i *stddev) Process(t internal.Tick) internal.Tick { return t }
func (i *stddev) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		_, sd := rollingMoment(src, i.Period, false)
		if i.NbDev != 1 {
			for j := range sd {
				sd[j] *= i.NbDev
			}
		}
		return sd
	})
}

// ---------- VARIANCE ----------

type variance struct{ Name, Source string; Period int }

func newVARIANCE(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &variance{
		Name: c.GetStr("name", "variance"), Source: c.GetStr("source", "close"),
		Period: c.GetInt("period", 5),
	}, nil
}
func (i *variance) Reset() error                          { return nil }
func (i *variance) Ready() bool                           { return true }
func (i *variance) Process(t internal.Tick) internal.Tick { return t }
func (i *variance) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		v, _ := rollingMoment(src, i.Period, false)
		return v
	})
}

// linreg returns slope and intercept for a window of length period ending at index j.
// Uses x = 0..period-1.
func linreg(src []float64, j, period int) (slope, intercept float64) {
	start := j - period + 1
	var sumX, sumY, sumXY, sumX2 float64
	pf := float64(period)
	for k := 0; k < period; k++ {
		x := float64(k)
		y := src[start+k]
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
	}
	denom := pf*sumX2 - sumX*sumX
	if denom == 0 {
		return 0, sumY / pf
	}
	slope = (pf*sumXY - sumX*sumY) / denom
	intercept = (sumY - slope*sumX) / pf
	return
}

// ---------- LINEARREG ----------

type linearreg struct {
	Name, Source string
	Period       int
	mode         string // "value" | "slope" | "intercept" | "angle" | "tsf"
}

func (i *linearreg) Reset() error                          { return nil }
func (i *linearreg) Ready() bool                           { return true }
func (i *linearreg) Process(t internal.Tick) internal.Tick { return t }
func (i *linearreg) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		n := len(src)
		out := make([]float64, n)
		if i.Period <= 1 || n < i.Period {
			return out
		}
		for j := i.Period - 1; j < n; j++ {
			slope, intercept := linreg(src, j, i.Period)
			switch i.mode {
			case "slope":
				out[j] = slope
			case "intercept":
				out[j] = intercept
			case "angle":
				out[j] = math.Atan(slope) * 180 / math.Pi
			case "tsf":
				out[j] = intercept + slope*float64(i.Period) // one step ahead
			default:
				out[j] = intercept + slope*float64(i.Period-1) // line value at the latest point
			}
		}
		return out
	})
}

func linregCtor(name, mode string) func(opts ...internal.ConfigOption) (internal.Indicator, error) {
	return func(opts ...internal.ConfigOption) (internal.Indicator, error) {
		c := util.Cfg(opts...)
		return &linearreg{
			Name: c.GetStr("name", name), Source: c.GetStr("source", "close"),
			Period: c.GetInt("period", 14),
			mode:   mode,
		}, nil
	}
}

// ---------- BETA / CORREL (binary) ----------

type binaryStat struct {
	Name, Source1, Source2 string
	Period                 int
	fn                     func(a, b []float64, period int) []float64
}

func (i *binaryStat) Reset() error                          { return nil }
func (i *binaryStat) Ready() bool                           { return true }
func (i *binaryStat) Process(t internal.Tick) internal.Tick { return t }
func (i *binaryStat) Compute(ts internal.TimeSeries) internal.TimeSeries {
	a := util.FieldOf(ts, i.Source1)
	b := util.FieldOf(ts, i.Source2)
	return util.AttachField(ts, i.Name, i.fn(a, b, i.Period))
}

func rollingBeta(a, b []float64, period int) []float64 {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	out := make([]float64, n)
	if period <= 1 || n < period {
		return out
	}
	for j := period - 1; j < n; j++ {
		var sa, sb, sab, sb2 float64
		for k := j - period + 1; k <= j; k++ {
			sa += a[k]
			sb += b[k]
			sab += a[k] * b[k]
			sb2 += b[k] * b[k]
		}
		pf := float64(period)
		denom := pf*sb2 - sb*sb
		if denom != 0 {
			out[j] = (pf*sab - sa*sb) / denom
		}
	}
	return out
}

func rollingCorrel(a, b []float64, period int) []float64 {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	out := make([]float64, n)
	if period <= 1 || n < period {
		return out
	}
	for j := period - 1; j < n; j++ {
		var sa, sb, sab, sa2, sb2 float64
		for k := j - period + 1; k <= j; k++ {
			sa += a[k]
			sb += b[k]
			sab += a[k] * b[k]
			sa2 += a[k] * a[k]
			sb2 += b[k] * b[k]
		}
		pf := float64(period)
		denom := math.Sqrt((pf*sa2 - sa*sa) * (pf*sb2 - sb*sb))
		if denom != 0 {
			out[j] = (pf*sab - sa*sb) / denom
		}
	}
	return out
}

func binaryStatCtor(name string, fn func([]float64, []float64, int) []float64) func(opts ...internal.ConfigOption) (internal.Indicator, error) {
	return func(opts ...internal.ConfigOption) (internal.Indicator, error) {
		c := util.Cfg(opts...)
		return &binaryStat{
			Name: c.GetStr("name", name),
			Source1: c.GetStr("source1", "close"),
			Source2: c.GetStr("source2", "open"),
			Period: c.GetInt("period", 5),
			fn: fn,
		}, nil
	}
}

// All runs every implemented Statistic indicator against ts.
func All(ts internal.TimeSeries, opts ...internal.ConfigOption) internal.TimeSeries {
	return util.RunGroup(ts, indicators.STATISTIC, opts...)
}

func init() {
	util.Must(indicators.Add("STDDEV", newSTDDEV, indicators.STATISTIC))
	util.Must(indicators.Add("VARIANCE", newVARIANCE, indicators.STATISTIC))
	util.Must(indicators.Add("LINEARREG", linregCtor("linearreg", "value"), indicators.STATISTIC))
	util.Must(indicators.Add("LINEARREG_SLOPE", linregCtor("linearreg_slope", "slope"), indicators.STATISTIC))
	util.Must(indicators.Add("LINEARREG_INTERCEPT", linregCtor("linearreg_intercept", "intercept"), indicators.STATISTIC))
	util.Must(indicators.Add("LINEARREG_ANGLE", linregCtor("linearreg_angle", "angle"), indicators.STATISTIC))
	util.Must(indicators.Add("TSF", linregCtor("tsf", "tsf"), indicators.STATISTIC))
	util.Must(indicators.Add("BETA", binaryStatCtor("beta", rollingBeta), indicators.STATISTIC))
	util.Must(indicators.Add("CORREL", binaryStatCtor("correl", rollingCorrel), indicators.STATISTIC))
}
