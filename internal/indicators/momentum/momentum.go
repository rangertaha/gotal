// Package momentum implements TA-Lib's Momentum indicators.
package momentum

import (
	"math"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/indicators"
	"github.com/rangertaha/gotal/internal/indicators/util"
)

func compute(name, source string, period int, fn func(src []float64) []float64) func(internal.TimeSeries) internal.TimeSeries {
	return func(ts internal.TimeSeries) internal.TimeSeries {
		return util.Unary(ts, source, name, fn)
	}
}

// ---------- MOM ----------

type mom struct{ Name, Source string; Period int }

func newMOM(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &mom{
		Name: c.GetStr("name", "mom"), Source: c.GetStr("source", "close"),
		Period: c.GetInt("period", 10),
	}, nil
}
func (i *mom) Reset() error                          { return nil }
func (i *mom) Ready() bool                           { return true }
func (i *mom) Process(t internal.Tick) internal.Tick { return t }
func (i *mom) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		out := make([]float64, len(src))
		for j := i.Period; j < len(src); j++ {
			out[j] = src[j] - src[j-i.Period]
		}
		return out
	})
}

// ---------- ROC ----------

type roc struct{ Name, Source string; Period int }

func newROC(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &roc{
		Name: c.GetStr("name", "roc"), Source: c.GetStr("source", "close"),
		Period: c.GetInt("period", 10),
	}, nil
}
func (i *roc) Reset() error                          { return nil }
func (i *roc) Ready() bool                           { return true }
func (i *roc) Process(t internal.Tick) internal.Tick { return t }
func (i *roc) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		out := make([]float64, len(src))
		for j := i.Period; j < len(src); j++ {
			if src[j-i.Period] != 0 {
				out[j] = (src[j] - src[j-i.Period]) / src[j-i.Period] * 100
			}
		}
		return out
	})
}

// ---------- ROCP ----------

type rocp struct{ Name, Source string; Period int }

func newROCP(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &rocp{
		Name: c.GetStr("name", "rocp"), Source: c.GetStr("source", "close"),
		Period: c.GetInt("period", 10),
	}, nil
}
func (i *rocp) Reset() error                          { return nil }
func (i *rocp) Ready() bool                           { return true }
func (i *rocp) Process(t internal.Tick) internal.Tick { return t }
func (i *rocp) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		out := make([]float64, len(src))
		for j := i.Period; j < len(src); j++ {
			if src[j-i.Period] != 0 {
				out[j] = (src[j] - src[j-i.Period]) / src[j-i.Period]
			}
		}
		return out
	})
}

// ---------- ROCR ----------

type rocr struct{ Name, Source string; Period int }

func newROCR(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &rocr{
		Name: c.GetStr("name", "rocr"), Source: c.GetStr("source", "close"),
		Period: c.GetInt("period", 10),
	}, nil
}
func (i *rocr) Reset() error                          { return nil }
func (i *rocr) Ready() bool                           { return true }
func (i *rocr) Process(t internal.Tick) internal.Tick { return t }
func (i *rocr) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		out := make([]float64, len(src))
		for j := i.Period; j < len(src); j++ {
			if src[j-i.Period] != 0 {
				out[j] = src[j] / src[j-i.Period]
			}
		}
		return out
	})
}

// ---------- ROCR100 ----------

type rocr100 struct{ Name, Source string; Period int }

func newROCR100(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &rocr100{
		Name: c.GetStr("name", "rocr100"), Source: c.GetStr("source", "close"),
		Period: c.GetInt("period", 10),
	}, nil
}
func (i *rocr100) Reset() error                          { return nil }
func (i *rocr100) Ready() bool                           { return true }
func (i *rocr100) Process(t internal.Tick) internal.Tick { return t }
func (i *rocr100) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		out := make([]float64, len(src))
		for j := i.Period; j < len(src); j++ {
			if src[j-i.Period] != 0 {
				out[j] = src[j] / src[j-i.Period] * 100
			}
		}
		return out
	})
}

// ---------- WILLR (Williams' %R) ----------

type willr struct{ Name string; Period int }

