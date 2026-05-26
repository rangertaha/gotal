// Package signal implements buy/sell/hold signal generators derived from the
// standard indicators. Each emits +100 (buy), -100 (sell), or 0 (hold), the
// same convention TA-Lib uses for candlestick pattern outputs.
package signal

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/indicators"
	"github.com/rangertaha/gotal/internal/indicators/util"
)

// All runs every implemented signal generator against ts.
func All(ts internal.TimeSeries, opts ...internal.ConfigOption) internal.TimeSeries {
	return util.RunGroup(ts, indicators.SIGNAL, opts...)
}

// ---------- threshold-cross helpers ----------

type thresholdSig struct {
	Name        string
	upper       float64
	lower       float64
	underlying  func(ts internal.TimeSeries, opts []internal.ConfigOption) []float64
}

func (i *thresholdSig) Reset() error                          { return nil }
func (i *thresholdSig) Ready() bool                           { return true }
func (i *thresholdSig) Process(t internal.Tick) internal.Tick { return t }

// thresholdCtor builds an indicator that computes the underlying indicator's
// scalar series via the given id and emits a threshold-cross signal.
func thresholdCtor(name, sourceID string, upper, lower float64) func(opts ...internal.ConfigOption) (internal.Indicator, error) {
	return func(opts ...internal.ConfigOption) (internal.Indicator, error) {
		c := util.Cfg(opts...)
		nm := c.GetStr("name", name)
		return &runFn{
			Name: nm,
			run: func(ts internal.TimeSeries) []float64 {
				// Reuse the underlying indicator by calling its Compute.
				if !indicators.Has(sourceID) || indicators.IsStub(sourceID) {
					return make([]float64, ts.Len())
				}
				ind, err := indicators.Get(sourceID)(opts...)
				if err != nil {
					return make([]float64, ts.Len())
				}
				out := ind.Compute(ts)
				vec := out.Fields().Get(c.GetStr("source", sourceLower(sourceID)))
				if vec == nil {
					return make([]float64, ts.Len())
				}
				return util.ThresholdCross(vec.Values(), upper, lower)
			},
		}, nil
	}
}

// ---------- zero-cross helpers ----------

func zeroCrossCtor(name, sourceID string) func(opts ...internal.ConfigOption) (internal.Indicator, error) {
	return func(opts ...internal.ConfigOption) (internal.Indicator, error) {
		c := util.Cfg(opts...)
		nm := c.GetStr("name", name)
		return &runFn{
			Name: nm,
			run: func(ts internal.TimeSeries) []float64 {
				if !indicators.Has(sourceID) || indicators.IsStub(sourceID) {
					return make([]float64, ts.Len())
				}
				ind, err := indicators.Get(sourceID)(opts...)
				if err != nil {
					return make([]float64, ts.Len())
				}
				out := ind.Compute(ts)
				vec := out.Fields().Get(c.GetStr("source", sourceLower(sourceID)))
				if vec == nil {
					return make([]float64, ts.Len())
				}
				return util.ZeroCross(vec.Values())
			},
		}, nil
	}
}

// ---------- line-crossover helpers ----------

// macdCross runs MACD then emits crossover of macd line vs macd_signal line.
type macdCrossSig struct {
	Name string
}

func newMACDCROSS(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &macdCrossSig{Name: c.GetStr("name", "macdcross")}, nil
}
func (i *macdCrossSig) Reset() error                          { return nil }
func (i *macdCrossSig) Ready() bool                           { return true }
func (i *macdCrossSig) Process(t internal.Tick) internal.Tick { return t }
func (i *macdCrossSig) Compute(ts internal.TimeSeries) internal.TimeSeries {
	ind, err := indicators.Get("macd")()
	if err != nil {
		return util.AttachField(ts, i.Name, make([]float64, ts.Len()))
	}
	out := ind.Compute(ts)
	line := out.Fields().Get("macd")
	sig := out.Fields().Get("macd_signal")
	if line == nil || sig == nil {
		return util.AttachField(ts, i.Name, make([]float64, ts.Len()))
	}
	return util.AttachField(ts, i.Name, util.LineCrossover(line.Values(), sig.Values()))
}

// stochCross emits crossover of stoch_k vs stoch_d.
type stochCrossSig struct {
	Name string
}