func newWILLR(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &willr{Name: c.GetStr("name", "willr"), Period: c.GetInt("period", 14)}, nil
}
func (i *willr) Reset() error                          { return nil }
func (i *willr) Ready() bool                           { return true }
func (i *willr) Process(t internal.Tick) internal.Tick { return t }
func (i *willr) Compute(ts internal.TimeSeries) internal.TimeSeries {
	highs := util.FieldOf(ts, "high")
	lows := util.FieldOf(ts, "low")
	closes := util.FieldOf(ts, "close")
	hi := util.RollingMax(highs, i.Period)
	lo := util.RollingMin(lows, i.Period)
	out := make([]float64, len(closes))
	for j := i.Period - 1; j < len(closes); j++ {
		if hi[j]-lo[j] != 0 {
			out[j] = (hi[j] - closes[j]) / (hi[j] - lo[j]) * -100
		}
	}
	return util.AttachField(ts, i.Name, out)
}

// ---------- RSI ----------

type rsi struct{ Name, Source string; Period int }

func newRSI(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &rsi{
		Name: c.GetStr("name", "rsi"), Source: c.GetStr("source", "close"),
		Period: c.GetInt("period", 14),
	}, nil
}
func (i *rsi) Reset() error                          { return nil }
func (i *rsi) Ready() bool                           { return true }
func (i *rsi) Process(t internal.Tick) internal.Tick { return t }
func (i *rsi) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		n := len(src)
		out := make([]float64, n)
		if i.Period <= 0 || n <= i.Period {
			return out
		}
		var gain, loss float64
		for j := 1; j <= i.Period; j++ {
			d := src[j] - src[j-1]
			if d >= 0 {
				gain += d
			} else {
				loss -= d
			}
		}
		avgGain := gain / float64(i.Period)
		avgLoss := loss / float64(i.Period)
		out[i.Period] = rsiFromAvgs(avgGain, avgLoss)
		for j := i.Period + 1; j < n; j++ {
			d := src[j] - src[j-1]
			if d >= 0 {
				avgGain = (avgGain*float64(i.Period-1) + d) / float64(i.Period)
				avgLoss = (avgLoss * float64(i.Period-1)) / float64(i.Period)
			} else {
				avgGain = (avgGain * float64(i.Period-1)) / float64(i.Period)
				avgLoss = (avgLoss*float64(i.Period-1) + (-d)) / float64(i.Period)
			}
			out[j] = rsiFromAvgs(avgGain, avgLoss)
		}
		return out
	})
}

func rsiFromAvgs(g, l float64) float64 {
	if l == 0 {
		if g == 0 {
			return 50
		}
		return 100
	}
	rs := g / l
	return 100 - 100/(1+rs)
}

// ---------- CCI ----------

type cci struct{ Name string; Period int }

func newCCI(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &cci{Name: c.GetStr("name", "cci"), Period: c.GetInt("period", 14)}, nil
}
func (i *cci) Reset() error                          { return nil }
func (i *cci) Ready() bool                           { return true }
func (i *cci) Process(t internal.Tick) internal.Tick { return t }
func (i *cci) Compute(ts internal.TimeSeries) internal.TimeSeries {
	highs := util.FieldOf(ts, "high")
	lows := util.FieldOf(ts, "low")
	closes := util.FieldOf(ts, "close")
	n := len(closes)
	tp := make([]float64, n)
	for j := 0; j < n; j++ {
		tp[j] = (highs[j] + lows[j] + closes[j]) / 3
	}
	smaTP := util.SMA(tp, i.Period)
	out := make([]float64, n)
	for j := i.Period - 1; j < n; j++ {
		var mad float64
		for k := j - i.Period + 1; k <= j; k++ {
			mad += math.Abs(tp[k] - smaTP[j])
		}
		mad /= float64(i.Period)
		if mad != 0 {
			out[j] = (tp[j] - smaTP[j]) / (0.015 * mad)
		}
	}
	return util.AttachField(ts, i.Name, out)
}

// ---------- MACD ----------
//
// Outputs three signal fields: <name> (the MACD line), <name>_signal, <name>_hist.

type macd struct {
	Name, Source       string
	Fast, Slow, Signal int
}

func newMACD(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &macd{
		Name:   c.GetStr("name", "macd"),
		Source: c.GetStr("source", "close"),
		Fast:   c.GetInt("fast", 12),
		Slow:   c.GetInt("slow", 26),
		Signal: c.GetInt("signal", 9),
	}, nil
}
func (i *macd) Reset() error                          { return nil }
func (i *macd) Ready() bool                           { return true }
func (i *macd) Process(t internal.Tick) internal.Tick { return t }
func (i *macd) Compute(ts internal.TimeSeries) internal.TimeSeries {
	src := util.FieldOf(ts, i.Source)
	fast := util.EMA(src, i.Fast, 0)
	slow := util.EMA(src, i.Slow, 0)
	line := make([]float64, len(src))
	for j := range line {
		line[j] = fast[j] - slow[j]
	}
	signal := util.EMA(line, i.Signal, 0)
	hist := make([]float64, len(src))
	for j := range hist {
		hist[j] = line[j] - signal[j]
	}
	util.AttachField(ts, i.Name, line)
	util.AttachField(ts, i.Name+"_signal", signal)
	util.AttachField(ts, i.Name+"_hist", hist)
	return ts
}

// ---------- APO ----------

type apo struct {
	Name, Source string
	Fast, Slow   int
}

func newAPO(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &apo{
		Name:   c.GetStr("name", "apo"),
		Source: c.GetStr("source", "close"),
		Fast:   c.GetInt("fast", 12), Slow: c.GetInt("slow", 26),
	}, nil
}
func (i *apo) Reset() error                          { return nil }
func (i *apo) Ready() bool                           { return true }
func (i *apo) Process(t internal.Tick) internal.Tick { return t }
func (i *apo) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		fast := util.EMA(src, i.Fast, 0)
		slow := util.EMA(src, i.Slow, 0)
		out := make([]float64, len(src))
		for j := range out {
			out[j] = fast[j] - slow[j]
		}
		return out
	})
}

// ---------- PPO ----------

type ppo struct {
	Name, Source string
	Fast, Slow   int
}

func newPPO(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &ppo{
		Name:   c.GetStr("name", "ppo"),
		Source: c.GetStr("source", "close"),
		Fast:   c.GetInt("fast", 12), Slow: c.GetInt("slow", 26),
	}, nil
}
func (i *ppo) Reset() error                          { return nil }
func (i *ppo) Ready() bool                           { return true }
func (i *ppo) Process(t internal.Tick) internal.Tick { return t }
func (i *ppo) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		fast := util.EMA(src, i.Fast, 0)
		slow := util.EMA(src, i.Slow, 0)
		out := make([]float64, len(src))
		for j := range out {
			if slow[j] != 0 {
				out[j] = (fast[j] - slow[j]) / slow[j] * 100
			}
		}
		return out
	})
}

// ---------- CMO (Chande Momentum Oscillator) ----------

type cmo struct {
	Name, Source string
	Period       int
}

func newCMO(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &cmo{
		Name: c.GetStr("name", "cmo"), Source: c.GetStr("source", "close"),
		Period: c.GetInt("period", 14),
	}, nil
}
func (i *cmo) Reset() error                          { return nil }
func (i *cmo) Ready() bool                           { return true }
func (i *cmo) Process(t internal.Tick) internal.Tick { return t }
func (i *cmo) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		n := len(src)
		out := make([]float64, n)
		if i.Period <= 0 || n <= i.Period {
			return out
		}
		ups := make([]float64, n)
		dns := make([]float64, n)
		for j := 1; j < n; j++ {
			d := src[j] - src[j-1]
			if d > 0 {
				ups[j] = d
			} else {
				dns[j] = -d
			}
		}
		for j := i.Period; j < n; j++ {
			var su, sd float64
			for k := j - i.Period + 1; k <= j; k++ {
				su += ups[k]
				sd += dns[k]
			}
			if su+sd != 0 {
				out[j] = (su - sd) / (su + sd) * 100
			}
		}
		return out
	})
}

// ---------- TRIX ----------

type trix struct {
	Name, Source string
	Period       int
}