func newSTOCHCROSS(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &stochCrossSig{Name: c.GetStr("name", "stochcross")}, nil
}
func (i *stochCrossSig) Reset() error                          { return nil }
func (i *stochCrossSig) Ready() bool                           { return true }
func (i *stochCrossSig) Process(t internal.Tick) internal.Tick { return t }
func (i *stochCrossSig) Compute(ts internal.TimeSeries) internal.TimeSeries {
	ind, err := indicators.Get("stoch")()
	if err != nil {
		return util.AttachField(ts, i.Name, make([]float64, ts.Len()))
	}
	out := ind.Compute(ts)
	k := out.Fields().Get("stoch_k")
	d := out.Fields().Get("stoch_d")
	if k == nil || d == nil {
		return util.AttachField(ts, i.Name, make([]float64, ts.Len()))
	}
	return util.AttachField(ts, i.Name, util.LineCrossover(k.Values(), d.Values()))
}

// adxSig: ADX > 25 means "trending"; ADX < 20 means "non-trending". Output
// +100 when entering trending state, -100 when entering non-trending state.
type adxSig struct {
	Name string
}

func newADXSIGNAL(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := util.Cfg(opts...)
	return &adxSig{Name: c.GetStr("name", "adxsignal")}, nil
}
func (i *adxSig) Reset() error                          { return nil }
func (i *adxSig) Ready() bool                           { return true }
func (i *adxSig) Process(t internal.Tick) internal.Tick { return t }
func (i *adxSig) Compute(ts internal.TimeSeries) internal.TimeSeries {
	// ADX is itself a stub today, so this signal also produces zeros.
	return util.AttachField(ts, i.Name, make([]float64, ts.Len()))
}

// ---------- runFn helper struct (wraps a closure as an Indicator) ----------

type runFn struct {
	Name string
	run  func(internal.TimeSeries) []float64
}

func (i *runFn) Reset() error                          { return nil }
func (i *runFn) Ready() bool                           { return true }
func (i *runFn) Process(t internal.Tick) internal.Tick { return t }
func (i *runFn) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.AttachField(ts, i.Name, i.run(ts))
}

func sourceLower(id string) string {
	out := []byte(id)
	for i := range out {
		if out[i] >= 'A' && out[i] <= 'Z' {
			out[i] += 'a' - 'A'
		}
	}
	return string(out)
}

func init() {
	// Threshold-cross signals on bounded oscillators.
	util.Must(indicators.Add("RSISIGNAL", thresholdCtor("rsisignal", "RSI", 70, 30), indicators.SIGNAL))
	util.Must(indicators.Add("MFISIGNAL", thresholdCtor("mfisignal", "MFI", 80, 20), indicators.SIGNAL))
	util.Must(indicators.Add("CCISIGNAL", thresholdCtor("ccisignal", "CCI", 100, -100), indicators.SIGNAL))
	util.Must(indicators.Add("WILLRSIGNAL", thresholdCtor("willrsignal", "WILLR", -20, -80), indicators.SIGNAL))

	// Zero-cross signals on unbounded momentum oscillators.
	util.Must(indicators.Add("MOMSIGNAL", zeroCrossCtor("momsignal", "MOM"), indicators.SIGNAL))
	util.Must(indicators.Add("ROCSIGNAL", zeroCrossCtor("rocsignal", "ROC"), indicators.SIGNAL))
	util.Must(indicators.Add("TRIXSIGNAL", zeroCrossCtor("trixsignal", "TRIX"), indicators.SIGNAL))
	util.Must(indicators.Add("PPOSIGNAL", zeroCrossCtor("pposignal", "PPO"), indicators.SIGNAL))
	util.Must(indicators.Add("CMOSIGNAL", zeroCrossCtor("cmosignal", "CMO"), indicators.SIGNAL))

	// Line-crossover signals.
	util.Must(indicators.Add("MACDSIGNAL_CROSS", newMACDCROSS, indicators.SIGNAL))
	util.Must(indicators.Add("STOCHSIGNAL", newSTOCHCROSS, indicators.SIGNAL))

	// State-style signals.
	util.Must(indicators.Add("ADXSIGNAL", newADXSIGNAL, indicators.SIGNAL))
}