func newTRIX(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &trix{
		Name: c.GetStr("name", "trix"), Source: c.GetStr("source", "close"),
		Period: c.GetInt("period", 14),
	}, nil
}
func (i *trix) Reset() error                          { return nil }
func (i *trix) Ready() bool                           { return true }
func (i *trix) Process(t internal.Tick) internal.Tick { return t }
func (i *trix) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		e1 := util.EMA(src, i.Period, 0)
		e2 := util.EMA(e1, i.Period, 0)
		e3 := util.EMA(e2, i.Period, 0)
		out := make([]float64, len(src))
		for j := 1; j < len(src); j++ {
			if e3[j-1] != 0 {
				out[j] = (e3[j] - e3[j-1]) / e3[j-1] * 100
			}
		}
		return out
	})
}

// ---------- AROON ----------
//
// Outputs <name>_up and <name>_down.

type aroon struct {
	Name   string
	Period int
}

func newAROON(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &aroon{Name: c.GetStr("name", "aroon"), Period: c.GetInt("period", 14)}, nil
}
func (i *aroon) Reset() error                          { return nil }
func (i *aroon) Ready() bool                           { return true }
func (i *aroon) Process(t internal.Tick) internal.Tick { return t }
func (i *aroon) Compute(ts internal.TimeSeries) internal.TimeSeries {
	highs := util.FieldOf(ts, "high")
	lows := util.FieldOf(ts, "low")
	n := len(highs)
	up := make([]float64, n)
	dn := make([]float64, n)
	if i.Period <= 0 || n <= i.Period {
		util.AttachField(ts, i.Name+"_up", up)
		util.AttachField(ts, i.Name+"_down", dn)
		return ts
	}
	for j := i.Period; j < n; j++ {
		hiIdx, loIdx := j-i.Period, j-i.Period
		for k := j - i.Period + 1; k <= j; k++ {
			if highs[k] >= highs[hiIdx] {
				hiIdx = k
			}
			if lows[k] <= lows[loIdx] {
				loIdx = k
			}
		}
		up[j] = float64(i.Period-(j-hiIdx)) / float64(i.Period) * 100
		dn[j] = float64(i.Period-(j-loIdx)) / float64(i.Period) * 100
	}
	util.AttachField(ts, i.Name+"_up", up)
	util.AttachField(ts, i.Name+"_down", dn)
	return ts
}

// ---------- AROONOSC ----------

type aroonOsc struct {
	Name   string
	Period int
}

func newAROONOSC(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &aroonOsc{Name: c.GetStr("name", "aroonosc"), Period: c.GetInt("period", 14)}, nil
}
func (i *aroonOsc) Reset() error                          { return nil }
func (i *aroonOsc) Ready() bool                           { return true }
func (i *aroonOsc) Process(t internal.Tick) internal.Tick { return t }
func (i *aroonOsc) Compute(ts internal.TimeSeries) internal.TimeSeries {
	highs := util.FieldOf(ts, "high")
	lows := util.FieldOf(ts, "low")
	n := len(highs)
	out := make([]float64, n)
	if i.Period <= 0 || n <= i.Period {
		return util.AttachField(ts, i.Name, out)
	}
	for j := i.Period; j < n; j++ {
		hiIdx, loIdx := j-i.Period, j-i.Period
		for k := j - i.Period + 1; k <= j; k++ {
			if highs[k] >= highs[hiIdx] {
				hiIdx = k
			}
			if lows[k] <= lows[loIdx] {
				loIdx = k
			}
		}
		up := float64(i.Period-(j-hiIdx)) / float64(i.Period) * 100
		dn := float64(i.Period-(j-loIdx)) / float64(i.Period) * 100
		out[j] = up - dn
	}
	return util.AttachField(ts, i.Name, out)
}

// ---------- STOCH ----------
//
// Outputs <name>_k (fast %K smoothed) and <name>_d (signal).

type stoch struct {
	Name             string
	FastK, SlowK, SlowD int
}

func newSTOCH(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &stoch{
		Name:  c.GetStr("name", "stoch"),
		FastK: c.GetInt("fastk", 5),
		SlowK: c.GetInt("slowk", 3),
		SlowD: c.GetInt("slowd", 3),
	}, nil
}
func (i *stoch) Reset() error                          { return nil }
func (i *stoch) Ready() bool                           { return true }
func (i *stoch) Process(t internal.Tick) internal.Tick { return t }
func (i *stoch) Compute(ts internal.TimeSeries) internal.TimeSeries {
	highs := util.FieldOf(ts, "high")
	lows := util.FieldOf(ts, "low")
	closes := util.FieldOf(ts, "close")
	n := len(closes)
	fastK := make([]float64, n)
	hi := util.RollingMax(highs, i.FastK)
	lo := util.RollingMin(lows, i.FastK)
	for j := i.FastK - 1; j < n; j++ {
		if hi[j]-lo[j] != 0 {
			fastK[j] = (closes[j] - lo[j]) / (hi[j] - lo[j]) * 100
		}
	}
	slowK := util.SMA(fastK, i.SlowK)
	slowD := util.SMA(slowK, i.SlowD)
	util.AttachField(ts, i.Name+"_k", slowK)
	util.AttachField(ts, i.Name+"_d", slowD)
	return ts
}

// ---------- STOCHF (Fast Stochastic) ----------

type stochF struct {
	Name        string
	FastK, FastD int
}

func newSTOCHF(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &stochF{
		Name:  c.GetStr("name", "stochf"),
		FastK: c.GetInt("fastk", 5),
		FastD: c.GetInt("fastd", 3),
	}, nil
}
func (i *stochF) Reset() error                          { return nil }
func (i *stochF) Ready() bool                           { return true }
func (i *stochF) Process(t internal.Tick) internal.Tick { return t }
func (i *stochF) Compute(ts internal.TimeSeries) internal.TimeSeries {
	highs := util.FieldOf(ts, "high")
	lows := util.FieldOf(ts, "low")
	closes := util.FieldOf(ts, "close")
	n := len(closes)
	fastK := make([]float64, n)
	hi := util.RollingMax(highs, i.FastK)
	lo := util.RollingMin(lows, i.FastK)
	for j := i.FastK - 1; j < n; j++ {
		if hi[j]-lo[j] != 0 {
			fastK[j] = (closes[j] - lo[j]) / (hi[j] - lo[j]) * 100
		}
	}
	fastD := util.SMA(fastK, i.FastD)
	util.AttachField(ts, i.Name+"_k", fastK)
	util.AttachField(ts, i.Name+"_d", fastD)
	return ts
}

// All runs every implemented Momentum indicator against ts.
func All(ts internal.TimeSeries, opts ...internal.ConfigOption) internal.TimeSeries {
	return util.RunGroup(ts, indicators.MOMENTUM, opts...)
}

func init() {
	util.Must(indicators.Add("MOM", newMOM, indicators.MOMENTUM))
	util.Must(indicators.Add("ROC", newROC, indicators.MOMENTUM))
	util.Must(indicators.Add("ROCP", newROCP, indicators.MOMENTUM))
	util.Must(indicators.Add("ROCR", newROCR, indicators.MOMENTUM))
	util.Must(indicators.Add("ROCR100", newROCR100, indicators.MOMENTUM))
	util.Must(indicators.Add("WILLR", newWILLR, indicators.MOMENTUM))
	util.Must(indicators.Add("RSI", newRSI, indicators.MOMENTUM))
	util.Must(indicators.Add("CCI", newCCI, indicators.MOMENTUM))
	util.Must(indicators.Add("MACD", newMACD, indicators.MOMENTUM))
	util.Must(indicators.Add("APO", newAPO, indicators.MOMENTUM))
	util.Must(indicators.Add("PPO", newPPO, indicators.MOMENTUM))
	util.Must(indicators.Add("CMO", newCMO, indicators.MOMENTUM))
	util.Must(indicators.Add("TRIX", newTRIX, indicators.MOMENTUM))
	util.Must(indicators.Add("AROON", newAROON, indicators.MOMENTUM))
	util.Must(indicators.Add("AROONOSC", newAROONOSC, indicators.MOMENTUM))
	util.Must(indicators.Add("STOCH", newSTOCH, indicators.MOMENTUM))
	util.Must(indicators.Add("STOCHF", newSTOCHF, indicators.MOMENTUM))
}

var _ = compute // keep helper available for future indicators
